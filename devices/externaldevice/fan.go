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
func (d Fan) GetRawId() string {
	return "fan"
}
func (d Fan) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Fan) GetUniqueId() string {
	return d.UniqueId
}
func (d Fan) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Fan struct {
	AvailabilityMode     string                          `json:"availability_mode"`
	AvailabilityTemplate string                          `json:"availability_template"`
	AvailabilityTopic    string                          `json:"availability_topic"`
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
		Viadevice        string   `json:"viadevice"`
	} `json:"device"`
	EnabledByDefault           bool                            `json:"enabled_by_default"`
	Encoding                   string                          `json:"encoding"`
	EntityCategory             string                          `json:"entity_category"`
	Icon                       string                          `json:"icon"`
	Name                       string                          `json:"name"`
	ObjectId                   string                          `json:"object_id"`
	Optimistic                 bool                            `json:"optimistic"`
	OscillationCommandTemplate string                          `json:"oscillation_command_template"`
	OscillationCommandTopic    string                          `json:"oscillation_command_topic"`
	OscillationCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	OscillationStateTopic      string                          `json:"oscillation_state_topic"`
	OscillationStateFunc       func() string                   `json:"-"`
	OscillationValueTemplate   string                          `json:"oscillation_value_template"`
	PayloadAvailable           string                          `json:"payload_available"`
	PayloadNotAvailable        string                          `json:"payload_not_available"`
	PayloadOff                 string                          `json:"payload_off"`
	PayloadOn                  string                          `json:"payload_on"`
	PayloadOscillationOff      string                          `json:"payload_oscillation_off"`
	PayloadOscillationOn       string                          `json:"payload_oscillation_on"`
	PayloadResetPercentage     string                          `json:"payload_reset_percentage"`
	PayloadResetPresetMode     string                          `json:"payload_reset_preset_mode"`
	PercentageCommandTemplate  string                          `json:"percentage_command_template"`
	PercentageCommandTopic     string                          `json:"percentage_command_topic"`
	PercentageCommandFunc      func(mqtt.Message, mqtt.Client) `json:"-"`
	PercentageStateTopic       string                          `json:"percentage_state_topic"`
	PercentageStateFunc        func() string                   `json:"-"`
	PercentageValueTemplate    string                          `json:"percentage_value_template"`
	PresetModeCommandTemplate  string                          `json:"preset_mode_command_template"`
	PresetModeCommandTopic     string                          `json:"preset_mode_command_topic"`
	PresetModeCommandFunc      func(mqtt.Message, mqtt.Client) `json:"-"`
	PresetModeStateTopic       string                          `json:"preset_mode_state_topic"`
	PresetModeStateFunc        func() string                   `json:"-"`
	PresetModeValueTemplate    string                          `json:"preset_mode_value_template"`
	PresetModes                []string                        `json:"preset_modes"`
	Qos                        int                             `json:"qos"`
	Retain                     bool                            `json:"retain"`
	SpeedRangeMax              int                             `json:"speed_range_max"`
	SpeedRangeMin              int                             `json:"speed_range_min"`
	StateTopic                 string                          `json:"state_topic"`
	StateFunc                  func() string                   `json:"-"`
	StateValueTemplate         string                          `json:"state_value_template"`
	UniqueId                   string                          `json:"unique_id"`
	MQTT                       MQTTFields                      `json:"-"`
}

func (d Fan) UpdateState() {
	if d.OscillationStateTopic != "" {
		state := d.OscillationStateFunc()
		if state != stateStore.Fan.OscillationState[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.OscillationStateTopic, common.QoS, common.Retain, state)
			stateStore.Fan.OscillationState[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.PercentageStateTopic != "" {
		state := d.PercentageStateFunc()
		if state != stateStore.Fan.PercentageState[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.PercentageStateTopic, common.QoS, common.Retain, state)
			stateStore.Fan.PercentageState[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.PresetModeStateTopic != "" {
		state := d.PresetModeStateFunc()
		if state != stateStore.Fan.PresetModeState[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.PresetModeStateTopic, common.QoS, common.Retain, state)
			stateStore.Fan.PresetModeState[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.StateTopic != "" {
		state := d.StateFunc()
		if state != stateStore.Fan.State[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.StateTopic, common.QoS, common.Retain, state)
			stateStore.Fan.State[d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d Fan) Subscribe() {
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
	if d.OscillationCommandTopic != "" {
		t := c.Subscribe(d.OscillationCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.PercentageCommandTopic != "" {
		t := c.Subscribe(d.PercentageCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.PresetModeCommandTopic != "" {
		t := c.Subscribe(d.PresetModeCommandTopic, 0, d.MQTT.MessageHandler)
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
func (d Fan) UnSubscribe() {
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
	if d.OscillationCommandTopic != "" {
		t := c.Unsubscribe(d.OscillationCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.PercentageCommandTopic != "" {
		t := c.Unsubscribe(d.PercentageCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.PresetModeCommandTopic != "" {
		t := c.Unsubscribe(d.PresetModeCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d Fan) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d Fan) Initialize() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d Fan) PopulateTopics() {
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		store.TopicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.OscillationCommandFunc != nil {
		d.OscillationCommandTopic = GetTopic(d, "oscillation_command_topic")
		store.TopicStore[d.OscillationCommandTopic] = &d.OscillationCommandFunc
	}
	if d.OscillationStateFunc != nil {
		d.OscillationStateTopic = GetTopic(d, "oscillation_state_topic")
	}
	if d.PercentageCommandFunc != nil {
		d.PercentageCommandTopic = GetTopic(d, "percentage_command_topic")
		store.TopicStore[d.PercentageCommandTopic] = &d.PercentageCommandFunc
	}
	if d.PercentageStateFunc != nil {
		d.PercentageStateTopic = GetTopic(d, "percentage_state_topic")
	}
	if d.PresetModeCommandFunc != nil {
		d.PresetModeCommandTopic = GetTopic(d, "preset_mode_command_topic")
		store.TopicStore[d.PresetModeCommandTopic] = &d.PresetModeCommandFunc
	}
	if d.PresetModeStateFunc != nil {
		d.PresetModeStateTopic = GetTopic(d, "preset_mode_state_topic")
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
