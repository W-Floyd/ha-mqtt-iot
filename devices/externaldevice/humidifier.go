package ExternalDevice

import (
	"encoding/json"
	common "github.com/W-Floyd/ha-mqtt-iot/common"
	store "github.com/W-Floyd/ha-mqtt-iot/store"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (d Humidifier) GetRawId() string {
	return "humidifier"
}
func (d Humidifier) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Humidifier) GetUniqueId() string {
	return d.UniqueId
}
func (d Humidifier) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Humidifier struct {
	AvailabilityMode     string                          `json:"availability_mode"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate string                          `json:"availability_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    string                          `json:"availability_topic"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	CommandTemplate      string                          `json:"command_template"`      // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `command_topic`."
	CommandTopic         string                          `json:"command_topic"`         // "The MQTT topic to publish commands to change the humidifier state."
	CommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
	Device               struct {
		ConfigurationUrl string `json:"configuration_url"` // "A link to the webpage that can manage the configuration of this device. Can be either an HTTP or HTTPS link."
		Connections      string `json:"connections"`       // "A list of connections of the device to the outside world as a list of tuples `[connection_type, connection_identifier]`. For example the MAC address of a network interface: `\"connections\": [[\"mac\", \"02:5b:26:a8:dc:12\"]]`."
		Identifiers      string `json:"identifiers"`       // "A list of IDs that uniquely identify the device. For example a serial number."
		Manufacturer     string `json:"manufacturer"`      // "The manufacturer of the device."
		Model            string `json:"model"`             // "The model of the device."
		Name             string `json:"name"`              // "The name of the device."
		SuggestedArea    string `json:"suggested_area"`    // "Suggest an area if the device isn’t in one yet."
		SwVersion        string `json:"sw_version"`        // "The firmware version of the device."
		Viadevice        string `json:"viadevice"`         // null
	} `json:"device"`
	DeviceClass                   string                          `json:"device_class"`          // "The device class of the MQTT device. Must be either `humidifier` or `dehumidifier`."
	EnabledByDefault              bool                            `json:"enabled_by_default"`    // "Flag which defines if the entity should be enabled when first added."
	Encoding                      string                          `json:"encoding"`              // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory                string                          `json:"entity_category"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                          string                          `json:"icon"`                  // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	MaxHumidity                   int                             `json:"max_humidity"`          // "The minimum target humidity percentage that can be set."
	MinHumidity                   int                             `json:"min_humidity"`          // "The maximum target humidity percentage that can be set."
	ModeCommandTemplate           string                          `json:"mode_command_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `mode_command_topic`."
	ModeCommandTopic              string                          `json:"mode_command_topic"`    // "The MQTT topic to publish commands to change the `mode` on the humidifier. This attribute ust be configured together with the `modes` attribute."
	ModeCommandFunc               func(mqtt.Message, mqtt.Client) `json:"-"`
	ModeStateTemplate             string                          `json:"mode_state_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract a value for the humidifier `mode` state."
	ModeStateTopic                string                          `json:"mode_state_topic"`    // "The MQTT topic subscribed to receive the humidifier `mode`."
	ModeStateFunc                 func() string                   `json:"-"`
	Modes                         []string                        `json:"modes"`                  // "List of available modes this humidifier is capable of running at. Common examples include `normal`, `eco`, `away`, `boost`, `comfort`, `home`, `sleep`, `auto` and `baby`. These examples offer built-in translations but other custom modes are allowed as well.  This attribute ust be configured together with the `mode_command_topic` attribute."
	Name                          string                          `json:"name"`                   // "The name of the humidifier."
	ObjectId                      string                          `json:"object_id"`              // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic                    bool                            `json:"optimistic"`             // "Flag that defines if humidifier works in optimistic mode"
	PayloadAvailable              string                          `json:"payload_available"`      // "The payload that represents the available state."
	PayloadNotAvailable           string                          `json:"payload_not_available"`  // "The payload that represents the unavailable state."
	PayloadOff                    string                          `json:"payload_off"`            // "The payload that represents the stop state."
	PayloadOn                     string                          `json:"payload_on"`             // "The payload that represents the running state."
	PayloadResetHumidity          string                          `json:"payload_reset_humidity"` // "A special payload that resets the `target_humidity` state attribute to `None` when received at the `target_humidity_state_topic`."
	PayloadResetMode              string                          `json:"payload_reset_mode"`     // "A special payload that resets the `mode` state attribute to `None` when received at the `mode_state_topic`."
	Qos                           int                             `json:"qos"`                    // "The maximum QoS level of the state topic."
	Retain                        bool                            `json:"retain"`                 // "If the published message should have the retain flag on or not."
	StateTopic                    string                          `json:"state_topic"`            // "The MQTT topic subscribed to receive state updates."
	StateFunc                     func() string                   `json:"-"`
	StateValueTemplate            string                          `json:"state_value_template"`             // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract a value from the state."
	TargetHumidityCommandTemplate string                          `json:"target_humidity_command_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `target_humidity_command_topic`."
	TargetHumidityCommandTopic    string                          `json:"target_humidity_command_topic"`    // "The MQTT topic to publish commands to change the humidifier target humidity state based on a percentage."
	TargetHumidityCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	TargetHumidityStateTemplate   string                          `json:"target_humidity_state_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract a value for the humidifier `target_humidity` state."
	TargetHumidityStateTopic      string                          `json:"target_humidity_state_topic"`    // "The MQTT topic subscribed to receive humidifier target humidity."
	TargetHumidityStateFunc       func() string                   `json:"-"`
	UniqueId                      string                          `json:"unique_id"` // "An ID that uniquely identifies this humidifier. If two humidifiers have the same unique ID, Home Assistant will raise an exception."
	MQTT                          MQTTFields                      `json:"-"`
}

