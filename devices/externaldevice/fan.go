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
func (d *Fan) GetRawId() string {
	return "fan"
}
func (d *Fan) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Fan) GetUniqueId() string {
	return d.UniqueId
}
func (d *Fan) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Fan struct {
	AvailabilityMode     string                          `json:"availability_mode"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate string                          `json:"availability_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    string                          `json:"availability_topic"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	CommandTemplate      string                          `json:"command_template"`      // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `command_topic`."
	CommandTopic         string                          `json:"command_topic"`         // "The MQTT topic to publish commands to change the fan state."
	CommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
	Device               struct {
		ConfigurationUrl string `json:"configuration_url"` // "A link to the webpage that can manage the configuration of this device. Can be either an HTTP or HTTPS link."
		Connections      string `json:"connections"`       // "A list of connections of the device to the outside world as a list of tuples `[connection_type, connection_identifier]`. For example the MAC address of a network interface: `\"connections\": [[\"mac\", \"02:5b:26:a8:dc:12\"]]`."
		Identifiers      string `json:"identifiers"`       // "A list of IDs that uniquely identify the device. For example a serial number."
		Manufacturer     string `json:"manufacturer"`      // "The manufacturer of the device."
		Model            string `json:"model"`             // "The model of the device."
		Name             string `json:"name"`              // "The name of the device."
		SuggestedArea    string `json:"suggested_area"`    // "Suggest an area if the device isn’t in one yet."
		SwVersion        string `json:"sw_version"`        // "The firmware version of the device."
		Viadevice        string `json:"viadevice"`         // null
	} `json:"device"`
	EnabledByDefault           bool                            `json:"enabled_by_default"`           // "Flag which defines if the entity should be enabled when first added."
	Encoding                   string                          `json:"encoding"`                     // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory             string                          `json:"entity_category"`              // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                       string                          `json:"icon"`                         // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Name                       string                          `json:"name"`                         // "The name of the fan."
	ObjectId                   string                          `json:"object_id"`                    // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic                 bool                            `json:"optimistic"`                   // "Flag that defines if fan works in optimistic mode"
	OscillationCommandTemplate string                          `json:"oscillation_command_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `oscillation_command_topic`."
	OscillationCommandTopic    string                          `json:"oscillation_command_topic"`    // "The MQTT topic to publish commands to change the oscillation state."
	OscillationCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	OscillationStateTopic      string                          `json:"oscillation_state_topic"` // "The MQTT topic subscribed to receive oscillation state updates."
	OscillationStateFunc       func() string                   `json:"-"`
	OscillationValueTemplate   string                          `json:"oscillation_value_template"`  // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract a value from the oscillation."
	PayloadAvailable           string                          `json:"payload_available"`           // "The payload that represents the available state."
	PayloadNotAvailable        string                          `json:"payload_not_available"`       // "The payload that represents the unavailable state."
	PayloadOff                 string                          `json:"payload_off"`                 // "The payload that represents the stop state."
	PayloadOn                  string                          `json:"payload_on"`                  // "The payload that represents the running state."
	PayloadOscillationOff      string                          `json:"payload_oscillation_off"`     // "The payload that represents the oscillation off state."
	PayloadOscillationOn       string                          `json:"payload_oscillation_on"`      // "The payload that represents the oscillation on state."
	PayloadResetPercentage     string                          `json:"payload_reset_percentage"`    // "A special payload that resets the `percentage` state attribute to `None` when received at the `percentage_state_topic`."
	PayloadResetPresetMode     string                          `json:"payload_reset_preset_mode"`   // "A special payload that resets the `preset_mode` state attribute to `None` when received at the `preset_mode_state_topic`."
	PercentageCommandTemplate  string                          `json:"percentage_command_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `percentage_command_topic`."
	PercentageCommandTopic     string                          `json:"percentage_command_topic"`    // "The MQTT topic to publish commands to change the fan speed state based on a percentage."
	PercentageCommandFunc      func(mqtt.Message, mqtt.Client) `json:"-"`
	PercentageStateTopic       string                          `json:"percentage_state_topic"` // "The MQTT topic subscribed to receive fan speed based on percentage."
	PercentageStateFunc        func() string                   `json:"-"`
	PercentageValueTemplate    string                          `json:"percentage_value_template"`    // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the `percentage` value from the payload received on `percentage_state_topic`."
	PresetModeCommandTemplate  string                          `json:"preset_mode_command_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `preset_mode_command_topic`."
	PresetModeCommandTopic     string                          `json:"preset_mode_command_topic"`    // "The MQTT topic to publish commands to change the preset mode."
	PresetModeCommandFunc      func(mqtt.Message, mqtt.Client) `json:"-"`
	PresetModeStateTopic       string                          `json:"preset_mode_state_topic"` // "The MQTT topic subscribed to receive fan speed based on presets."
	PresetModeStateFunc        func() string                   `json:"-"`
	PresetModeValueTemplate    string                          `json:"preset_mode_value_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the `preset_mode` value from the payload received on `preset_mode_state_topic`."
	PresetModes                []string                        `json:"preset_modes"`               // "List of preset modes this fan is capable of running at. Common examples include `auto`, `smart`, `whoosh`, `eco` and `breeze`."
	Qos                        int                             `json:"qos"`                        // "The maximum QoS level of the state topic."
	Retain                     bool                            `json:"retain"`                     // "If the published message should have the retain flag on or not."
	SpeedRangeMax              int                             `json:"speed_range_max"`            // "The maximum of numeric output range (representing 100 %). The number of speeds within the `speed_range` / `100` will determine the `percentage_step`."
	SpeedRangeMin              int                             `json:"speed_range_min"`            // "The minimum of numeric output range (`off` not included, so `speed_range_min` - `1` represents 0 %). The number of speeds within the speed_range / 100 will determine the `percentage_step`."
	StateTopic                 string                          `json:"state_topic"`                // "The MQTT topic subscribed to receive state updates."
	StateFunc                  func() string                   `json:"-"`
	StateValueTemplate         string                          `json:"state_value_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract a value from the state."
	UniqueId                   string                          `json:"unique_id"`            // "An ID that uniquely identifies this fan. If two fans have the same unique ID, Home Assistant will raise an exception."
	MQTT                       MQTTFields                      `json:"-"`
}

func (d *Fan) UpdateState() {
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
func (d *Fan) Subscribe() {
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
func (d *Fan) UnSubscribe() {
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
func (d *Fan) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d *Fan) Initialize() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d *Fan) PopulateTopics() {
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
func (d *Fan) SetMQTTFields(fields MQTTFields) {
	d.MQTT = fields
}
func (d *Fan) GetMQTTFields() (fields MQTTFields) {
	return d.MQTT
}
