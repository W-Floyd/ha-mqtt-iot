package battery

import (
	"strconv"

	"../../hadiscovery"
	"../../iotconfig/config"
)

var Sconfig = config.Config{}

func (batteries Batteries) Translate() (batteriesOutput []hadiscovery.Sensor, batteriesConnectedOutput []hadiscovery.BinarySensor) {

	for k, battery := range batteries {

		bSensor := hadiscovery.Sensor{}
		if len(batteries) > 1 {
			bSensor.UniqueID = "builtin-battery-" + strconv.Itoa(k) + "_" + hadiscovery.NodeID
			bSensor.Name = Sconfig.Builtin.Prefix + "Battery " + strconv.Itoa(k)
		} else {
			bSensor.UniqueID = "builtin-battery_" + hadiscovery.NodeID
			bSensor.Name = Sconfig.Builtin.Prefix + "Battery"
		}

		bSensor.StateFunc = func() string {
			val, _ := strconv.Atoi(battery.GetCharge())
			return strconv.FormatFloat(float64(val*100)/float64(battery.MaxCapacity), 'f', 1, 32)
		}
		bSensor.UnitOfMeasurement = "%"
		bSensor.UpdateInterval = 10
		bSensor.Icon = "mdi:battery"

		bSensor.Initialize()

		batteriesOutput = append(batteriesOutput, bSensor)

		bBSensor := hadiscovery.BinarySensor{}
		if len(batteries) > 1 {
			bBSensor.UniqueID = "builtin-battery-" + strconv.Itoa(k) + "-plugged-in_" + hadiscovery.NodeID
			bBSensor.Name = Sconfig.Builtin.Prefix + "Battery " + strconv.Itoa(k) + " Plugged In"
		} else {
			bBSensor.UniqueID = "builtin-battery-plugged-in_" + hadiscovery.NodeID
			bBSensor.Name = Sconfig.Builtin.Prefix + "Battery Plugged In"
		}
		bBSensor.StateFunc = func() string {
			if battery.IsPluggedIn() {
				return "ON"
			}
			return "OFF"
		}
		bBSensor.UpdateInterval = 1
		bBSensor.DeviceClass = "battery_charging"

		bBSensor.Initialize()

		batteriesConnectedOutput = append(batteriesConnectedOutput, bBSensor)

	}

	return

}
