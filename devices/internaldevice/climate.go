package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
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
	ActionTemplate                 string   `json:"action_template"`
	Action                         []string `json:"action"`
	AuxCommand                     []string `json:"aux_command"`
	AuxStateTemplate               string   `json:"aux_state_template"`
	AuxState                       []string `json:"aux_state"`
	CurrentTemperatureTemplate     string   `json:"current_temperature_template"`
	CurrentTemperature             []string `json:"current_temperature"`
	EnabledByDefault               bool     `json:"enabled_by_default"`
	Encoding                       string   `json:"encoding"`
	EntityCategory                 string   `json:"entity_category"`
	FanModeCommandTemplate         string   `json:"fan_mode_command_template"`
	FanModeCommand                 []string `json:"fan_mode_command"`
	FanModeStateTemplate           string   `json:"fan_mode_state_template"`
	FanModeState                   []string `json:"fan_mode_state"`
	FanModes                       []string `json:"fan_modes"`
	Icon                           string   `json:"icon"`
	Initial                        int      `json:"initial"`
	MaxTemp                        float64  `json:"max_temp"`
	MinTemp                        float64  `json:"min_temp"`
	ModeCommandTemplate            string   `json:"mode_command_template"`
	ModeCommand                    []string `json:"mode_command"`
	ModeStateTemplate              string   `json:"mode_state_template"`
	ModeState                      []string `json:"mode_state"`
	Modes                          []string `json:"modes"`
	Name                           string   `json:"name"`
	ObjectId                       string   `json:"object_id"`
	PayloadAvailable               string   `json:"payload_available"`
	PayloadNotAvailable            string   `json:"payload_not_available"`
	PayloadOff                     string   `json:"payload_off"`
	PayloadOn                      string   `json:"payload_on"`
	PowerCommand                   []string `json:"power_command"`
	Precision                      float64  `json:"precision"`
	PresetModeCommandTemplate      string   `json:"preset_mode_command_template"`
	PresetModeCommand              []string `json:"preset_mode_command"`
	PresetModeState                []string `json:"preset_mode_state"`
	PresetModeValueTemplate        string   `json:"preset_mode_value_template"`
	PresetModes                    []string `json:"preset_modes"`
	Qos                            int      `json:"qos"`
	Retain                         bool     `json:"retain"`
	SwingModeCommandTemplate       string   `json:"swing_mode_command_template"`
	SwingModeCommand               []string `json:"swing_mode_command"`
	SwingModeStateTemplate         string   `json:"swing_mode_state_template"`
	SwingModeState                 []string `json:"swing_mode_state"`
	SwingModes                     []string `json:"swing_modes"`
	TempStep                       float64  `json:"temp_step"`
	TemperatureCommandTemplate     string   `json:"temperature_command_template"`
	TemperatureCommand             []string `json:"temperature_command"`
	TemperatureHighCommandTemplate string   `json:"temperature_high_command_template"`
	TemperatureHighCommand         []string `json:"temperature_high_command"`
	TemperatureHighStateTemplate   string   `json:"temperature_high_state_template"`
	TemperatureHighState           []string `json:"temperature_high_state"`
	TemperatureLowCommandTemplate  string   `json:"temperature_low_command_template"`
	TemperatureLowCommand          []string `json:"temperature_low_command"`
	TemperatureLowStateTemplate    string   `json:"temperature_low_state_template"`
	TemperatureLowState            []string `json:"temperature_low_state"`
	TemperatureStateTemplate       string   `json:"temperature_state_template"`
	TemperatureState               []string `json:"temperature_state"`
	TemperatureUnit                string   `json:"temperature_unit"`
	UniqueId                       string   `json:"unique_id"`
	ValueTemplate                  string   `json:"value_template"`
	MQTT                           struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
