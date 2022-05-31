package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
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
	EnabledByDefault    bool     `json:"enabled_by_default"`    // "Flag which defines if the entity should be enabled when first added."
	EntityCategory      string   `json:"entity_category"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                string   `json:"icon"`                  // "Icon for the scene."
	Name                string   `json:"name"`                  // "The name to use when displaying this scene."
	ObjectId            string   `json:"object_id"`             // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable    string   `json:"payload_available"`     // "The payload that represents the available state."
	PayloadNotAvailable string   `json:"payload_not_available"` // "The payload that represents the unavailable state."
	PayloadOn           string   `json:"payload_on"`            // "The payload that will be sent to `command_topic` when activating the MQTT scene."
	Qos                 int      `json:"qos"`                   // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain              bool     `json:"retain"`                // "If the published message should have the retain flag on or not."
	UniqueId            string   `json:"unique_id"`             // "An ID that uniquely identifies this scene entity. If two scenes have the same unique ID, Home Assistant will raise an exception."
	MQTT                struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}