package iotconfig

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
type Button struct {
	Command []string `json:"command"`
	MQTT    struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    float64 `json:"force_update"`
	} `json:"mqtt"`
}
