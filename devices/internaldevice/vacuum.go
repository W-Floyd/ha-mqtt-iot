package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Vacuum struct {
	AvailabilityMode       *string     `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string     `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability           *([]string) `json:"availability,omitempty"`
	Command                *([]string) `json:"command,omitempty"`
	Encoding               *string     `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	FanSpeedList           *([]string) `json:"fan_speed_list,omitempty"`           // "List of possible fan speeds for the vacuum."
	JsonAttributesTemplate *string     `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributes         *([]string) `json:"json_attributes,omitempty"`
	Name                   *string     `json:"name,omitempty"`                   // "The name of the vacuum."
	ObjectId               *string     `json:"object_id,omitempty"`              // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable       *string     `json:"payload_available,omitempty"`      // "The payload that represents the available state."
	PayloadCleanSpot       *string     `json:"payload_clean_spot,omitempty"`     // "The payload to send to the `command_topic` to begin a spot cleaning cycle."
	PayloadLocate          *string     `json:"payload_locate,omitempty"`         // "The payload to send to the `command_topic` to locate the vacuum (typically plays a song)."
	PayloadNotAvailable    *string     `json:"payload_not_available,omitempty"`  // "The payload that represents the unavailable state."
	PayloadPause           *string     `json:"payload_pause,omitempty"`          // "The payload to send to the `command_topic` to pause the vacuum."
	PayloadReturnToBase    *string     `json:"payload_return_to_base,omitempty"` // "The payload to send to the `command_topic` to tell the vacuum to return to base."
	PayloadStart           *string     `json:"payload_start,omitempty"`          // "The payload to send to the `command_topic` to begin the cleaning cycle."
	PayloadStop            *string     `json:"payload_stop,omitempty"`           // "The payload to send to the `command_topic` to stop cleaning."
	Qos                    *int        `json:"qos,omitempty"`                    // "The maximum QoS level of the state topic."
	Retain                 *bool       `json:"retain,omitempty"`                 // "If the published message should have the retain flag on or not."
	Schema                 *string     `json:"schema,omitempty"`                 // "The schema to use. Must be `state` to select the state schema."
	SendCommand            *([]string) `json:"send_command,omitempty"`
	SetFanSpeed            *([]string) `json:"set_fan_speed,omitempty"`
	State                  *([]string) `json:"state,omitempty"`
	SupportedFeatures      *([]string) `json:"supported_features,omitempty"` // "List of features that the vacuum supports (possible values are `start`, `stop`, `pause`, `return_home`, `battery`, `status`, `locate`, `clean_spot`, `fan_speed`, `send_command`)."
	UniqueId               *string     `json:"unique_id,omitempty"`          // "An ID that uniquely identifies this vacuum. If two vacuums have the same unique ID, Home Assistant will raise an exception."
	MQTT                   struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}

func (iDevice Vacuum) Translate() externaldevice.Vacuum {
	eDevice := externaldevice.Vacuum{}
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
	if iDevice.Encoding != nil {
		eDevice.Encoding = iDevice.Encoding
	}
	if iDevice.FanSpeedList != nil {
		eDevice.FanSpeedList = iDevice.FanSpeedList
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
	if iDevice.PayloadAvailable != nil {
		eDevice.PayloadAvailable = iDevice.PayloadAvailable
	}
	if iDevice.PayloadCleanSpot != nil {
		eDevice.PayloadCleanSpot = iDevice.PayloadCleanSpot
	}
	if iDevice.PayloadLocate != nil {
		eDevice.PayloadLocate = iDevice.PayloadLocate
	}
	if iDevice.PayloadNotAvailable != nil {
		eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	}
	if iDevice.PayloadPause != nil {
		eDevice.PayloadPause = iDevice.PayloadPause
	}
	if iDevice.PayloadReturnToBase != nil {
		eDevice.PayloadReturnToBase = iDevice.PayloadReturnToBase
	}
	if iDevice.PayloadStart != nil {
		eDevice.PayloadStart = iDevice.PayloadStart
	}
	if iDevice.PayloadStop != nil {
		eDevice.PayloadStop = iDevice.PayloadStop
	}
	if iDevice.Qos != nil {
		eDevice.Qos = iDevice.Qos
	}
	if iDevice.Retain != nil {
		eDevice.Retain = iDevice.Retain
	}
	if iDevice.Schema != nil {
		eDevice.Schema = iDevice.Schema
	}
	if iDevice.SendCommand != nil {
		eDevice.SendCommandFunc = common.ConstructCommandFunc(*iDevice.SendCommand)
	}
	if iDevice.SetFanSpeed != nil {
		eDevice.SetFanSpeedFunc = common.ConstructCommandFunc(*iDevice.SetFanSpeed)
	}
	if iDevice.State != nil {
		eDevice.StateFunc = common.ConstructStateFunc(*iDevice.State)
	}
	if iDevice.SupportedFeatures != nil {
		eDevice.SupportedFeatures = iDevice.SupportedFeatures
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
