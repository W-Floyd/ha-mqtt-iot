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

// DiscoveryPrefix is the prefix that HA uses to discover on.
// Does not usually need changing.
const DiscoveryPrefix = "homeassistant"

// SWVersion is the software version.
// TODO - Move this elsewhere maybe?
const SWVersion = "0.2.0"

// InstanceName is the instance name, helpful for identifying a given client
var InstanceName = "Homeassistant MQTT IOT"

// NodeID is the Node ID, that is, what that node connects under.
var NodeID = "ha-mqtt-iot"

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

// Light matches a light
type Light struct {
	AvailabilityTopic        string   `json:"availability_topic,omitempty"`
	BrightnessCommandTopic   string   `json:"brightness_command_topic,omitempty"`
	BrightnessScale          int      `json:"brightness_scale,omitempty"`
	BrightnessStateTopic     string   `json:"brightness_state_topic,omitempty"`
	BrightnessValueTemplate  string   `json:"brightness_value_template,omitempty"`
	ColorTempCommandTemplate string   `json:"color_temp_command_template,omitempty"`
	ColorTempCommandTopic    string   `json:"color_temp_command_topic,omitempty"`
	ColorTempStateTopic      string   `json:"color_temp_state_topic,omitempty"`
	ColorTempValueTemplate   string   `json:"color_temp_value_template,omitempty"`
	CommandTopic             string   `json:"command_topic"`
	Device                   device   `json:"device,omitempty"`
	EffectCommandTopic       string   `json:"effect_command_topic,omitempty"`
	EffectList               []string `json:"effect_list,omitempty"`
	EffectStateTopic         string   `json:"effect_state_topic,omitempty"`
	EffectValueTemplate      string   `json:"effect_value_template,omitempty"`
	HsCommandTopic           string   `json:"hs_command_topic,omitempty"`
	HsStateTopic             string   `json:"hs_state_topic,omitempty"`
	HsValueTemplate          string   `json:"hs_value_template,omitempty"`
	JSONAttributesTemplate   string   `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic      string   `json:"json_attributes_topic,omitempty"`
	MaxMireds                int      `json:"max_mireds,omitempty"`
	MinMireds                int      `json:"min_mireds,omitempty"`
	Name                     string   `json:"name,omitempty"`
	OnCommandType            string   `json:"on_command_type,omitempty"`
	Optimistic               bool     `json:"opt,omitempty"`
	PayloadAvailable         string   `json:"payload_available,omitempty"`
	PayloadNotAvailable      string   `json:"payload_not_available,omitempty"`
	PayloadOff               string   `json:"pl_off,omitempty"`
	PayloadOn                string   `json:"pl_on,omitempty"`
	QOS                      int      `json:"qos,omitempty"`
	Retain                   bool     `json:"ret,omitempty"`
	RgbCommandTemplate       string   `json:"rgb_command_template,omitempty"`
	RgbCommandTopic          string   `json:"rgb_command_topic,omitempty"`
	RgbStateTopic            string   `json:"rgb_state_topic,omitempty"`
	RgbValueTemplate         string   `json:"rgb_value_template,omitempty"`
	Schema                   string   `json:"schema,omitempty"`
	StateTopic               string   `json:"state_topic,omitempty"`
	StateValueTemplate       string   `json:"state_value_template,omitempty"`
	UniqueID                 string   `json:"unique_id,omitempty"`
	WhiteValueCommandTopic   string   `json:"white_value_command_topic,omitempty"`
	WhiteValueScale          int      `json:"white_value_scale,omitempty"`
	WhiteValueStateTopic     string   `json:"white_value_state_topic,omitempty"`
	WhiteValueTemplate       string   `json:"white_value_template,omitempty"`
	XyCommandTopic           string   `json:"xy_command_topic,omitempty"`
	XyStateTopic             string   `json:"xy_state_topic,omitempty"`
	XyValueTemplate          string   `json:"xy_value_template,omitempty"`
	ValueTemplate            string   `json:"value_template,omitempty"`

	BrightnessStateFunc func() string `json:"-"`
	ColorTempStateFunc  func() string `json:"-"`
	EffectStateFunc     func() string `json:"-"`
	HsStateFunc         func() string `json:"-"`
	RgbStateFunc        func() string `json:"-"`
	StateFunc           func() string `json:"-"`
	WhiteValueStateFunc func() string `json:"-"`
	XyStateFunc         func() string `json:"-"`

	BrightnessCommandFunc func(mqtt.Message, mqtt.Client) `json:"-"`
	ColorTempCommandFunc  func(mqtt.Message, mqtt.Client) `json:"-"`
	CommandFunc           func(mqtt.Message, mqtt.Client) `json:"-"`
	EffectCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	HsCommandFunc         func(mqtt.Message, mqtt.Client) `json:"-"`
	RgbCommandFunc        func(mqtt.Message, mqtt.Client) `json:"-"`
	WhiteValueCommandFunc func(mqtt.Message, mqtt.Client) `json:"-"`
	XyCommandFunc         func(mqtt.Message, mqtt.Client) `json:"-"`

	UpdateInterval  float64 `json:"-"`
	ForceUpdateMQTT bool    `json:"-"`

	messageHandler mqtt.MessageHandler
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

	messageHandler mqtt.MessageHandler
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
	d.SWVersion = SWVersion

	return
}

///////////////////

type store struct {
	BinarySensor map[string]string
	Light        struct {
		BrightnessState map[string]string
		ColorTempState  map[string]string
		EffectState     map[string]string
		HsState         map[string]string
		RgbState        map[string]string
		State           map[string]string
		WhiteValueState map[string]string
		XyState         map[string]string
	}
	Sensor map[string]string
	Switch map[string]string
}

func initStore() store {
	var s store
	s.BinarySensor = make(map[string]string)
	s.Light.BrightnessState = make(map[string]string)
	s.Light.ColorTempState = make(map[string]string)
	s.Light.EffectState = make(map[string]string)
	s.Light.HsState = make(map[string]string)
	s.Light.RgbState = make(map[string]string)
	s.Light.State = make(map[string]string)
	s.Light.WhiteValueState = make(map[string]string)
	s.Light.XyState = make(map[string]string)
	s.Sensor = make(map[string]string)
	s.Switch = make(map[string]string)
	return s
}

var topicStore = make(map[string]*func(mqtt.Message, mqtt.Client))
var stateStore = initStore()

///////////////////

// GetTopicPrefix gets the prefix for all state/command topics
// This is for a light
func (device Light) GetTopicPrefix() string {
	return NodeID + "/light/" + device.UniqueID + "/"
}

// GetTopicPrefix gets the prefix for all state/command topics
// This is for a sensor
func (device Sensor) GetTopicPrefix() string {
	return NodeID + "/sensor/" + device.UniqueID + "/"
}

// GetTopicPrefix gets the prefix for all state/command topics
// This is for a switch
func (device Switch) GetTopicPrefix() string {
	return NodeID + "/switch/" + device.UniqueID + "/"
}

// GetTopicPrefix gets the prefix for all state/command topics
// This is for a binary sensor
func (device BinarySensor) GetTopicPrefix() string {
	return NodeID + "/binary_sensor/" + device.UniqueID + "/"
}

// GetDiscoveryTopic gets the topic for a device's discovery topic.
// This is for a light
func (device Light) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/light/" + NodeID + "/" + device.UniqueID + "/" + "config"
}

// GetDiscoveryTopic gets the topic for a device's discovery topic.
// This is for a sensor
func (device Sensor) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/sensor/" + NodeID + "/" + device.UniqueID + "/" + "config"
}

// GetDiscoveryTopic gets the topic for a device's discovery topic.
// This is for a switch
func (device Switch) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/switch/" + NodeID + "/" + device.UniqueID + "/" + "config"
}

// GetDiscoveryTopic gets the topic for a device's discovery topic.
// This is for a binary sensor
func (device BinarySensor) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/binary_sensor/" + NodeID + "/" + device.UniqueID + "/" + "config"
}

// GetCommandTopic gets the command topic for a device
// This is for a light
func (device Light) GetCommandTopic() string {
	return device.GetTopicPrefix() + "command"
}

// GetCommandTopic gets the command topic for a device
// This is for a switch
func (device Switch) GetCommandTopic() string {
	return device.GetTopicPrefix() + "command"
}

// GetStateTopic gets the state topic for a device
// This is for a light
func (device Light) GetStateTopic() string {
	return device.GetTopicPrefix() + "state"
}

// GetStateTopic gets the state topic for a device
// This is for a sensor
func (device Sensor) GetStateTopic() string {
	return device.GetTopicPrefix() + "state"
}

// GetStateTopic gets the state topic for a device
// This is for a switch
func (device Switch) GetStateTopic() string {
	return device.GetTopicPrefix() + "state"
}

// GetStateTopic gets the state topic for a device
// This is for a binary sensor
func (device BinarySensor) GetStateTopic() string {
	return device.GetTopicPrefix() + "state"
}

// GetAvailabilityTopic gets the availability topic for a device
// This is for a light
func (device Light) GetAvailabilityTopic() string {
	return device.GetTopicPrefix() + "availability"
}

// GetAvailabilityTopic gets the availability topic for a device
// This is for a sensor
func (device Sensor) GetAvailabilityTopic() string {
	return device.GetTopicPrefix() + "availability"
}

// GetAvailabilityTopic gets the availability topic for a device
// This is for a switch
func (device Switch) GetAvailabilityTopic() string {
	return device.GetTopicPrefix() + "availability"
}

// GetAvailabilityTopic gets the availability topic for a device
// This is for a binary sensor
func (device BinarySensor) GetAvailabilityTopic() string {
	return device.GetTopicPrefix() + "availability"
}

// Initialize sets topics as needed on a Light
func (device *Light) Initialize() {
	device.Retain = false
	device.Device = getDevice()

	device.AvailabilityTopic = device.GetAvailabilityTopic()

	// Brightness
	if device.BrightnessCommandFunc != nil {
		device.BrightnessCommandTopic = device.GetTopicPrefix() + "brightness/set"
		topicStore[device.BrightnessCommandTopic] = &device.BrightnessCommandFunc
	}
	if device.BrightnessStateFunc != nil {
		device.BrightnessStateTopic = device.GetTopicPrefix() + "brightness/get"
	}

	// ColorTemp
	if device.ColorTempCommandFunc != nil {
		device.ColorTempCommandTopic = device.GetTopicPrefix() + "color-temp/set"
		topicStore[device.ColorTempCommandTopic] = &device.ColorTempCommandFunc
	}
	if device.ColorTempStateFunc != nil {
		device.ColorTempStateTopic = device.GetTopicPrefix() + "color-temp/get"
	}

	// Command/State
	if device.CommandFunc != nil {
		device.CommandTopic = device.GetCommandTopic()
		topicStore[device.CommandTopic] = &device.CommandFunc
	}
	if device.StateFunc != nil {
		device.StateTopic = device.GetStateTopic()
	}

	// Effect
	if device.EffectCommandFunc != nil {
		device.EffectCommandTopic = device.GetTopicPrefix() + "effect/set"
		topicStore[device.EffectCommandTopic] = &device.EffectCommandFunc
	}
	if device.EffectStateFunc != nil {
		device.EffectStateTopic = device.GetTopicPrefix() + "effect/get"
	}

	// Hs
	if device.HsCommandFunc != nil {
		device.HsCommandTopic = device.GetTopicPrefix() + "hs/set"
		topicStore[device.HsCommandTopic] = &device.HsCommandFunc
	}
	if device.HsStateFunc != nil {
		device.HsStateTopic = device.GetTopicPrefix() + "hs/get"
	}

	// Rgb
	if device.RgbCommandFunc != nil {
		device.RgbCommandTopic = device.GetTopicPrefix() + "rgb/set"
		topicStore[device.RgbCommandTopic] = &device.RgbCommandFunc
	}
	if device.RgbStateFunc != nil {
		device.RgbStateTopic = device.GetTopicPrefix() + "rgb/get"
	}

	// WhiteValue
	if device.WhiteValueCommandFunc != nil {
		device.WhiteValueCommandTopic = device.GetTopicPrefix() + "white-value/set"
		topicStore[device.WhiteValueCommandTopic] = &device.WhiteValueCommandFunc
	}
	if device.WhiteValueStateFunc != nil {
		device.WhiteValueStateTopic = device.GetTopicPrefix() + "white-value/get"
	}

	// Xy
	if device.XyCommandFunc != nil {
		device.XyCommandTopic = device.GetTopicPrefix() + "xy/set"
		topicStore[device.XyCommandTopic] = &device.XyCommandFunc
	}
	if device.XyStateFunc != nil {
		device.XyStateTopic = device.GetTopicPrefix() + "xy/get"
	}

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

// Initialize sets topics as needed on a Binary Sensor
func (device *BinarySensor) Initialize() {
	device.StateTopic = device.GetStateTopic()
	device.AvailabilityTopic = device.GetAvailabilityTopic()
	device.Device = getDevice()
}

// UpdateState publishes any new state for a device
// This is for a light
func (device Light) UpdateState(client mqtt.Client) {
	if device.BrightnessStateTopic != "" {
		state := device.BrightnessStateFunc()
		if state != stateStore.Light.BrightnessState[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.BrightnessStateTopic, 0, false, state)
			stateStore.Light.BrightnessState[device.UniqueID] = state
			token.Wait()
		}
	}
	if device.ColorTempStateTopic != "" {
		state := device.ColorTempStateFunc()
		if state != stateStore.Light.ColorTempState[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.ColorTempStateTopic, 0, false, state)
			stateStore.Light.ColorTempState[device.UniqueID] = state
			token.Wait()
		}
	}
	if device.EffectStateTopic != "" {
		state := device.EffectStateFunc()
		if state != stateStore.Light.EffectState[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.EffectStateTopic, 0, false, state)
			stateStore.Light.EffectState[device.UniqueID] = state
			token.Wait()
		}
	}
	if device.HsStateTopic != "" {
		state := device.HsStateFunc()
		if state != stateStore.Light.HsState[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.HsStateTopic, 0, false, state)
			stateStore.Light.HsState[device.UniqueID] = state
			token.Wait()
		}
	}
	if device.RgbStateTopic != "" {
		state := device.RgbStateFunc()
		if state != stateStore.Light.RgbState[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.RgbStateTopic, 0, false, state)
			stateStore.Light.RgbState[device.UniqueID] = state
			token.Wait()
		}
	}
	if device.StateTopic != "" {
		state := device.StateFunc()
		if state != stateStore.Light.State[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.StateTopic, 0, false, state)
			stateStore.Light.State[device.UniqueID] = state
			token.Wait()
		}
	}
	if device.WhiteValueStateTopic != "" {
		state := device.WhiteValueStateFunc()
		if state != stateStore.Light.WhiteValueState[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.WhiteValueStateTopic, 0, false, state)
			stateStore.Light.WhiteValueState[device.UniqueID] = state
			token.Wait()
		}
	}
	if device.XyStateTopic != "" {
		state := device.XyStateFunc()
		if state != stateStore.Light.XyState[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.XyStateTopic, 0, false, state)
			stateStore.Light.XyState[device.UniqueID] = state
			token.Wait()
		}
	}
}

// UpdateState publishes any new state for a device
// This is for a switch
func (device Switch) UpdateState(client mqtt.Client) {
	if device.StateFunc != nil {
		state := device.StateFunc()
		if state != stateStore.Switch[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.StateTopic, 0, false, state)
			stateStore.Switch[device.UniqueID] = state
			token.Wait()
		}
	} else {
		log.Println("No statefunc for " + device.UniqueID + strconv.FormatFloat(float64(device.UpdateInterval), 'g', 1, 64))
	}
}

// UpdateState publishes any new state for a device
// This is for a sensor
func (device Sensor) UpdateState(client mqtt.Client) {
	if device.StateFunc != nil {
		state := device.StateFunc()
		if state != stateStore.Sensor[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.StateTopic, 0, false, state)
			token.Wait()
			stateStore.Sensor[device.UniqueID] = state
		}
	} else {
		log.Fatalln("No statefunc, this makes no sensor for a sensor!")
	}
}

// UpdateState publishes any new state for a device
// This is for a binary sensor
func (device BinarySensor) UpdateState(client mqtt.Client) {
	if device.StateFunc != nil {
		state := device.StateFunc()
		if state != stateStore.BinarySensor[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.StateTopic, 0, false, state)
			token.Wait()
			stateStore.BinarySensor[device.UniqueID] = state
		}
	} else {
		log.Fatalln("No statefunc, this makes no sensor for a binary sensor!")
	}
}

// Subscribe announces and starts listening to MQTT topics appropriate for a device
// This is for a light
func (device Light) Subscribe(client mqtt.Client) {

	message, err := json.Marshal(device)
	if err != nil {
		log.Fatal(err)
	}

	if device.BrightnessCommandTopic != "" {
		if token := client.Subscribe(device.BrightnessCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}
	}
	if device.ColorTempCommandTopic != "" {
		if token := client.Subscribe(device.ColorTempCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}
	}
	if device.EffectCommandTopic != "" {
		if token := client.Subscribe(device.EffectCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}
	}
	if device.HsCommandTopic != "" {
		if token := client.Subscribe(device.HsCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}
	}
	if device.RgbCommandTopic != "" {
		if token := client.Subscribe(device.RgbCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}
	}
	if device.CommandTopic != "" {
		if token := client.Subscribe(device.CommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}
	}
	if device.WhiteValueCommandTopic != "" {
		if token := client.Subscribe(device.WhiteValueCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}
	}
	if device.XyCommandTopic != "" {
		if token := client.Subscribe(device.XyCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}
	}

	token := client.Publish(device.GetDiscoveryTopic(), 0, true, message)
	token.Wait()

	// HA needs time to process
	time.Sleep(500 * time.Millisecond)

	device.AnnounceAvailable(client)

	device.UpdateState(client)

	if token := client.Subscribe(device.GetCommandTopic(), 0, device.messageHandler); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		os.Exit(1)
	}

}

// Subscribe announces and starts listening to MQTT topics appropriate for a device
// This is for a switch
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

// Subscribe announces and MQTT topics appropriate for a device
// This is for a sensor
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

// Subscribe announces and MQTT topics appropriate for a device
// This is for a binary sensor
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

// UnSubscribe from MQTT topics appropriate for a device, and publishes unavailability
// This is for a light
func (device Light) UnSubscribe(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), 0, false, "offline")
	token.Wait()

	if device.BrightnessCommandTopic != "" {
		if token := client.Unsubscribe(device.BrightnessCommandTopic); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}

	}
	if device.ColorTempCommandTopic != "" {
		if token := client.Unsubscribe(device.ColorTempCommandTopic); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}

	}
	if device.EffectCommandTopic != "" {
		if token := client.Unsubscribe(device.EffectCommandTopic); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}

	}
	if device.HsCommandTopic != "" {
		if token := client.Unsubscribe(device.HsCommandTopic); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}

	}
	if device.RgbCommandTopic != "" {
		if token := client.Unsubscribe(device.RgbCommandTopic); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}

	}
	if device.CommandTopic != "" {
		if token := client.Unsubscribe(device.CommandTopic); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}

	}
	if device.WhiteValueCommandTopic != "" {
		if token := client.Unsubscribe(device.WhiteValueCommandTopic); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}

	}
	if device.XyCommandTopic != "" {
		if token := client.Unsubscribe(device.XyCommandTopic); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}

	}

}

// UnSubscribe from MQTT topics appropriate for a device, and publishes unavailability
// This is for a switch
func (device Switch) UnSubscribe(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), 0, false, "offline")
	token.Wait()

	if token := client.Unsubscribe(device.GetCommandTopic()); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		os.Exit(1)
	}
}

// UnSubscribe publishes unavailability for a device
// This is for a sensor
func (device Sensor) UnSubscribe(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), 0, false, "offline")
	token.Wait()
}

// UnSubscribe publishes unavailability for a device
// This is for a binary sensor
func (device BinarySensor) UnSubscribe(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), 0, false, "offline")
	token.Wait()
}

// AnnounceAvailable publishes availability for a device
// This is for a light
func (device Light) AnnounceAvailable(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), 0, false, "online")
	token.Wait()
}

// AnnounceAvailable publishes availability for a device
// This is for a switch
func (device Switch) AnnounceAvailable(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), 0, false, "online")
	token.Wait()
}

// AnnounceAvailable publishes availability for a device
// This is for a sensor
func (device Sensor) AnnounceAvailable(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), 0, false, "online")
	token.Wait()
}

// AnnounceAvailable publishes availability for a device
// This is for a binary sensor
func (device BinarySensor) AnnounceAvailable(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), 0, false, "online")
	token.Wait()
}

///////////////////
