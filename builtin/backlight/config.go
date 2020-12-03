package backlight

import (
	"log"
	"math"
	"os/exec"
	"strconv"
	"strings"

	"../../hadiscovery"
	"../../iotconfig"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var sconfig = iotconfig.Config{}

func (backlights []Backlight) Translate() (lightsOutput []hadiscovery.Light) {

	for k, backlight := range backlights {

		bLight := hadiscovery.Light{}
		bLight.BrightnessScale = backlight.MaxBrightness
		if len(backlights) > 1 {
			bLight.UniqueID = "builtin-backlight-" + strconv.Itoa(k) + "_" + hadiscovery.NodeID
			bLight.Name = sconfig.Builtin.Prefix + "Backlight " + strconv.Itoa(k)
		} else {
			bLight.UniqueID = "builtin-backlight_" + hadiscovery.NodeID
			bLight.Name = sconfig.Builtin.Prefix + "Backlight"
		}
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

		lightsOutput = append(lightsOutput, bLight)

	}

	return

}
