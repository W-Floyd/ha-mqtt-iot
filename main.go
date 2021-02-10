package main

import (
	"encoding/json"
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

	configFiles := [...]string{"./config.json", "./secrets.json"}

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

	opts, switches, sensors, binarySensors, lights := sconfig.Convert()

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

	sub := func(client mqtt.Client) {
		for _, sw := range switches {
			go sw.Subscribe(client)
		}
		for _, se := range sensors {
			go se.Subscribe(client)
		}
		for _, bse := range binarySensors {
			go bse.Subscribe(client)
		}
		for _, li := range lights {
			go li.Subscribe(client)
		}
	}

	opts.SetOnConnectHandler(sub)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logError(token.Error())
	}

	updatingSwitches := 0

	for _, sw := range switches {
		if !almostEqual(sw.UpdateInterval, 0) {
			updatingSwitches++
		}
	}

	updatingLights := 0

	for _, li := range lights {
		if !almostEqual(li.UpdateInterval, 0) {
			updatingLights++
		}
	}

	tickers := make([]*time.Ticker, len(sensors)+len(binarySensors)+updatingSwitches+updatingLights)

	tickerN := 0
	for _, se := range sensors {
		tickers[tickerN] = time.NewTicker(time.Duration(se.UpdateInterval*1000) * time.Millisecond)
		go func(t *time.Ticker, sen hadiscovery.Sensor) {
			for range t.C {
				go sen.UpdateState(client)
			}
		}(tickers[tickerN], se)
		tickerN++
	}
	for _, bse := range binarySensors {
		tickers[tickerN] = time.NewTicker(time.Duration(bse.UpdateInterval*1000) * time.Millisecond)
		go func(t *time.Ticker, bsen hadiscovery.BinarySensor) {
			for range t.C {
				go bsen.UpdateState(client)
			}
		}(tickers[tickerN], bse)
		tickerN++
	}
	for _, sw := range switches {
		if !almostEqual(sw.UpdateInterval, 0) {
			tickers[tickerN] = time.NewTicker(time.Duration(sw.UpdateInterval*1000) * time.Millisecond)
			go func(t *time.Ticker, swi hadiscovery.Switch) {
				for range t.C {
					go swi.UpdateState(client)
				}
			}(tickers[tickerN], sw)
			tickerN++
		}
	}
	for _, li := range lights {
		if !almostEqual(li.UpdateInterval, 0) {
			tickers[tickerN] = time.NewTicker(time.Duration(li.UpdateInterval*1000) * time.Millisecond)
			go func(t *time.Ticker, lig hadiscovery.Light) {
				for range t.C {
					go lig.UpdateState(client)
				}
			}(tickers[tickerN], li)
			tickerN++
		}
	}

	availableTicker := time.NewTicker(60 * time.Second)
	go func() {
		for range availableTicker.C {
			for _, sw := range switches {
				go sw.AnnounceAvailable(client)
			}
			for _, se := range sensors {
				go se.AnnounceAvailable(client)
			}
			for _, bse := range binarySensors {
				go bse.AnnounceAvailable(client)
			}
			for _, li := range lights {
				go li.AnnounceAvailable(client)
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

	for _, sw := range switches {
		sw.UnSubscribe(client)
	}
	for _, se := range sensors {
		se.UnSubscribe(client)
	}
	for _, bse := range binarySensors {
		bse.UnSubscribe(client)
	}
	for _, li := range lights {
		li.UnSubscribe(client)
	}

	client.Disconnect(250)

	time.Sleep(1 * time.Second)
}
