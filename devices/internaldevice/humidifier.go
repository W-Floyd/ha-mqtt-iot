package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Humidifier struct {
	ActionTemplate                *string     `json:"action_template,omitempty"`                  // "A template to render the value received on the `action_topic` with."
	Action                        *([]string) `json:"action,omitempty"`                           // Action for the Humidifier
	AvailabilityMode              *string     `json:"availability_mode,omitempty"`                // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate          *string     `json:"availability_template,omitempty"`            // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability                  *([]string) `json:"availability,omitempty"`                     // Availability for the Humidifier
	CommandTemplate               *string     `json:"command_template,omitempty"`                 // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	Command                       *([]string) `json:"command,omitempty"`                          // Command for the Humidifier
	CurrentHumidityTemplate       *string     `json:"current_humidity_template,omitempty"`        // "A template with which the value received on `current_humidity_topic` will be rendered."
	CurrentHumidity               *([]string) `json:"current_humidity,omitempty"`                 // CurrentHumidity for the Humidifier
	DeviceClass                   *string     `json:"device_class,omitempty"`                     // "The device class of the MQTT device. Must be either `humidifier`, `dehumidifier` or `null`."
	EnabledByDefault              *bool       `json:"enabled_by_default,omitempty"`               // "Flag which defines if the entity should be enabled when first added."
	Encoding                      *string     `json:"encoding,omitempty"`                         // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory                *string     `json:"entity_category,omitempty"`                  // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                          *string     `json:"icon,omitempty"`                             // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate        *string     `json:"json_attributes_template,omitempty"`         // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributes                *([]string) `json:"json_attributes,omitempty"`                  // JsonAttributes for the Humidifier
	MaxHumidity                   *int        `json:"max_humidity,omitempty"`                     // "The minimum target humidity percentage that can be set."
	MinHumidity                   *int        `json:"min_humidity,omitempty"`                     // "The maximum target humidity percentage that can be set."
	ModeCommandTemplate           *string     `json:"mode_command_template,omitempty"`            // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `mode_command_topic`."
	ModeCommand                   *([]string) `json:"mode_command,omitempty"`                     // ModeCommand for the Humidifier
	ModeStateTemplate             *string     `json:"mode_state_template,omitempty"`              // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value for the humidifier `mode` state."
	ModeState                     *([]string) `json:"mode_state,omitempty"`                       // ModeState for the Humidifier
	Modes                         *([]string) `json:"modes,omitempty"`                            // "List of available modes this humidifier is capable of running at. Common examples include `normal`, `eco`, `away`, `boost`, `comfort`, `home`, `sleep`, `auto` and `baby`. These examples offer built-in translations but other custom modes are allowed as well.  This attribute ust be configured together with the `mode_command_topic` attribute."
	Name                          *string     `json:"name,omitempty"`                             // "The name of the humidifier."
	ObjectId                      *string     `json:"object_id,omitempty"`                        // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic                    *bool       `json:"optimistic,omitempty"`                       // "Flag that defines if humidifier works in optimistic mode"
	PayloadAvailable              *string     `json:"payload_available,omitempty"`                // "The payload that represents the available state."
	PayloadNotAvailable           *string     `json:"payload_not_available,omitempty"`            // "The payload that represents the unavailable state."
	PayloadOff                    *string     `json:"payload_off,omitempty"`                      // "The payload that represents the stop state."
	PayloadOn                     *string     `json:"payload_on,omitempty"`                       // "The payload that represents the running state."
	PayloadResetHumidity          *string     `json:"payload_reset_humidity,omitempty"`           // "A special payload that resets the `target_humidity` state attribute to `None` when received at the `target_humidity_state_topic`."
	PayloadResetMode              *string     `json:"payload_reset_mode,omitempty"`               // "A special payload that resets the `mode` state attribute to `None` when received at the `mode_state_topic`. When received at `current_humidity_topic` it will reset the current humidity state."
	Qos                           *int        `json:"qos,omitempty"`                              // "The maximum QoS level of the state topic."
	Retain                        *bool       `json:"retain,omitempty"`                           // "If the published message should have the retain flag on or not."
	State                         *([]string) `json:"state,omitempty"`                            // State for the Humidifier
	StateValueTemplate            *string     `json:"state_value_template,omitempty"`             // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value from the state."
	TargetHumidityCommandTemplate *string     `json:"target_humidity_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `target_humidity_command_topic`."
	TargetHumidityCommand         *([]string) `json:"target_humidity_command,omitempty"`          // TargetHumidityCommand for the Humidifier
	TargetHumidityStateTemplate   *string     `json:"target_humidity_state_template,omitempty"`   // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a value for the humidifier `target_humidity` state."
	TargetHumidityState           *([]string) `json:"target_humidity_state,omitempty"`            // TargetHumidityState for the Humidifier
	UniqueId                      *string     `json:"unique_id,omitempty"`                        // "An ID that uniquely identifies this humidifier. If two humidifiers have the same unique ID, Home Assistant will raise an exception."
	MQTT                          struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}

func (iDevice Humidifier) Translate() externaldevice.Humidifier {
	eDevice := externaldevice.Humidifier{}
	eDevice.MQTT = new(externaldevice.MQTTFields)
	if iDevice.MQTT.ForceUpdate != nil {
		eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	}
	if iDevice.MQTT.UpdateInterval != nil {
		eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	}
	if iDevice.ActionTemplate != nil {
		eDevice.ActionTemplate = iDevice.ActionTemplate
	}
	if iDevice.Action != nil {
		eDevice.ActionFunc = common.ConstructCommandFunc(*iDevice.Action)
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
	if iDevice.CurrentHumidityTemplate != nil {
		eDevice.CurrentHumidityTemplate = iDevice.CurrentHumidityTemplate
	}
	if iDevice.CurrentHumidity != nil {
		eDevice.CurrentHumidityFunc = common.ConstructStateFunc(*iDevice.CurrentHumidity)
	}
	if iDevice.DeviceClass != nil {
		eDevice.DeviceClass = iDevice.DeviceClass
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
	if iDevice.MaxHumidity != nil {
		eDevice.MaxHumidity = iDevice.MaxHumidity
	}
	if iDevice.MinHumidity != nil {
		eDevice.MinHumidity = iDevice.MinHumidity
	}
	if iDevice.ModeCommandTemplate != nil {
		eDevice.ModeCommandTemplate = iDevice.ModeCommandTemplate
	}
	if iDevice.ModeCommand != nil {
		eDevice.ModeCommandFunc = common.ConstructCommandFunc(*iDevice.ModeCommand)
	}
	if iDevice.ModeStateTemplate != nil {
		eDevice.ModeStateTemplate = iDevice.ModeStateTemplate
	}
	if iDevice.ModeState != nil {
		eDevice.ModeStateFunc = common.ConstructStateFunc(*iDevice.ModeState)
	}
	if iDevice.Modes != nil {
		eDevice.Modes = iDevice.Modes
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
	if iDevice.PayloadAvailable != nil {
		eDevice.PayloadAvailable = iDevice.PayloadAvailable
	}
	if iDevice.PayloadNotAvailable != nil {
		eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	}
	if iDevice.PayloadOff != nil {
		eDevice.PayloadOff = iDevice.PayloadOff
	}
	if iDevice.PayloadOn != nil {
		eDevice.PayloadOn = iDevice.PayloadOn
	}
	if iDevice.PayloadResetHumidity != nil {
		eDevice.PayloadResetHumidity = iDevice.PayloadResetHumidity
	}
	if iDevice.PayloadResetMode != nil {
		eDevice.PayloadResetMode = iDevice.PayloadResetMode
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
	if iDevice.StateValueTemplate != nil {
		eDevice.StateValueTemplate = iDevice.StateValueTemplate
	}
	if iDevice.TargetHumidityCommandTemplate != nil {
		eDevice.TargetHumidityCommandTemplate = iDevice.TargetHumidityCommandTemplate
	}
	if iDevice.TargetHumidityCommand != nil {
		eDevice.TargetHumidityCommandFunc = common.ConstructCommandFunc(*iDevice.TargetHumidityCommand)
	}
	if iDevice.TargetHumidityStateTemplate != nil {
		eDevice.TargetHumidityStateTemplate = iDevice.TargetHumidityStateTemplate
	}
	if iDevice.TargetHumidityState != nil {
		eDevice.TargetHumidityStateFunc = common.ConstructStateFunc(*iDevice.TargetHumidityState)
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
