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
type WaterHeater struct {
	AvailabilityMode           *string       `json:"availability_mode,omitempty"`            // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate       *string       `json:"availability_template,omitempty"`        // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic          *string       `json:"availability_topic,omitempty"`           // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	AvailabilityFunc           func() string `json:"-"`                                      // Function for availability
	CurrentTemperatureTemplate *string       `json:"current_temperature_template,omitempty"` // "A template with which the value received on `current_temperature_topic` will be rendered."
	CurrentTemperatureTopic    *string       `json:"current_temperature_topic,omitempty"`    // "The MQTT topic on which to listen for the current temperature. A `\"None\"` value received will reset the current temperature. Empty values (`'''`) will be ignored."
	CurrentTemperatureFunc     func() string `json:"-"`                                      // Function for current temperature
	Device                     struct {
		ConfigurationUrl *string `json:"configuration_url,omitempty"` // "A link to the webpage that can manage the configuration of this device. Can be either an HTTP or HTTPS link."
		Connections      *string `json:"connections,omitempty"`       // "A list of connections of the device to the outside world as a list of tuples `[connection_type, connection_identifier]`. For example the MAC address of a network interface: `\"connections\": [[\"mac\", \"02:5b:26:a8:dc:12\"]]`."
		Identifiers      *string `json:"identifiers,omitempty"`       // "A list of IDs that uniquely identify the device. For example a serial number."
		Manufacturer     *string `json:"manufacturer,omitempty"`      // "The manufacturer of the device."
		Model            *string `json:"model,omitempty"`             // "The model of the device."
		Name             *string `json:"name,omitempty"`              // "The name of the device."
		SuggestedArea    *string `json:"suggested_area,omitempty"`    // "Suggest an area if the device isn’t in one yet."
		SwVersion        *string `json:"sw_version,omitempty"`        // "The firmware version of the device."
		ViaDevice        *string `json:"via_device,omitempty"`        // "Identifier of a device that routes messages between this device and Home Assistant. Examples of such devices are hubs, or parent devices of a sub-device. This is used to show device topology in Home Assistant."
	} `json:"device,omitempty"` // Device configuration parameters
	EnabledByDefault           *bool                           `json:"enabled_by_default,omitempty"`           // "Flag which defines if the entity should be enabled when first added."
	Encoding                   *string                         `json:"encoding,omitempty"`                     // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory             *string                         `json:"entity_category,omitempty"`              // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                       *string                         `json:"icon,omitempty"`                         // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Initial                    *int                            `json:"initial,omitempty"`                      // "Set the initial target temperature. The default value depends on the temperature unit, and will be 43.3°C or 110°F."
	JsonAttributesTemplate     *string                         `json:"json_attributes_template,omitempty"`     // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic        *string                         `json:"json_attributes_topic,omitempty"`        // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	JsonAttributesFunc         func(mqtt.Message, mqtt.Client) `json:"-"`                                      // Function for json attributes
	MaxTemp                    *float64                        `json:"max_temp,omitempty"`                     // "Maximum set point available. The default value depends on the temperature unit, and will be 60°C or 140°F."
	MinTemp                    *float64                        `json:"min_temp,omitempty"`                     // "Minimum set point available. The default value depends on the temperature unit, and will be 43.3°C or 110°F."
	ModeCommandTemplate        *string                         `json:"mode_command_template,omitempty"`        // "A template to render the value sent to the `mode_command_topic` with."
	ModeCommandTopic           *string                         `json:"mode_command_topic,omitempty"`           // "The MQTT topic to publish commands to change the Water Heater operation mode. Use with `mode_command_template` if you only want to publish the power state."
	ModeCommandFunc            func(mqtt.Message, mqtt.Client) `json:"-"`                                      // Function for mode command
	ModeStateTemplate          *string                         `json:"mode_state_template,omitempty"`          // "A template to render the value received on the `mode_state_topic` with."
	ModeStateTopic             *string                         `json:"mode_state_topic,omitempty"`             // "The MQTT topic to subscribe for changes of the Water Heater operation mode. If this is not set, the operation mode works in optimistic mode (see below)."
	ModeStateFunc              func() string                   `json:"-"`                                      // Function for mode state
	Modes                      *([]string)                     `json:"modes,omitempty"`                        // "A list of supported modes. Needs to be a subset of the default values."
	Name                       *string                         `json:"name,omitempty"`                         // "The name of the Water Heater."
	ObjectId                   *string                         `json:"object_id,omitempty"`                    // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic                 *bool                           `json:"optimistic,omitempty"`                   // "Flag that defines if the water heater works in optimistic mode"
	PayloadAvailable           *string                         `json:"payload_available,omitempty"`            // "The payload that represents the available state."
	PayloadNotAvailable        *string                         `json:"payload_not_available,omitempty"`        // "The payload that represents the unavailable state."
	PayloadOff                 *string                         `json:"payload_off,omitempty"`                  // "The payload that represents disabled state."
	PayloadOn                  *string                         `json:"payload_on,omitempty"`                   // "The payload that represents enabled state."
	Precision                  *float64                        `json:"precision,omitempty"`                    // "The desired precision for this device. Can be used to match your actual water heater's precision. Supported values are `0.1`, `0.5` and `1.0`."
	Qos                        *int                            `json:"qos,omitempty"`                          // "The maximum QoS level to be used when receiving and publishing messages."
	Retain                     *bool                           `json:"retain,omitempty"`                       // "Defines if published messages should have the retain flag set."
	TemperatureCommandTemplate *string                         `json:"temperature_command_template,omitempty"` // "A template to render the value sent to the `temperature_command_topic` with."
	TemperatureCommandTopic    *string                         `json:"temperature_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the target temperature."
	TemperatureCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`                                      // Function for temperature command
	TemperatureStateTemplate   *string                         `json:"temperature_state_template,omitempty"`   // "A template to render the value received on the `temperature_state_topic` with."
	TemperatureStateTopic      *string                         `json:"temperature_state_topic,omitempty"`      // "The MQTT topic to subscribe for changes in the target temperature. If this is not set, the target temperature works in optimistic mode (see below). A `\"None\"` value received will reset the temperature set point. Empty values (`'''`) will be ignored."
	TemperatureStateFunc       func() string                   `json:"-"`                                      // Function for temperature state
	TemperatureUnit            *string                         `json:"temperature_unit,omitempty"`             // "Defines the temperature unit of the device, `C` or `F`. If this is not set, the temperature unit is set to the system temperature unit."
	UniqueId                   *string                         `json:"unique_id,omitempty"`                    // "An ID that uniquely identifies this Water Heater device. If two Water Heater devices have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate              *string                         `json:"value_template,omitempty"`               // "Default template to render the payloads on *all* `*_state_topic`s with."
	MQTT                       *MQTTFields                     `json:"-"`                                      // MQTT configuration parameters
}

