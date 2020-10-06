package hadiscovery

import (
	"encoding/json"
	"log"
	"os"

	"github.com/denisbrodbeck/machineid"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const DiscoveryPrefix = "homeassistant"
const NodeID = "linux-iot-ha"
const InstanceName = "Linux IOT HA"

/////////////////// Components of config

type device struct {
	Connections  map[string]string `json:"connections,omitempty"`
	Identifiers  []string          `json:"identifiers,omitempty"`
	Manufacturer string            `json:"manufacturer,omitempty"`
	Model        string            `json:"model,omitempty"`
	Name         string            `json:"name,omitempty"`
	SWVersion    string            `json:"sw_version,omitempty"`
	ViaDevice    string            `json:"via_device,omitempty"`
}

///////////////////

// // BinarySensor matches a binary_sensor
// type BinarySensor struct {
// }

// // Sensor matches a sensor
// type Sensor struct {
// 	AvailabilityTopic      string `json:"availability_topic,omitempty"`
// 	Device                 device `json:"device,omitempty"`
// 	DeviceClass            string `json:"device_class,omitempty"`
// 	ExpireAfter            int    `json:"expire_after,omitempty"`
// 	ForceUpdate            bool   `json:"force_update,omitempty"`
// 	Icon                   string `json:"icon,omitempty"`
// 	JSONAttributesTemplate string `json:"json_attributes_template,omitempty"`
// 	JSONAttributesTopic    string `json:"json_attributes_topic,omitempty"`
// 	Name                   string `json:"name,omitempty"`
// 	PayloadAvailable       string `json:"payload_available,omitempty"`
// 	PayloadNotAvailable    string `json:"payload_not_available,omitempty"`
// 	QOS                    int    `json:"qos,omitempty"`
// 	StateOff               string `json:"state_off,omitempty"`
// 	StateOn                string `json:"state_on,omitempty"`
// 	StateTopic             string `json:"state_topic"`
// 	UniqueID               string `json:"unique_id,omitempty"`
// 	UnitOfMeasurement      string `json:"unit_of_measurement,omitempty"`
// 	ValueTemplate          string `json:"value_template,omitempty"`
// }

// Switch matches a switch
type Switch struct {
	AvailabilityTopic      string `json:"avty_t,omitempty"`
	CommandTopic           string `json:"cmd_t"`
	Device                 device `json:"dev,omitempty"`
	Icon                   string `json:"ic,omitempty"`
	JSONAttributesTemplate string `json:"json_attr_tpl,omitempty"`
	JSONAttributesTopic    string `json:"json_attr_t,omitempty"`
	Name                   string `json:"name,omitempty"`
	Optimistic             bool   `json:"opt,omitempty"`
	PayloadAvailable       string `json:"pl_avail,omitempty"`
	PayloadNotAvailable    string `json:"pl_not_avail,omitempty"`
	PayloadOff             string `json:"pl_off,omitempty"`
	PayloadOn              string `json:"pl_on,omitempty"`
	Platform               string `json:"platform"`
	QOS                    int    `json:"qos,omitempty"`
	Retain                 bool   `json:"ret,omitempty"`
	StateOff               string `json:"stat_off,omitempty"`
	StateOn                string `json:"stat_on,omitempty"`
	StateTopic             string `json:"stat_t,omitempty"`
	UniqueID               string `json:"uniq_id,omitempty"`
	ValueTemplate          string `json:"val_tpl,omitempty"`

	CommandFunc func(mqtt.Message, mqtt.Client) `json:"-"`
	StateFunc   func() string                   `json:"-"`

	messageHandler mqtt.MessageHandler `json:"-"`
}

///////////////////

func getDevice() (d device) {

	id, err := machineid.ProtectedID(NodeID)
	if err != nil {
		log.Fatal(err)
	}

	d.Identifiers = []string{id}
	d.Manufacturer = "William Floyd"
	d.Model = NodeID
	d.Name = InstanceName
	d.SWVersion = "0.0.1"

	return
}

///////////////////

var topicStore = make(map[string]*func(mqtt.Message, mqtt.Client))

var Connection mqtt.Client

///////////////////

func (device Switch) GetTopicPrefix() string {
	return NodeID + "/switch/" + device.UniqueID + "/"
}

func (device Switch) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/switch/" + NodeID + "/" + device.UniqueID + "/" + "config"
}

func (device Switch) GetCommandTopic() string {
	return device.GetTopicPrefix() + "command"
}

func (device Switch) GetStateTopic() string {
	return device.GetTopicPrefix() + "state"
}

func (device Switch) GetAvailabilityTopic() string {
	return device.GetTopicPrefix() + "availability"
}

// Initialize sets topics as needed on a Switch
func (device *Switch) Initialize() {
	device.CommandTopic = device.GetCommandTopic()
	device.StateTopic = device.GetStateTopic()
	device.AvailabilityTopic = device.GetAvailabilityTopic()
	device.Device = getDevice()
	device.Retain = true
	device.Platform = "mqtt"

	topicStore[device.CommandTopic] = &device.CommandFunc

	device.messageHandler = func(client mqtt.Client, msg mqtt.Message) {

		topicFound := false

		for topic, f := range topicStore {
			if msg.Topic() == topic {
				topicFound = true
				(*f)(msg, Connection)
				device.UpdateState()
			}
		}

		if !topicFound {
			log.Println("Unknown Message on topic " + msg.Topic())
			log.Println(msg.Payload())
		}
	}

}

func (device Switch) UpdateState() {
	if device.StateFunc != nil {
		token := Connection.Publish(device.GetStateTopic(), 0, true, device.StateFunc())
		token.Wait()
	} else {
		log.Println("No statefunc")
	}
}

func (device Switch) Subscribe() {

	message, err := json.Marshal(device)
	if err != nil {
		log.Fatal(err)
	}

	token := Connection.Publish(device.GetDiscoveryTopic(), 0, true, message)
	token.Wait()

	device.UpdateState()

	if token := Connection.Subscribe(device.GetCommandTopic(), 0, device.messageHandler); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		os.Exit(1)
	}

	device.AnnounceAvailable()

}

func (device Switch) UnSubscribe() {
	token := Connection.Publish(device.GetAvailabilityTopic(), 0, false, "offline")
	token.Wait()

	if token := Connection.Unsubscribe(device.GetCommandTopic()); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		os.Exit(1)
	}
}

func (device Switch) AnnounceAvailable() {
	token := Connection.Publish(device.GetAvailabilityTopic(), 0, false, "online")
	token.Wait()
}

///////////////////
