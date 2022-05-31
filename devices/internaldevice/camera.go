package InternalDevice

import externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
func (iDevice Camera) Translate() externaldevice.Camera {
	eDevice := externaldevice.Camera{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.Icon = iDevice.Icon
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.Topic = iDevice.Topic
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.Initialize()
	return eDevice
}

type Camera struct {
	EnabledByDefault bool   `json:"enabled_by_default"` // "Flag which defines if the entity should be enabled when first added."
	Encoding         string `json:"encoding"`           // "The encoding of the image payloads received. Set to `\"b64\"` to enable base64 decoding of image payload. If not set, or if set to `null`, the image payload must be raw binary data."
	EntityCategory   string `json:"entity_category"`    // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon             string `json:"icon"`               // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Name             string `json:"name"`               // "The name of the camera."
	ObjectId         string `json:"object_id"`          // "Used instead of `name` for automatic generation of `entity_id`"
	Topic            string `json:"topic"`              // "The MQTT topic to subscribe to."
	UniqueId         string `json:"unique_id"`          // "An ID that uniquely identifies this camera. If two cameras have the same unique ID Home Assistant will raise an exception."
	MQTT             struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
