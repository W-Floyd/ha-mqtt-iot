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
type Light struct {
	AvailabilityMode          *string                         `json:"availability_mode,omitempty"`     // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate      *string                         `json:"availability_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	AvailabilityTopic         *string                         `json:"availability_topic,omitempty"`    // "The MQTT topic subscribed to receive availability (online/offline) updates. Must not be used together with `availability`."
	AvailabilityFunc          func() string                   `json:"-"`
	BrightnessCommandTemplate *string                         `json:"brightness_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `brightness_command_topic`. Available variables: `value`."
	BrightnessCommandTopic    *string                         `json:"brightness_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the light’s brightness."
	BrightnessCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	BrightnessScale           *int                            `json:"brightness_scale,omitempty"`       // "Defines the maximum brightness value (i.e., 100%) of the MQTT device."
	BrightnessStateTopic      *string                         `json:"brightness_state_topic,omitempty"` // "The MQTT topic subscribed to receive brightness state updates."
	BrightnessStateFunc       func() string                   `json:"-"`
	BrightnessValueTemplate   *string                         `json:"brightness_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the brightness value."
	ColorModeStateTopic       *string                         `json:"color_mode_state_topic,omitempty"`    // "The MQTT topic subscribed to receive color mode updates. If this is not configured, `color_mode` will be automatically set according to the last received valid color or color temperature"
	ColorModeStateFunc        func() string                   `json:"-"`
	ColorModeValueTemplate    *string                         `json:"color_mode_value_template,omitempty"`   // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the color mode."
	ColorTempCommandTemplate  *string                         `json:"color_temp_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `color_temp_command_topic`. Available variables: `value`."
	ColorTempCommandTopic     *string                         `json:"color_temp_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the light’s color temperature state. The color temperature command slider has a range of 153 to 500 mireds (micro reciprocal degrees)."
	ColorTempCommandFunc      func(mqtt.Message, mqtt.Client) `json:"-"`
	ColorTempStateTopic       *string                         `json:"color_temp_state_topic,omitempty"` // "The MQTT topic subscribed to receive color temperature state updates."
	ColorTempStateFunc        func() string                   `json:"-"`
	ColorTempValueTemplate    *string                         `json:"color_temp_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the color temperature value."
	CommandTopic              *string                         `json:"command_topic,omitempty"`             // "The MQTT topic to publish commands to change the switch state."
	CommandFunc               func(mqtt.Message, mqtt.Client) `json:"-"`
	Device                    struct {
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
	EffectCommandTemplate  *string                         `json:"effect_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `effect_command_topic`. Available variables: `value`."
	EffectCommandTopic     *string                         `json:"effect_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the light's effect state."
	EffectCommandFunc      func(mqtt.Message, mqtt.Client) `json:"-"`
	EffectList             *([]string)                     `json:"effect_list,omitempty"`        // "The list of effects the light supports."
	EffectStateTopic       *string                         `json:"effect_state_topic,omitempty"` // "The MQTT topic subscribed to receive effect state updates."
	EffectStateFunc        func() string                   `json:"-"`
	EffectValueTemplate    *string                         `json:"effect_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the effect value."
	EnabledByDefault       *bool                           `json:"enabled_by_default,omitempty"`    // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string                         `json:"encoding,omitempty"`              // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string                         `json:"entity_category,omitempty"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	HsCommandTemplate      *string                         `json:"hs_command_template,omitempty"`   // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `hs_command_topic`. Available variables: `hue` and `sat`."
	HsCommandTopic         *string                         `json:"hs_command_topic,omitempty"`      // "The MQTT topic to publish commands to change the light's color state in HS format (Hue Saturation). Range for Hue: 0° .. 360°, Range of Saturation: 0..100. Note: Brightness is sent separately in the `brightness_command_topic`."
	HsCommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
	HsStateTopic           *string                         `json:"hs_state_topic,omitempty"` // "The MQTT topic subscribed to receive color state updates in HS format. The expected payload is the hue and saturation values separated by commas, for example, `359.5,100.0`. Note: Brightness is received separately in the `brightness_state_topic`."
	HsStateFunc            func() string                   `json:"-"`
	HsValueTemplate        *string                         `json:"hs_value_template,omitempty"`        // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the HS value."
	Icon                   *string                         `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string                         `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration) documentation."
	JsonAttributesTopic    *string                         `json:"json_attributes_topic,omitempty"`    // "The MQTT topic subscribed to receive a JSON dictionary payload and then set as sensor attributes. Usage example can be found in [MQTT sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration) documentation."
	JsonAttributesFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	MaxMireds              *int                            `json:"max_mireds,omitempty"`            // "The maximum color temperature in mireds."
	MinMireds              *int                            `json:"min_mireds,omitempty"`            // "The minimum color temperature in mireds."
	Name                   *string                         `json:"name,omitempty"`                  // "The name of the light."
	ObjectId               *string                         `json:"object_id,omitempty"`             // "Used instead of `name` for automatic generation of `entity_id`"
	OnCommandType          *string                         `json:"on_command_type,omitempty"`       // "Defines when on the payload_on is sent. Using `last` (the default) will send any style (brightness, color, etc) topics first and then a `payload_on` to the `command_topic`. Using `first` will send the `payload_on` and then any style topics. Using `brightness` will only send brightness commands instead of the `payload_on` to turn the light on."
	Optimistic             *bool                           `json:"optimistic,omitempty"`            // "Flag that defines if switch works in optimistic mode."
	PayloadAvailable       *string                         `json:"payload_available,omitempty"`     // "The payload that represents the available state."
	PayloadNotAvailable    *string                         `json:"payload_not_available,omitempty"` // "The payload that represents the unavailable state."
	PayloadOff             *string                         `json:"payload_off,omitempty"`           // "The payload that represents disabled state."
	PayloadOn              *string                         `json:"payload_on,omitempty"`            // "The payload that represents enabled state."
	Qos                    *int                            `json:"qos,omitempty"`                   // "The maximum QoS level of the state topic."
	Retain                 *bool                           `json:"retain,omitempty"`                // "If the published message should have the retain flag on or not."
	RgbCommandTemplate     *string                         `json:"rgb_command_template,omitempty"`  // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `rgb_command_topic`. Available variables: `red`, `green` and `blue`."
	RgbCommandTopic        *string                         `json:"rgb_command_topic,omitempty"`     // "The MQTT topic to publish commands to change the light's RGB state."
	RgbCommandFunc         func(mqtt.Message, mqtt.Client) `json:"-"`
	RgbStateTopic          *string                         `json:"rgb_state_topic,omitempty"` // "The MQTT topic subscribed to receive RGB state updates. The expected payload is the RGB values separated by commas, for example, `255,0,127`."
	RgbStateFunc           func() string                   `json:"-"`
	RgbValueTemplate       *string                         `json:"rgb_value_template,omitempty"`    // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the RGB value."
	RgbwCommandTemplate    *string                         `json:"rgbw_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `rgbw_command_topic`. Available variables: `red`, `green`, `blue` and `white`."
	RgbwCommandTopic       *string                         `json:"rgbw_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the light's RGBW state."
	RgbwCommandFunc        func(mqtt.Message, mqtt.Client) `json:"-"`
	RgbwStateTopic         *string                         `json:"rgbw_state_topic,omitempty"` // "The MQTT topic subscribed to receive RGBW state updates. The expected payload is the RGBW values separated by commas, for example, `255,0,127,64`."
	RgbwStateFunc          func() string                   `json:"-"`
	RgbwValueTemplate      *string                         `json:"rgbw_value_template,omitempty"`    // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the RGBW value."
	RgbwwCommandTemplate   *string                         `json:"rgbww_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `rgbww_command_topic`. Available variables: `red`, `green`, `blue`, `cold_white` and `warm_white`."
	RgbwwCommandTopic      *string                         `json:"rgbww_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the light's RGBWW state."
	RgbwwCommandFunc       func(mqtt.Message, mqtt.Client) `json:"-"`
	RgbwwStateTopic        *string                         `json:"rgbww_state_topic,omitempty"` // "The MQTT topic subscribed to receive RGBWW state updates. The expected payload is the RGBWW values separated by commas, for example, `255,0,127,64,32`."
	RgbwwStateFunc         func() string                   `json:"-"`
	RgbwwValueTemplate     *string                         `json:"rgbww_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the RGBWW value."
	Schema                 *string                         `json:"schema,omitempty"`               // "The schema to use. Must be `default` or omitted to select the default schema."
	StateTopic             *string                         `json:"state_topic,omitempty"`          // "The MQTT topic subscribed to receive state updates."
	StateFunc              func() string                   `json:"-"`
	StateValueTemplate     *string                         `json:"state_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the state value. The template should match the payload `on` and `off` values, so if your light uses `power on` to turn on, your `state_value_template` string should return `power on` when the switch is on. For example if the message is just `on`, your `state_value_template` should be `power {{ value }}`."
	UniqueId               *string                         `json:"unique_id,omitempty"`            // "An ID that uniquely identifies this light. If two lights have the same unique ID, Home Assistant will raise an exception."
	WhiteCommandTopic      *string                         `json:"white_command_topic,omitempty"`  // "The MQTT topic to publish commands to change the light to white mode with a given brightness."
	WhiteCommandFunc       func(mqtt.Message, mqtt.Client) `json:"-"`
	WhiteScale             *int                            `json:"white_scale,omitempty"`         // "Defines the maximum white level (i.e., 100%) of the MQTT device."
	XyCommandTemplate      *string                         `json:"xy_command_template,omitempty"` // "Defines a [template](/docs/configuration/templating/) to compose message which will be sent to `xy_command_topic`. Available variables: `x` and `y`."
	XyCommandTopic         *string                         `json:"xy_command_topic,omitempty"`    // "The MQTT topic to publish commands to change the light's XY state."
	XyCommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
	XyStateTopic           *string                         `json:"xy_state_topic,omitempty"` // "The MQTT topic subscribed to receive XY state updates. The expected payload is the X and Y color values separated by commas, for example, `0.675,0.322`."
	XyStateFunc            func() string                   `json:"-"`
	XyValueTemplate        *string                         `json:"xy_value_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the XY value."
	MQTT                   *MQTTFields                     `json:"-"`                           // MQTT configuration parameters
}

func (d *Light) GetRawId() string {
	return "light"
}
func (d *Light) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d *Light) GetUniqueId() string {
	return *d.UniqueId
}
func (d *Light) PopulateDevice() {
	d.Device.Manufacturer = &Manufacturer
	d.Device.Model = &SoftwareName
	d.Device.Name = &InstanceName
	d.Device.SwVersion = &SWVersion
	d.Device.Identifiers = &common.MachineID
}
func (d *Light) UpdateState() {
	if d.AvailabilityTopic != nil {
		state := d.AvailabilityFunc()
		if state != stateStore.Light.Availability[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.AvailabilityTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Light.Availability[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.BrightnessStateTopic != nil {
		state := d.BrightnessStateFunc()
		if state != stateStore.Light.BrightnessState[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.BrightnessStateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Light.BrightnessState[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.ColorModeStateTopic != nil {
		state := d.ColorModeStateFunc()
		if state != stateStore.Light.ColorModeState[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.ColorModeStateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Light.ColorModeState[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.ColorTempStateTopic != nil {
		state := d.ColorTempStateFunc()
		if state != stateStore.Light.ColorTempState[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.ColorTempStateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Light.ColorTempState[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.EffectStateTopic != nil {
		state := d.EffectStateFunc()
		if state != stateStore.Light.EffectState[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.EffectStateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Light.EffectState[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.HsStateTopic != nil {
		state := d.HsStateFunc()
		if state != stateStore.Light.HsState[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.HsStateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Light.HsState[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.RgbStateTopic != nil {
		state := d.RgbStateFunc()
		if state != stateStore.Light.RgbState[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.RgbStateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Light.RgbState[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.RgbwStateTopic != nil {
		state := d.RgbwStateFunc()
		if state != stateStore.Light.RgbwState[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.RgbwStateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Light.RgbwState[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.RgbwwStateTopic != nil {
		state := d.RgbwwStateFunc()
		if state != stateStore.Light.RgbwwState[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.RgbwwStateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Light.RgbwwState[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.StateTopic != nil {
		state := d.StateFunc()
		if state != stateStore.Light.State[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.StateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Light.State[d.GetUniqueId()] = state
			token.Wait()
		}
	}
	if d.XyStateTopic != nil {
		state := d.XyStateFunc()
		if state != stateStore.Light.XyState[d.GetUniqueId()] || (d.MQTT.ForceUpdate != nil && *d.MQTT.ForceUpdate) {
			token := (*d.MQTT.Client).Publish(*d.XyStateTopic, byte(*d.Qos), *d.Retain, state)
			stateStore.Light.XyState[d.GetUniqueId()] = state
			token.Wait()
		}
	}
}
func (d *Light) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	if d.BrightnessCommandTopic != nil {
		t := c.Subscribe(*d.BrightnessCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.ColorTempCommandTopic != nil {
		t := c.Subscribe(*d.ColorTempCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.CommandTopic != nil {
		t := c.Subscribe(*d.CommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.EffectCommandTopic != nil {
		t := c.Subscribe(*d.EffectCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.HsCommandTopic != nil {
		t := c.Subscribe(*d.HsCommandTopic, 0, d.MQTT.MessageHandler)
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
	if d.RgbCommandTopic != nil {
		t := c.Subscribe(*d.RgbCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.RgbwCommandTopic != nil {
		t := c.Subscribe(*d.RgbwCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.RgbwwCommandTopic != nil {
		t := c.Subscribe(*d.RgbwwCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.WhiteCommandTopic != nil {
		t := c.Subscribe(*d.WhiteCommandTopic, 0, d.MQTT.MessageHandler)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.XyCommandTopic != nil {
		t := c.Subscribe(*d.XyCommandTopic, 0, d.MQTT.MessageHandler)
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
func (d *Light) UnSubscribe() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "offline")
	token.Wait()
	if d.BrightnessCommandTopic != nil {
		t := c.Unsubscribe(*d.BrightnessCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.ColorTempCommandTopic != nil {
		t := c.Unsubscribe(*d.ColorTempCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.CommandTopic != nil {
		t := c.Unsubscribe(*d.CommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.EffectCommandTopic != nil {
		t := c.Unsubscribe(*d.EffectCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.HsCommandTopic != nil {
		t := c.Unsubscribe(*d.HsCommandTopic)
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
	if d.RgbCommandTopic != nil {
		t := c.Unsubscribe(*d.RgbCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.RgbwCommandTopic != nil {
		t := c.Unsubscribe(*d.RgbwCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.RgbwwCommandTopic != nil {
		t := c.Unsubscribe(*d.RgbwwCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.WhiteCommandTopic != nil {
		t := c.Unsubscribe(*d.WhiteCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
	if d.XyCommandTopic != nil {
		t := c.Unsubscribe(*d.XyCommandTopic)
		t.Wait()
		if t.Error() != nil {
			log.Fatal(t.Error())
		}
	}
}
func (d *Light) AnnounceAvailable() {
	c := *d.MQTT.Client
	token := c.Publish(*d.AvailabilityTopic, common.QoS, common.Retain, "online")
	token.Wait()
}
func (d *Light) Initialize() {
	if d.Qos == nil {
		d.Qos = new(int)
		*d.Qos = int(common.QoS)
	}
	if d.Retain == nil {
		d.Retain = new(bool)
		*d.Retain = common.Retain
	}
	if d.UniqueId == nil {
		d.UniqueId = new(string)
		*d.UniqueId = strcase.ToDelimited(*d.Name, uint8(0x2d))
	}
	d.PopulateDevice()
	d.AddMessageHandler()
	d.PopulateTopics()
}
func (d *Light) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = new(string)
		*d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.BrightnessCommandFunc != nil {
		d.BrightnessCommandTopic = new(string)
		*d.BrightnessCommandTopic = GetTopic(d, "brightness_command_topic")
		store.TopicStore[*d.BrightnessCommandTopic] = &d.BrightnessCommandFunc
	}
	if d.BrightnessStateFunc != nil {
		d.BrightnessStateTopic = new(string)
		*d.BrightnessStateTopic = GetTopic(d, "brightness_state_topic")
	}
	if d.ColorModeStateFunc != nil {
		d.ColorModeStateTopic = new(string)
		*d.ColorModeStateTopic = GetTopic(d, "color_mode_state_topic")
	}
	if d.ColorTempCommandFunc != nil {
		d.ColorTempCommandTopic = new(string)
		*d.ColorTempCommandTopic = GetTopic(d, "color_temp_command_topic")
		store.TopicStore[*d.ColorTempCommandTopic] = &d.ColorTempCommandFunc
	}
	if d.ColorTempStateFunc != nil {
		d.ColorTempStateTopic = new(string)
		*d.ColorTempStateTopic = GetTopic(d, "color_temp_state_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = new(string)
		*d.CommandTopic = GetTopic(d, "command_topic")
		store.TopicStore[*d.CommandTopic] = &d.CommandFunc
	}
	if d.EffectCommandFunc != nil {
		d.EffectCommandTopic = new(string)
		*d.EffectCommandTopic = GetTopic(d, "effect_command_topic")
		store.TopicStore[*d.EffectCommandTopic] = &d.EffectCommandFunc
	}
	if d.EffectStateFunc != nil {
		d.EffectStateTopic = new(string)
		*d.EffectStateTopic = GetTopic(d, "effect_state_topic")
	}
	if d.HsCommandFunc != nil {
		d.HsCommandTopic = new(string)
		*d.HsCommandTopic = GetTopic(d, "hs_command_topic")
		store.TopicStore[*d.HsCommandTopic] = &d.HsCommandFunc
	}
	if d.HsStateFunc != nil {
		d.HsStateTopic = new(string)
		*d.HsStateTopic = GetTopic(d, "hs_state_topic")
	}
	if d.JsonAttributesFunc != nil {
		d.JsonAttributesTopic = new(string)
		*d.JsonAttributesTopic = GetTopic(d, "json_attributes_topic")
		store.TopicStore[*d.JsonAttributesTopic] = &d.JsonAttributesFunc
	}
	if d.RgbCommandFunc != nil {
		d.RgbCommandTopic = new(string)
		*d.RgbCommandTopic = GetTopic(d, "rgb_command_topic")
		store.TopicStore[*d.RgbCommandTopic] = &d.RgbCommandFunc
	}
	if d.RgbStateFunc != nil {
		d.RgbStateTopic = new(string)
		*d.RgbStateTopic = GetTopic(d, "rgb_state_topic")
	}
	if d.RgbwCommandFunc != nil {
		d.RgbwCommandTopic = new(string)
		*d.RgbwCommandTopic = GetTopic(d, "rgbw_command_topic")
		store.TopicStore[*d.RgbwCommandTopic] = &d.RgbwCommandFunc
	}
	if d.RgbwStateFunc != nil {
		d.RgbwStateTopic = new(string)
		*d.RgbwStateTopic = GetTopic(d, "rgbw_state_topic")
	}
	if d.RgbwwCommandFunc != nil {
		d.RgbwwCommandTopic = new(string)
		*d.RgbwwCommandTopic = GetTopic(d, "rgbww_command_topic")
		store.TopicStore[*d.RgbwwCommandTopic] = &d.RgbwwCommandFunc
	}
	if d.RgbwwStateFunc != nil {
		d.RgbwwStateTopic = new(string)
		*d.RgbwwStateTopic = GetTopic(d, "rgbww_state_topic")
	}
	if d.StateFunc != nil {
		d.StateTopic = new(string)
		*d.StateTopic = GetTopic(d, "state_topic")
	}
	if d.WhiteCommandFunc != nil {
		d.WhiteCommandTopic = new(string)
		*d.WhiteCommandTopic = GetTopic(d, "white_command_topic")
		store.TopicStore[*d.WhiteCommandTopic] = &d.WhiteCommandFunc
	}
	if d.XyCommandFunc != nil {
		d.XyCommandTopic = new(string)
		*d.XyCommandTopic = GetTopic(d, "xy_command_topic")
		store.TopicStore[*d.XyCommandTopic] = &d.XyCommandFunc
	}
	if d.XyStateFunc != nil {
		d.XyStateTopic = new(string)
		*d.XyStateTopic = GetTopic(d, "xy_state_topic")
	}
}
func (d *Light) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Light) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
