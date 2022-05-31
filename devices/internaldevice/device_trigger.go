package InternalDevice

import externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
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
	AutomationType string `json:"automation_type"` // "The type of automation, must be 'trigger'."
	Payload        string `json:"payload"`         // "Optional payload to match the payload being sent over the topic."
	Qos            int    `json:"qos"`             // "The maximum QoS level to be used when receiving messages."
	Subtype        string `json:"subtype"`         // "The subtype of the trigger, e.g. `button_1`. Entries supported by the frontend: `turn_on`, `turn_off`, `button_1`, `button_2`, `button_3`, `button_4`, `button_5`, `button_6`. If set to an unsupported value, will render as `subtype type`, e.g. `left_button pressed` with `type` set to `button_short_press` and `subtype` set to `left_button`"
	Topic          string `json:"topic"`           // "The MQTT topic subscribed to receive trigger events."
	Type           string `json:"type"`            // "The type of the trigger, e.g. `button_short_press`. Entries supported by the frontend: `button_short_press`, `button_short_release`, `button_long_press`, `button_long_release`, `button_double_press`, `button_triple_press`, `button_quadruple_press`, `button_quintuple_press`. If set to an unsupported value, will render as `subtype type`, e.g. `button_1 spammed` with `type` set to `spammed` and `subtype` set to `button_1`"
	ValueTemplate  string `json:"value_template"`  // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the value."
	MQTT           struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}