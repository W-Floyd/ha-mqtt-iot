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
func (d *AlarmControlPanel) GetRawId() string {
	return "alarm_control_panel"
}
func (d *AlarmControlPanel) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *AlarmControlPanel) GetUniqueId() string {
	return *d.UniqueId
}
func (d *AlarmControlPanel) PopulateDevice() {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &common.MachineID
}

type AlarmControlPanel struct {
	AvailabilityMode     *string                         `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate *string                         `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    *string                         `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	AvailabilityFunc     func() string                   `json:"-"`
	Code                 *string                         `json:"code,omitempty"`                  // "If defined, specifies a code to enable or disable the alarm in the frontend. Note that the code is validated locally and blocks sending MQTT messages to the remote device. For remote code validation, the code can be configured to either of the special values `REMOTE_CODE` (numeric code) or `REMOTE_CODE_TEXT` (text code). In this case, local code validation is bypassed but the frontend will still show a numeric or text code dialog. Use `command_template` to send the code to the remote device. Example configurations for remote code validation [can be found here](./#configurations-with-remote-code-validation)."
	CodeArmRequired      *bool                           `json:"code_arm_required,omitempty"`     // "If true the code is required to arm the alarm. If false the code is not validated."
	CodeDisarmRequired   *bool                           `json:"code_disarm_required,omitempty"`  // "If true the code is required to disarm the alarm. If false the code is not validated."
	CodeTriggerRequired  *bool                           `json:"code_trigger_required,omitempty"` // "If true the code is required to trigger the alarm. If false the code is not validated."
	CommandTemplate      *string                         `json:"command_template,omitempty"`      // "The [template](/docs/configuration/templating/#processing-incoming-data) used for the command payload. Available variables: `action` and `code`."
	CommandTopic         *string                         `json:"command_topic,omitempty"`         // "The MQTT topic to publish commands to change the alarm state."
	CommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
	Device               struct {
		ConfigurationUrl *string `json:"configuration_url,omitempty"` // "A link to the webpage that can manage the configuration of this device. Can be either an HTTP or HTTPS link."
		Connections      *string `json:"connections,omitempty"`       // "A list of connections of the device to the outside world as a list of tuples `[connection_type, connection_identifier]`. For example the MAC address of a network interface: `\"connections\": [[\"mac\", \"02:5b:26:a8:dc:12\"]]`."
		Identifiers      *string `json:"identifiers,omitempty"`       // "A list of IDs that uniquely identify the device. For example a serial number."
		Manufacturer     *string `json:"manufacturer,omitempty"`      // "The manufacturer of the device."
		Model            *string `json:"model,omitempty"`             // "The model of the device."
		Name             *string `json:"name,omitempty"`              // "The name of the device."
		SuggestedArea    *string `json:"suggested_area,omitempty"`    // "Suggest an area if the device isn’t in one yet."
		SwVersion        *string `json:"sw_version,omitempty"`        // "The firmware version of the device."
		Viadevice        *string `json:"viadevice,omitempty"`         // null
	} `json:"device,omitempty"`
	EnabledByDefault       *bool         `json:"enabled_by_default,omitempty"`        // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string       `json:"encoding,omitempty"`                  // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string       `json:"entity_category,omitempty"`           // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   *string       `json:"icon,omitempty"`                      // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Name                   *string       `json:"name,omitempty"`                      // "The name of the alarm."
	ObjectId               *string       `json:"object_id,omitempty"`                 // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadArmAway         *string       `json:"payload_arm_away,omitempty"`          // "The payload to set armed-away mode on your Alarm Panel."
	PayloadArmCustomBypass *string       `json:"payload_arm_custom_bypass,omitempty"` // "The payload to set armed-custom-bypass mode on your Alarm Panel."
	PayloadArmHome         *string       `json:"payload_arm_home,omitempty"`          // "The payload to set armed-home mode on your Alarm Panel."
	PayloadArmNight        *string       `json:"payload_arm_night,omitempty"`         // "The payload to set armed-night mode on your Alarm Panel."
	PayloadArmVacation     *string       `json:"payload_arm_vacation,omitempty"`      // "The payload to set armed-vacation mode on your Alarm Panel."
	PayloadAvailable       *string       `json:"payload_available,omitempty"`         // "The payload that represents the available state."
	PayloadDisarm          *string       `json:"payload_disarm,omitempty"`            // "The payload to disarm your Alarm Panel."
	PayloadNotAvailable    *string       `json:"payload_not_available,omitempty"`     // "The payload that represents the unavailable state."
	PayloadTrigger         *string       `json:"payload_trigger,omitempty"`           // "The payload to trigger the alarm on your Alarm Panel."
	Qos                    *int          `json:"qos,omitempty"`                       // "The maximum QoS level of the state topic."
	Retain                 *bool         `json:"retain,omitempty"`                    // "If the published message should have the retain flag on or not."
	StateTopic             *string       `json:"state_topic,omitempty"`               // "The MQTT topic subscribed to receive state updates."
	StateFunc              func() string `json:"-"`
	UniqueId               *string       `json:"unique_id,omitempty"`      // "An ID that uniquely identifies this alarm panel. If two alarm panels have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate          *string       `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the value."
	MQTT                   *MQTTFields   `json:"-"`
}

func (d *AlarmControlPanel) UpdateState() {
	if d.AvailabilityTopic != nil {
		state := d.AvailabilityFunc()
		if state != stateStore.AlarmControlPanel.Availability[*d.UniqueId] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			c := *d.MQTT.Client
			token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, state)
			stateStore.AlarmControlPanel.Availability[*d.UniqueId] = state
			token.Wait()
		}
	}
	if d.StateTopic != nil {
		state := d.StateFunc()
		if state != stateStore.AlarmControlPanel.State[*d.UniqueId] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			c := *d.MQTT.Client
			token := c.Publish(*d.StateTopic, common.QoS, common.Retain, state)
			stateStore.AlarmControlPanel.State[*d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d *AlarmControlPanel) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	if *d.CommandTopic != "" {
		t := c.Subscribe(*d.CommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	token := c.Publish(GetDiscoveryTopic(d), 0, true, message)
	token.Wait()
	time.Sleep(common.HADiscoveryDelay)
	d.AvailabilityFunc()
	d.UpdateState()
}
func (d *AlarmControlPanel) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "offline")
	token.Wait()
	if *d.CommandTopic != "" {
		t := c.Unsubscribe(*d.CommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d *AlarmControlPanel) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d *AlarmControlPanel) Initialize() {
	*d.Retain = false
	d.PopulateDevice()
	d.AddMessageHandler()
	d.PopulateTopics()
}
func (d *AlarmControlPanel) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = new(string)
		*d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = new(string)
		*d.CommandTopic = GetTopic(d, "command_topic")
		store.TopicStore[*d.CommandTopic] = &d.CommandFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = new(string)
		*d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d *AlarmControlPanel) SetMQTTFields(fields MQTTFields) {
	d.MQTT = &fields
}
func (d *AlarmControlPanel) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
