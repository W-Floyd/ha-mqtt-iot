package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Sensor) Translate() externaldevice.Sensor {
	eDevice := externaldevice.Sensor{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.DeviceClass = iDevice.DeviceClass
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.ExpireAfter = iDevice.ExpireAfter
	eDevice.ForceUpdate = iDevice.ForceUpdate
	eDevice.Icon = iDevice.Icon
	eDevice.LastResetValueTemplate = iDevice.LastResetValueTemplate
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.Qos = iDevice.Qos
	eDevice.StateClass = iDevice.StateClass
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.UnitOfMeasurement = iDevice.UnitOfMeasurement
	eDevice.ValueTemplate = iDevice.ValueTemplate
	eDevice.Initialize()
	return eDevice
}

type Sensor struct {
	DeviceClass            string   `json:"device_class"`
	EnabledByDefault       bool     `json:"enabled_by_default"`
	Encoding               string   `json:"encoding"`
	EntityCategory         string   `json:"entity_category"`
	ExpireAfter            int      `json:"expire_after"`
	ForceUpdate            bool     `json:"force_update"`
	Icon                   string   `json:"icon"`
	LastResetValueTemplate string   `json:"last_reset_value_template"`
	Name                   string   `json:"name"`
	ObjectId               string   `json:"object_id"`
	PayloadAvailable       string   `json:"payload_available"`
	PayloadNotAvailable    string   `json:"payload_not_available"`
	Qos                    int      `json:"qos"`
	StateClass             string   `json:"state_class"`
	State                  []string `json:"state"`
	UniqueId               string   `json:"unique_id"`
	UnitOfMeasurement      string   `json:"unit_of_measurement"`
	ValueTemplate          string   `json:"value_template"`
	MQTT                   struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
