package devices

import (
	"bufio"
	"log"
	"os/exec"

	"github.com/W-Floyd/ha-mqtt-iot/common"
	"github.com/W-Floyd/ha-mqtt-iot/logging"
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

func ConstructStateWatcher(command *[]string) (output chan string, stop chan bool) {
	output = make(chan string)
	stop = make(chan bool)
	go func() {

		var cmd *exec.Cmd
		if len(*command) > 1 {
			cmd = exec.Command((*command)[0], (*command)[1:]...)
		} else {
			cmd = exec.Command((*command)[0])
		}

		stdout, _ := cmd.StdoutPipe()

		// Make a new channel which will be used to ensure we get all output
		done := make(chan bool)

		// Create a scanner which scans r in a line-by-line fashion
		scanner := bufio.NewScanner(stdout)

		// Use the scanner to scan the output line by line and log it
		// It's running in a goroutine so that it doesn't block
		go func() {

			// Read line by line and process it
			for scanner.Scan() {
				output <- scanner.Text()
			}

			// We're all done, unblock the channel
			done <- true

		}()

		cmd.Start()

		for {
			switch {
			case <-done:
				logging.LogDebug("Command ended: ", command)
			case <-stop:
				cmd.Process.Kill()
				logging.LogDebug("Command killed: ", command)
			}
		}

	}()
	return
}

func ConstructStateWatcherPublisher(command *[]string, publishTopic string, client mqtt.Client) chan bool {

	quitPublisher := make(chan bool)

	outputWatcher, quitWatcher := ConstructStateWatcher(command)

	go func() {
		for {
			select {
			case <-quitPublisher:
				quitWatcher <- true
				return
			case <-outputWatcher:
				client.Publish(publishTopic, common.MqttQos, common.MqttRetain, outputWatcher)
			}
		}
	}()

	return quitWatcher
}

func TopicAvailability(base string) string {
	return base + "availability"
}

func TopicCommand(base string) string {
	return base + "command"
}
