package InternalDevice

import ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice DeviceTracker) Translate() ExternalDevice.DeviceTracker {
	eDevice := ExternalDevice.DeviceTracker{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.Devices = iDevice.Devices
	eDevice.PayloadHome = iDevice.PayloadHome
	eDevice.PayloadNotHome = iDevice.PayloadNotHome
	eDevice.Qos = iDevice.Qos
	eDevice.SourceType = iDevice.SourceType
	eDevice.Initialize()
	return eDevice
}

type DeviceTracker struct {
	Devices        []string `json:"devices"`
	PayloadHome    string   `json:"payload_home"`
	PayloadNotHome string   `json:"payload_not_home"`
	Qos            int      `json:"qos"`
	SourceType     string   `json:"source_type"`
	MQTT           struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
