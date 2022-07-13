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

	"github.com/W-Floyd/ha-mqtt-iot/config"
	ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/imdario/mergo"
)

//go:generate go run ./helpers/

var debugLog = log.New(os.Stdout, "DEBUG   ", 0)
var errorLog = log.New(os.Stdout, "ERROR   ", 0)
var warnLog = log.New(os.Stdout, "WARN    ", 0)
var criticalLog = log.New(os.Stdout, "CRITICAL", 0)

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

const logPrefix = "[ha-mqtt]  "

func logError(message ...interface{}) {
	if len(message) > 1 {
		for _, mes := range message[:len(message)-2] {
			errorLog.Printf("%v%v\n", logPrefix, mes)
		}
	}
	errorLog.Fatalf("%v%v\n", logPrefix, message[len(message)-1])
}

func logDebug(message ...interface{}) {
	for _, mes := range message {
		debugLog.Printf("%v%v\n", logPrefix, mes)
	}
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
			logError("Error reading "+configFile, err)
		}

		d := json.NewDecoder(strings.NewReader(string(data)))
		d.DisallowUnknownFields()

		// unmarshall it
		err = d.Decode(&tConfig)
		if err != nil {
			logError("Error parsing config", err)
		}

		mergo.Merge(&sconfig, tConfig)

	}

	devices, opts := sconfig.Convert()

	if sconfig.Logging.Debug {
		mqtt.DEBUG = debugLog
	}
	if sconfig.Logging.Warn {
		mqtt.WARN = warnLog
	}
	if sconfig.Logging.Error {
		mqtt.ERROR = errorLog
	}
	if sconfig.Logging.Critical {
		mqtt.CRITICAL = criticalLog
	}

	sub := func(c mqtt.Client) {
		fields := ExternalDevice.MQTTFields{
			Client: &c,
		}
		for _, d := range devices {
			d.SetMQTTFields(fields)
			go d.Subscribe()
		}
	}

	opts.SetOnConnectHandler(sub)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logError(token.Error())
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
			tickers[tickerN] = time.NewTicker(time.Duration(d.GetMQTTFields().UpdateInterval*1000) * time.Millisecond)
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

	<-done
	availableTicker.Stop()
	for _, t := range tickers {
		t.Stop()
	}
	logDebug("Server Stopped")

	for _, d := range devices {
		d.UnSubscribe()
	}

	client.Disconnect(250)

	time.Sleep(1 * time.Second)
}
