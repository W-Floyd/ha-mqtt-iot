package InternalDevice

import ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/external"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
func (iDevice Camera) Translate() ExternalDevice.Camera {
	eDevice := ExternalDevice.Camera{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
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
	EnabledByDefault bool   `json:"enabled_by_default"`
	EntityCategory   string `json:"entity_category"`
	Icon             string `json:"icon"`
	Name             string `json:"name"`
	ObjectId         string `json:"object_id"`
	Topic            string `json:"topic"`
	UniqueId         string `json:"unique_id"`
	MQTT             struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
