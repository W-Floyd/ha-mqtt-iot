package hadiscovery

import (
	"encoding/json"
	"log"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (d DeviceTrigger) GetRawId() string {
	return "device_trigger"
}
func (d DeviceTrigger) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d DeviceTrigger) GetUniqueId() string {
	return ""
}
func (d DeviceTrigger) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type DeviceTrigger struct {
	AutomationType string `json:"automation_type"`
	Device         struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	Payload       string     `json:"payload"`
	Qos           int        `json:"qos"`
	Subtype       string     `json:"subtype"`
	Topic         string     `json:"topic"`
	Type          string     `json:"type"`
	ValueTemplate string     `json:"value_template"`
	MQTT          MQTTFields `json:"-"`
}

func (d DeviceTrigger) UpdateState() {}
func (d DeviceTrigger) Subscribe() {
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
func (d DeviceTrigger) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, qos, retain, "offline")
	token.Wait()
}
func (d DeviceTrigger) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, qos, retain, "online")
	token.Wait()
}
func (d DeviceTrigger) Initialize() {
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d DeviceTrigger) PopulateTopics() {}
