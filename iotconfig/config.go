package iotconfig

import (
	"log"
	"math"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	backlightP "../builtin/backlight"
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

type LightHA struct {
	Info                   InfoIcon `json:"info"`
	Command                []string `json:"command"`
	CommandState           []string `json:"command_state"`
	CommandBrightness      []string `json:"command_brightness"`
	CommandBrightnessState []string `json:"command_brightness_state"`
	BrightnessScale        int      `json:"brightness_scale"`
	CommandColorTemp       []string `json:"command_color_temp"`
	CommandColorTempState  []string `json:"command_color_temp_state"`
	CommandEffect          []string `json:"command_effect"`
	CommandEffectState     []string `json:"command_effect_state"`
	CommandHs              []string `json:"command_hs"`
	CommandHsState         []string `json:"command_hs_state"`
	CommandRgb             []string `json:"command_rgb"`
	CommandRgbState        []string `json:"command_rgb_state"`
	CommandWhiteValue      []string `json:"command_white_value"`
	CommandWhiteValueState []string `json:"command_white_value_state"`
	CommandXy              []string `json:"command_xy"`
	CommandXyState         []string `json:"command_xy_state"`
	UpdateInterval         float64  `json:"update_interval"`
	ForceUpdateMQTT        bool     `json:"force_update"`
}

type SwitchHA struct {
	Info            InfoIcon `json:"info"`
	CommandOn       []string `json:"command_on"`
	CommandOff      []string `json:"command_off"`
	CommandState    []string `json:"command_state"`
	UpdateInterval  float64  `json:"update_interval"`
	ForceUpdateMQTT bool     `json:"force_update"`
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
		Backlight struct {
			Enable      bool `json:"enable"`
			Temperature bool `json:"temperature"`
		} `json:"backlight"`
	} `json:"builtin"`
	Lights        []LightHA         `json:"lights"`
	Switches      []SwitchHA        `json:"switches"`
	Sensors       []SensorHA        `json:"sensors"`
	BinarySensors []BinarySensorsHA `json:"binary_sensors"`
}

