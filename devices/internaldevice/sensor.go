package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
func (iDevice Sensor) Translate() externaldevice.Sensor {
	eDevice := externaldevice.Sensor{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.AvailabilityMode = iDevice.AvailabilityMode
	eDevice.AvailabilityTemplate = iDevice.AvailabilityTemplate
	eDevice.AvailabilityFunc = common.ConstructStateFunc(iDevice.Availability)
	eDevice.DeviceClass = iDevice.DeviceClass
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.ExpireAfter = iDevice.ExpireAfter
	eDevice.ForceUpdate = iDevice.ForceUpdate
	eDevice.Icon = iDevice.Icon
	eDevice.LastResetValueTemplate = iDevice.LastResetValueTemplate
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.Qos = iDevice.Qos
	eDevice.StateClass = iDevice.StateClass
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.UnitOfMeasurement = iDevice.UnitOfMeasurement
	eDevice.ValueTemplate = iDevice.ValueTemplate
	if len(iDevice.Availability) == 0 {
		eDevice.AvailabilityFunc = common.AvailabilityFunc
	}
	eDevice.Initialize()
	return eDevice
}

type Sensor struct {
	AvailabilityMode       string   `json:"availability_mode"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   string   `json:"availability_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability           []string `json:"availability"`
	DeviceClass            string   `json:"device_class"`              // "The [type/class](/integrations/sensor/#device-class) of the sensor to set the icon in the frontend."
	EnabledByDefault       bool     `json:"enabled_by_default"`        // "Flag which defines if the entity should be enabled when first added."
	Encoding               string   `json:"encoding"`                  // "The encoding of the payloads received. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         string   `json:"entity_category"`           // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	ExpireAfter            int      `json:"expire_after"`              // "Defines the number of seconds after the sensor's state expires, if it's not updated. After expiry, the sensor's state becomes `unavailable`."
	ForceUpdate            bool     `json:"force_update"`              // "Sends update events even if the value hasn't changed. Useful if you want to have meaningful value graphs in history."
	Icon                   string   `json:"icon"`                      // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	LastResetValueTemplate string   `json:"last_reset_value_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the last_reset. Available variables: `entity_id`. The `entity_id` can be used to reference the entity's attributes."
	Name                   string   `json:"name"`                      // "The name of the MQTT sensor."
	ObjectId               string   `json:"object_id"`                 // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadAvailable       string   `json:"payload_available"`         // "The payload that represents the available state."
	PayloadNotAvailable    string   `json:"payload_not_available"`     // "The payload that represents the unavailable state."
	Qos                    int      `json:"qos"`                       // "The maximum QoS level of the state topic."
	StateClass             string   `json:"state_class"`               // "The [state_class](https://developers.home-assistant.io/docs/core/entity/sensor#available-state-classes) of the sensor."
	State                  []string `json:"state"`
	UniqueId               string   `json:"unique_id"`           // "An ID that uniquely identifies this sensor. If two sensors have the same unique ID, Home Assistant will raise an exception."
	UnitOfMeasurement      string   `json:"unit_of_measurement"` // "Defines the units of measurement of the sensor, if any."
	ValueTemplate          string   `json:"value_template"`      // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the value. Available variables: `entity_id`. The `entity_id` can be used to reference the entity's attributes. If the template throws an error, the current state will be used instead."
	MQTT                   struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
