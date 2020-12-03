package iotconfig

import "../hadiscovery"

type LightHA struct {
	Info                   InfoIcon `json:"info"`
	Command                []string `json:"command"`
	CommandState           []string `json:"command_state"`
	CommandBrightness      []string `json:"command_brightness"`
	CommandBrightnessState []string `json:"command_brightness_state"`
	BrightnessScale        int      `json:"brightness_scale"`
	CommandColorTemp       []string `json:"command_color_temp"`
	CommandColorTempState  []string `json:"command_color_temp_state"`
	CommandEffect          []string `json:"command_effect"`
	CommandEffectState     []string `json:"command_effect_state"`
	CommandHs              []string `json:"command_hs"`
	CommandHsState         []string `json:"command_hs_state"`
	CommandRgb             []string `json:"command_rgb"`
	CommandRgbState        []string `json:"command_rgb_state"`
	CommandWhiteValue      []string `json:"command_white_value"`
	CommandWhiteValueState []string `json:"command_white_value_state"`
	CommandXy              []string `json:"command_xy"`
	CommandXyState         []string `json:"command_xy_state"`
	UpdateInterval         float64  `json:"update_interval"`
	ForceUpdateMQTT        bool     `json:"force_update"`
}

func (li LightHA) Translate() hadiscovery.Light {
	nli := hadiscovery.Light{}
	nli.UpdateInterval = li.UpdateInterval
	nli.ForceUpdateMQTT = li.ForceUpdateMQTT

	nli.Name = li.Info.Name
	nli.UniqueID = li.Info.ID + "_" + hadiscovery.NodeID

	if !almostEqual(li.UpdateInterval, 0.0) {
		nli.UpdateInterval = li.UpdateInterval
	} else {
		nli.UpdateInterval = 1
	}

	if li.BrightnessScale != 0 {
		nli.BrightnessScale = li.BrightnessScale
	}

	nli.ForceUpdateMQTT = li.ForceUpdateMQTT

	if len(li.CommandState) > 0 {
		nli.StateFunc = constructStateFunc(li.CommandState)
	}
	if len(li.CommandBrightnessState) > 0 {
		nli.BrightnessStateFunc = constructStateFunc(li.CommandBrightnessState)
	}
	if len(li.CommandColorTempState) > 0 {
		nli.ColorTempStateFunc = constructStateFunc(li.CommandColorTempState)
	}
	if len(li.CommandEffectState) > 0 {
		nli.EffectStateFunc = constructStateFunc(li.CommandEffectState)
	}
	if len(li.CommandHsState) > 0 {
		nli.HsStateFunc = constructStateFunc(li.CommandHsState)
	}
	if len(li.CommandRgbState) > 0 {
		nli.RgbStateFunc = constructStateFunc(li.CommandRgbState)
	}
	if len(li.CommandWhiteValueState) > 0 {
		nli.WhiteValueStateFunc = constructStateFunc(li.CommandWhiteValueState)
	}
	if len(li.CommandXyState) > 0 {
		nli.XyStateFunc = constructStateFunc(li.CommandXyState)
	}

	if len(li.Command) > 0 {
		nli.CommandFunc = constructCommandFunc(li.Command)
	}
	if len(li.CommandBrightness) > 0 {
		nli.BrightnessCommandFunc = constructCommandFunc(li.CommandBrightness)
	}
	if len(li.CommandColorTemp) > 0 {
		nli.ColorTempCommandFunc = constructCommandFunc(li.CommandColorTemp)
	}
	if len(li.CommandEffect) > 0 {
		nli.EffectCommandFunc = constructCommandFunc(li.CommandEffect)
	}
	if len(li.CommandHs) > 0 {
		nli.HsCommandFunc = constructCommandFunc(li.CommandHs)
	}
	if len(li.CommandRgb) > 0 {
		nli.RgbCommandFunc = constructCommandFunc(li.CommandRgb)
	}
	if len(li.CommandWhiteValue) > 0 {
		nli.WhiteValueCommandFunc = constructCommandFunc(li.CommandWhiteValue)
	}
	if len(li.CommandXy) > 0 {
		nli.XyCommandFunc = constructCommandFunc(li.CommandXy)
	}

	nli.Initialize()

	return nli
}
