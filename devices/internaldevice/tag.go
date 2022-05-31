package InternalDevice

import externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
func (iDevice Tag) Translate() externaldevice.Tag {
	eDevice := externaldevice.Tag{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.Topic = iDevice.Topic
	eDevice.ValueTemplate = iDevice.ValueTemplate
	eDevice.Initialize()
	return eDevice
}

type Tag struct {
	Topic         string `json:"topic"`          // "The MQTT topic subscribed to receive tag scanned events."
	ValueTemplate string `json:"value_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) that returns a tag ID."
	MQTT          struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}