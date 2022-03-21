package config

import internaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/internaldevice"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
type Config struct {
	MQTT struct {
		Broker       string
		Username     string
		Password     string
		NodeID       string
		InstanceName string
	}
	Logging struct {
		Critical bool
		Debug    bool
		Error    bool
		Warn     bool
	}
	AlarmControlPanel []internaldevice.AlarmControlPanel
	BinarySensor      []internaldevice.BinarySensor
	Button            []internaldevice.Button
	Camera            []internaldevice.Camera
	Cover             []internaldevice.Cover
	DeviceTracker     []internaldevice.DeviceTracker
	DeviceTrigger     []internaldevice.DeviceTrigger
	Fan               []internaldevice.Fan
	Humidifier        []internaldevice.Humidifier
	Climate           []internaldevice.Climate
	Light             []internaldevice.Light
	Lock              []internaldevice.Lock
	Number            []internaldevice.Number
	Scene             []internaldevice.Scene
	Select            []internaldevice.Select
	Sensor            []internaldevice.Sensor
	Siren             []internaldevice.Siren
	Switch            []internaldevice.Switch
	Tag               []internaldevice.Tag
	Vacuum            []internaldevice.Vacuum
}
