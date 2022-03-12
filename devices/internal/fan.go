package InternalDevice

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
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
		ForceUpdate    float64 `json:"force_update"`
	} `json:"mqtt"`
}
