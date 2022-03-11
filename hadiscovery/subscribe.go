package hadiscovery

import (
	"encoding/json"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Subscribe announces and starts listening to MQTT topics appropriate for a device
// This is for a light
func (device Light) Subscribe(client mqtt.Client) {

	message, err := json.Marshal(device)
	if err != nil {
		log.Fatal(err)
	}

	if device.BrightnessCommandTopic != "" {
		if token := client.Subscribe(device.BrightnessCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}
	}
	if device.ColorTempCommandTopic != "" {
		if token := client.Subscribe(device.ColorTempCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}
	}
	if device.EffectCommandTopic != "" {
		if token := client.Subscribe(device.EffectCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}
	}
	if device.HsCommandTopic != "" {
		if token := client.Subscribe(device.HsCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}
	}
	if device.RgbCommandTopic != "" {
		if token := client.Subscribe(device.RgbCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}
	}
	if device.CommandTopic != "" {
		if token := client.Subscribe(device.CommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}
	}
	if device.WhiteValueCommandTopic != "" {
		if token := client.Subscribe(device.WhiteValueCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}
	}
	if device.XyCommandTopic != "" {
		if token := client.Subscribe(device.XyCommandTopic, 0, device.messageHandler); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}
	}

	token := client.Publish(device.GetDiscoveryTopic(), 0, true, message)
	token.Wait()

	// HA needs time to process
	time.Sleep(500 * time.Millisecond)

	device.AnnounceAvailable(client)

	device.UpdateState(client)

	if token := client.Subscribe(device.GetCommandTopic(), 0, device.messageHandler); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		os.Exit(1)
	}

}

// UnSubscribe from MQTT topics appropriate for a device, and publishes unavailability
// This is for a light
func (device Light) UnSubscribe(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), qos, retain, "offline")
	token.Wait()

	if device.BrightnessCommandTopic != "" {
		if token := client.Unsubscribe(device.BrightnessCommandTopic); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}

	}
	if device.ColorTempCommandTopic != "" {
		if token := client.Unsubscribe(device.ColorTempCommandTopic); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}

	}
	if device.EffectCommandTopic != "" {
		if token := client.Unsubscribe(device.EffectCommandTopic); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}

	}
	if device.HsCommandTopic != "" {
		if token := client.Unsubscribe(device.HsCommandTopic); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}

	}
	if device.RgbCommandTopic != "" {
		if token := client.Unsubscribe(device.RgbCommandTopic); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}

	}
	if device.CommandTopic != "" {
		if token := client.Unsubscribe(device.CommandTopic); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}

	}
	if device.WhiteValueCommandTopic != "" {
		if token := client.Unsubscribe(device.WhiteValueCommandTopic); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}

	}
	if device.XyCommandTopic != "" {
		if token := client.Unsubscribe(device.XyCommandTopic); token.Wait() && token.Error() != nil {
			log.Println(token.Error())
			os.Exit(1)
		}

	}

}
