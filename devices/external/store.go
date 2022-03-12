package ExternalDevice

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
type store struct {
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
	}
	Cover struct {
		Availability map[string]string
		Position     map[string]string
		State        map[string]string
		TiltStatus   map[string]string
	}
	DeviceTracker struct{}
	DeviceTrigger struct{}
	Fan           struct {
		Availability     map[string]string
		OscillationState map[string]string
		PercentageState  map[string]string
		PresetModeState  map[string]string
		State            map[string]string
	}
	Humidifier struct {
		Availability        map[string]string
		ModeState           map[string]string
		State               map[string]string
		TargetHumidityState map[string]string
	}
	Climate struct {
		AuxState             map[string]string
		Availability         map[string]string
		CurrentTemperature   map[string]string
		FanModeState         map[string]string
		ModeState            map[string]string
		PresetModeState      map[string]string
		SwingModeState       map[string]string
		TemperatureHighState map[string]string
		TemperatureLowState  map[string]string
		TemperatureState     map[string]string
	}
	Light struct {
		Availability map[string]string
		State        map[string]string
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
	Tag    struct{}
	Vacuum struct {
		Availability map[string]string
		State        map[string]string
	}
}

func initStore() store {
	s := store{}
	s.AlarmControlPanel.Availability = make(map[string]string)
	s.AlarmControlPanel.State = make(map[string]string)
	s.BinarySensor.Availability = make(map[string]string)
	s.BinarySensor.State = make(map[string]string)
	s.Button.Availability = make(map[string]string)
	s.Camera.Availability = make(map[string]string)
	s.Cover.Availability = make(map[string]string)
	s.Cover.Position = make(map[string]string)
	s.Cover.State = make(map[string]string)
	s.Cover.TiltStatus = make(map[string]string)
	s.Fan.Availability = make(map[string]string)
	s.Fan.OscillationState = make(map[string]string)
	s.Fan.PercentageState = make(map[string]string)
	s.Fan.PresetModeState = make(map[string]string)
	s.Fan.State = make(map[string]string)
	s.Humidifier.Availability = make(map[string]string)
	s.Humidifier.ModeState = make(map[string]string)
	s.Humidifier.State = make(map[string]string)
	s.Humidifier.TargetHumidityState = make(map[string]string)
	s.Climate.AuxState = make(map[string]string)
	s.Climate.Availability = make(map[string]string)
	s.Climate.CurrentTemperature = make(map[string]string)
	s.Climate.FanModeState = make(map[string]string)
	s.Climate.ModeState = make(map[string]string)
	s.Climate.PresetModeState = make(map[string]string)
	s.Climate.SwingModeState = make(map[string]string)
	s.Climate.TemperatureHighState = make(map[string]string)
	s.Climate.TemperatureLowState = make(map[string]string)
	s.Climate.TemperatureState = make(map[string]string)
	s.Light.Availability = make(map[string]string)
	s.Light.State = make(map[string]string)
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
	s.Vacuum.Availability = make(map[string]string)
	s.Vacuum.State = make(map[string]string)
	return s
}
