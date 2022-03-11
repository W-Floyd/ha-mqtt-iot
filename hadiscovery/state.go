package hadiscovery

import (
	"log"
	"strconv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// UpdateState publishes any new state for a device
// This is for a light
func (device Light) UpdateState(client mqtt.Client) {
	if device.BrightnessStateTopic != "" {
		state := device.BrightnessStateFunc()
		if state != stateStore.Light.BrightnessState[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.BrightnessStateTopic, qos, retain, state)
			stateStore.Light.BrightnessState[device.UniqueID] = state
			token.Wait()
		}
	}
	if device.ColorTempStateTopic != "" {
		state := device.ColorTempStateFunc()
		if state != stateStore.Light.ColorTempState[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.ColorTempStateTopic, qos, retain, state)
			stateStore.Light.ColorTempState[device.UniqueID] = state
			token.Wait()
		}
	}
	if device.EffectStateTopic != "" {
		state := device.EffectStateFunc()
		if state != stateStore.Light.EffectState[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.EffectStateTopic, qos, retain, state)
			stateStore.Light.EffectState[device.UniqueID] = state
			token.Wait()
		}
	}
	if device.HsStateTopic != "" {
		state := device.HsStateFunc()
		if state != stateStore.Light.HsState[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.HsStateTopic, qos, retain, state)
			stateStore.Light.HsState[device.UniqueID] = state
			token.Wait()
		}
	}
	if device.RgbStateTopic != "" {
		state := device.RgbStateFunc()
		if state != stateStore.Light.RgbState[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.RgbStateTopic, qos, retain, state)
			stateStore.Light.RgbState[device.UniqueID] = state
			token.Wait()
		}
	}
	if device.StateTopic != "" {
		state := device.StateFunc()
		if state != stateStore.Light.State[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.StateTopic, qos, retain, state)
			stateStore.Light.State[device.UniqueID] = state
			token.Wait()
		}
	}
	if device.WhiteValueStateTopic != "" {
		state := device.WhiteValueStateFunc()
		if state != stateStore.Light.WhiteValueState[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.WhiteValueStateTopic, qos, retain, state)
			stateStore.Light.WhiteValueState[device.UniqueID] = state
			token.Wait()
		}
	}
	if device.XyStateTopic != "" {
		state := device.XyStateFunc()
		if state != stateStore.Light.XyState[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.XyStateTopic, qos, retain, state)
			stateStore.Light.XyState[device.UniqueID] = state
			token.Wait()
		}
	}
}

// UpdateState publishes any new state for a device
// This is for a switch
func (device Switch) UpdateState(client mqtt.Client) {
	if device.StateFunc != nil {
		state := device.StateFunc()
		if state != stateStore.Switch[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.StateTopic, qos, retain, state)
			stateStore.Switch[device.UniqueID] = state
			token.Wait()
		}
	} else {
		log.Println("No statefunc for " + device.UniqueID + strconv.FormatFloat(float64(device.UpdateInterval), 'g', 1, 64))
	}
}

// UpdateState publishes any new state for a device
// This is for a sensor
func (device Sensor) UpdateState(client mqtt.Client) {
	if device.StateFunc != nil {
		state := device.StateFunc()
		if state != stateStore.Sensor[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.StateTopic, qos, retain, state)
			token.Wait()
			stateStore.Sensor[device.UniqueID] = state
		}
	} else {
		log.Fatalln("No statefunc, this makes no sensor for a sensor!")
	}
}

// UpdateState publishes any new state for a device
// This is for a binary sensor
func (device BinarySensor) UpdateState(client mqtt.Client) {
	if device.StateFunc != nil {
		state := device.StateFunc()
		if state != stateStore.BinarySensor[device.UniqueID] || device.ForceUpdateMQTT {
			token := client.Publish(device.StateTopic, qos, retain, state)
			token.Wait()
			stateStore.BinarySensor[device.UniqueID] = state
		}
	} else {
		log.Fatalln("No statefunc, this makes no sensor for a binary sensor!")
	}
}
