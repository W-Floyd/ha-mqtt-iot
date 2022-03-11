package hadiscovery

import mqtt "github.com/eclipse/paho.mqtt.golang"

/////////////////// Components of config

type device struct {
	Connections  map[string]string `json:"cns,omitempty"`
	Identifiers  []string          `json:"ids,omitempty"`
	Manufacturer string            `json:"mf,omitempty"`
	Model        string            `json:"mdl,omitempty"`
	Name         string            `json:"name,omitempty"`
	SWVersion    string            `json:"sw,omitempty"`
	ViaDevice    string            `json:"via_device,omitempty"`
}

///////////////////

// BinarySensor matches a binary_sensor
type BinarySensor struct {
	AvailabilityTopic      string `json:"availability_topic,omitempty"`
	Device                 device `json:"device,omitempty"`
	DeviceClass            string `json:"device_class,omitempty"`
	ExpireAfter            int    `json:"expire_after,omitempty"`
	ForceUpdate            bool   `json:"force_update,omitempty"`
	JSONAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic    string `json:"json_attributes_topic,omitempty"`
	Name                   string `json:"name,omitempty"`
	OffDelay               int    `json:"off_delay,omitempty"`
	PayloadAvailable       string `json:"payload_available,omitempty"`
	PayloadNotAvailable    string `json:"payload_not_available,omitempty"`
	PayloadOff             string `json:"pl_off,omitempty"`
	PayloadOn              string `json:"pl_on,omitempty"`
	QOS                    int    `json:"qos,omitempty"`
	StateTopic             string `json:"state_topic"`
	UniqueID               string `json:"unique_id,omitempty"`
	ValueTemplate          string `json:"value_template,omitempty"`

	StateFunc func() string `json:"-"`

	UpdateInterval  float64 `json:"-"`
	ForceUpdateMQTT bool    `json:"-"`
}

