package InternalDevice

import ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Tag) Translate() ExternalDevice.Tag {
	eDevice := ExternalDevice.Tag{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	return eDevice
}

type Tag struct {
	MQTT struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
