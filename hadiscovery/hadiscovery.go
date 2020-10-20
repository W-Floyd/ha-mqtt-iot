package hadiscovery

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/denisbrodbeck/machineid"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const DiscoveryPrefix = "homeassistant"
const NodeID = "ha-mqtt-iot"
const InstanceName = "Homeassistant MQTT IOT"
const SWVersion = "0.0.1"

/////////////////// Components of config

type device struct {
	Connections  map[string]string `json:"cns,omitempty"`
	Identifiers  []string          `json:"ids,omitempty"`
	Manufacturer string            `json:"mf,omitempty"`
	Model        string            `json:"mdl,omitempty"`
	Name         string            `json:"name,omitempty"`
	SWVersion    string            `json:"sw,omitempty"`
	ViaDevice    string            `json:"via_device,omitempty"`
}

///////////////////

// BinarySensor matches a binary_sensor
type BinarySensor struct {
	AvailabilityTopic      string `json:"availability_topic,omitempty"`
	Device                 device `json:"device,omitempty"`
	DeviceClass            string `json:"device_class,omitempty"`
	ExpireAfter            int    `json:"expire_after,omitempty"`
	ForceUpdate            bool   `json:"force_update,omitempty"`
	JSONAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic    string `json:"json_attributes_topic,omitempty"`
	Name                   string `json:"name,omitempty"`
	OffDelay               int    `json:"off_delay,omitempty"`
	PayloadAvailable       string `json:"payload_available,omitempty"`
	PayloadNotAvailable    string `json:"payload_not_available,omitempty"`
	PayloadOff             string `json:"pl_off,omitempty"`
	PayloadOn              string `json:"pl_on,omitempty"`
	QOS                    int    `json:"qos,omitempty"`
	StateTopic             string `json:"state_topic"`
	UniqueID               string `json:"unique_id,omitempty"`
	ValueTemplate          string `json:"value_template,omitempty"`

	StateFunc func() string `json:"-"`

	UpdateInterval  float64 `json:"-"`
	ForceUpdateMQTT bool    `json:"-"`
}

// Sensor matches a sensor
type Sensor struct {
	AvailabilityTopic      string `json:"availability_topic,omitempty"`
	Device                 device `json:"device,omitempty"`
	DeviceClass            string `json:"device_class,omitempty"`
	ExpireAfter            int    `json:"expire_after,omitempty"`
	ForceUpdate            bool   `json:"force_update,omitempty"`
	Icon                   string `json:"icon,omitempty"`
	JSONAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic    string `json:"json_attributes_topic,omitempty"`
	Name                   string `json:"name,omitempty"`
	PayloadAvailable       string `json:"payload_available,omitempty"`
	PayloadNotAvailable    string `json:"payload_not_available,omitempty"`
	QOS                    int    `json:"qos,omitempty"`
	StateTopic             string `json:"state_topic"`
	UniqueID               string `json:"unique_id,omitempty"`
	UnitOfMeasurement      string `json:"unit_of_measurement,omitempty"`
	ValueTemplate          string `json:"value_template,omitempty"`

	StateFunc func() string `json:"-"`

	UpdateInterval  float64 `json:"-"`
	ForceUpdateMQTT bool    `json:"-"`
}

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
	QOS                    int    `json:"qos,omitempty"`
	Retain                 bool   `json:"ret,omitempty"`
	StateOff               string `json:"stat_off,omitempty"`
	StateOn                string `json:"stat_on,omitempty"`
	StateTopic             string `json:"stat_t,omitempty"`
	UniqueID               string `json:"uniq_id,omitempty"`
	ValueTemplate          string `json:"val_tpl,omitempty"`

	CommandFunc func(mqtt.Message, mqtt.Client) `json:"-"`
	StateFunc   func() string                   `json:"-"`

	UpdateInterval  float64 `json:"-"`
	ForceUpdateMQTT bool    `json:"-"`

	messageHandler mqtt.MessageHandler `json:"-"`
}

///////////////////

func getDevice() (d device) {

	id, err := machineid.ProtectedID(NodeID + SWVersion)
	if err != nil {
		log.Fatal(err)
	}

	d.Identifiers = []string{id}
	d.Manufacturer = "William Floyd"
	d.Model = NodeID
	d.Name = InstanceName
	d.SWVersion = SWVersion

	return
}

///////////////////

var topicStore = make(map[string]*func(mqtt.Message, mqtt.Client))
var binarySensorStateStore = make(map[string]string)
var sensorStateStore = make(map[string]string)
var switchStateStore = make(map[string]string)

///////////////////

func (device Sensor) GetTopicPrefix() string {
	return NodeID + "/sensor/" + device.UniqueID + "/"
}
func (device Switch) GetTopicPrefix() string {
	return NodeID + "/switch/" + device.UniqueID + "/"
}
func (device BinarySensor) GetTopicPrefix() string {
	return NodeID + "/binary_sensor/" + device.UniqueID + "/"
}

func (device Sensor) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/sensor/" + NodeID + "/" + device.UniqueID + "/" + "config"
}
func (device Switch) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/switch/" + NodeID + "/" + device.UniqueID + "/" + "config"
}
func (device BinarySensor) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/binary_sensor/" + NodeID + "/" + device.UniqueID + "/" + "config"
}

func (device Switch) GetCommandTopic() string {
	return device.GetTopicPrefix() + "command"
}

