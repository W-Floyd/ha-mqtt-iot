package ExternalDevice

import (
	"encoding/json"
	common "github.com/W-Floyd/ha-mqtt-iot/common"
	"log"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (d *DeviceTracker) GetRawId() string {
	return "device_tracker"
}
func (d *DeviceTracker) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d DeviceTracker) GetUniqueId() string {
	return ""
}
func (d *DeviceTracker) PopulateDevice() {}

type DeviceTracker struct {
	Devices        []string   `json:"devices"`          // "List of devices with their topic."
	PayloadHome    string     `json:"payload_home"`     // "The payload value that represents the 'home' state for the device."
	PayloadNotHome string     `json:"payload_not_home"` // "The payload value that represents the 'not_home' state for the device."
	Qos            int        `json:"qos"`              // "The QoS level of the topic."
	SourceType     string     `json:"source_type"`      // "Attribute of a device tracker that affects state when being used to track a [person](/integrations/person/). Valid options are `gps`, `router`, `bluetooth`, or `bluetooth_le`."
	MQTT           MQTTFields `json:"-"`
}

func (d *DeviceTracker) UpdateState() {}
func (d *DeviceTracker) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	token := c.Publish(GetDiscoveryTopic(d), 0, true, message)
	token.Wait()
	time.Sleep(common.HADiscoveryDelay)
	d.UpdateState()
}
func (d *DeviceTracker) UnSubscribe()       {}
func (d *DeviceTracker) AnnounceAvailable() {}
func (d *DeviceTracker) Initialize() {
	d.PopulateDevice()
	d.AddMessageHandler()
	d.PopulateTopics()
}
func (d *DeviceTracker) PopulateTopics() {}
func (d *DeviceTracker) SetMQTTFields(fields MQTTFields) {
	d.MQTT = fields
}
func (d *DeviceTracker) GetMQTTFields() (fields MQTTFields) {
	return d.MQTT
}
