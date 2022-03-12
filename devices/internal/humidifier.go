package InternalDevice

import ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Humidifier) Translate() ExternalDevice.Humidifier {
	eDevice := ExternalDevice.Humidifier{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	return eDevice
}

type Humidifier struct {
	Command               []string `json:"command"`
	ModeCommand           []string `json:"mode_command"`
	ModeState             []string `json:"mode_state"`
	State                 []string `json:"state"`
	TargetHumidityCommand []string `json:"target_humidity_command"`
	TargetHumidityState   []string `json:"target_humidity_state"`
	MQTT                  struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
