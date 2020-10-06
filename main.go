package main

import (
	"fmt"
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
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(Broker)
	opts.SetClientID(hadiscovery.NodeID)
	opts.SetUsername(Username)
	opts.SetPassword(Password)
	opts.SetKeepAlive(30 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	hadiscovery.Connection = mqtt.NewClient(opts)
	if token := hadiscovery.Connection.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	config := hadiscovery.Switch{}
	config.Name = "dark-mode"
	config.UniqueID = "dark-mode"
	config.Initialize()
	fmt.Println(config)
	config.CommandFunc = func(message mqtt.Message, connection mqtt.Client) {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		if string(message.Payload()) == "ON" {

			out, err := exec.Command(usr.HomeDir+"/.theme.sh", "set", "dark").Output()
			if err != nil {
				fmt.Printf("%s", err)
			}
			output := string(out[:])
			fmt.Println(output)

		} else if string(message.Payload()) == "OFF" {

			out, err := exec.Command(usr.HomeDir+"/.theme.sh", "set", "light").Output()
			if err != nil {
				fmt.Printf("%s", err)
			}
			output := string(out[:])
			fmt.Println(output)

		} else {

			fmt.Println("Unknown payload: " + string(message.Payload()))

		}
	}
	config.StateFunc = func() string {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		out, err := exec.Command(usr.HomeDir+"/.theme.sh", "get").Output()
		if err != nil {
			fmt.Printf("%s", err)
		}

		return string(out)

	}
	config.Subscribe()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			config.UpdateState()
		}
	}()

	<-done
	ticker.Stop()
	log.Print("Server Stopped")

	config.UnSubscribe()

	hadiscovery.Connection.Disconnect(250)

	time.Sleep(1 * time.Second)
}
