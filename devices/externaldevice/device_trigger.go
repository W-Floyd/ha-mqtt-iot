package ExternalDevice

import (
	"encoding/json"
	common "github.com/W-Floyd/ha-mqtt-iot/common"
	"log"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (d DeviceTrigger) GetRawId() string {
	return "device_trigger"
}
func (d DeviceTrigger) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d DeviceTrigger) GetUniqueId() string {
	return ""
}
func (d DeviceTrigger) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type DeviceTrigger struct {
	AutomationType string `json:"automation_type"` // "The type of automation, must be 'trigger'."
	Device         struct {
		ConfigurationUrl string `json:"configuration_url"` // "A link to the webpage that can manage the configuration of this device. Can be either an HTTP or HTTPS link."
		Connections      string `json:"connections"`       // "A list of connections of the device to the outside world as a list of tuples `[connection_type, connection_identifier]`. For example the MAC address of a network interface: `'connections': ['mac', '02:5b:26:a8:dc:12']`."
		Identifiers      string `json:"identifiers"`       // "A list of IDs that uniquely identify the device. For example a serial number."
		Manufacturer     string `json:"manufacturer"`      // "The manufacturer of the device."
		Model            string `json:"model"`             // "The model of the device."
		Name             string `json:"name"`              // "The name of the device."
		SuggestedArea    string `json:"suggested_area"`    // "Suggest an area if the device isn’t in one yet."
		SwVersion        string `json:"sw_version"`        // "The firmware version of the device."
		Viadevice        string `json:"viadevice"`         // null
	} `json:"device"`
	Payload       string     `json:"payload"`        // "Optional payload to match the payload being sent over the topic."
	Qos           int        `json:"qos"`            // "The maximum QoS level to be used when receiving messages."
	Subtype       string     `json:"subtype"`        // "The subtype of the trigger, e.g. `button_1`. Entries supported by the frontend: `turn_on`, `turn_off`, `button_1`, `button_2`, `button_3`, `button_4`, `button_5`, `button_6`. If set to an unsupported value, will render as `subtype type`, e.g. `left_button pressed` with `type` set to `button_short_press` and `subtype` set to `left_button`"
	Topic         string     `json:"topic"`          // "The MQTT topic subscribed to receive trigger events."
	Type          string     `json:"type"`           // "The type of the trigger, e.g. `button_short_press`. Entries supported by the frontend: `button_short_press`, `button_short_release`, `button_long_press`, `button_long_release`, `button_double_press`, `button_triple_press`, `button_quadruple_press`, `button_quintuple_press`. If set to an unsupported value, will render as `subtype type`, e.g. `button_1 spammed` with `type` set to `spammed` and `subtype` set to `button_1`"
	ValueTemplate string     `json:"value_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the value."
	MQTT          MQTTFields `json:"-"`
}

func (d DeviceTrigger) UpdateState() {}
func (d DeviceTrigger) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	token := c.Publish(GetDiscoveryTopic(d), 0, true, message)
	token.Wait()
	time.Sleep(common.HADiscoveryDelay)
	d.AnnounceAvailable()
	d.UpdateState()
}
func (d DeviceTrigger) UnSubscribe()       {}
func (d DeviceTrigger) AnnounceAvailable() {}
func (d DeviceTrigger) Initialize() {
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d DeviceTrigger) PopulateTopics() {}
