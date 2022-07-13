package ExternalDevice

import (
	"encoding/json"
	common "github.com/W-Floyd/ha-mqtt-iot/common"
	"log"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (d *BinarySensor) GetRawId() string {
	return "binary_sensor"
}
func (d *BinarySensor) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *BinarySensor) GetUniqueId() string {
	return d.UniqueId
}
func (d *BinarySensor) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type BinarySensor struct {
	AvailabilityMode     string `json:"availability_mode"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate string `json:"availability_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    string `json:"availability_topic"`    // "The MQTT topic subscribed to receive birth and LWT messages from the MQTT device. If `availability` is not defined, the binary sensor will always be considered `available` and its state will be `on`, `off` or `unknown`. If `availability` is defined, the binary sensor will be considered as `unavailable` by default and the sensor's initial state will be `unavailable`. Must not be used together with `availability`."
	Device               struct {
		ConfigurationUrl string `json:"configuration_url"` // "A link to the webpage that can manage the configuration of this device. Can be either an HTTP or HTTPS link."
		Connections      string `json:"connections"`       // "A list of connections of the device to the outside world as a list of tuples `[connection_type, connection_identifier]`. For example the MAC address of a network interface: `'connections': ['mac', '02:5b:26:a8:dc:12']`."
		Identifiers      string `json:"identifiers"`       // "A list of IDs that uniquely identify the device. For example a serial number."
		Manufacturer     string `json:"manufacturer"`      // "The manufacturer of the device."
		Model            string `json:"model"`             // "The model of the device."
		Name             string `json:"name"`              // "The name of the device."
		SuggestedArea    string `json:"suggested_area"`    // "Suggest an area if the device isn’t in one yet."
		SwVersion        string `json:"sw_version"`        // "The firmware version of the device."
		Viadevice        string `json:"viadevice"`         // null
	} `json:"device"`
	DeviceClass         string        `json:"device_class"`          // "Sets the [class of the device](/integrations/binary_sensor/#device-class), changing the device state and icon that is displayed on the frontend."
	EnabledByDefault    bool          `json:"enabled_by_default"`    // "Flag which defines if the entity should be enabled when first added."
	Encoding            string        `json:"encoding"`              // "The encoding of the payloads received. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory      string        `json:"entity_category"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	ExpireAfter         int           `json:"expire_after"`          // "Defines the number of seconds after the sensor's state expires, if it's not updated. After expiry, the sensor's state becomes `unavailable`."
	ForceUpdate         bool          `json:"force_update"`          // "Sends update events (which results in update of [state object](/docs/configuration/state_object/)'s `last_changed`) even if the sensor's state hasn't changed. Useful if you want to have meaningful value graphs in history or want to create an automation that triggers on *every* incoming state message (not only when the sensor's new state is different to the current one)."
	Icon                string        `json:"icon"`                  // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Name                string        `json:"name"`                  // "The name of the binary sensor."
	ObjectId            string        `json:"object_id"`             // "Used instead of `name` for automatic generation of `entity_id`"
	OffDelay            int           `json:"off_delay"`             // "For sensors that only send `on` state updates (like PIRs), this variable sets a delay in seconds after which the sensor's state will be updated back to `off`."
	PayloadAvailable    string        `json:"payload_available"`     // "The string that represents the `online` state."
	PayloadNotAvailable string        `json:"payload_not_available"` // "The string that represents the `offline` state."
	PayloadOff          string        `json:"payload_off"`           // "The string that represents the `off` state. It will be compared to the message in the `state_topic` (see `value_template` for details)"
	PayloadOn           string        `json:"payload_on"`            // "The string that represents the `on` state. It will be compared to the message in the `state_topic` (see `value_template` for details)"
	Qos                 int           `json:"qos"`                   // "The maximum QoS level to be used when receiving messages."
	StateTopic          string        `json:"state_topic"`           // "The MQTT topic subscribed to receive sensor's state."
	StateFunc           func() string `json:"-"`
	UniqueId            string        `json:"unique_id"`      // "An ID that uniquely identifies this sensor. If two sensors have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate       string        `json:"value_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) that returns a string to be compared to `payload_on`/`payload_off` or an empty string, in which case the MQTT message will be removed. Available variables: `entity_id`. Remove this option when 'payload_on' and 'payload_off' are sufficient to match your payloads (i.e no pre-processing of original message is required)."
	MQTT                MQTTFields    `json:"-"`
}

func (d *BinarySensor) UpdateState() {
	if d.StateTopic != "" {
		state := d.StateFunc()
		if state != stateStore.BinarySensor.State[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.StateTopic, common.QoS, common.Retain, state)
			stateStore.BinarySensor.State[d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d *BinarySensor) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	token := c.Publish(GetDiscoveryTopic(d), 0, true, message)
	token.Wait()
	time.Sleep(common.HADiscoveryDelay)
	d.AnnounceAvailable()
	d.UpdateState()
}
func (d *BinarySensor) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, "offline")
	token.Wait()
}
func (d *BinarySensor) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d *BinarySensor) Initialize() {
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d *BinarySensor) PopulateTopics() {
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d *BinarySensor) SetMQTTFields(fields MQTTFields) {
	d.MQTT = fields
}
func (d *BinarySensor) GetMQTTFields() (fields MQTTFields) {
	return d.MQTT
}
