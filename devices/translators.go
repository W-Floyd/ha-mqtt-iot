package devices

func (entity HADeviceAlarmControlPanelFunctionsConfig) Translate() (output HADeviceAlarmControlPanelFunctions) {
	AddStateFunction(entity.Availability.State, &output.Availability.State)
	AddCommandFunction(entity.Command, &output.Command)
	AddStateFunction(entity.State, &output.State)
	return
}

func (entity HADeviceBinarySensorFunctionsConfig) Translate() (output HADeviceBinarySensorFunctions) {
	AddStateFunction(entity.Availability.State, &output.Availability.State)
	AddStateFunction(entity.State, &output.State)
	return
}

func (entity HADeviceCameraFunctionsConfig) Translate() (output HADeviceCameraFunctions) {
	AddStateFunction(entity.Availability.State, &output.Availability.State)
	AddStateFunction(entity.State, &output.State)
	return
}

func (entity HADeviceCoverFunctionsConfig) Translate() (output HADeviceCoverFunctions) {
	AddStateFunction(entity.Availability.State, &output.Availability.State)
	AddCommandFunction(entity.Command, &output.Command)
	AddCommandFunction(entity.SetPosition, &output.SetPosition)
	AddStateFunction(entity.State, &output.State)
	AddCommandFunction(entity.TiltCommand, &output.TiltCommand)
	AddStateFunction(entity.TiltStatus, &output.TiltStatus)
	return
}

func (entity HADeviceDeviceTrackerFunctionsConfig) Translate() (output HADeviceDeviceTrackerFunctions) {
	return
}

func (entity HADeviceDeviceTriggerFunctionsConfig) Translate() (output HADeviceDeviceTriggerFunctions) {
	AddStateFunction(entity.State, &output.State)
	return
}

func (entity HADeviceFanFunctionsConfig) Translate() (output HADeviceFanFunctions) {
	AddStateFunction(entity.Availability.State, &output.Availability.State)
	AddCommandFunction(entity.Command, &output.Command)
	AddCommandFunction(entity.OscillationCommand, &output.OscillationCommand)
	AddStateFunction(entity.OscillationState, &output.OscillationState)
	AddCommandFunction(entity.PercentageCommand, &output.PercentageCommand)
	AddStateFunction(entity.PercentageState, &output.PercentageState)
	AddCommandFunction(entity.PresetModeCommand, &output.PresetModeCommand)
	AddStateFunction(entity.PresetModeState, &output.PresetModeState)
	AddStateFunction(entity.State, &output.State)
	return
}

func (entity HADeviceHumidifierFunctionsConfig) Translate() (output HADeviceHumidifierFunctions) {
	AddStateFunction(entity.Availability.State, &output.Availability.State)
	AddCommandFunction(entity.Command, &output.Command)
	AddCommandFunction(entity.ModeCommand, &output.ModeCommand)
	AddStateFunction(entity.ModeState, &output.ModeState)
	AddStateFunction(entity.State, &output.State)
	AddCommandFunction(entity.TargetHumidityCommand, &output.TargetHumidityCommand)
	AddStateFunction(entity.TargetHumidityState, &output.TargetHumidityState)
	return
}

func (entity HADeviceClimateFunctionsConfig) Translate() (output HADeviceClimateFunctions) {
	AddCommandFunction(entity.AuxCommand, &output.AuxCommand)
	AddStateFunction(entity.AuxState, &output.AuxState)
	AddStateFunction(entity.Availability.State, &output.Availability.State)
	AddCommandFunction(entity.AwayModeCommand, &output.AwayModeCommand)
	AddStateFunction(entity.AwayModeState, &output.AwayModeState)
	AddCommandFunction(entity.FanModeCommand, &output.FanModeCommand)
	AddStateFunction(entity.FanModeState, &output.FanModeState)
	AddCommandFunction(entity.HoldCommand, &output.HoldCommand)
	AddStateFunction(entity.HoldState, &output.HoldState)
	AddCommandFunction(entity.ModeCommand, &output.ModeCommand)
	AddStateFunction(entity.ModeState, &output.ModeState)
	AddCommandFunction(entity.PowerCommand, &output.PowerCommand)
	AddCommandFunction(entity.SwingModeCommand, &output.SwingModeCommand)
	AddStateFunction(entity.SwingModeState, &output.SwingModeState)
	AddCommandFunction(entity.TemperatureCommand, &output.TemperatureCommand)
	AddCommandFunction(entity.TemperatureHighCommand, &output.TemperatureHighCommand)
	AddStateFunction(entity.TemperatureHighState, &output.TemperatureHighState)
	AddCommandFunction(entity.TemperatureLowCommand, &output.TemperatureLowCommand)
	AddStateFunction(entity.TemperatureLowState, &output.TemperatureLowState)
	AddStateFunction(entity.TemperatureState, &output.TemperatureState)
	return
}

