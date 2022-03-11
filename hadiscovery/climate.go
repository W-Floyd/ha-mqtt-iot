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
func (d Climate) GetRawId() string {
	return "climate"
}
func (d Climate) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Climate) GetUniqueId() string {
	return d.UniqueId
}
func (d Climate) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Climate struct {
	ActionTemplate             string                          `json:"action_template"`
	ActionTopic                string                          `json:"action_topic"`
	ActionFunc                 func(mqtt.Message, mqtt.Client) `json:"-"`
	AuxCommandTopic            string                          `json:"aux_command_topic"`
	AuxCommandFunc             func(mqtt.Message, mqtt.Client) `json:"-"`
	AuxStateTemplate           string                          `json:"aux_state_template"`
	AuxStateTopic              string                          `json:"aux_state_topic"`
	AuxStateFunc               func() string                   `json:"-"`
	AvailabilityMode           string                          `json:"availability_mode"`
	AvailabilityTemplate       string                          `json:"availability_template"`
	AvailabilityTopic          string                          `json:"availability_topic"`
	AvailabilityFunc           func() string                   `json:"-"`
	CurrentTemperatureTemplate string                          `json:"current_temperature_template"`
	CurrentTemperatureTopic    string                          `json:"current_temperature_topic"`
	CurrentTemperatureFunc     func() string                   `json:"-"`
	Device                     struct {
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
	EnabledByDefault               bool                            `json:"enabled_by_default"`
	Encoding                       string                          `json:"encoding"`
	EntityCategory                 string                          `json:"entity_category"`
	FanModeCommandTemplate         string                          `json:"fan_mode_command_template"`
	FanModeCommandTopic            string                          `json:"fan_mode_command_topic"`
	FanModeCommandFunc             func(mqtt.Message, mqtt.Client) `json:"-"`
	FanModeStateTemplate           string                          `json:"fan_mode_state_template"`
	FanModeStateTopic              string                          `json:"fan_mode_state_topic"`
	FanModeStateFunc               func() string                   `json:"-"`
	FanModes                       []string                        `json:"fan_modes"`
	Icon                           string                          `json:"icon"`
	Initial                        int                             `json:"initial"`
	MaxTemp                        float64                         `json:"max_temp"`
	MinTemp                        float64                         `json:"min_temp"`
	ModeCommandTemplate            string                          `json:"mode_command_template"`
	ModeCommandTopic               string                          `json:"mode_command_topic"`
	ModeCommandFunc                func(mqtt.Message, mqtt.Client) `json:"-"`
	ModeStateTemplate              string                          `json:"mode_state_template"`
	ModeStateTopic                 string                          `json:"mode_state_topic"`
	ModeStateFunc                  func() string                   `json:"-"`
	Modes                          []string                        `json:"modes"`
	Name                           string                          `json:"name"`
	ObjectId                       string                          `json:"object_id"`
	PayloadAvailable               string                          `json:"payload_available"`
	PayloadNotAvailable            string                          `json:"payload_not_available"`
	PayloadOff                     string                          `json:"payload_off"`
	PayloadOn                      string                          `json:"payload_on"`
	PowerCommandTopic              string                          `json:"power_command_topic"`
	PowerCommandFunc               func(mqtt.Message, mqtt.Client) `json:"-"`
	Precision                      float64                         `json:"precision"`
	PresetModeCommandTemplate      string                          `json:"preset_mode_command_template"`
	PresetModeCommandTopic         string                          `json:"preset_mode_command_topic"`
	PresetModeCommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
	PresetModeStateTopic           string                          `json:"preset_mode_state_topic"`
	PresetModeStateFunc            func() string                   `json:"-"`
	PresetModeValueTemplate        string                          `json:"preset_mode_value_template"`
	PresetModes                    []string                        `json:"preset_modes"`
	Qos                            int                             `json:"qos"`
	Retain                         bool                            `json:"retain"`
	SwingModeCommandTemplate       string                          `json:"swing_mode_command_template"`
	SwingModeCommandTopic          string                          `json:"swing_mode_command_topic"`
	SwingModeCommandFunc           func(mqtt.Message, mqtt.Client) `json:"-"`
	SwingModeStateTemplate         string                          `json:"swing_mode_state_template"`
	SwingModeStateTopic            string                          `json:"swing_mode_state_topic"`
	SwingModeStateFunc             func() string                   `json:"-"`
	SwingModes                     []string                        `json:"swing_modes"`
	TempStep                       float64                         `json:"temp_step"`
	TemperatureCommandTemplate     string                          `json:"temperature_command_template"`
	TemperatureCommandTopic        string                          `json:"temperature_command_topic"`
	TemperatureCommandFunc         func(mqtt.Message, mqtt.Client) `json:"-"`
	TemperatureHighCommandTemplate string                          `json:"temperature_high_command_template"`
	TemperatureHighCommandTopic    string                          `json:"temperature_high_command_topic"`
	TemperatureHighCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	TemperatureHighStateTemplate   string                          `json:"temperature_high_state_template"`
	TemperatureHighStateTopic      string                          `json:"temperature_high_state_topic"`
	TemperatureHighStateFunc       func() string                   `json:"-"`
	TemperatureLowCommandTemplate  string                          `json:"temperature_low_command_template"`
	TemperatureLowCommandTopic     string                          `json:"temperature_low_command_topic"`
	TemperatureLowCommandFunc      func(mqtt.Message, mqtt.Client) `json:"-"`
	TemperatureLowStateTemplate    string                          `json:"temperature_low_state_template"`
	TemperatureLowStateTopic       string                          `json:"temperature_low_state_topic"`
	TemperatureLowStateFunc        func() string                   `json:"-"`
	TemperatureStateTemplate       string                          `json:"temperature_state_template"`
	TemperatureStateTopic          string                          `json:"temperature_state_topic"`
	TemperatureStateFunc           func() string                   `json:"-"`
	TemperatureUnit                string                          `json:"temperature_unit"`
	UniqueId                       string                          `json:"unique_id"`
	ValueTemplate                  string                          `json:"value_template"`
	MQTT                           MQTTFields                      `json:"-"`
}

func (d Climate) UpdateState() {
	if d.AuxStateTopic != "" {
		state := d.AuxStateFunc()
		if state != stateStore.Climate.AuxState[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.AuxStateTopic, qos, retain, state)
			stateStore.Climate.AuxState[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.AvailabilityTopic != "" {
		state := d.AvailabilityFunc()
		if state != stateStore.Climate.Availability[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.AvailabilityTopic, qos, retain, state)
			stateStore.Climate.Availability[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.CurrentTemperatureTopic != "" {
		state := d.CurrentTemperatureFunc()
		if state != stateStore.Climate.CurrentTemperature[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.CurrentTemperatureTopic, qos, retain, state)
			stateStore.Climate.CurrentTemperature[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.FanModeStateTopic != "" {
		state := d.FanModeStateFunc()
		if state != stateStore.Climate.FanModeState[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.FanModeStateTopic, qos, retain, state)
			stateStore.Climate.FanModeState[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.ModeStateTopic != "" {
		state := d.ModeStateFunc()
		if state != stateStore.Climate.ModeState[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.ModeStateTopic, qos, retain, state)
			stateStore.Climate.ModeState[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.PresetModeStateTopic != "" {
		state := d.PresetModeStateFunc()
		if state != stateStore.Climate.PresetModeState[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.PresetModeStateTopic, qos, retain, state)
			stateStore.Climate.PresetModeState[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.SwingModeStateTopic != "" {
		state := d.SwingModeStateFunc()
		if state != stateStore.Climate.SwingModeState[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.SwingModeStateTopic, qos, retain, state)
			stateStore.Climate.SwingModeState[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.TemperatureHighStateTopic != "" {
		state := d.TemperatureHighStateFunc()
		if state != stateStore.Climate.TemperatureHighState[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.TemperatureHighStateTopic, qos, retain, state)
			stateStore.Climate.TemperatureHighState[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.TemperatureLowStateTopic != "" {
		state := d.TemperatureLowStateFunc()
		if state != stateStore.Climate.TemperatureLowState[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.TemperatureLowStateTopic, qos, retain, state)
			stateStore.Climate.TemperatureLowState[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.TemperatureStateTopic != "" {
		state := d.TemperatureStateFunc()
		if state != stateStore.Climate.TemperatureState[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.TemperatureStateTopic, qos, retain, state)
			stateStore.Climate.TemperatureState[d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d Climate) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	if d.ActionTopic != "" {
		t := c.Subscribe(d.ActionTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.AuxCommandTopic != "" {
		t := c.Subscribe(d.AuxCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.FanModeCommandTopic != "" {
		t := c.Subscribe(d.FanModeCommandTopic, 0, d.MQTT.MessageHandler)
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
	if d.PowerCommandTopic != "" {
		t := c.Subscribe(d.PowerCommandTopic, 0, d.MQTT.MessageHandler)
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
	if d.SwingModeCommandTopic != "" {
		t := c.Subscribe(d.SwingModeCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TemperatureCommandTopic != "" {
		t := c.Subscribe(d.TemperatureCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TemperatureHighCommandTopic != "" {
		t := c.Subscribe(d.TemperatureHighCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TemperatureLowCommandTopic != "" {
		t := c.Subscribe(d.TemperatureLowCommandTopic, 0, d.MQTT.MessageHandler)
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
func (d Climate) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, qos, retain, "offline")
	token.Wait()
	if d.ActionTopic != "" {
		t := c.Unsubscribe(d.ActionTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.AuxCommandTopic != "" {
		t := c.Unsubscribe(d.AuxCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.FanModeCommandTopic != "" {
		t := c.Unsubscribe(d.FanModeCommandTopic)
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
	if d.PowerCommandTopic != "" {
		t := c.Unsubscribe(d.PowerCommandTopic)
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
	if d.SwingModeCommandTopic != "" {
		t := c.Unsubscribe(d.SwingModeCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TemperatureCommandTopic != "" {
		t := c.Unsubscribe(d.TemperatureCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TemperatureHighCommandTopic != "" {
		t := c.Unsubscribe(d.TemperatureHighCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TemperatureLowCommandTopic != "" {
		t := c.Unsubscribe(d.TemperatureLowCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d Climate) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, qos, retain, "online")
	token.Wait()
}
func (d Climate) Initialize() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d Climate) PopulateTopics() {
	if d.ActionFunc != nil {
		d.ActionTopic = GetTopic(d, "action_topic")
		topicStore[d.ActionTopic] = &d.ActionFunc
	}
	if d.AuxCommandFunc != nil {
		d.AuxCommandTopic = GetTopic(d, "aux_command_topic")
		topicStore[d.AuxCommandTopic] = &d.AuxCommandFunc
	}
	if d.AuxStateFunc != nil {
		d.AuxStateTopic = GetTopic(d, "aux_state_topic")
	}
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CurrentTemperatureFunc != nil {
		d.CurrentTemperatureTopic = GetTopic(d, "current_temperature_topic")
	}
	if d.FanModeCommandFunc != nil {
		d.FanModeCommandTopic = GetTopic(d, "fan_mode_command_topic")
		topicStore[d.FanModeCommandTopic] = &d.FanModeCommandFunc
	}
	if d.FanModeStateFunc != nil {
		d.FanModeStateTopic = GetTopic(d, "fan_mode_state_topic")
	}
	if d.ModeCommandFunc != nil {
		d.ModeCommandTopic = GetTopic(d, "mode_command_topic")
		topicStore[d.ModeCommandTopic] = &d.ModeCommandFunc
	}
	if d.ModeStateFunc != nil {
		d.ModeStateTopic = GetTopic(d, "mode_state_topic")
	}
	if d.PowerCommandFunc != nil {
		d.PowerCommandTopic = GetTopic(d, "power_command_topic")
		topicStore[d.PowerCommandTopic] = &d.PowerCommandFunc
	}
	if d.PresetModeCommandFunc != nil {
		d.PresetModeCommandTopic = GetTopic(d, "preset_mode_command_topic")
		topicStore[d.PresetModeCommandTopic] = &d.PresetModeCommandFunc
	}
	if d.PresetModeStateFunc != nil {
		d.PresetModeStateTopic = GetTopic(d, "preset_mode_state_topic")
	}
	if d.SwingModeCommandFunc != nil {
		d.SwingModeCommandTopic = GetTopic(d, "swing_mode_command_topic")
		topicStore[d.SwingModeCommandTopic] = &d.SwingModeCommandFunc
	}
	if d.SwingModeStateFunc != nil {
		d.SwingModeStateTopic = GetTopic(d, "swing_mode_state_topic")
	}
	if d.TemperatureCommandFunc != nil {
		d.TemperatureCommandTopic = GetTopic(d, "temperature_command_topic")
		topicStore[d.TemperatureCommandTopic] = &d.TemperatureCommandFunc
	}
	if d.TemperatureHighCommandFunc != nil {
		d.TemperatureHighCommandTopic = GetTopic(d, "temperature_high_command_topic")
		topicStore[d.TemperatureHighCommandTopic] = &d.TemperatureHighCommandFunc
	}
	if d.TemperatureHighStateFunc != nil {
		d.TemperatureHighStateTopic = GetTopic(d, "temperature_high_state_topic")
	}
	if d.TemperatureLowCommandFunc != nil {
		d.TemperatureLowCommandTopic = GetTopic(d, "temperature_low_command_topic")
		topicStore[d.TemperatureLowCommandTopic] = &d.TemperatureLowCommandFunc
	}
	if d.TemperatureLowStateFunc != nil {
		d.TemperatureLowStateTopic = GetTopic(d, "temperature_low_state_topic")
	}
	if d.TemperatureStateFunc != nil {
		d.TemperatureStateTopic = GetTopic(d, "temperature_state_topic")
	}
}
