package ExternalDevice

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type StateStore struct {
	AlarmControlPanel struct {
		Availability map[string]string
		State        map[string]string
	}
	BinarySensor struct {
		Availability map[string]string
		State        map[string]string
	}
	Button struct {
		Availability map[string]string
	}
	Camera struct {
		Availability map[string]string
		State        map[string]string
	}
	Climate struct {
		AuxState             map[string]string
		Availability         map[string]string
		CurrentHumidity      map[string]string
		CurrentTemperature   map[string]string
		FanModeState         map[string]string
		ModeState            map[string]string
		PresetModeState      map[string]string
		SwingModeState       map[string]string
		TargetHumidityState  map[string]string
		TemperatureHighState map[string]string
		TemperatureLowState  map[string]string
		TemperatureState     map[string]string
	}
	Cover struct {
		Availability map[string]string
		Position     map[string]string
		State        map[string]string
		TiltStatus   map[string]string
	}
	DeviceTracker struct {
		Availability map[string]string
		State        map[string]string
	}
	DeviceTrigger struct {
		State map[string]string
	}
	Fan struct {
		Availability     map[string]string
		DirectionState   map[string]string
		OscillationState map[string]string
		PercentageState  map[string]string
		PresetModeState  map[string]string
		State            map[string]string
	}
	Humidifier struct {
		Availability        map[string]string
		CurrentHumidity     map[string]string
		ModeState           map[string]string
		State               map[string]string
		TargetHumidityState map[string]string
	}
	Image struct {
		Availability map[string]string
	}
	Light struct {
		Availability    map[string]string
		BrightnessState map[string]string
		ColorModeState  map[string]string
		ColorTempState  map[string]string
		EffectState     map[string]string
		HsState         map[string]string
		RgbState        map[string]string
		RgbwState       map[string]string
		RgbwwState      map[string]string
		State           map[string]string
		XyState         map[string]string
	}
	Lock struct {
		Availability map[string]string
		State        map[string]string
	}
	Number struct {
		Availability map[string]string
		State        map[string]string
	}
	Scene struct {
		Availability map[string]string
	}
	Select struct {
		Availability map[string]string
		State        map[string]string
	}
	Sensor struct {
		Availability map[string]string
		State        map[string]string
	}
	Siren struct {
		Availability map[string]string
		State        map[string]string
	}
	Switch struct {
		Availability map[string]string
		State        map[string]string
	}
	Tag struct {
		State map[string]string
	}
	Text struct {
		Availability map[string]string
		State        map[string]string
	}
	Update struct {
		Availability map[string]string
		State        map[string]string
	}
	Vacuum struct {
		Availability map[string]string
		State        map[string]string
	}
	WaterHeater struct {
		Availability       map[string]string
		CurrentTemperature map[string]string
		ModeState          map[string]string
		TemperatureState   map[string]string
	}
}

func initStore() StateStore {
	s := StateStore{}
	s.AlarmControlPanel.Availability = make(map[string]string)
	s.AlarmControlPanel.State = make(map[string]string)
	s.BinarySensor.Availability = make(map[string]string)
	s.BinarySensor.State = make(map[string]string)
	s.Button.Availability = make(map[string]string)
	s.Camera.Availability = make(map[string]string)
	s.Camera.State = make(map[string]string)
	s.Climate.AuxState = make(map[string]string)
	s.Climate.Availability = make(map[string]string)
	s.Climate.CurrentHumidity = make(map[string]string)
	s.Climate.CurrentTemperature = make(map[string]string)
	s.Climate.FanModeState = make(map[string]string)
	s.Climate.ModeState = make(map[string]string)
	s.Climate.PresetModeState = make(map[string]string)
	s.Climate.SwingModeState = make(map[string]string)
	s.Climate.TargetHumidityState = make(map[string]string)
	s.Climate.TemperatureHighState = make(map[string]string)
	s.Climate.TemperatureLowState = make(map[string]string)
	s.Climate.TemperatureState = make(map[string]string)
	s.Cover.Availability = make(map[string]string)
	s.Cover.Position = make(map[string]string)
	s.Cover.State = make(map[string]string)
	s.Cover.TiltStatus = make(map[string]string)
	s.DeviceTracker.Availability = make(map[string]string)
	s.DeviceTracker.State = make(map[string]string)
	s.DeviceTrigger.State = make(map[string]string)
	s.Fan.Availability = make(map[string]string)
	s.Fan.DirectionState = make(map[string]string)
	s.Fan.OscillationState = make(map[string]string)
	s.Fan.PercentageState = make(map[string]string)
	s.Fan.PresetModeState = make(map[string]string)
	s.Fan.State = make(map[string]string)
	s.Humidifier.Availability = make(map[string]string)
	s.Humidifier.CurrentHumidity = make(map[string]string)
	s.Humidifier.ModeState = make(map[string]string)
	s.Humidifier.State = make(map[string]string)
	s.Humidifier.TargetHumidityState = make(map[string]string)
	s.Image.Availability = make(map[string]string)
	s.Light.Availability = make(map[string]string)
	s.Light.BrightnessState = make(map[string]string)
	s.Light.ColorModeState = make(map[string]string)
	s.Light.ColorTempState = make(map[string]string)
	s.Light.EffectState = make(map[string]string)
	s.Light.HsState = make(map[string]string)
	s.Light.RgbState = make(map[string]string)
	s.Light.RgbwState = make(map[string]string)
	s.Light.RgbwwState = make(map[string]string)
	s.Light.State = make(map[string]string)
	s.Light.XyState = make(map[string]string)
	s.Lock.Availability = make(map[string]string)
	s.Lock.State = make(map[string]string)
	s.Number.Availability = make(map[string]string)
	s.Number.State = make(map[string]string)
	s.Scene.Availability = make(map[string]string)
	s.Select.Availability = make(map[string]string)
	s.Select.State = make(map[string]string)
	s.Sensor.Availability = make(map[string]string)
	s.Sensor.State = make(map[string]string)
	s.Siren.Availability = make(map[string]string)
	s.Siren.State = make(map[string]string)
	s.Switch.Availability = make(map[string]string)
	s.Switch.State = make(map[string]string)
	s.Tag.State = make(map[string]string)
	s.Text.Availability = make(map[string]string)
	s.Text.State = make(map[string]string)
	s.Update.Availability = make(map[string]string)
	s.Update.State = make(map[string]string)
	s.Vacuum.Availability = make(map[string]string)
	s.Vacuum.State = make(map[string]string)
	s.WaterHeater.Availability = make(map[string]string)
	s.WaterHeater.CurrentTemperature = make(map[string]string)
	s.WaterHeater.ModeState = make(map[string]string)
	s.WaterHeater.TemperatureState = make(map[string]string)
	return s
}
