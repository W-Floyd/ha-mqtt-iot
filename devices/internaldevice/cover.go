package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Cover struct {
	AvailabilityMode       *string     `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string     `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability           *([]string) `json:"availability,omitempty"`
	Command                *([]string) `json:"command,omitempty"`
	DeviceClass            *string     `json:"device_class,omitempty"`             // "Sets the [class of the device](/integrations/cover/), changing the device state and icon that is displayed on the frontend."
	EnabledByDefault       *bool       `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string     `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string     `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   *string     `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string     `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributes         *([]string) `json:"json_attributes,omitempty"`
	Name                   *string     `json:"name,omitempty"`                  // "The name of the cover."
	ObjectId               *string     `json:"object_id,omitempty"`             // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic             *bool       `json:"optimistic,omitempty"`            // "Flag that defines if switch works in optimistic mode."
	PayloadAvailable       *string     `json:"payload_available,omitempty"`     // "The payload that represents the online state."
	PayloadClose           *string     `json:"payload_close,omitempty"`         // "The command payload that closes the cover."
	PayloadNotAvailable    *string     `json:"payload_not_available,omitempty"` // "The payload that represents the offline state."
	PayloadOpen            *string     `json:"payload_open,omitempty"`          // "The command payload that opens the cover."
	PayloadStop            *string     `json:"payload_stop,omitempty"`          // "The command payload that stops the cover."
	PositionClosed         *int        `json:"position_closed,omitempty"`       // "Number which represents closed position."
	PositionOpen           *int        `json:"position_open,omitempty"`         // "Number which represents open position."
	PositionTemplate       *string     `json:"position_template,omitempty"`     // "Defines a [template](/topics/templating/) that can be used to extract the payload for the `position_topic` topic. Within the template the following variables are available: `entity_id`, `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function;"
	Position               *([]string) `json:"position,omitempty"`
	Qos                    *int        `json:"qos,omitempty"`                   // "The maximum QoS level to be used when receiving and publishing messages."
	Retain                 *bool       `json:"retain,omitempty"`                // "Defines if published messages should have the retain flag set."
	SetPositionTemplate    *string     `json:"set_position_template,omitempty"` // "Defines a [template](/topics/templating/) to define the position to be sent to the `set_position_topic` topic. Incoming position value is available for use in the template `{% raw %}{{ position }}{% endraw %}`. Within the template the following variables are available: `entity_id`, `position`, the target position in percent; `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function;"
	SetPosition            *([]string) `json:"set_position,omitempty"`
	StateClosed            *string     `json:"state_closed,omitempty"`  // "The payload that represents the closed state."
	StateClosing           *string     `json:"state_closing,omitempty"` // "The payload that represents the closing state."
	StateOpen              *string     `json:"state_open,omitempty"`    // "The payload that represents the open state."
	StateOpening           *string     `json:"state_opening,omitempty"` // "The payload that represents the opening state."
	StateStopped           *string     `json:"state_stopped,omitempty"` // "The payload that represents the stopped state (for covers that do not report `open`/`closed` state)."
	State                  *([]string) `json:"state,omitempty"`
	TiltClosedValue        *int        `json:"tilt_closed_value,omitempty"`     // "The value that will be sent on a `close_cover_tilt` command."
	TiltCommandTemplate    *string     `json:"tilt_command_template,omitempty"` // "Defines a [template](/topics/templating/) that can be used to extract the payload for the `tilt_command_topic` topic. Within the template the following variables are available: `entity_id`, `tilt_position`, the target tilt position in percent; `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function;"
	TiltCommand            *([]string) `json:"tilt_command,omitempty"`
	TiltMax                *int        `json:"tilt_max,omitempty"`             // "The maximum tilt value."
	TiltMin                *int        `json:"tilt_min,omitempty"`             // "The minimum tilt value."
	TiltOpenedValue        *int        `json:"tilt_opened_value,omitempty"`    // "The value that will be sent on an `open_cover_tilt` command."
	TiltOptimistic         *bool       `json:"tilt_optimistic,omitempty"`      // "Flag that determines if tilt works in optimistic mode."
	TiltStatusTemplate     *string     `json:"tilt_status_template,omitempty"` // "Defines a [template](/topics/templating/) that can be used to extract the payload for the `tilt_status_topic` topic. Within the template the following variables are available: `entity_id`, `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function;"
	TiltStatus             *([]string) `json:"tilt_status,omitempty"`
	UniqueId               *string     `json:"unique_id,omitempty"`      // "An ID that uniquely identifies this cover. If two covers have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          *string     `json:"value_template,omitempty"` // "Defines a [template](/topics/templating/) that can be used to extract the payload for the `state_topic` topic."
	MQTT                   struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}

func (iDevice Cover) Translate() externaldevice.Cover {
	eDevice := externaldevice.Cover{}
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
	if iDevice.PositionTemplate != nil {
		eDevice.PositionTemplate = iDevice.PositionTemplate
	}
	if iDevice.Position != nil {
		eDevice.PositionFunc = common.ConstructStateFunc(*iDevice.Position)
	}
	if iDevice.Qos != nil {
		eDevice.Qos = iDevice.Qos
	}
	if iDevice.Retain != nil {
		eDevice.Retain = iDevice.Retain
	}
	if iDevice.SetPositionTemplate != nil {
		eDevice.SetPositionTemplate = iDevice.SetPositionTemplate
	}
	if iDevice.SetPosition != nil {
		eDevice.SetPositionFunc = common.ConstructCommandFunc(*iDevice.SetPosition)
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
	if iDevice.StateStopped != nil {
		eDevice.StateStopped = iDevice.StateStopped
	}
	if iDevice.State != nil {
		eDevice.StateFunc = common.ConstructStateFunc(*iDevice.State)
	}
	if iDevice.TiltClosedValue != nil {
		eDevice.TiltClosedValue = iDevice.TiltClosedValue
	}
	if iDevice.TiltCommandTemplate != nil {
		eDevice.TiltCommandTemplate = iDevice.TiltCommandTemplate
	}
	if iDevice.TiltCommand != nil {
		eDevice.TiltCommandFunc = common.ConstructCommandFunc(*iDevice.TiltCommand)
	}
	if iDevice.TiltMax != nil {
		eDevice.TiltMax = iDevice.TiltMax
	}
	if iDevice.TiltMin != nil {
		eDevice.TiltMin = iDevice.TiltMin
	}
	if iDevice.TiltOpenedValue != nil {
		eDevice.TiltOpenedValue = iDevice.TiltOpenedValue
	}
	if iDevice.TiltOptimistic != nil {
		eDevice.TiltOptimistic = iDevice.TiltOptimistic
	}
	if iDevice.TiltStatusTemplate != nil {
		eDevice.TiltStatusTemplate = iDevice.TiltStatusTemplate
	}
	if iDevice.TiltStatus != nil {
		eDevice.TiltStatusFunc = common.ConstructStateFunc(*iDevice.TiltStatus)
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
