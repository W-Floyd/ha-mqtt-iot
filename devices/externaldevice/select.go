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
func (d Select) GetRawId() string {
	return "select"
}
func (d Select) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Select) GetUniqueId() string {
	return d.UniqueId
}
func (d Select) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Select struct {
	AvailabilityMode     string                          `json:"availability_mode"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate string                          `json:"availability_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    string                          `json:"availability_topic"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	CommandTemplate      string                          `json:"command_template"`      // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `command_topic`."
	CommandTopic         string                          `json:"command_topic"`         // "The MQTT topic to publish commands to change the selected option."
	CommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
	Device               struct {
		ConfigurationUrl string `json:"configuration_url"` // "A link to the webpage that can manage the configuration of this device. Can be either an HTTP or HTTPS link."
		Connections      string `json:"connections"`       // "A list of connections of the device to the outside world as a list of tuples `[connection_type, connection_identifier]`. For example the MAC address of a network interface: `\"connections\": [\"mac\", \"02:5b:26:a8:dc:12\"]`."
		Identifiers      string `json:"identifiers"`       // "A list of IDs that uniquely identify the device. For example a serial number."
		Manufacturer     string `json:"manufacturer"`      // "The manufacturer of the device."
		Model            string `json:"model"`             // "The model of the device."
		Name             string `json:"name"`              // "The name of the device."
		SuggestedArea    string `json:"suggested_area"`    // "Suggest an area if the device isn’t in one yet."
		SwVersion        string `json:"sw_version"`        // "The firmware version of the device."
		Viadevice        string `json:"viadevice"`         // null
	} `json:"device"`
	EnabledByDefault bool          `json:"enabled_by_default"` // "Flag which defines if the entity should be enabled when first added."
	Encoding         string        `json:"encoding"`           // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory   string        `json:"entity_category"`    // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon             string        `json:"icon"`               // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Name             string        `json:"name"`               // "The name of the Select."
	ObjectId         string        `json:"object_id"`          // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic       bool          `json:"optimistic"`         // "Flag that defines if the select works in optimistic mode."
	Options          []string      `json:"options"`            // "List of options that can be selected. An empty list or a list with a single item is allowed."
	Qos              int           `json:"qos"`                // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain           bool          `json:"retain"`             // "If the published message should have the retain flag on or not."
	StateTopic       string        `json:"state_topic"`        // "The MQTT topic subscribed to receive update of the selected option."
	StateFunc        func() string `json:"-"`
	UniqueId         string        `json:"unique_id"`      // "An ID that uniquely identifies this Select. If two Selects have the same unique ID Home Assistant will raise an exception."
	ValueTemplate    string        `json:"value_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the value."
	MQTT             MQTTFields    `json:"-"`
}

func (d Select) UpdateState() {
	if d.StateTopic != "" {
		state := d.StateFunc()
		if state != stateStore.Select.State[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.StateTopic, common.QoS, common.Retain, state)
			stateStore.Select.State[d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d Select) Subscribe() {
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
func (d Select) UnSubscribe() {
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
func (d Select) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d Select) Initialize() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d Select) PopulateTopics() {
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		store.TopicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
