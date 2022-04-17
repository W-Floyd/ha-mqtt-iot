package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Scene) Translate() externaldevice.Scene {
	eDevice := externaldevice.Scene{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.CommandFunc = common.ConstructCommandFunc(iDevice.Command)
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.Icon = iDevice.Icon
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.PayloadOn = iDevice.PayloadOn
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.Initialize()
	return eDevice
}

type Scene struct {
	Command             []string `json:"command"`
	EnabledByDefault    bool     `json:"enabled_by_default"`
	EntityCategory      string   `json:"entity_category"`
	Icon                string   `json:"icon"`
	Name                string   `json:"name"`
	ObjectId            string   `json:"object_id"`
	PayloadAvailable    string   `json:"payload_available"`
	PayloadNotAvailable string   `json:"payload_not_available"`
	PayloadOn           string   `json:"payload_on"`
	Qos                 int      `json:"qos"`
	Retain              bool     `json:"retain"`
	UniqueId            string   `json:"unique_id"`
	MQTT                struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
