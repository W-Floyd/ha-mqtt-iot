package batterywindows

import (
	"github.com/W-Floyd/ha-mqtt-iot/hadiscovery"
	"github.com/W-Floyd/ha-mqtt-iot/iotconfig/config"
)

var Sconfig = config.Config{}

func Create() (bSensor hadiscovery.Sensor, bBSensor hadiscovery.BinarySensor) {

	bSensor.UniqueID = "builtin-battery_" + hadiscovery.NodeID
	bSensor.Name = Sconfig.Builtin.Prefix + "Battery"

	bSensor.StateFunc = GetLevel
	bSensor.UnitOfMeasurement = "%"
	bSensor.UpdateInterval = 10
	bSensor.Icon = "mdi:battery"

	bSensor.Initialize()

	bBSensor.UniqueID = "builtin-battery-plugged-in_" + hadiscovery.NodeID
	bBSensor.Name = Sconfig.Builtin.Prefix + "Battery Plugged In"
	bBSensor.StateFunc = func() string {
		if IsPluggedIn() {
			return "ON"
		}
		return "OFF"
	}
	bBSensor.UpdateInterval = 5
	bBSensor.DeviceClass = "battery_charging"

	bBSensor.Initialize()

	return

}
