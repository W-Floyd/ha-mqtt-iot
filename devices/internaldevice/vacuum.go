package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
func (iDevice Vacuum) Translate() externaldevice.Vacuum {
	eDevice := externaldevice.Vacuum{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.CommandFunc = common.ConstructCommandFunc(iDevice.Command)
	eDevice.Encoding = iDevice.Encoding
	eDevice.FanSpeedList = iDevice.FanSpeedList
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadCleanSpot = iDevice.PayloadCleanSpot
	eDevice.PayloadLocate = iDevice.PayloadLocate
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.PayloadPause = iDevice.PayloadPause
	eDevice.PayloadReturnToBase = iDevice.PayloadReturnToBase
	eDevice.PayloadStart = iDevice.PayloadStart
	eDevice.PayloadStop = iDevice.PayloadStop
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.Schema = iDevice.Schema
	eDevice.SendCommandFunc = common.ConstructCommandFunc(iDevice.SendCommand)
	eDevice.SetFanSpeedFunc = common.ConstructCommandFunc(iDevice.SetFanSpeed)
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.SupportedFeatures = iDevice.SupportedFeatures
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.Initialize()
	return eDevice
}

type Vacuum struct {
	Command             []string `json:"command"`
	Encoding            string   `json:"encoding"`               // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	FanSpeedList        []string `json:"fan_speed_list"`         // "List of possible fan speeds for the vacuum."
	Name                string   `json:"name"`                   // "The name of the vacuum."
	ObjectId            string   `json:"object_id"`              // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable    string   `json:"payload_available"`      // "The payload that represents the available state."
	PayloadCleanSpot    string   `json:"payload_clean_spot"`     // "The payload to send to the `command_topic` to begin a spot cleaning cycle."
	PayloadLocate       string   `json:"payload_locate"`         // "The payload to send to the `command_topic` to locate the vacuum (typically plays a song)."
	PayloadNotAvailable string   `json:"payload_not_available"`  // "The payload that represents the unavailable state."
	PayloadPause        string   `json:"payload_pause"`          // "The payload to send to the `command_topic` to pause the vacuum."
	PayloadReturnToBase string   `json:"payload_return_to_base"` // "The payload to send to the `command_topic` to tell the vacuum to return to base."
	PayloadStart        string   `json:"payload_start"`          // "The payload to send to the `command_topic` to begin the cleaning cycle."
	PayloadStop         string   `json:"payload_stop"`           // "The payload to send to the `command_topic` to stop cleaning."
	Qos                 int      `json:"qos"`                    // "The maximum QoS level of the state topic."
	Retain              bool     `json:"retain"`                 // "If the published message should have the retain flag on or not."
	Schema              string   `json:"schema"`                 // "The schema to use. Must be `state` to select the state schema."
	SendCommand         []string `json:"send_command"`
	SetFanSpeed         []string `json:"set_fan_speed"`
	State               []string `json:"state"`
	SupportedFeatures   []string `json:"supported_features"` // "List of features that the vacuum supports (possible values are `start`, `stop`, `pause`, `return_home`, `battery`, `status`, `locate`, `clean_spot`, `fan_speed`, `send_command`)."
	UniqueId            string   `json:"unique_id"`          // "An ID that uniquely identifies this vacuum. If two vacuums have the same unique ID, Home Assistant will raise an exception."
	MQTT                struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
