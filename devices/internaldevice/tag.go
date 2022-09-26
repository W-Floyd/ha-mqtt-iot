package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Tag struct {
	State         *([]string) `json:"state,omitempty"`
	ValueTemplate *string     `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) that returns a tag ID."
	MQTT          struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}

func (iDevice Tag) Translate() externaldevice.Tag {
	eDevice := externaldevice.Tag{}
	eDevice.MQTT = new(externaldevice.MQTTFields)
	if iDevice.MQTT.ForceUpdate != nil {
		eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	}
	if iDevice.MQTT.UpdateInterval != nil {
		eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	}
	if iDevice.State != nil {
		eDevice.StateFunc = common.ConstructStateFunc(*iDevice.State)
	}
	if iDevice.ValueTemplate != nil {
		eDevice.ValueTemplate = iDevice.ValueTemplate
	}
	eDevice.Initialize()
	return eDevice
}
