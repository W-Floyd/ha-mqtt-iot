package ExternalDevice

import (
	"encoding/json"
	"log"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (d Camera) GetRawId() string {
	return "camera"
}
func (d Camera) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Camera) GetUniqueId() string {
	return d.UniqueId
}
func (d Camera) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Camera struct {
	AvailabilityMode     string `json:"availability_mode"`
	AvailabilityTemplate string `json:"availability_template"`
	AvailabilityTopic    string `json:"availability_topic"`
	Device               struct {
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
	EnabledByDefault bool       `json:"enabled_by_default"`
	EntityCategory   string     `json:"entity_category"`
	Icon             string     `json:"icon"`
	Name             string     `json:"name"`
	ObjectId         string     `json:"object_id"`
	Topic            string     `json:"topic"`
	UniqueId         string     `json:"unique_id"`
	MQTT             MQTTFields `json:"-"`
}

func (d Camera) UpdateState() {}
func (d Camera) Subscribe() {
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
func (d Camera) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, qos, retain, "offline")
	token.Wait()
}
func (d Camera) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, qos, retain, "online")
	token.Wait()
}
func (d Camera) Initialize() {
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d Camera) PopulateTopics() {}
