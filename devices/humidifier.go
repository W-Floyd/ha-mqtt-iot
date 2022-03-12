package iotconfig

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
type Humidifier struct {
	Command               []string `json:"command"`
	ModeCommand           []string `json:"mode_command"`
	ModeState             []string `json:"mode_state"`
	State                 []string `json:"state"`
	TargetHumidityCommand []string `json:"target_humidity_command"`
	TargetHumidityState   []string `json:"target_humidity_state"`
	MQTT                  struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    float64 `json:"force_update"`
	} `json:"mqtt"`
}
