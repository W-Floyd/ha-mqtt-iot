package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
func (iDevice Fan) Translate() externaldevice.Fan {
	eDevice := externaldevice.Fan{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.CommandTemplate = iDevice.CommandTemplate
	eDevice.CommandFunc = common.ConstructCommandFunc(iDevice.Command)
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.Icon = iDevice.Icon
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.Optimistic = iDevice.Optimistic
	eDevice.OscillationCommandTemplate = iDevice.OscillationCommandTemplate
	eDevice.OscillationCommandFunc = common.ConstructCommandFunc(iDevice.OscillationCommand)
	eDevice.OscillationStateFunc = common.ConstructStateFunc(iDevice.OscillationState)
	eDevice.OscillationValueTemplate = iDevice.OscillationValueTemplate
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.PayloadOff = iDevice.PayloadOff
	eDevice.PayloadOn = iDevice.PayloadOn
	eDevice.PayloadOscillationOff = iDevice.PayloadOscillationOff
	eDevice.PayloadOscillationOn = iDevice.PayloadOscillationOn
	eDevice.PayloadResetPercentage = iDevice.PayloadResetPercentage
	eDevice.PayloadResetPresetMode = iDevice.PayloadResetPresetMode
	eDevice.PercentageCommandTemplate = iDevice.PercentageCommandTemplate
	eDevice.PercentageCommandFunc = common.ConstructCommandFunc(iDevice.PercentageCommand)
	eDevice.PercentageStateFunc = common.ConstructStateFunc(iDevice.PercentageState)
	eDevice.PercentageValueTemplate = iDevice.PercentageValueTemplate
	eDevice.PresetModeCommandTemplate = iDevice.PresetModeCommandTemplate
	eDevice.PresetModeCommandFunc = common.ConstructCommandFunc(iDevice.PresetModeCommand)
	eDevice.PresetModeStateFunc = common.ConstructStateFunc(iDevice.PresetModeState)
	eDevice.PresetModeValueTemplate = iDevice.PresetModeValueTemplate
	eDevice.PresetModes = iDevice.PresetModes
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.SpeedRangeMax = iDevice.SpeedRangeMax
	eDevice.SpeedRangeMin = iDevice.SpeedRangeMin
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.StateValueTemplate = iDevice.StateValueTemplate
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.Initialize()
	return eDevice
}

type Fan struct {
	CommandTemplate            string   `json:"command_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `command_topic`."
	Command                    []string `json:"command"`
	EnabledByDefault           bool     `json:"enabled_by_default"`           // "Flag which defines if the entity should be enabled when first added."
	Encoding                   string   `json:"encoding"`                     // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory             string   `json:"entity_category"`              // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                       string   `json:"icon"`                         // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Name                       string   `json:"name"`                         // "The name of the fan."
	ObjectId                   string   `json:"object_id"`                    // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic                 bool     `json:"optimistic"`                   // "Flag that defines if fan works in optimistic mode"
	OscillationCommandTemplate string   `json:"oscillation_command_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `oscillation_command_topic`."
	OscillationCommand         []string `json:"oscillation_command"`
	OscillationState           []string `json:"oscillation_state"`
	OscillationValueTemplate   string   `json:"oscillation_value_template"`  // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract a value from the oscillation."
	PayloadAvailable           string   `json:"payload_available"`           // "The payload that represents the available state."
	PayloadNotAvailable        string   `json:"payload_not_available"`       // "The payload that represents the unavailable state."
	PayloadOff                 string   `json:"payload_off"`                 // "The payload that represents the stop state."
	PayloadOn                  string   `json:"payload_on"`                  // "The payload that represents the running state."
	PayloadOscillationOff      string   `json:"payload_oscillation_off"`     // "The payload that represents the oscillation off state."
	PayloadOscillationOn       string   `json:"payload_oscillation_on"`      // "The payload that represents the oscillation on state."
	PayloadResetPercentage     string   `json:"payload_reset_percentage"`    // "A special payload that resets the `percentage` state attribute to `None` when received at the `percentage_state_topic`."
	PayloadResetPresetMode     string   `json:"payload_reset_preset_mode"`   // "A special payload that resets the `preset_mode` state attribute to `None` when received at the `preset_mode_state_topic`."
	PercentageCommandTemplate  string   `json:"percentage_command_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `percentage_command_topic`."
	PercentageCommand          []string `json:"percentage_command"`
	PercentageState            []string `json:"percentage_state"`
	PercentageValueTemplate    string   `json:"percentage_value_template"`    // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the `percentage` value from the payload received on `percentage_state_topic`."
	PresetModeCommandTemplate  string   `json:"preset_mode_command_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `preset_mode_command_topic`."
	PresetModeCommand          []string `json:"preset_mode_command"`
	PresetModeState            []string `json:"preset_mode_state"`
	PresetModeValueTemplate    string   `json:"preset_mode_value_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the `preset_mode` value from the payload received on `preset_mode_state_topic`."
	PresetModes                []string `json:"preset_modes"`               // "List of preset modes this fan is capable of running at. Common examples include `auto`, `smart`, `whoosh`, `eco` and `breeze`."
	Qos                        int      `json:"qos"`                        // "The maximum QoS level of the state topic."
	Retain                     bool     `json:"retain"`                     // "If the published message should have the retain flag on or not."
	SpeedRangeMax              int      `json:"speed_range_max"`            // "The maximum of numeric output range (representing 100 %). The number of speeds within the `speed_range` / `100` will determine the `percentage_step`."
	SpeedRangeMin              int      `json:"speed_range_min"`            // "The minimum of numeric output range (`off` not included, so `speed_range_min` - `1` represents 0 %). The number of speeds within the speed_range / 100 will determine the `percentage_step`."
	State                      []string `json:"state"`
	StateValueTemplate         string   `json:"state_value_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract a value from the state."
	UniqueId                   string   `json:"unique_id"`            // "An ID that uniquely identifies this fan. If two fans have the same unique ID, Home Assistant will raise an exception."
	MQTT                       struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
