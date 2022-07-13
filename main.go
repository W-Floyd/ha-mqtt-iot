package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"math"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/W-Floyd/ha-mqtt-iot/common"
	"github.com/W-Floyd/ha-mqtt-iot/config"
	ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/imdario/mergo"
)

//go:generate go run ./helpers/

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func main() {
	configFile := flag.String("config", "./config.json", "path to config file")
	secretsFile := flag.String("secrets", "./secrets.json", "path to secrets file")
	flag.Parse()

	configFiles := [...]string{*configFile, *secretsFile}

	var sconfig config.Config

	for _, configFile := range configFiles {
		var tConfig config.Config

		// read file
		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			common.LogError("Error reading "+configFile, err)
		}

		d := json.NewDecoder(strings.NewReader(string(data)))
		d.DisallowUnknownFields()

		// unmarshall it
		err = d.Decode(&tConfig)
		if err != nil {
			common.LogError("Error parsing config", err)
		}

		mergo.Merge(&sconfig, tConfig)

	}

	devices, opts := sconfig.Convert()

	if sconfig.Logging.Debug {
		mqtt.DEBUG = common.DebugLog
	}
	if sconfig.Logging.Warn {
		mqtt.WARN = common.WarnLog
	}
	if sconfig.Logging.Error {
		mqtt.ERROR = common.ErrorLog
	}
	if sconfig.Logging.Critical {
		mqtt.CRITICAL = common.CriticalLog
	}

	sub := func(c mqtt.Client) {
		fields := ExternalDevice.MQTTFields{
			Client: &c,
		}
		for _, d := range devices {
			d.SetMQTTFields(fields)
			// spew.Dump(d.GetMQTTFields())
			// spew.Dump(fields)
			go d.Subscribe()
		}
	}

	opts.SetOnConnectHandler(sub)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		common.LogError(token.Error())
	}

	updatingDevices := 0

	for _, d := range devices {
		if !almostEqual(d.GetMQTTFields().UpdateInterval, 0) {
			updatingDevices++
		}
	}

	tickers := make([]*time.Ticker, updatingDevices)

	tickerN := 0

	for _, d := range devices {
		if !almostEqual(d.GetMQTTFields().UpdateInterval, 0) {
			tickers[tickerN] = time.NewTicker(time.Duration(d.GetMQTTFields().UpdateInterval) * time.Second)
			go func(t *time.Ticker, device ExternalDevice.Device) {
				for range t.C {
					go device.UpdateState()
				}
			}(tickers[tickerN], d)
			tickerN++
		}
	}

	availableTicker := time.NewTicker(60 * time.Second)
	go func() {
		for range availableTicker.C {
			for _, d := range devices {
				go d.Subscribe()
			}
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	common.LogDebug("Everything is set up")

	<-done
	availableTicker.Stop()
	for _, t := range tickers {
		t.Stop()
	}
	common.LogDebug("Server Stopped")

	for _, d := range devices {
		d.UnSubscribe()
		common.LogDebug(d.GetRawId() + " Unsubscribed")
	}
	common.LogDebug("All Devices Unsubscribed")

	client.Disconnect(250)

	time.Sleep(1 * time.Second)
}
