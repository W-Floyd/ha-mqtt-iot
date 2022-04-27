package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
func (iDevice Climate) Translate() externaldevice.Climate {
	eDevice := externaldevice.Climate{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.ActionTemplate = iDevice.ActionTemplate
	eDevice.ActionFunc = common.ConstructCommandFunc(iDevice.Action)
	eDevice.AuxCommandFunc = common.ConstructCommandFunc(iDevice.AuxCommand)
	eDevice.AuxStateTemplate = iDevice.AuxStateTemplate
	eDevice.AuxStateFunc = common.ConstructStateFunc(iDevice.AuxState)
	eDevice.CurrentTemperatureTemplate = iDevice.CurrentTemperatureTemplate
	eDevice.CurrentTemperatureFunc = common.ConstructStateFunc(iDevice.CurrentTemperature)
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.FanModeCommandTemplate = iDevice.FanModeCommandTemplate
	eDevice.FanModeCommandFunc = common.ConstructCommandFunc(iDevice.FanModeCommand)
	eDevice.FanModeStateTemplate = iDevice.FanModeStateTemplate
	eDevice.FanModeStateFunc = common.ConstructStateFunc(iDevice.FanModeState)
	eDevice.FanModes = iDevice.FanModes
	eDevice.Icon = iDevice.Icon
	eDevice.Initial = iDevice.Initial
	eDevice.MaxTemp = iDevice.MaxTemp
	eDevice.MinTemp = iDevice.MinTemp
	eDevice.ModeCommandTemplate = iDevice.ModeCommandTemplate
	eDevice.ModeCommandFunc = common.ConstructCommandFunc(iDevice.ModeCommand)
	eDevice.ModeStateTemplate = iDevice.ModeStateTemplate
	eDevice.ModeStateFunc = common.ConstructStateFunc(iDevice.ModeState)
	eDevice.Modes = iDevice.Modes
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.PayloadOff = iDevice.PayloadOff
	eDevice.PayloadOn = iDevice.PayloadOn
	eDevice.PowerCommandFunc = common.ConstructCommandFunc(iDevice.PowerCommand)
	eDevice.Precision = iDevice.Precision
	eDevice.PresetModeCommandTemplate = iDevice.PresetModeCommandTemplate
	eDevice.PresetModeCommandFunc = common.ConstructCommandFunc(iDevice.PresetModeCommand)
	eDevice.PresetModeStateFunc = common.ConstructStateFunc(iDevice.PresetModeState)
	eDevice.PresetModeValueTemplate = iDevice.PresetModeValueTemplate
	eDevice.PresetModes = iDevice.PresetModes
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.SwingModeCommandTemplate = iDevice.SwingModeCommandTemplate
	eDevice.SwingModeCommandFunc = common.ConstructCommandFunc(iDevice.SwingModeCommand)
	eDevice.SwingModeStateTemplate = iDevice.SwingModeStateTemplate
	eDevice.SwingModeStateFunc = common.ConstructStateFunc(iDevice.SwingModeState)
	eDevice.SwingModes = iDevice.SwingModes
	eDevice.TempStep = iDevice.TempStep
	eDevice.TemperatureCommandTemplate = iDevice.TemperatureCommandTemplate
	eDevice.TemperatureCommandFunc = common.ConstructCommandFunc(iDevice.TemperatureCommand)
	eDevice.TemperatureHighCommandTemplate = iDevice.TemperatureHighCommandTemplate
	eDevice.TemperatureHighCommandFunc = common.ConstructCommandFunc(iDevice.TemperatureHighCommand)
	eDevice.TemperatureHighStateTemplate = iDevice.TemperatureHighStateTemplate
	eDevice.TemperatureHighStateFunc = common.ConstructStateFunc(iDevice.TemperatureHighState)
	eDevice.TemperatureLowCommandTemplate = iDevice.TemperatureLowCommandTemplate
	eDevice.TemperatureLowCommandFunc = common.ConstructCommandFunc(iDevice.TemperatureLowCommand)
	eDevice.TemperatureLowStateTemplate = iDevice.TemperatureLowStateTemplate
	eDevice.TemperatureLowStateFunc = common.ConstructStateFunc(iDevice.TemperatureLowState)
	eDevice.TemperatureStateTemplate = iDevice.TemperatureStateTemplate
	eDevice.TemperatureStateFunc = common.ConstructStateFunc(iDevice.TemperatureState)
	eDevice.TemperatureUnit = iDevice.TemperatureUnit
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.ValueTemplate = iDevice.ValueTemplate
	eDevice.Initialize()
	return eDevice
}

type Climate struct {
	ActionTemplate                 string   `json:"action_template"` // "A template to render the value received on the `action_topic` with."
	Action                         []string `json:"action"`
	AuxCommand                     []string `json:"aux_command"`
	AuxStateTemplate               string   `json:"aux_state_template"` // "A template to render the value received on the `aux_state_topic` with."
	AuxState                       []string `json:"aux_state"`
	CurrentTemperatureTemplate     string   `json:"current_temperature_template"` // "A template with which the value received on `current_temperature_topic` will be rendered."
	CurrentTemperature             []string `json:"current_temperature"`
	EnabledByDefault               bool     `json:"enabled_by_default"`        // "Flag which defines if the entity should be enabled when first added."
	Encoding                       string   `json:"encoding"`                  // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory                 string   `json:"entity_category"`           // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	FanModeCommandTemplate         string   `json:"fan_mode_command_template"` // "A template to render the value sent to the `fan_mode_command_topic` with."
	FanModeCommand                 []string `json:"fan_mode_command"`
	FanModeStateTemplate           string   `json:"fan_mode_state_template"` // "A template to render the value received on the `fan_mode_state_topic` with."
	FanModeState                   []string `json:"fan_mode_state"`
	FanModes                       []string `json:"fan_modes"`             // "A list of supported fan modes."
	Icon                           string   `json:"icon"`                  // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Initial                        int      `json:"initial"`               // "Set the initial target temperature."
	MaxTemp                        float64  `json:"max_temp"`              // "Maximum set point available."
	MinTemp                        float64  `json:"min_temp"`              // "Minimum set point available."
	ModeCommandTemplate            string   `json:"mode_command_template"` // "A template to render the value sent to the `mode_command_topic` with."
	ModeCommand                    []string `json:"mode_command"`
	ModeStateTemplate              string   `json:"mode_state_template"` // "A template to render the value received on the `mode_state_topic` with."
	ModeState                      []string `json:"mode_state"`
	Modes                          []string `json:"modes"`                 // "A list of supported modes. Needs to be a subset of the default values."
	Name                           string   `json:"name"`                  // "The name of the HVAC."
	ObjectId                       string   `json:"object_id"`             // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable               string   `json:"payload_available"`     // "The payload that represents the available state."
	PayloadNotAvailable            string   `json:"payload_not_available"` // "The payload that represents the unavailable state."
	PayloadOff                     string   `json:"payload_off"`           // "The payload that represents disabled state."
	PayloadOn                      string   `json:"payload_on"`            // "The payload that represents enabled state."
	PowerCommand                   []string `json:"power_command"`
	Precision                      float64  `json:"precision"`                    // "The desired precision for this device. Can be used to match your actual thermostat's precision. Supported values are `0.1`, `0.5` and `1.0`."
	PresetModeCommandTemplate      string   `json:"preset_mode_command_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `preset_mode_command_topic`."
	PresetModeCommand              []string `json:"preset_mode_command"`
	PresetModeState                []string `json:"preset_mode_state"`
	PresetModeValueTemplate        string   `json:"preset_mode_value_template"`  // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the `preset_mode` value from the payload received on `preset_mode_state_topic`."
	PresetModes                    []string `json:"preset_modes"`                // "List of preset modes this climate is supporting. Common examples include `eco`, `away`, `boost`, `comfort`, `home`, `sleep` and `activity`."
	Qos                            int      `json:"qos"`                         // "The maximum QoS level to be used when receiving and publishing messages."
	Retain                         bool     `json:"retain"`                      // "Defines if published messages should have the retain flag set."
	SwingModeCommandTemplate       string   `json:"swing_mode_command_template"` // "A template to render the value sent to the `swing_mode_command_topic` with."
	SwingModeCommand               []string `json:"swing_mode_command"`
	SwingModeStateTemplate         string   `json:"swing_mode_state_template"` // "A template to render the value received on the `swing_mode_state_topic` with."
	SwingModeState                 []string `json:"swing_mode_state"`
	SwingModes                     []string `json:"swing_modes"`                  // "A list of supported swing modes."
	TempStep                       float64  `json:"temp_step"`                    // "Step size for temperature set point."
	TemperatureCommandTemplate     string   `json:"temperature_command_template"` // "A template to render the value sent to the `temperature_command_topic` with."
	TemperatureCommand             []string `json:"temperature_command"`
	TemperatureHighCommandTemplate string   `json:"temperature_high_command_template"` // "A template to render the value sent to the `temperature_high_command_topic` with."
	TemperatureHighCommand         []string `json:"temperature_high_command"`
	TemperatureHighStateTemplate   string   `json:"temperature_high_state_template"` // "A template to render the value received on the `temperature_high_state_topic` with."
	TemperatureHighState           []string `json:"temperature_high_state"`
	TemperatureLowCommandTemplate  string   `json:"temperature_low_command_template"` // "A template to render the value sent to the `temperature_low_command_topic` with."
	TemperatureLowCommand          []string `json:"temperature_low_command"`
	TemperatureLowStateTemplate    string   `json:"temperature_low_state_template"` // "A template to render the value received on the `temperature_low_state_topic` with."
	TemperatureLowState            []string `json:"temperature_low_state"`
	TemperatureStateTemplate       string   `json:"temperature_state_template"` // "A template to render the value received on the `temperature_state_topic` with."
	TemperatureState               []string `json:"temperature_state"`
	TemperatureUnit                string   `json:"temperature_unit"` // "Defines the temperature unit of the device, `C` or `F`. If this is not set, the temperature unit is set to the system temperature unit."
	UniqueId                       string   `json:"unique_id"`        // "An ID that uniquely identifies this HVAC device. If two HVAC devices have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate                  string   `json:"value_template"`   // "Default template to render the payloads on *all* `*_state_topic`s with."
	MQTT                           struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
