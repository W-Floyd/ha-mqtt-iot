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
func (d *Sensor) GetRawId() string {
	return "sensor"
}
func (d *Sensor) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Sensor) GetUniqueId() string {
	return d.UniqueId
}
func (d *Sensor) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Sensor struct {
	AvailabilityMode     string        `json:"availability_mode"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate string        `json:"availability_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    string        `json:"availability_topic"`    // "The MQTT topic subscribed to receive availability (online/offline) updates."
	AvailabilityFunc     func() string `json:"-"`
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
	DeviceClass            string        `json:"device_class"`              // "The [type/class](/integrations/sensor/#device-class) of the sensor to set the icon in the frontend."
	EnabledByDefault       bool          `json:"enabled_by_default"`        // "Flag which defines if the entity should be enabled when first added."
	Encoding               string        `json:"encoding"`                  // "The encoding of the payloads received. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         string        `json:"entity_category"`           // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	ExpireAfter            int           `json:"expire_after"`              // "Defines the number of seconds after the sensor's state expires, if it's not updated. After expiry, the sensor's state becomes `unavailable`."
	ForceUpdate            bool          `json:"force_update"`              // "Sends update events even if the value hasn't changed. Useful if you want to have meaningful value graphs in history."
	Icon                   string        `json:"icon"`                      // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	LastResetValueTemplate string        `json:"last_reset_value_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the last_reset. Available variables: `entity_id`. The `entity_id` can be used to reference the entity's attributes."
	Name                   string        `json:"name"`                      // "The name of the MQTT sensor."
	ObjectId               string        `json:"object_id"`                 // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable       string        `json:"payload_available"`         // "The payload that represents the available state."
	PayloadNotAvailable    string        `json:"payload_not_available"`     // "The payload that represents the unavailable state."
	Qos                    int           `json:"qos"`                       // "The maximum QoS level of the state topic."
	StateClass             string        `json:"state_class"`               // "The [state_class](https://developers.home-assistant.io/docs/core/entity/sensor#available-state-classes) of the sensor."
	StateTopic             string        `json:"state_topic"`               // "The MQTT topic subscribed to receive sensor values."
	StateFunc              func() string `json:"-"`
	UniqueId               string        `json:"unique_id"`           // "An ID that uniquely identifies this sensor. If two sensors have the same unique ID, Home Assistant will raise an exception."
	UnitOfMeasurement      string        `json:"unit_of_measurement"` // "Defines the units of measurement of the sensor, if any."
	ValueTemplate          string        `json:"value_template"`      // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the value. Available variables: `entity_id`. The `entity_id` can be used to reference the entity's attributes. If the template throws an error, the current state will be used instead."
	MQTT                   MQTTFields    `json:"-"`
}

func (d *Sensor) UpdateState() {
	if d.AvailabilityTopic != "" {
		state := d.AvailabilityFunc()
		if state != stateStore.Sensor.Availability[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, state)
			stateStore.Sensor.Availability[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.StateTopic != "" {
		state := d.StateFunc()
		if state != stateStore.Sensor.State[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.StateTopic, common.QoS, common.Retain, state)
			stateStore.Sensor.State[d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d *Sensor) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	token := c.Publish(GetDiscoveryTopic(d), 0, true, message)
	token.Wait()
	time.Sleep(common.HADiscoveryDelay)
	d.AvailabilityFunc()
	d.UpdateState()
}
func (d *Sensor) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, "offline")
	token.Wait()
}
func (d *Sensor) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d *Sensor) Initialize() {
	d.PopulateDevice()
	d.AddMessageHandler()
	d.PopulateTopics()
}
func (d *Sensor) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d *Sensor) SetMQTTFields(fields MQTTFields) {
	d.MQTT = fields
}
func (d *Sensor) GetMQTTFields() (fields MQTTFields) {
	return d.MQTT
}
