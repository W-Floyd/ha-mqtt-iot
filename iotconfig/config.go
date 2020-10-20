package iotconfig

import (
	"log"
	"os/exec"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"../hadiscovery"
)

type InfoIcon struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Icon string `json:"icon"`
}

type InfoClass struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	DeviceClass string `json:"device_class"`
}

type SwitchHA struct {
	Info            InfoIcon `json:"info"`
	CommandOn       []string `json:"command_on"`
	CommandOff      []string `json:"command_off"`
	CommandState    []string `json:"command_state"`
	UpdateInterval  float64  `json:"update_interval"`
	ForceUpdateMQTT bool     `json:"force_update"`
}

type SensorHA struct {
	Info            InfoIcon `json:"info"`
	CommandState    []string `json:"command_state"`
	UpdateInterval  float64  `json:"update_interval"`
	ForceUpdateMQTT bool     `json:"force_update"`
}

type BinarySensorsHA struct {
	Info            InfoClass `json:"info"`
	CommandState    []string  `json:"command_state"`
	UpdateInterval  float64   `json:"update_interval"`
	ForceUpdateMQTT bool      `json:"force_update"`
}

type Config struct {
	MQTT struct {
		Broker       string `json:"broker"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		NodeID       string `json:"node_id"`
		InstanceName string `json:"instance_name"`
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

func (sw SensorHA) constructStateFunc() (f func() string) {
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
func (sw BinarySensorsHA) constructStateFunc() (f func() string) {
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
func (sconfig Config) Convert() (opts *mqtt.ClientOptions, switches []hadiscovery.Switch, sensors []hadiscovery.Sensor, binarySensors []hadiscovery.BinarySensor) {
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
	if sconfig.MQTT.NodeID != "" {
		hadiscovery.NodeID = sconfig.MQTT.NodeID
	}
	if sconfig.MQTT.InstanceName != "" {
		hadiscovery.InstanceName = sconfig.MQTT.InstanceName
	}
	for _, sw := range sconfig.Switches {
		nsw := hadiscovery.Switch{}
		nsw.UpdateInterval = sw.UpdateInterval
		nsw.ForceUpdateMQTT = sw.ForceUpdateMQTT
		nsw.Name = sw.Info.Name
		nsw.UniqueID = sw.Info.ID
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

		switches = append(switches, nsw)
	}
	for _, se := range sconfig.Sensors {
		nse := hadiscovery.Sensor{}
		nse.UpdateInterval = se.UpdateInterval
		nse.ExpireAfter = int(nse.UpdateInterval + 1)
		nse.ForceUpdateMQTT = se.ForceUpdateMQTT
		if !se.ForceUpdateMQTT {
			nse.ExpireAfter = 0
		}
		nse.Name = se.Info.Name
		nse.UniqueID = se.Info.ID
		if se.Info.Icon != "" {
			nse.Icon = se.Info.Icon
		}

		if len(se.CommandState) > 0 {
			nse.StateFunc = se.constructStateFunc()
		} else {
			log.Fatalln("Missing state func, needed for sensors!")
		}

		nse.Initialize()

		sensors = append(sensors, nse)
	}
	for _, bse := range sconfig.BinarySensors {
		nse := hadiscovery.BinarySensor{}
		nse.UpdateInterval = bse.UpdateInterval
		nse.ExpireAfter = int(nse.UpdateInterval + 1)
		nse.ForceUpdateMQTT = bse.ForceUpdateMQTT
		if !bse.ForceUpdateMQTT {
			nse.ExpireAfter = 0
		}
		nse.Name = bse.Info.Name
		nse.UniqueID = bse.Info.ID
		if bse.Info.DeviceClass != "" {
			nse.DeviceClass = bse.Info.DeviceClass
		}

		if len(bse.CommandState) > 0 {
			nse.StateFunc = bse.constructStateFunc()
		} else {
			log.Fatalln("Missing state func, needed for binary sensors!")
		}

		nse.Initialize()

		binarySensors = append(binarySensors, nse)
	}
	return
}
