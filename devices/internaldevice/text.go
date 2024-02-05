package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Text struct {
	AvailabilityMode       *string     `json:"availability_mode,omitempty"`        // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string     `json:"availability_template,omitempty"`    // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability           *([]string) `json:"availability,omitempty"`             // Availability for the Text
	CommandTemplate        *string     `json:"command_template,omitempty"`         // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	Command                *([]string) `json:"command,omitempty"`                  // Command for the Text
	EnabledByDefault       *bool       `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string     `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string     `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	JsonAttributesTemplate *string     `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributes         *([]string) `json:"json_attributes,omitempty"`          // JsonAttributes for the Text
	Max                    *int        `json:"max,omitempty"`                      // "The maximum size of a text being set or received (maximum is 255)."
	Min                    *int        `json:"min,omitempty"`                      // "The minimum size of a text being set or received."
	Mode                   *string     `json:"mode,omitempty"`                     // "The mode off the text entity. Must be either `text` or `password`."
	Name                   *string     `json:"name,omitempty"`                     // "The name of the text entity. Can be set to `null` if only the device name is relevant."
	ObjectId               *string     `json:"object_id,omitempty"`                // "Used instead of `name` for automatic generation of `entity_id`"
	Pattern                *string     `json:"pattern,omitempty"`                  // "A valid regular expression the text being set or received must match with."
	Qos                    *int        `json:"qos,omitempty"`                      // "The maximum QoS level to be used when receiving and publishing messages."
	Retain                 *bool       `json:"retain,omitempty"`                   // "If the published message should have the retain flag on or not."
	State                  *([]string) `json:"state,omitempty"`                    // State for the Text
	UniqueId               *string     `json:"unique_id,omitempty"`                // "An ID that uniquely identifies this Select. If two Selects have the same unique ID Home Assistant will raise an exception."
	ValueTemplate          *string     `json:"value_template,omitempty"`           // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the text state value from the payload received on `state_topic`."
	MQTT                   struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}

func (iDevice Text) Translate() externaldevice.Text {
	eDevice := externaldevice.Text{}
	eDevice.MQTT = new(externaldevice.MQTTFields)
	if iDevice.MQTT.ForceUpdate != nil {
		eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	}
	if iDevice.MQTT.UpdateInterval != nil {
		eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	}
	if iDevice.AvailabilityMode != nil {
		eDevice.AvailabilityMode = iDevice.AvailabilityMode
	}
	if iDevice.AvailabilityTemplate != nil {
		eDevice.AvailabilityTemplate = iDevice.AvailabilityTemplate
	}
	if iDevice.Availability != nil {
		eDevice.AvailabilityFunc = common.ConstructStateFunc(*iDevice.Availability)
	}
	if iDevice.CommandTemplate != nil {
		eDevice.CommandTemplate = iDevice.CommandTemplate
	}
	if iDevice.Command != nil {
		eDevice.CommandFunc = common.ConstructCommandFunc(*iDevice.Command)
	}
	if iDevice.EnabledByDefault != nil {
		eDevice.EnabledByDefault = iDevice.EnabledByDefault
	}
	if iDevice.Encoding != nil {
		eDevice.Encoding = iDevice.Encoding
	}
	if iDevice.EntityCategory != nil {
		eDevice.EntityCategory = iDevice.EntityCategory
	}
	if iDevice.JsonAttributesTemplate != nil {
		eDevice.JsonAttributesTemplate = iDevice.JsonAttributesTemplate
	}
	if iDevice.JsonAttributes != nil {
		eDevice.JsonAttributesFunc = common.ConstructCommandFunc(*iDevice.JsonAttributes)
	}
	if iDevice.Max != nil {
		eDevice.Max = iDevice.Max
	}
	if iDevice.Min != nil {
		eDevice.Min = iDevice.Min
	}
	if iDevice.Mode != nil {
		eDevice.Mode = iDevice.Mode
	}
	if iDevice.Name != nil {
		eDevice.Name = iDevice.Name
	}
	if iDevice.ObjectId != nil {
		eDevice.ObjectId = iDevice.ObjectId
	}
	if iDevice.Pattern != nil {
		eDevice.Pattern = iDevice.Pattern
	}
	if iDevice.Qos != nil {
		eDevice.Qos = iDevice.Qos
	}
	if iDevice.Retain != nil {
		eDevice.Retain = iDevice.Retain
	}
	if iDevice.State != nil {
		eDevice.StateFunc = common.ConstructStateFunc(*iDevice.State)
	}
	if iDevice.UniqueId != nil {
		eDevice.UniqueId = iDevice.UniqueId
	}
	if iDevice.ValueTemplate != nil {
		eDevice.ValueTemplate = iDevice.ValueTemplate
	}
	if iDevice.Availability == nil {
		eDevice.AvailabilityFunc = common.AvailabilityFunc
	}
	eDevice.Initialize()
	return eDevice
}
