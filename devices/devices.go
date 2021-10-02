package devices

import (
	"log"
	"os/exec"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//go:generate go run ../helpers/ha-devices/main.go
//go:generate go run ../helpers/translators/main.go
//go:generate go run ../helpers/generators/main.go
//go:generate go run ../helpers/topics/main.go
//go:generate gofmt -w structs.go
//go:generate gofmt -w translators.go
//go:generate gofmt -w generators.go
//go:generate gofmt -w topics.go

func AddStateFunction(command *[]string, target *func() string) {
	if command != nil {
		*target = ConstructStateFunc(command)
	}
}

func AddCommandFunction(command *[]string, target *func(message mqtt.Message, connection mqtt.Client)) {
	if command != nil {
		*target = ConstructCommandFunc(command)
	}
}

func ConstructStateFunc(command *[]string) (f func() string) {
	var err error
	return func() string {
		var out []byte
		if len(*command) > 1 {
			out, err = exec.Command((*command)[0], (*command)[1:]...).Output()
		} else {

			out, err = exec.Command((*command)[0]).Output()
		}
		if err != nil {
			log.Printf("%s", err)
		}
		return string(out)
	}
}

func ConstructCommandFunc(command *[]string) (f func(message mqtt.Message, connection mqtt.Client)) {
	var err error
	return func(message mqtt.Message, connection mqtt.Client) {
		localcom := *command
		localcom = append(localcom, string(message.Payload()))
		if len(*command) > 0 {

			if len(*command) > 1 {
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

func TopicAvailability(base string) string {
	return base + "availability"
}

func TopicCommand(base string) string {
	return base + "command"
}
