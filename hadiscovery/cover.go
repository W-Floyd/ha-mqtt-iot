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
func (d Cover) GetRawId() string {
	return "cover"
}
func (d Cover) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Cover) GetUniqueId() string {
	return d.UniqueId
}
func (d Cover) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Cover struct {
	AvailabilityMode     string                          `json:"availability_mode"`
	AvailabilityTemplate string                          `json:"availability_template"`
	AvailabilityTopic    string                          `json:"availability_topic"`
	AvailabilityFunc     func() string                   `json:"-"`
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
	DeviceClass         string                          `json:"device_class"`
	EnabledByDefault    bool                            `json:"enabled_by_default"`
	Encoding            string                          `json:"encoding"`
	EntityCategory      string                          `json:"entity_category"`
	Icon                string                          `json:"icon"`
	Name                string                          `json:"name"`
	ObjectId            string                          `json:"object_id"`
	Optimistic          bool                            `json:"optimistic"`
	PayloadAvailable    string                          `json:"payload_available"`
	PayloadClose        string                          `json:"payload_close"`
	PayloadNotAvailable string                          `json:"payload_not_available"`
	PayloadOpen         string                          `json:"payload_open"`
	PayloadStop         string                          `json:"payload_stop"`
	PositionClosed      int                             `json:"position_closed"`
	PositionOpen        int                             `json:"position_open"`
	PositionTemplate    string                          `json:"position_template"`
	PositionTopic       string                          `json:"position_topic"`
	PositionFunc        func() string                   `json:"-"`
	Qos                 int                             `json:"qos"`
	Retain              bool                            `json:"retain"`
	SetPositionTemplate string                          `json:"set_position_template"`
	SetPositionTopic    string                          `json:"set_position_topic"`
	SetPositionFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	StateClosed         string                          `json:"state_closed"`
	StateClosing        string                          `json:"state_closing"`
	StateOpen           string                          `json:"state_open"`
	StateOpening        string                          `json:"state_opening"`
	StateStopped        string                          `json:"state_stopped"`
	StateTopic          string                          `json:"state_topic"`
	StateFunc           func() string                   `json:"-"`
	TiltClosedValue     int                             `json:"tilt_closed_value"`
	TiltCommandTemplate string                          `json:"tilt_command_template"`
	TiltCommandTopic    string                          `json:"tilt_command_topic"`
	TiltCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	TiltMax             int                             `json:"tilt_max"`
	TiltMin             int                             `json:"tilt_min"`
	TiltOpenedValue     int                             `json:"tilt_opened_value"`
	TiltOptimistic      bool                            `json:"tilt_optimistic"`
	TiltStatusTemplate  string                          `json:"tilt_status_template"`
	TiltStatusTopic     string                          `json:"tilt_status_topic"`
	TiltStatusFunc      func() string                   `json:"-"`
	UniqueId            string                          `json:"unique_id"`
	ValueTemplate       string                          `json:"value_template"`
	MQTT                MQTTFields                      `json:"-"`
}

func (d Cover) UpdateState() {
	if d.AvailabilityTopic != "" {
		state := d.AvailabilityFunc()
		if state != stateStore.Cover.Availability[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.AvailabilityTopic, qos, retain, state)
			stateStore.Cover.Availability[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.PositionTopic != "" {
		state := d.PositionFunc()
		if state != stateStore.Cover.Position[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.PositionTopic, qos, retain, state)
			stateStore.Cover.Position[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.StateTopic != "" {
		state := d.StateFunc()
		if state != stateStore.Cover.State[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.StateTopic, qos, retain, state)
			stateStore.Cover.State[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.TiltStatusTopic != "" {
		state := d.TiltStatusFunc()
		if state != stateStore.Cover.TiltStatus[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.TiltStatusTopic, qos, retain, state)
			stateStore.Cover.TiltStatus[d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d Cover) Subscribe() {
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
	if d.SetPositionTopic != "" {
		t := c.Subscribe(d.SetPositionTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TiltCommandTopic != "" {
		t := c.Subscribe(d.TiltCommandTopic, 0, d.MQTT.MessageHandler)
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
func (d Cover) UnSubscribe() {
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
	if d.SetPositionTopic != "" {
		t := c.Unsubscribe(d.SetPositionTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TiltCommandTopic != "" {
		t := c.Unsubscribe(d.TiltCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d Cover) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, qos, retain, "online")
	token.Wait()
}
func (d Cover) Initialize() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d Cover) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.PositionFunc != nil {
		d.PositionTopic = GetTopic(d, "position_topic")
	}
	if d.SetPositionFunc != nil {
		d.SetPositionTopic = GetTopic(d, "set_position_topic")
		topicStore[d.SetPositionTopic] = &d.SetPositionFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
	if d.TiltCommandFunc != nil {
		d.TiltCommandTopic = GetTopic(d, "tilt_command_topic")
		topicStore[d.TiltCommandTopic] = &d.TiltCommandFunc
	}
	if d.TiltStatusFunc != nil {
		d.TiltStatusTopic = GetTopic(d, "tilt_status_topic")
	}
}
