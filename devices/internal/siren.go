package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Siren) Translate() ExternalDevice.Siren {
	eDevice := ExternalDevice.Siren{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.AvailableTones = iDevice.AvailableTones
	eDevice.CommandOffTemplate = iDevice.CommandOffTemplate
	eDevice.CommandTemplate = iDevice.CommandTemplate
	eDevice.CommandFunc = common.ConstructCommandFunc(iDevice.Command)
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.Icon = iDevice.Icon
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.Optimistic = iDevice.Optimistic
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.PayloadOff = iDevice.PayloadOff
	eDevice.PayloadOn = iDevice.PayloadOn
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.StateOff = iDevice.StateOff
	eDevice.StateOn = iDevice.StateOn
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.StateValueTemplate = iDevice.StateValueTemplate
	eDevice.SupportDuration = iDevice.SupportDuration
	eDevice.SupportVolumeSet = iDevice.SupportVolumeSet
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.Initialize()
	return eDevice
}

type Siren struct {
	AvailableTones      []string `json:"available_tones"`
	CommandOffTemplate  string   `json:"command_off_template"`
	CommandTemplate     string   `json:"command_template"`
	Command             []string `json:"command"`
	EnabledByDefault    bool     `json:"enabled_by_default"`
	Encoding            string   `json:"encoding"`
	EntityCategory      string   `json:"entity_category"`
	Icon                string   `json:"icon"`
	Name                string   `json:"name"`
	ObjectId            string   `json:"object_id"`
	Optimistic          bool     `json:"optimistic"`
	PayloadAvailable    string   `json:"payload_available"`
	PayloadNotAvailable string   `json:"payload_not_available"`
	PayloadOff          string   `json:"payload_off"`
	PayloadOn           string   `json:"payload_on"`
	Qos                 int      `json:"qos"`
	Retain              bool     `json:"retain"`
	StateOff            string   `json:"state_off"`
	StateOn             string   `json:"state_on"`
	State               []string `json:"state"`
	StateValueTemplate  string   `json:"state_value_template"`
	SupportDuration     bool     `json:"support_duration"`
	SupportVolumeSet    bool     `json:"support_volume_set"`
	UniqueId            string   `json:"unique_id"`
	MQTT                struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
