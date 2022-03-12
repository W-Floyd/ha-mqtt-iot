package InternalDevice

import ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Climate) Translate() ExternalDevice.Climate {
	eDevice := ExternalDevice.Climate{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	return eDevice
}

type Climate struct {
	Action                 []string `json:"action"`
	AuxCommand             []string `json:"aux_command"`
	AuxState               []string `json:"aux_state"`
	CurrentTemperature     []string `json:"current_temperature"`
	FanModeCommand         []string `json:"fan_mode_command"`
	FanModeState           []string `json:"fan_mode_state"`
	ModeCommand            []string `json:"mode_command"`
	ModeState              []string `json:"mode_state"`
	PowerCommand           []string `json:"power_command"`
	PresetModeCommand      []string `json:"preset_mode_command"`
	PresetModeState        []string `json:"preset_mode_state"`
	SwingModeCommand       []string `json:"swing_mode_command"`
	SwingModeState         []string `json:"swing_mode_state"`
	TemperatureCommand     []string `json:"temperature_command"`
	TemperatureHighCommand []string `json:"temperature_high_command"`
	TemperatureHighState   []string `json:"temperature_high_state"`
	TemperatureLowCommand  []string `json:"temperature_low_command"`
	TemperatureLowState    []string `json:"temperature_low_state"`
	TemperatureState       []string `json:"temperature_state"`
	MQTT                   struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
