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
type Event struct {
	AvailabilityMode     *string       `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate *string       `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    *string       `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	AvailabilityFunc     func() string `json:"-"`                               // Function for availability
	Device               struct {
		ConfigurationUrl *string `json:"configuration_url,omitempty"` // "A link to the webpage that can manage the configuration of this device. Can be either an HTTP or HTTPS link."
		Connections      *string `json:"connections,omitempty"`       // "A list of connections of the device to the outside world as a list of tuples `[connection_type, connection_identifier]`. For example the MAC address of a network interface: `\"connections\": [[\"mac\", \"02:5b:26:a8:dc:12\"]]`."
		Identifiers      *string `json:"identifiers,omitempty"`       // "A list of IDs that uniquely identify the device. For example a serial number."
		Manufacturer     *string `json:"manufacturer,omitempty"`      // "The manufacturer of the device."
		Model            *string `json:"model,omitempty"`             // "The model of the device."
		Name             *string `json:"name,omitempty"`              // "The name of the device."
		SuggestedArea    *string `json:"suggested_area,omitempty"`    // "Suggest an area if the device isn’t in one yet."
		SwVersion        *string `json:"sw_version,omitempty"`        // "The firmware version of the device."
		ViaDevice        *string `json:"via_device,omitempty"`        // "Identifier of a device that routes messages between this device and Home Assistant. Examples of such devices are hubs, or parent devices of a sub-device. This is used to show device topology in Home Assistant."
	} `json:"device,omitempty"` // Device configuration parameters
	DeviceClass            *string                         `json:"device_class,omitempty"`             // "The [type/class](/integrations/event/#device-class) of the event to set the icon in the frontend. The `device_class` can be `null`."
	EnabledByDefault       *bool                           `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string                         `json:"encoding,omitempty"`                 // "The encoding of the published messages."
	EntityCategory         *string                         `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	EventTypes             *([]string)                     `json:"event_types,omitempty"`              // "A list of valid `event_type` strings."
	Icon                   *string                         `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string                         `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic    *string                         `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	JsonAttributesFunc     func(mqtt.Message, mqtt.Client) `json:"-"`                                  // Function for json attributes
	Name                   *string                         `json:"name,omitempty"`                     // "The name to use when displaying this event."
	ObjectId               *string                         `json:"object_id,omitempty"`                // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable       *string                         `json:"payload_available,omitempty"`        // "The payload that represents the available state."
	PayloadNotAvailable    *string                         `json:"payload_not_available,omitempty"`    // "The payload that represents the unavailable state."
	Qos                    *int                            `json:"qos,omitempty"`                      // "The maximum QoS level to be used when receiving and publishing messages."
	StateTopic             *string                         `json:"state_topic,omitempty"`              // "The MQTT topic subscribed to receive JSON event payloads. The JSON payload should contain the `event_type` element. The event type should be one of the configured `event_types`. Note that replayed retained messages will be discarded."
	StateFunc              func() string                   `json:"-"`                                  // Function for state
	UniqueId               *string                         `json:"unique_id,omitempty"`                // "An ID that uniquely identifies this event entity. If two events have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          *string                         `json:"value_template,omitempty"`           // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the value and render it to a valid JSON event payload. If the template throws an error, the current state will be used instead."
	MQTT                   *MQTTFields                     `json:"-"`                                  // MQTT configuration parameters
}

func (d *Event) GetRawId() string {
	return "event"
}
func (d *Event) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Event) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Event) PopulateDevice() {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &common.MachineID
}
func (d *Event) UpdateState() {
	if d.AvailabilityTopic != nil {
		state := d.AvailabilityFunc()
		if state != stateStore.Event.Availability[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, byte(*d.Qos), common.Retain, state)
			stateStore.Event.Availability[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.StateTopic != nil {
		state := d.StateFunc()
		if state != stateStore.Event.State[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, byte(*d.Qos), common.Retain, state)
			stateStore.Event.State[d.GetUniqueId()] = state
			token.Wait()
		}
	}
}
func (d *Event) Subscribe() {
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
func (d *Event) UnSubscribe() {
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
func (d *Event) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d *Event) Initialize() {
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
func (d *Event) PopulateTopics() {
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
func (d *Event) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Event) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
