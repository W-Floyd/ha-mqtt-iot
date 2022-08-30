package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Light struct {
	AvailabilityMode          *string     `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate      *string     `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability              *([]string) `json:"availability,omitempty"`
	BrightnessCommandTemplate *string     `json:"brightness_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `brightness_command_topic`. Available variables: `value`."
	BrightnessCommand         *([]string) `json:"brightness_command,omitempty"`
	BrightnessScale           *int        `json:"brightness_scale,omitempty"` // "Defines the maximum brightness value (i.e., 100%) of the MQTT device."
	BrightnessState           *([]string) `json:"brightness_state,omitempty"`
	BrightnessValueTemplate   *string     `json:"brightness_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the brightness value."
	ColorModeState            *([]string) `json:"color_mode_state,omitempty"`
	ColorModeValueTemplate    *string     `json:"color_mode_value_template,omitempty"`   // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the color mode."
	ColorTempCommandTemplate  *string     `json:"color_temp_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `color_temp_command_topic`. Available variables: `value`."
	ColorTempCommand          *([]string) `json:"color_temp_command,omitempty"`
	ColorTempState            *([]string) `json:"color_temp_state,omitempty"`
	ColorTempValueTemplate    *string     `json:"color_temp_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the color temperature value."
	Command                   *([]string) `json:"command,omitempty"`
	EffectCommandTemplate     *string     `json:"effect_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `effect_command_topic`. Available variables: `value`."
	EffectCommand             *([]string) `json:"effect_command,omitempty"`
	EffectList                *([]string) `json:"effect_list,omitempty"` // "The list of effects the light supports."
	EffectState               *([]string) `json:"effect_state,omitempty"`
	EffectValueTemplate       *string     `json:"effect_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the effect value."
	EnabledByDefault          *bool       `json:"enabled_by_default,omitempty"`    // "Flag which defines if the entity should be enabled when first added."
	Encoding                  *string     `json:"encoding,omitempty"`              // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory            *string     `json:"entity_category,omitempty"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	HsCommand                 *([]string) `json:"hs_command,omitempty"`
	HsState                   *([]string) `json:"hs_state,omitempty"`
	HsValueTemplate           *string     `json:"hs_value_template,omitempty"`        // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the HS value."
	Icon                      *string     `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate    *string     `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributes            *([]string) `json:"json_attributes,omitempty"`
	MaxMireds                 *int        `json:"max_mireds,omitempty"`            // "The maximum color temperature in mireds."
	MinMireds                 *int        `json:"min_mireds,omitempty"`            // "The minimum color temperature in mireds."
	Name                      *string     `json:"name,omitempty"`                  // "The name of the light."
	ObjectId                  *string     `json:"object_id,omitempty"`             // "Used instead of `name` for automatic generation of `entity_id`"
	OnCommandType             *string     `json:"on_command_type,omitempty"`       // "Defines when on the payload_on is sent. Using `last` (the default) will send any style (brightness, color, etc) topics first and then a `payload_on` to the `command_topic`. Using `first` will send the `payload_on` and then any style topics. Using `brightness` will only send brightness commands instead of the `payload_on` to turn the light on."
	Optimistic                *bool       `json:"optimistic,omitempty"`            // "Flag that defines if switch works in optimistic mode."
	PayloadAvailable          *string     `json:"payload_available,omitempty"`     // "The payload that represents the available state."
	PayloadNotAvailable       *string     `json:"payload_not_available,omitempty"` // "The payload that represents the unavailable state."
	PayloadOff                *string     `json:"payload_off,omitempty"`           // "The payload that represents disabled state."
	PayloadOn                 *string     `json:"payload_on,omitempty"`            // "The payload that represents enabled state."
	Qos                       *int        `json:"qos,omitempty"`                   // "The maximum QoS level of the state topic."
	Retain                    *bool       `json:"retain,omitempty"`                // "If the published message should have the retain flag on or not."
	RgbCommandTemplate        *string     `json:"rgb_command_template,omitempty"`  // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `rgb_command_topic`. Available variables: `red`, `green` and `blue`."
	RgbCommand                *([]string) `json:"rgb_command,omitempty"`
	RgbState                  *([]string) `json:"rgb_state,omitempty"`
	RgbValueTemplate          *string     `json:"rgb_value_template,omitempty"`    // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the RGB value."
	RgbwCommandTemplate       *string     `json:"rgbw_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `rgbw_command_topic`. Available variables: `red`, `green`, `blue` and `white`."
	RgbwCommand               *([]string) `json:"rgbw_command,omitempty"`
	RgbwState                 *([]string) `json:"rgbw_state,omitempty"`
	RgbwValueTemplate         *string     `json:"rgbw_value_template,omitempty"`    // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the RGBW value."
	RgbwwCommandTemplate      *string     `json:"rgbww_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `rgbww_command_topic`. Available variables: `red`, `green`, `blue`, `cold_white` and `warm_white`."
	RgbwwCommand              *([]string) `json:"rgbww_command,omitempty"`
	RgbwwState                *([]string) `json:"rgbww_state,omitempty"`
	RgbwwValueTemplate        *string     `json:"rgbww_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the RGBWW value."
	Schema                    *string     `json:"schema,omitempty"`               // "The schema to use. Must be `default` or omitted to select the default schema."
	State                     *([]string) `json:"state,omitempty"`
	StateValueTemplate        *string     `json:"state_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the state value. The template should match the payload `on` and `off` values, so if your light uses `power on` to turn on, your `state_value_template` string should return `power on` when the switch is on. For example if the message is just `on`, your `state_value_template` should be `power {{ value }}`."
	UniqueId                  *string     `json:"unique_id,omitempty"`            // "An ID that uniquely identifies this light. If two lights have the same unique ID, Home Assistant will raise an exception."
	WhiteCommand              *([]string) `json:"white_command,omitempty"`
	WhiteScale                *int        `json:"white_scale,omitempty"` // "Defines the maximum white level (i.e., 100%) of the MQTT device."
	XyCommand                 *([]string) `json:"xy_command,omitempty"`
	XyState                   *([]string) `json:"xy_state,omitempty"`
	XyValueTemplate           *string     `json:"xy_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the XY value."
	MQTT                      struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}

func (iDevice Light) Translate() externaldevice.Light {
	eDevice := externaldevice.Light{}
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
	if iDevice.BrightnessCommandTemplate != nil {
		eDevice.BrightnessCommandTemplate = iDevice.BrightnessCommandTemplate
	}
	if iDevice.BrightnessCommand != nil {
		eDevice.BrightnessCommandFunc = common.ConstructCommandFunc(*iDevice.BrightnessCommand)
	}
	if iDevice.BrightnessScale != nil {
		eDevice.BrightnessScale = iDevice.BrightnessScale
	}
	if iDevice.BrightnessState != nil {
		eDevice.BrightnessStateFunc = common.ConstructStateFunc(*iDevice.BrightnessState)
	}
	if iDevice.BrightnessValueTemplate != nil {
		eDevice.BrightnessValueTemplate = iDevice.BrightnessValueTemplate
	}
	if iDevice.ColorModeState != nil {
		eDevice.ColorModeStateFunc = common.ConstructStateFunc(*iDevice.ColorModeState)
	}
	if iDevice.ColorModeValueTemplate != nil {
		eDevice.ColorModeValueTemplate = iDevice.ColorModeValueTemplate
	}
	if iDevice.ColorTempCommandTemplate != nil {
		eDevice.ColorTempCommandTemplate = iDevice.ColorTempCommandTemplate
	}
	if iDevice.ColorTempCommand != nil {
		eDevice.ColorTempCommandFunc = common.ConstructCommandFunc(*iDevice.ColorTempCommand)
	}
	if iDevice.ColorTempState != nil {
		eDevice.ColorTempStateFunc = common.ConstructStateFunc(*iDevice.ColorTempState)
	}
	if iDevice.ColorTempValueTemplate != nil {
		eDevice.ColorTempValueTemplate = iDevice.ColorTempValueTemplate
	}
	if iDevice.Command != nil {
		eDevice.CommandFunc = common.ConstructCommandFunc(*iDevice.Command)
	}
	if iDevice.EffectCommandTemplate != nil {
		eDevice.EffectCommandTemplate = iDevice.EffectCommandTemplate
	}
	if iDevice.EffectCommand != nil {
		eDevice.EffectCommandFunc = common.ConstructCommandFunc(*iDevice.EffectCommand)
	}
	if iDevice.EffectList != nil {
		eDevice.EffectList = iDevice.EffectList
	}
	if iDevice.EffectState != nil {
		eDevice.EffectStateFunc = common.ConstructStateFunc(*iDevice.EffectState)
	}
	if iDevice.EffectValueTemplate != nil {
		eDevice.EffectValueTemplate = iDevice.EffectValueTemplate
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
	if iDevice.HsCommand != nil {
		eDevice.HsCommandFunc = common.ConstructCommandFunc(*iDevice.HsCommand)
	}
	if iDevice.HsState != nil {
		eDevice.HsStateFunc = common.ConstructStateFunc(*iDevice.HsState)
	}
	if iDevice.HsValueTemplate != nil {
		eDevice.HsValueTemplate = iDevice.HsValueTemplate
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
	if iDevice.MaxMireds != nil {
		eDevice.MaxMireds = iDevice.MaxMireds
	}
	if iDevice.MinMireds != nil {
		eDevice.MinMireds = iDevice.MinMireds
	}
	if iDevice.Name != nil {
		eDevice.Name = iDevice.Name
	}
	if iDevice.ObjectId != nil {
		eDevice.ObjectId = iDevice.ObjectId
	}
	if iDevice.OnCommandType != nil {
		eDevice.OnCommandType = iDevice.OnCommandType
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
	if iDevice.Qos != nil {
		eDevice.Qos = iDevice.Qos
	}
	if iDevice.Retain != nil {
		eDevice.Retain = iDevice.Retain
	}
	if iDevice.RgbCommandTemplate != nil {
		eDevice.RgbCommandTemplate = iDevice.RgbCommandTemplate
	}
	if iDevice.RgbCommand != nil {
		eDevice.RgbCommandFunc = common.ConstructCommandFunc(*iDevice.RgbCommand)
	}
	if iDevice.RgbState != nil {
		eDevice.RgbStateFunc = common.ConstructStateFunc(*iDevice.RgbState)
	}
	if iDevice.RgbValueTemplate != nil {
		eDevice.RgbValueTemplate = iDevice.RgbValueTemplate
	}
	if iDevice.RgbwCommandTemplate != nil {
		eDevice.RgbwCommandTemplate = iDevice.RgbwCommandTemplate
	}
	if iDevice.RgbwCommand != nil {
		eDevice.RgbwCommandFunc = common.ConstructCommandFunc(*iDevice.RgbwCommand)
	}
	if iDevice.RgbwState != nil {
		eDevice.RgbwStateFunc = common.ConstructStateFunc(*iDevice.RgbwState)
	}
	if iDevice.RgbwValueTemplate != nil {
		eDevice.RgbwValueTemplate = iDevice.RgbwValueTemplate
	}
	if iDevice.RgbwwCommandTemplate != nil {
		eDevice.RgbwwCommandTemplate = iDevice.RgbwwCommandTemplate
	}
	if iDevice.RgbwwCommand != nil {
		eDevice.RgbwwCommandFunc = common.ConstructCommandFunc(*iDevice.RgbwwCommand)
	}
	if iDevice.RgbwwState != nil {
		eDevice.RgbwwStateFunc = common.ConstructStateFunc(*iDevice.RgbwwState)
	}
	if iDevice.RgbwwValueTemplate != nil {
		eDevice.RgbwwValueTemplate = iDevice.RgbwwValueTemplate
	}
	if iDevice.Schema != nil {
		eDevice.Schema = iDevice.Schema
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
	if iDevice.WhiteCommand != nil {
		eDevice.WhiteCommandFunc = common.ConstructCommandFunc(*iDevice.WhiteCommand)
	}
	if iDevice.WhiteScale != nil {
		eDevice.WhiteScale = iDevice.WhiteScale
	}
	if iDevice.XyCommand != nil {
		eDevice.XyCommandFunc = common.ConstructCommandFunc(*iDevice.XyCommand)
	}
	if iDevice.XyState != nil {
		eDevice.XyStateFunc = common.ConstructStateFunc(*iDevice.XyState)
	}
	if iDevice.XyValueTemplate != nil {
		eDevice.XyValueTemplate = iDevice.XyValueTemplate
	}
	if iDevice.Availability == nil {
		eDevice.AvailabilityFunc = common.AvailabilityFunc
	}
	eDevice.Initialize()
	return eDevice
}
