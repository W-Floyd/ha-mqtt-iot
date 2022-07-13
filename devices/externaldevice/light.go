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
func (d *Light) GetRawId() string {
	return "light"
}
func (d *Light) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Light) GetUniqueId() string {
	return d.UniqueId
}
func (d *Light) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Light struct {
	AvailabilityMode     string                          `json:"availability_mode"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate string                          `json:"availability_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    string                          `json:"availability_topic"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	Brightness           bool                            `json:"brightness"`            // "Flag that defines if the light supports brightness."
	BrightnessScale      int                             `json:"brightness_scale"`      // "Defines the maximum brightness value (i.e., 100%) of the MQTT device."
	ColorMode            bool                            `json:"color_mode"`            // "Flag that defines if the light supports color modes."
	CommandTopic         string                          `json:"command_topic"`         // "The MQTT topic to publish commands to change the light’s state."
	CommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
	Device               struct {
		ConfigurationUrl string `json:"configuration_url"` // "A link to the webpage that can manage the configuration of this device. Can be either an HTTP or HTTPS link."
		Connections      string `json:"connections"`       // "A list of connections of the device to the outside world as a list of tuples `[connection_type, connection_identifier]`. For example the MAC address of a network interface: `\"connections\": [[\"mac\", \"02:5b:26:a8:dc:12\"]]`."
		Identifiers      string `json:"identifiers"`       // "A list of IDs that uniquely identify the device. For example a serial number."
		Manufacturer     string `json:"manufacturer"`      // "The manufacturer of the device."
		Model            string `json:"model"`             // "The model of the device."
		Name             string `json:"name"`              // "The name of the device."
		SuggestedArea    string `json:"suggested_area"`    // null
		SwVersion        string `json:"sw_version"`        // "The firmware version of the device."
		Viadevice        string `json:"viadevice"`         // null
	} `json:"device"`
	Effect              bool          `json:"effect"`                // "Flag that defines if the light supports effects."
	EffectList          []string      `json:"effect_list"`           // "The list of effects the light supports."
	EnabledByDefault    bool          `json:"enabled_by_default"`    // "Flag which defines if the entity should be enabled when first added."
	Encoding            string        `json:"encoding"`              // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory      string        `json:"entity_category"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	FlashTimeLong       int           `json:"flash_time_long"`       // "The duration, in seconds, of a “long” flash."
	FlashTimeShort      int           `json:"flash_time_short"`      // "The duration, in seconds, of a “short” flash."
	Icon                string        `json:"icon"`                  // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	MaxMireds           int           `json:"max_mireds"`            // "The maximum color temperature in mireds."
	MinMireds           int           `json:"min_mireds"`            // "The minimum color temperature in mireds."
	Name                string        `json:"name"`                  // "The name of the light."
	ObjectId            string        `json:"object_id"`             // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic          bool          `json:"optimistic"`            // "Flag that defines if the light works in optimistic mode."
	PayloadAvailable    string        `json:"payload_available"`     // "The payload that represents the available state."
	PayloadNotAvailable string        `json:"payload_not_available"` // "The payload that represents the unavailable state."
	Qos                 int           `json:"qos"`                   // "The maximum QoS level of the state topic."
	Retain              bool          `json:"retain"`                // "If the published message should have the retain flag on or not."
	Schema              string        `json:"schema"`                // "The schema to use. Must be `json` to select the JSON schema."
	StateTopic          string        `json:"state_topic"`           // "The MQTT topic subscribed to receive state updates."
	StateFunc           func() string `json:"-"`
	SupportedColorModes []string      `json:"supported_color_modes"` // "A list of color modes supported by the list. This is required if `color_mode` is `True`. Possible color modes are `onoff`, `brightness`, `color_temp`, `hs`, `xy`, `rgb`, `rgbw`, `rgbww`."
	UniqueId            string        `json:"unique_id"`             // "An ID that uniquely identifies this light. If two lights have the same unique ID, Home Assistant will raise an exception."
	MQTT                MQTTFields    `json:"-"`
}

func (d *Light) UpdateState() {
	if d.StateTopic != "" {
		state := d.StateFunc()
		if state != stateStore.Light.State[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.StateTopic, common.QoS, common.Retain, state)
			stateStore.Light.State[d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d *Light) Subscribe() {
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
func (d *Light) UnSubscribe() {
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
func (d *Light) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d *Light) Initialize() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d *Light) PopulateTopics() {
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		store.TopicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d *Light) SetMQTTFields(fields MQTTFields) {
	d.MQTT = fields
}
func (d *Light) GetMQTTFields() (fields MQTTFields) {
	return d.MQTT
}
