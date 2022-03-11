package hadiscovery

import (
	"encoding/json"
	"log"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (d DeviceTracker) GetRawId() string {
	return "device_tracker"
}
func (d DeviceTracker) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d DeviceTracker) GetUniqueId() string {
	return ""
}
func (d DeviceTracker) PopulateDevice() {}

type DeviceTracker struct {
	Devices        []string   `json:"devices"`
	PayloadHome    string     `json:"payload_home"`
	PayloadNotHome string     `json:"payload_not_home"`
	Qos            int        `json:"qos"`
	SourceType     string     `json:"source_type"`
	MQTT           MQTTFields `json:"-"`
}

func (d DeviceTracker) UpdateState() {}
func (d DeviceTracker) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	token := c.Publish(GetDiscoveryTopic(d), 0, true, message)
	token.Wait()
	time.Sleep(500 * time.Millisecond)
	d.AnnounceAvailable()
	d.UpdateState()
}
func (d DeviceTracker) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, qos, retain, "offline")
	token.Wait()
}
func (d DeviceTracker) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, qos, retain, "online")
	token.Wait()
}
func (d DeviceTracker) Initialize() {
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d DeviceTracker) PopulateTopics() {}
