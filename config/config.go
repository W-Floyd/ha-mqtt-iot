package config

import (
	ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
	internaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/internaldevice"
)

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

func (c Config) Translate() (output []ExternalDevice.Device) {
	for _, d := range c.AlarmControlPanel {
		output = append(output, d.Translate())
	}
	for _, d := range c.BinarySensor {
		output = append(output, d.Translate())
	}
	for _, d := range c.Button {
		output = append(output, d.Translate())
	}
	for _, d := range c.Camera {
		output = append(output, d.Translate())
	}
	for _, d := range c.Cover {
		output = append(output, d.Translate())
	}
	for _, d := range c.DeviceTracker {
		output = append(output, d.Translate())
	}
	for _, d := range c.DeviceTrigger {
		output = append(output, d.Translate())
	}
	for _, d := range c.Fan {
		output = append(output, d.Translate())
	}
	for _, d := range c.Humidifier {
		output = append(output, d.Translate())
	}
	for _, d := range c.Climate {
		output = append(output, d.Translate())
	}
	for _, d := range c.Light {
		output = append(output, d.Translate())
	}
	for _, d := range c.Lock {
		output = append(output, d.Translate())
	}
	for _, d := range c.Number {
		output = append(output, d.Translate())
	}
	for _, d := range c.Scene {
		output = append(output, d.Translate())
	}
	for _, d := range c.Select {
		output = append(output, d.Translate())
	}
	for _, d := range c.Sensor {
		output = append(output, d.Translate())
	}
	for _, d := range c.Siren {
		output = append(output, d.Translate())
	}
	for _, d := range c.Switch {
		output = append(output, d.Translate())
	}
	for _, d := range c.Tag {
		output = append(output, d.Translate())
	}
	for _, d := range c.Vacuum {
		output = append(output, d.Translate())
	}
	return
}
