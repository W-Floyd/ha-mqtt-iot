package hadiscovery

import mqtt "github.com/eclipse/paho.mqtt.golang"

type BinarySensor struct {
	AvailabilityMode     string        `json:"availability_mode"`
	AvailabilityTemplate string        `json:"availability_template"`
	AvailabilityTopic    string        `json:"availability_topic"`
	AvailabilityFunc     func() string `json:"-"`
	Device               struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	DeviceClass         string        `json:"device_class"`
	EnabledByDefault    bool          `json:"enabled_by_default"`
	Encoding            string        `json:"encoding"`
	EntityCategory      string        `json:"entity_category"`
	ExpireAfter         int           `json:"expire_after"`
	ForceUpdate         bool          `json:"force_update"`
	Icon                string        `json:"icon"`
	Name                string        `json:"name"`
	ObjectId            string        `json:"object_id"`
	OffDelay            int           `json:"off_delay"`
	PayloadAvailable    string        `json:"payload_available"`
	PayloadNotAvailable string        `json:"payload_not_available"`
	PayloadOff          string        `json:"payload_off"`
	PayloadOn           string        `json:"payload_on"`
	Qos                 int           `json:"qos"`
	StateTopic          string        `json:"state_topic"`
	StateFunc           func() string `json:"-"`
	UniqueId            string        `json:"unique_id"`
	ValueTemplate       string        `json:"value_template"`
	RawId               string        `json:"-"`
	MQTTClient          *mqtt.Client  `json:"-"`
}
type Light struct {
	AvailabilityMode          string                          `json:"availability_mode"`
	AvailabilityTemplate      string                          `json:"availability_template"`
	AvailabilityTopic         string                          `json:"availability_topic"`
	AvailabilityFunc          func() string                   `json:"-"`
	BrightnessCommandTemplate string                          `json:"brightness_command_template"`
	BrightnessCommandTopic    string                          `json:"brightness_command_topic"`
	BrightnessCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	BrightnessScale           int                             `json:"brightness_scale"`
	BrightnessStateTopic      string                          `json:"brightness_state_topic"`
	BrightnessStateFunc       func() string                   `json:"-"`
	BrightnessValueTemplate   string                          `json:"brightness_value_template"`
	ColorModeStateTopic       string                          `json:"color_mode_state_topic"`
	ColorModeStateFunc        func() string                   `json:"-"`
	ColorModeValueTemplate    string                          `json:"color_mode_value_template"`
	ColorTempCommandTemplate  string                          `json:"color_temp_command_template"`
	ColorTempCommandTopic     string                          `json:"color_temp_command_topic"`
	ColorTempCommandFunc      func(mqtt.Message, mqtt.Client) `json:"-"`
	ColorTempStateTopic       string                          `json:"color_temp_state_topic"`
	ColorTempStateFunc        func() string                   `json:"-"`
	ColorTempValueTemplate    string                          `json:"color_temp_value_template"`
	CommandTopic              string                          `json:"command_topic"`
	CommandFunc               func(mqtt.Message, mqtt.Client) `json:"-"`
	Device                    struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	EffectCommandTemplate string                          `json:"effect_command_template"`
	EffectCommandTopic    string                          `json:"effect_command_topic"`
	EffectCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	EffectList            []string                        `json:"effect_list"`
	EffectStateTopic      string                          `json:"effect_state_topic"`
	EffectStateFunc       func() string                   `json:"-"`
	EffectValueTemplate   string                          `json:"effect_value_template"`
	EnabledByDefault      bool                            `json:"enabled_by_default"`
	Encoding              string                          `json:"encoding"`
	EntityCategory        string                          `json:"entity_category"`
	HsCommandTopic        string                          `json:"hs_command_topic"`
	HsCommandFunc         func(mqtt.Message, mqtt.Client) `json:"-"`
	HsStateTopic          string                          `json:"hs_state_topic"`
	HsStateFunc           func() string                   `json:"-"`
	HsValueTemplate       string                          `json:"hs_value_template"`
	Icon                  string                          `json:"icon"`
	MaxMireds             int                             `json:"max_mireds"`
	MinMireds             int                             `json:"min_mireds"`
	Name                  string                          `json:"name"`
	ObjectId              string                          `json:"object_id"`
	OnCommandType         string                          `json:"on_command_type"`
	Optimistic            bool                            `json:"optimistic"`
	PayloadAvailable      string                          `json:"payload_available"`
	PayloadNotAvailable   string                          `json:"payload_not_available"`
	PayloadOff            string                          `json:"payload_off"`
	PayloadOn             string                          `json:"payload_on"`
	Qos                   int                             `json:"qos"`
	Retain                bool                            `json:"retain"`
	RgbCommandTemplate    string                          `json:"rgb_command_template"`
	RgbCommandTopic       string                          `json:"rgb_command_topic"`
	RgbCommandFunc        func(mqtt.Message, mqtt.Client) `json:"-"`
	RgbStateTopic         string                          `json:"rgb_state_topic"`
	RgbStateFunc          func() string                   `json:"-"`
	RgbValueTemplate      string                          `json:"rgb_value_template"`
	Schema                string                          `json:"schema"`
	StateTopic            string                          `json:"state_topic"`
	StateFunc             func() string                   `json:"-"`
	StateValueTemplate    string                          `json:"state_value_template"`
	UniqueId              string                          `json:"unique_id"`
	WhiteCommandTopic     string                          `json:"white_command_topic"`
	WhiteCommandFunc      func(mqtt.Message, mqtt.Client) `json:"-"`
	WhiteScale            int                             `json:"white_scale"`
	XyCommandTopic        string                          `json:"xy_command_topic"`
	XyCommandFunc         func(mqtt.Message, mqtt.Client) `json:"-"`
	XyStateTopic          string                          `json:"xy_state_topic"`
	XyStateFunc           func() string                   `json:"-"`
	XyValueTemplate       string                          `json:"xy_value_template"`
	RawId                 string                          `json:"-"`
	MQTTClient            *mqtt.Client                    `json:"-"`
}
type Sensor struct {
	AvailabilityMode     string        `json:"availability_mode"`
	AvailabilityTemplate string        `json:"availability_template"`
	AvailabilityTopic    string        `json:"availability_topic"`
	AvailabilityFunc     func() string `json:"-"`
	Device               struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	DeviceClass            string        `json:"device_class"`
	EnabledByDefault       bool          `json:"enabled_by_default"`
	Encoding               string        `json:"encoding"`
	EntityCategory         string        `json:"entity_category"`
	ExpireAfter            int           `json:"expire_after"`
	ForceUpdate            bool          `json:"force_update"`
	Icon                   string        `json:"icon"`
	LastResetValueTemplate string        `json:"last_reset_value_template"`
	Name                   string        `json:"name"`
	ObjectId               string        `json:"object_id"`
	PayloadAvailable       string        `json:"payload_available"`
	PayloadNotAvailable    string        `json:"payload_not_available"`
	Qos                    int           `json:"qos"`
	StateClass             string        `json:"state_class"`
	StateTopic             string        `json:"state_topic"`
	StateFunc              func() string `json:"-"`
	UniqueId               string        `json:"unique_id"`
	UnitOfMeasurement      string        `json:"unit_of_measurement"`
	ValueTemplate          string        `json:"value_template"`
	RawId                  string        `json:"-"`
	MQTTClient             *mqtt.Client  `json:"-"`
}
type Switch struct {
	AvailabilityMode     string                          `json:"availability_mode"`
	AvailabilityTemplate string                          `json:"availability_template"`
	AvailabilityTopic    string                          `json:"availability_topic"`
	AvailabilityFunc     func() string                   `json:"-"`
	CommandTopic         string                          `json:"command_topic"`
	CommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
	Device               struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	DeviceClass         string        `json:"device_class"`
	EnabledByDefault    bool          `json:"enabled_by_default"`
	Encoding            string        `json:"encoding"`
	EntityCategory      string        `json:"entity_category"`
	Icon                string        `json:"icon"`
	Name                string        `json:"name"`
	ObjectId            string        `json:"object_id"`
	Optimistic          bool          `json:"optimistic"`
	PayloadAvailable    string        `json:"payload_available"`
	PayloadNotAvailable string        `json:"payload_not_available"`
	PayloadOff          string        `json:"payload_off"`
	PayloadOn           string        `json:"payload_on"`
	Qos                 int           `json:"qos"`
	Retain              bool          `json:"retain"`
	StateOff            string        `json:"state_off"`
	StateOn             string        `json:"state_on"`
	StateTopic          string        `json:"state_topic"`
	StateFunc           func() string `json:"-"`
	UniqueId            string        `json:"unique_id"`
	ValueTemplate       string        `json:"value_template"`
	RawId               string        `json:"-"`
	MQTTClient          *mqtt.Client  `json:"-"`
}
type Device interface {
	GetRawId() string
	GetUniqueId() string
	PopulateDevice()
	PopulateTopics()
	UpdateState(*mqtt.Client)
	Subscribe(*mqtt.Client)
}
