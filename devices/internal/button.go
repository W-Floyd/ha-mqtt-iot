package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Button) Translate() ExternalDevice.Button {
	eDevice := ExternalDevice.Button{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.CommandTemplate = iDevice.CommandTemplate
	eDevice.CommandFunc = common.ConstructCommandFunc(iDevice.Command)
	eDevice.DeviceClass = iDevice.DeviceClass
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.Icon = iDevice.Icon
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.PayloadPress = iDevice.PayloadPress
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.Initialize()
	return eDevice
}

type Button struct {
	CommandTemplate     string   `json:"command_template"`
	Command             []string `json:"command"`
	DeviceClass         string   `json:"device_class"`
	EnabledByDefault    bool     `json:"enabled_by_default"`
	Encoding            string   `json:"encoding"`
	EntityCategory      string   `json:"entity_category"`
	Icon                string   `json:"icon"`
	Name                string   `json:"name"`
	ObjectId            string   `json:"object_id"`
	PayloadAvailable    string   `json:"payload_available"`
	PayloadNotAvailable string   `json:"payload_not_available"`
	PayloadPress        string   `json:"payload_press"`
	Qos                 int      `json:"qos"`
	Retain              bool     `json:"retain"`
	UniqueId            string   `json:"unique_id"`
	MQTT                struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
