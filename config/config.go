package config

import (
	ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
	internaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/internaldevice"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Config struct {
	MQTT struct {
		Broker       string `json:"broker,omitempty"`
		Username     string `json:"username,omitempty"`
		Password     string `json:"password,omitempty"`
		NodeId       string `json:"node_id,omitempty"`
		InstanceName string `json:"instance_name,omitempty"`
	}
	Logging struct {
		Critical bool `json:"critical,omitempty"`
		Debug    bool `json:"debug,omitempty"`
		Error    bool `json:"error,omitempty"`
		Warn     bool `json:"warn,omitempty"`
		Mqtt     bool `json:"mqtt,omitempty"`
	}
	AlarmControlPanel []internaldevice.AlarmControlPanel `json:"alarm_control_panel,omitempty"`
	BinarySensor      []internaldevice.BinarySensor      `json:"binary_sensor,omitempty"`
	Button            []internaldevice.Button            `json:"button,omitempty"`
	Camera            []internaldevice.Camera            `json:"camera,omitempty"`
	Climate           []internaldevice.Climate           `json:"climate,omitempty"`
	Cover             []internaldevice.Cover             `json:"cover,omitempty"`
	DeviceTracker     []internaldevice.DeviceTracker     `json:"device_tracker,omitempty"`
	DeviceTrigger     []internaldevice.DeviceTrigger     `json:"device_trigger,omitempty"`
	Fan               []internaldevice.Fan               `json:"fan,omitempty"`
	Humidifier        []internaldevice.Humidifier        `json:"humidifier,omitempty"`
	Image             []internaldevice.Image             `json:"image,omitempty"`
	Light             []internaldevice.Light             `json:"light,omitempty"`
	Lock              []internaldevice.Lock              `json:"lock,omitempty"`
	Number            []internaldevice.Number            `json:"number,omitempty"`
	Scene             []internaldevice.Scene             `json:"scene,omitempty"`
	Select            []internaldevice.Select            `json:"select,omitempty"`
	Sensor            []internaldevice.Sensor            `json:"sensor,omitempty"`
	Siren             []internaldevice.Siren             `json:"siren,omitempty"`
	Switch            []internaldevice.Switch            `json:"switch,omitempty"`
	Tag               []internaldevice.Tag               `json:"tag,omitempty"`
	Text              []internaldevice.Text              `json:"text,omitempty"`
	Update            []internaldevice.Update            `json:"update,omitempty"`
	Vacuum            []internaldevice.Vacuum            `json:"vacuum,omitempty"`
	WaterHeater       []internaldevice.WaterHeater       `json:"water_heater,omitempty"`
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
	for _, d := range c.Climate {
		newClimate := d.Translate()
		newDevice := ExternalDevice.Device(&newClimate)
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
	for _, d := range c.Image {
		newImage := d.Translate()
		newDevice := ExternalDevice.Device(&newImage)
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
	for _, d := range c.Text {
		newText := d.Translate()
		newDevice := ExternalDevice.Device(&newText)
		output = append(output, newDevice)
	}
	for _, d := range c.Update {
		newUpdate := d.Translate()
		newDevice := ExternalDevice.Device(&newUpdate)
		output = append(output, newDevice)
	}
	for _, d := range c.Vacuum {
		newVacuum := d.Translate()
		newDevice := ExternalDevice.Device(&newVacuum)
		output = append(output, newDevice)
	}
	for _, d := range c.WaterHeater {
		newWaterHeater := d.Translate()
		newDevice := ExternalDevice.Device(&newWaterHeater)
		output = append(output, newDevice)
	}
	return
}
