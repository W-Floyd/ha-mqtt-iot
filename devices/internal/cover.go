package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Cover) Translate() ExternalDevice.Cover {
	eDevice := ExternalDevice.Cover{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.CommandFunc = common.ConstructCommandFunc(iDevice.Command)
	eDevice.DeviceClass = iDevice.DeviceClass
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.Icon = iDevice.Icon
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.Optimistic = iDevice.Optimistic
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadClose = iDevice.PayloadClose
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.PayloadOpen = iDevice.PayloadOpen
	eDevice.PayloadStop = iDevice.PayloadStop
	eDevice.PositionClosed = iDevice.PositionClosed
	eDevice.PositionOpen = iDevice.PositionOpen
	eDevice.PositionTemplate = iDevice.PositionTemplate
	eDevice.PositionFunc = common.ConstructStateFunc(iDevice.Position)
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.SetPositionTemplate = iDevice.SetPositionTemplate
	eDevice.SetPositionFunc = common.ConstructCommandFunc(iDevice.SetPosition)
	eDevice.StateClosed = iDevice.StateClosed
	eDevice.StateClosing = iDevice.StateClosing
	eDevice.StateOpen = iDevice.StateOpen
	eDevice.StateOpening = iDevice.StateOpening
	eDevice.StateStopped = iDevice.StateStopped
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.TiltClosedValue = iDevice.TiltClosedValue
	eDevice.TiltCommandTemplate = iDevice.TiltCommandTemplate
	eDevice.TiltCommandFunc = common.ConstructCommandFunc(iDevice.TiltCommand)
	eDevice.TiltMax = iDevice.TiltMax
	eDevice.TiltMin = iDevice.TiltMin
	eDevice.TiltOpenedValue = iDevice.TiltOpenedValue
	eDevice.TiltOptimistic = iDevice.TiltOptimistic
	eDevice.TiltStatusTemplate = iDevice.TiltStatusTemplate
	eDevice.TiltStatusFunc = common.ConstructStateFunc(iDevice.TiltStatus)
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.ValueTemplate = iDevice.ValueTemplate
	eDevice.Initialize()
	return eDevice
}

type Cover struct {
	Command             []string `json:"command"`
	DeviceClass         string   `json:"device_class"`
	EnabledByDefault    bool     `json:"enabled_by_default"`
	Encoding            string   `json:"encoding"`
	EntityCategory      string   `json:"entity_category"`
	Icon                string   `json:"icon"`
	Name                string   `json:"name"`
	ObjectId            string   `json:"object_id"`
	Optimistic          bool     `json:"optimistic"`
	PayloadAvailable    string   `json:"payload_available"`
	PayloadClose        string   `json:"payload_close"`
	PayloadNotAvailable string   `json:"payload_not_available"`
	PayloadOpen         string   `json:"payload_open"`
	PayloadStop         string   `json:"payload_stop"`
	PositionClosed      int      `json:"position_closed"`
	PositionOpen        int      `json:"position_open"`
	PositionTemplate    string   `json:"position_template"`
	Position            []string `json:"position"`
	Qos                 int      `json:"qos"`
	Retain              bool     `json:"retain"`
	SetPositionTemplate string   `json:"set_position_template"`
	SetPosition         []string `json:"set_position"`
	StateClosed         string   `json:"state_closed"`
	StateClosing        string   `json:"state_closing"`
	StateOpen           string   `json:"state_open"`
	StateOpening        string   `json:"state_opening"`
	StateStopped        string   `json:"state_stopped"`
	State               []string `json:"state"`
	TiltClosedValue     int      `json:"tilt_closed_value"`
	TiltCommandTemplate string   `json:"tilt_command_template"`
	TiltCommand         []string `json:"tilt_command"`
	TiltMax             int      `json:"tilt_max"`
	TiltMin             int      `json:"tilt_min"`
	TiltOpenedValue     int      `json:"tilt_opened_value"`
	TiltOptimistic      bool     `json:"tilt_optimistic"`
	TiltStatusTemplate  string   `json:"tilt_status_template"`
	TiltStatus          []string `json:"tilt_status"`
	UniqueId            string   `json:"unique_id"`
	ValueTemplate       string   `json:"value_template"`
	MQTT                struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
