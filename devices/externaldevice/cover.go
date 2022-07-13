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
func (d *Cover) GetRawId() string {
	return "cover"
}
func (d *Cover) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Cover) GetUniqueId() string {
	return d.UniqueId
}
func (d *Cover) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Cover struct {
	AvailabilityMode     string                          `json:"availability_mode"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate string                          `json:"availability_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    string                          `json:"availability_topic"`    // "The MQTT topic subscribed to to receive birth and LWT messages from the MQTT cover device. If an `availability` topic is not defined, the cover availability state will always be `available`. If an `availability` topic is defined, the cover availability state will be `unavailable` by default. Must not be used together with `availability`."
	CommandTopic         string                          `json:"command_topic"`         // "The MQTT topic to publish commands to control the cover."
	CommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
	Device               struct {
		ConfigurationUrl string `json:"configuration_url"` // "A link to the webpage that can manage the configuration of this device. Can be either an HTTP or HTTPS link."
		Connections      string `json:"connections"`       // "A list of connections of the device to the outside world as a list of tuples `[connection_type, connection_identifier]`. For example the MAC address of a network interface: `\"connections\": [\"mac\", \"02:5b:26:a8:dc:12\"]`."
		Identifiers      string `json:"identifiers"`       // "A list of IDs that uniquely identify the device. For example a serial number."
		Manufacturer     string `json:"manufacturer"`      // "The manufacturer of the device."
		Model            string `json:"model"`             // "The model of the device."
		Name             string `json:"name"`              // "The name of the device."
		SuggestedArea    string `json:"suggested_area"`    // "Suggest an area if the device isn’t in one yet."
		SwVersion        string `json:"sw_version"`        // "The firmware version of the device."
		Viadevice        string `json:"viadevice"`         // null
	} `json:"device"`
	DeviceClass         string                          `json:"device_class"`          // "Sets the [class of the device](/integrations/cover/), changing the device state and icon that is displayed on the frontend."
	EnabledByDefault    bool                            `json:"enabled_by_default"`    // "Flag which defines if the entity should be enabled when first added."
	Encoding            string                          `json:"encoding"`              // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory      string                          `json:"entity_category"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                string                          `json:"icon"`                  // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Name                string                          `json:"name"`                  // "The name of the cover."
	ObjectId            string                          `json:"object_id"`             // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic          bool                            `json:"optimistic"`            // "Flag that defines if switch works in optimistic mode."
	PayloadAvailable    string                          `json:"payload_available"`     // "The payload that represents the online state."
	PayloadClose        string                          `json:"payload_close"`         // "The command payload that closes the cover."
	PayloadNotAvailable string                          `json:"payload_not_available"` // "The payload that represents the offline state."
	PayloadOpen         string                          `json:"payload_open"`          // "The command payload that opens the cover."
	PayloadStop         string                          `json:"payload_stop"`          // "The command payload that stops the cover."
	PositionClosed      int                             `json:"position_closed"`       // "Number which represents closed position."
	PositionOpen        int                             `json:"position_open"`         // "Number which represents open position."
	PositionTemplate    string                          `json:"position_template"`     // "Defines a [template](/topics/templating/) that can be used to extract the payload for the `position_topic` topic. Within the template the following variables are available: `entity_id`, `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function;"
	PositionTopic       string                          `json:"position_topic"`        // "The MQTT topic subscribed to receive cover position messages."
	PositionFunc        func() string                   `json:"-"`
	Qos                 int                             `json:"qos"`                   // "The maximum QoS level to be used when receiving and publishing messages."
	Retain              bool                            `json:"retain"`                // "Defines if published messages should have the retain flag set."
	SetPositionTemplate string                          `json:"set_position_template"` // "Defines a [template](/topics/templating/) to define the position to be sent to the `set_position_topic` topic. Incoming position value is available for use in the template `{% raw %}{{ position }}{% endraw %}`. Within the template the following variables are available: `entity_id`, `position`, the target position in percent; `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function;"
	SetPositionTopic    string                          `json:"set_position_topic"`    // "The MQTT topic to publish position commands to. You need to set position_topic as well if you want to use position topic. Use template if position topic wants different values than within range `position_closed` - `position_open`. If template is not defined and `position_closed != 100` and `position_open != 0` then proper position value is calculated from percentage position."
	SetPositionFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	StateClosed         string                          `json:"state_closed"`  // "The payload that represents the closed state."
	StateClosing        string                          `json:"state_closing"` // "The payload that represents the closing state."
	StateOpen           string                          `json:"state_open"`    // "The payload that represents the open state."
	StateOpening        string                          `json:"state_opening"` // "The payload that represents the opening state."
	StateStopped        string                          `json:"state_stopped"` // "The payload that represents the stopped state (for covers that do not report `open`/`closed` state)."
	StateTopic          string                          `json:"state_topic"`   // "The MQTT topic subscribed to receive cover state messages. State topic can only read (`open`, `opening`, `closed`, `closing` or `stopped`) state."
	StateFunc           func() string                   `json:"-"`
	TiltClosedValue     int                             `json:"tilt_closed_value"`     // "The value that will be sent on a `close_cover_tilt` command."
	TiltCommandTemplate string                          `json:"tilt_command_template"` // "Defines a [template](/topics/templating/) that can be used to extract the payload for the `tilt_command_topic` topic. Within the template the following variables are available: `entity_id`, `tilt_position`, the target tilt position in percent; `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function;"
	TiltCommandTopic    string                          `json:"tilt_command_topic"`    // "The MQTT topic to publish commands to control the cover tilt."
	TiltCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	TiltMax             int                             `json:"tilt_max"`             // "The maximum tilt value."
	TiltMin             int                             `json:"tilt_min"`             // "The minimum tilt value."
	TiltOpenedValue     int                             `json:"tilt_opened_value"`    // "The value that will be sent on an `open_cover_tilt` command."
	TiltOptimistic      bool                            `json:"tilt_optimistic"`      // "Flag that determines if tilt works in optimistic mode."
	TiltStatusTemplate  string                          `json:"tilt_status_template"` // "Defines a [template](/topics/templating/) that can be used to extract the payload for the `tilt_status_topic` topic. Within the template the following variables are available: `entity_id`, `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function;"
	TiltStatusTopic     string                          `json:"tilt_status_topic"`    // "The MQTT topic subscribed to receive tilt status update values."
	TiltStatusFunc      func() string                   `json:"-"`
	UniqueId            string                          `json:"unique_id"`      // "An ID that uniquely identifies this cover. If two covers have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate       string                          `json:"value_template"` // "Defines a [template](/topics/templating/) that can be used to extract the payload for the `state_topic` topic."
	MQTT                MQTTFields                      `json:"-"`
}

