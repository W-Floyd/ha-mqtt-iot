package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
func (iDevice Humidifier) Translate() externaldevice.Humidifier {
	eDevice := externaldevice.Humidifier{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.AvailabilityMode = iDevice.AvailabilityMode
	eDevice.AvailabilityTemplate = iDevice.AvailabilityTemplate
	eDevice.AvailabilityFunc = common.ConstructStateFunc(iDevice.Availability)
	eDevice.CommandTemplate = iDevice.CommandTemplate
	eDevice.CommandFunc = common.ConstructCommandFunc(iDevice.Command)
	eDevice.DeviceClass = iDevice.DeviceClass
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.Icon = iDevice.Icon
	eDevice.MaxHumidity = iDevice.MaxHumidity
	eDevice.MinHumidity = iDevice.MinHumidity
	eDevice.ModeCommandTemplate = iDevice.ModeCommandTemplate
	eDevice.ModeCommandFunc = common.ConstructCommandFunc(iDevice.ModeCommand)
	eDevice.ModeStateTemplate = iDevice.ModeStateTemplate
	eDevice.ModeStateFunc = common.ConstructStateFunc(iDevice.ModeState)
	eDevice.Modes = iDevice.Modes
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.Optimistic = iDevice.Optimistic
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.PayloadOff = iDevice.PayloadOff
	eDevice.PayloadOn = iDevice.PayloadOn
	eDevice.PayloadResetHumidity = iDevice.PayloadResetHumidity
	eDevice.PayloadResetMode = iDevice.PayloadResetMode
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.StateValueTemplate = iDevice.StateValueTemplate
	eDevice.TargetHumidityCommandTemplate = iDevice.TargetHumidityCommandTemplate
	eDevice.TargetHumidityCommandFunc = common.ConstructCommandFunc(iDevice.TargetHumidityCommand)
	eDevice.TargetHumidityStateTemplate = iDevice.TargetHumidityStateTemplate
	eDevice.TargetHumidityStateFunc = common.ConstructStateFunc(iDevice.TargetHumidityState)
	eDevice.UniqueId = iDevice.UniqueId
	if len(iDevice.Availability) == 0 {
		eDevice.AvailabilityFunc = common.AvailabilityFunc
	}
	eDevice.Initialize()
	return eDevice
}

type Humidifier struct {
	AvailabilityMode              string   `json:"availability_mode"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate          string   `json:"availability_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability                  []string `json:"availability"`
	CommandTemplate               string   `json:"command_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `command_topic`."
	Command                       []string `json:"command"`
	DeviceClass                   string   `json:"device_class"`          // "The device class of the MQTT device. Must be either `humidifier` or `dehumidifier`."
	EnabledByDefault              bool     `json:"enabled_by_default"`    // "Flag which defines if the entity should be enabled when first added."
	Encoding                      string   `json:"encoding"`              // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory                string   `json:"entity_category"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                          string   `json:"icon"`                  // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	MaxHumidity                   int      `json:"max_humidity"`          // "The minimum target humidity percentage that can be set."
	MinHumidity                   int      `json:"min_humidity"`          // "The maximum target humidity percentage that can be set."
	ModeCommandTemplate           string   `json:"mode_command_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `mode_command_topic`."
	ModeCommand                   []string `json:"mode_command"`
	ModeStateTemplate             string   `json:"mode_state_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract a value for the humidifier `mode` state."
	ModeState                     []string `json:"mode_state"`
	Modes                         []string `json:"modes"`                  // "List of available modes this humidifier is capable of running at. Common examples include `normal`, `eco`, `away`, `boost`, `comfort`, `home`, `sleep`, `auto` and `baby`. These examples offer built-in translations but other custom modes are allowed as well.  This attribute ust be configured together with the `mode_command_topic` attribute."
	Name                          string   `json:"name"`                   // "The name of the humidifier."
	ObjectId                      string   `json:"object_id"`              // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic                    bool     `json:"optimistic"`             // "Flag that defines if humidifier works in optimistic mode"
	PayloadAvailable              string   `json:"payload_available"`      // "The payload that represents the available state."
	PayloadNotAvailable           string   `json:"payload_not_available"`  // "The payload that represents the unavailable state."
	PayloadOff                    string   `json:"payload_off"`            // "The payload that represents the stop state."
	PayloadOn                     string   `json:"payload_on"`             // "The payload that represents the running state."
	PayloadResetHumidity          string   `json:"payload_reset_humidity"` // "A special payload that resets the `target_humidity` state attribute to `None` when received at the `target_humidity_state_topic`."
	PayloadResetMode              string   `json:"payload_reset_mode"`     // "A special payload that resets the `mode` state attribute to `None` when received at the `mode_state_topic`."
	Qos                           int      `json:"qos"`                    // "The maximum QoS level of the state topic."
	Retain                        bool     `json:"retain"`                 // "If the published message should have the retain flag on or not."
	State                         []string `json:"state"`
	StateValueTemplate            string   `json:"state_value_template"`             // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract a value from the state."
	TargetHumidityCommandTemplate string   `json:"target_humidity_command_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `target_humidity_command_topic`."
	TargetHumidityCommand         []string `json:"target_humidity_command"`
	TargetHumidityStateTemplate   string   `json:"target_humidity_state_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract a value for the humidifier `target_humidity` state."
	TargetHumidityState           []string `json:"target_humidity_state"`
	UniqueId                      string   `json:"unique_id"` // "An ID that uniquely identifies this humidifier. If two humidifiers have the same unique ID, Home Assistant will raise an exception."
	MQTT                          struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