func (sw SwitchHA) constructCommandFunc() (f func(message mqtt.Message, connection mqtt.Client)) {
	var err error
	return func(message mqtt.Message, connection mqtt.Client) {

		if len(sw.CommandOn) > 0 {
			if string(message.Payload()) == "ON" {

				if len(sw.CommandOn) > 1 {
					_, err = exec.Command(sw.CommandOn[0], sw.CommandOn[1:]...).Output()
				} else {
					_, err = exec.Command(sw.CommandOn[0]).Output()
				}
				if err != nil {
					log.Printf("%s", err)
				}
			}
		}
		if len(sw.CommandOff) > 0 {
			if string(message.Payload()) == "OFF" {

				if len(sw.CommandOff) > 1 {
					_, err = exec.Command(sw.CommandOff[0], sw.CommandOff[1:]...).Output()
				} else {
					_, err = exec.Command(sw.CommandOff[0]).Output()
				}
				if err != nil {
					log.Printf("%s", err)
				}
			}
		}
		if string(message.Payload()) != "OFF" && string(message.Payload()) != "ON" {

			log.Println("Unknown payload: " + string(message.Payload()))

		}
	}
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

func (sw SensorHA) constructStateFunc() (f func() string) {
	var err error
	return func() string {
		var out []byte
		if len(sw.CommandState) > 1 {
			out, err = exec.Command(sw.CommandState[0], sw.CommandState[1:]...).Output()
		} else {

			out, err = exec.Command(sw.CommandState[0]).Output()
		}
		if err != nil {
			log.Printf("%s", err)
		}
		return string(out)
	}
}
func (sw BinarySensorsHA) constructStateFunc() (f func() string) {
	var err error
	return func() string {
		var out []byte
		if len(sw.CommandState) > 1 {
			out, err = exec.Command(sw.CommandState[0], sw.CommandState[1:]...).Output()
		} else {

			out, err = exec.Command(sw.CommandState[0]).Output()
		}
		if err != nil {
			log.Printf("%s", err)
		}
		return string(out)
	}
}
func (sw SwitchHA) constructStateFunc() (f func() string) {
	var err error
	return func() string {
		var out []byte
		if len(sw.CommandState) > 1 {
			out, err = exec.Command(sw.CommandState[0], sw.CommandState[1:]...).Output()
		} else {

			out, err = exec.Command(sw.CommandState[0]).Output()
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
		nsw := hadiscovery.Switch{}
		nsw.UpdateInterval = sw.UpdateInterval
		nsw.ForceUpdateMQTT = sw.ForceUpdateMQTT
		nsw.Name = sw.Info.Name
		nsw.UniqueID = sw.Info.ID + "_" + hadiscovery.NodeID
		if sw.Info.Icon != "" {
			nsw.Icon = sw.Info.Icon
		}

		if len(sw.CommandOn) > 0 || len(sw.CommandOff) > 0 {
			nsw.CommandFunc = sw.constructCommandFunc()
		} else {
			log.Fatalln("Missing command func, needed for switches!")
		}
		if len(sw.CommandState) > 0 {
			nsw.StateFunc = sw.constructStateFunc()
		}

		nsw.Initialize()

		switches = append(switches, nsw)
	}
	for _, se := range sconfig.Sensors {
		nse := hadiscovery.Sensor{}
		nse.UpdateInterval = se.UpdateInterval
		nse.ExpireAfter = int(nse.UpdateInterval + 1)
		nse.ForceUpdateMQTT = se.ForceUpdateMQTT
		nse.UnitOfMeasurement = se.UnitOfMeasurement
		if !se.ForceUpdateMQTT {
			nse.ExpireAfter = 0
		}
		nse.Name = se.Info.Name
		nse.UniqueID = se.Info.ID
		if se.Info.Icon != "" {
			nse.Icon = se.Info.Icon
		}

		if len(se.CommandState) > 0 {
			nse.StateFunc = se.constructStateFunc()
		} else {
			log.Fatalln("Missing state func, needed for sensors!")
		}

		nse.Initialize()

		sensors = append(sensors, nse)
	}
	for _, bse := range sconfig.BinarySensors {
		nse := hadiscovery.BinarySensor{}
		nse.UpdateInterval = bse.UpdateInterval
		nse.ExpireAfter = int(nse.UpdateInterval + 1)
		nse.ForceUpdateMQTT = bse.ForceUpdateMQTT
		if !bse.ForceUpdateMQTT {
			nse.ExpireAfter = 0
		}
		nse.Name = bse.Info.Name
		nse.UniqueID = bse.Info.ID
		if bse.Info.DeviceClass != "" {
			nse.DeviceClass = bse.Info.DeviceClass
		}

		if len(bse.CommandState) > 0 {
			nse.StateFunc = bse.constructStateFunc()
		} else {
			log.Fatalln("Missing state func, needed for binary sensors!")
		}

		nse.Initialize()

		binarySensors = append(binarySensors, nse)
	}

	for _, li := range sconfig.Lights {
		nli := hadiscovery.Light{}
		nli.UpdateInterval = li.UpdateInterval
		nli.ForceUpdateMQTT = li.ForceUpdateMQTT

		nli.Name = li.Info.Name
		nli.UniqueID = li.Info.ID

		if !almostEqual(li.UpdateInterval, 0.0) {
			nli.UpdateInterval = li.UpdateInterval
		} else {
			nli.UpdateInterval = 1
		}

		if li.BrightnessScale != 0 {
			nli.BrightnessScale = li.BrightnessScale
		}

		nli.ForceUpdateMQTT = li.ForceUpdateMQTT

		if len(li.CommandState) > 0 {
			nli.StateFunc = constructStateFunc(li.CommandState)
		}
		if len(li.CommandBrightnessState) > 0 {
			nli.BrightnessStateFunc = constructStateFunc(li.CommandBrightnessState)
		}
		if len(li.CommandColorTempState) > 0 {
			nli.ColorTempStateFunc = constructStateFunc(li.CommandColorTempState)
		}
		if len(li.CommandEffectState) > 0 {
			nli.EffectStateFunc = constructStateFunc(li.CommandEffectState)
		}
		if len(li.CommandHsState) > 0 {
			nli.HsStateFunc = constructStateFunc(li.CommandHsState)
		}
		if len(li.CommandRgbState) > 0 {
			nli.RgbStateFunc = constructStateFunc(li.CommandRgbState)
		}
		if len(li.CommandWhiteValueState) > 0 {
			nli.WhiteValueStateFunc = constructStateFunc(li.CommandWhiteValueState)
		}
		if len(li.CommandXyState) > 0 {
			nli.XyStateFunc = constructStateFunc(li.CommandXyState)
		}

		if len(li.Command) > 0 {
			nli.CommandFunc = constructCommandFunc(li.Command)
		}
		if len(li.CommandBrightness) > 0 {
			nli.BrightnessCommandFunc = constructCommandFunc(li.CommandBrightness)
		}
		if len(li.CommandColorTemp) > 0 {
			nli.ColorTempCommandFunc = constructCommandFunc(li.CommandColorTemp)
		}
		if len(li.CommandEffect) > 0 {
			nli.EffectCommandFunc = constructCommandFunc(li.CommandEffect)
		}
		if len(li.CommandHs) > 0 {
			nli.HsCommandFunc = constructCommandFunc(li.CommandHs)
		}
		if len(li.CommandRgb) > 0 {
			nli.RgbCommandFunc = constructCommandFunc(li.CommandRgb)
		}
		if len(li.CommandWhiteValue) > 0 {
			nli.WhiteValueCommandFunc = constructCommandFunc(li.CommandWhiteValue)
		}
		if len(li.CommandXy) > 0 {
			nli.XyCommandFunc = constructCommandFunc(li.CommandXy)
		}

		nli.Initialize()

		lights = append(lights, nli)
	}

	if sconfig.Builtin.Backlight.Enable {

		if runtime.GOOS != "linux" {
			log.Println("Backlight not supported on non-linux platforms")
		} else {

			backlights := backlightP.PopulateBacklights()

			for k, backlight := range backlights {

				bLight := hadiscovery.Light{}
				bLight.BrightnessScale = backlight.MaxBrightness
				bLight.UniqueID = "builtin-backlight-" + strconv.Itoa(k)
				bLight.Name = "Backlight " + strconv.Itoa(k)
				bLight.BrightnessCommandFunc = func(message mqtt.Message, client mqtt.Client) {
					backlight.SetBrightness(string(message.Payload()))
				}
				bLight.BrightnessStateFunc = backlight.GetBrightness
				bLight.CommandFunc = func(message mqtt.Message, client mqtt.Client) {
					if string(message.Payload()) == "ON" {
						_, err := exec.Command("/bin/sh", "-c", "xset dpms force on").Output()

						if err != nil {
							log.Printf("%s", err)
						}
					} else if string(message.Payload()) == "OFF" {
						_, err := exec.Command("/bin/sh", "-c", "xset dpms force off").Output()
						if err != nil {
							log.Printf("%s", err)
						}
					} else {
						log.Println("Unknown payload: " + string(message.Payload()))
					}
				}
				bLight.StateFunc = func() string {
					if backlight.IsOn() {
						return "ON"
					}
					return "OFF"
				}
				bLight.UpdateInterval = 1

				if sconfig.Builtin.Backlight.Temperature {
					bLight.ColorTempCommandFunc = func(message mqtt.Message, client mqtt.Client) {
						minred, _ := strconv.ParseInt(string(message.Payload()), 10, 64)

						_, err := exec.Command("./scripts/run-in-user-session.sh", "gsettings", "set", "org.gnome.settings-daemon.plugins.color", "night-light-temperature", strconv.Itoa(int(math.Round(backlightP.MinredToKelvin(float64(minred)))))).Output()
						if err != nil {
							log.Printf("%s", err)
						}
					}
					bLight.ColorTempStateFunc = func() string {
						var out []byte
						out, err := exec.Command("./scripts/run-in-user-session.sh", "gsettings", "get", "org.gnome.settings-daemon.plugins.color", "night-light-temperature").Output()
						if err != nil {
							log.Printf("%s", err)
						}
						fields := strings.Fields(string(out))
						str, err := strconv.ParseInt(fields[1], 10, 64)
						if err != nil {
							log.Println("Error parsing color temp state")
							log.Println(err)
						}
						return strconv.Itoa(int(backlightP.KelvinToMinred(float64(str))))
					}
					bLight.MinMireds = int(math.Round(backlightP.KelvinToMinred(10000)))
					bLight.MaxMireds = int(math.Round(backlightP.KelvinToMinred(3000)))
				}

				bLight.Initialize()

				lights = append(lights, bLight)
			}
		}

	}

	return
}
