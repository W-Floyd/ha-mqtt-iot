package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Select) Translate() ExternalDevice.Select {
	eDevice := ExternalDevice.Select{}
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
	eDevice.Options = iDevice.Options
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.ValueTemplate = iDevice.ValueTemplate
	eDevice.Initialize()
	return eDevice
}

type Select struct {
	CommandTemplate  string   `json:"command_template"`
	Command          []string `json:"command"`
	EnabledByDefault bool     `json:"enabled_by_default"`
	Encoding         string   `json:"encoding"`
	EntityCategory   string   `json:"entity_category"`
	Icon             string   `json:"icon"`
	Name             string   `json:"name"`
	ObjectId         string   `json:"object_id"`
	Optimistic       bool     `json:"optimistic"`
	Options          []string `json:"options"`
	Qos              int      `json:"qos"`
	Retain           bool     `json:"retain"`
	State            []string `json:"state"`
	UniqueId         string   `json:"unique_id"`
	ValueTemplate    string   `json:"value_template"`
	MQTT             struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
