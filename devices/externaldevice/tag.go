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
func (d Tag) GetRawId() string {
	return "tag"
}
func (d Tag) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Tag) GetUniqueId() string {
	return ""
}
func (d Tag) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Tag struct {
	Device struct {
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
	Topic         string     `json:"topic"`          // "The MQTT topic subscribed to receive tag scanned events."
	ValueTemplate string     `json:"value_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) that returns a tag ID."
	MQTT          MQTTFields `json:"-"`
}

func (d Tag) UpdateState() {}
func (d Tag) Subscribe() {
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
func (d Tag) UnSubscribe()       {}
func (d Tag) AnnounceAvailable() {}
func (d Tag) Initialize() {
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d Tag) PopulateTopics() {}
