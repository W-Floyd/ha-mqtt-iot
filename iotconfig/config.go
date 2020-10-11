package iotconfig

import (
	"log"
	"os/exec"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"../hadiscovery"
)

type Info struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Icon string `json:"icon"`
}

type SwitchHA struct {
	Info         Info     `json:"info"`
	CommandOn    []string `json:"command_on"`
	CommandOff   []string `json:"command_off"`
	CommandState []string `json:"command_state"`
}

type SensorHA struct {
	Info         Info     `json:"info"`
	CommandState []string `json:"command_state"`
}

type BinarySensorsHA struct {
	Info         Info     `json:"info"`
	CommandState []string `json:"command_state"`
}

type Config struct {
	MQTT struct {
		Broker   string `json:"broker"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"mqtt"`
	Switches      []SwitchHA        `json:"switches"`
	Sensors       []SensorHA        `json:"sensors"`
	BinarySensors []BinarySensorsHA `json:"binary_sensors"`
}

func (sw SwitchHA) constructCommandFunc() (f func(message mqtt.Message, connection mqtt.Client)) {
	var err error
	return func(message mqtt.Message, connection mqtt.Client) {

		if len(sw.CommandOn) > 0 {
			if string(message.Payload()) == "ON" {

				if len(sw.CommandOn) > 1 {
					_, err = exec.Command(sw.CommandOn[0], sw.CommandOn[1:]...).Output()
				} else {
					_, err = exec.Command(sw.CommandOn[0]).Output()
				}
				if err != nil {
					log.Printf("%s", err)
				}
			}
		}
		if len(sw.CommandOff) > 0 {
			if string(message.Payload()) == "OFF" {

				if len(sw.CommandOff) > 1 {
					_, err = exec.Command(sw.CommandOff[0], sw.CommandOff[1:]...).Output()
				} else {
					_, err = exec.Command(sw.CommandOff[0]).Output()
				}
				if err != nil {
					log.Printf("%s", err)
				}
			}
		}
		if string(message.Payload()) != "OFF" && string(message.Payload()) != "ON" {

			log.Println("Unknown payload: " + string(message.Payload()))

		}
	}
}

func (sw SwitchHA) constructStateFunc() (f func() string) {
	var err error
	return func() string {
		var out []byte
		if len(sw.CommandState) > 1 {
			out, err = exec.Command(sw.CommandState[0], sw.CommandState[1:]...).Output()
		} else {

			out, err = exec.Command(sw.CommandState[0]).Output()
		}
		if err != nil {
			log.Printf("%s", err)
		}
		return string(out)
	}
}

// Convert takes a config and turns it into a format that can be used for HA.
func (sconfig Config) Convert() (opts *mqtt.ClientOptions, switches []hadiscovery.Switch) {
	opts = mqtt.NewClientOptions()
	opts.AddBroker(sconfig.MQTT.Broker)
	opts.SetClientID(hadiscovery.NodeID)
	opts.SetUsername(sconfig.MQTT.Username)
	opts.SetPassword(sconfig.MQTT.Password)
	opts.SetKeepAlive(30 * time.Second)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("TOPIC: %s\n", msg.Topic())
		log.Printf("MSG: %s\n", msg.Payload())
	})
	opts.SetPingTimeout(1 * time.Second)
	opts.SetAutoReconnect(true)
	for _, sw := range sconfig.Switches {
		nsw := hadiscovery.Switch{}
		nsw.Name = sw.Info.Name
		nsw.UniqueID = sw.Info.ID
		nsw.Icon = sw.Info.Icon
		nsw.Initialize()

		if len(sw.CommandOn) > 0 || len(sw.CommandOff) > 0 {
			nsw.CommandFunc = sw.constructCommandFunc()
		}
		if len(sw.CommandState) > 0 {

			nsw.StateFunc = sw.constructStateFunc()
		}

		switches = append(switches, nsw)
	}
	return
}
