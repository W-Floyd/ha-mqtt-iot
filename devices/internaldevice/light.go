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
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.AvailabilityMode = iDevice.AvailabilityMode
	eDevice.AvailabilityTemplate = iDevice.AvailabilityTemplate
	eDevice.AvailabilityFunc = common.ConstructStateFunc(iDevice.Availability)
	eDevice.Brightness = iDevice.Brightness
	eDevice.BrightnessScale = iDevice.BrightnessScale
	eDevice.ColorMode = iDevice.ColorMode
	eDevice.CommandFunc = common.ConstructCommandFunc(iDevice.Command)
	eDevice.Effect = iDevice.Effect
	eDevice.EffectList = iDevice.EffectList
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.FlashTimeLong = iDevice.FlashTimeLong
	eDevice.FlashTimeShort = iDevice.FlashTimeShort
	eDevice.Icon = iDevice.Icon
	eDevice.MaxMireds = iDevice.MaxMireds
	eDevice.MinMireds = iDevice.MinMireds
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.Optimistic = iDevice.Optimistic
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.Schema = iDevice.Schema
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.SupportedColorModes = iDevice.SupportedColorModes
	eDevice.UniqueId = iDevice.UniqueId
	if len(iDevice.Availability) == 0 {
		eDevice.AvailabilityFunc = common.AvailabilityFunc
	}
	eDevice.Initialize()
	return eDevice
}

type Light struct {
	AvailabilityMode     string   `json:"availability_mode"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate string   `json:"availability_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability         []string `json:"availability"`
	Brightness           bool     `json:"brightness"`       // "Flag that defines if the light supports brightness."
	BrightnessScale      int      `json:"brightness_scale"` // "Defines the maximum brightness value (i.e., 100%) of the MQTT device."
	ColorMode            bool     `json:"color_mode"`       // "Flag that defines if the light supports color modes."
	Command              []string `json:"command"`
	Effect               bool     `json:"effect"`                // "Flag that defines if the light supports effects."
	EffectList           []string `json:"effect_list"`           // "The list of effects the light supports."
	EnabledByDefault     bool     `json:"enabled_by_default"`    // "Flag which defines if the entity should be enabled when first added."
	Encoding             string   `json:"encoding"`              // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory       string   `json:"entity_category"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	FlashTimeLong        int      `json:"flash_time_long"`       // "The duration, in seconds, of a “long” flash."
	FlashTimeShort       int      `json:"flash_time_short"`      // "The duration, in seconds, of a “short” flash."
	Icon                 string   `json:"icon"`                  // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	MaxMireds            int      `json:"max_mireds"`            // "The maximum color temperature in mireds."
	MinMireds            int      `json:"min_mireds"`            // "The minimum color temperature in mireds."
	Name                 string   `json:"name"`                  // "The name of the light."
	ObjectId             string   `json:"object_id"`             // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic           bool     `json:"optimistic"`            // "Flag that defines if the light works in optimistic mode."
	PayloadAvailable     string   `json:"payload_available"`     // "The payload that represents the available state."
	PayloadNotAvailable  string   `json:"payload_not_available"` // "The payload that represents the unavailable state."
	Qos                  int      `json:"qos"`                   // "The maximum QoS level of the state topic."
	Retain               bool     `json:"retain"`                // "If the published message should have the retain flag on or not."
	Schema               string   `json:"schema"`                // "The schema to use. Must be `json` to select the JSON schema."
	State                []string `json:"state"`
	SupportedColorModes  []string `json:"supported_color_modes"` // "A list of color modes supported by the list. This is required if `color_mode` is `True`. Possible color modes are `onoff`, `brightness`, `color_temp`, `hs`, `xy`, `rgb`, `rgbw`, `rgbww`."
	UniqueId             string   `json:"unique_id"`             // "An ID that uniquely identifies this light. If two lights have the same unique ID, Home Assistant will raise an exception."
	MQTT                 struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
