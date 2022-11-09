package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/W-Floyd/ha-mqtt-iot/common"
	"github.com/W-Floyd/ha-mqtt-iot/config"
	ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
	"github.com/denisbrodbeck/machineid"
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

	id, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}

	common.MachineID = id

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

	if sconfig.Logging.Debug && sconfig.Logging.Mqtt {
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

	common.LogState.Debug = sconfig.Logging.Debug
	common.LogState.Warn = sconfig.Logging.Warn
	common.LogState.Error = sconfig.Logging.Error
	common.LogState.Critical = sconfig.Logging.Critical

	opts.SetOnConnectHandler(
		func(c mqtt.Client) {
			for _, d := range devices {
				common.LogDebug("Connecting " + d.GetRawId() + "." + d.GetUniqueId())
				go d.Subscribe()
			}
		},
	)

        opts.SetPingTimeout(10 * time.Second)
        opts.SetKeepAlive(10 * time.Second)
        opts.SetMaxReconnectInterval(10 * time.Second)
        opts.SetAutoReconnect(true)

	client := mqtt.NewClient(opts)

	for _, d := range devices {
		f := d.GetMQTTFields()
		f.Client = &client
		d.SetMQTTFields(f)
	}

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		common.LogError(token.Error())
	}

	updatingDevices := 0

	for _, d := range devices {
		if d.GetMQTTFields().UpdateInterval != nil && !almostEqual(*d.GetMQTTFields().UpdateInterval, 0) {
			updatingDevices++
		}
	}

	tickers := make([]*time.Ticker, updatingDevices)

	tickerN := 0

	for _, d := range devices {
		if d.GetMQTTFields().UpdateInterval != nil && !almostEqual(*d.GetMQTTFields().UpdateInterval, 0) {
			common.LogDebug("Starting ticker for " + d.GetRawId())
			tickers[tickerN] = time.NewTicker(time.Duration(*d.GetMQTTFields().UpdateInterval) * time.Second)
			go func(t *time.Ticker, device ExternalDevice.Device) {
				for range t.C {
					go device.UpdateState()
				}
			}(tickers[tickerN], d)
			tickerN++
		}
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	common.LogDebug("Everything is set up")
	<-done

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
