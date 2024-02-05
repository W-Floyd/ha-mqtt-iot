package ExternalDevice

import (
	"encoding/json"
	common "github.com/W-Floyd/ha-mqtt-iot/common"
	store "github.com/W-Floyd/ha-mqtt-iot/store"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	strcase "github.com/iancoleman/strcase"
	"log"
	"time"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Update struct {
	AvailabilityMode     *string                         `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate *string                         `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    *string                         `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	AvailabilityFunc     func() string                   `json:"-"`                               // Function for availability
	CommandTopic         *string                         `json:"command_topic,omitempty"`         // "The MQTT topic to publish `payload_install` to start installing process."
	CommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`                               // Function for command
	Device               struct {
		ConfigurationUrl *string `json:"configuration_url,omitempty"` // "A link to the webpage that can manage the configuration of this device. Can be either an `http://`, `https://` or an internal `homeassistant://` URL."
		Connections      *string `json:"connections,omitempty"`       // "A list of connections of the device to the outside world as a list of tuples `[connection_type, connection_identifier]`. For example the MAC address of a network interface: `\"connections\": [[\"mac\", \"02:5b:26:a8:dc:12\"]]`."
		Identifiers      *string `json:"identifiers,omitempty"`       // "A list of IDs that uniquely identify the device. For example a serial number."
		Manufacturer     *string `json:"manufacturer,omitempty"`      // "The manufacturer of the device."
		Model            *string `json:"model,omitempty"`             // "The model of the device."
		Name             *string `json:"name,omitempty"`              // "The name of the device."
		SuggestedArea    *string `json:"suggested_area,omitempty"`    // "Suggest an area if the device isn’t in one yet."
		SwVersion        *string `json:"sw_version,omitempty"`        // "The firmware version of the device."
		ViaDevice        *string `json:"via_device,omitempty"`        // "Identifier of a device that routes messages between this device and Home Assistant. Examples of such devices are hubs, or parent devices of a sub-device. This is used to show device topology in Home Assistant."
	} `json:"device,omitempty"` // Device configuration parameters
	DeviceClass            *string                         `json:"device_class,omitempty"`             // "The [type/class](/integrations/update/#device-classes) of the update to set the icon in the frontend. The `device_class` can be `null`."
	EnabledByDefault       *bool                           `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string                         `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string                         `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	EntityPicture          *string                         `json:"entity_picture,omitempty"`           // "Picture URL for the entity."
	Icon                   *string                         `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string                         `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributesTopic    *string                         `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as entity attributes. Implies `force_update` of the current select state when a message is received on this topic."
	JsonAttributesFunc     func(mqtt.Message, mqtt.Client) `json:"-"`                                  // Function for json attributes
	LatestVersionTemplate  *string                         `json:"latest_version_template,omitempty"`  // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the latest version value."
	LatestVersionTopic     *string                         `json:"latest_version_topic,omitempty"`     // "The MQTT topic subscribed to receive an update of the latest version."
	LatestVersionFunc      func(mqtt.Message, mqtt.Client) `json:"-"`                                  // Function for latest version
	Name                   *string                         `json:"name,omitempty"`                     // "The name of the Update. Can be set to `null` if only the device name is relevant."
	ObjectId               *string                         `json:"object_id,omitempty"`                // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadInstall         *string                         `json:"payload_install,omitempty"`          // "The MQTT payload to start installing process."
	Qos                    *int                            `json:"qos,omitempty"`                      // "The maximum QoS level to be used when receiving and publishing messages."
	ReleaseSummary         *string                         `json:"release_summary,omitempty"`          // "Summary of the release notes or changelog. This is suitable a brief update description of max 255 characters."
	ReleaseUrl             *string                         `json:"release_url,omitempty"`              // "URL to the full release notes of the latest version available."
	Retain                 *bool                           `json:"retain,omitempty"`                   // "If the published message should have the retain flag on or not."
	StateTopic             *string                         `json:"state_topic,omitempty"`              // "The MQTT topic subscribed to receive state updates. The state update may be either JSON or a simple string with `installed_version` value. When a JSON payload is detected, the state value of the JSON payload should supply the `installed_version` and can optional supply: `latest_version`, `title`, `release_summary`, `release_url` or an `entity_picture` URL."
	StateFunc              func() string                   `json:"-"`                                  // Function for state
	Title                  *string                         `json:"title,omitempty"`                    // "Title of the software, or firmware update. This helps to differentiate between the device or entity name versus the title of the software installed."
	UniqueId               *string                         `json:"unique_id,omitempty"`                // "An ID that uniquely identifies this Update. If two Updates have the same unique ID Home Assistant will raise an exception."
	ValueTemplate          *string                         `json:"value_template,omitempty"`           // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the `installed_version` state value or to render to a valid JSON payload on from the payload received on `state_topic`."
	MQTT                   *MQTTFields                     `json:"-"`                                  // MQTT configuration parameters
}

func (d *Update) GetRawId() string {
	return "update"
}
func (d *Update) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Update) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Update) PopulateDevice() {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &common.MachineID
}
func (d *Update) UpdateState() {
	if d.AvailabilityTopic != nil {
		state := d.AvailabilityFunc()
		if state != stateStore.Update.Availability[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Update.Availability[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.StateTopic != nil {
		state := d.StateFunc()
		if state != stateStore.Update.State[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Update.State[d.GetUniqueId()] = state
			token.Wait()
		}
	}
}
func (d *Update) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	if d.CommandTopic != nil {
		t := c.Subscribe(*d.CommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.JsonAttributesTopic != nil {
		t := c.Subscribe(*d.JsonAttributesTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.LatestVersionTopic != nil {
		t := c.Subscribe(*d.LatestVersionTopic, 0, d.MQTT.MessageHandler)
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
func (d *Update) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "offline")
	token.Wait()
	if d.CommandTopic != nil {
		t := c.Unsubscribe(*d.CommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.JsonAttributesTopic != nil {
		t := c.Unsubscribe(*d.JsonAttributesTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.LatestVersionTopic != nil {
		t := c.Unsubscribe(*d.LatestVersionTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d *Update) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d *Update) Initialize() {
	if d.Qos == nil {
		d.Qos = new(int)
		*d.Qos = int(common.QoS)
	}
	if d.Retain == nil {
		d.Retain = new(bool)
		*d.Retain = common.RetainClient
	}
	if d.UniqueId == nil {
		d.UniqueId = new(string)
		*d.UniqueId = strcase.ToDelimited(*d.Name, uint8(0x2d))
	}
	d.PopulateDevice()
	d.AddMessageHandler()
	d.PopulateTopics()
}
func (d *Update) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = new(string)
		*d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = new(string)
		*d.CommandTopic = GetTopic(d, "command_topic")
		store.TopicStore[*d.CommandTopic] = &d.CommandFunc
	}
	if d.JsonAttributesFunc != nil {
		d.JsonAttributesTopic = new(string)
		*d.JsonAttributesTopic = GetTopic(d, "json_attributes_topic")
		store.TopicStore[*d.JsonAttributesTopic] = &d.JsonAttributesFunc
	}
	if d.LatestVersionFunc != nil {
		d.LatestVersionTopic = new(string)
		*d.LatestVersionTopic = GetTopic(d, "latest_version_topic")
		store.TopicStore[*d.LatestVersionTopic] = &d.LatestVersionFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = new(string)
		*d.StateTopic = GetTopic(d, "state_topic")
	}
}
func (d *Update) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Update) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