func (device Sensor) GetStateTopic() string {
	return device.GetTopicPrefix() + "state"
}
func (device Switch) GetStateTopic() string {
	return device.GetTopicPrefix() + "state"
}
func (device BinarySensor) GetStateTopic() string {
	return device.GetTopicPrefix() + "state"
}

func (device Sensor) GetAvailabilityTopic() string {
	return device.GetTopicPrefix() + "availability"
}
func (device Switch) GetAvailabilityTopic() string {
	return device.GetTopicPrefix() + "availability"
}
func (device BinarySensor) GetAvailabilityTopic() string {
	return device.GetTopicPrefix() + "availability"
}

// Initialize sets topics as needed on a Switch
func (device *Switch) Initialize() {
	device.CommandTopic = device.GetCommandTopic()
	if device.StateFunc != nil {
		device.StateTopic = device.GetStateTopic()
	}
	device.AvailabilityTopic = device.GetAvailabilityTopic()
	device.Device = getDevice()
	device.Retain = false

	topicStore[device.CommandTopic] = &device.CommandFunc

	device.messageHandler = func(client mqtt.Client, msg mqtt.Message) {

		topicFound := false

		for topic, f := range topicStore {
			if msg.Topic() == topic {
				topicFound = true
				(*f)(msg, client)
				device.UpdateState(client)
			}
		}

		if !topicFound {
			log.Println("Unknown Message on topic " + msg.Topic())
			log.Println(msg.Payload())
		}
	}

}

// Initialize sets topics as needed on a Sensor
func (device *Sensor) Initialize() {
	device.StateTopic = device.GetStateTopic()
	device.AvailabilityTopic = device.GetAvailabilityTopic()
	device.Device = getDevice()
}

// Initialize sets topics as needed on a Sensor
func (device *BinarySensor) Initialize() {
	device.StateTopic = device.GetStateTopic()
	device.AvailabilityTopic = device.GetAvailabilityTopic()
	device.Device = getDevice()
}

func (device Switch) UpdateState(client mqtt.Client) {
	if device.StateFunc != nil {
		state := device.StateFunc()
		if state != switchStateStore[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.GetStateTopic(), 0, false, device.StateFunc())
			switchStateStore[device.UniqueID] = state
			token.Wait()
		}
	} else {
		log.Println("No statefunc for " + device.UniqueID + strconv.FormatFloat(float64(device.UpdateInterval), 'g', 1, 64))
	}
}
func (device Sensor) UpdateState(client mqtt.Client) {
	if device.StateFunc != nil {
		state := device.StateFunc()
		if state != sensorStateStore[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.GetStateTopic(), 0, false, device.StateFunc())
			token.Wait()
			sensorStateStore[device.UniqueID] = state
		}
	} else {
		log.Fatalln("No statefunc, this makes no sensor for a sensor!")
	}
}
func (device BinarySensor) UpdateState(client mqtt.Client) {
	if device.StateFunc != nil {
		state := device.StateFunc()
		if state != binarySensorStateStore[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.GetStateTopic(), 0, false, device.StateFunc())
			token.Wait()
			binarySensorStateStore[device.UniqueID] = state
		}
	} else {
		log.Fatalln("No statefunc, this makes no sensor for a binary sensor!")
	}
}

func (device Switch) Subscribe(client mqtt.Client) {

	message, err := json.Marshal(device)
	if err != nil {
		log.Fatal(err)
	}

	token := client.Publish(device.GetDiscoveryTopic(), 0, true, message)
	token.Wait()

	// HA needs time to process
	time.Sleep(500 * time.Millisecond)

	device.AnnounceAvailable(client)

	if device.StateFunc != nil {
		device.UpdateState(client)
	}

	if token := client.Subscribe(device.GetCommandTopic(), 0, device.messageHandler); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		os.Exit(1)
	}

}
func (device Sensor) Subscribe(client mqtt.Client) {

	message, err := json.Marshal(device)
	if err != nil {
		log.Fatal(err)
	}

	token := client.Publish(device.GetDiscoveryTopic(), 0, true, message)
	token.Wait()

	// HA needs time to process
	time.Sleep(500 * time.Millisecond)

	device.UpdateState(client)
	device.AnnounceAvailable(client)

}
func (device BinarySensor) Subscribe(client mqtt.Client) {

	message, err := json.Marshal(device)
	if err != nil {
		log.Fatal(err)
	}

	token := client.Publish(device.GetDiscoveryTopic(), 0, true, message)
	token.Wait()

	// HA needs time to process
	time.Sleep(500 * time.Millisecond)

	device.UpdateState(client)
	device.AnnounceAvailable(client)

}

func (device Switch) UnSubscribe(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), 0, false, "offline")
	token.Wait()

	if token := client.Unsubscribe(device.GetCommandTopic()); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		os.Exit(1)
	}
}
func (device Sensor) UnSubscribe(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), 0, false, "offline")
	token.Wait()
}
func (device BinarySensor) UnSubscribe(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), 0, false, "offline")
	token.Wait()
}

func (device Switch) AnnounceAvailable(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), 0, false, "online")
	token.Wait()
}
func (device Sensor) AnnounceAvailable(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), 0, false, "online")
	token.Wait()
}
func (device BinarySensor) AnnounceAvailable(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), 0, false, "online")
	token.Wait()
}

///////////////////
