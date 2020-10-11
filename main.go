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

	opts, switches := sconfig.Convert()

	mqtt.DEBUG = log.New(os.Stdout, "DEBUG", 0)
	mqtt.ERROR = log.New(os.Stdout, "ERROR", 0)

	sub := func(client mqtt.Client) {
		for _, sw := range switches {
			sw.Subscribe(client)
		}
	}

	opts.SetOnConnectHandler(sub)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(60 * time.Second)
	go func() {
		for range ticker.C {
			for _, sw := range switches {
				sw.AnnounceAvailable(client)
			}
		}
	}()

	<-done
	// ticker.Stop()
	log.Print("Server Stopped")

	for _, sw := range switches {
		sw.UnSubscribe(client)
	}

	client.Disconnect(250)

	time.Sleep(1 * time.Second)
}
