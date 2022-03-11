package hadiscovery

func (d AlarmControlPanel) Init() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *AlarmControlPanel) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d BinarySensor) Init() {
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *BinarySensor) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d Button) Init() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *Button) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
}
func (d Camera) Init() {
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *Camera) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
}
func (d Cover) Init() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *Cover) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.PositionFunc != nil {
		d.PositionTopic = GetTopic(d, "position_topic")
	}
	if d.SetPositionFunc != nil {
		d.SetPositionTopic = GetTopic(d, "set_position_topic")
		topicStore[d.SetPositionTopic] = &d.SetPositionFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
	if d.TiltCommandFunc != nil {
		d.TiltCommandTopic = GetTopic(d, "tilt_command_topic")
		topicStore[d.TiltCommandTopic] = &d.TiltCommandFunc
	}
	if d.TiltStatusFunc != nil {
		d.TiltStatusTopic = GetTopic(d, "tilt_status_topic")
	}
}
func (d DeviceTracker) Init() {
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *DeviceTracker) PopulateTopics() {}
func (d DeviceTrigger) Init() {
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *DeviceTrigger) PopulateTopics() {}
func (d Fan) Init() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *Fan) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.OscillationCommandFunc != nil {
		d.OscillationCommandTopic = GetTopic(d, "oscillation_command_topic")
		topicStore[d.OscillationCommandTopic] = &d.OscillationCommandFunc
	}
	if d.OscillationStateFunc != nil {
		d.OscillationStateTopic = GetTopic(d, "oscillation_state_topic")
	}
	if d.PercentageCommandFunc != nil {
		d.PercentageCommandTopic = GetTopic(d, "percentage_command_topic")
		topicStore[d.PercentageCommandTopic] = &d.PercentageCommandFunc
	}
	if d.PercentageStateFunc != nil {
		d.PercentageStateTopic = GetTopic(d, "percentage_state_topic")
	}
	if d.PresetModeCommandFunc != nil {
		d.PresetModeCommandTopic = GetTopic(d, "preset_mode_command_topic")
		topicStore[d.PresetModeCommandTopic] = &d.PresetModeCommandFunc
	}
	if d.PresetModeStateFunc != nil {
		d.PresetModeStateTopic = GetTopic(d, "preset_mode_state_topic")
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d Humidifier) Init() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *Humidifier) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.ModeCommandFunc != nil {
		d.ModeCommandTopic = GetTopic(d, "mode_command_topic")
		topicStore[d.ModeCommandTopic] = &d.ModeCommandFunc
	}
	if d.ModeStateFunc != nil {
		d.ModeStateTopic = GetTopic(d, "mode_state_topic")
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
	if d.TargetHumidityCommandFunc != nil {
		d.TargetHumidityCommandTopic = GetTopic(d, "target_humidity_command_topic")
		topicStore[d.TargetHumidityCommandTopic] = &d.TargetHumidityCommandFunc
	}
	if d.TargetHumidityStateFunc != nil {
		d.TargetHumidityStateTopic = GetTopic(d, "target_humidity_state_topic")
	}
}
func (d Climate) Init() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *Climate) PopulateTopics() {
	if d.ActionFunc != nil {
		d.ActionTopic = GetTopic(d, "action_topic")
		topicStore[d.ActionTopic] = &d.ActionFunc
	}
	if d.AuxCommandFunc != nil {
		d.AuxCommandTopic = GetTopic(d, "aux_command_topic")
		topicStore[d.AuxCommandTopic] = &d.AuxCommandFunc
	}
	if d.AuxStateFunc != nil {
		d.AuxStateTopic = GetTopic(d, "aux_state_topic")
	}
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CurrentTemperatureFunc != nil {
		d.CurrentTemperatureTopic = GetTopic(d, "current_temperature_topic")
	}
	if d.FanModeCommandFunc != nil {
		d.FanModeCommandTopic = GetTopic(d, "fan_mode_command_topic")
		topicStore[d.FanModeCommandTopic] = &d.FanModeCommandFunc
	}
	if d.FanModeStateFunc != nil {
		d.FanModeStateTopic = GetTopic(d, "fan_mode_state_topic")
	}
	if d.ModeCommandFunc != nil {
		d.ModeCommandTopic = GetTopic(d, "mode_command_topic")
		topicStore[d.ModeCommandTopic] = &d.ModeCommandFunc
	}
	if d.ModeStateFunc != nil {
		d.ModeStateTopic = GetTopic(d, "mode_state_topic")
	}
	if d.PowerCommandFunc != nil {
		d.PowerCommandTopic = GetTopic(d, "power_command_topic")
		topicStore[d.PowerCommandTopic] = &d.PowerCommandFunc
	}
	if d.PresetModeCommandFunc != nil {
		d.PresetModeCommandTopic = GetTopic(d, "preset_mode_command_topic")
		topicStore[d.PresetModeCommandTopic] = &d.PresetModeCommandFunc
	}
	if d.PresetModeStateFunc != nil {
		d.PresetModeStateTopic = GetTopic(d, "preset_mode_state_topic")
	}
	if d.SwingModeCommandFunc != nil {
		d.SwingModeCommandTopic = GetTopic(d, "swing_mode_command_topic")
		topicStore[d.SwingModeCommandTopic] = &d.SwingModeCommandFunc
	}
	if d.SwingModeStateFunc != nil {
		d.SwingModeStateTopic = GetTopic(d, "swing_mode_state_topic")
	}
	if d.TemperatureCommandFunc != nil {
		d.TemperatureCommandTopic = GetTopic(d, "temperature_command_topic")
		topicStore[d.TemperatureCommandTopic] = &d.TemperatureCommandFunc
	}
	if d.TemperatureHighCommandFunc != nil {
		d.TemperatureHighCommandTopic = GetTopic(d, "temperature_high_command_topic")
		topicStore[d.TemperatureHighCommandTopic] = &d.TemperatureHighCommandFunc
	}
	if d.TemperatureHighStateFunc != nil {
		d.TemperatureHighStateTopic = GetTopic(d, "temperature_high_state_topic")
	}
	if d.TemperatureLowCommandFunc != nil {
		d.TemperatureLowCommandTopic = GetTopic(d, "temperature_low_command_topic")
		topicStore[d.TemperatureLowCommandTopic] = &d.TemperatureLowCommandFunc
	}
	if d.TemperatureLowStateFunc != nil {
		d.TemperatureLowStateTopic = GetTopic(d, "temperature_low_state_topic")
	}
	if d.TemperatureStateFunc != nil {
		d.TemperatureStateTopic = GetTopic(d, "temperature_state_topic")
	}
}
func (d Light) Init() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *Light) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.BrightnessCommandFunc != nil {
		d.BrightnessCommandTopic = GetTopic(d, "brightness_command_topic")
		topicStore[d.BrightnessCommandTopic] = &d.BrightnessCommandFunc
	}
	if d.BrightnessStateFunc != nil {
		d.BrightnessStateTopic = GetTopic(d, "brightness_state_topic")
	}
	if d.ColorModeStateFunc != nil {
		d.ColorModeStateTopic = GetTopic(d, "color_mode_state_topic")
	}
	if d.ColorTempCommandFunc != nil {
		d.ColorTempCommandTopic = GetTopic(d, "color_temp_command_topic")
		topicStore[d.ColorTempCommandTopic] = &d.ColorTempCommandFunc
	}
	if d.ColorTempStateFunc != nil {
		d.ColorTempStateTopic = GetTopic(d, "color_temp_state_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.EffectCommandFunc != nil {
		d.EffectCommandTopic = GetTopic(d, "effect_command_topic")
		topicStore[d.EffectCommandTopic] = &d.EffectCommandFunc
	}
	if d.EffectStateFunc != nil {
		d.EffectStateTopic = GetTopic(d, "effect_state_topic")
	}
	if d.HsCommandFunc != nil {
		d.HsCommandTopic = GetTopic(d, "hs_command_topic")
		topicStore[d.HsCommandTopic] = &d.HsCommandFunc
	}
	if d.HsStateFunc != nil {
		d.HsStateTopic = GetTopic(d, "hs_state_topic")
	}
	if d.RgbCommandFunc != nil {
		d.RgbCommandTopic = GetTopic(d, "rgb_command_topic")
		topicStore[d.RgbCommandTopic] = &d.RgbCommandFunc
	}
	if d.RgbStateFunc != nil {
		d.RgbStateTopic = GetTopic(d, "rgb_state_topic")
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
	if d.WhiteCommandFunc != nil {
		d.WhiteCommandTopic = GetTopic(d, "white_command_topic")
		topicStore[d.WhiteCommandTopic] = &d.WhiteCommandFunc
	}
	if d.XyCommandFunc != nil {
		d.XyCommandTopic = GetTopic(d, "xy_command_topic")
		topicStore[d.XyCommandTopic] = &d.XyCommandFunc
	}
	if d.XyStateFunc != nil {
		d.XyStateTopic = GetTopic(d, "xy_state_topic")
	}
}
func (d Lock) Init() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *Lock) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d Number) Init() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *Number) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d Scene) Init() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *Scene) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
}
func (d Select) Init() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *Select) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d Sensor) Init() {
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *Sensor) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d Siren) Init() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *Siren) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d Switch) Init() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *Switch) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d Tag) Init() {
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *Tag) PopulateTopics() {}
func (d Vacuum) Init() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
}
func (d *Vacuum) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.SendCommandFunc != nil {
		d.SendCommandTopic = GetTopic(d, "send_command_topic")
		topicStore[d.SendCommandTopic] = &d.SendCommandFunc
	}
	if d.SetFanSpeedFunc != nil {
		d.SetFanSpeedTopic = GetTopic(d, "set_fan_speed_topic")
		topicStore[d.SetFanSpeedTopic] = &d.SetFanSpeedFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
