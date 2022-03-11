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
func (d Humidifier) GetRawId() string {
	return "humidifier"
}
func (d Humidifier) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Humidifier) GetUniqueId() string {
	return d.UniqueId
}
func (d Humidifier) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Humidifier struct {
	AvailabilityMode     string                          `json:"availability_mode"`
	AvailabilityTemplate string                          `json:"availability_template"`
	AvailabilityTopic    string                          `json:"availability_topic"`
	AvailabilityFunc     func() string                   `json:"-"`
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
	DeviceClass                   string                          `json:"device_class"`
	EnabledByDefault              bool                            `json:"enabled_by_default"`
	Encoding                      string                          `json:"encoding"`
	EntityCategory                string                          `json:"entity_category"`
	Icon                          string                          `json:"icon"`
	MaxHumidity                   int                             `json:"max_humidity"`
	MinHumidity                   int                             `json:"min_humidity"`
	ModeCommandTemplate           string                          `json:"mode_command_template"`
	ModeCommandTopic              string                          `json:"mode_command_topic"`
	ModeCommandFunc               func(mqtt.Message, mqtt.Client) `json:"-"`
	ModeStateTemplate             string                          `json:"mode_state_template"`
	ModeStateTopic                string                          `json:"mode_state_topic"`
	ModeStateFunc                 func() string                   `json:"-"`
	Modes                         []string                        `json:"modes"`
	Name                          string                          `json:"name"`
	ObjectId                      string                          `json:"object_id"`
	Optimistic                    bool                            `json:"optimistic"`
	PayloadAvailable              string                          `json:"payload_available"`
	PayloadNotAvailable           string                          `json:"payload_not_available"`
	PayloadOff                    string                          `json:"payload_off"`
	PayloadOn                     string                          `json:"payload_on"`
	PayloadResetHumidity          string                          `json:"payload_reset_humidity"`
	PayloadResetMode              string                          `json:"payload_reset_mode"`
	Qos                           int                             `json:"qos"`
	Retain                        bool                            `json:"retain"`
	StateTopic                    string                          `json:"state_topic"`
	StateFunc                     func() string                   `json:"-"`
	StateValueTemplate            string                          `json:"state_value_template"`
	TargetHumidityCommandTemplate string                          `json:"target_humidity_command_template"`
	TargetHumidityCommandTopic    string                          `json:"target_humidity_command_topic"`
	TargetHumidityCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	TargetHumidityStateTemplate   string                          `json:"target_humidity_state_template"`
	TargetHumidityStateTopic      string                          `json:"target_humidity_state_topic"`
	TargetHumidityStateFunc       func() string                   `json:"-"`
	UniqueId                      string                          `json:"unique_id"`
	MQTT                          MQTTFields                      `json:"-"`
}

func (d Humidifier) UpdateState() {
	if d.AvailabilityTopic != "" {
		state := d.AvailabilityFunc()
		if state != stateStore.Humidifier.Availability[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.AvailabilityTopic, qos, retain, state)
			stateStore.Humidifier.Availability[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.ModeStateTopic != "" {
		state := d.ModeStateFunc()
		if state != stateStore.Humidifier.ModeState[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.ModeStateTopic, qos, retain, state)
			stateStore.Humidifier.ModeState[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.StateTopic != "" {
		state := d.StateFunc()
		if state != stateStore.Humidifier.State[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.StateTopic, qos, retain, state)
			stateStore.Humidifier.State[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.TargetHumidityStateTopic != "" {
		state := d.TargetHumidityStateFunc()
		if state != stateStore.Humidifier.TargetHumidityState[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.TargetHumidityStateTopic, qos, retain, state)
			stateStore.Humidifier.TargetHumidityState[d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d Humidifier) Subscribe() {
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
	if d.ModeCommandTopic != "" {
		t := c.Subscribe(d.ModeCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TargetHumidityCommandTopic != "" {
		t := c.Subscribe(d.TargetHumidityCommandTopic, 0, d.MQTT.MessageHandler)
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
func (d Humidifier) UnSubscribe() {
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
	if d.ModeCommandTopic != "" {
		t := c.Unsubscribe(d.ModeCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TargetHumidityCommandTopic != "" {
		t := c.Unsubscribe(d.TargetHumidityCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d Humidifier) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, qos, retain, "online")
	token.Wait()
}
func (d Humidifier) Initialize() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d Humidifier) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.ModeCommandFunc != nil {
		d.ModeCommandTopic = GetTopic(d, "mode_command_topic")
		topicStore[d.ModeCommandTopic] = &d.ModeCommandFunc
	}
	if d.ModeStateFunc != nil {
		d.ModeStateTopic = GetTopic(d, "mode_state_topic")
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
	if d.TargetHumidityCommandFunc != nil {
		d.TargetHumidityCommandTopic = GetTopic(d, "target_humidity_command_topic")
		topicStore[d.TargetHumidityCommandTopic] = &d.TargetHumidityCommandFunc
	}
	if d.TargetHumidityStateFunc != nil {
		d.TargetHumidityStateTopic = GetTopic(d, "target_humidity_state_topic")
	}
}
