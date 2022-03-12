package iotconfig

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
type Sensor struct {
	State []string `json:"state"`
	MQTT  struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    float64 `json:"force_update"`
	} `json:"mqtt"`
}
