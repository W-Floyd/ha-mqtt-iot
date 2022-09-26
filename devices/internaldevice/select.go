package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Select struct {
	AvailabilityMode       *string     `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string     `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability           *([]string) `json:"availability,omitempty"`
	CommandTemplate        *string     `json:"command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `command_topic`."
	Command                *([]string) `json:"command,omitempty"`
	EnabledByDefault       *bool       `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string     `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string     `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   *string     `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string     `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributes         *([]string) `json:"json_attributes,omitempty"`
	Name                   *string     `json:"name,omitempty"`       // "The name of the Select."
	ObjectId               *string     `json:"object_id,omitempty"`  // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic             *bool       `json:"optimistic,omitempty"` // "Flag that defines if the select works in optimistic mode."
	Options                *([]string) `json:"options,omitempty"`    // "List of options that can be selected. An empty list or a list with a single item is allowed."
	Qos                    *int        `json:"qos,omitempty"`        // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain                 *bool       `json:"retain,omitempty"`     // "If the published message should have the retain flag on or not."
	State                  *([]string) `json:"state,omitempty"`
	UniqueId               *string     `json:"unique_id,omitempty"`      // "An ID that uniquely identifies this Select. If two Selects have the same unique ID Home Assistant will raise an exception."
	ValueTemplate          *string     `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the value."
	MQTT                   struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}

func (iDevice Select) Translate() externaldevice.Select {
	eDevice := externaldevice.Select{}
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
	if iDevice.Icon != nil {
		eDevice.Icon = iDevice.Icon
	}
	if iDevice.JsonAttributesTemplate != nil {
		eDevice.JsonAttributesTemplate = iDevice.JsonAttributesTemplate
	}
	if iDevice.JsonAttributes != nil {
		eDevice.JsonAttributesFunc = common.ConstructCommandFunc(*iDevice.JsonAttributes)
	}
	if iDevice.Name != nil {
		eDevice.Name = iDevice.Name
	}
	if iDevice.ObjectId != nil {
		eDevice.ObjectId = iDevice.ObjectId
	}
	if iDevice.Optimistic != nil {
		eDevice.Optimistic = iDevice.Optimistic
	}
	if iDevice.Options != nil {
		eDevice.Options = iDevice.Options
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
