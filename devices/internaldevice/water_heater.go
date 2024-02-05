package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type WaterHeater struct {
	AvailabilityMode           *string     `json:"availability_mode,omitempty"`            // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate       *string     `json:"availability_template,omitempty"`        // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability               *([]string) `json:"availability,omitempty"`                 // Availability for the WaterHeater
	CurrentTemperatureTemplate *string     `json:"current_temperature_template,omitempty"` // "A template with which the value received on `current_temperature_topic` will be rendered."
	CurrentTemperature         *([]string) `json:"current_temperature,omitempty"`          // CurrentTemperature for the WaterHeater
	EnabledByDefault           *bool       `json:"enabled_by_default,omitempty"`           // "Flag which defines if the entity should be enabled when first added."
	Encoding                   *string     `json:"encoding,omitempty"`                     // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory             *string     `json:"entity_category,omitempty"`              // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                       *string     `json:"icon,omitempty"`                         // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Initial                    *int        `json:"initial,omitempty"`                      // "Set the initial target temperature. The default value depends on the temperature unit, and will be 43.3°C or 110°F."
	JsonAttributesTemplate     *string     `json:"json_attributes_template,omitempty"`     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributes             *([]string) `json:"json_attributes,omitempty"`              // JsonAttributes for the WaterHeater
	MaxTemp                    *float64    `json:"max_temp,omitempty"`                     // "Maximum set point available. The default value depends on the temperature unit, and will be 60°C or 140°F."
	MinTemp                    *float64    `json:"min_temp,omitempty"`                     // "Minimum set point available. The default value depends on the temperature unit, and will be 43.3°C or 110°F."
	ModeCommandTemplate        *string     `json:"mode_command_template,omitempty"`        // "A template to render the value sent to the `mode_command_topic` with."
	ModeCommand                *([]string) `json:"mode_command,omitempty"`                 // ModeCommand for the WaterHeater
	ModeStateTemplate          *string     `json:"mode_state_template,omitempty"`          // "A template to render the value received on the `mode_state_topic` with."
	ModeState                  *([]string) `json:"mode_state,omitempty"`                   // ModeState for the WaterHeater
	Modes                      *([]string) `json:"modes,omitempty"`                        // "A list of supported modes. Needs to be a subset of the default values."
	Name                       *string     `json:"name,omitempty"`                         // "The name of the water heater. Can be set to `null` if only the device name is relevant."
	ObjectId                   *string     `json:"object_id,omitempty"`                    // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic                 *bool       `json:"optimistic,omitempty"`                   // "Flag that defines if the water heater works in optimistic mode"
	PayloadAvailable           *string     `json:"payload_available,omitempty"`            // "The payload that represents the available state."
	PayloadNotAvailable        *string     `json:"payload_not_available,omitempty"`        // "The payload that represents the unavailable state."
	PayloadOff                 *string     `json:"payload_off,omitempty"`                  // "The payload that represents disabled state."
	PayloadOn                  *string     `json:"payload_on,omitempty"`                   // "The payload that represents enabled state."
	PowerCommandTemplate       *string     `json:"power_command_template,omitempty"`       // "A template to render the value sent to the `power_command_topic` with. The `value` parameter is the payload set for `payload_on` or `payload_off`."
	PowerCommand               *([]string) `json:"power_command,omitempty"`                // PowerCommand for the WaterHeater
	Precision                  *float64    `json:"precision,omitempty"`                    // "The desired precision for this device. Can be used to match your actual water heater's precision. Supported values are `0.1`, `0.5` and `1.0`."
	Qos                        *int        `json:"qos,omitempty"`                          // "The maximum QoS level to be used when receiving and publishing messages."
	Retain                     *bool       `json:"retain,omitempty"`                       // "Defines if published messages should have the retain flag set."
	TemperatureCommandTemplate *string     `json:"temperature_command_template,omitempty"` // "A template to render the value sent to the `temperature_command_topic` with."
	TemperatureCommand         *([]string) `json:"temperature_command,omitempty"`          // TemperatureCommand for the WaterHeater
	TemperatureStateTemplate   *string     `json:"temperature_state_template,omitempty"`   // "A template to render the value received on the `temperature_state_topic` with."
	TemperatureState           *([]string) `json:"temperature_state,omitempty"`            // TemperatureState for the WaterHeater
	TemperatureUnit            *string     `json:"temperature_unit,omitempty"`             // "Defines the temperature unit of the device, `C` or `F`. If this is not set, the temperature unit is set to the system temperature unit."
	UniqueId                   *string     `json:"unique_id,omitempty"`                    // "An ID that uniquely identifies this water heater device. If two water heater devices have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate              *string     `json:"value_template,omitempty"`               // "Default template to render the payloads on *all* `*_state_topic`s with."
	MQTT                       struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}

