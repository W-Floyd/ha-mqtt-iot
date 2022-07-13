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
		Mqtt     bool `json:"mqtt"`
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
		newAlarmControlPanel := d.Translate()
		newDevice := ExternalDevice.Device(&newAlarmControlPanel)
		output = append(output, newDevice)
	}
	for _, d := range c.BinarySensor {
		newBinarySensor := d.Translate()
		newDevice := ExternalDevice.Device(&newBinarySensor)
		output = append(output, newDevice)
	}
	for _, d := range c.Button {
		newButton := d.Translate()
		newDevice := ExternalDevice.Device(&newButton)
		output = append(output, newDevice)
	}
	for _, d := range c.Camera {
		newCamera := d.Translate()
		newDevice := ExternalDevice.Device(&newCamera)
		output = append(output, newDevice)
	}
	for _, d := range c.Cover {
		newCover := d.Translate()
		newDevice := ExternalDevice.Device(&newCover)
		output = append(output, newDevice)
	}
	for _, d := range c.DeviceTracker {
		newDeviceTracker := d.Translate()
		newDevice := ExternalDevice.Device(&newDeviceTracker)
		output = append(output, newDevice)
	}
	for _, d := range c.DeviceTrigger {
		newDeviceTrigger := d.Translate()
		newDevice := ExternalDevice.Device(&newDeviceTrigger)
		output = append(output, newDevice)
	}
	for _, d := range c.Fan {
		newFan := d.Translate()
		newDevice := ExternalDevice.Device(&newFan)
		output = append(output, newDevice)
	}
	for _, d := range c.Humidifier {
		newHumidifier := d.Translate()
		newDevice := ExternalDevice.Device(&newHumidifier)
		output = append(output, newDevice)
	}
	for _, d := range c.Climate {
		newClimate := d.Translate()
		newDevice := ExternalDevice.Device(&newClimate)
		output = append(output, newDevice)
	}
	for _, d := range c.Light {
		newLight := d.Translate()
		newDevice := ExternalDevice.Device(&newLight)
		output = append(output, newDevice)
	}
	for _, d := range c.Lock {
		newLock := d.Translate()
		newDevice := ExternalDevice.Device(&newLock)
		output = append(output, newDevice)
	}
	for _, d := range c.Number {
		newNumber := d.Translate()
		newDevice := ExternalDevice.Device(&newNumber)
		output = append(output, newDevice)
	}
	for _, d := range c.Scene {
		newScene := d.Translate()
		newDevice := ExternalDevice.Device(&newScene)
		output = append(output, newDevice)
	}
	for _, d := range c.Select {
		newSelect := d.Translate()
		newDevice := ExternalDevice.Device(&newSelect)
		output = append(output, newDevice)
	}
	for _, d := range c.Sensor {
		newSensor := d.Translate()
		newDevice := ExternalDevice.Device(&newSensor)
		output = append(output, newDevice)
	}
	for _, d := range c.Siren {
		newSiren := d.Translate()
		newDevice := ExternalDevice.Device(&newSiren)
		output = append(output, newDevice)
	}
	for _, d := range c.Switch {
		newSwitch := d.Translate()
		newDevice := ExternalDevice.Device(&newSwitch)
		output = append(output, newDevice)
	}
	for _, d := range c.Tag {
		newTag := d.Translate()
		newDevice := ExternalDevice.Device(&newTag)
		output = append(output, newDevice)
	}
	for _, d := range c.Vacuum {
		newVacuum := d.Translate()
		newDevice := ExternalDevice.Device(&newVacuum)
		output = append(output, newDevice)
	}
	return
}
