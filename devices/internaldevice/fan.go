package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Fan struct {
	AvailabilityMode           *string     `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate       *string     `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability               *([]string) `json:"availability,omitempty"`
	CommandTemplate            *string     `json:"command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `command_topic`."
	Command                    *([]string) `json:"command,omitempty"`
	EnabledByDefault           *bool       `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding                   *string     `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory             *string     `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                       *string     `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate     *string     `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributes             *([]string) `json:"json_attributes,omitempty"`
	Name                       *string     `json:"name,omitempty"`                         // "The name of the fan."
	ObjectId                   *string     `json:"object_id,omitempty"`                    // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic                 *bool       `json:"optimistic,omitempty"`                   // "Flag that defines if fan works in optimistic mode"
	OscillationCommandTemplate *string     `json:"oscillation_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `oscillation_command_topic`."
	OscillationCommand         *([]string) `json:"oscillation_command,omitempty"`
	OscillationState           *([]string) `json:"oscillation_state,omitempty"`
	OscillationValueTemplate   *string     `json:"oscillation_value_template,omitempty"`  // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract a value from the oscillation."
	PayloadAvailable           *string     `json:"payload_available,omitempty"`           // "The payload that represents the available state."
	PayloadNotAvailable        *string     `json:"payload_not_available,omitempty"`       // "The payload that represents the unavailable state."
	PayloadOff                 *string     `json:"payload_off,omitempty"`                 // "The payload that represents the stop state."
	PayloadOn                  *string     `json:"payload_on,omitempty"`                  // "The payload that represents the running state."
	PayloadOscillationOff      *string     `json:"payload_oscillation_off,omitempty"`     // "The payload that represents the oscillation off state."
	PayloadOscillationOn       *string     `json:"payload_oscillation_on,omitempty"`      // "The payload that represents the oscillation on state."
	PayloadResetPercentage     *string     `json:"payload_reset_percentage,omitempty"`    // "A special payload that resets the `percentage` state attribute to `None` when received at the `percentage_state_topic`."
	PayloadResetPresetMode     *string     `json:"payload_reset_preset_mode,omitempty"`   // "A special payload that resets the `preset_mode` state attribute to `None` when received at the `preset_mode_state_topic`."
	PercentageCommandTemplate  *string     `json:"percentage_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `percentage_command_topic`."
	PercentageCommand          *([]string) `json:"percentage_command,omitempty"`
	PercentageState            *([]string) `json:"percentage_state,omitempty"`
	PercentageValueTemplate    *string     `json:"percentage_value_template,omitempty"`    // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the `percentage` value from the payload received on `percentage_state_topic`."
	PresetModeCommandTemplate  *string     `json:"preset_mode_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `preset_mode_command_topic`."
	PresetModeCommand          *([]string) `json:"preset_mode_command,omitempty"`
	PresetModeState            *([]string) `json:"preset_mode_state,omitempty"`
	PresetModeValueTemplate    *string     `json:"preset_mode_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the `preset_mode` value from the payload received on `preset_mode_state_topic`."
	PresetModes                *([]string) `json:"preset_modes,omitempty"`               // "List of preset modes this fan is capable of running at. Common examples include `auto`, `smart`, `whoosh`, `eco` and `breeze`."
	Qos                        *int        `json:"qos,omitempty"`                        // "The maximum QoS level of the state topic."
	Retain                     *bool       `json:"retain,omitempty"`                     // "If the published message should have the retain flag on or not."
	SpeedRangeMax              *int        `json:"speed_range_max,omitempty"`            // "The maximum of numeric output range (representing 100 %). The number of speeds within the `speed_range` / `100` will determine the `percentage_step`."
	SpeedRangeMin              *int        `json:"speed_range_min,omitempty"`            // "The minimum of numeric output range (`off` not included, so `speed_range_min` - `1` represents 0 %). The number of speeds within the speed_range / 100 will determine the `percentage_step`."
	State                      *([]string) `json:"state,omitempty"`
	StateValueTemplate         *string     `json:"state_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract a value from the state."
	UniqueId                   *string     `json:"unique_id,omitempty"`            // "An ID that uniquely identifies this fan. If two fans have the same unique ID, Home Assistant will raise an exception."
	MQTT                       struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}

func (iDevice Fan) Translate() externaldevice.Fan {
	eDevice := externaldevice.Fan{}
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
	if iDevice.OscillationCommandTemplate != nil {
		eDevice.OscillationCommandTemplate = iDevice.OscillationCommandTemplate
	}
	if iDevice.OscillationCommand != nil {
		eDevice.OscillationCommandFunc = common.ConstructCommandFunc(*iDevice.OscillationCommand)
	}
	if iDevice.OscillationState != nil {
		eDevice.OscillationStateFunc = common.ConstructStateFunc(*iDevice.OscillationState)
	}
	if iDevice.OscillationValueTemplate != nil {
		eDevice.OscillationValueTemplate = iDevice.OscillationValueTemplate
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
	if iDevice.PayloadOscillationOff != nil {
		eDevice.PayloadOscillationOff = iDevice.PayloadOscillationOff
	}
	if iDevice.PayloadOscillationOn != nil {
		eDevice.PayloadOscillationOn = iDevice.PayloadOscillationOn
	}
	if iDevice.PayloadResetPercentage != nil {
		eDevice.PayloadResetPercentage = iDevice.PayloadResetPercentage
	}
	if iDevice.PayloadResetPresetMode != nil {
		eDevice.PayloadResetPresetMode = iDevice.PayloadResetPresetMode
	}
	if iDevice.PercentageCommandTemplate != nil {
		eDevice.PercentageCommandTemplate = iDevice.PercentageCommandTemplate
	}
	if iDevice.PercentageCommand != nil {
		eDevice.PercentageCommandFunc = common.ConstructCommandFunc(*iDevice.PercentageCommand)
	}
	if iDevice.PercentageState != nil {
		eDevice.PercentageStateFunc = common.ConstructStateFunc(*iDevice.PercentageState)
	}
	if iDevice.PercentageValueTemplate != nil {
		eDevice.PercentageValueTemplate = iDevice.PercentageValueTemplate
	}
	if iDevice.PresetModeCommandTemplate != nil {
		eDevice.PresetModeCommandTemplate = iDevice.PresetModeCommandTemplate
	}
	if iDevice.PresetModeCommand != nil {
		eDevice.PresetModeCommandFunc = common.ConstructCommandFunc(*iDevice.PresetModeCommand)
	}
	if iDevice.PresetModeState != nil {
		eDevice.PresetModeStateFunc = common.ConstructStateFunc(*iDevice.PresetModeState)
	}
	if iDevice.PresetModeValueTemplate != nil {
		eDevice.PresetModeValueTemplate = iDevice.PresetModeValueTemplate
	}
	if iDevice.PresetModes != nil {
		eDevice.PresetModes = iDevice.PresetModes
	}
	if iDevice.Qos != nil {
		eDevice.Qos = iDevice.Qos
	}
	if iDevice.Retain != nil {
		eDevice.Retain = iDevice.Retain
	}
	if iDevice.SpeedRangeMax != nil {
		eDevice.SpeedRangeMax = iDevice.SpeedRangeMax
	}
	if iDevice.SpeedRangeMin != nil {
		eDevice.SpeedRangeMin = iDevice.SpeedRangeMin
	}
	if iDevice.State != nil {
		eDevice.StateFunc = common.ConstructStateFunc(*iDevice.State)
	}
	if iDevice.StateValueTemplate != nil {
		eDevice.StateValueTemplate = iDevice.StateValueTemplate
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
