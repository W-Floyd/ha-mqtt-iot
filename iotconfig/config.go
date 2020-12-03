package iotconfig

import (
	"log"
	"runtime"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	backlightP "../builtin/backlight"
	batteryP "../builtin/battery"
	"../builtin/batterywindows"
	"../hadiscovery"
	"./config"
)

type Config config.Config

// Convert takes a config and turns it into a format that can be used for HA.
func (sconfig Config) Convert() (opts *mqtt.ClientOptions, switches []hadiscovery.Switch, sensors []hadiscovery.Sensor, binarySensors []hadiscovery.BinarySensor, lights []hadiscovery.Light) {
	opts = mqtt.NewClientOptions()
	opts.AddBroker(sconfig.MQTT.Broker)
	opts.SetUsername(sconfig.MQTT.Username)
	opts.SetPassword(sconfig.MQTT.Password)
	opts.SetKeepAlive(30 * time.Second)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("TOPIC: %s\n", msg.Topic())
		log.Printf("MSG: %s\n", msg.Payload())
	})
	opts.SetPingTimeout(1 * time.Second)
	opts.SetAutoReconnect(true)
	if sconfig.MQTT.NodeID != "" {
		hadiscovery.NodeID = sconfig.MQTT.NodeID
	}
	opts.SetClientID(hadiscovery.NodeID)
	if sconfig.MQTT.InstanceName != "" {
		hadiscovery.InstanceName = sconfig.MQTT.InstanceName
	}
	for _, sw := range sconfig.Switches {
		switches = append(switches, sw.Translate())
	}
	for _, se := range sconfig.Sensors {
		sensors = append(sensors, se.Translate())
	}
	for _, bse := range sconfig.BinarySensors {
		binarySensors = append(binarySensors, bse.Translate())
	}

	for _, li := range sconfig.Lights {
		lights = append(lights, li.Translate())
	}

	if sconfig.Builtin.Backlight.Enable {

		backlightP.Sconfig = config.Config(sconfig)

		if runtime.GOOS != "linux" {
			log.Fatalln("Backlight not supported on non-linux platforms")
		} else {

			backlights := backlightP.Backlights(backlightP.PopulateBacklights())
			lights = append(lights, backlights.Translate()...)

		}

	}

	if sconfig.Builtin.Battery.Enable {

		if runtime.GOOS == "linux" {

			batteries := batteryP.PopulateBatteries()

			for k, battery := range batteries {

				bSensor := hadiscovery.Sensor{}
				if len(batteries) > 1 {
					bSensor.UniqueID = "builtin-battery-" + strconv.Itoa(k) + "_" + hadiscovery.NodeID
					bSensor.Name = sconfig.Builtin.Prefix + "Battery " + strconv.Itoa(k)
				} else {
					bSensor.UniqueID = "builtin-battery_" + hadiscovery.NodeID
					bSensor.Name = sconfig.Builtin.Prefix + "Battery"
				}

				bSensor.StateFunc = func() string {
					val, _ := strconv.Atoi(battery.GetCharge())
					return strconv.FormatFloat(float64(val*100)/float64(battery.MaxCapacity), 'f', 1, 32)
				}
				bSensor.UnitOfMeasurement = "%"
				bSensor.UpdateInterval = 10
				bSensor.Icon = "mdi:battery"

				bSensor.Initialize()

				sensors = append(sensors, bSensor)

				bBSensor := hadiscovery.BinarySensor{}
				if len(batteries) > 1 {
					bBSensor.UniqueID = "builtin-battery-" + strconv.Itoa(k) + "-plugged-in_" + hadiscovery.NodeID
					bBSensor.Name = sconfig.Builtin.Prefix + "Battery " + strconv.Itoa(k) + " Plugged In"
				} else {
					bBSensor.UniqueID = "builtin-battery-plugged-in_" + hadiscovery.NodeID
					bBSensor.Name = sconfig.Builtin.Prefix + "Battery Plugged In"
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

				binarySensors = append(binarySensors, bBSensor)

			}
		} else if runtime.GOOS == "windows" {

			bSensor := hadiscovery.Sensor{}

			bSensor.UniqueID = "builtin-battery_" + hadiscovery.NodeID
			bSensor.Name = sconfig.Builtin.Prefix + "Battery"

			bSensor.StateFunc = batterywindows.GetLevel
			bSensor.UnitOfMeasurement = "%"
			bSensor.UpdateInterval = 10
			bSensor.Icon = "mdi:battery"

			bSensor.Initialize()

			sensors = append(sensors, bSensor)

			bBSensor := hadiscovery.BinarySensor{}
			bBSensor.UniqueID = "builtin-battery-plugged-in_" + hadiscovery.NodeID
			bBSensor.Name = sconfig.Builtin.Prefix + "Battery Plugged In"
			bBSensor.StateFunc = func() string {
				if batterywindows.IsPluggedIn() {
					return "ON"
				}
				return "OFF"
			}
			bBSensor.UpdateInterval = 5
			bBSensor.DeviceClass = "battery_charging"

			bBSensor.Initialize()

			binarySensors = append(binarySensors, bBSensor)

		}

	} else {
		log.Fatalln("Battery not supported on this platform")
	}

	return
}
