package hadiscovery

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (d Scene) GetRawId() string {
	return "scene"
}
func (d Scene) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Scene) GetUniqueId() string {
	return d.UniqueId
}
func (d Scene) PopulateDevice() {}

type Scene struct {
	AvailabilityMode     string                          `json:"availability_mode"`
	AvailabilityTemplate string                          `json:"availability_template"`
	AvailabilityTopic    string                          `json:"availability_topic"`
	AvailabilityFunc     func() string                   `json:"-"`
	CommandTopic         string                          `json:"command_topic"`
	CommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
	EnabledByDefault     bool                            `json:"enabled_by_default"`
	EntityCategory       string                          `json:"entity_category"`
	Icon                 string                          `json:"icon"`
	Name                 string                          `json:"name"`
	ObjectId             string                          `json:"object_id"`
	PayloadAvailable     string                          `json:"payload_available"`
	PayloadNotAvailable  string                          `json:"payload_not_available"`
	PayloadOn            string                          `json:"payload_on"`
	Qos                  int                             `json:"qos"`
	Retain               bool                            `json:"retain"`
	UniqueId             string                          `json:"unique_id"`
	MQTT                 MQTTFields                      `json:"-"`
}

func (d Scene) UpdateState() {
	if d.AvailabilityTopic != "" {
		state := d.AvailabilityFunc()
		if state != stateStore.Scene.Availability[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.AvailabilityTopic, qos, retain, state)
			stateStore.Scene.Availability[d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d Scene) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	if d.CommandTopic != "" {
		t := c.Subscribe(d.CommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	token := c.Publish(GetDiscoveryTopic(d), 0, true, message)
	token.Wait()
	time.Sleep(500 * time.Millisecond)
	d.AnnounceAvailable()
	d.UpdateState()
}
func (d Scene) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, qos, retain, "offline")
	token.Wait()
	if d.CommandTopic != "" {
		t := c.Unsubscribe(d.CommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d Scene) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, qos, retain, "online")
	token.Wait()
}
func (d Scene) Initialize() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d Scene) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
}
