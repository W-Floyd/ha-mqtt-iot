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
func (d AlarmControlPanel) GetRawId() string {
	return "alarm_control_panel"
}
func (d AlarmControlPanel) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d AlarmControlPanel) GetUniqueId() string {
	return d.UniqueId
}
func (d AlarmControlPanel) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type AlarmControlPanel struct {
	AvailabilityMode     string                          `json:"availability_mode"`
	AvailabilityTemplate string                          `json:"availability_template"`
	AvailabilityTopic    string                          `json:"availability_topic"`
	Code                 string                          `json:"code"`
	CodeArmRequired      bool                            `json:"code_arm_required"`
	CodeDisarmRequired   bool                            `json:"code_disarm_required"`
	CodeTriggerRequired  bool                            `json:"code_trigger_required"`
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
	EnabledByDefault       bool          `json:"enabled_by_default"`
	Encoding               string        `json:"encoding"`
	EntityCategory         string        `json:"entity_category"`
	Icon                   string        `json:"icon"`
	Name                   string        `json:"name"`
	ObjectId               string        `json:"object_id"`
	PayloadArmAway         string        `json:"payload_arm_away"`
	PayloadArmCustomBypass string        `json:"payload_arm_custom_bypass"`
	PayloadArmHome         string        `json:"payload_arm_home"`
	PayloadArmNight        string        `json:"payload_arm_night"`
	PayloadArmVacation     string        `json:"payload_arm_vacation"`
	PayloadAvailable       string        `json:"payload_available"`
	PayloadDisarm          string        `json:"payload_disarm"`
	PayloadNotAvailable    string        `json:"payload_not_available"`
	PayloadTrigger         string        `json:"payload_trigger"`
	Qos                    int           `json:"qos"`
	Retain                 bool          `json:"retain"`
	StateTopic             string        `json:"state_topic"`
	StateFunc              func() string `json:"-"`
	UniqueId               string        `json:"unique_id"`
	ValueTemplate          string        `json:"value_template"`
	MQTT                   MQTTFields    `json:"-"`
}

func (d AlarmControlPanel) UpdateState() {
	if d.AvailabilityTopic != "" {
		state := d.AvailabilityFunc()
		if state != stateStore.AlarmControlPanel.Availability[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.AvailabilityTopic, qos, retain, state)
			stateStore.AlarmControlPanel.Availability[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.StateTopic != "" {
		state := d.StateFunc()
		if state != stateStore.AlarmControlPanel.State[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.StateTopic, qos, retain, state)
			stateStore.AlarmControlPanel.State[d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d AlarmControlPanel) Subscribe() {
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
func (d AlarmControlPanel) UnSubscribe() {
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
func (d AlarmControlPanel) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, qos, retain, "online")
	token.Wait()
}
func (d AlarmControlPanel) Initialize() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d AlarmControlPanel) PopulateTopics() {
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
