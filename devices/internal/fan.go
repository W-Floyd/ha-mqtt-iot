package InternalDevice

import ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Fan) Translate() ExternalDevice.Fan {
	eDevice := ExternalDevice.Fan{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	return eDevice
}

type Fan struct {
	Command            []string `json:"command"`
	OscillationCommand []string `json:"oscillation_command"`
	OscillationState   []string `json:"oscillation_state"`
	PercentageCommand  []string `json:"percentage_command"`
	PercentageState    []string `json:"percentage_state"`
	PresetModeCommand  []string `json:"preset_mode_command"`
	PresetModeState    []string `json:"preset_mode_state"`
	State              []string `json:"state"`
	MQTT               struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
