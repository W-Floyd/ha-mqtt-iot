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

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
type Cover struct {
	AvailabilityMode     *string                         `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate *string                         `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    *string                         `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to to receive birth and LWT messages from the MQTT cover device. If an `availability` topic is not defined, the cover availability state will always be `available`. If an `availability` topic is defined, the cover availability state will be `unavailable` by default. Must not be used together with `availability`."
	AvailabilityFunc     func() string                   `json:"-"`
	CommandTopic         *string                         `json:"command_topic,omitempty"` // "The MQTT topic to publish commands to control the cover."
	CommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
	Device               struct {
		ConfigurationUrl *string `json:"configuration_url,omitempty"` // "A link to the webpage that can manage the configuration of this device. Can be either an HTTP or HTTPS link."
		Connections      *string `json:"connections,omitempty"`       // "A list of connections of the device to the outside world as a list of tuples `[connection_type, connection_identifier]`. For example the MAC address of a network interface: `\"connections\": [\"mac\", \"02:5b:26:a8:dc:12\"]`."
		Identifiers      *string `json:"identifiers,omitempty"`       // "A list of IDs that uniquely identify the device. For example a serial number."
		Manufacturer     *string `json:"manufacturer,omitempty"`      // "The manufacturer of the device."
		Model            *string `json:"model,omitempty"`             // "The model of the device."
		Name             *string `json:"name,omitempty"`              // "The name of the device."
		SuggestedArea    *string `json:"suggested_area,omitempty"`    // "Suggest an area if the device isn’t in one yet."
		SwVersion        *string `json:"sw_version,omitempty"`        // "The firmware version of the device."
		Viadevice        *string `json:"viadevice,omitempty"`         // null
	} `json:"device,omitempty"`
	DeviceClass            *string                         `json:"device_class,omitempty"`             // "Sets the [class of the device](/integrations/cover/), changing the device state and icon that is displayed on the frontend."
	EnabledByDefault       *bool                           `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string                         `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string                         `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   *string                         `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string                         `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic    *string                         `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	JsonAttributesFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	Name                   *string                         `json:"name,omitempty"`                  // "The name of the cover."
	ObjectId               *string                         `json:"object_id,omitempty"`             // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic             *bool                           `json:"optimistic,omitempty"`            // "Flag that defines if switch works in optimistic mode."
	PayloadAvailable       *string                         `json:"payload_available,omitempty"`     // "The payload that represents the online state."
	PayloadClose           *string                         `json:"payload_close,omitempty"`         // "The command payload that closes the cover."
	PayloadNotAvailable    *string                         `json:"payload_not_available,omitempty"` // "The payload that represents the offline state."
	PayloadOpen            *string                         `json:"payload_open,omitempty"`          // "The command payload that opens the cover."
	PayloadStop            *string                         `json:"payload_stop,omitempty"`          // "The command payload that stops the cover."
	PositionClosed         *int                            `json:"position_closed,omitempty"`       // "Number which represents closed position."
	PositionOpen           *int                            `json:"position_open,omitempty"`         // "Number which represents open position."
	PositionTemplate       *string                         `json:"position_template,omitempty"`     // "Defines a [template](/topics/templating/) that can be used to extract the payload for the `position_topic` topic. Within the template the following variables are available: `entity_id`, `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function;"
	PositionTopic          *string                         `json:"position_topic,omitempty"`        // "The MQTT topic subscribed to receive cover position messages."
	PositionFunc           func() string                   `json:"-"`
	Qos                    *int                            `json:"qos,omitempty"`                   // "The maximum QoS level to be used when receiving and publishing messages."
	Retain                 *bool                           `json:"retain,omitempty"`                // "Defines if published messages should have the retain flag set."
	SetPositionTemplate    *string                         `json:"set_position_template,omitempty"` // "Defines a [template](/topics/templating/) to define the position to be sent to the `set_position_topic` topic. Incoming position value is available for use in the template `{% raw %}{{ position }}{% endraw %}`. Within the template the following variables are available: `entity_id`, `position`, the target position in percent; `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function;"
	SetPositionTopic       *string                         `json:"set_position_topic,omitempty"`    // "The MQTT topic to publish position commands to. You need to set position_topic as well if you want to use position topic. Use template if position topic wants different values than within range `position_closed` - `position_open`. If template is not defined and `position_closed != 100` and `position_open != 0` then proper position value is calculated from percentage position."
	SetPositionFunc        func(mqtt.Message, mqtt.Client) `json:"-"`
	StateClosed            *string                         `json:"state_closed,omitempty"`  // "The payload that represents the closed state."
	StateClosing           *string                         `json:"state_closing,omitempty"` // "The payload that represents the closing state."
	StateOpen              *string                         `json:"state_open,omitempty"`    // "The payload that represents the open state."
	StateOpening           *string                         `json:"state_opening,omitempty"` // "The payload that represents the opening state."
	StateStopped           *string                         `json:"state_stopped,omitempty"` // "The payload that represents the stopped state (for covers that do not report `open`/`closed` state)."
	StateTopic             *string                         `json:"state_topic,omitempty"`   // "The MQTT topic subscribed to receive cover state messages. State topic can only read (`open`, `opening`, `closed`, `closing` or `stopped`) state."
	StateFunc              func() string                   `json:"-"`
	TiltClosedValue        *int                            `json:"tilt_closed_value,omitempty"`     // "The value that will be sent on a `close_cover_tilt` command."
	TiltCommandTemplate    *string                         `json:"tilt_command_template,omitempty"` // "Defines a [template](/topics/templating/) that can be used to extract the payload for the `tilt_command_topic` topic. Within the template the following variables are available: `entity_id`, `tilt_position`, the target tilt position in percent; `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function;"
	TiltCommandTopic       *string                         `json:"tilt_command_topic,omitempty"`    // "The MQTT topic to publish commands to control the cover tilt."
	TiltCommandFunc        func(mqtt.Message, mqtt.Client) `json:"-"`
	TiltMax                *int                            `json:"tilt_max,omitempty"`             // "The maximum tilt value."
	TiltMin                *int                            `json:"tilt_min,omitempty"`             // "The minimum tilt value."
	TiltOpenedValue        *int                            `json:"tilt_opened_value,omitempty"`    // "The value that will be sent on an `open_cover_tilt` command."
	TiltOptimistic         *bool                           `json:"tilt_optimistic,omitempty"`      // "Flag that determines if tilt works in optimistic mode."
	TiltStatusTemplate     *string                         `json:"tilt_status_template,omitempty"` // "Defines a [template](/topics/templating/) that can be used to extract the payload for the `tilt_status_topic` topic. Within the template the following variables are available: `entity_id`, `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function;"
	TiltStatusTopic        *string                         `json:"tilt_status_topic,omitempty"`    // "The MQTT topic subscribed to receive tilt status update values."
	TiltStatusFunc         func() string                   `json:"-"`
	UniqueId               *string                         `json:"unique_id,omitempty"`      // "An ID that uniquely identifies this cover. If two covers have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          *string                         `json:"value_template,omitempty"` // "Defines a [template](/topics/templating/) that can be used to extract the payload for the `state_topic` topic."
	MQTT                   *MQTTFields                     `json:"-"`
}

