package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Lock struct {
	AvailabilityMode       *string     `json:"availability_mode,omitempty"`        // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string     `json:"availability_template,omitempty"`    // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability           *([]string) `json:"availability,omitempty"`             // Availability for the Lock
	CodeFormat             *string     `json:"code_format,omitempty"`              // "A regular expression to validate a supplied code when it is set during the service call to `open`, `lock` or `unlock` the MQTT lock."
	CommandTemplate        *string     `json:"command_template,omitempty"`         // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to generate the payload to send to `command_topic`. The lock command template accepts the parameters `value` and `code`. The `value` parameter will contain the configured value for either `payload_open`, `payload_lock` or `payload_unlock`. The `code` parameter is set during the service call to `open`, `lock` or `unlock` the MQTT lock and will be set `None` if no code was passed."
	Command                *([]string) `json:"command,omitempty"`                  // Command for the Lock
	EnabledByDefault       *bool       `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string     `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string     `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   *string     `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string     `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributes         *([]string) `json:"json_attributes,omitempty"`          // JsonAttributes for the Lock
	Name                   *string     `json:"name,omitempty"`                     // "The name of the lock."
	ObjectId               *string     `json:"object_id,omitempty"`                // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic             *bool       `json:"optimistic,omitempty"`               // "Flag that defines if lock works in optimistic mode."
	PayloadAvailable       *string     `json:"payload_available,omitempty"`        // "The payload that represents the available state."
	PayloadLock            *string     `json:"payload_lock,omitempty"`             // "The payload sent to the lock to lock it."
	PayloadNotAvailable    *string     `json:"payload_not_available,omitempty"`    // "The payload that represents the unavailable state."
	PayloadOpen            *string     `json:"payload_open,omitempty"`             // "The payload sent to the lock to open it."
	PayloadUnlock          *string     `json:"payload_unlock,omitempty"`           // "The payload sent to the lock to unlock it."
	Qos                    *int        `json:"qos,omitempty"`                      // "The maximum QoS level of the state topic. It will also be used for messages published to command topic."
	Retain                 *bool       `json:"retain,omitempty"`                   // "If the published message should have the retain flag on or not."
	StateJammed            *string     `json:"state_jammed,omitempty"`             // "The payload sent to `state_topic` by the lock when it's jammed."
	StateLocked            *string     `json:"state_locked,omitempty"`             // "The payload sent to `state_topic` by the lock when it's locked."
	StateLocking           *string     `json:"state_locking,omitempty"`            // "The payload sent to `state_topic` by the lock when it's locking."
	State                  *([]string) `json:"state,omitempty"`                    // State for the Lock
	StateUnlocked          *string     `json:"state_unlocked,omitempty"`           // "The payload sent to `state_topic` by the lock when it's unlocked."
	StateUnlocking         *string     `json:"state_unlocking,omitempty"`          // "The payload sent to `state_topic` by the lock when it's unlocking."
	UniqueId               *string     `json:"unique_id,omitempty"`                // "An ID that uniquely identifies this lock. If two locks have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          *string     `json:"value_template,omitempty"`           // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract a state value from the payload."
	MQTT                   struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}

func (iDevice Lock) Translate() externaldevice.Lock {
	eDevice := externaldevice.Lock{}
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
	if iDevice.CodeFormat != nil {
		eDevice.CodeFormat = iDevice.CodeFormat
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
	if iDevice.PayloadAvailable != nil {
		eDevice.PayloadAvailable = iDevice.PayloadAvailable
	}
	if iDevice.PayloadLock != nil {
		eDevice.PayloadLock = iDevice.PayloadLock
	}
	if iDevice.PayloadNotAvailable != nil {
		eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	}
	if iDevice.PayloadOpen != nil {
		eDevice.PayloadOpen = iDevice.PayloadOpen
	}
	if iDevice.PayloadUnlock != nil {
		eDevice.PayloadUnlock = iDevice.PayloadUnlock
	}
	if iDevice.Qos != nil {
		eDevice.Qos = iDevice.Qos
	}
	if iDevice.Retain != nil {
		eDevice.Retain = iDevice.Retain
	}
	if iDevice.StateJammed != nil {
		eDevice.StateJammed = iDevice.StateJammed
	}
	if iDevice.StateLocked != nil {
		eDevice.StateLocked = iDevice.StateLocked
	}
	if iDevice.StateLocking != nil {
		eDevice.StateLocking = iDevice.StateLocking
	}
	if iDevice.State != nil {
		eDevice.StateFunc = common.ConstructStateFunc(*iDevice.State)
	}
	if iDevice.StateUnlocked != nil {
		eDevice.StateUnlocked = iDevice.StateUnlocked
	}
	if iDevice.StateUnlocking != nil {
		eDevice.StateUnlocking = iDevice.StateUnlocking
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
