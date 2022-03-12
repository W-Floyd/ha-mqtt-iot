package iotconfig

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
type Cover struct {
	Command     []string `json:"command"`
	Position    []string `json:"position"`
	SetPosition []string `json:"set_position"`
	State       []string `json:"state"`
	TiltCommand []string `json:"tilt_command"`
	TiltStatus  []string `json:"tilt_status"`
	MQTT        struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    float64 `json:"force_update"`
	} `json:"mqtt"`
}
