package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type LawnMower struct {
	ActivityState          *([]string) `json:"activity_state,omitempty"`           // ActivityState for the LawnMower
	ActivityValueTemplate  *string     `json:"activity_value_template,omitempty"`  // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the value."
	AvailabilityMode       *string     `json:"availability_mode,omitempty"`        // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string     `json:"availability_template,omitempty"`    // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability, the result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability           *([]string) `json:"availability,omitempty"`             // Availability for the LawnMower
	DockCommandTemplate    *string     `json:"dock_command_template,omitempty"`    // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `dock_command_topic`. The `value` parameter in the template will be set to `dock`."
	DockCommand            *([]string) `json:"dock_command,omitempty"`             // DockCommand for the LawnMower
	EnabledByDefault       *bool       `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string     `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of the incoming payload."
	EntityCategory         *string     `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   *string     `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string     `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributes         *([]string) `json:"json_attributes,omitempty"`          // JsonAttributes for the LawnMower
	Name                   *string     `json:"name,omitempty"`                     // "The name of the lawn mower. Can be set to `null` if only the device name is relevant."
	ObjectId               *string     `json:"object_id,omitempty"`                // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic             *bool       `json:"optimistic,omitempty"`               // "Flag that defines if the lawn mower works in optimistic mode."
	PauseCommandTemplate   *string     `json:"pause_command_template,omitempty"`   // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `pause_command_topic`. The `value` parameter in the template will be set to `pause`."
	PauseCommand           *([]string) `json:"pause_command,omitempty"`            // PauseCommand for the LawnMower
	Qos                    *int        `json:"qos,omitempty"`                      // "The maximum QoS level to be used when receiving and publishing messages."
	Retain                 *bool       `json:"retain,omitempty"`                   // "If the published message should have the retain flag on or not."
	StartMowingCommand     *([]string) `json:"start_mowing_command,omitempty"`     // StartMowingCommand for the LawnMower
	StartMowingTemplate    *string     `json:"start_mowing_template,omitempty"`    // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `dock_command_topic`. The `value` parameter in the template will be set to `dock`."
	UniqueId               *string     `json:"unique_id,omitempty"`                // "An ID that uniquely identifies this lawn mower. If two lawn mowers have the same unique ID, Home Assistant will raise an exception."
	MQTT                   struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}

func (iDevice LawnMower) Translate() externaldevice.LawnMower {
	eDevice := externaldevice.LawnMower{}
	eDevice.MQTT = new(externaldevice.MQTTFields)
	if iDevice.MQTT.ForceUpdate != nil {
		eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	}
	if iDevice.MQTT.UpdateInterval != nil {
		eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	}
	if iDevice.ActivityState != nil {
		eDevice.ActivityStateFunc = common.ConstructStateFunc(*iDevice.ActivityState)
	}
	if iDevice.ActivityValueTemplate != nil {
		eDevice.ActivityValueTemplate = iDevice.ActivityValueTemplate
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
	if iDevice.DockCommandTemplate != nil {
		eDevice.DockCommandTemplate = iDevice.DockCommandTemplate
	}
	if iDevice.DockCommand != nil {
		eDevice.DockCommandFunc = common.ConstructCommandFunc(*iDevice.DockCommand)
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
	if iDevice.PauseCommandTemplate != nil {
		eDevice.PauseCommandTemplate = iDevice.PauseCommandTemplate
	}
	if iDevice.PauseCommand != nil {
		eDevice.PauseCommandFunc = common.ConstructCommandFunc(*iDevice.PauseCommand)
	}
	if iDevice.Qos != nil {
		eDevice.Qos = iDevice.Qos
	}
	if iDevice.Retain != nil {
		eDevice.Retain = iDevice.Retain
	}
	if iDevice.StartMowingCommand != nil {
		eDevice.StartMowingCommandFunc = common.ConstructCommandFunc(*iDevice.StartMowingCommand)
	}
	if iDevice.StartMowingTemplate != nil {
		eDevice.StartMowingTemplate = iDevice.StartMowingTemplate
	}
	if iDevice.UniqueId != nil {
		eDevice.UniqueId = iDevice.UniqueId
	}
	if iDevice.Availability == nil {
		eDevice.AvailabilityFunc = common.AvailabilityFunc
	}
	eDevice.Initialize()
	return eDevice
}
