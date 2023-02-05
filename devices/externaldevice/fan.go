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
type Fan struct {
	AvailabilityMode     *string                         `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate *string                         `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    *string                         `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	AvailabilityFunc     func() string                   `json:"-"`
	CommandTemplate      *string                         `json:"command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	CommandTopic         *string                         `json:"command_topic,omitempty"`    // "The MQTT topic to publish commands to change the fan state."
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
		ViaDevice        *string `json:"via_device,omitempty"`        // "Identifier of a device that routes messages between this device and Home Assistant. Examples of such devices are hubs, or parent devices of a sub-device. This is used to show device topology in Home Assistant."
	} `json:"device,omitempty"` // Device configuration parameters
	EnabledByDefault           *bool                           `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding                   *string                         `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory             *string                         `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                       *string                         `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate     *string                         `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic        *string                         `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	JsonAttributesFunc         func(mqtt.Message, mqtt.Client) `json:"-"`
	Name                       *string                         `json:"name,omitempty"`                         // "The name of the fan."
	ObjectId                   *string                         `json:"object_id,omitempty"`                    // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic                 *bool                           `json:"optimistic,omitempty"`                   // "Flag that defines if fan works in optimistic mode"
	OscillationCommandTemplate *string                         `json:"oscillation_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `oscillation_command_topic`."
	OscillationCommandTopic    *string                         `json:"oscillation_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the oscillation state."
	OscillationCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	OscillationStateTopic      *string                         `json:"oscillation_state_topic,omitempty"` // "The MQTT topic subscribed to receive oscillation state updates."
	OscillationStateFunc       func() string                   `json:"-"`
	OscillationValueTemplate   *string                         `json:"oscillation_value_template,omitempty"`  // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value from the oscillation."
	PayloadAvailable           *string                         `json:"payload_available,omitempty"`           // "The payload that represents the available state."
	PayloadNotAvailable        *string                         `json:"payload_not_available,omitempty"`       // "The payload that represents the unavailable state."
	PayloadOff                 *string                         `json:"payload_off,omitempty"`                 // "The payload that represents the stop state."
	PayloadOn                  *string                         `json:"payload_on,omitempty"`                  // "The payload that represents the running state."
	PayloadOscillationOff      *string                         `json:"payload_oscillation_off,omitempty"`     // "The payload that represents the oscillation off state."
	PayloadOscillationOn       *string                         `json:"payload_oscillation_on,omitempty"`      // "The payload that represents the oscillation on state."
	PayloadResetPercentage     *string                         `json:"payload_reset_percentage,omitempty"`    // "A special payload that resets the `percentage` state attribute to `None` when received at the `percentage_state_topic`."
	PayloadResetPresetMode     *string                         `json:"payload_reset_preset_mode,omitempty"`   // "A special payload that resets the `preset_mode` state attribute to `None` when received at the `preset_mode_state_topic`."
	PercentageCommandTemplate  *string                         `json:"percentage_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `percentage_command_topic`."
	PercentageCommandTopic     *string                         `json:"percentage_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the fan speed state based on a percentage."
	PercentageCommandFunc      func(mqtt.Message, mqtt.Client) `json:"-"`
	PercentageStateTopic       *string                         `json:"percentage_state_topic,omitempty"` // "The MQTT topic subscribed to receive fan speed based on percentage."
	PercentageStateFunc        func() string                   `json:"-"`
	PercentageValueTemplate    *string                         `json:"percentage_value_template,omitempty"`    // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the `percentage` value from the payload received on `percentage_state_topic`."
	PresetModeCommandTemplate  *string                         `json:"preset_mode_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `preset_mode_command_topic`."
	PresetModeCommandTopic     *string                         `json:"preset_mode_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the preset mode."
	PresetModeCommandFunc      func(mqtt.Message, mqtt.Client) `json:"-"`
	PresetModeStateTopic       *string                         `json:"preset_mode_state_topic,omitempty"` // "The MQTT topic subscribed to receive fan speed based on presets."
	PresetModeStateFunc        func() string                   `json:"-"`
	PresetModeValueTemplate    *string                         `json:"preset_mode_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the `preset_mode` value from the payload received on `preset_mode_state_topic`."
	PresetModes                *([]string)                     `json:"preset_modes,omitempty"`               // "List of preset modes this fan is capable of running at. Common examples include `auto`, `smart`, `whoosh`, `eco` and `breeze`."
	Qos                        *int                            `json:"qos,omitempty"`                        // "The maximum QoS level of the state topic."
	Retain                     *bool                           `json:"retain,omitempty"`                     // "If the published message should have the retain flag on or not."
	SpeedRangeMax              *int                            `json:"speed_range_max,omitempty"`            // "The maximum of numeric output range (representing 100 %). The number of speeds within the `speed_range` / `100` will determine the `percentage_step`."
	SpeedRangeMin              *int                            `json:"speed_range_min,omitempty"`            // "The minimum of numeric output range (`off` not included, so `speed_range_min` - `1` represents 0 %). The number of speeds within the speed_range / 100 will determine the `percentage_step`."
	StateTopic                 *string                         `json:"state_topic,omitempty"`                // "The MQTT topic subscribed to receive state updates."
	StateFunc                  func() string                   `json:"-"`
	StateValueTemplate         *string                         `json:"state_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value from the state."
	UniqueId                   *string                         `json:"unique_id,omitempty"`            // "An ID that uniquely identifies this fan. If two fans have the same unique ID, Home Assistant will raise an exception."
	MQTT                       *MQTTFields                     `json:"-"`                              // MQTT configuration parameters
}

