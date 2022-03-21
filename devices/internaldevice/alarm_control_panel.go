package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice AlarmControlPanel) Translate() externaldevice.AlarmControlPanel {
	eDevice := externaldevice.AlarmControlPanel{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.Code = iDevice.Code
	eDevice.CodeArmRequired = iDevice.CodeArmRequired
	eDevice.CodeDisarmRequired = iDevice.CodeDisarmRequired
	eDevice.CodeTriggerRequired = iDevice.CodeTriggerRequired
	eDevice.CommandTemplate = iDevice.CommandTemplate
	eDevice.CommandFunc = common.ConstructCommandFunc(iDevice.Command)
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.Icon = iDevice.Icon
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.PayloadArmAway = iDevice.PayloadArmAway
	eDevice.PayloadArmCustomBypass = iDevice.PayloadArmCustomBypass
	eDevice.PayloadArmHome = iDevice.PayloadArmHome
	eDevice.PayloadArmNight = iDevice.PayloadArmNight
	eDevice.PayloadArmVacation = iDevice.PayloadArmVacation
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadDisarm = iDevice.PayloadDisarm
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.PayloadTrigger = iDevice.PayloadTrigger
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.ValueTemplate = iDevice.ValueTemplate
	eDevice.Initialize()
	return eDevice
}

type AlarmControlPanel struct {
	Code                   string   `json:"code"`
	CodeArmRequired        bool     `json:"code_arm_required"`
	CodeDisarmRequired     bool     `json:"code_disarm_required"`
	CodeTriggerRequired    bool     `json:"code_trigger_required"`
	CommandTemplate        string   `json:"command_template"`
	Command                []string `json:"command"`
	EnabledByDefault       bool     `json:"enabled_by_default"`
	Encoding               string   `json:"encoding"`
	EntityCategory         string   `json:"entity_category"`
	Icon                   string   `json:"icon"`
	Name                   string   `json:"name"`
	ObjectId               string   `json:"object_id"`
	PayloadArmAway         string   `json:"payload_arm_away"`
	PayloadArmCustomBypass string   `json:"payload_arm_custom_bypass"`
	PayloadArmHome         string   `json:"payload_arm_home"`
	PayloadArmNight        string   `json:"payload_arm_night"`
	PayloadArmVacation     string   `json:"payload_arm_vacation"`
	PayloadAvailable       string   `json:"payload_available"`
	PayloadDisarm          string   `json:"payload_disarm"`
	PayloadNotAvailable    string   `json:"payload_not_available"`
	PayloadTrigger         string   `json:"payload_trigger"`
	Qos                    int      `json:"qos"`
	Retain                 bool     `json:"retain"`
	State                  []string `json:"state"`
	UniqueId               string   `json:"unique_id"`
	ValueTemplate          string   `json:"value_template"`
	MQTT                   struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
