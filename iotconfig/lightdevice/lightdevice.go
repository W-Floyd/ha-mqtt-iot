package lightdevice

import (
	ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"
	"github.com/W-Floyd/ha-mqtt-iot/iotconfig/common"
)

func (li LightHA) Translate() ExternalDevice.Light {
	nli := ExternalDevice.Light{}
	nli.MQTT.ForceUpdate = li.ForceUpdateMQTT

	nli.Name = li.Info.Name
	nli.UniqueId = li.Info.ID + "_" + ExternalDevice.NodeID

	if !common.AlmostEqual(li.UpdateInterval, 0.0) {
		nli.MQTT.UpdateInterval = li.UpdateInterval
	} else {
		nli.MQTT.UpdateInterval = 1
	}

	if li.BrightnessScale != 0 {
		nli.BrightnessScale = li.BrightnessScale
	}

	if li.MaxMireds != 0 {
		nli.MaxMireds = li.MaxMireds
	}

	if li.MinMireds != 0 {
		nli.MinMireds = li.MinMireds
	}

	if len(li.CommandState) > 0 {
		nli.StateFunc = common.ConstructStateFunc(li.CommandState)
	}

	if len(li.Command) > 0 {
		nli.CommandFunc = common.ConstructCommandFunc(li.Command)
	}

	nli.Initialize()

	return nli
}
