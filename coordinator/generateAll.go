package coordinator

import (
	"errors"
	"github.com/fatih/structs"
	"strconv"
)

// GenerateAll iterates through all types of devices, as declared in the config,
// and runs the Generate() function on each of them. This fills in any gaps that
// these devices have (topics, for example), and raises errors if required.
func GenerateAll() error {
	n := 0
	if len(Config.Devices.AlarmControlPanel) > 0 {
		for k := range Config.Devices.AlarmControlPanel {
			if structs.IsStruct(Config.Devices.AlarmControlPanel[k].Configuration) {
				Config.Devices.AlarmControlPanel[k].Configuration.Generate()
				n += 1
			} else {
				return errors.New("device Config.Devices.AlarmControlPanel[" + strconv.Itoa(k) + "].Configuration does not exist")
			}
		}
	}
	if len(Config.Devices.BinarySensor) > 0 {
		for k := range Config.Devices.BinarySensor {
			if structs.IsStruct(Config.Devices.BinarySensor[k].Configuration) {
				Config.Devices.BinarySensor[k].Configuration.Generate()
				n += 1
			} else {
				return errors.New("device Config.Devices.BinarySensor[" + strconv.Itoa(k) + "].Configuration does not exist")
			}
		}
	}
	if len(Config.Devices.Camera) > 0 {
		for k := range Config.Devices.Camera {
			if structs.IsStruct(Config.Devices.Camera[k].Configuration) {
				Config.Devices.Camera[k].Configuration.Generate()
				n += 1
			} else {
				return errors.New("device Config.Devices.Camera[" + strconv.Itoa(k) + "].Configuration does not exist")
			}
		}
	}
	if len(Config.Devices.Cover) > 0 {
		for k := range Config.Devices.Cover {
			if structs.IsStruct(Config.Devices.Cover[k].Configuration) {
				Config.Devices.Cover[k].Configuration.Generate()
				n += 1
			} else {
				return errors.New("device Config.Devices.Cover[" + strconv.Itoa(k) + "].Configuration does not exist")
			}
		}
	}
	if len(Config.Devices.DeviceTracker) > 0 {
		for k := range Config.Devices.DeviceTracker {
			if structs.IsStruct(Config.Devices.DeviceTracker[k].Configuration) {
				Config.Devices.DeviceTracker[k].Configuration.Generate()
				n += 1
			} else {
				return errors.New("device Config.Devices.DeviceTracker[" + strconv.Itoa(k) + "].Configuration does not exist")
			}
		}
	}
	if len(Config.Devices.DeviceTrigger) > 0 {
		for k := range Config.Devices.DeviceTrigger {
			if structs.IsStruct(Config.Devices.DeviceTrigger[k].Configuration) {
				Config.Devices.DeviceTrigger[k].Configuration.Generate()
				n += 1
			} else {
				return errors.New("device Config.Devices.DeviceTrigger[" + strconv.Itoa(k) + "].Configuration does not exist")
			}
		}
	}
	if len(Config.Devices.Fan) > 0 {
		for k := range Config.Devices.Fan {
			if structs.IsStruct(Config.Devices.Fan[k].Configuration) {
				Config.Devices.Fan[k].Configuration.Generate()
				n += 1
			} else {
				return errors.New("device Config.Devices.Fan[" + strconv.Itoa(k) + "].Configuration does not exist")
			}
		}
	}
	if len(Config.Devices.Humidifier) > 0 {
		for k := range Config.Devices.Humidifier {
			if structs.IsStruct(Config.Devices.Humidifier[k].Configuration) {
				Config.Devices.Humidifier[k].Configuration.Generate()
				n += 1
			} else {
				return errors.New("device Config.Devices.Humidifier[" + strconv.Itoa(k) + "].Configuration does not exist")
			}
		}
	}
	if len(Config.Devices.Climate) > 0 {
		for k := range Config.Devices.Climate {
			if structs.IsStruct(Config.Devices.Climate[k].Configuration) {
				Config.Devices.Climate[k].Configuration.Generate()
				n += 1
			} else {
				return errors.New("device Config.Devices.Climate[" + strconv.Itoa(k) + "].Configuration does not exist")
			}
		}
	}
	if len(Config.Devices.Light) > 0 {
		for k := range Config.Devices.Light {
			if structs.IsStruct(Config.Devices.Light[k].Configuration) {
				Config.Devices.Light[k].Configuration.Generate()
				n += 1
			} else {
				return errors.New("device Config.Devices.Light[" + strconv.Itoa(k) + "].Configuration does not exist")
			}
		}
	}
	if len(Config.Devices.Lock) > 0 {
		for k := range Config.Devices.Lock {
			if structs.IsStruct(Config.Devices.Lock[k].Configuration) {
				Config.Devices.Lock[k].Configuration.Generate()
				n += 1
			} else {
				return errors.New("device Config.Devices.Lock[" + strconv.Itoa(k) + "].Configuration does not exist")
			}
		}
	}
	if len(Config.Devices.Number) > 0 {
		for k := range Config.Devices.Number {
			if structs.IsStruct(Config.Devices.Number[k].Configuration) {
				Config.Devices.Number[k].Configuration.Generate()
				n += 1
			} else {
				return errors.New("device Config.Devices.Number[" + strconv.Itoa(k) + "].Configuration does not exist")
			}
		}
	}
	if len(Config.Devices.Scene) > 0 {
		for k := range Config.Devices.Scene {
			if structs.IsStruct(Config.Devices.Scene[k].Configuration) {
				Config.Devices.Scene[k].Configuration.Generate()
				n += 1
			} else {
				return errors.New("device Config.Devices.Scene[" + strconv.Itoa(k) + "].Configuration does not exist")
			}
		}
	}
	if len(Config.Devices.Select) > 0 {
		for k := range Config.Devices.Select {
			if structs.IsStruct(Config.Devices.Select[k].Configuration) {
				Config.Devices.Select[k].Configuration.Generate()
				n += 1
			} else {
				return errors.New("device Config.Devices.Select[" + strconv.Itoa(k) + "].Configuration does not exist")
			}
		}
	}
	if len(Config.Devices.Sensor) > 0 {
		for k := range Config.Devices.Sensor {
			if structs.IsStruct(Config.Devices.Sensor[k].Configuration) {
				Config.Devices.Sensor[k].Configuration.Generate()
				n += 1
			} else {
				return errors.New("device Config.Devices.Sensor[" + strconv.Itoa(k) + "].Configuration does not exist")
			}
		}
	}
	if len(Config.Devices.Switch) > 0 {
		for k := range Config.Devices.Switch {
			if structs.IsStruct(Config.Devices.Switch[k].Configuration) {
				Config.Devices.Switch[k].Configuration.Generate()
				n += 1
			} else {
				return errors.New("device Config.Devices.Switch[" + strconv.Itoa(k) + "].Configuration does not exist")
			}
		}
	}
	if len(Config.Devices.Tag) > 0 {
		for k := range Config.Devices.Tag {
			if structs.IsStruct(Config.Devices.Tag[k].Configuration) {
				Config.Devices.Tag[k].Configuration.Generate()
				n += 1
			} else {
				return errors.New("device Config.Devices.Tag[" + strconv.Itoa(k) + "].Configuration does not exist")
			}
		}
	}
	if len(Config.Devices.Vacuum) > 0 {
		for k := range Config.Devices.Vacuum {
			if structs.IsStruct(Config.Devices.Vacuum[k].Configuration) {
				Config.Devices.Vacuum[k].Configuration.Generate()
				n += 1
			} else {
				return errors.New("device Config.Devices.Vacuum[" + strconv.Itoa(k) + "].Configuration does not exist")
			}
		}
	}
	if n == 0 {
		return errors.New("no devices generated.")
	}
	return nil
}
