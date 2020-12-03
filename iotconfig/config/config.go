package config

import (
	"../binarysensordevice"
	"../lightdevice"
	"../sensordevice"
	"../switchdevice"
)

type Config struct {
	MQTT struct {
		Broker       string `json:"broker"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		NodeID       string `json:"node_id"`
		InstanceName string `json:"instance_name"`
	} `json:"mqtt"`
	Builtin struct {
		Prefix    string `json:"prefix"`
		Backlight struct {
			Enable      bool `json:"enable"`
			Temperature bool `json:"temperature"`
		} `json:"backlight"`
		Battery struct {
			Enable bool `json:"enable"`
		} `json:"battery"`
	} `json:"builtin"`
	Lights        []lightdevice.LightHA                `json:"lights"`
	Switches      []switchdevice.SwitchHA              `json:"switches"`
	Sensors       []sensordevice.SensorHA              `json:"sensors"`
	BinarySensors []binarysensordevice.BinarySensorsHA `json:"binary_sensors"`
}
