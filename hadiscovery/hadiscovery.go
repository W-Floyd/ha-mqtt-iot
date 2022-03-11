package hadiscovery

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//go:generate go run ../helpers/

var (
	retain bool = true
	qos    byte = 0
)

var topicStore = make(map[string]*func(mqtt.Message, mqtt.Client))

// DiscoveryPrefix is the prefix that HA uses to discover on.
// Does not usually need changing.
const DiscoveryPrefix = "homeassistant"

// SWVersion is the software version.
// TODO - Move this elsewhere maybe?
const SWVersion = "0.4.5"

var SoftwareName = "Homeassistant MQTT IOT"

// InstanceName is the instance name, helpful for identifying a given client
var InstanceName = "Homeassistant MQTT IOT"

// NodeID is the Node ID, that is, what that node connects under.
var NodeID = "ha-mqtt-iot"

var Manufacturer = "William Floyd"

///////////////////

type store struct {
	BinarySensor map[string]string
	Light        struct {
		BrightnessState map[string]string
		ColorTempState  map[string]string
		EffectState     map[string]string
		HsState         map[string]string
		RgbState        map[string]string
		State           map[string]string
		WhiteValueState map[string]string
		XyState         map[string]string
	}
	Sensor map[string]string
	Switch map[string]string
}

// UpdateState publishes any new state for a device
// This is for a light
// func (device Light) UpdateState(client mqtt.Client) {
// 	if device.BrightnessStateTopic != "" {
// 		state := device.BrightnessStateFunc()
// 		if state != stateStore.Light.BrightnessState[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.BrightnessStateTopic, qos, retain, state)
// 			stateStore.Light.BrightnessState[device.UniqueID] = state
// 			token.Wait()
// 		}
// 	}
// 	if device.ColorTempStateTopic != "" {
// 		state := device.ColorTempStateFunc()
// 		if state != stateStore.Light.ColorTempState[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.ColorTempStateTopic, qos, retain, state)
// 			stateStore.Light.ColorTempState[device.UniqueID] = state
// 			token.Wait()
// 		}
// 	}
// 	if device.EffectStateTopic != "" {
// 		state := device.EffectStateFunc()
// 		if state != stateStore.Light.EffectState[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.EffectStateTopic, qos, retain, state)
// 			stateStore.Light.EffectState[device.UniqueID] = state
// 			token.Wait()
// 		}
// 	}
// 	if device.HsStateTopic != "" {
// 		state := device.HsStateFunc()
// 		if state != stateStore.Light.HsState[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.HsStateTopic, qos, retain, state)
// 			stateStore.Light.HsState[device.UniqueID] = state
// 			token.Wait()
// 		}
// 	}
// 	if device.RgbStateTopic != "" {
// 		state := device.RgbStateFunc()
// 		if state != stateStore.Light.RgbState[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.RgbStateTopic, qos, retain, state)
// 			stateStore.Light.RgbState[device.UniqueID] = state
// 			token.Wait()
// 		}
// 	}
// 	if device.StateTopic != "" {
// 		state := device.StateFunc()
// 		if state != stateStore.Light.State[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.StateTopic, qos, retain, state)
// 			stateStore.Light.State[device.UniqueID] = state
// 			token.Wait()
// 		}
// 	}
// 	if device.WhiteValueStateTopic != "" {
// 		state := device.WhiteValueStateFunc()
// 		if state != stateStore.Light.WhiteValueState[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.WhiteValueStateTopic, qos, retain, state)
// 			stateStore.Light.WhiteValueState[device.UniqueID] = state
// 			token.Wait()
// 		}
// 	}
// 	if device.XyStateTopic != "" {
// 		state := device.XyStateFunc()
// 		if state != stateStore.Light.XyState[device.UniqueID] || device.ForceUpdateMQTT {
// 			token := client.Publish(device.XyStateTopic, qos, retain, state)
// 			stateStore.Light.XyState[device.UniqueID] = state
// 			token.Wait()
// 		}
// 	}
// }

// Subscribe announces and starts listening to MQTT topics appropriate for a device
// This is for a light
// func (device Light) Subscribe(client mqtt.Client) {

// 	message, err := json.Marshal(device)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if device.BrightnessCommandTopic != "" {
// 		if token := client.Subscribe(device.BrightnessCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
// 	}
// 	if device.ColorTempCommandTopic != "" {
// 		if token := client.Subscribe(device.ColorTempCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
// 	}
// 	if device.EffectCommandTopic != "" {
// 		if token := client.Subscribe(device.EffectCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
// 	}
// 	if device.HsCommandTopic != "" {
// 		if token := client.Subscribe(device.HsCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
// 	}
// 	if device.RgbCommandTopic != "" {
// 		if token := client.Subscribe(device.RgbCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
// 	}
// 	if device.CommandTopic != "" {
// 		if token := client.Subscribe(device.CommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
// 	}
// 	if device.WhiteValueCommandTopic != "" {
// 		if token := client.Subscribe(device.WhiteValueCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
// 	}
// 	if device.XyCommandTopic != "" {
// 		if token := client.Subscribe(device.XyCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}
// 	}

// 	token := client.Publish(device.GetDiscoveryTopic(), 0, true, message)
// 	token.Wait()

// 	// HA needs time to process
// 	time.Sleep(500 * time.Millisecond)

// 	device.AnnounceAvailable(client)

// 	device.UpdateState(client)

// 	if token := client.Subscribe(GetCommandTopic(device), 0, device.messageHandler); token.Wait() && token.Error() != nil {
// 		log.Println(token.Error())
// 		os.Exit(1)
// 	}

// }

// UnSubscribe from MQTT topics appropriate for a device, and publishes unavailability
// This is for a light
// func (device Light) UnSubscribe(client mqtt.Client) {
// 	token := client.Publish(device.GetAvailabilityTopic(), qos, retain, "offline")
// 	token.Wait()

// 	if device.BrightnessCommandTopic != "" {
// 		if token := client.Unsubscribe(device.BrightnessCommandTopic); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}

// 	}
// 	if device.ColorTempCommandTopic != "" {
// 		if token := client.Unsubscribe(device.ColorTempCommandTopic); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}

// 	}
// 	if device.EffectCommandTopic != "" {
// 		if token := client.Unsubscribe(device.EffectCommandTopic); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}

// 	}
// 	if device.HsCommandTopic != "" {
// 		if token := client.Unsubscribe(device.HsCommandTopic); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}

// 	}
// 	if device.RgbCommandTopic != "" {
// 		if token := client.Unsubscribe(device.RgbCommandTopic); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}

// 	}
// 	if device.CommandTopic != "" {
// 		if token := client.Unsubscribe(device.CommandTopic); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}

// 	}
// 	if device.WhiteValueCommandTopic != "" {
// 		if token := client.Unsubscribe(device.WhiteValueCommandTopic); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}

// 	}
// 	if device.XyCommandTopic != "" {
// 		if token := client.Unsubscribe(device.XyCommandTopic); token.Wait() && token.Error() != nil {
// 			log.Println(token.Error())
// 			os.Exit(1)
// 		}

// 	}

// }

// AnnounceAvailable publishes availability for a device
// This is for a light
// func (device Light) AnnounceAvailable(client mqtt.Client) {
// 	token := client.Publish(device.GetAvailabilityTopic(), qos, retain, "online")
// 	token.Wait()
// }

///////////////////
