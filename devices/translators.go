package devices

func (entity HADeviceAlarmControlPanelFunctionsConfig) Translate() (functions HADeviceAlarmControlPanelFunctions) {
	AddStateFunction(entity.Availability.State, &functions.Availability.State)
	AddCommandFunction(entity.Command, &functions.Command)
	AddStateFunction(entity.State, &functions.State)
	return
}

func (entity HADeviceBinarySensorFunctionsConfig) Translate() (functions HADeviceBinarySensorFunctions) {
	AddStateFunction(entity.Availability.State, &functions.Availability.State)
	AddStateFunction(entity.State, &functions.State)
	return
}

func (entity HADeviceCameraFunctionsConfig) Translate() (functions HADeviceCameraFunctions) {
	AddStateFunction(entity.Availability.State, &functions.Availability.State)
	AddStateFunction(entity.State, &functions.State)
	return
}

func (entity HADeviceCoverFunctionsConfig) Translate() (functions HADeviceCoverFunctions) {
	AddStateFunction(entity.Availability.State, &functions.Availability.State)
	AddCommandFunction(entity.Command, &functions.Command)
	AddCommandFunction(entity.SetPosition, &functions.SetPosition)
	AddStateFunction(entity.State, &functions.State)
	AddCommandFunction(entity.TiltCommand, &functions.TiltCommand)
	AddStateFunction(entity.TiltStatus, &functions.TiltStatus)
	return
}

func (entity HADeviceDeviceTrackerFunctionsConfig) Translate() (functions HADeviceDeviceTrackerFunctions) {
	return
}

func (entity HADeviceDeviceTriggerFunctionsConfig) Translate() (functions HADeviceDeviceTriggerFunctions) {
	AddStateFunction(entity.State, &functions.State)
	return
}

func (entity HADeviceFanFunctionsConfig) Translate() (functions HADeviceFanFunctions) {
	AddStateFunction(entity.Availability.State, &functions.Availability.State)
	AddCommandFunction(entity.Command, &functions.Command)
	AddCommandFunction(entity.OscillationCommand, &functions.OscillationCommand)
	AddStateFunction(entity.OscillationState, &functions.OscillationState)
	AddCommandFunction(entity.PercentageCommand, &functions.PercentageCommand)
	AddStateFunction(entity.PercentageState, &functions.PercentageState)
	AddCommandFunction(entity.PresetModeCommand, &functions.PresetModeCommand)
	AddStateFunction(entity.PresetModeState, &functions.PresetModeState)
	AddStateFunction(entity.State, &functions.State)
	return
}

func (entity HADeviceHumidifierFunctionsConfig) Translate() (functions HADeviceHumidifierFunctions) {
	AddStateFunction(entity.Availability.State, &functions.Availability.State)
	AddCommandFunction(entity.Command, &functions.Command)
	AddCommandFunction(entity.ModeCommand, &functions.ModeCommand)
	AddStateFunction(entity.ModeState, &functions.ModeState)
	AddStateFunction(entity.State, &functions.State)
	AddCommandFunction(entity.TargetHumidityCommand, &functions.TargetHumidityCommand)
	AddStateFunction(entity.TargetHumidityState, &functions.TargetHumidityState)
	return
}

func (entity HADeviceClimateFunctionsConfig) Translate() (functions HADeviceClimateFunctions) {
	AddCommandFunction(entity.AuxCommand, &functions.AuxCommand)
	AddStateFunction(entity.AuxState, &functions.AuxState)
	AddStateFunction(entity.Availability.State, &functions.Availability.State)
	AddCommandFunction(entity.FanModeCommand, &functions.FanModeCommand)
	AddStateFunction(entity.FanModeState, &functions.FanModeState)
	AddCommandFunction(entity.ModeCommand, &functions.ModeCommand)
	AddStateFunction(entity.ModeState, &functions.ModeState)
	AddCommandFunction(entity.PowerCommand, &functions.PowerCommand)
	AddCommandFunction(entity.PresetModeCommand, &functions.PresetModeCommand)
	AddStateFunction(entity.PresetModeState, &functions.PresetModeState)
	AddCommandFunction(entity.SwingModeCommand, &functions.SwingModeCommand)
	AddStateFunction(entity.SwingModeState, &functions.SwingModeState)
	AddCommandFunction(entity.TemperatureCommand, &functions.TemperatureCommand)
	AddCommandFunction(entity.TemperatureHighCommand, &functions.TemperatureHighCommand)
	AddStateFunction(entity.TemperatureHighState, &functions.TemperatureHighState)
	AddCommandFunction(entity.TemperatureLowCommand, &functions.TemperatureLowCommand)
	AddStateFunction(entity.TemperatureLowState, &functions.TemperatureLowState)
	AddStateFunction(entity.TemperatureState, &functions.TemperatureState)
	return
}

