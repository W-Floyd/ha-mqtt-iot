package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
func (iDevice Camera) Translate() externaldevice.Camera {
	eDevice := externaldevice.Camera{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.AvailabilityMode = iDevice.AvailabilityMode
	eDevice.AvailabilityTemplate = iDevice.AvailabilityTemplate
	eDevice.AvailabilityFunc = common.ConstructStateFunc(iDevice.Availability)
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.Icon = iDevice.Icon
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.Topic = iDevice.Topic
	eDevice.UniqueId = iDevice.UniqueId
	if len(iDevice.Availability) == 0 {
		eDevice.AvailabilityFunc = common.AvailabilityFunc
	}
	eDevice.Initialize()
	return eDevice
}

type Camera struct {
	AvailabilityMode     string   `json:"availability_mode"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate string   `json:"availability_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability         []string `json:"availability"`
	EnabledByDefault     bool     `json:"enabled_by_default"` // "Flag which defines if the entity should be enabled when first added."
	Encoding             string   `json:"encoding"`           // "The encoding of the image payloads received. Set to `\"b64\"` to enable base64 decoding of image payload. If not set, or if set to `null`, the image payload must be raw binary data."
	EntityCategory       string   `json:"entity_category"`    // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                 string   `json:"icon"`               // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Name                 string   `json:"name"`               // "The name of the camera."
	ObjectId             string   `json:"object_id"`          // "Used instead of `name` for automatic generation of `entity_id`"
	Topic                string   `json:"topic"`              // "The MQTT topic to subscribe to."
	UniqueId             string   `json:"unique_id"`          // "An ID that uniquely identifies this camera. If two cameras have the same unique ID Home Assistant will raise an exception."
	MQTT                 struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