func (d *Fan) GetRawId() string {
	return "fan"
}
func (d *Fan) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Fan) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Fan) PopulateDevice() {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &common.MachineID
}
func (d *Fan) UpdateState() {
	if d.AvailabilityTopic != nil {
		state := d.AvailabilityFunc()
		if state != stateStore.Fan.Availability[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Fan.Availability[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.OscillationStateTopic != nil {
		state := d.OscillationStateFunc()
		if state != stateStore.Fan.OscillationState[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.OscillationStateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Fan.OscillationState[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.PercentageStateTopic != nil {
		state := d.PercentageStateFunc()
		if state != stateStore.Fan.PercentageState[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.PercentageStateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Fan.PercentageState[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.PresetModeStateTopic != nil {
		state := d.PresetModeStateFunc()
		if state != stateStore.Fan.PresetModeState[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.PresetModeStateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Fan.PresetModeState[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.StateTopic != nil {
		state := d.StateFunc()
		if state != stateStore.Fan.State[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Fan.State[d.GetUniqueId()] = state
			token.Wait()
		}
	}
}
func (d *Fan) Subscribe() {
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
	if d.OscillationCommandTopic != nil {
		t := c.Subscribe(*d.OscillationCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.PercentageCommandTopic != nil {
		t := c.Subscribe(*d.PercentageCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.PresetModeCommandTopic != nil {
		t := c.Subscribe(*d.PresetModeCommandTopic, 0, d.MQTT.MessageHandler)
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
func (d *Fan) UnSubscribe() {
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
	if d.OscillationCommandTopic != nil {
		t := c.Unsubscribe(*d.OscillationCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.PercentageCommandTopic != nil {
		t := c.Unsubscribe(*d.PercentageCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.PresetModeCommandTopic != nil {
		t := c.Unsubscribe(*d.PresetModeCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d *Fan) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d *Fan) Initialize() {
	if d.Qos == nil {
		d.Qos = new(int)
		*d.Qos = int(common.QoS)
	}
	if d.Retain == nil {
		d.Retain = new(bool)
		*d.Retain = common.RetainClient
	}
	if d.UniqueId == nil {
		d.UniqueId = new(string)
		*d.UniqueId = strcase.ToDelimited(*d.Name, uint8(0x2d))
	}
	d.PopulateDevice()
	d.AddMessageHandler()
	d.PopulateTopics()
}
func (d *Fan) PopulateTopics() {
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
	if d.OscillationCommandFunc != nil {
		d.OscillationCommandTopic = new(string)
		*d.OscillationCommandTopic = GetTopic(d, "oscillation_command_topic")
		store.TopicStore[*d.OscillationCommandTopic] = &d.OscillationCommandFunc
	}
	if d.OscillationStateFunc != nil {
		d.OscillationStateTopic = new(string)
		*d.OscillationStateTopic = GetTopic(d, "oscillation_state_topic")
	}
	if d.PercentageCommandFunc != nil {
		d.PercentageCommandTopic = new(string)
		*d.PercentageCommandTopic = GetTopic(d, "percentage_command_topic")
		store.TopicStore[*d.PercentageCommandTopic] = &d.PercentageCommandFunc
	}
	if d.PercentageStateFunc != nil {
		d.PercentageStateTopic = new(string)
		*d.PercentageStateTopic = GetTopic(d, "percentage_state_topic")
	}
	if d.PresetModeCommandFunc != nil {
		d.PresetModeCommandTopic = new(string)
		*d.PresetModeCommandTopic = GetTopic(d, "preset_mode_command_topic")
		store.TopicStore[*d.PresetModeCommandTopic] = &d.PresetModeCommandFunc
	}
	if d.PresetModeStateFunc != nil {
		d.PresetModeStateTopic = new(string)
		*d.PresetModeStateTopic = GetTopic(d, "preset_mode_state_topic")
	}
	if d.StateFunc != nil {
		d.StateTopic = new(string)
		*d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d *Fan) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Fan) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
