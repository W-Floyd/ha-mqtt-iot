package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/imdario/mergo"

	"github.com/W-Floyd/ha-mqtt-iot/hadiscovery"
	"github.com/W-Floyd/ha-mqtt-iot/iotconfig"
)

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

	var sconfig iotconfig.Config

	for _, configFile := range configFiles {
		var tConfig iotconfig.Config

		// read file
		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			logError("Error reading "+configFile, err)
		}

		// unmarshall it
		err = json.Unmarshal(data, &tConfig)
		if err != nil {
			logError("Error parsing config", err)
		}

		mergo.Merge(&sconfig, tConfig)

	}

	opts, lights := sconfig.Convert()

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
		for _, li := range lights {
			li.MQTT.Client = &c
			go li.Subscribe()
		}
	}

	opts.SetOnConnectHandler(sub)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logError(token.Error())
	}

	updatingLights := 0

	for _, li := range lights {
		if !almostEqual(li.MQTT.UpdateInterval, 0) {
			updatingLights++
		}
	}

	tickers := make([]*time.Ticker, updatingLights)

	tickerN := 0

	for _, li := range lights {
		if !almostEqual(li.MQTT.UpdateInterval, 0) {
			tickers[tickerN] = time.NewTicker(time.Duration(li.MQTT.UpdateInterval*1000) * time.Millisecond)
			go func(t *time.Ticker, lig hadiscovery.Light) {
				for range t.C {
					go lig.UpdateState()
				}
			}(tickers[tickerN], li)
			tickerN++
		}
	}

	availableTicker := time.NewTicker(60 * time.Second)
	go func() {
		for range availableTicker.C {
			for _, li := range lights {
				go li.AnnounceAvailable()
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

	for _, li := range lights {
		li.UnSubscribe()
	}

	client.Disconnect(250)

	time.Sleep(1 * time.Second)
}
