package hadiscovery

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Initialize sets topics as needed on a Light
func (device *Light) Initialize() {
	device.Retain = false
	device.Device = getDevice()

	device.AvailabilityTopic = device.GetAvailabilityTopic()

	// Brightness
	if device.BrightnessCommandFunc != nil {
		device.BrightnessCommandTopic = device.GetTopicPrefix() + "brightness/set"
		topicStore[device.BrightnessCommandTopic] = &device.BrightnessCommandFunc
	}
	if device.BrightnessStateFunc != nil {
		device.BrightnessStateTopic = device.GetTopicPrefix() + "brightness/get"
	}

	// ColorTemp
	if device.ColorTempCommandFunc != nil {
		device.ColorTempCommandTopic = device.GetTopicPrefix() + "color-temp/set"
		topicStore[device.ColorTempCommandTopic] = &device.ColorTempCommandFunc
	}
	if device.ColorTempStateFunc != nil {
		device.ColorTempStateTopic = device.GetTopicPrefix() + "color-temp/get"
	}

	// Command/State
	if device.CommandFunc != nil {
		device.CommandTopic = device.GetCommandTopic()
		topicStore[device.CommandTopic] = &device.CommandFunc
	}
	if device.StateFunc != nil {
		device.StateTopic = device.GetStateTopic()
	}

	// Effect
	if device.EffectCommandFunc != nil {
		device.EffectCommandTopic = device.GetTopicPrefix() + "effect/set"
		topicStore[device.EffectCommandTopic] = &device.EffectCommandFunc
	}
	if device.EffectStateFunc != nil {
		device.EffectStateTopic = device.GetTopicPrefix() + "effect/get"
	}

	// Hs
	if device.HsCommandFunc != nil {
		device.HsCommandTopic = device.GetTopicPrefix() + "hs/set"
		topicStore[device.HsCommandTopic] = &device.HsCommandFunc
	}
	if device.HsStateFunc != nil {
		device.HsStateTopic = device.GetTopicPrefix() + "hs/get"
	}

	// Rgb
	if device.RgbCommandFunc != nil {
		device.RgbCommandTopic = device.GetTopicPrefix() + "rgb/set"
		topicStore[device.RgbCommandTopic] = &device.RgbCommandFunc
	}
	if device.RgbStateFunc != nil {
		device.RgbStateTopic = device.GetTopicPrefix() + "rgb/get"
	}

	// WhiteValue
	if device.WhiteValueCommandFunc != nil {
		device.WhiteValueCommandTopic = device.GetTopicPrefix() + "white-value/set"
		topicStore[device.WhiteValueCommandTopic] = &device.WhiteValueCommandFunc
	}
	if device.WhiteValueStateFunc != nil {
		device.WhiteValueStateTopic = device.GetTopicPrefix() + "white-value/get"
	}

	// Xy
	if device.XyCommandFunc != nil {
		device.XyCommandTopic = device.GetTopicPrefix() + "xy/set"
		topicStore[device.XyCommandTopic] = &device.XyCommandFunc
	}
	if device.XyStateFunc != nil {
		device.XyStateTopic = device.GetTopicPrefix() + "xy/get"
	}

	device.messageHandler = func(client mqtt.Client, msg mqtt.Message) {

		topicFound := false

		for topic, f := range topicStore {
			if msg.Topic() == topic {
				topicFound = true
				(*f)(msg, client)
				device.UpdateState(client)
			}
		}

		if !topicFound {
			log.Println("Unknown Message on topic " + msg.Topic())
			log.Println(msg.Payload())
		}
	}
}

// Initialize sets topics as needed on a Switch
func (device *Switch) Initialize() {
	device.CommandTopic = device.GetCommandTopic()
	if device.StateFunc != nil {
		device.StateTopic = device.GetStateTopic()
	}
	device.AvailabilityTopic = device.GetAvailabilityTopic()
	device.Device = getDevice()
	device.Retain = false

	topicStore[device.CommandTopic] = &device.CommandFunc

	device.messageHandler = func(client mqtt.Client, msg mqtt.Message) {

		topicFound := false

		for topic, f := range topicStore {
			if msg.Topic() == topic {
				topicFound = true
				(*f)(msg, client)
				device.UpdateState(client)
			}
		}

		if !topicFound {
			log.Println("Unknown Message on topic " + msg.Topic())
			log.Println(msg.Payload())
		}
	}

}

// Initialize sets topics as needed on a Sensor
func (device *Sensor) Initialize() {
	device.StateTopic = device.GetStateTopic()
	device.AvailabilityTopic = device.GetAvailabilityTopic()
	device.Device = getDevice()
}

// Initialize sets topics as needed on a Binary Sensor
func (device *BinarySensor) Initialize() {
	device.StateTopic = device.GetStateTopic()
	device.AvailabilityTopic = device.GetAvailabilityTopic()
	device.Device = getDevice()
}
