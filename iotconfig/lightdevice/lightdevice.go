package lightdevice

import (
	"github.com/W-Floyd/ha-mqtt-iot/hadiscovery"
	"github.com/W-Floyd/ha-mqtt-iot/iotconfig/common"
)

type LightHA struct {
	Info                   common.InfoIcon `json:"info"`
	Command                []string        `json:"command"`
	CommandState           []string        `json:"command_state"`
	CommandBrightness      []string        `json:"command_brightness"`
	CommandBrightnessState []string        `json:"command_brightness_state"`
	BrightnessScale        int             `json:"brightness_scale"`
	MaxMireds              int             `json:"max_mireds,omitempty"`
	MinMireds              int             `json:"min_mireds,omitempty"`
	CommandColorTemp       []string        `json:"command_color_temp"`
	CommandColorTempState  []string        `json:"command_color_temp_state"`
	CommandEffect          []string        `json:"command_effect"`
	CommandEffectState     []string        `json:"command_effect_state"`
	CommandHs              []string        `json:"command_hs"`
	CommandHsState         []string        `json:"command_hs_state"`
	CommandRgb             []string        `json:"command_rgb"`
	CommandRgbState        []string        `json:"command_rgb_state"`
	CommandWhiteValue      []string        `json:"command_white_value"`
	CommandWhiteValueState []string        `json:"command_white_value_state"`
	CommandXy              []string        `json:"command_xy"`
	CommandXyState         []string        `json:"command_xy_state"`
	UpdateInterval         float64         `json:"update_interval"`
	ForceUpdateMQTT        bool            `json:"force_update"`
}

func (li LightHA) Translate() hadiscovery.Light {
	nli := hadiscovery.Light{}
	nli.MQTT.ForceUpdate = li.ForceUpdateMQTT

	nli.Name = li.Info.Name
	nli.UniqueId = li.Info.ID + "_" + hadiscovery.NodeID

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
