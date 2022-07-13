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
		Broker       string `json:"broker"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		NodeId       string `json:"node_id"`
		InstanceName string `json:"instance_name"`
	}
	Logging struct {
		Critical bool `json:"critical"`
		Debug    bool `json:"debug"`
		Error    bool `json:"error"`
		Warn     bool `json:"warn"`
	}
	AlarmControlPanel []internaldevice.AlarmControlPanel `json:"alarm_control_panel"`
	BinarySensor      []internaldevice.BinarySensor      `json:"binary_sensor"`
	Button            []internaldevice.Button            `json:"button"`
	Camera            []internaldevice.Camera            `json:"camera"`
	Cover             []internaldevice.Cover             `json:"cover"`
	DeviceTracker     []internaldevice.DeviceTracker     `json:"device_tracker"`
	DeviceTrigger     []internaldevice.DeviceTrigger     `json:"device_trigger"`
	Fan               []internaldevice.Fan               `json:"fan"`
	Humidifier        []internaldevice.Humidifier        `json:"humidifier"`
	Climate           []internaldevice.Climate           `json:"climate"`
	Light             []internaldevice.Light             `json:"light"`
	Lock              []internaldevice.Lock              `json:"lock"`
	Number            []internaldevice.Number            `json:"number"`
	Scene             []internaldevice.Scene             `json:"scene"`
	Select            []internaldevice.Select            `json:"select"`
	Sensor            []internaldevice.Sensor            `json:"sensor"`
	Siren             []internaldevice.Siren             `json:"siren"`
	Switch            []internaldevice.Switch            `json:"switch"`
	Tag               []internaldevice.Tag               `json:"tag"`
	Vacuum            []internaldevice.Vacuum            `json:"vacuum"`
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
