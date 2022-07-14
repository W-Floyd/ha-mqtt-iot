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
func (d *Climate) GetRawId() string {
	return "climate"
}
func (d *Climate) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Climate) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Climate) PopulateDevice() {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &common.MachineID
}

type Climate struct {
	ActionTemplate             *string                         `json:"action_template,omitempty"` // "A template to render the value received on the `action_topic` with."
	ActionTopic                *string                         `json:"action_topic,omitempty"`    // "The MQTT topic to subscribe for changes of the current action. If this is set, the climate graph uses the value received as data source. Valid values: `off`, `heating`, `cooling`, `drying`, `idle`, `fan`."
	ActionFunc                 func(mqtt.Message, mqtt.Client) `json:"-"`
	AuxCommandTopic            *string                         `json:"aux_command_topic,omitempty"` // "The MQTT topic to publish commands to switch auxiliary heat."
	AuxCommandFunc             func(mqtt.Message, mqtt.Client) `json:"-"`
	AuxStateTemplate           *string                         `json:"aux_state_template,omitempty"` // "A template to render the value received on the `aux_state_topic` with."
	AuxStateTopic              *string                         `json:"aux_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes of the auxiliary heat mode. If this is not set, the auxiliary heat mode works in optimistic mode (see below)."
	AuxStateFunc               func() string                   `json:"-"`
	AvailabilityMode           *string                         `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate       *string                         `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic          *string                         `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	AvailabilityFunc           func() string                   `json:"-"`
	CurrentTemperatureTemplate *string                         `json:"current_temperature_template,omitempty"` // "A template with which the value received on `current_temperature_topic` will be rendered."
	CurrentTemperatureTopic    *string                         `json:"current_temperature_topic,omitempty"`    // "The MQTT topic on which to listen for the current temperature."
	CurrentTemperatureFunc     func() string                   `json:"-"`
	Device                     struct {
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
	EnabledByDefault               *bool                           `json:"enabled_by_default,omitempty"`        // "Flag which defines if the entity should be enabled when first added."
	Encoding                       *string                         `json:"encoding,omitempty"`                  // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory                 *string                         `json:"entity_category,omitempty"`           // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	FanModeCommandTemplate         *string                         `json:"fan_mode_command_template,omitempty"` // "A template to render the value sent to the `fan_mode_command_topic` with."
	FanModeCommandTopic            *string                         `json:"fan_mode_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the fan mode."
	FanModeCommandFunc             func(mqtt.Message, mqtt.Client) `json:"-"`
	FanModeStateTemplate           *string                         `json:"fan_mode_state_template,omitempty"` // "A template to render the value received on the `fan_mode_state_topic` with."
	FanModeStateTopic              *string                         `json:"fan_mode_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes of the HVAC fan mode. If this is not set, the fan mode works in optimistic mode (see below)."
	FanModeStateFunc               func() string                   `json:"-"`
	FanModes                       *([]string)                     `json:"fan_modes,omitempty"`             // "A list of supported fan modes."
	Icon                           *string                         `json:"icon,omitempty"`                  // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Initial                        *int                            `json:"initial,omitempty"`               // "Set the initial target temperature."
	MaxTemp                        *float64                        `json:"max_temp,omitempty"`              // "Maximum set point available."
	MinTemp                        *float64                        `json:"min_temp,omitempty"`              // "Minimum set point available."
	ModeCommandTemplate            *string                         `json:"mode_command_template,omitempty"` // "A template to render the value sent to the `mode_command_topic` with."
	ModeCommandTopic               *string                         `json:"mode_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the HVAC operation mode."
	ModeCommandFunc                func(mqtt.Message, mqtt.Client) `json:"-"`
	ModeStateTemplate              *string                         `json:"mode_state_template,omitempty"` // "A template to render the value received on the `mode_state_topic` with."
	ModeStateTopic                 *string                         `json:"mode_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes of the HVAC operation mode. If this is not set, the operation mode works in optimistic mode (see below)."
	ModeStateFunc                  func() string                   `json:"-"`
	Modes                          *([]string)                     `json:"modes,omitempty"`                 // "A list of supported modes. Needs to be a subset of the default values."
	Name                           *string                         `json:"name,omitempty"`                  // "The name of the HVAC."
	ObjectId                       *string                         `json:"object_id,omitempty"`             // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable               *string                         `json:"payload_available,omitempty"`     // "The payload that represents the available state."
	PayloadNotAvailable            *string                         `json:"payload_not_available,omitempty"` // "The payload that represents the unavailable state."
	PayloadOff                     *string                         `json:"payload_off,omitempty"`           // "The payload that represents disabled state."
	PayloadOn                      *string                         `json:"payload_on,omitempty"`            // "The payload that represents enabled state."
	PowerCommandTopic              *string                         `json:"power_command_topic,omitempty"`   // "The MQTT topic to publish commands to change the power state. This is useful if your device has a separate power toggle in addition to mode."
	PowerCommandFunc               func(mqtt.Message, mqtt.Client) `json:"-"`
	Precision                      *float64                        `json:"precision,omitempty"`                    // "The desired precision for this device. Can be used to match your actual thermostat's precision. Supported values are `0.1`, `0.5` and `1.0`."
	PresetModeCommandTemplate      *string                         `json:"preset_mode_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `preset_mode_command_topic`."
	PresetModeCommandTopic         *string                         `json:"preset_mode_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the preset mode."
	PresetModeCommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
	PresetModeStateTopic           *string                         `json:"preset_mode_state_topic,omitempty"` // "The MQTT topic subscribed to receive climate speed based on presets. When preset 'none' is received or `None` the `preset_mode` will be reset."
	PresetModeStateFunc            func() string                   `json:"-"`
	PresetModeValueTemplate        *string                         `json:"preset_mode_value_template,omitempty"`  // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the `preset_mode` value from the payload received on `preset_mode_state_topic`."
	PresetModes                    *([]string)                     `json:"preset_modes,omitempty"`                // "List of preset modes this climate is supporting. Common examples include `eco`, `away`, `boost`, `comfort`, `home`, `sleep` and `activity`."
	Qos                            *int                            `json:"qos,omitempty"`                         // "The maximum QoS level to be used when receiving and publishing messages."
	Retain                         *bool                           `json:"retain,omitempty"`                      // "Defines if published messages should have the retain flag set."
	SwingModeCommandTemplate       *string                         `json:"swing_mode_command_template,omitempty"` // "A template to render the value sent to the `swing_mode_command_topic` with."
	SwingModeCommandTopic          *string                         `json:"swing_mode_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the swing mode."
	SwingModeCommandFunc           func(mqtt.Message, mqtt.Client) `json:"-"`
	SwingModeStateTemplate         *string                         `json:"swing_mode_state_template,omitempty"` // "A template to render the value received on the `swing_mode_state_topic` with."
	SwingModeStateTopic            *string                         `json:"swing_mode_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes of the HVAC swing mode. If this is not set, the swing mode works in optimistic mode (see below)."
	SwingModeStateFunc             func() string                   `json:"-"`
	SwingModes                     *([]string)                     `json:"swing_modes,omitempty"`                  // "A list of supported swing modes."
	TempStep                       *float64                        `json:"temp_step,omitempty"`                    // "Step size for temperature set point."
	TemperatureCommandTemplate     *string                         `json:"temperature_command_template,omitempty"` // "A template to render the value sent to the `temperature_command_topic` with."
	TemperatureCommandTopic        *string                         `json:"temperature_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the target temperature."
	TemperatureCommandFunc         func(mqtt.Message, mqtt.Client) `json:"-"`
	TemperatureHighCommandTemplate *string                         `json:"temperature_high_command_template,omitempty"` // "A template to render the value sent to the `temperature_high_command_topic` with."
	TemperatureHighCommandTopic    *string                         `json:"temperature_high_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the high target temperature."
	TemperatureHighCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	TemperatureHighStateTemplate   *string                         `json:"temperature_high_state_template,omitempty"` // "A template to render the value received on the `temperature_high_state_topic` with."
	TemperatureHighStateTopic      *string                         `json:"temperature_high_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes in the target high temperature. If this is not set, the target high temperature works in optimistic mode (see below)."
	TemperatureHighStateFunc       func() string                   `json:"-"`
	TemperatureLowCommandTemplate  *string                         `json:"temperature_low_command_template,omitempty"` // "A template to render the value sent to the `temperature_low_command_topic` with."
	TemperatureLowCommandTopic     *string                         `json:"temperature_low_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the target low temperature."
	TemperatureLowCommandFunc      func(mqtt.Message, mqtt.Client) `json:"-"`
	TemperatureLowStateTemplate    *string                         `json:"temperature_low_state_template,omitempty"` // "A template to render the value received on the `temperature_low_state_topic` with."
	TemperatureLowStateTopic       *string                         `json:"temperature_low_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes in the target low temperature. If this is not set, the target low temperature works in optimistic mode (see below)."
	TemperatureLowStateFunc        func() string                   `json:"-"`
	TemperatureStateTemplate       *string                         `json:"temperature_state_template,omitempty"` // "A template to render the value received on the `temperature_state_topic` with."
	TemperatureStateTopic          *string                         `json:"temperature_state_topic,omitempty"`    // "The MQTT topic to subscribe for changes in the target temperature. If this is not set, the target temperature works in optimistic mode (see below)."
	TemperatureStateFunc           func() string                   `json:"-"`
	TemperatureUnit                *string                         `json:"temperature_unit,omitempty"` // "Defines the temperature unit of the device, `C` or `F`. If this is not set, the temperature unit is set to the system temperature unit."
	UniqueId                       *string                         `json:"unique_id,omitempty"`        // "An ID that uniquely identifies this HVAC device. If two HVAC devices have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate                  *string                         `json:"value_template,omitempty"`   // "Default template to render the payloads on *all* `*_state_topic`s with."
	MQTT                           *MQTTFields                     `json:"-"`
}

func (d *Climate) UpdateState() {
	if d.AuxStateTopic != nil {
		state := d.AuxStateFunc()
		if state != stateStore.Climate.AuxState[*d.UniqueId] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			c := *d.MQTT.Client
			token := c.Publish(*d.AuxStateTopic, common.QoS, common.Retain, state)
			stateStore.Climate.AuxState[*d.UniqueId] = state
			token.Wait()
		}
	}
	if d.AvailabilityTopic != nil {
		state := d.AvailabilityFunc()
		if state != stateStore.Climate.Availability[*d.UniqueId] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			c := *d.MQTT.Client
			token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, state)
			stateStore.Climate.Availability[*d.UniqueId] = state
			token.Wait()
		}
	}
	if d.CurrentTemperatureTopic != nil {
		state := d.CurrentTemperatureFunc()
		if state != stateStore.Climate.CurrentTemperature[*d.UniqueId] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			c := *d.MQTT.Client
			token := c.Publish(*d.CurrentTemperatureTopic, common.QoS, common.Retain, state)
			stateStore.Climate.CurrentTemperature[*d.UniqueId] = state
			token.Wait()
		}
	}
	if d.FanModeStateTopic != nil {
		state := d.FanModeStateFunc()
		if state != stateStore.Climate.FanModeState[*d.UniqueId] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			c := *d.MQTT.Client
			token := c.Publish(*d.FanModeStateTopic, common.QoS, common.Retain, state)
			stateStore.Climate.FanModeState[*d.UniqueId] = state
			token.Wait()
		}
	}
	if d.ModeStateTopic != nil {
		state := d.ModeStateFunc()
		if state != stateStore.Climate.ModeState[*d.UniqueId] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			c := *d.MQTT.Client
			token := c.Publish(*d.ModeStateTopic, common.QoS, common.Retain, state)
			stateStore.Climate.ModeState[*d.UniqueId] = state
			token.Wait()
		}
	}
	if d.PresetModeStateTopic != nil {
		state := d.PresetModeStateFunc()
		if state != stateStore.Climate.PresetModeState[*d.UniqueId] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			c := *d.MQTT.Client
			token := c.Publish(*d.PresetModeStateTopic, common.QoS, common.Retain, state)
			stateStore.Climate.PresetModeState[*d.UniqueId] = state
			token.Wait()
		}
	}
	if d.SwingModeStateTopic != nil {
		state := d.SwingModeStateFunc()
		if state != stateStore.Climate.SwingModeState[*d.UniqueId] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			c := *d.MQTT.Client
			token := c.Publish(*d.SwingModeStateTopic, common.QoS, common.Retain, state)
			stateStore.Climate.SwingModeState[*d.UniqueId] = state
			token.Wait()
		}
	}
	if d.TemperatureHighStateTopic != nil {
		state := d.TemperatureHighStateFunc()
		if state != stateStore.Climate.TemperatureHighState[*d.UniqueId] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			c := *d.MQTT.Client
			token := c.Publish(*d.TemperatureHighStateTopic, common.QoS, common.Retain, state)
			stateStore.Climate.TemperatureHighState[*d.UniqueId] = state
			token.Wait()
		}
	}
	if d.TemperatureLowStateTopic != nil {
		state := d.TemperatureLowStateFunc()
		if state != stateStore.Climate.TemperatureLowState[*d.UniqueId] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			c := *d.MQTT.Client
			token := c.Publish(*d.TemperatureLowStateTopic, common.QoS, common.Retain, state)
			stateStore.Climate.TemperatureLowState[*d.UniqueId] = state
			token.Wait()
		}
	}
	if d.TemperatureStateTopic != nil {
		state := d.TemperatureStateFunc()
		if state != stateStore.Climate.TemperatureState[*d.UniqueId] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			c := *d.MQTT.Client
			token := c.Publish(*d.TemperatureStateTopic, common.QoS, common.Retain, state)
			stateStore.Climate.TemperatureState[*d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d *Climate) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	if *d.ActionTopic != "" {
		t := c.Subscribe(*d.ActionTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if *d.AuxCommandTopic != "" {
		t := c.Subscribe(*d.AuxCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if *d.FanModeCommandTopic != "" {
		t := c.Subscribe(*d.FanModeCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if *d.ModeCommandTopic != "" {
		t := c.Subscribe(*d.ModeCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if *d.PowerCommandTopic != "" {
		t := c.Subscribe(*d.PowerCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if *d.PresetModeCommandTopic != "" {
		t := c.Subscribe(*d.PresetModeCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if *d.SwingModeCommandTopic != "" {
		t := c.Subscribe(*d.SwingModeCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if *d.TemperatureCommandTopic != "" {
		t := c.Subscribe(*d.TemperatureCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if *d.TemperatureHighCommandTopic != "" {
		t := c.Subscribe(*d.TemperatureHighCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if *d.TemperatureLowCommandTopic != "" {
		t := c.Subscribe(*d.TemperatureLowCommandTopic, 0, d.MQTT.MessageHandler)
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
func (d *Climate) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "offline")
	token.Wait()
	if *d.ActionTopic != "" {
		t := c.Unsubscribe(*d.ActionTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if *d.AuxCommandTopic != "" {
		t := c.Unsubscribe(*d.AuxCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if *d.FanModeCommandTopic != "" {
		t := c.Unsubscribe(*d.FanModeCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if *d.ModeCommandTopic != "" {
		t := c.Unsubscribe(*d.ModeCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if *d.PowerCommandTopic != "" {
		t := c.Unsubscribe(*d.PowerCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if *d.PresetModeCommandTopic != "" {
		t := c.Unsubscribe(*d.PresetModeCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if *d.SwingModeCommandTopic != "" {
		t := c.Unsubscribe(*d.SwingModeCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if *d.TemperatureCommandTopic != "" {
		t := c.Unsubscribe(*d.TemperatureCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if *d.TemperatureHighCommandTopic != "" {
		t := c.Unsubscribe(*d.TemperatureHighCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if *d.TemperatureLowCommandTopic != "" {
		t := c.Unsubscribe(*d.TemperatureLowCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d *Climate) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d *Climate) Initialize() {
	*d.Retain = false
	d.PopulateDevice()
	d.AddMessageHandler()
	d.PopulateTopics()
}
func (d *Climate) PopulateTopics() {
	if d.ActionFunc != nil {
		d.ActionTopic = new(string)
		*d.ActionTopic = GetTopic(d, "action_topic")
		store.TopicStore[*d.ActionTopic] = &d.ActionFunc
	}
	if d.AuxCommandFunc != nil {
		d.AuxCommandTopic = new(string)
		*d.AuxCommandTopic = GetTopic(d, "aux_command_topic")
		store.TopicStore[*d.AuxCommandTopic] = &d.AuxCommandFunc
	}
	if d.AuxStateFunc != nil {
		d.AuxStateTopic = new(string)
		*d.AuxStateTopic = GetTopic(d, "aux_state_topic")
	}
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = new(string)
		*d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CurrentTemperatureFunc != nil {
		d.CurrentTemperatureTopic = new(string)
		*d.CurrentTemperatureTopic = GetTopic(d, "current_temperature_topic")
	}
	if d.FanModeCommandFunc != nil {
		d.FanModeCommandTopic = new(string)
		*d.FanModeCommandTopic = GetTopic(d, "fan_mode_command_topic")
		store.TopicStore[*d.FanModeCommandTopic] = &d.FanModeCommandFunc
	}
	if d.FanModeStateFunc != nil {
		d.FanModeStateTopic = new(string)
		*d.FanModeStateTopic = GetTopic(d, "fan_mode_state_topic")
	}
	if d.ModeCommandFunc != nil {
		d.ModeCommandTopic = new(string)
		*d.ModeCommandTopic = GetTopic(d, "mode_command_topic")
		store.TopicStore[*d.ModeCommandTopic] = &d.ModeCommandFunc
	}
	if d.ModeStateFunc != nil {
		d.ModeStateTopic = new(string)
		*d.ModeStateTopic = GetTopic(d, "mode_state_topic")
	}
	if d.PowerCommandFunc != nil {
		d.PowerCommandTopic = new(string)
		*d.PowerCommandTopic = GetTopic(d, "power_command_topic")
		store.TopicStore[*d.PowerCommandTopic] = &d.PowerCommandFunc
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
	if d.SwingModeCommandFunc != nil {
		d.SwingModeCommandTopic = new(string)
		*d.SwingModeCommandTopic = GetTopic(d, "swing_mode_command_topic")
		store.TopicStore[*d.SwingModeCommandTopic] = &d.SwingModeCommandFunc
	}
	if d.SwingModeStateFunc != nil {
		d.SwingModeStateTopic = new(string)
		*d.SwingModeStateTopic = GetTopic(d, "swing_mode_state_topic")
	}
	if d.TemperatureCommandFunc != nil {
		d.TemperatureCommandTopic = new(string)
		*d.TemperatureCommandTopic = GetTopic(d, "temperature_command_topic")
		store.TopicStore[*d.TemperatureCommandTopic] = &d.TemperatureCommandFunc
	}
	if d.TemperatureHighCommandFunc != nil {
		d.TemperatureHighCommandTopic = new(string)
		*d.TemperatureHighCommandTopic = GetTopic(d, "temperature_high_command_topic")
		store.TopicStore[*d.TemperatureHighCommandTopic] = &d.TemperatureHighCommandFunc
	}
	if d.TemperatureHighStateFunc != nil {
		d.TemperatureHighStateTopic = new(string)
		*d.TemperatureHighStateTopic = GetTopic(d, "temperature_high_state_topic")
	}
	if d.TemperatureLowCommandFunc != nil {
		d.TemperatureLowCommandTopic = new(string)
		*d.TemperatureLowCommandTopic = GetTopic(d, "temperature_low_command_topic")
		store.TopicStore[*d.TemperatureLowCommandTopic] = &d.TemperatureLowCommandFunc
	}
	if d.TemperatureLowStateFunc != nil {
		d.TemperatureLowStateTopic = new(string)
		*d.TemperatureLowStateTopic = GetTopic(d, "temperature_low_state_topic")
	}
	if d.TemperatureStateFunc != nil {
		d.TemperatureStateTopic = new(string)
		*d.TemperatureStateTopic = GetTopic(d, "temperature_state_topic")
	}
}
func (d *Climate) SetMQTTFields(fields MQTTFields) {
	d.MQTT = &fields
}
func (d *Climate) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
