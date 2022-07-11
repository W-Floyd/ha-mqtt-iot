package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
func (iDevice Number) Translate() externaldevice.Number {
	eDevice := externaldevice.Number{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.CommandTemplate = iDevice.CommandTemplate
	eDevice.CommandFunc = common.ConstructCommandFunc(iDevice.Command)
	eDevice.DeviceClass = iDevice.DeviceClass
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.Icon = iDevice.Icon
	eDevice.Max = iDevice.Max
	eDevice.Min = iDevice.Min
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.Optimistic = iDevice.Optimistic
	eDevice.PayloadReset = iDevice.PayloadReset
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.Step = iDevice.Step
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.UnitOfMeasurement = iDevice.UnitOfMeasurement
	eDevice.ValueTemplate = iDevice.ValueTemplate
	eDevice.Initialize()
	return eDevice
}

type Number struct {
	CommandTemplate   string   `json:"command_template"` // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to generate the payload to send to `command_topic`."
	Command           []string `json:"command"`
	DeviceClass       string   `json:"device_class"`       // "The [type/class](/integrations/number/#device-class) of the number."
	EnabledByDefault  bool     `json:"enabled_by_default"` // "Flag which defines if the entity should be enabled when first added."
	Encoding          string   `json:"encoding"`           // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory    string   `json:"entity_category"`    // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon              string   `json:"icon"`               // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Max               float64  `json:"max"`                // "Maximum value."
	Min               float64  `json:"min"`                // "Minimum value."
	Name              string   `json:"name"`               // "The name of the Number."
	ObjectId          string   `json:"object_id"`          // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic        bool     `json:"optimistic"`         // "Flag that defines if number works in optimistic mode."
	PayloadReset      string   `json:"payload_reset"`      // "A special payload that resets the state to `None` when received on the `state_topic`."
	Qos               int      `json:"qos"`                // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	Retain            bool     `json:"retain"`             // "If the published message should have the retain flag on or not."
	State             []string `json:"state"`
	Step              float64  `json:"step"`                // "Step value. Smallest value `0.001`."
	UniqueId          string   `json:"unique_id"`           // "An ID that uniquely identifies this Number. If two Numbers have the same unique ID Home Assistant will raise an exception."
	UnitOfMeasurement string   `json:"unit_of_measurement"` // "Defines the unit of measurement of the sensor, if any."
	ValueTemplate     string   `json:"value_template"`      // "Defines a [template](/docs/configuration/templating/#processing-incoming-data) to extract the value."
	MQTT              struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
