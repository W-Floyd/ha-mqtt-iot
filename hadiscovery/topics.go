package hadiscovery

import mqtt "github.com/eclipse/paho.mqtt.golang"

// GetTopicPrefix gets the prefix for all state/command topics
// This is for a light
func (device Light) GetTopicPrefix() string {
	return NodeID + "/light/" + device.UniqueID + "/"
}

// GetTopicPrefix gets the prefix for all state/command topics
// This is for a sensor
func (device Sensor) GetTopicPrefix() string {
	return NodeID + "/sensor/" + device.UniqueID + "/"
}

// GetTopicPrefix gets the prefix for all state/command topics
// This is for a switch
func (device Switch) GetTopicPrefix() string {
	return NodeID + "/switch/" + device.UniqueID + "/"
}

// GetTopicPrefix gets the prefix for all state/command topics
// This is for a binary sensor
func (device BinarySensor) GetTopicPrefix() string {
	return NodeID + "/binary_sensor/" + device.UniqueID + "/"
}

// GetDiscoveryTopic gets the topic for a device's discovery topic.
// This is for a light
func (device Light) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/light/" + NodeID + "/" + device.UniqueID + "/" + "config"
}

// GetDiscoveryTopic gets the topic for a device's discovery topic.
// This is for a sensor
func (device Sensor) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/sensor/" + NodeID + "/" + device.UniqueID + "/" + "config"
}

// GetDiscoveryTopic gets the topic for a device's discovery topic.
// This is for a switch
func (device Switch) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/switch/" + NodeID + "/" + device.UniqueID + "/" + "config"
}

// GetDiscoveryTopic gets the topic for a device's discovery topic.
// This is for a binary sensor
func (device BinarySensor) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/binary_sensor/" + NodeID + "/" + device.UniqueID + "/" + "config"
}

// GetCommandTopic gets the command topic for a device
// This is for a light
func (device Light) GetCommandTopic() string {
	return device.GetTopicPrefix() + "command"
}

// GetCommandTopic gets the command topic for a device
// This is for a switch
func (device Switch) GetCommandTopic() string {
	return device.GetTopicPrefix() + "command"
}

// GetStateTopic gets the state topic for a device
// This is for a light
func (device Light) GetStateTopic() string {
	return device.GetTopicPrefix() + "state"
}

// GetStateTopic gets the state topic for a device
// This is for a sensor
func (device Sensor) GetStateTopic() string {
	return device.GetTopicPrefix() + "state"
}

// GetStateTopic gets the state topic for a device
// This is for a switch
func (device Switch) GetStateTopic() string {
	return device.GetTopicPrefix() + "state"
}

// GetStateTopic gets the state topic for a device
// This is for a binary sensor
func (device BinarySensor) GetStateTopic() string {
	return device.GetTopicPrefix() + "state"
}

// GetAvailabilityTopic gets the availability topic for a device
// This is for a light
func (device Light) GetAvailabilityTopic() string {
	return device.GetTopicPrefix() + "availability"
}

// GetAvailabilityTopic gets the availability topic for a device
// This is for a sensor
func (device Sensor) GetAvailabilityTopic() string {
	return device.GetTopicPrefix() + "availability"
}

// GetAvailabilityTopic gets the availability topic for a device
// This is for a switch
func (device Switch) GetAvailabilityTopic() string {
	return device.GetTopicPrefix() + "availability"
}

// GetAvailabilityTopic gets the availability topic for a device
// This is for a binary sensor
func (device BinarySensor) GetAvailabilityTopic() string {
	return device.GetTopicPrefix() + "availability"
}

// AnnounceAvailable publishes availability for a device
// This is for a light
func (device Light) AnnounceAvailable(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), qos, retain, "online")
	token.Wait()
}

// AnnounceAvailable publishes availability for a device
// This is for a switch
func (device Switch) AnnounceAvailable(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), qos, retain, "online")
	token.Wait()
}

// AnnounceAvailable publishes availability for a device
// This is for a sensor
func (device Sensor) AnnounceAvailable(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), qos, retain, "online")
	token.Wait()
}

// AnnounceAvailable publishes availability for a device
// This is for a binary sensor
func (device BinarySensor) AnnounceAvailable(client mqtt.Client) {
	token := client.Publish(device.GetAvailabilityTopic(), qos, retain, "online")
	token.Wait()
}
