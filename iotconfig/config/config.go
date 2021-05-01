package config

import (
	"github.com/W-Floyd/ha-mqtt-iot/iotconfig/binarysensordevice"
	"github.com/W-Floyd/ha-mqtt-iot/iotconfig/lightdevice"
	"github.com/W-Floyd/ha-mqtt-iot/iotconfig/sensordevice"
	"github.com/W-Floyd/ha-mqtt-iot/iotconfig/switchdevice"
)

type Config struct {
	Logging struct {
		Debug    bool `json:"debug"`
		Error    bool `json:"error"`
		Warn     bool `json:"warn"`
		Critical bool `json:"critical"`
	}
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
			Range       struct {
				Minimum float64 `json:"minimum"` // E.g 0.01 for  1% minimum
				Maximum float64 `json:"maximum"` // E.g 0.95 for 95% maximum

			}
		} `json:"backlight"`
		Battery struct {
			Enable bool `json:"enable"`
		} `json:"battery"`
		Crypto []struct {
			CoinName       string  `json:"coin_name"`
			CurrencyName   string  `json:"currency_name"`
			UpdateInterval float64 `json:"update_interval"`
			Icon           string  `json:"icon"`
		} `json:"crypto"`
	} `json:"builtin"`
	Lights        []lightdevice.LightHA                `json:"lights"`
	Switches      []switchdevice.SwitchHA              `json:"switches"`
	Sensors       []sensordevice.SensorHA              `json:"sensors"`
	BinarySensors []binarysensordevice.BinarySensorsHA `json:"binary_sensors"`
}
