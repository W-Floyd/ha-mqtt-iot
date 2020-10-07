package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"os/user"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"./hadiscovery"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("TOPIC: %s\n", msg.Topic())
	log.Printf("MSG: %s\n", msg.Payload())
}

func main() {

	config := hadiscovery.Switch{}
	config.Name = "Dark Mode"
	config.UniqueID = "dark-mode"
	config.Icon = "mdi:theme-light-dark"
	config.Initialize()
	config.CommandFunc = func(message mqtt.Message, connection mqtt.Client) {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		if string(message.Payload()) == "ON" {

			_, err := exec.Command(usr.HomeDir+"/.theme.sh", "set", "dark").Output()
			if err != nil {
				log.Printf("%s", err)
			}

		} else if string(message.Payload()) == "OFF" {

			_, err := exec.Command(usr.HomeDir+"/.theme.sh", "set", "light").Output()
			if err != nil {
				log.Printf("%s", err)
			}

		} else {

			log.Println("Unknown payload: " + string(message.Payload()))

		}
	}
	config.StateFunc = func() string {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		out, err := exec.Command(usr.HomeDir+"/.theme.sh", "get").Output()
		if err != nil {
			log.Printf("%s", err)
		}

		return string(out)

	}

	config2 := hadiscovery.Switch{}
	config2.Name = "Monitor Awake"
	config2.UniqueID = "desktop-monitor-awake"
	config2.Icon = "mdi:monitor"
	config2.Optimistic = true
	config2.Initialize()
	config2.CommandFunc = func(message mqtt.Message, connection mqtt.Client) {

		if string(message.Payload()) == "ON" {

			_, err := exec.Command("xset", "dpms", "force", "on").Output()
			if err != nil {
				log.Printf("%s", err)
			}

		} else if string(message.Payload()) == "OFF" {

			_, err := exec.Command("xset", "dpms", "force", "off").Output()
			if err != nil {
				log.Printf("%s", err)
			}

		} else {

			log.Println("Unknown payload: " + string(message.Payload()))

		}
	}

	mqtt.DEBUG = log.New(os.Stdout, "DEBUG", 0)
	mqtt.ERROR = log.New(os.Stdout, "ERROR", 0)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(Broker)
	opts.SetClientID(hadiscovery.NodeID)
	opts.SetUsername(Username)
	opts.SetPassword(Password)
	opts.SetKeepAlive(30 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetAutoReconnect(true)

	sub := func(client mqtt.Client) {
		config.Subscribe(client)
		config2.Subscribe(client)
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
			config.AnnounceAvailable(client)
			config2.AnnounceAvailable(client)
		}
	}()

	<-done
	// ticker.Stop()
	log.Print("Server Stopped")

	config.UnSubscribe(client)
	config2.UnSubscribe(client)

	client.Disconnect(250)

	time.Sleep(1 * time.Second)
}
