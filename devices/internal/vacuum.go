package InternalDevice

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
type Vacuum struct {
	Command     []string `json:"command"`
	SendCommand []string `json:"send_command"`
	SetFanSpeed []string `json:"set_fan_speed"`
	State       []string `json:"state"`
	MQTT        struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    float64 `json:"force_update"`
	} `json:"mqtt"`
}
