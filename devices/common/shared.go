package CommonDevices

import (
	"os/exec"
	"strings"

	"github.com/W-Floyd/ha-mqtt-iot/common"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ConstructCommandFunc(command []string) (f func(message mqtt.Message, connection mqtt.Client)) {
	var err error
	return func(message mqtt.Message, connection mqtt.Client) {
		localCommand := command
		if string(message.Payload()) != "" {
			localCommand = append(localCommand, string(message.Payload()))
		}
		if len(localCommand) > 0 {
			var out []byte
			if len(localCommand) > 1 {
				out, err = exec.Command(localCommand[0], localCommand[1:]...).Output()
			} else {
				out, err = exec.Command(localCommand[0]).Output()
			}
			if err != nil {
				common.LogWarning("Error running command", localCommand)
				common.LogWarning(string(out))
				common.LogWarning(err)
			} else {
				common.LogDebug("Ran", localCommand)
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