func (d Humidifier) UpdateState() {
	if d.ModeStateTopic != "" {
		state := d.ModeStateFunc()
		if state != stateStore.Humidifier.ModeState[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.ModeStateTopic, common.QoS, common.Retain, state)
			stateStore.Humidifier.ModeState[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.StateTopic != "" {
		state := d.StateFunc()
		if state != stateStore.Humidifier.State[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.StateTopic, common.QoS, common.Retain, state)
			stateStore.Humidifier.State[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.TargetHumidityStateTopic != "" {
		state := d.TargetHumidityStateFunc()
		if state != stateStore.Humidifier.TargetHumidityState[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.TargetHumidityStateTopic, common.QoS, common.Retain, state)
			stateStore.Humidifier.TargetHumidityState[d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d Humidifier) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	if d.CommandTopic != "" {
		t := c.Subscribe(d.CommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.ModeCommandTopic != "" {
		t := c.Subscribe(d.ModeCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TargetHumidityCommandTopic != "" {
		t := c.Subscribe(d.TargetHumidityCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	token := c.Publish(GetDiscoveryTopic(d), 0, true, message)
	token.Wait()
	time.Sleep(common.HADiscoveryDelay)
	d.AnnounceAvailable()
	d.UpdateState()
}
func (d Humidifier) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, "offline")
	token.Wait()
	if d.CommandTopic != "" {
		t := c.Unsubscribe(d.CommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.ModeCommandTopic != "" {
		t := c.Unsubscribe(d.ModeCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TargetHumidityCommandTopic != "" {
		t := c.Unsubscribe(d.TargetHumidityCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d Humidifier) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d Humidifier) Initialize() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d Humidifier) PopulateTopics() {
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		store.TopicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.ModeCommandFunc != nil {
		d.ModeCommandTopic = GetTopic(d, "mode_command_topic")
		store.TopicStore[d.ModeCommandTopic] = &d.ModeCommandFunc
	}
	if d.ModeStateFunc != nil {
		d.ModeStateTopic = GetTopic(d, "mode_state_topic")
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
	if d.TargetHumidityCommandFunc != nil {
		d.TargetHumidityCommandTopic = GetTopic(d, "target_humidity_command_topic")
		store.TopicStore[d.TargetHumidityCommandTopic] = &d.TargetHumidityCommandFunc
	}
	if d.TargetHumidityStateFunc != nil {
		d.TargetHumidityStateTopic = GetTopic(d, "target_humidity_state_topic")
	}
}
func (d Humidifier) SetMQTTFields(fields MQTTFields) {
	d.MQTT = fields
}
func (d Humidifier) GetMQTTFields() (fields MQTTFields) {
	return d.MQTT
}
