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
type Image struct {
	AvailabilityMode     *string       `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate *string       `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic    *string       `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	AvailabilityFunc     func() string `json:"-"`                               // Function for availability
	ContentType          *string       `json:"content_type,omitempty"`          // "The content type of and image data message received on `image_topic`. This option cannot be used with the `url_topic` because the content type is derived when downloading the image."
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
	EnabledByDefault       *bool                           `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string                         `json:"encoding,omitempty"`                 // "The encoding of the payloads received. Set to `\"\"` to disable decoding of incoming payload. Use `image_encoding` to enable `Base64` decoding on `image_topic`."
	EntityCategory         *string                         `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                   *string                         `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	ImageEncoding          *string                         `json:"image_encoding,omitempty"`           // "The encoding of the image payloads received. Set to `\"b64\"` to enable base64 decoding of image payload. If not set, the image payload must be raw binary data."
	ImageTopic             *string                         `json:"image_topic,omitempty"`              // "The MQTT topic to subscribe to receive the image payload of the image to be downloaded. Ensure the `content_type` type option is set to the corresponding content type. This option cannot be used together with the `url_topic` option. But at least one of these option is required."
	ImageFunc              func(mqtt.Message, mqtt.Client) `json:"-"`                                  // Function for image
	JsonAttributesTemplate *string                         `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributesTopic    *string                         `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Implies `force_update` of the current sensor state when a message is received on this topic."
	JsonAttributesFunc     func(mqtt.Message, mqtt.Client) `json:"-"`                                  // Function for json attributes
	Name                   *string                         `json:"name,omitempty"`                     // "The name of the image. Can be set to `null` if only the device name is relevant."
	ObjectId               *string                         `json:"object_id,omitempty"`                // "Used instead of `name` for automatic generation of `entity_id`"
	UniqueId               *string                         `json:"unique_id,omitempty"`                // "An ID that uniquely identifies this image. If two images have the same unique ID Home Assistant will raise an exception."
	UrlTemplate            *string                         `json:"url_template,omitempty"`             // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the image URL from a message received at `url_topic`."
	UrlTopic               *string                         `json:"url_topic,omitempty"`                // "The MQTT topic to subscribe to receive an image URL. A `url_template` option can extract the URL from the message. The `content_type` will be derived from the image when downloaded. This option cannot be used together with the `image_topic` option, but at least one of these options is required."
	UrlFunc                func(mqtt.Message, mqtt.Client) `json:"-"`                                  // Function for url
	MQTT                   *MQTTFields                     `json:"-"`                                  // MQTT configuration parameters
}

func (d *Image) GetRawId() string {
	return "image"
}
func (d *Image) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Image) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Image) PopulateDevice() {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &common.MachineID
}
func (d *Image) UpdateState() {
	if d.AvailabilityTopic != nil {
		state := d.AvailabilityFunc()
		if state != stateStore.Image.Availability[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, common.QoS, common.Retain, state)
			stateStore.Image.Availability[d.GetUniqueId()] = state
			token.Wait()
		}
	}
}
func (d *Image) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	if d.ImageTopic != nil {
		t := c.Subscribe(*d.ImageTopic, 0, d.MQTT.MessageHandler)
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
	if d.UrlTopic != nil {
		t := c.Subscribe(*d.UrlTopic, 0, d.MQTT.MessageHandler)
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
func (d *Image) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "offline")
	token.Wait()
	if d.ImageTopic != nil {
		t := c.Unsubscribe(*d.ImageTopic)
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
	if d.UrlTopic != nil {
		t := c.Unsubscribe(*d.UrlTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d *Image) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d *Image) Initialize() {
	if d.UniqueId == nil {
		d.UniqueId = new(string)
		*d.UniqueId = strcase.ToDelimited(*d.Name, uint8(0x2d))
	}
	d.PopulateDevice()
	d.AddMessageHandler()
	d.PopulateTopics()
}
func (d *Image) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = new(string)
		*d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.ImageFunc != nil {
		d.ImageTopic = new(string)
		*d.ImageTopic = GetTopic(d, "image_topic")
		store.TopicStore[*d.ImageTopic] = &d.ImageFunc
	}
	if d.JsonAttributesFunc != nil {
		d.JsonAttributesTopic = new(string)
		*d.JsonAttributesTopic = GetTopic(d, "json_attributes_topic")
		store.TopicStore[*d.JsonAttributesTopic] = &d.JsonAttributesFunc
	}
	if d.UrlFunc != nil {
		d.UrlTopic = new(string)
		*d.UrlTopic = GetTopic(d, "url_topic")
		store.TopicStore[*d.UrlTopic] = &d.UrlFunc
	}
}
func (d *Image) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Image) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