func (entity HADeviceLightFunctionsConfig) Translate() (functions HADeviceLightFunctions) {
	AddStateFunction(entity.Availability.State, &functions.Availability.State)
	AddCommandFunction(entity.BrightnessCommand, &functions.BrightnessCommand)
	AddStateFunction(entity.BrightnessState, &functions.BrightnessState)
	AddStateFunction(entity.ColorModeState, &functions.ColorModeState)
	AddCommandFunction(entity.ColorTempCommand, &functions.ColorTempCommand)
	AddStateFunction(entity.ColorTempState, &functions.ColorTempState)
	AddCommandFunction(entity.Command, &functions.Command)
	AddCommandFunction(entity.EffectCommand, &functions.EffectCommand)
	AddStateFunction(entity.EffectState, &functions.EffectState)
	AddCommandFunction(entity.HsCommand, &functions.HsCommand)
	AddStateFunction(entity.HsState, &functions.HsState)
	AddCommandFunction(entity.RgbCommand, &functions.RgbCommand)
	AddStateFunction(entity.RgbState, &functions.RgbState)
	AddStateFunction(entity.State, &functions.State)
	AddCommandFunction(entity.WhiteCommand, &functions.WhiteCommand)
	AddCommandFunction(entity.XyCommand, &functions.XyCommand)
	AddStateFunction(entity.XyState, &functions.XyState)
	return
}

func (entity HADeviceLockFunctionsConfig) Translate() (functions HADeviceLockFunctions) {
	AddStateFunction(entity.Availability.State, &functions.Availability.State)
	AddCommandFunction(entity.Command, &functions.Command)
	AddStateFunction(entity.State, &functions.State)
	return
}

func (entity HADeviceNumberFunctionsConfig) Translate() (functions HADeviceNumberFunctions) {
	AddStateFunction(entity.Availability.State, &functions.Availability.State)
	AddCommandFunction(entity.Command, &functions.Command)
	AddStateFunction(entity.State, &functions.State)
	return
}

func (entity HADeviceSceneFunctionsConfig) Translate() (functions HADeviceSceneFunctions) {
	AddStateFunction(entity.Availability.State, &functions.Availability.State)
	AddCommandFunction(entity.Command, &functions.Command)
	return
}

func (entity HADeviceSelectFunctionsConfig) Translate() (functions HADeviceSelectFunctions) {
	AddStateFunction(entity.Availability.State, &functions.Availability.State)
	AddCommandFunction(entity.Command, &functions.Command)
	AddStateFunction(entity.State, &functions.State)
	return
}

func (entity HADeviceSensorFunctionsConfig) Translate() (functions HADeviceSensorFunctions) {
	AddStateFunction(entity.Availability.State, &functions.Availability.State)
	AddStateFunction(entity.State, &functions.State)
	return
}

func (entity HADeviceSwitchFunctionsConfig) Translate() (functions HADeviceSwitchFunctions) {
	AddStateFunction(entity.Availability.State, &functions.Availability.State)
	AddCommandFunction(entity.Command, &functions.Command)
	AddStateFunction(entity.State, &functions.State)
	return
}

func (entity HADeviceTagFunctionsConfig) Translate() (functions HADeviceTagFunctions) {
	AddStateFunction(entity.State, &functions.State)
	return
}

func (entity HADeviceVacuumFunctionsConfig) Translate() (functions HADeviceVacuumFunctions) {
	AddStateFunction(entity.Availability.State, &functions.Availability.State)
	AddCommandFunction(entity.Command, &functions.Command)
	AddCommandFunction(entity.SendCommand, &functions.SendCommand)
	AddCommandFunction(entity.SetFanSpeed, &functions.SetFanSpeed)
	return
}
