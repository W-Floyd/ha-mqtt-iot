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
	return *d.UniqueId
}
func (d *Switch) PopulateDevice() {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &common.MachineID
}

type Switch struct {
	AvailabilityMode     *string                         `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate *string                         `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    *string                         `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	AvailabilityFunc     func() string                   `json:"-"`
	CommandTopic         *string                         `json:"command_topic,omitempty"` // "The MQTT topic to publish commands to change the switch state."
	CommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
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
	DeviceClass         *string       `json:"device_class,omitempty"`          // "The [type/class](/integrations/switch/#device-class) of the switch to set the icon in the frontend."
	EnabledByDefault    *bool         `json:"enabled_by_default,omitempty"`    // "Flag which defines if the entity should be enabled when first added."
	Encoding            *string       `json:"encoding,omitempty"`              // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory      *string       `json:"entity_category,omitempty"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                *string       `json:"icon,omitempty"`                  // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Name                *string       `json:"name,omitempty"`                  // "The name to use when displaying this switch."
	ObjectId            *string       `json:"object_id,omitempty"`             // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic          *bool         `json:"optimistic,omitempty"`            // "Flag that defines if switch works in optimistic mode."
	PayloadAvailable    *string       `json:"payload_available,omitempty"`     // "The payload that represents the available state."
	PayloadNotAvailable *string       `json:"payload_not_available,omitempty"` // "The payload that represents the unavailable state."
	PayloadOff          *string       `json:"payload_off,omitempty"`           // "The payload that represents `off` state. If specified, will be used for both comparing to the value in the `state_topic` (see `value_template` and `state_off` for details) and sending as `off` command to the `command_topic`."
	PayloadOn           *string       `json:"payload_on,omitempty"`            // "The payload that represents `on` state. If specified, will be used for both comparing to the value in the `state_topic` (see `value_template` and `state_on`  for details) and sending as `on` command to the `command_topic`."
	Qos                 *int          `json:"qos,omitempty"`                   // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain              *bool         `json:"retain,omitempty"`                // "If the published message should have the retain flag on or not."
	StateOff            *string       `json:"state_off,omitempty"`             // "The payload that represents the `off` state. Used when value that represents `off` state in the `state_topic` is different from value that should be sent to the `command_topic` to turn the device `off`."
	StateOn             *string       `json:"state_on,omitempty"`              // "The payload that represents the `on` state. Used when value that represents `on` state in the `state_topic` is different from value that should be sent to the `command_topic` to turn the device `on`."
	StateTopic          *string       `json:"state_topic,omitempty"`           // "The MQTT topic subscribed to receive state updates."
	StateFunc           func() string `json:"-"`
	UniqueId            *string       `json:"unique_id,omitempty"`      // "An ID that uniquely identifies this switch device. If two switches have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate       *string       `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's state from the `state_topic`. To determine the switches's state result of this template will be compared to `state_on` and `state_off`."
	MQTT                *MQTTFields   `json:"-"`
}

func (d *Switch) UpdateState() {
	if d.AvailabilityTopic != nil {
		state := d.AvailabilityFunc()
		if state != stateStore.Switch.Availability[*d.UniqueId] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Switch.Availability[*d.UniqueId] = state
			token.Wait()
		}
	}
	if d.StateTopic != nil {
		state := d.StateFunc()
		if state != stateStore.Switch.State[*d.UniqueId] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Switch.State[*d.UniqueId] = state
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
	if d.CommandTopic != nil {
		t := c.Subscribe(*d.CommandTopic, 0, d.MQTT.MessageHandler)
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
func (d *Switch) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "offline")
	token.Wait()
	if d.CommandTopic != nil {
		t := c.Unsubscribe(*d.CommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d *Switch) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d *Switch) Initialize() {
	if d.Qos == nil {
		d.Qos = new(int)
		*d.Qos = int(common.QoS)
	}
	if d.Retain == nil {
		d.Retain = new(bool)
		*d.Retain = common.Retain
	}
	d.PopulateDevice()
	d.AddMessageHandler()
	d.PopulateTopics()
}
func (d *Switch) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = new(string)
		*d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = new(string)
		*d.CommandTopic = GetTopic(d, "command_topic")
		store.TopicStore[*d.CommandTopic] = &d.CommandFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = new(string)
		*d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d *Switch) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Switch) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