func (iDevice WaterHeater) Translate() externaldevice.WaterHeater {
	eDevice := externaldevice.WaterHeater{}
	eDevice.MQTT = new(externaldevice.MQTTFields)
	if iDevice.MQTT.ForceUpdate != nil {
		eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	}
	if iDevice.MQTT.UpdateInterval != nil {
		eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	}
	if iDevice.AvailabilityMode != nil {
		eDevice.AvailabilityMode = iDevice.AvailabilityMode
	}
	if iDevice.AvailabilityTemplate != nil {
		eDevice.AvailabilityTemplate = iDevice.AvailabilityTemplate
	}
	if iDevice.Availability != nil {
		eDevice.AvailabilityFunc = common.ConstructStateFunc(*iDevice.Availability)
	}
	if iDevice.CurrentTemperatureTemplate != nil {
		eDevice.CurrentTemperatureTemplate = iDevice.CurrentTemperatureTemplate
	}
	if iDevice.CurrentTemperature != nil {
		eDevice.CurrentTemperatureFunc = common.ConstructStateFunc(*iDevice.CurrentTemperature)
	}
	if iDevice.EnabledByDefault != nil {
		eDevice.EnabledByDefault = iDevice.EnabledByDefault
	}
	if iDevice.Encoding != nil {
		eDevice.Encoding = iDevice.Encoding
	}
	if iDevice.EntityCategory != nil {
		eDevice.EntityCategory = iDevice.EntityCategory
	}
	if iDevice.Icon != nil {
		eDevice.Icon = iDevice.Icon
	}
	if iDevice.Initial != nil {
		eDevice.Initial = iDevice.Initial
	}
	if iDevice.JsonAttributesTemplate != nil {
		eDevice.JsonAttributesTemplate = iDevice.JsonAttributesTemplate
	}
	if iDevice.JsonAttributes != nil {
		eDevice.JsonAttributesFunc = common.ConstructCommandFunc(*iDevice.JsonAttributes)
	}
	if iDevice.MaxTemp != nil {
		eDevice.MaxTemp = iDevice.MaxTemp
	}
	if iDevice.MinTemp != nil {
		eDevice.MinTemp = iDevice.MinTemp
	}
	if iDevice.ModeCommandTemplate != nil {
		eDevice.ModeCommandTemplate = iDevice.ModeCommandTemplate
	}
	if iDevice.ModeCommand != nil {
		eDevice.ModeCommandFunc = common.ConstructCommandFunc(*iDevice.ModeCommand)
	}
	if iDevice.ModeStateTemplate != nil {
		eDevice.ModeStateTemplate = iDevice.ModeStateTemplate
	}
	if iDevice.ModeState != nil {
		eDevice.ModeStateFunc = common.ConstructStateFunc(*iDevice.ModeState)
	}
	if iDevice.Modes != nil {
		eDevice.Modes = iDevice.Modes
	}
	if iDevice.Name != nil {
		eDevice.Name = iDevice.Name
	}
	if iDevice.ObjectId != nil {
		eDevice.ObjectId = iDevice.ObjectId
	}
	if iDevice.Optimistic != nil {
		eDevice.Optimistic = iDevice.Optimistic
	}
	if iDevice.PayloadAvailable != nil {
		eDevice.PayloadAvailable = iDevice.PayloadAvailable
	}
	if iDevice.PayloadNotAvailable != nil {
		eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	}
	if iDevice.PayloadOff != nil {
		eDevice.PayloadOff = iDevice.PayloadOff
	}
	if iDevice.PayloadOn != nil {
		eDevice.PayloadOn = iDevice.PayloadOn
	}
	if iDevice.PowerCommandTemplate != nil {
		eDevice.PowerCommandTemplate = iDevice.PowerCommandTemplate
	}
	if iDevice.PowerCommand != nil {
		eDevice.PowerCommandFunc = common.ConstructCommandFunc(*iDevice.PowerCommand)
	}
	if iDevice.Precision != nil {
		eDevice.Precision = iDevice.Precision
	}
	if iDevice.Qos != nil {
		eDevice.Qos = iDevice.Qos
	}
	if iDevice.Retain != nil {
		eDevice.Retain = iDevice.Retain
	}
	if iDevice.TemperatureCommandTemplate != nil {
		eDevice.TemperatureCommandTemplate = iDevice.TemperatureCommandTemplate
	}
	if iDevice.TemperatureCommand != nil {
		eDevice.TemperatureCommandFunc = common.ConstructCommandFunc(*iDevice.TemperatureCommand)
	}
	if iDevice.TemperatureStateTemplate != nil {
		eDevice.TemperatureStateTemplate = iDevice.TemperatureStateTemplate
	}
	if iDevice.TemperatureState != nil {
		eDevice.TemperatureStateFunc = common.ConstructStateFunc(*iDevice.TemperatureState)
	}
	if iDevice.TemperatureUnit != nil {
		eDevice.TemperatureUnit = iDevice.TemperatureUnit
	}
	if iDevice.UniqueId != nil {
		eDevice.UniqueId = iDevice.UniqueId
	}
	if iDevice.ValueTemplate != nil {
		eDevice.ValueTemplate = iDevice.ValueTemplate
	}
	if iDevice.Availability == nil {
		eDevice.AvailabilityFunc = common.AvailabilityFunc
	}
	eDevice.Initialize()
	return eDevice
}
