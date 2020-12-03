package iotconfig

import (
	"log"
	"os/exec"

	"../hadiscovery"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type SwitchHA struct {
	Info            InfoIcon `json:"info"`
	CommandOn       []string `json:"command_on"`
	CommandOff      []string `json:"command_off"`
	CommandState    []string `json:"command_state"`
	UpdateInterval  float64  `json:"update_interval"`
	ForceUpdateMQTT bool     `json:"force_update"`
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

func (sw SwitchHA) Translate() hadiscovery.Switch {
	nsw := hadiscovery.Switch{}
	nsw.UpdateInterval = sw.UpdateInterval
	nsw.ForceUpdateMQTT = sw.ForceUpdateMQTT
	nsw.Name = sw.Info.Name
	nsw.UniqueID = sw.Info.ID + "_" + hadiscovery.NodeID
	if sw.Info.Icon != "" {
		nsw.Icon = sw.Info.Icon
	}

	if len(sw.CommandOn) > 0 || len(sw.CommandOff) > 0 {
		nsw.CommandFunc = sw.constructCommandFunc()
	} else {
		log.Fatalln("Missing command func, needed for switches!")
	}
	if len(sw.CommandState) > 0 {
		nsw.StateFunc = sw.constructStateFunc()
	}

	nsw.Initialize()

	return nsw
}
