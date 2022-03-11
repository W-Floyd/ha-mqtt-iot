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
func (d Siren) GetRawId() string {
	return "siren"
}
func (d Siren) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Siren) GetUniqueId() string {
	return d.UniqueId
}
func (d Siren) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Siren struct {
	AvailabilityMode     string                          `json:"availability_mode"`
	AvailabilityTemplate string                          `json:"availability_template"`
	AvailabilityTopic    string                          `json:"availability_topic"`
	AvailabilityFunc     func() string                   `json:"-"`
	AvailableTones       []string                        `json:"available_tones"`
	CommandOffTemplate   string                          `json:"command_off_template"`
	CommandTemplate      string                          `json:"command_template"`
	CommandTopic         string                          `json:"command_topic"`
	CommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
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
	EnabledByDefault    bool          `json:"enabled_by_default"`
	Encoding            string        `json:"encoding"`
	EntityCategory      string        `json:"entity_category"`
	Icon                string        `json:"icon"`
	Name                string        `json:"name"`
	ObjectId            string        `json:"object_id"`
	Optimistic          bool          `json:"optimistic"`
	PayloadAvailable    string        `json:"payload_available"`
	PayloadNotAvailable string        `json:"payload_not_available"`
	PayloadOff          string        `json:"payload_off"`
	PayloadOn           string        `json:"payload_on"`
	Qos                 int           `json:"qos"`
	Retain              bool          `json:"retain"`
	StateOff            string        `json:"state_off"`
	StateOn             string        `json:"state_on"`
	StateTopic          string        `json:"state_topic"`
	StateFunc           func() string `json:"-"`
	StateValueTemplate  string        `json:"state_value_template"`
	SupportDuration     bool          `json:"support_duration"`
	SupportVolumeSet    bool          `json:"support_volume_set"`
	UniqueId            string        `json:"unique_id"`
	MQTT                MQTTFields    `json:"-"`
}

func (d Siren) UpdateState() {
	if d.AvailabilityTopic != "" {
		state := d.AvailabilityFunc()
		if state != stateStore.Siren.Availability[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.AvailabilityTopic, qos, retain, state)
			stateStore.Siren.Availability[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.StateTopic != "" {
		state := d.StateFunc()
		if state != stateStore.Siren.State[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.StateTopic, qos, retain, state)
			stateStore.Siren.State[d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d Siren) Subscribe() {
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
func (d Siren) UnSubscribe() {
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
func (d Siren) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, qos, retain, "online")
	token.Wait()
}
func (d Siren) Initialize() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d Siren) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
