package InternalDevice

import externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
func (iDevice Tag) Translate() externaldevice.Tag {
	eDevice := externaldevice.Tag{}
	eDevice.MQTT = new(externaldevice.MQTTFields)
	if iDevice.MQTT.ForceUpdate != nil {
		eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	}
	if iDevice.MQTT.UpdateInterval != nil {
		eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	}
	if iDevice.Topic != nil {
		eDevice.Topic = iDevice.Topic
	}
	if iDevice.ValueTemplate != nil {
		eDevice.ValueTemplate = iDevice.ValueTemplate
	}
	eDevice.Initialize()
	return eDevice
}

type Tag struct {
	Topic         *string `json:"topic,omitempty"`          // "The MQTT topic subscribed to receive tag scanned events."
	ValueTemplate *string `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) that returns a tag ID."
	MQTT          struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}
