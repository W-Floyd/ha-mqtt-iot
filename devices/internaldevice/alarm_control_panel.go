package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
func (iDevice AlarmControlPanel) Translate() externaldevice.AlarmControlPanel {
	eDevice := externaldevice.AlarmControlPanel{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.AvailabilityMode = iDevice.AvailabilityMode
	eDevice.AvailabilityTemplate = iDevice.AvailabilityTemplate
	eDevice.AvailabilityFunc = common.ConstructStateFunc(iDevice.Availability)
	eDevice.Code = iDevice.Code
	eDevice.CodeArmRequired = iDevice.CodeArmRequired
	eDevice.CodeDisarmRequired = iDevice.CodeDisarmRequired
	eDevice.CodeTriggerRequired = iDevice.CodeTriggerRequired
	eDevice.CommandTemplate = iDevice.CommandTemplate
	eDevice.CommandFunc = common.ConstructCommandFunc(iDevice.Command)
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.Icon = iDevice.Icon
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.PayloadArmAway = iDevice.PayloadArmAway
	eDevice.PayloadArmCustomBypass = iDevice.PayloadArmCustomBypass
	eDevice.PayloadArmHome = iDevice.PayloadArmHome
	eDevice.PayloadArmNight = iDevice.PayloadArmNight
	eDevice.PayloadArmVacation = iDevice.PayloadArmVacation
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadDisarm = iDevice.PayloadDisarm
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.PayloadTrigger = iDevice.PayloadTrigger
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.ValueTemplate = iDevice.ValueTemplate
	if len(iDevice.Availability) == 0 {
		eDevice.AvailabilityFunc = common.AvailabilityFunc
	}
	eDevice.Initialize()
	return eDevice
}

type AlarmControlPanel struct {
	AvailabilityMode       string   `json:"availability_mode"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   string   `json:"availability_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability           []string `json:"availability"`
	Code                   string   `json:"code"`                  // "If defined, specifies a code to enable or disable the alarm in the frontend. Note that the code is validated locally and blocks sending MQTT messages to the remote device. For remote code validation, the code can be configured to either of the special values `REMOTE_CODE` (numeric code) or `REMOTE_CODE_TEXT` (text code). In this case, local code validation is bypassed but the frontend will still show a numeric or text code dialog. Use `command_template` to send the code to the remote device. Example configurations for remote code validation [can be found here](./#configurations-with-remote-code-validation)."
	CodeArmRequired        bool     `json:"code_arm_required"`     // "If true the code is required to arm the alarm. If false the code is not validated."
	CodeDisarmRequired     bool     `json:"code_disarm_required"`  // "If true the code is required to disarm the alarm. If false the code is not validated."
	CodeTriggerRequired    bool     `json:"code_trigger_required"` // "If true the code is required to trigger the alarm. If false the code is not validated."
	CommandTemplate        string   `json:"command_template"`      // "The [template](/docs/configuration/templating/#processing-incoming-data) used for the command payload. Available variables: `action` and `code`."
	Command                []string `json:"command"`
	EnabledByDefault       bool     `json:"enabled_by_default"`        // "Flag which defines if the entity should be enabled when first added."
	Encoding               string   `json:"encoding"`                  // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         string   `json:"entity_category"`           // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   string   `json:"icon"`                      // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Name                   string   `json:"name"`                      // "The name of the alarm."
	ObjectId               string   `json:"object_id"`                 // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadArmAway         string   `json:"payload_arm_away"`          // "The payload to set armed-away mode on your Alarm Panel."
	PayloadArmCustomBypass string   `json:"payload_arm_custom_bypass"` // "The payload to set armed-custom-bypass mode on your Alarm Panel."
	PayloadArmHome         string   `json:"payload_arm_home"`          // "The payload to set armed-home mode on your Alarm Panel."
	PayloadArmNight        string   `json:"payload_arm_night"`         // "The payload to set armed-night mode on your Alarm Panel."
	PayloadArmVacation     string   `json:"payload_arm_vacation"`      // "The payload to set armed-vacation mode on your Alarm Panel."
	PayloadAvailable       string   `json:"payload_available"`         // "The payload that represents the available state."
	PayloadDisarm          string   `json:"payload_disarm"`            // "The payload to disarm your Alarm Panel."
	PayloadNotAvailable    string   `json:"payload_not_available"`     // "The payload that represents the unavailable state."
	PayloadTrigger         string   `json:"payload_trigger"`           // "The payload to trigger the alarm on your Alarm Panel."
	Qos                    int      `json:"qos"`                       // "The maximum QoS level of the state topic."
	Retain                 bool     `json:"retain"`                    // "If the published message should have the retain flag on or not."
	State                  []string `json:"state"`
	UniqueId               string   `json:"unique_id"`      // "An ID that uniquely identifies this alarm panel. If two alarm panels have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          string   `json:"value_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the value."
	MQTT                   struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
