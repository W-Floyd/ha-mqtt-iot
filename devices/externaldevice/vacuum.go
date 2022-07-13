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
func (d Vacuum) GetRawId() string {
	return "vacuum"
}
func (d Vacuum) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Vacuum) GetUniqueId() string {
	return d.UniqueId
}
func (d Vacuum) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Vacuum struct {
	AvailabilityMode     string                          `json:"availability_mode"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate string                          `json:"availability_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    string                          `json:"availability_topic"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	CommandTopic         string                          `json:"command_topic"`         // "The MQTT topic to publish commands to control the vacuum."
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
	Encoding            string                          `json:"encoding"`               // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	FanSpeedList        []string                        `json:"fan_speed_list"`         // "List of possible fan speeds for the vacuum."
	Name                string                          `json:"name"`                   // "The name of the vacuum."
	ObjectId            string                          `json:"object_id"`              // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable    string                          `json:"payload_available"`      // "The payload that represents the available state."
	PayloadCleanSpot    string                          `json:"payload_clean_spot"`     // "The payload to send to the `command_topic` to begin a spot cleaning cycle."
	PayloadLocate       string                          `json:"payload_locate"`         // "The payload to send to the `command_topic` to locate the vacuum (typically plays a song)."
	PayloadNotAvailable string                          `json:"payload_not_available"`  // "The payload that represents the unavailable state."
	PayloadPause        string                          `json:"payload_pause"`          // "The payload to send to the `command_topic` to pause the vacuum."
	PayloadReturnToBase string                          `json:"payload_return_to_base"` // "The payload to send to the `command_topic` to tell the vacuum to return to base."
	PayloadStart        string                          `json:"payload_start"`          // "The payload to send to the `command_topic` to begin the cleaning cycle."
	PayloadStop         string                          `json:"payload_stop"`           // "The payload to send to the `command_topic` to stop cleaning."
	Qos                 int                             `json:"qos"`                    // "The maximum QoS level of the state topic."
	Retain              bool                            `json:"retain"`                 // "If the published message should have the retain flag on or not."
	Schema              string                          `json:"schema"`                 // "The schema to use. Must be `state` to select the state schema."
	SendCommandTopic    string                          `json:"send_command_topic"`     // "The MQTT topic to publish custom commands to the vacuum."
	SendCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	SetFanSpeedTopic    string                          `json:"set_fan_speed_topic"` // "The MQTT topic to publish commands to control the vacuum's fan speed."
	SetFanSpeedFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	StateTopic          string                          `json:"state_topic"` // "The MQTT topic subscribed to receive state messages from the vacuum. Messages received on the `state_topic` must be a valid JSON dictionary, with a mandatory `state` key and optionally `battery_level` and `fan_speed` keys as shown in the [example](#state-mqtt-protocol)."
	StateFunc           func() string                   `json:"-"`
	SupportedFeatures   []string                        `json:"supported_features"` // "List of features that the vacuum supports (possible values are `start`, `stop`, `pause`, `return_home`, `battery`, `status`, `locate`, `clean_spot`, `fan_speed`, `send_command`)."
	UniqueId            string                          `json:"unique_id"`          // "An ID that uniquely identifies this vacuum. If two vacuums have the same unique ID, Home Assistant will raise an exception."
	MQTT                MQTTFields                      `json:"-"`
}

func (d Vacuum) UpdateState() {
	if d.StateTopic != "" {
		state := d.StateFunc()
		if state != stateStore.Vacuum.State[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.StateTopic, common.QoS, common.Retain, state)
			stateStore.Vacuum.State[d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d Vacuum) Subscribe() {
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
	if d.SendCommandTopic != "" {
		t := c.Subscribe(d.SendCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.SetFanSpeedTopic != "" {
		t := c.Subscribe(d.SetFanSpeedTopic, 0, d.MQTT.MessageHandler)
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
func (d Vacuum) UnSubscribe() {
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
	if d.SendCommandTopic != "" {
		t := c.Unsubscribe(d.SendCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.SetFanSpeedTopic != "" {
		t := c.Unsubscribe(d.SetFanSpeedTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d Vacuum) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d Vacuum) Initialize() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d Vacuum) PopulateTopics() {
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		store.TopicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.SendCommandFunc != nil {
		d.SendCommandTopic = GetTopic(d, "send_command_topic")
		store.TopicStore[d.SendCommandTopic] = &d.SendCommandFunc
	}
	if d.SetFanSpeedFunc != nil {
		d.SetFanSpeedTopic = GetTopic(d, "set_fan_speed_topic")
		store.TopicStore[d.SetFanSpeedTopic] = &d.SetFanSpeedFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d Vacuum) SetMQTTFields(fields MQTTFields) {
	d.MQTT = fields
}
func (d Vacuum) GetMQTTFields() (fields MQTTFields) {
	return d.MQTT
}
