package ExternalDevice

import (
	"encoding/json"
	common "github.com/W-Floyd/ha-mqtt-iot/common"
	"log"
	"time"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type DeviceTrigger struct {
	AutomationType *string `json:"automation_type,omitempty"` // "The type of automation, must be 'trigger'."
	Device         struct {
		ConfigurationUrl *string `json:"configuration_url,omitempty"` // "A link to the webpage that can manage the configuration of this device. Can be either an HTTP or HTTPS link."
		Connections      *string `json:"connections,omitempty"`       // "A list of connections of the device to the outside world as a list of tuples `[connection_type, connection_identifier]`. For example the MAC address of a network interface: `'connections': ['mac', '02:5b:26:a8:dc:12']`."
		Identifiers      *string `json:"identifiers,omitempty"`       // "A list of IDs that uniquely identify the device. For example a serial number."
		Manufacturer     *string `json:"manufacturer,omitempty"`      // "The manufacturer of the device."
		Model            *string `json:"model,omitempty"`             // "The model of the device."
		Name             *string `json:"name,omitempty"`              // "The name of the device."
		SuggestedArea    *string `json:"suggested_area,omitempty"`    // "Suggest an area if the device isn’t in one yet."
		SwVersion        *string `json:"sw_version,omitempty"`        // "The firmware version of the device."
		Viadevice        *string `json:"viadevice,omitempty"`         // null
	} `json:"device,omitempty"`
	Payload       *string       `json:"payload,omitempty"` // "Optional payload to match the payload being sent over the topic."
	Qos           *int          `json:"qos,omitempty"`     // "The maximum QoS level to be used when receiving messages."
	Subtype       *string       `json:"subtype,omitempty"` // "The subtype of the trigger, e.g. `button_1`. Entries supported by the frontend: `turn_on`, `turn_off`, `button_1`, `button_2`, `button_3`, `button_4`, `button_5`, `button_6`. If set to an unsupported value, will render as `subtype type`, e.g. `left_button pressed` with `type` set to `button_short_press` and `subtype` set to `left_button`"
	StateTopic    *string       `json:"topic,omitempty"`   // "The MQTT topic subscribed to receive trigger events."
	StateFunc     func() string `json:"-"`
	Type          *string       `json:"type,omitempty"`           // "The type of the trigger, e.g. `button_short_press`. Entries supported by the frontend: `button_short_press`, `button_short_release`, `button_long_press`, `button_long_release`, `button_double_press`, `button_triple_press`, `button_quadruple_press`, `button_quintuple_press`. If set to an unsupported value, will render as `subtype type`, e.g. `button_1 spammed` with `type` set to `spammed` and `subtype` set to `button_1`"
	ValueTemplate *string       `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the value."
	MQTT          *MQTTFields   `json:"-"`
}

func (d *DeviceTrigger) GetRawId() string {
	return "device_trigger"
}
func (d *DeviceTrigger) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d DeviceTrigger) GetUniqueId() string {
	return ""
}
func (d *DeviceTrigger) PopulateDevice() {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &common.MachineID
}
func (d *DeviceTrigger) UpdateState() {
	if d.StateTopic != nil {
		state := d.StateFunc()
		if state != stateStore.DeviceTrigger.State[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, byte(*d.Qos), common.Retain, state)
			stateStore.DeviceTrigger.State[d.GetUniqueId()] = state
			token.Wait()
		}
	}
}
func (d *DeviceTrigger) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	token := c.Publish(GetDiscoveryTopic(d), 0, true, message)
	token.Wait()
	time.Sleep(common.HADiscoveryDelay)
	d.UpdateState()
}
func (d *DeviceTrigger) UnSubscribe()       {}
func (d *DeviceTrigger) AnnounceAvailable() {}
func (d *DeviceTrigger) Initialize() {
	if d.Qos == nil {
		d.Qos = new(int)
		*d.Qos = int(common.QoS)
	}
	d.PopulateDevice()
	d.AddMessageHandler()
	d.PopulateTopics()
}
func (d *DeviceTrigger) PopulateTopics() {
	if d.StateFunc != nil {
		d.StateTopic = new(string)
		*d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d *DeviceTrigger) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *DeviceTrigger) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
