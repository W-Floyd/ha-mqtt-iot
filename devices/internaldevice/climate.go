package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Climate struct {
	ActionTemplate                 *string     `json:"action_template,omitempty"` // "A template to render the value received on the `action_topic` with."
	Action                         *([]string) `json:"action,omitempty"`
	AuxCommand                     *([]string) `json:"aux_command,omitempty"`
	AuxStateTemplate               *string     `json:"aux_state_template,omitempty"` // "A template to render the value received on the `aux_state_topic` with."
	AuxState                       *([]string) `json:"aux_state,omitempty"`
	AvailabilityMode               *string     `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate           *string     `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability                   *([]string) `json:"availability,omitempty"`
	CurrentTemperatureTemplate     *string     `json:"current_temperature_template,omitempty"` // "A template with which the value received on `current_temperature_topic` will be rendered."
	CurrentTemperature             *([]string) `json:"current_temperature,omitempty"`
	EnabledByDefault               *bool       `json:"enabled_by_default,omitempty"`        // "Flag which defines if the entity should be enabled when first added."
	Encoding                       *string     `json:"encoding,omitempty"`                  // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory                 *string     `json:"entity_category,omitempty"`           // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	FanModeCommandTemplate         *string     `json:"fan_mode_command_template,omitempty"` // "A template to render the value sent to the `fan_mode_command_topic` with."
	FanModeCommand                 *([]string) `json:"fan_mode_command,omitempty"`
	FanModeStateTemplate           *string     `json:"fan_mode_state_template,omitempty"` // "A template to render the value received on the `fan_mode_state_topic` with."
	FanModeState                   *([]string) `json:"fan_mode_state,omitempty"`
	FanModes                       *([]string) `json:"fan_modes,omitempty"`                // "A list of supported fan modes."
	Icon                           *string     `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Initial                        *int        `json:"initial,omitempty"`                  // "Set the initial target temperature."
	JsonAttributesTemplate         *string     `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributes                 *([]string) `json:"json_attributes,omitempty"`
	MaxTemp                        *float64    `json:"max_temp,omitempty"`              // "Maximum set point available."
	MinTemp                        *float64    `json:"min_temp,omitempty"`              // "Minimum set point available."
	ModeCommandTemplate            *string     `json:"mode_command_template,omitempty"` // "A template to render the value sent to the `mode_command_topic` with."
	ModeCommand                    *([]string) `json:"mode_command,omitempty"`
	ModeStateTemplate              *string     `json:"mode_state_template,omitempty"` // "A template to render the value received on the `mode_state_topic` with."
	ModeState                      *([]string) `json:"mode_state,omitempty"`
	Modes                          *([]string) `json:"modes,omitempty"`                 // "A list of supported modes. Needs to be a subset of the default values."
	Name                           *string     `json:"name,omitempty"`                  // "The name of the HVAC."
	ObjectId                       *string     `json:"object_id,omitempty"`             // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable               *string     `json:"payload_available,omitempty"`     // "The payload that represents the available state."
	PayloadNotAvailable            *string     `json:"payload_not_available,omitempty"` // "The payload that represents the unavailable state."
	PayloadOff                     *string     `json:"payload_off,omitempty"`           // "The payload that represents disabled state."
	PayloadOn                      *string     `json:"payload_on,omitempty"`            // "The payload that represents enabled state."
	PowerCommand                   *([]string) `json:"power_command,omitempty"`
	Precision                      *float64    `json:"precision,omitempty"`                    // "The desired precision for this device. Can be used to match your actual thermostat's precision. Supported values are `0.1`, `0.5` and `1.0`."
	PresetModeCommandTemplate      *string     `json:"preset_mode_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `preset_mode_command_topic`."
	PresetModeCommand              *([]string) `json:"preset_mode_command,omitempty"`
	PresetModeState                *([]string) `json:"preset_mode_state,omitempty"`
	PresetModeValueTemplate        *string     `json:"preset_mode_value_template,omitempty"`  // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the `preset_mode` value from the payload received on `preset_mode_state_topic`."
	PresetModes                    *([]string) `json:"preset_modes,omitempty"`                // "List of preset modes this climate is supporting. Common examples include `eco`, `away`, `boost`, `comfort`, `home`, `sleep` and `activity`."
	Qos                            *int        `json:"qos,omitempty"`                         // "The maximum QoS level to be used when receiving and publishing messages."
	Retain                         *bool       `json:"retain,omitempty"`                      // "Defines if published messages should have the retain flag set."
	SwingModeCommandTemplate       *string     `json:"swing_mode_command_template,omitempty"` // "A template to render the value sent to the `swing_mode_command_topic` with."
	SwingModeCommand               *([]string) `json:"swing_mode_command,omitempty"`
	SwingModeStateTemplate         *string     `json:"swing_mode_state_template,omitempty"` // "A template to render the value received on the `swing_mode_state_topic` with."
	SwingModeState                 *([]string) `json:"swing_mode_state,omitempty"`
	SwingModes                     *([]string) `json:"swing_modes,omitempty"`                  // "A list of supported swing modes."
	TempStep                       *float64    `json:"temp_step,omitempty"`                    // "Step size for temperature set point."
	TemperatureCommandTemplate     *string     `json:"temperature_command_template,omitempty"` // "A template to render the value sent to the `temperature_command_topic` with."
	TemperatureCommand             *([]string) `json:"temperature_command,omitempty"`
	TemperatureHighCommandTemplate *string     `json:"temperature_high_command_template,omitempty"` // "A template to render the value sent to the `temperature_high_command_topic` with."
	TemperatureHighCommand         *([]string) `json:"temperature_high_command,omitempty"`
	TemperatureHighStateTemplate   *string     `json:"temperature_high_state_template,omitempty"` // "A template to render the value received on the `temperature_high_state_topic` with."
	TemperatureHighState           *([]string) `json:"temperature_high_state,omitempty"`
	TemperatureLowCommandTemplate  *string     `json:"temperature_low_command_template,omitempty"` // "A template to render the value sent to the `temperature_low_command_topic` with."
	TemperatureLowCommand          *([]string) `json:"temperature_low_command,omitempty"`
	TemperatureLowStateTemplate    *string     `json:"temperature_low_state_template,omitempty"` // "A template to render the value received on the `temperature_low_state_topic` with."
	TemperatureLowState            *([]string) `json:"temperature_low_state,omitempty"`
	TemperatureStateTemplate       *string     `json:"temperature_state_template,omitempty"` // "A template to render the value received on the `temperature_state_topic` with."
	TemperatureState               *([]string) `json:"temperature_state,omitempty"`
	TemperatureUnit                *string     `json:"temperature_unit,omitempty"` // "Defines the temperature unit of the device, `C` or `F`. If this is not set, the temperature unit is set to the system temperature unit."
	UniqueId                       *string     `json:"unique_id,omitempty"`        // "An ID that uniquely identifies this HVAC device. If two HVAC devices have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate                  *string     `json:"value_template,omitempty"`   // "Default template to render the payloads on *all* `*_state_topic`s with."
	MQTT                           struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}

func (iDevice Climate) Translate() externaldevice.Climate {
	eDevice := externaldevice.Climate{}
	eDevice.MQTT = new(externaldevice.MQTTFields)
	if iDevice.MQTT.ForceUpdate != nil {
		eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	}
	if iDevice.MQTT.UpdateInterval != nil {
		eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	}
	if iDevice.ActionTemplate != nil {
		eDevice.ActionTemplate = iDevice.ActionTemplate
	}
	if iDevice.Action != nil {
		eDevice.ActionFunc = common.ConstructCommandFunc(*iDevice.Action)
	}
	if iDevice.AuxCommand != nil {
		eDevice.AuxCommandFunc = common.ConstructCommandFunc(*iDevice.AuxCommand)
	}
	if iDevice.AuxStateTemplate != nil {
		eDevice.AuxStateTemplate = iDevice.AuxStateTemplate
	}
	if iDevice.AuxState != nil {
		eDevice.AuxStateFunc = common.ConstructStateFunc(*iDevice.AuxState)
	}
	if iDevice.AvailabilityMode != nil {
		eDevice.AvailabilityMode = iDevice.AvailabilityMode
	}
	if iDevice.AvailabilityTemplate != nil {
		eDevice.AvailabilityTemplate = iDevice.AvailabilityTemplate
	}
	if iDevice.Availability != nil {
		eDevice.AvailabilityFunc = common.ConstructStateFunc(*iDevice.Availability)
	}
	if iDevice.CurrentTemperatureTemplate != nil {
		eDevice.CurrentTemperatureTemplate = iDevice.CurrentTemperatureTemplate
	}
	if iDevice.CurrentTemperature != nil {
		eDevice.CurrentTemperatureFunc = common.ConstructStateFunc(*iDevice.CurrentTemperature)
	}
	if iDevice.EnabledByDefault != nil {
		eDevice.EnabledByDefault = iDevice.EnabledByDefault
	}
	if iDevice.Encoding != nil {
		eDevice.Encoding = iDevice.Encoding
	}
	if iDevice.EntityCategory != nil {
		eDevice.EntityCategory = iDevice.EntityCategory
	}
	if iDevice.FanModeCommandTemplate != nil {
		eDevice.FanModeCommandTemplate = iDevice.FanModeCommandTemplate
	}
	if iDevice.FanModeCommand != nil {
		eDevice.FanModeCommandFunc = common.ConstructCommandFunc(*iDevice.FanModeCommand)
	}
	if iDevice.FanModeStateTemplate != nil {
		eDevice.FanModeStateTemplate = iDevice.FanModeStateTemplate
	}
	if iDevice.FanModeState != nil {
		eDevice.FanModeStateFunc = common.ConstructStateFunc(*iDevice.FanModeState)
	}
	if iDevice.FanModes != nil {
		eDevice.FanModes = iDevice.FanModes
	}
	if iDevice.Icon != nil {
		eDevice.Icon = iDevice.Icon
	}
	if iDevice.Initial != nil {
		eDevice.Initial = iDevice.Initial
	}
	if iDevice.JsonAttributesTemplate != nil {
		eDevice.JsonAttributesTemplate = iDevice.JsonAttributesTemplate
	}
	if iDevice.JsonAttributes != nil {
		eDevice.JsonAttributesFunc = common.ConstructCommandFunc(*iDevice.JsonAttributes)
	}
	if iDevice.MaxTemp != nil {
		eDevice.MaxTemp = iDevice.MaxTemp
	}
	if iDevice.MinTemp != nil {
		eDevice.MinTemp = iDevice.MinTemp
	}
	if iDevice.ModeCommandTemplate != nil {
		eDevice.ModeCommandTemplate = iDevice.ModeCommandTemplate
	}
	if iDevice.ModeCommand != nil {
		eDevice.ModeCommandFunc = common.ConstructCommandFunc(*iDevice.ModeCommand)
	}
	if iDevice.ModeStateTemplate != nil {
		eDevice.ModeStateTemplate = iDevice.ModeStateTemplate
	}
	if iDevice.ModeState != nil {
		eDevice.ModeStateFunc = common.ConstructStateFunc(*iDevice.ModeState)
	}
	if iDevice.Modes != nil {
		eDevice.Modes = iDevice.Modes
	}
	if iDevice.Name != nil {
		eDevice.Name = iDevice.Name
	}
	if iDevice.ObjectId != nil {
		eDevice.ObjectId = iDevice.ObjectId
	}
	if iDevice.PayloadAvailable != nil {
		eDevice.PayloadAvailable = iDevice.PayloadAvailable
	}
	if iDevice.PayloadNotAvailable != nil {
		eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	}
	if iDevice.PayloadOff != nil {
		eDevice.PayloadOff = iDevice.PayloadOff
	}
	if iDevice.PayloadOn != nil {
		eDevice.PayloadOn = iDevice.PayloadOn
	}
	if iDevice.PowerCommand != nil {
		eDevice.PowerCommandFunc = common.ConstructCommandFunc(*iDevice.PowerCommand)
	}
	if iDevice.Precision != nil {
		eDevice.Precision = iDevice.Precision
	}
	if iDevice.PresetModeCommandTemplate != nil {
		eDevice.PresetModeCommandTemplate = iDevice.PresetModeCommandTemplate
	}
	if iDevice.PresetModeCommand != nil {
		eDevice.PresetModeCommandFunc = common.ConstructCommandFunc(*iDevice.PresetModeCommand)
	}
	if iDevice.PresetModeState != nil {
		eDevice.PresetModeStateFunc = common.ConstructStateFunc(*iDevice.PresetModeState)
	}
	if iDevice.PresetModeValueTemplate != nil {
		eDevice.PresetModeValueTemplate = iDevice.PresetModeValueTemplate
	}
	if iDevice.PresetModes != nil {
		eDevice.PresetModes = iDevice.PresetModes
	}
	if iDevice.Qos != nil {
		eDevice.Qos = iDevice.Qos
	}
	if iDevice.Retain != nil {
		eDevice.Retain = iDevice.Retain
	}
	if iDevice.SwingModeCommandTemplate != nil {
		eDevice.SwingModeCommandTemplate = iDevice.SwingModeCommandTemplate
	}
	if iDevice.SwingModeCommand != nil {
		eDevice.SwingModeCommandFunc = common.ConstructCommandFunc(*iDevice.SwingModeCommand)
	}
	if iDevice.SwingModeStateTemplate != nil {
		eDevice.SwingModeStateTemplate = iDevice.SwingModeStateTemplate
	}
	if iDevice.SwingModeState != nil {
		eDevice.SwingModeStateFunc = common.ConstructStateFunc(*iDevice.SwingModeState)
	}
	if iDevice.SwingModes != nil {
		eDevice.SwingModes = iDevice.SwingModes
	}
	if iDevice.TempStep != nil {
		eDevice.TempStep = iDevice.TempStep
	}
	if iDevice.TemperatureCommandTemplate != nil {
		eDevice.TemperatureCommandTemplate = iDevice.TemperatureCommandTemplate
	}
	if iDevice.TemperatureCommand != nil {
		eDevice.TemperatureCommandFunc = common.ConstructCommandFunc(*iDevice.TemperatureCommand)
	}
	if iDevice.TemperatureHighCommandTemplate != nil {
		eDevice.TemperatureHighCommandTemplate = iDevice.TemperatureHighCommandTemplate
	}
	if iDevice.TemperatureHighCommand != nil {
		eDevice.TemperatureHighCommandFunc = common.ConstructCommandFunc(*iDevice.TemperatureHighCommand)
	}
	if iDevice.TemperatureHighStateTemplate != nil {
		eDevice.TemperatureHighStateTemplate = iDevice.TemperatureHighStateTemplate
	}
	if iDevice.TemperatureHighState != nil {
		eDevice.TemperatureHighStateFunc = common.ConstructStateFunc(*iDevice.TemperatureHighState)
	}
	if iDevice.TemperatureLowCommandTemplate != nil {
		eDevice.TemperatureLowCommandTemplate = iDevice.TemperatureLowCommandTemplate
	}
	if iDevice.TemperatureLowCommand != nil {
		eDevice.TemperatureLowCommandFunc = common.ConstructCommandFunc(*iDevice.TemperatureLowCommand)
	}
	if iDevice.TemperatureLowStateTemplate != nil {
		eDevice.TemperatureLowStateTemplate = iDevice.TemperatureLowStateTemplate
	}
	if iDevice.TemperatureLowState != nil {
		eDevice.TemperatureLowStateFunc = common.ConstructStateFunc(*iDevice.TemperatureLowState)
	}
	if iDevice.TemperatureStateTemplate != nil {
		eDevice.TemperatureStateTemplate = iDevice.TemperatureStateTemplate
	}
	if iDevice.TemperatureState != nil {
		eDevice.TemperatureStateFunc = common.ConstructStateFunc(*iDevice.TemperatureState)
	}
	if iDevice.TemperatureUnit != nil {
		eDevice.TemperatureUnit = iDevice.TemperatureUnit
	}
	if iDevice.UniqueId != nil {
		eDevice.UniqueId = iDevice.UniqueId
	}
	if iDevice.ValueTemplate != nil {
		eDevice.ValueTemplate = iDevice.ValueTemplate
	}
	if iDevice.Availability == nil {
		eDevice.AvailabilityFunc = common.AvailabilityFunc
	}
	eDevice.Initialize()
	return eDevice
}
