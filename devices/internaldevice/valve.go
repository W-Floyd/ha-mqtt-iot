package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Valve struct {
	AvailabilityMode       *string     `json:"availability_mode,omitempty"`        // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string     `json:"availability_template,omitempty"`    // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the device's availability from the `availability_topic`. To determine the devices's availability, the result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability           *([]string) `json:"availability,omitempty"`             // Availability for the Valve
	CommandTemplate        *string     `json:"command_template,omitempty"`         // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`."
	Command                *([]string) `json:"command,omitempty"`                  // Command for the Valve
	DeviceClass            *string     `json:"device_class,omitempty"`             // "Sets the [class of the device](/integrations/valve/), changing the device state and icon that is displayed on the frontend. The `device_class` can be `null`."
	EnabledByDefault       *bool       `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string     `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string     `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   *string     `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string     `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. A usage example can be found in the [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributes         *([]string) `json:"json_attributes,omitempty"`          // JsonAttributes for the Valve
	Name                   *string     `json:"name,omitempty"`                     // "The name of the valve. Can be set to `null` if only the device name is relevant."
	ObjectId               *string     `json:"object_id,omitempty"`                // "Used instead of `name` to have the `entity_id` generated automatically."
	Optimistic             *bool       `json:"optimistic,omitempty"`               // "Flag that defines if a switch works in optimistic mode."
	PayloadAvailable       *string     `json:"payload_available,omitempty"`        // "The payload that represents the online state."
	PayloadClose           *string     `json:"payload_close,omitempty"`            // "The command payload that closes the valve. Is only used when `reports_position` is set to `false` (default). The `payload_close` is not allowed if `reports_position` is set to `true`. Can be set to `null` to disable the valve's close option."
	PayloadNotAvailable    *string     `json:"payload_not_available,omitempty"`    // "The payload that represents the offline state."
	PayloadOpen            *string     `json:"payload_open,omitempty"`             // "The command payload that opens the valve. Is only used when `reports_position` is set to `false` (default). The `payload_open` is not allowed if `reports_position` is set to `true`. Can be set to `null` to disable the valve's open option."
	PayloadStop            *string     `json:"payload_stop,omitempty"`             // "The command payload that stops the valve. When not configured, the valve will not support the `valve.stop` service."
	PositionClosed         *int        `json:"position_closed,omitempty"`          // "Number which represents closed position. The valve's position will be scaled to the(`position_closed`...`position_open`) range when a service is called and scaled back when a value is received."
	PositionOpen           *int        `json:"position_open,omitempty"`            // "Number which represents open position. The valve's position will be scaled to (`position_closed`...`position_open`) range when a service is called and scaled back when a value is received."
	Qos                    *int        `json:"qos,omitempty"`                      // "The maximum QoS level to be used when receiving and publishing messages."
	ReportsPosition        *bool       `json:"reports_position,omitempty"`         // "Set to `true` if the value reports the position or supports setting the position. Enabling the `reports_position` option will cause the position to be published instead of a payload defined by `payload_open`, `payload_close` or `payload_stop`. When receiving messages, `state_topic` will accept numeric payloads or one of the following state messages: `open`, `opening`, `closed`, or `closing`."
	Retain                 *bool       `json:"retain,omitempty"`                   // "Defines if published messages should have the retain flag set."
	StateClosed            *string     `json:"state_closed,omitempty"`             // "The payload that represents the closed state. Is only allowed when `reports_position` is set to `False` (default)."
	StateClosing           *string     `json:"state_closing,omitempty"`            // "The payload that represents the closing state."
	StateOpen              *string     `json:"state_open,omitempty"`               // "The payload that represents the open state. Is only allowed when `reports_position` is set to `False` (default)."
	StateOpening           *string     `json:"state_opening,omitempty"`            // "The payload that represents the opening state."
	State                  *([]string) `json:"state,omitempty"`                    // State for the Valve
	UniqueId               *string     `json:"unique_id,omitempty"`                // "An ID that uniquely identifies this valve. If two valves have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          *string     `json:"value_template,omitempty"`           // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) that can be used to extract the payload for the `state_topic` topic. The rendered value should be a defined state payload or, if reporting a `position` is supported and `reports_position` is set to `true`, a numeric value is expected representing the position. See also `state_topic`."
	MQTT                   struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}

func (iDevice Valve) Translate() externaldevice.Valve {
	eDevice := externaldevice.Valve{}
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
	if iDevice.PayloadClose != nil {
		eDevice.PayloadClose = iDevice.PayloadClose
	}
	if iDevice.PayloadNotAvailable != nil {
		eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	}
	if iDevice.PayloadOpen != nil {
		eDevice.PayloadOpen = iDevice.PayloadOpen
	}
	if iDevice.PayloadStop != nil {
		eDevice.PayloadStop = iDevice.PayloadStop
	}
	if iDevice.PositionClosed != nil {
		eDevice.PositionClosed = iDevice.PositionClosed
	}
	if iDevice.PositionOpen != nil {
		eDevice.PositionOpen = iDevice.PositionOpen
	}
	if iDevice.Qos != nil {
		eDevice.Qos = iDevice.Qos
	}
	if iDevice.ReportsPosition != nil {
		eDevice.ReportsPosition = iDevice.ReportsPosition
	}
	if iDevice.Retain != nil {
		eDevice.Retain = iDevice.Retain
	}
	if iDevice.StateClosed != nil {
		eDevice.StateClosed = iDevice.StateClosed
	}
	if iDevice.StateClosing != nil {
		eDevice.StateClosing = iDevice.StateClosing
	}
	if iDevice.StateOpen != nil {
		eDevice.StateOpen = iDevice.StateOpen
	}
	if iDevice.StateOpening != nil {
		eDevice.StateOpening = iDevice.StateOpening
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
