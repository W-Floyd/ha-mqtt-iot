package hadiscovery

import mqtt "github.com/eclipse/paho.mqtt.golang"

// GetTopicPrefix gets the prefix for all state/command topics
// This is for a light
func (device Light) GetTopicPrefix() string {
	return NodeID + "/light/" + device.UniqueID + "/"
}

// GetDiscoveryTopic gets the topic for a device's discovery topic.
// This is for a light
func (device Light) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/light/" + NodeID + "/" + device.UniqueID + "/" + "config"
}

// GetCommandTopic gets the command topic for a device
// This is for a light
func (device Light) GetCommandTopic() string {
	return device.GetTopicPrefix() + "command"
}

// GetStateTopic gets the state topic for a device
// This is for a light
func (device Light) GetStateTopic() string {
	return device.GetTopicPrefix() + "state"
}

// GetAvailabilityTopic gets the availability topic for a device
// This is for a light
func (device Light) GetAvailabilityTopic() string {
	return device.GetTopicPrefix() + "availability"
}

// AnnounceAvailable publishes availability for a device
// This is for a light
func (device Light) AnnounceAvailable(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), qos, retain, "online")
	token.Wait()
}