func (d *Cover) UpdateState() {
	if d.PositionTopic != "" {
		state := d.PositionFunc()
		if state != stateStore.Cover.Position[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.PositionTopic, common.QoS, common.Retain, state)
			stateStore.Cover.Position[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.StateTopic != "" {
		state := d.StateFunc()
		if state != stateStore.Cover.State[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.StateTopic, common.QoS, common.Retain, state)
			stateStore.Cover.State[d.UniqueId] = state
			token.Wait()
		}
	}
	if d.TiltStatusTopic != "" {
		state := d.TiltStatusFunc()
		if state != stateStore.Cover.TiltStatus[d.UniqueId] || d.MQTT.ForceUpdate {
			c := *d.MQTT.Client
			token := c.Publish(d.TiltStatusTopic, common.QoS, common.Retain, state)
			stateStore.Cover.TiltStatus[d.UniqueId] = state
			token.Wait()
		}
	}
}
func (d *Cover) Subscribe() {
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
	time.Sleep(common.HADiscoveryDelay)
	d.AnnounceAvailable()
	d.UpdateState()
}
func (d *Cover) UnSubscribe() {
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
func (d *Cover) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d *Cover) Initialize() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d *Cover) PopulateTopics() {
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		store.TopicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.PositionFunc != nil {
		d.PositionTopic = GetTopic(d, "position_topic")
	}
	if d.SetPositionFunc != nil {
		d.SetPositionTopic = GetTopic(d, "set_position_topic")
		store.TopicStore[d.SetPositionTopic] = &d.SetPositionFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
	if d.TiltCommandFunc != nil {
		d.TiltCommandTopic = GetTopic(d, "tilt_command_topic")
		store.TopicStore[d.TiltCommandTopic] = &d.TiltCommandFunc
	}
	if d.TiltStatusFunc != nil {
		d.TiltStatusTopic = GetTopic(d, "tilt_status_topic")
	}
}
func (d *Cover) SetMQTTFields(fields MQTTFields) {
	d.MQTT = fields
}
func (d *Cover) GetMQTTFields() (fields MQTTFields) {
	return d.MQTT
}
