package InternalDevice

import ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Cover) Translate() ExternalDevice.Cover {
	eDevice := ExternalDevice.Cover{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	return eDevice
}

type Cover struct {
	Command     []string `json:"command"`
	Position    []string `json:"position"`
	SetPosition []string `json:"set_position"`
	State       []string `json:"state"`
	TiltCommand []string `json:"tilt_command"`
	TiltStatus  []string `json:"tilt_status"`
	MQTT        struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
