package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
func (iDevice Lock) Translate() externaldevice.Lock {
	eDevice := externaldevice.Lock{}
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
	EnabledByDefault    bool     `json:"enabled_by_default"`    // "Flag which defines if the entity should be enabled when first added."
	Encoding            string   `json:"encoding"`              // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory      string   `json:"entity_category"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                string   `json:"icon"`                  // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Name                string   `json:"name"`                  // "The name of the lock."
	ObjectId            string   `json:"object_id"`             // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic          bool     `json:"optimistic"`            // "Flag that defines if lock works in optimistic mode."
	PayloadAvailable    string   `json:"payload_available"`     // "The payload that represents the available state."
	PayloadLock         string   `json:"payload_lock"`          // "The payload sent to the lock to lock it."
	PayloadNotAvailable string   `json:"payload_not_available"` // "The payload that represents the unavailable state."
	PayloadOpen         string   `json:"payload_open"`          // "The payload sent to the lock to open it."
	PayloadUnlock       string   `json:"payload_unlock"`        // "The payload sent to the lock to unlock it."
	Qos                 int      `json:"qos"`                   // "The maximum QoS level of the state topic."
	Retain              bool     `json:"retain"`                // "If the published message should have the retain flag on or not."
	StateLocked         string   `json:"state_locked"`          // "The payload sent to by the lock when it's locked."
	State               []string `json:"state"`
	StateUnlocked       string   `json:"state_unlocked"` // "The payload sent to by the lock when it's unlocked."
	UniqueId            string   `json:"unique_id"`      // "An ID that uniquely identifies this lock. If two locks have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate       string   `json:"value_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract a value from the payload."
	MQTT                struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