func (d *WaterHeater) GetRawId() string {
	return "water_heater"
}
func (d *WaterHeater) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *WaterHeater) GetUniqueId() string {
	return *d.UniqueId
}
func (d *WaterHeater) PopulateDevice() {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &common.MachineID
}
func (d *WaterHeater) UpdateState() {
	if d.AvailabilityTopic != nil {
		state := d.AvailabilityFunc()
		if state != stateStore.WaterHeater.Availability[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.WaterHeater.Availability[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.CurrentTemperatureTopic != nil {
		state := d.CurrentTemperatureFunc()
		if state != stateStore.WaterHeater.CurrentTemperature[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.CurrentTemperatureTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.WaterHeater.CurrentTemperature[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.ModeStateTopic != nil {
		state := d.ModeStateFunc()
		if state != stateStore.WaterHeater.ModeState[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.ModeStateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.WaterHeater.ModeState[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.TemperatureStateTopic != nil {
		state := d.TemperatureStateFunc()
		if state != stateStore.WaterHeater.TemperatureState[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.TemperatureStateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.WaterHeater.TemperatureState[d.GetUniqueId()] = state
			token.Wait()
		}
	}
}
func (d *WaterHeater) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	if d.JsonAttributesTopic != nil {
		t := c.Subscribe(*d.JsonAttributesTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.ModeCommandTopic != nil {
		t := c.Subscribe(*d.ModeCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TemperatureCommandTopic != nil {
		t := c.Subscribe(*d.TemperatureCommandTopic, 0, d.MQTT.MessageHandler)
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
func (d *WaterHeater) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "offline")
	token.Wait()
	if d.JsonAttributesTopic != nil {
		t := c.Unsubscribe(*d.JsonAttributesTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.ModeCommandTopic != nil {
		t := c.Unsubscribe(*d.ModeCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.TemperatureCommandTopic != nil {
		t := c.Unsubscribe(*d.TemperatureCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d *WaterHeater) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d *WaterHeater) Initialize() {
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
func (d *WaterHeater) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = new(string)
		*d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CurrentTemperatureFunc != nil {
		d.CurrentTemperatureTopic = new(string)
		*d.CurrentTemperatureTopic = GetTopic(d, "current_temperature_topic")
	}
	if d.JsonAttributesFunc != nil {
		d.JsonAttributesTopic = new(string)
		*d.JsonAttributesTopic = GetTopic(d, "json_attributes_topic")
		store.TopicStore[*d.JsonAttributesTopic] = &d.JsonAttributesFunc
	}
	if d.ModeCommandFunc != nil {
		d.ModeCommandTopic = new(string)
		*d.ModeCommandTopic = GetTopic(d, "mode_command_topic")
		store.TopicStore[*d.ModeCommandTopic] = &d.ModeCommandFunc
	}
	if d.ModeStateFunc != nil {
		d.ModeStateTopic = new(string)
		*d.ModeStateTopic = GetTopic(d, "mode_state_topic")
	}
	if d.TemperatureCommandFunc != nil {
		d.TemperatureCommandTopic = new(string)
		*d.TemperatureCommandTopic = GetTopic(d, "temperature_command_topic")
		store.TopicStore[*d.TemperatureCommandTopic] = &d.TemperatureCommandFunc
	}
	if d.TemperatureStateFunc != nil {
		d.TemperatureStateTopic = new(string)
		*d.TemperatureStateTopic = GetTopic(d, "temperature_state_topic")
	}
}
func (d *WaterHeater) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *WaterHeater) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
