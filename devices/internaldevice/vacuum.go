package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
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
	Encoding            string   `json:"encoding"`
	FanSpeedList        []string `json:"fan_speed_list"`
	Name                string   `json:"name"`
	ObjectId            string   `json:"object_id"`
	PayloadAvailable    string   `json:"payload_available"`
	PayloadCleanSpot    string   `json:"payload_clean_spot"`
	PayloadLocate       string   `json:"payload_locate"`
	PayloadNotAvailable string   `json:"payload_not_available"`
	PayloadPause        string   `json:"payload_pause"`
	PayloadReturnToBase string   `json:"payload_return_to_base"`
	PayloadStart        string   `json:"payload_start"`
	PayloadStop         string   `json:"payload_stop"`
	Qos                 int      `json:"qos"`
	Retain              bool     `json:"retain"`
	Schema              string   `json:"schema"`
	SendCommand         []string `json:"send_command"`
	SetFanSpeed         []string `json:"set_fan_speed"`
	State               []string `json:"state"`
	SupportedFeatures   []string `json:"supported_features"`
	UniqueId            string   `json:"unique_id"`
	MQTT                struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
