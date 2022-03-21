package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Number) Translate() externaldevice.Number {
	eDevice := externaldevice.Number{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.CommandTemplate = iDevice.CommandTemplate
	eDevice.CommandFunc = common.ConstructCommandFunc(iDevice.Command)
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.Icon = iDevice.Icon
	eDevice.Max = iDevice.Max
	eDevice.Min = iDevice.Min
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.Optimistic = iDevice.Optimistic
	eDevice.PayloadReset = iDevice.PayloadReset
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.Step = iDevice.Step
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.UnitOfMeasurement = iDevice.UnitOfMeasurement
	eDevice.ValueTemplate = iDevice.ValueTemplate
	eDevice.Initialize()
	return eDevice
}

type Number struct {
	CommandTemplate   string   `json:"command_template"`
	Command           []string `json:"command"`
	EnabledByDefault  bool     `json:"enabled_by_default"`
	Encoding          string   `json:"encoding"`
	EntityCategory    string   `json:"entity_category"`
	Icon              string   `json:"icon"`
	Max               float64  `json:"max"`
	Min               float64  `json:"min"`
	Name              string   `json:"name"`
	ObjectId          string   `json:"object_id"`
	Optimistic        bool     `json:"optimistic"`
	PayloadReset      string   `json:"payload_reset"`
	Qos               int      `json:"qos"`
	Retain            bool     `json:"retain"`
	State             []string `json:"state"`
	Step              float64  `json:"step"`
	UniqueId          string   `json:"unique_id"`
	UnitOfMeasurement string   `json:"unit_of_measurement"`
	ValueTemplate     string   `json:"value_template"`
	MQTT              struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
