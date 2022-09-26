package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type DeviceTrigger struct {
	AutomationType *string     `json:"automation_type,omitempty"` // "The type of automation, must be 'trigger'."
	Payload        *string     `json:"payload,omitempty"`         // "Optional payload to match the payload being sent over the topic."
	Qos            *int        `json:"qos,omitempty"`             // "The maximum QoS level to be used when receiving messages."
	Subtype        *string     `json:"subtype,omitempty"`         // "The subtype of the trigger, e.g. `button_1`. Entries supported by the frontend: `turn_on`, `turn_off`, `button_1`, `button_2`, `button_3`, `button_4`, `button_5`, `button_6`. If set to an unsupported value, will render as `subtype type`, e.g. `left_button pressed` with `type` set to `button_short_press` and `subtype` set to `left_button`"
	State          *([]string) `json:"state,omitempty"`
	Type           *string     `json:"type,omitempty"`           // "The type of the trigger, e.g. `button_short_press`. Entries supported by the frontend: `button_short_press`, `button_short_release`, `button_long_press`, `button_long_release`, `button_double_press`, `button_triple_press`, `button_quadruple_press`, `button_quintuple_press`. If set to an unsupported value, will render as `subtype type`, e.g. `button_1 spammed` with `type` set to `spammed` and `subtype` set to `button_1`"
	ValueTemplate  *string     `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the value."
	MQTT           struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}

func (iDevice DeviceTrigger) Translate() externaldevice.DeviceTrigger {
	eDevice := externaldevice.DeviceTrigger{}
	eDevice.MQTT = new(externaldevice.MQTTFields)
	if iDevice.MQTT.ForceUpdate != nil {
		eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	}
	if iDevice.MQTT.UpdateInterval != nil {
		eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	}
	if iDevice.AutomationType != nil {
		eDevice.AutomationType = iDevice.AutomationType
	}
	if iDevice.Payload != nil {
		eDevice.Payload = iDevice.Payload
	}
	if iDevice.Qos != nil {
		eDevice.Qos = iDevice.Qos
	}
	if iDevice.Subtype != nil {
		eDevice.Subtype = iDevice.Subtype
	}
	if iDevice.State != nil {
		eDevice.StateFunc = common.ConstructStateFunc(*iDevice.State)
	}
	if iDevice.Type != nil {
		eDevice.Type = iDevice.Type
	}
	if iDevice.ValueTemplate != nil {
		eDevice.ValueTemplate = iDevice.ValueTemplate
	}
	eDevice.Initialize()
	return eDevice
}