func (entity HADeviceLightFunctionsConfig) Translate() (output HADeviceLightFunctions) {
	AddStateFunction(entity.Availability.State, &output.Availability.State)
	AddCommandFunction(entity.BrightnessCommand, &output.BrightnessCommand)
	AddStateFunction(entity.BrightnessState, &output.BrightnessState)
	AddStateFunction(entity.ColorModeState, &output.ColorModeState)
	AddCommandFunction(entity.ColorTempCommand, &output.ColorTempCommand)
	AddStateFunction(entity.ColorTempState, &output.ColorTempState)
	AddCommandFunction(entity.Command, &output.Command)
	AddCommandFunction(entity.EffectCommand, &output.EffectCommand)
	AddStateFunction(entity.EffectState, &output.EffectState)
	AddCommandFunction(entity.HsCommand, &output.HsCommand)
	AddStateFunction(entity.HsState, &output.HsState)
	AddCommandFunction(entity.RgbCommand, &output.RgbCommand)
	AddStateFunction(entity.RgbState, &output.RgbState)
	AddStateFunction(entity.State, &output.State)
	AddCommandFunction(entity.WhiteCommand, &output.WhiteCommand)
	AddCommandFunction(entity.XyCommand, &output.XyCommand)
	AddStateFunction(entity.XyState, &output.XyState)
	return
}

func (entity HADeviceLockFunctionsConfig) Translate() (output HADeviceLockFunctions) {
	AddStateFunction(entity.Availability.State, &output.Availability.State)
	AddCommandFunction(entity.Command, &output.Command)
	AddStateFunction(entity.State, &output.State)
	return
}

func (entity HADeviceNumberFunctionsConfig) Translate() (output HADeviceNumberFunctions) {
	AddStateFunction(entity.Availability.State, &output.Availability.State)
	AddCommandFunction(entity.Command, &output.Command)
	AddStateFunction(entity.State, &output.State)
	return
}

func (entity HADeviceSceneFunctionsConfig) Translate() (output HADeviceSceneFunctions) {
	AddStateFunction(entity.Availability.State, &output.Availability.State)
	AddCommandFunction(entity.Command, &output.Command)
	return
}

func (entity HADeviceSelectFunctionsConfig) Translate() (output HADeviceSelectFunctions) {
	AddStateFunction(entity.Availability.State, &output.Availability.State)
	AddCommandFunction(entity.Command, &output.Command)
	AddStateFunction(entity.State, &output.State)
	return
}

func (entity HADeviceSensorFunctionsConfig) Translate() (output HADeviceSensorFunctions) {
	AddStateFunction(entity.Availability.State, &output.Availability.State)
	AddStateFunction(entity.State, &output.State)
	return
}

func (entity HADeviceSwitchFunctionsConfig) Translate() (output HADeviceSwitchFunctions) {
	AddStateFunction(entity.Availability.State, &output.Availability.State)
	AddCommandFunction(entity.Command, &output.Command)
	AddStateFunction(entity.State, &output.State)
	return
}

func (entity HADeviceTagFunctionsConfig) Translate() (output HADeviceTagFunctions) {
	AddStateFunction(entity.State, &output.State)
	return
}

func (entity HADeviceVacuumFunctionsConfig) Translate() (output HADeviceVacuumFunctions) {
	AddStateFunction(entity.Availability.State, &output.Availability.State)
	AddCommandFunction(entity.Command, &output.Command)
	AddCommandFunction(entity.SendCommand, &output.SendCommand)
	AddCommandFunction(entity.SetFanSpeed, &output.SetFanSpeed)
	return
}
