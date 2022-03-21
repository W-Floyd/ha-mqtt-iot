package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Humidifier) Translate() externaldevice.Humidifier {
	eDevice := externaldevice.Humidifier{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.CommandTemplate = iDevice.CommandTemplate
	eDevice.CommandFunc = common.ConstructCommandFunc(iDevice.Command)
	eDevice.DeviceClass = iDevice.DeviceClass
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.Icon = iDevice.Icon
	eDevice.MaxHumidity = iDevice.MaxHumidity
	eDevice.MinHumidity = iDevice.MinHumidity
	eDevice.ModeCommandTemplate = iDevice.ModeCommandTemplate
	eDevice.ModeCommandFunc = common.ConstructCommandFunc(iDevice.ModeCommand)
	eDevice.ModeStateTemplate = iDevice.ModeStateTemplate
	eDevice.ModeStateFunc = common.ConstructStateFunc(iDevice.ModeState)
	eDevice.Modes = iDevice.Modes
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.Optimistic = iDevice.Optimistic
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.PayloadOff = iDevice.PayloadOff
	eDevice.PayloadOn = iDevice.PayloadOn
	eDevice.PayloadResetHumidity = iDevice.PayloadResetHumidity
	eDevice.PayloadResetMode = iDevice.PayloadResetMode
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.StateValueTemplate = iDevice.StateValueTemplate
	eDevice.TargetHumidityCommandTemplate = iDevice.TargetHumidityCommandTemplate
	eDevice.TargetHumidityCommandFunc = common.ConstructCommandFunc(iDevice.TargetHumidityCommand)
	eDevice.TargetHumidityStateTemplate = iDevice.TargetHumidityStateTemplate
	eDevice.TargetHumidityStateFunc = common.ConstructStateFunc(iDevice.TargetHumidityState)
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.Initialize()
	return eDevice
}

type Humidifier struct {
	CommandTemplate               string   `json:"command_template"`
	Command                       []string `json:"command"`
	DeviceClass                   string   `json:"device_class"`
	EnabledByDefault              bool     `json:"enabled_by_default"`
	Encoding                      string   `json:"encoding"`
	EntityCategory                string   `json:"entity_category"`
	Icon                          string   `json:"icon"`
	MaxHumidity                   int      `json:"max_humidity"`
	MinHumidity                   int      `json:"min_humidity"`
	ModeCommandTemplate           string   `json:"mode_command_template"`
	ModeCommand                   []string `json:"mode_command"`
	ModeStateTemplate             string   `json:"mode_state_template"`
	ModeState                     []string `json:"mode_state"`
	Modes                         []string `json:"modes"`
	Name                          string   `json:"name"`
	ObjectId                      string   `json:"object_id"`
	Optimistic                    bool     `json:"optimistic"`
	PayloadAvailable              string   `json:"payload_available"`
	PayloadNotAvailable           string   `json:"payload_not_available"`
	PayloadOff                    string   `json:"payload_off"`
	PayloadOn                     string   `json:"payload_on"`
	PayloadResetHumidity          string   `json:"payload_reset_humidity"`
	PayloadResetMode              string   `json:"payload_reset_mode"`
	Qos                           int      `json:"qos"`
	Retain                        bool     `json:"retain"`
	State                         []string `json:"state"`
	StateValueTemplate            string   `json:"state_value_template"`
	TargetHumidityCommandTemplate string   `json:"target_humidity_command_template"`
	TargetHumidityCommand         []string `json:"target_humidity_command"`
	TargetHumidityStateTemplate   string   `json:"target_humidity_state_template"`
	TargetHumidityState           []string `json:"target_humidity_state"`
	UniqueId                      string   `json:"unique_id"`
	MQTT                          struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
