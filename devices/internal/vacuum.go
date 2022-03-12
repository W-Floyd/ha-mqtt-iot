package InternalDevice

import ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Vacuum) Translate() ExternalDevice.Vacuum {
	eDevice := ExternalDevice.Vacuum{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	return eDevice
}

type Vacuum struct {
	Command     []string `json:"command"`
	SendCommand []string `json:"send_command"`
	SetFanSpeed []string `json:"set_fan_speed"`
	State       []string `json:"state"`
	MQTT        struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
