package iotconfig

import (
	"log"
	"math"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	backlightP "../builtin/backlight"
	batteryP "../builtin/battery"
	"../builtin/batterywindows"
	"../hadiscovery"
)

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

type InfoIcon struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Icon string `json:"icon"`
}

type InfoClass struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	DeviceClass string `json:"device_class"`
}

type SensorHA struct {
	Info              InfoIcon `json:"info"`
	CommandState      []string `json:"command_state"`
	UnitOfMeasurement string   `json:"unit_of_measurement,omitempty"`
	UpdateInterval    float64  `json:"update_interval"`
	ForceUpdateMQTT   bool     `json:"force_update"`
}

type BinarySensorsHA struct {
	Info            InfoClass `json:"info"`
	CommandState    []string  `json:"command_state"`
	UpdateInterval  float64   `json:"update_interval"`
	ForceUpdateMQTT bool      `json:"force_update"`
}

type Config struct {
	MQTT struct {
		Broker       string `json:"broker"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		NodeID       string `json:"node_id"`
		InstanceName string `json:"instance_name"`
	} `json:"mqtt"`
	Builtin struct {
		Prefix    string `json:"prefix"`
		Backlight struct {
			Enable      bool `json:"enable"`
			Temperature bool `json:"temperature"`
		} `json:"backlight"`
		Battery struct {
			Enable bool `json:"enable"`
		} `json:"battery"`
	} `json:"builtin"`
	Lights        []LightHA         `json:"lights"`
	Switches      []SwitchHA        `json:"switches"`
	Sensors       []SensorHA        `json:"sensors"`
	BinarySensors []BinarySensorsHA `json:"binary_sensors"`
}

func constructCommandFunc(command []string) (f func(message mqtt.Message, connection mqtt.Client)) {
	var err error
	return func(message mqtt.Message, connection mqtt.Client) {
		localcom := command
		localcom = append(localcom, string(message.Payload()))
		if len(command) > 0 {

			if len(command) > 1 {
				_, err = exec.Command(localcom[0], localcom[1:]...).Output()
			} else {
				_, err = exec.Command(localcom[0]).Output()
			}
			if err != nil {
				log.Printf("%s", err)
			}

		}
	}
}

func constructStateFunc(command []string) (f func() string) {
	var err error
	return func() string {
		var out []byte
		if len(command) > 1 {
			out, err = exec.Command(command[0], command[1:]...).Output()
		} else {

			out, err = exec.Command(command[0]).Output()
		}
		if err != nil {
			log.Printf("%s", err)
		}
		return string(out)
	}
}

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

		backlightP.sconfig = sconfig

		if runtime.GOOS != "linux" {
			log.Fatalln("Backlight not supported on non-linux platforms")
		} else {

			backlights := backlightP.PopulateBacklights()
			lights = append(lights, backlights.Translate())

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