// Light matches a light
type Light struct {
	AvailabilityTopic        string   `json:"availability_topic,omitempty"`
	BrightnessCommandTopic   string   `json:"brightness_command_topic,omitempty"`
	BrightnessScale          int      `json:"brightness_scale,omitempty"`
	BrightnessStateTopic     string   `json:"brightness_state_topic,omitempty"`
	BrightnessValueTemplate  string   `json:"brightness_value_template,omitempty"`
	ColorTempCommandTemplate string   `json:"color_temp_command_template,omitempty"`
	ColorTempCommandTopic    string   `json:"color_temp_command_topic,omitempty"`
	ColorTempStateTopic      string   `json:"color_temp_state_topic,omitempty"`
	ColorTempValueTemplate   string   `json:"color_temp_value_template,omitempty"`
	CommandTopic             string   `json:"command_topic"`
	Device                   device   `json:"device,omitempty"`
	EffectCommandTopic       string   `json:"effect_command_topic,omitempty"`
	EffectList               []string `json:"effect_list,omitempty"`
	EffectStateTopic         string   `json:"effect_state_topic,omitempty"`
	EffectValueTemplate      string   `json:"effect_value_template,omitempty"`
	HsCommandTopic           string   `json:"hs_command_topic,omitempty"`
	HsStateTopic             string   `json:"hs_state_topic,omitempty"`
	HsValueTemplate          string   `json:"hs_value_template,omitempty"`
	JSONAttributesTemplate   string   `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic      string   `json:"json_attributes_topic,omitempty"`
	MaxMireds                int      `json:"max_mireds,omitempty"`
	MinMireds                int      `json:"min_mireds,omitempty"`
	Name                     string   `json:"name,omitempty"`
	OnCommandType            string   `json:"on_command_type,omitempty"`
	Optimistic               bool     `json:"opt,omitempty"`
	PayloadAvailable         string   `json:"payload_available,omitempty"`
	PayloadNotAvailable      string   `json:"payload_not_available,omitempty"`
	PayloadOff               string   `json:"pl_off,omitempty"`
	PayloadOn                string   `json:"pl_on,omitempty"`
	QOS                      int      `json:"qos,omitempty"`
	Retain                   bool     `json:"ret,omitempty"`
	RgbCommandTemplate       string   `json:"rgb_command_template,omitempty"`
	RgbCommandTopic          string   `json:"rgb_command_topic,omitempty"`
	RgbStateTopic            string   `json:"rgb_state_topic,omitempty"`
	RgbValueTemplate         string   `json:"rgb_value_template,omitempty"`
	Schema                   string   `json:"schema,omitempty"`
	StateTopic               string   `json:"state_topic,omitempty"`
	StateValueTemplate       string   `json:"state_value_template,omitempty"`
	UniqueID                 string   `json:"unique_id,omitempty"`
	WhiteValueCommandTopic   string   `json:"white_value_command_topic,omitempty"`
	WhiteValueScale          int      `json:"white_value_scale,omitempty"`
	WhiteValueStateTopic     string   `json:"white_value_state_topic,omitempty"`
	WhiteValueTemplate       string   `json:"white_value_template,omitempty"`
	XyCommandTopic           string   `json:"xy_command_topic,omitempty"`
	XyStateTopic             string   `json:"xy_state_topic,omitempty"`
	XyValueTemplate          string   `json:"xy_value_template,omitempty"`
	ValueTemplate            string   `json:"value_template,omitempty"`

	BrightnessStateFunc func() string `json:"-"`
	ColorTempStateFunc  func() string `json:"-"`
	EffectStateFunc     func() string `json:"-"`
	HsStateFunc         func() string `json:"-"`
	RgbStateFunc        func() string `json:"-"`
	StateFunc           func() string `json:"-"`
	WhiteValueStateFunc func() string `json:"-"`
	XyStateFunc         func() string `json:"-"`

	BrightnessCommandFunc func(mqtt.Message, mqtt.Client) `json:"-"`
	ColorTempCommandFunc  func(mqtt.Message, mqtt.Client) `json:"-"`
	CommandFunc           func(mqtt.Message, mqtt.Client) `json:"-"`
	EffectCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	HsCommandFunc         func(mqtt.Message, mqtt.Client) `json:"-"`
	RgbCommandFunc        func(mqtt.Message, mqtt.Client) `json:"-"`
	WhiteValueCommandFunc func(mqtt.Message, mqtt.Client) `json:"-"`
	XyCommandFunc         func(mqtt.Message, mqtt.Client) `json:"-"`

	UpdateInterval  float64 `json:"-"`
	ForceUpdateMQTT bool    `json:"-"`

	messageHandler mqtt.MessageHandler
}

// Sensor matches a sensor
type Sensor struct {
	AvailabilityTopic      string `json:"availability_topic,omitempty"`
	Device                 device `json:"device,omitempty"`
	DeviceClass            string `json:"device_class,omitempty"`
	ExpireAfter            int    `json:"expire_after,omitempty"`
	ForceUpdate            bool   `json:"force_update,omitempty"`
	Icon                   string `json:"icon,omitempty"`
	JSONAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JSONAttributesTopic    string `json:"json_attributes_topic,omitempty"`
	Name                   string `json:"name,omitempty"`
	PayloadAvailable       string `json:"payload_available,omitempty"`
	PayloadNotAvailable    string `json:"payload_not_available,omitempty"`
	QOS                    int    `json:"qos,omitempty"`
	StateTopic             string `json:"state_topic"`
	UniqueID               string `json:"unique_id,omitempty"`
	UnitOfMeasurement      string `json:"unit_of_measurement,omitempty"`
	ValueTemplate          string `json:"value_template,omitempty"`

	StateFunc func() string `json:"-"`

	UpdateInterval  float64 `json:"-"`
	ForceUpdateMQTT bool    `json:"-"`
}

// Switch matches a switch
type Switch struct {
	AvailabilityTopic      string `json:"avty_t,omitempty"`
	CommandTopic           string `json:"cmd_t"`
	Device                 device `json:"dev,omitempty"`
	Icon                   string `json:"ic,omitempty"`
	JSONAttributesTemplate string `json:"json_attr_tpl,omitempty"`
	JSONAttributesTopic    string `json:"json_attr_t,omitempty"`
	Name                   string `json:"name,omitempty"`
	Optimistic             bool   `json:"opt,omitempty"`
	PayloadAvailable       string `json:"pl_avail,omitempty"`
	PayloadNotAvailable    string `json:"pl_not_avail,omitempty"`
	PayloadOff             string `json:"pl_off,omitempty"`
	PayloadOn              string `json:"pl_on,omitempty"`
	QOS                    int    `json:"qos,omitempty"`
	Retain                 bool   `json:"ret,omitempty"`
	StateOff               string `json:"stat_off,omitempty"`
	StateOn                string `json:"stat_on,omitempty"`
	StateTopic             string `json:"stat_t,omitempty"`
	UniqueID               string `json:"uniq_id,omitempty"`
	ValueTemplate          string `json:"val_tpl,omitempty"`

	CommandFunc func(mqtt.Message, mqtt.Client) `json:"-"`
	StateFunc   func() string                   `json:"-"`

	UpdateInterval  float64 `json:"-"`
	ForceUpdateMQTT bool    `json:"-"`

	messageHandler mqtt.MessageHandler
}
