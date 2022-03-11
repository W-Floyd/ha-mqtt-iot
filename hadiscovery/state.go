package hadiscovery

import (
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
