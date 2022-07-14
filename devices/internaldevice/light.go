package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
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
	if iDevice.Brightness != nil {
		eDevice.Brightness = iDevice.Brightness
	}
	if iDevice.BrightnessScale != nil {
		eDevice.BrightnessScale = iDevice.BrightnessScale
	}
	if iDevice.ColorMode != nil {
		eDevice.ColorMode = iDevice.ColorMode
	}
	if iDevice.Command != nil {
		eDevice.CommandFunc = common.ConstructCommandFunc(*iDevice.Command)
	}
	if iDevice.Effect != nil {
		eDevice.Effect = iDevice.Effect
	}
	if iDevice.EffectList != nil {
		eDevice.EffectList = iDevice.EffectList
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
	if iDevice.FlashTimeLong != nil {
		eDevice.FlashTimeLong = iDevice.FlashTimeLong
	}
	if iDevice.FlashTimeShort != nil {
		eDevice.FlashTimeShort = iDevice.FlashTimeShort
	}
	if iDevice.Icon != nil {
		eDevice.Icon = iDevice.Icon
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
	if iDevice.Optimistic != nil {
		eDevice.Optimistic = iDevice.Optimistic
	}
	if iDevice.PayloadAvailable != nil {
		eDevice.PayloadAvailable = iDevice.PayloadAvailable
	}
	if iDevice.PayloadNotAvailable != nil {
		eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
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
	if iDevice.State != nil {
		eDevice.StateFunc = common.ConstructStateFunc(*iDevice.State)
	}
	if iDevice.SupportedColorModes != nil {
		eDevice.SupportedColorModes = iDevice.SupportedColorModes
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

type Light struct {
	AvailabilityMode     *string     `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate *string     `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability         *([]string) `json:"availability,omitempty"`
	Brightness           *bool       `json:"brightness,omitempty"`       // "Flag that defines if the light supports brightness."
	BrightnessScale      *int        `json:"brightness_scale,omitempty"` // "Defines the maximum brightness value (i.e., 100%) of the MQTT device."
	ColorMode            *bool       `json:"color_mode,omitempty"`       // "Flag that defines if the light supports color modes."
	Command              *([]string) `json:"command,omitempty"`
	Effect               *bool       `json:"effect,omitempty"`                // "Flag that defines if the light supports effects."
	EffectList           *([]string) `json:"effect_list,omitempty"`           // "The list of effects the light supports."
	EnabledByDefault     *bool       `json:"enabled_by_default,omitempty"`    // "Flag which defines if the entity should be enabled when first added."
	Encoding             *string     `json:"encoding,omitempty"`              // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory       *string     `json:"entity_category,omitempty"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	FlashTimeLong        *int        `json:"flash_time_long,omitempty"`       // "The duration, in seconds, of a “long” flash."
	FlashTimeShort       *int        `json:"flash_time_short,omitempty"`      // "The duration, in seconds, of a “short” flash."
	Icon                 *string     `json:"icon,omitempty"`                  // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	MaxMireds            *int        `json:"max_mireds,omitempty"`            // "The maximum color temperature in mireds."
	MinMireds            *int        `json:"min_mireds,omitempty"`            // "The minimum color temperature in mireds."
	Name                 *string     `json:"name,omitempty"`                  // "The name of the light."
	ObjectId             *string     `json:"object_id,omitempty"`             // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic           *bool       `json:"optimistic,omitempty"`            // "Flag that defines if the light works in optimistic mode."
	PayloadAvailable     *string     `json:"payload_available,omitempty"`     // "The payload that represents the available state."
	PayloadNotAvailable  *string     `json:"payload_not_available,omitempty"` // "The payload that represents the unavailable state."
	Qos                  *int        `json:"qos,omitempty"`                   // "The maximum QoS level of the state topic."
	Retain               *bool       `json:"retain,omitempty"`                // "If the published message should have the retain flag on or not."
	Schema               *string     `json:"schema,omitempty"`                // "The schema to use. Must be `json` to select the JSON schema."
	State                *([]string) `json:"state,omitempty"`
	SupportedColorModes  *([]string) `json:"supported_color_modes,omitempty"` // "A list of color modes supported by the list. This is required if `color_mode` is `True`. Possible color modes are `onoff`, `brightness`, `color_temp`, `hs`, `xy`, `rgb`, `rgbw`, `rgbww`."
	UniqueId             *string     `json:"unique_id,omitempty"`             // "An ID that uniquely identifies this light. If two lights have the same unique ID, Home Assistant will raise an exception."
	MQTT                 struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}
