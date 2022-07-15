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
func (d *Tag) GetRawId() string {
	return "tag"
}
func (d *Tag) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Tag) GetUniqueId() string {
	return ""
}
func (d *Tag) PopulateDevice() {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &common.MachineID
}

type Tag struct {
	Device struct {
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
	Topic         *string     `json:"topic,omitempty"`          // "The MQTT topic subscribed to receive tag scanned events."
	ValueTemplate *string     `json:"value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) that returns a tag ID."
	MQTT          *MQTTFields `json:"-"`
}

func (d *Tag) UpdateState() {}
func (d *Tag) Subscribe() {
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
func (d *Tag) UnSubscribe()       {}
func (d *Tag) AnnounceAvailable() {}
func (d *Tag) Initialize() {
	d.PopulateDevice()
	d.AddMessageHandler()
	d.PopulateTopics()
}
func (d *Tag) PopulateTopics() {}
func (d *Tag) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Tag) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
