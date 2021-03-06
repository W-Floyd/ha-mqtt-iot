package iotconfig

import (
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	backlightP "github.com/W-Floyd/ha-mqtt-iot/builtin/backlight"
	batteryP "github.com/W-Floyd/ha-mqtt-iot/builtin/battery"
	"github.com/W-Floyd/ha-mqtt-iot/builtin/batterywindows"
	"github.com/W-Floyd/ha-mqtt-iot/builtin/crypto"
	"github.com/W-Floyd/ha-mqtt-iot/hadiscovery"
	"github.com/W-Floyd/ha-mqtt-iot/iotconfig/common"
	"github.com/W-Floyd/ha-mqtt-iot/iotconfig/config"
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

		if common.AlmostEqual(sconfig.Builtin.Backlight.Range.Maximum, 0.0) {
			sconfig.Builtin.Backlight.Range.Maximum = 1.0
		}

		if common.AlmostEqual(sconfig.Builtin.Backlight.Range.Minimum, 0.0) {
			sconfig.Builtin.Backlight.Range.Minimum = 0.0
		}

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

			batteryP.Sconfig = config.Config(sconfig)

			batterySensors, batteryBinarySensors := batteryP.PopulateBatteries().Translate()

			sensors = append(sensors, batterySensors...)
			binarySensors = append(binarySensors, batteryBinarySensors...)

		} else if runtime.GOOS == "windows" {

			batterywindows.Sconfig = config.Config(sconfig)

			bs, bbs := batterywindows.Create()

			sensors = append(sensors, bs)
			binarySensors = append(binarySensors, bbs)

		} else {
			log.Fatalln("Battery not supported on this platform")
		}
	}

	if len(sconfig.Builtin.Crypto) > 0 {

		crypto.Setup()

		for k, entry := range sconfig.Builtin.Crypto {
			cSensor := hadiscovery.Sensor{}
			cSensor.UniqueID = "builtin-crypto-" + entry.CoinName + "-" + entry.CurrencyName
			cSensor.Name = strings.ToUpper(entry.CurrencyName) + "/" + strings.Title(entry.CoinName)

			f := func(i int) func() string {
				return func() string {
					localEntry := sconfig.Builtin.Crypto[i]
					val := crypto.GetPrice(localEntry.CoinName, localEntry.CurrencyName)
					return strconv.FormatFloat(float64(val), 'f', 6, 32)
				}
			}(k)

			cSensor.StateFunc = f
			cSensor.UnitOfMeasurement = strings.ToUpper(entry.CurrencyName)
			if !common.AlmostEqual(entry.UpdateInterval, 0.0) {
				cSensor.UpdateInterval = entry.UpdateInterval
			} else {
				entry.UpdateInterval = float64(len(sconfig.Builtin.Crypto))
			}
			cSensor.UpdateInterval = entry.UpdateInterval
			if len(entry.Icon) > 0 {
				cSensor.Icon = entry.Icon
			} else {
				cSensor.Icon = "mdi:currency-usd"
			}
			cSensor.Initialize()
			sensors = append(sensors, cSensor)
		}

	}

	return
}
