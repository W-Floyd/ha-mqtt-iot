package InternalDevice

import ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Button) Translate() ExternalDevice.Button {
	eDevice := ExternalDevice.Button{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	return eDevice
}

type Button struct {
	Command []string `json:"command"`
	MQTT    struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
