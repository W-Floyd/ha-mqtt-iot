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
func (d Camera) GetRawId() string {
	return "camera"
}
func (d Camera) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Camera) GetUniqueId() string {
	return d.UniqueId
}
func (d Camera) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Camera struct {
	AvailabilityMode     string `json:"availability_mode"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate string `json:"availability_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    string `json:"availability_topic"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
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
	EnabledByDefault bool       `json:"enabled_by_default"` // "Flag which defines if the entity should be enabled when first added."
	EntityCategory   string     `json:"entity_category"`    // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon             string     `json:"icon"`               // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Name             string     `json:"name"`               // "The name of the camera."
	ObjectId         string     `json:"object_id"`          // "Used instead of `name` for automatic generation of `entity_id`"
	Topic            string     `json:"topic"`              // "The MQTT topic to subscribe to."
	UniqueId         string     `json:"unique_id"`          // "An ID that uniquely identifies this camera. If two cameras have the same unique ID Home Assistant will raise an exception."
	MQTT             MQTTFields `json:"-"`
}

func (d Camera) UpdateState() {}
func (d Camera) Subscribe() {
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
func (d Camera) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, "offline")
	token.Wait()
}
func (d Camera) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d Camera) Initialize() {
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d Camera) PopulateTopics() {}
