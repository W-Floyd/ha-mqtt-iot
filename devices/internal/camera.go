package InternalDevice

import ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Camera) Translate() ExternalDevice.Camera {
	eDevice := ExternalDevice.Camera{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	return eDevice
}

type Camera struct {
	MQTT struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
