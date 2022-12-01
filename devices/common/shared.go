package CommonDevices

import (
	"os/exec"
	"strings"

	"github.com/W-Floyd/ha-mqtt-iot/common"
	mqtt "tinygo.org/x/drivers/net/mqtt"
)

func ConstructCommandFunc(command []string) (f func(message mqtt.Message, connection mqtt.Client)) {
	var err error
	return func(message mqtt.Message, connection mqtt.Client) {
		localcom := command
		if string(message.Payload()) != "" {
			localcom = append(localcom, string(message.Payload()))
		}
		if len(localcom) > 0 {
			var out []byte
			if len(localcom) > 1 {
				out, err = exec.Command(localcom[0], localcom[1:]...).Output()
			} else {
				out, err = exec.Command(localcom[0]).Output()
			}
			if err != nil {
				common.LogWarning("Error running command", localcom)
				common.LogWarning(string(out))
				common.LogWarning(err)
			} else {
				common.LogDebug("Ran", localcom)
				common.LogDebug(string(out))
			}

		}
	}
}

func ConstructStateFunc(command []string) (f func() string) {
	var err error
	return func() string {
		var out []byte
		if len(command) > 1 {
			out, err = exec.Command(command[0], command[1:]...).Output()
		} else {
			out, err = exec.Command(command[0]).Output()
		}
		if err != nil {
			common.LogWarning("Error running state command", command)
			common.LogWarning(string(out))
			common.LogWarning(err)
		} else {
			common.LogDebug("Ran", command)
			common.LogDebug(string(out))
		}

		return string(strings.TrimRight(string(out), "\n"))
	}
}

func AvailabilityFunc() string {
	return "online"
}