func (d *Cover) GetRawId() string {
	return "cover"
}
func (d *Cover) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Cover) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Cover) PopulateDevice() {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &common.MachineID
}
func (d *Cover) UpdateState() {
	if d.AvailabilityTopic != nil {
		state := d.AvailabilityFunc()
		if state != stateStore.Cover.Availability[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Cover.Availability[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.PositionTopic != nil {
		state := d.PositionFunc()
		if state != stateStore.Cover.Position[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.PositionTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Cover.Position[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.StateTopic != nil {
		state := d.StateFunc()
		if state != stateStore.Cover.State[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Cover.State[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.TiltStatusTopic != nil {
		state := d.TiltStatusFunc()
		if state != stateStore.Cover.TiltStatus[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.TiltStatusTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Cover.TiltStatus[d.GetUniqueId()] = state
			token.Wait()
		}
	}
}
func (d *Cover) Subscribe() {
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
	if d.JsonAttributesTopic != nil {
		t := c.Subscribe(*d.JsonAttributesTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.SetPositionTopic != nil {
		t := c.Subscribe(*d.SetPositionTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TiltCommandTopic != nil {
		t := c.Subscribe(*d.TiltCommandTopic, 0, d.MQTT.MessageHandler)
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
func (d *Cover) UnSubscribe() {
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
	if d.JsonAttributesTopic != nil {
		t := c.Unsubscribe(*d.JsonAttributesTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.SetPositionTopic != nil {
		t := c.Unsubscribe(*d.SetPositionTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TiltCommandTopic != nil {
		t := c.Unsubscribe(*d.TiltCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d *Cover) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d *Cover) Initialize() {
	if d.Qos == nil {
		d.Qos = new(int)
		*d.Qos = int(common.QoS)
	}
	if d.Retain == nil {
		d.Retain = new(bool)
		*d.Retain = common.Retain
	}
	if d.UniqueId == nil {
		d.UniqueId = new(string)
		*d.UniqueId = strcase.ToDelimited(*d.Name, uint8(0x2d))
	}
	d.PopulateDevice()
	d.AddMessageHandler()
	d.PopulateTopics()
}
func (d *Cover) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = new(string)
		*d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = new(string)
		*d.CommandTopic = GetTopic(d, "command_topic")
		store.TopicStore[*d.CommandTopic] = &d.CommandFunc
	}
	if d.JsonAttributesFunc != nil {
		d.JsonAttributesTopic = new(string)
		*d.JsonAttributesTopic = GetTopic(d, "json_attributes_topic")
		store.TopicStore[*d.JsonAttributesTopic] = &d.JsonAttributesFunc
	}
	if d.PositionFunc != nil {
		d.PositionTopic = new(string)
		*d.PositionTopic = GetTopic(d, "position_topic")
	}
	if d.SetPositionFunc != nil {
		d.SetPositionTopic = new(string)
		*d.SetPositionTopic = GetTopic(d, "set_position_topic")
		store.TopicStore[*d.SetPositionTopic] = &d.SetPositionFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = new(string)
		*d.StateTopic = GetTopic(d, "state_topic")
	}
	if d.TiltCommandFunc != nil {
		d.TiltCommandTopic = new(string)
		*d.TiltCommandTopic = GetTopic(d, "tilt_command_topic")
		store.TopicStore[*d.TiltCommandTopic] = &d.TiltCommandFunc
	}
	if d.TiltStatusFunc != nil {
		d.TiltStatusTopic = new(string)
		*d.TiltStatusTopic = GetTopic(d, "tilt_status_topic")
	}
}
func (d *Cover) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Cover) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
