package config

import (
	"github.com/W-Floyd/ha-mqtt-iot/iotconfig/lightdevice"
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
	Lights []lightdevice.LightHA `json:"lights"`
}
