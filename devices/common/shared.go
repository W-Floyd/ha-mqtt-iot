package CommonDevices

import (
	"log"
	"os/exec"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ConstructCommandFunc(command []string) (f func(message mqtt.Message, connection mqtt.Client)) {
	var err error
	return func(message mqtt.Message, connection mqtt.Client) {
		localcom := command
		localcom = append(localcom, string(message.Payload()))
		if len(command) > 0 {
			var out []byte
			if len(command) > 1 {
				out, err = exec.Command(localcom[0], localcom[1:]...).Output()
			} else {
				out, err = exec.Command(localcom[0]).Output()
			}
			if err != nil {
				log.Println("Error running command", localcom)
				log.Println(string(out))
				log.Printf("%s", err)
			} else {
				log.Println("Ran", localcom)
				log.Println(string(out))
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
			log.Println("Error running state command", command)
			log.Println(string(out))
			log.Printf("%s", err)
		} else {
			log.Println("Ran", command)
			log.Println(string(out))
		}

		return string(strings.TrimRight(string(out), "\n"))
	}
}

func AvailabilityFunc() string {
	return "online"
}
