package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Light) Translate() ExternalDevice.Light {
	eDevice := ExternalDevice.Light{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
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
	eDevice.Initialize()
	return eDevice
}

type Light struct {
	Brightness          bool     `json:"brightness"`
	BrightnessScale     int      `json:"brightness_scale"`
	ColorMode           bool     `json:"color_mode"`
	Command             []string `json:"command"`
	Effect              bool     `json:"effect"`
	EffectList          []string `json:"effect_list"`
	EnabledByDefault    bool     `json:"enabled_by_default"`
	Encoding            string   `json:"encoding"`
	EntityCategory      string   `json:"entity_category"`
	FlashTimeLong       int      `json:"flash_time_long"`
	FlashTimeShort      int      `json:"flash_time_short"`
	Icon                string   `json:"icon"`
	MaxMireds           int      `json:"max_mireds"`
	MinMireds           int      `json:"min_mireds"`
	Name                string   `json:"name"`
	ObjectId            string   `json:"object_id"`
	Optimistic          bool     `json:"optimistic"`
	PayloadAvailable    string   `json:"payload_available"`
	PayloadNotAvailable string   `json:"payload_not_available"`
	Qos                 int      `json:"qos"`
	Retain              bool     `json:"retain"`
	Schema              string   `json:"schema"`
	State               []string `json:"state"`
	SupportedColorModes []string `json:"supported_color_modes"`
	UniqueId            string   `json:"unique_id"`
	MQTT                struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
