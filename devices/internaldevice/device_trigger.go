package InternalDevice

import externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice DeviceTrigger) Translate() externaldevice.DeviceTrigger {
	eDevice := externaldevice.DeviceTrigger{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.AutomationType = iDevice.AutomationType
	eDevice.Payload = iDevice.Payload
	eDevice.Qos = iDevice.Qos
	eDevice.Subtype = iDevice.Subtype
	eDevice.Topic = iDevice.Topic
	eDevice.Type = iDevice.Type
	eDevice.ValueTemplate = iDevice.ValueTemplate
	eDevice.Initialize()
	return eDevice
}

type DeviceTrigger struct {
	AutomationType string `json:"automation_type"`
	Payload        string `json:"payload"`
	Qos            int    `json:"qos"`
	Subtype        string `json:"subtype"`
	Topic          string `json:"topic"`
	Type           string `json:"type"`
	ValueTemplate  string `json:"value_template"`
	MQTT           struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
