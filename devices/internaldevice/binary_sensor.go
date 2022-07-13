package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
func (iDevice BinarySensor) Translate() externaldevice.BinarySensor {
	eDevice := externaldevice.BinarySensor{}
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
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.OffDelay = iDevice.OffDelay
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.PayloadOff = iDevice.PayloadOff
	eDevice.PayloadOn = iDevice.PayloadOn
	eDevice.Qos = iDevice.Qos
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.ValueTemplate = iDevice.ValueTemplate
	if len(iDevice.Availability) == 0 {
		eDevice.AvailabilityFunc = common.AvailabilityFunc
	}
	eDevice.Initialize()
	return eDevice
}

type BinarySensor struct {
	AvailabilityMode     string   `json:"availability_mode"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate string   `json:"availability_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability         []string `json:"availability"`
	DeviceClass          string   `json:"device_class"`          // "Sets the [class of the device](/integrations/binary_sensor/#device-class), changing the device state and icon that is displayed on the frontend."
	EnabledByDefault     bool     `json:"enabled_by_default"`    // "Flag which defines if the entity should be enabled when first added."
	Encoding             string   `json:"encoding"`              // "The encoding of the payloads received. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory       string   `json:"entity_category"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	ExpireAfter          int      `json:"expire_after"`          // "Defines the number of seconds after the sensor's state expires, if it's not updated. After expiry, the sensor's state becomes `unavailable`."
	ForceUpdate          bool     `json:"force_update"`          // "Sends update events (which results in update of [state object](/docs/configuration/state_object/)'s `last_changed`) even if the sensor's state hasn't changed. Useful if you want to have meaningful value graphs in history or want to create an automation that triggers on *every* incoming state message (not only when the sensor's new state is different to the current one)."
	Icon                 string   `json:"icon"`                  // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Name                 string   `json:"name"`                  // "The name of the binary sensor."
	ObjectId             string   `json:"object_id"`             // "Used instead of `name` for automatic generation of `entity_id`"
	OffDelay             int      `json:"off_delay"`             // "For sensors that only send `on` state updates (like PIRs), this variable sets a delay in seconds after which the sensor's state will be updated back to `off`."
	PayloadAvailable     string   `json:"payload_available"`     // "The string that represents the `online` state."
	PayloadNotAvailable  string   `json:"payload_not_available"` // "The string that represents the `offline` state."
	PayloadOff           string   `json:"payload_off"`           // "The string that represents the `off` state. It will be compared to the message in the `state_topic` (see `value_template` for details)"
	PayloadOn            string   `json:"payload_on"`            // "The string that represents the `on` state. It will be compared to the message in the `state_topic` (see `value_template` for details)"
	Qos                  int      `json:"qos"`                   // "The maximum QoS level to be used when receiving messages."
	State                []string `json:"state"`
	UniqueId             string   `json:"unique_id"`      // "An ID that uniquely identifies this sensor. If two sensors have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate        string   `json:"value_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) that returns a string to be compared to `payload_on`/`payload_off` or an empty string, in which case the MQTT message will be removed. Available variables: `entity_id`. Remove this option when 'payload_on' and 'payload_off' are sufficient to match your payloads (i.e no pre-processing of original message is required)."
	MQTT                 struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
