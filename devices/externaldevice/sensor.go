package ExternalDevice

import (
	"encoding/json"
	common "github.com/W-Floyd/ha-mqtt-iot/common"
	store "github.com/W-Floyd/ha-mqtt-iot/store"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	strcase "github.com/iancoleman/strcase"
	"log"
	"time"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Sensor struct {
	AvailabilityMode     *string       `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate *string       `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    *string       `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates."
	AvailabilityFunc     func() string `json:"-"`
	Device               struct {
		ConfigurationUrl *string `json:"configuration_url,omitempty"` // "A link to the webpage that can manage the configuration of this device. Can be either an HTTP or HTTPS link."
		Connections      *string `json:"connections,omitempty"`       // "A list of connections of the device to the outside world as a list of tuples `[connection_type, connection_identifier]`. For example the MAC address of a network interface: `\"connections\": [[\"mac\", \"02:5b:26:a8:dc:12\"]]`."
		Identifiers      *string `json:"identifiers,omitempty"`       // "A list of IDs that uniquely identify the device. For example a serial number."
		Manufacturer     *string `json:"manufacturer,omitempty"`      // "The manufacturer of the device."
		Model            *string `json:"model,omitempty"`             // "The model of the device."
		Name             *string `json:"name,omitempty"`              // "The name of the device."
		SuggestedArea    *string `json:"suggested_area,omitempty"`    // "Suggest an area if the device isn’t in one yet."
		SwVersion        *string `json:"sw_version,omitempty"`        // "The firmware version of the device."
		Viadevice        *string `json:"viadevice,omitempty"`         // null
	} `json:"device,omitempty"`
	DeviceClass            *string                         `json:"device_class,omitempty"`             // "The [type/class](/integrations/sensor/#device-class) of the sensor to set the icon in the frontend."
	EnabledByDefault       *bool                           `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string                         `json:"encoding,omitempty"`                 // "The encoding of the payloads received. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string                         `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	ExpireAfter            *int                            `json:"expire_after,omitempty"`             // "Defines the number of seconds after the sensor's state expires, if it's not updated. After expiry, the sensor's state becomes `unavailable`."
	ForceUpdate            *bool                           `json:"force_update,omitempty"`             // "Sends update events even if the value hasn't changed. Useful if you want to have meaningful value graphs in history."
	Icon                   *string                         `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string                         `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributesTopic    *string                         `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Implies `force_update` of the current sensor state when a message is received on this topic."
	JsonAttributesFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	LastResetValueTemplate *string                         `json:"last_reset_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the last_reset. Available variables: `entity_id`. The `entity_id` can be used to reference the entity's attributes."
	Name                   *string                         `json:"name,omitempty"`                      // "The name of the MQTT sensor."
	ObjectId               *string                         `json:"object_id,omitempty"`                 // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable       *string                         `json:"payload_available,omitempty"`         // "The payload that represents the available state."
	PayloadNotAvailable    *string                         `json:"payload_not_available,omitempty"`     // "The payload that represents the unavailable state."
	Qos                    *int                            `json:"qos,omitempty"`                       // "The maximum QoS level of the state topic."
	StateClass             *string                         `json:"state_class,omitempty"`               // "The [state_class](https://developers.home-assistant.io/docs/core/entity/sensor#available-state-classes) of the sensor."
	StateTopic             *string                         `json:"state_topic,omitempty"`               // "The MQTT topic subscribed to receive sensor values."
	StateFunc              func() string                   `json:"-"`
	UniqueId               *string                         `json:"unique_id,omitempty"`           // "An ID that uniquely identifies this sensor. If two sensors have the same unique ID, Home Assistant will raise an exception."
	UnitOfMeasurement      *string                         `json:"unit_of_measurement,omitempty"` // "Defines the units of measurement of the sensor, if any."
	ValueTemplate          *string                         `json:"value_template,omitempty"`      // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the value. Available variables: `entity_id`. The `entity_id` can be used to reference the entity's attributes. If the template throws an error, the current state will be used instead."
	MQTT                   *MQTTFields                     `json:"-"`
}

func (d *Sensor) GetRawId() string {
	return "sensor"
}
func (d *Sensor) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Sensor) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Sensor) PopulateDevice() {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &common.MachineID
}
func (d *Sensor) UpdateState() {
	if d.AvailabilityTopic != nil {
		state := d.AvailabilityFunc()
		if state != stateStore.Sensor.Availability[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, byte(*d.Qos), common.Retain, state)
			stateStore.Sensor.Availability[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.StateTopic != nil {
		state := d.StateFunc()
		if state != stateStore.Sensor.State[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, byte(*d.Qos), common.Retain, state)
			stateStore.Sensor.State[d.GetUniqueId()] = state
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
	if d.JsonAttributesTopic != nil {
		t := c.Subscribe(*d.JsonAttributesTopic, 0, d.MQTT.MessageHandler)
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
func (d *Sensor) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "offline")
	token.Wait()
	if d.JsonAttributesTopic != nil {
		t := c.Unsubscribe(*d.JsonAttributesTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d *Sensor) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d *Sensor) Initialize() {
	if d.Qos == nil {
		d.Qos = new(int)
		*d.Qos = int(common.QoS)
	}
	if d.UniqueId == nil {
		d.UniqueId = new(string)
		*d.UniqueId = strcase.ToDelimited(*d.Name, uint8(0x2d))
	}
	d.PopulateDevice()
	d.AddMessageHandler()
	d.PopulateTopics()
}
func (d *Sensor) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = new(string)
		*d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.JsonAttributesFunc != nil {
		d.JsonAttributesTopic = new(string)
		*d.JsonAttributesTopic = GetTopic(d, "json_attributes_topic")
		store.TopicStore[*d.JsonAttributesTopic] = &d.JsonAttributesFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = new(string)
		*d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d *Sensor) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Sensor) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
