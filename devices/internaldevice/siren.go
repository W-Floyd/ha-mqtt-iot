package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
func (iDevice Siren) Translate() externaldevice.Siren {
	eDevice := externaldevice.Siren{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.AvailableTones = iDevice.AvailableTones
	eDevice.CommandOffTemplate = iDevice.CommandOffTemplate
	eDevice.CommandTemplate = iDevice.CommandTemplate
	eDevice.CommandFunc = common.ConstructCommandFunc(iDevice.Command)
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.Icon = iDevice.Icon
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.Optimistic = iDevice.Optimistic
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.PayloadOff = iDevice.PayloadOff
	eDevice.PayloadOn = iDevice.PayloadOn
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.StateOff = iDevice.StateOff
	eDevice.StateOn = iDevice.StateOn
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.StateValueTemplate = iDevice.StateValueTemplate
	eDevice.SupportDuration = iDevice.SupportDuration
	eDevice.SupportVolumeSet = iDevice.SupportVolumeSet
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.Initialize()
	return eDevice
}

type Siren struct {
	AvailableTones      []string `json:"available_tones"`      // "A list of available tones the siren supports. When configured, this enables the support for setting a `tone` and enables the `tone` state attribute."
	CommandOffTemplate  string   `json:"command_off_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate a custom payload to send to `command_topic` when the siren turn off service is called. By default `command_template` will be used as template for service turn off. The variable `value` will be assigned with the configured `payload_off` setting."
	CommandTemplate     string   `json:"command_template"`     // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate a custom payload to send to `command_topic`. The variable `value` will be assigned with the configured `payload_on` or `payload_off` setting. The siren turn on service parameters `tone`, `volume_level` or `duration` can be used as variables in the template. When operation in optimistic mode the corresponding state attributes will be set. Turn on parameters will be filtered if a device misses the support."
	Command             []string `json:"command"`
	EnabledByDefault    bool     `json:"enabled_by_default"`    // "Flag which defines if the entity should be enabled when first added."
	Encoding            string   `json:"encoding"`              // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory      string   `json:"entity_category"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                string   `json:"icon"`                  // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Name                string   `json:"name"`                  // "The name to use when displaying this siren."
	ObjectId            string   `json:"object_id"`             // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic          bool     `json:"optimistic"`            // "Flag that defines if siren works in optimistic mode."
	PayloadAvailable    string   `json:"payload_available"`     // "The payload that represents the available state."
	PayloadNotAvailable string   `json:"payload_not_available"` // "The payload that represents the unavailable state."
	PayloadOff          string   `json:"payload_off"`           // "The payload that represents `off` state. If specified, will be used for both comparing to the value in the `state_topic` (see `value_template` and `state_off` for details) and sending as `off` command to the `command_topic`."
	PayloadOn           string   `json:"payload_on"`            // "The payload that represents `on` state. If specified, will be used for both comparing to the value in the `state_topic` (see `value_template` and `state_on`  for details) and sending as `on` command to the `command_topic`."
	Qos                 int      `json:"qos"`                   // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain              bool     `json:"retain"`                // "If the published message should have the retain flag on or not."
	StateOff            string   `json:"state_off"`             // "The payload that represents the `off` state. Used when value that represents `off` state in the `state_topic` is different from value that should be sent to the `command_topic` to turn the device `off`."
	StateOn             string   `json:"state_on"`              // "The payload that represents the `on` state. Used when value that represents `on` state in the `state_topic` is different from value that should be sent to the `command_topic` to turn the device `on`."
	State               []string `json:"state"`
	StateValueTemplate  string   `json:"state_value_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's state from the `state_topic`. To determine the siren's state result of this template will be compared to `state_on` and `state_off`. Alternatively `value_template` can be used to render to a valid JSON payload."
	SupportDuration     bool     `json:"support_duration"`     // "Set to `true` if the MQTT siren supports the `duration` service turn on parameter and enables the `duration` state attribute."
	SupportVolumeSet    bool     `json:"support_volume_set"`   // "Set to `true` if the MQTT siren supports the `volume_set` service turn on parameter and enables the `volume_level` state attribute."
	UniqueId            string   `json:"unique_id"`            // "An ID that uniquely identifies this siren device. If two sirens have the same unique ID, Home Assistant will raise an exception."
	MQTT                struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
