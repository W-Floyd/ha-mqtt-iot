package ExternalDevice

import (
	"encoding/json"
	common "github.com/W-Floyd/ha-mqtt-iot/common"
	store "github.com/W-Floyd/ha-mqtt-iot/store"
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
	AvailabilityMode     string                          `json:"availability_mode"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate string                          `json:"availability_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    string                          `json:"availability_topic"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	CommandTopic         string                          `json:"command_topic"`         // "The MQTT topic to publish `payload_on` to activate the scene."
	CommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
	EnabledByDefault     bool                            `json:"enabled_by_default"`    // "Flag which defines if the entity should be enabled when first added."
	EntityCategory       string                          `json:"entity_category"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                 string                          `json:"icon"`                  // "Icon for the scene."
	Name                 string                          `json:"name"`                  // "The name to use when displaying this scene."
	ObjectId             string                          `json:"object_id"`             // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable     string                          `json:"payload_available"`     // "The payload that represents the available state."
	PayloadNotAvailable  string                          `json:"payload_not_available"` // "The payload that represents the unavailable state."
	PayloadOn            string                          `json:"payload_on"`            // "The payload that will be sent to `command_topic` when activating the MQTT scene."
	Qos                  int                             `json:"qos"`                   // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain               bool                            `json:"retain"`                // "If the published message should have the retain flag on or not."
	UniqueId             string                          `json:"unique_id"`             // "An ID that uniquely identifies this scene entity. If two scenes have the same unique ID, Home Assistant will raise an exception."
	MQTT                 MQTTFields                      `json:"-"`
}

func (d Scene) UpdateState() {}
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
	time.Sleep(common.HADiscoveryDelay)
	d.AnnounceAvailable()
	d.UpdateState()
}
func (d Scene) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, "offline")
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
	token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d Scene) Initialize() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d Scene) PopulateTopics() {
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		store.TopicStore[d.CommandTopic] = &d.CommandFunc
	}
}
func (d Scene) SetMQTTFields(fields MQTTFields) {
	d.MQTT = fields
}
func (d Scene) GetMQTTFields() (fields MQTTFields) {
	return d.MQTT
}
