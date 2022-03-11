package hadiscovery

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
