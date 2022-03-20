package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Fan) Translate() ExternalDevice.Fan {
	eDevice := ExternalDevice.Fan{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.CommandTemplate = iDevice.CommandTemplate
	eDevice.CommandFunc = common.ConstructCommandFunc(iDevice.Command)
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.Icon = iDevice.Icon
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.Optimistic = iDevice.Optimistic
	eDevice.OscillationCommandTemplate = iDevice.OscillationCommandTemplate
	eDevice.OscillationCommandFunc = common.ConstructCommandFunc(iDevice.OscillationCommand)
	eDevice.OscillationStateFunc = common.ConstructStateFunc(iDevice.OscillationState)
	eDevice.OscillationValueTemplate = iDevice.OscillationValueTemplate
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.PayloadOff = iDevice.PayloadOff
	eDevice.PayloadOn = iDevice.PayloadOn
	eDevice.PayloadOscillationOff = iDevice.PayloadOscillationOff
	eDevice.PayloadOscillationOn = iDevice.PayloadOscillationOn
	eDevice.PayloadResetPercentage = iDevice.PayloadResetPercentage
	eDevice.PayloadResetPresetMode = iDevice.PayloadResetPresetMode
	eDevice.PercentageCommandTemplate = iDevice.PercentageCommandTemplate
	eDevice.PercentageCommandFunc = common.ConstructCommandFunc(iDevice.PercentageCommand)
	eDevice.PercentageStateFunc = common.ConstructStateFunc(iDevice.PercentageState)
	eDevice.PercentageValueTemplate = iDevice.PercentageValueTemplate
	eDevice.PresetModeCommandTemplate = iDevice.PresetModeCommandTemplate
	eDevice.PresetModeCommandFunc = common.ConstructCommandFunc(iDevice.PresetModeCommand)
	eDevice.PresetModeStateFunc = common.ConstructStateFunc(iDevice.PresetModeState)
	eDevice.PresetModeValueTemplate = iDevice.PresetModeValueTemplate
	eDevice.PresetModes = iDevice.PresetModes
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.SpeedRangeMax = iDevice.SpeedRangeMax
	eDevice.SpeedRangeMin = iDevice.SpeedRangeMin
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.StateValueTemplate = iDevice.StateValueTemplate
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.Initialize()
	return eDevice
}

type Fan struct {
	CommandTemplate            string   `json:"command_template"`
	Command                    []string `json:"command"`
	EnabledByDefault           bool     `json:"enabled_by_default"`
	Encoding                   string   `json:"encoding"`
	EntityCategory             string   `json:"entity_category"`
	Icon                       string   `json:"icon"`
	Name                       string   `json:"name"`
	ObjectId                   string   `json:"object_id"`
	Optimistic                 bool     `json:"optimistic"`
	OscillationCommandTemplate string   `json:"oscillation_command_template"`
	OscillationCommand         []string `json:"oscillation_command"`
	OscillationState           []string `json:"oscillation_state"`
	OscillationValueTemplate   string   `json:"oscillation_value_template"`
	PayloadAvailable           string   `json:"payload_available"`
	PayloadNotAvailable        string   `json:"payload_not_available"`
	PayloadOff                 string   `json:"payload_off"`
	PayloadOn                  string   `json:"payload_on"`
	PayloadOscillationOff      string   `json:"payload_oscillation_off"`
	PayloadOscillationOn       string   `json:"payload_oscillation_on"`
	PayloadResetPercentage     string   `json:"payload_reset_percentage"`
	PayloadResetPresetMode     string   `json:"payload_reset_preset_mode"`
	PercentageCommandTemplate  string   `json:"percentage_command_template"`
	PercentageCommand          []string `json:"percentage_command"`
	PercentageState            []string `json:"percentage_state"`
	PercentageValueTemplate    string   `json:"percentage_value_template"`
	PresetModeCommandTemplate  string   `json:"preset_mode_command_template"`
	PresetModeCommand          []string `json:"preset_mode_command"`
	PresetModeState            []string `json:"preset_mode_state"`
	PresetModeValueTemplate    string   `json:"preset_mode_value_template"`
	PresetModes                []string `json:"preset_modes"`
	Qos                        int      `json:"qos"`
	Retain                     bool     `json:"retain"`
	SpeedRangeMax              int      `json:"speed_range_max"`
	SpeedRangeMin              int      `json:"speed_range_min"`
	State                      []string `json:"state"`
	StateValueTemplate         string   `json:"state_value_template"`
	UniqueId                   string   `json:"unique_id"`
	MQTT                       struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
