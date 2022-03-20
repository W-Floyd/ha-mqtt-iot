package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Lock) Translate() ExternalDevice.Lock {
	eDevice := ExternalDevice.Lock{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.CommandFunc = common.ConstructCommandFunc(iDevice.Command)
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.Icon = iDevice.Icon
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.Optimistic = iDevice.Optimistic
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadLock = iDevice.PayloadLock
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.PayloadOpen = iDevice.PayloadOpen
	eDevice.PayloadUnlock = iDevice.PayloadUnlock
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.StateLocked = iDevice.StateLocked
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.StateUnlocked = iDevice.StateUnlocked
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.ValueTemplate = iDevice.ValueTemplate
	eDevice.Initialize()
	return eDevice
}

type Lock struct {
	Command             []string `json:"command"`
	EnabledByDefault    bool     `json:"enabled_by_default"`
	Encoding            string   `json:"encoding"`
	EntityCategory      string   `json:"entity_category"`
	Icon                string   `json:"icon"`
	Name                string   `json:"name"`
	ObjectId            string   `json:"object_id"`
	Optimistic          bool     `json:"optimistic"`
	PayloadAvailable    string   `json:"payload_available"`
	PayloadLock         string   `json:"payload_lock"`
	PayloadNotAvailable string   `json:"payload_not_available"`
	PayloadOpen         string   `json:"payload_open"`
	PayloadUnlock       string   `json:"payload_unlock"`
	Qos                 int      `json:"qos"`
	Retain              bool     `json:"retain"`
	StateLocked         string   `json:"state_locked"`
	State               []string `json:"state"`
	StateUnlocked       string   `json:"state_unlocked"`
	UniqueId            string   `json:"unique_id"`
	ValueTemplate       string   `json:"value_template"`
	MQTT                struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
