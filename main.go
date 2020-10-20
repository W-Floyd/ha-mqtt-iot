package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"./hadiscovery"
	"./iotconfig"
)

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func main() {

	// read file
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Print(err)
	}

	// json data
	var sconfig iotconfig.Config

	// unmarshall it
	err = json.Unmarshal(data, &sconfig)
	if err != nil {
		log.Println("error:", err)
	}

	opts, switches, sensors, binarySensors := sconfig.Convert()

	mqtt.DEBUG = log.New(os.Stdout, "DEBUG", 0)
	mqtt.ERROR = log.New(os.Stdout, "ERROR", 0)

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
	}

	opts.SetOnConnectHandler(sub)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	updatingSwitches := 0

	for _, sw := range switches {
		if !almostEqual(sw.UpdateInterval, 0) {
			updatingSwitches++
		}
	}

	tickers := make([]*time.Ticker, len(sensors)+len(binarySensors)+updatingSwitches)

	tickerN := 0
	for _, se := range sensors {
		tickers[tickerN] = time.NewTicker(time.Duration(se.UpdateInterval) * time.Second)
		go func(t *time.Ticker, sen hadiscovery.Sensor) {
			for range t.C {
				go sen.UpdateState(client)
			}
		}(tickers[tickerN], se)
		tickerN++
	}
	for _, bse := range binarySensors {
		tickers[tickerN] = time.NewTicker(time.Duration(bse.UpdateInterval) * time.Second)
		go func(t *time.Ticker, bsen hadiscovery.BinarySensor) {
			for range t.C {
				go bsen.UpdateState(client)
			}
		}(tickers[tickerN], bse)
		tickerN++
	}
	for _, sw := range switches {
		if !almostEqual(sw.UpdateInterval, 0) {
			tickers[tickerN] = time.NewTicker(time.Duration(sw.UpdateInterval) * time.Second)
			go func(t *time.Ticker, swi hadiscovery.Switch) {
				for range t.C {
					go swi.UpdateState(client)
				}
			}(tickers[tickerN], sw)
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
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	availableTicker.Stop()
	for _, t := range tickers {
		t.Stop()
	}
	log.Print("Server Stopped")

	for _, sw := range switches {
		sw.UnSubscribe(client)
	}
	for _, se := range sensors {
		se.UnSubscribe(client)
	}
	for _, bse := range binarySensors {
		bse.UnSubscribe(client)
	}

	client.Disconnect(250)

	time.Sleep(1 * time.Second)
}
