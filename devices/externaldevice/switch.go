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
func (d *Switch) GetRawId() string {
	return "switch"
}
func (d *Switch) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Switch) GetUniqueId() string {
	return d.UniqueId
}
func (d *Switch) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Switch struct {
	AvailabilityMode     string                          `json:"availability_mode"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate string                          `json:"availability_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    string                          `json:"availability_topic"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	CommandTopic         string                          `json:"command_topic"`         // "The MQTT topic to publish commands to change the switch state."
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
	DeviceClass         string        `json:"device_class"`          // "The [type/class](/integrations/switch/#device-class) of the switch to set the icon in the frontend."
	EnabledByDefault    bool          `json:"enabled_by_default"`    // "Flag which defines if the entity should be enabled when first added."
	Encoding            string        `json:"encoding"`              // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory      string        `json:"entity_category"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                string        `json:"icon"`                  // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Name                string        `json:"name"`                  // "The name to use when displaying this switch."
	ObjectId            string        `json:"object_id"`             // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic          bool          `json:"optimistic"`            // "Flag that defines if switch works in optimistic mode."
	PayloadAvailable    string        `json:"payload_available"`     // "The payload that represents the available state."
	PayloadNotAvailable string        `json:"payload_not_available"` // "The payload that represents the unavailable state."
	PayloadOff          string        `json:"payload_off"`           // "The payload that represents `off` state. If specified, will be used for both comparing to the value in the `state_topic` (see `value_template` and `state_off` for details) and sending as `off` command to the `command_topic`."
	PayloadOn           string        `json:"payload_on"`            // "The payload that represents `on` state. If specified, will be used for both comparing to the value in the `state_topic` (see `value_template` and `state_on`  for details) and sending as `on` command to the `command_topic`."
	Qos                 int           `json:"qos"`                   // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain              bool          `json:"retain"`                // "If the published message should have the retain flag on or not."
	StateOff            string        `json:"state_off"`             // "The payload that represents the `off` state. Used when value that represents `off` state in the `state_topic` is different from value that should be sent to the `command_topic` to turn the device `off`."
	StateOn             string        `json:"state_on"`              // "The payload that represents the `on` state. Used when value that represents `on` state in the `state_topic` is different from value that should be sent to the `command_topic` to turn the device `on`."
	StateTopic          string        `json:"state_topic"`           // "The MQTT topic subscribed to receive state updates."
	StateFunc           func() string `json:"-"`
	UniqueId            string        `json:"unique_id"`      // "An ID that uniquely identifies this switch device. If two switches have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate       string        `json:"value_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's state from the `state_topic`. To determine the switches's state result of this template will be compared to `state_on` and `state_off`."
	MQTT                MQTTFields    `json:"-"`
}

func (d *Switch) UpdateState() {
	if d.StateTopic != "" {
		state := d.StateFunc()
		if state != stateStore.Switch.State[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.StateTopic, common.QoS, common.Retain, state)
			stateStore.Switch.State[d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d *Switch) Subscribe() {
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
	token := c.Publish(GetDiscoveryTopic(d), 0, true, message)
	token.Wait()
	time.Sleep(common.HADiscoveryDelay)
	d.AnnounceAvailable()
	d.UpdateState()
}
func (d *Switch) UnSubscribe() {
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
}
func (d *Switch) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d *Switch) Initialize() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d *Switch) PopulateTopics() {
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		store.TopicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d *Switch) SetMQTTFields(fields MQTTFields) {
	d.MQTT = fields
}
func (d *Switch) GetMQTTFields() (fields MQTTFields) {
	return d.MQTT
}
