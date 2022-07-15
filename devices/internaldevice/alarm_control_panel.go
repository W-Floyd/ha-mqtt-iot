package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
type AlarmControlPanel struct {
	AvailabilityMode       *string     `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string     `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability           *([]string) `json:"availability,omitempty"`
	Code                   *string     `json:"code,omitempty"`                  // "If defined, specifies a code to enable or disable the alarm in the frontend. Note that the code is validated locally and blocks sending MQTT messages to the remote device. For remote code validation, the code can be configured to either of the special values `REMOTE_CODE` (numeric code) or `REMOTE_CODE_TEXT` (text code). In this case, local code validation is bypassed but the frontend will still show a numeric or text code dialog. Use `command_template` to send the code to the remote device. Example configurations for remote code validation [can be found here](./#configurations-with-remote-code-validation)."
	CodeArmRequired        *bool       `json:"code_arm_required,omitempty"`     // "If true the code is required to arm the alarm. If false the code is not validated."
	CodeDisarmRequired     *bool       `json:"code_disarm_required,omitempty"`  // "If true the code is required to disarm the alarm. If false the code is not validated."
	CodeTriggerRequired    *bool       `json:"code_trigger_required,omitempty"` // "If true the code is required to trigger the alarm. If false the code is not validated."
	CommandTemplate        *string     `json:"command_template,omitempty"`      // "The [template](/docs/configuration/templating/#processing-incoming-data) used for the command payload. Available variables: `action` and `code`."
	Command                *([]string) `json:"command,omitempty"`
	EnabledByDefault       *bool       `json:"enabled_by_default,omitempty"`        // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string     `json:"encoding,omitempty"`                  // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string     `json:"entity_category,omitempty"`           // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   *string     `json:"icon,omitempty"`                      // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Name                   *string     `json:"name,omitempty"`                      // "The name of the alarm."
	ObjectId               *string     `json:"object_id,omitempty"`                 // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadArmAway         *string     `json:"payload_arm_away,omitempty"`          // "The payload to set armed-away mode on your Alarm Panel."
	PayloadArmCustomBypass *string     `json:"payload_arm_custom_bypass,omitempty"` // "The payload to set armed-custom-bypass mode on your Alarm Panel."
	PayloadArmHome         *string     `json:"payload_arm_home,omitempty"`          // "The payload to set armed-home mode on your Alarm Panel."
	PayloadArmNight        *string     `json:"payload_arm_night,omitempty"`         // "The payload to set armed-night mode on your Alarm Panel."
	PayloadArmVacation     *string     `json:"payload_arm_vacation,omitempty"`      // "The payload to set armed-vacation mode on your Alarm Panel."
	PayloadAvailable       *string     `json:"payload_available,omitempty"`         // "The payload that represents the available state."
	PayloadDisarm          *string     `json:"payload_disarm,omitempty"`            // "The payload to disarm your Alarm Panel."
	PayloadNotAvailable    *string     `json:"payload_not_available,omitempty"`     // "The payload that represents the unavailable state."
	PayloadTrigger         *string     `json:"payload_trigger,omitempty"`           // "The payload to trigger the alarm on your Alarm Panel."
	Qos                    *int        `json:"qos,omitempty"`                       // "The maximum QoS level of the state topic."
	Retain                 *bool       `json:"retain,omitempty"`                    // "If the published message should have the retain flag on or not."
	State                  *([]string) `json:"state,omitempty"`
	UniqueId               *string     `json:"unique_id,omitempty"`      // "An ID that uniquely identifies this alarm panel. If two alarm panels have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          *string     `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the value."
	MQTT                   struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}

func (iDevice AlarmControlPanel) Translate() externaldevice.AlarmControlPanel {
	eDevice := externaldevice.AlarmControlPanel{}
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
	if iDevice.Code != nil {
		eDevice.Code = iDevice.Code
	}
	if iDevice.CodeArmRequired != nil {
		eDevice.CodeArmRequired = iDevice.CodeArmRequired
	}
	if iDevice.CodeDisarmRequired != nil {
		eDevice.CodeDisarmRequired = iDevice.CodeDisarmRequired
	}
	if iDevice.CodeTriggerRequired != nil {
		eDevice.CodeTriggerRequired = iDevice.CodeTriggerRequired
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
	if iDevice.Name != nil {
		eDevice.Name = iDevice.Name
	}
	if iDevice.ObjectId != nil {
		eDevice.ObjectId = iDevice.ObjectId
	}
	if iDevice.PayloadArmAway != nil {
		eDevice.PayloadArmAway = iDevice.PayloadArmAway
	}
	if iDevice.PayloadArmCustomBypass != nil {
		eDevice.PayloadArmCustomBypass = iDevice.PayloadArmCustomBypass
	}
	if iDevice.PayloadArmHome != nil {
		eDevice.PayloadArmHome = iDevice.PayloadArmHome
	}
	if iDevice.PayloadArmNight != nil {
		eDevice.PayloadArmNight = iDevice.PayloadArmNight
	}
	if iDevice.PayloadArmVacation != nil {
		eDevice.PayloadArmVacation = iDevice.PayloadArmVacation
	}
	if iDevice.PayloadAvailable != nil {
		eDevice.PayloadAvailable = iDevice.PayloadAvailable
	}
	if iDevice.PayloadDisarm != nil {
		eDevice.PayloadDisarm = iDevice.PayloadDisarm
	}
	if iDevice.PayloadNotAvailable != nil {
		eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	}
	if iDevice.PayloadTrigger != nil {
		eDevice.PayloadTrigger = iDevice.PayloadTrigger
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
