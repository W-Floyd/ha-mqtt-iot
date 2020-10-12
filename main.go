package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"./iotconfig"
)

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

	opts, switches, sensors, _ := sconfig.Convert()

	mqtt.DEBUG = log.New(os.Stdout, "DEBUG", 0)
	mqtt.ERROR = log.New(os.Stdout, "ERROR", 0)

	sub := func(client mqtt.Client) {
		for _, sw := range switches {
			go sw.Subscribe(client)
		}
		for _, se := range sensors {
			go se.Subscribe(client)
		}
	}

	opts.SetOnConnectHandler(sub)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	tickers := make([]*time.Ticker, len(sensors))

	for k, se := range sensors {
		tickers[k] = time.NewTicker(time.Duration(se.UpdateInterval) * time.Second)
		go func(t *time.Ticker) {
			for range t.C {
				go se.UpdateState(client)
			}
		}(tickers[k])
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
		go sw.UnSubscribe(client)
	}
	for _, se := range sensors {
		go se.UnSubscribe(client)
	}

	client.Disconnect(250)

	time.Sleep(1 * time.Second)
}
