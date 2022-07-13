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
func (d *Number) GetRawId() string {
	return "number"
}
func (d *Number) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Number) GetUniqueId() string {
	return d.UniqueId
}
func (d *Number) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Number struct {
	AvailabilityMode  string                          `json:"availability_mode"`  // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTopic string                          `json:"availability_topic"` // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	AvailabilityFunc  func() string                   `json:"-"`
	CommandTemplate   string                          `json:"command_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `command_topic`."
	CommandTopic      string                          `json:"command_topic"`    // "The MQTT topic to publish commands to change the number."
	CommandFunc       func(mqtt.Message, mqtt.Client) `json:"-"`
	Device            struct {
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
	DeviceClass       string        `json:"device_class"`       // "The [type/class](/integrations/number/#device-class) of the number."
	EnabledByDefault  bool          `json:"enabled_by_default"` // "Flag which defines if the entity should be enabled when first added."
	Encoding          string        `json:"encoding"`           // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory    string        `json:"entity_category"`    // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon              string        `json:"icon"`               // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Max               float64       `json:"max"`                // "Maximum value."
	Min               float64       `json:"min"`                // "Minimum value."
	Name              string        `json:"name"`               // "The name of the Number."
	ObjectId          string        `json:"object_id"`          // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic        bool          `json:"optimistic"`         // "Flag that defines if number works in optimistic mode."
	PayloadReset      string        `json:"payload_reset"`      // "A special payload that resets the state to `None` when received on the `state_topic`."
	Qos               int           `json:"qos"`                // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain            bool          `json:"retain"`             // "If the published message should have the retain flag on or not."
	StateTopic        string        `json:"state_topic"`        // "The MQTT topic subscribed to receive number values."
	StateFunc         func() string `json:"-"`
	Step              float64       `json:"step"`                // "Step value. Smallest value `0.001`."
	UniqueId          string        `json:"unique_id"`           // "An ID that uniquely identifies this Number. If two Numbers have the same unique ID Home Assistant will raise an exception."
	UnitOfMeasurement string        `json:"unit_of_measurement"` // "Defines the unit of measurement of the sensor, if any."
	ValueTemplate     string        `json:"value_template"`      // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the value."
	MQTT              MQTTFields    `json:"-"`
}

func (d *Number) UpdateState() {
	if d.AvailabilityTopic != "" {
		state := d.AvailabilityFunc()
		if state != stateStore.Number.Availability[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, state)
			stateStore.Number.Availability[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.StateTopic != "" {
		state := d.StateFunc()
		if state != stateStore.Number.State[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.StateTopic, common.QoS, common.Retain, state)
			stateStore.Number.State[d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d *Number) Subscribe() {
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
	d.AvailabilityFunc()
	d.UpdateState()
}
func (d *Number) UnSubscribe() {
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
func (d *Number) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d *Number) Initialize() {
	d.Retain = false
	d.PopulateDevice()
	d.AddMessageHandler()
	d.PopulateTopics()
}
func (d *Number) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		store.TopicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d *Number) SetMQTTFields(fields MQTTFields) {
	d.MQTT = fields
}
func (d *Number) GetMQTTFields() (fields MQTTFields) {
	return d.MQTT
}
