package hadiscovery

import mqtt "github.com/eclipse/paho.mqtt.golang"

type AlarmControlPanel struct {
	AvailabilityMode     string                          `json:"availability_mode"`
	AvailabilityTemplate string                          `json:"availability_template"`
	AvailabilityTopic    string                          `json:"availability_topic"`
	AvailabilityFunc     func() string                   `json:"-"`
	Code                 string                          `json:"code"`
	CodeArmRequired      bool                            `json:"code_arm_required"`
	CodeDisarmRequired   bool                            `json:"code_disarm_required"`
	CodeTriggerRequired  bool                            `json:"code_trigger_required"`
	CommandTemplate      string                          `json:"command_template"`
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
	EnabledByDefault       bool          `json:"enabled_by_default"`
	Encoding               string        `json:"encoding"`
	EntityCategory         string        `json:"entity_category"`
	Icon                   string        `json:"icon"`
	Name                   string        `json:"name"`
	ObjectId               string        `json:"object_id"`
	PayloadArmAway         string        `json:"payload_arm_away"`
	PayloadArmCustomBypass string        `json:"payload_arm_custom_bypass"`
	PayloadArmHome         string        `json:"payload_arm_home"`
	PayloadArmNight        string        `json:"payload_arm_night"`
	PayloadArmVacation     string        `json:"payload_arm_vacation"`
	PayloadAvailable       string        `json:"payload_available"`
	PayloadDisarm          string        `json:"payload_disarm"`
	PayloadNotAvailable    string        `json:"payload_not_available"`
	PayloadTrigger         string        `json:"payload_trigger"`
	Qos                    int           `json:"qos"`
	Retain                 bool          `json:"retain"`
	StateTopic             string        `json:"state_topic"`
	StateFunc              func() string `json:"-"`
	UniqueId               string        `json:"unique_id"`
	ValueTemplate          string        `json:"value_template"`
	RawId                  string        `json:"-"`
}
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
}
type Button struct {
	AvailabilityMode     string                          `json:"availability_mode"`
	AvailabilityTemplate string                          `json:"availability_template"`
	AvailabilityTopic    string                          `json:"availability_topic"`
	AvailabilityFunc     func() string                   `json:"-"`
	CommandTemplate      string                          `json:"command_template"`
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
	DeviceClass         string `json:"device_class"`
	EnabledByDefault    bool   `json:"enabled_by_default"`
	Encoding            string `json:"encoding"`
	EntityCategory      string `json:"entity_category"`
	Icon                string `json:"icon"`
	Name                string `json:"name"`
	ObjectId            string `json:"object_id"`
	PayloadAvailable    string `json:"payload_available"`
	PayloadNotAvailable string `json:"payload_not_available"`
	PayloadPress        string `json:"payload_press"`
	Qos                 int    `json:"qos"`
	Retain              bool   `json:"retain"`
	UniqueId            string `json:"unique_id"`
	RawId               string `json:"-"`
}
type Camera struct {
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
	EnabledByDefault bool   `json:"enabled_by_default"`
	EntityCategory   string `json:"entity_category"`
	Icon             string `json:"icon"`
	Name             string `json:"name"`
	ObjectId         string `json:"object_id"`
	Topic            string `json:"topic"`
	UniqueId         string `json:"unique_id"`
	RawId            string `json:"-"`
}
type Cover struct {
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
	DeviceClass         string                          `json:"device_class"`
	EnabledByDefault    bool                            `json:"enabled_by_default"`
	Encoding            string                          `json:"encoding"`
	EntityCategory      string                          `json:"entity_category"`
	Icon                string                          `json:"icon"`
	Name                string                          `json:"name"`
	ObjectId            string                          `json:"object_id"`
	Optimistic          bool                            `json:"optimistic"`
	PayloadAvailable    string                          `json:"payload_available"`
	PayloadClose        string                          `json:"payload_close"`
	PayloadNotAvailable string                          `json:"payload_not_available"`
	PayloadOpen         string                          `json:"payload_open"`
	PayloadStop         string                          `json:"payload_stop"`
	PositionClosed      int                             `json:"position_closed"`
	PositionOpen        int                             `json:"position_open"`
	PositionTemplate    string                          `json:"position_template"`
	PositionTopic       string                          `json:"position_topic"`
	PositionFunc        func() string                   `json:"-"`
	Qos                 int                             `json:"qos"`
	Retain              bool                            `json:"retain"`
	SetPositionTemplate string                          `json:"set_position_template"`
	SetPositionTopic    string                          `json:"set_position_topic"`
	SetPositionFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	StateClosed         string                          `json:"state_closed"`
	StateClosing        string                          `json:"state_closing"`
	StateOpen           string                          `json:"state_open"`
	StateOpening        string                          `json:"state_opening"`
	StateStopped        string                          `json:"state_stopped"`
	StateTopic          string                          `json:"state_topic"`
	StateFunc           func() string                   `json:"-"`
	TiltClosedValue     int                             `json:"tilt_closed_value"`
	TiltCommandTemplate string                          `json:"tilt_command_template"`
	TiltCommandTopic    string                          `json:"tilt_command_topic"`
	TiltCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	TiltMax             int                             `json:"tilt_max"`
	TiltMin             int                             `json:"tilt_min"`
	TiltOpenedValue     int                             `json:"tilt_opened_value"`
	TiltOptimistic      bool                            `json:"tilt_optimistic"`
	TiltStatusTemplate  string                          `json:"tilt_status_template"`
	TiltStatusTopic     string                          `json:"tilt_status_topic"`
	TiltStatusFunc      func() string                   `json:"-"`
	UniqueId            string                          `json:"unique_id"`
	ValueTemplate       string                          `json:"value_template"`
	RawId               string                          `json:"-"`
}
type DeviceTracker struct {
	Devices        []string `json:"devices"`
	PayloadHome    string   `json:"payload_home"`
	PayloadNotHome string   `json:"payload_not_home"`
	Qos            int      `json:"qos"`
	SourceType     string   `json:"source_type"`
	RawId          string   `json:"-"`
}
type DeviceTrigger struct {
	AutomationType string `json:"automation_type"`
	Device         struct {
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
	Payload       string `json:"payload"`
	Qos           int    `json:"qos"`
	Subtype       string `json:"subtype"`
	Topic         string `json:"topic"`
	Type          string `json:"type"`
	ValueTemplate string `json:"value_template"`
	RawId         string `json:"-"`
}
type Fan struct {
	AvailabilityMode     string                          `json:"availability_mode"`
	AvailabilityTemplate string                          `json:"availability_template"`
	AvailabilityTopic    string                          `json:"availability_topic"`
	AvailabilityFunc     func() string                   `json:"-"`
	CommandTemplate      string                          `json:"command_template"`
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
	EnabledByDefault           bool                            `json:"enabled_by_default"`
	Encoding                   string                          `json:"encoding"`
	EntityCategory             string                          `json:"entity_category"`
	Icon                       string                          `json:"icon"`
	Name                       string                          `json:"name"`
	ObjectId                   string                          `json:"object_id"`
	Optimistic                 bool                            `json:"optimistic"`
	OscillationCommandTemplate string                          `json:"oscillation_command_template"`
	OscillationCommandTopic    string                          `json:"oscillation_command_topic"`
	OscillationCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	OscillationStateTopic      string                          `json:"oscillation_state_topic"`
	OscillationStateFunc       func() string                   `json:"-"`
	OscillationValueTemplate   string                          `json:"oscillation_value_template"`
	PayloadAvailable           string                          `json:"payload_available"`
	PayloadNotAvailable        string                          `json:"payload_not_available"`
	PayloadOff                 string                          `json:"payload_off"`
	PayloadOn                  string                          `json:"payload_on"`
	PayloadOscillationOff      string                          `json:"payload_oscillation_off"`
	PayloadOscillationOn       string                          `json:"payload_oscillation_on"`
	PayloadResetPercentage     string                          `json:"payload_reset_percentage"`
	PayloadResetPresetMode     string                          `json:"payload_reset_preset_mode"`
	PercentageCommandTemplate  string                          `json:"percentage_command_template"`
	PercentageCommandTopic     string                          `json:"percentage_command_topic"`
	PercentageCommandFunc      func(mqtt.Message, mqtt.Client) `json:"-"`
	PercentageStateTopic       string                          `json:"percentage_state_topic"`
	PercentageStateFunc        func() string                   `json:"-"`
	PercentageValueTemplate    string                          `json:"percentage_value_template"`
	PresetModeCommandTemplate  string                          `json:"preset_mode_command_template"`
	PresetModeCommandTopic     string                          `json:"preset_mode_command_topic"`
	PresetModeCommandFunc      func(mqtt.Message, mqtt.Client) `json:"-"`
	PresetModeStateTopic       string                          `json:"preset_mode_state_topic"`
	PresetModeStateFunc        func() string                   `json:"-"`
	PresetModeValueTemplate    string                          `json:"preset_mode_value_template"`
	PresetModes                []string                        `json:"preset_modes"`
	Qos                        int                             `json:"qos"`
	Retain                     bool                            `json:"retain"`
	SpeedRangeMax              int                             `json:"speed_range_max"`
	SpeedRangeMin              int                             `json:"speed_range_min"`
	StateTopic                 string                          `json:"state_topic"`
	StateFunc                  func() string                   `json:"-"`
	StateValueTemplate         string                          `json:"state_value_template"`
	UniqueId                   string                          `json:"unique_id"`
	RawId                      string                          `json:"-"`
}
type Humidifier struct {
	AvailabilityMode     string                          `json:"availability_mode"`
	AvailabilityTemplate string                          `json:"availability_template"`
	AvailabilityTopic    string                          `json:"availability_topic"`
	AvailabilityFunc     func() string                   `json:"-"`
	CommandTemplate      string                          `json:"command_template"`
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
	DeviceClass                   string                          `json:"device_class"`
	EnabledByDefault              bool                            `json:"enabled_by_default"`
	Encoding                      string                          `json:"encoding"`
	EntityCategory                string                          `json:"entity_category"`
	Icon                          string                          `json:"icon"`
	MaxHumidity                   int                             `json:"max_humidity"`
	MinHumidity                   int                             `json:"min_humidity"`
	ModeCommandTemplate           string                          `json:"mode_command_template"`
	ModeCommandTopic              string                          `json:"mode_command_topic"`
	ModeCommandFunc               func(mqtt.Message, mqtt.Client) `json:"-"`
	ModeStateTemplate             string                          `json:"mode_state_template"`
	ModeStateTopic                string                          `json:"mode_state_topic"`
	ModeStateFunc                 func() string                   `json:"-"`
	Modes                         []string                        `json:"modes"`
	Name                          string                          `json:"name"`
	ObjectId                      string                          `json:"object_id"`
	Optimistic                    bool                            `json:"optimistic"`
	PayloadAvailable              string                          `json:"payload_available"`
	PayloadNotAvailable           string                          `json:"payload_not_available"`
	PayloadOff                    string                          `json:"payload_off"`
	PayloadOn                     string                          `json:"payload_on"`
	PayloadResetHumidity          string                          `json:"payload_reset_humidity"`
	PayloadResetMode              string                          `json:"payload_reset_mode"`
	Qos                           int                             `json:"qos"`
	Retain                        bool                            `json:"retain"`
	StateTopic                    string                          `json:"state_topic"`
	StateFunc                     func() string                   `json:"-"`
	StateValueTemplate            string                          `json:"state_value_template"`
	TargetHumidityCommandTemplate string                          `json:"target_humidity_command_template"`
	TargetHumidityCommandTopic    string                          `json:"target_humidity_command_topic"`
	TargetHumidityCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	TargetHumidityStateTemplate   string                          `json:"target_humidity_state_template"`
	TargetHumidityStateTopic      string                          `json:"target_humidity_state_topic"`
	TargetHumidityStateFunc       func() string                   `json:"-"`
	UniqueId                      string                          `json:"unique_id"`
	RawId                         string                          `json:"-"`
}
type Climate struct {
	ActionTemplate             string                          `json:"action_template"`
	ActionTopic                string                          `json:"action_topic"`
	ActionFunc                 func(mqtt.Message, mqtt.Client) `json:"-"`
	AuxCommandTopic            string                          `json:"aux_command_topic"`
	AuxCommandFunc             func(mqtt.Message, mqtt.Client) `json:"-"`
	AuxStateTemplate           string                          `json:"aux_state_template"`
	AuxStateTopic              string                          `json:"aux_state_topic"`
	AuxStateFunc               func() string                   `json:"-"`
	AvailabilityMode           string                          `json:"availability_mode"`
	AvailabilityTemplate       string                          `json:"availability_template"`
	AvailabilityTopic          string                          `json:"availability_topic"`
	AvailabilityFunc           func() string                   `json:"-"`
	CurrentTemperatureTemplate string                          `json:"current_temperature_template"`
	CurrentTemperatureTopic    string                          `json:"current_temperature_topic"`
	CurrentTemperatureFunc     func() string                   `json:"-"`
	Device                     struct {
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
	EnabledByDefault               bool                            `json:"enabled_by_default"`
	Encoding                       string                          `json:"encoding"`
	EntityCategory                 string                          `json:"entity_category"`
	FanModeCommandTemplate         string                          `json:"fan_mode_command_template"`
	FanModeCommandTopic            string                          `json:"fan_mode_command_topic"`
	FanModeCommandFunc             func(mqtt.Message, mqtt.Client) `json:"-"`
	FanModeStateTemplate           string                          `json:"fan_mode_state_template"`
	FanModeStateTopic              string                          `json:"fan_mode_state_topic"`
	FanModeStateFunc               func() string                   `json:"-"`
	FanModes                       []string                        `json:"fan_modes"`
	Icon                           string                          `json:"icon"`
	Initial                        int                             `json:"initial"`
	MaxTemp                        float64                         `json:"max_temp"`
	MinTemp                        float64                         `json:"min_temp"`
	ModeCommandTemplate            string                          `json:"mode_command_template"`
	ModeCommandTopic               string                          `json:"mode_command_topic"`
	ModeCommandFunc                func(mqtt.Message, mqtt.Client) `json:"-"`
	ModeStateTemplate              string                          `json:"mode_state_template"`
	ModeStateTopic                 string                          `json:"mode_state_topic"`
	ModeStateFunc                  func() string                   `json:"-"`
	Modes                          []string                        `json:"modes"`
	Name                           string                          `json:"name"`
	ObjectId                       string                          `json:"object_id"`
	PayloadAvailable               string                          `json:"payload_available"`
	PayloadNotAvailable            string                          `json:"payload_not_available"`
	PayloadOff                     string                          `json:"payload_off"`
	PayloadOn                      string                          `json:"payload_on"`
	PowerCommandTopic              string                          `json:"power_command_topic"`
	PowerCommandFunc               func(mqtt.Message, mqtt.Client) `json:"-"`
	Precision                      float64                         `json:"precision"`
	PresetModeCommandTemplate      string                          `json:"preset_mode_command_template"`
	PresetModeCommandTopic         string                          `json:"preset_mode_command_topic"`
	PresetModeCommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
	PresetModeStateTopic           string                          `json:"preset_mode_state_topic"`
	PresetModeStateFunc            func() string                   `json:"-"`
	PresetModeValueTemplate        string                          `json:"preset_mode_value_template"`
	PresetModes                    []string                        `json:"preset_modes"`
	Qos                            int                             `json:"qos"`
	Retain                         bool                            `json:"retain"`
	SwingModeCommandTemplate       string                          `json:"swing_mode_command_template"`
	SwingModeCommandTopic          string                          `json:"swing_mode_command_topic"`
	SwingModeCommandFunc           func(mqtt.Message, mqtt.Client) `json:"-"`
	SwingModeStateTemplate         string                          `json:"swing_mode_state_template"`
	SwingModeStateTopic            string                          `json:"swing_mode_state_topic"`
	SwingModeStateFunc             func() string                   `json:"-"`
	SwingModes                     []string                        `json:"swing_modes"`
	TempStep                       float64                         `json:"temp_step"`
	TemperatureCommandTemplate     string                          `json:"temperature_command_template"`
	TemperatureCommandTopic        string                          `json:"temperature_command_topic"`
	TemperatureCommandFunc         func(mqtt.Message, mqtt.Client) `json:"-"`
	TemperatureHighCommandTemplate string                          `json:"temperature_high_command_template"`
	TemperatureHighCommandTopic    string                          `json:"temperature_high_command_topic"`
	TemperatureHighCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	TemperatureHighStateTemplate   string                          `json:"temperature_high_state_template"`
	TemperatureHighStateTopic      string                          `json:"temperature_high_state_topic"`
	TemperatureHighStateFunc       func() string                   `json:"-"`
	TemperatureLowCommandTemplate  string                          `json:"temperature_low_command_template"`
	TemperatureLowCommandTopic     string                          `json:"temperature_low_command_topic"`
	TemperatureLowCommandFunc      func(mqtt.Message, mqtt.Client) `json:"-"`
	TemperatureLowStateTemplate    string                          `json:"temperature_low_state_template"`
	TemperatureLowStateTopic       string                          `json:"temperature_low_state_topic"`
	TemperatureLowStateFunc        func() string                   `json:"-"`
	TemperatureStateTemplate       string                          `json:"temperature_state_template"`
	TemperatureStateTopic          string                          `json:"temperature_state_topic"`
	TemperatureStateFunc           func() string                   `json:"-"`
	TemperatureUnit                string                          `json:"temperature_unit"`
	UniqueId                       string                          `json:"unique_id"`
	ValueTemplate                  string                          `json:"value_template"`
	RawId                          string                          `json:"-"`
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
}
type Lock struct {
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
	EnabledByDefault    bool          `json:"enabled_by_default"`
	Encoding            string        `json:"encoding"`
	EntityCategory      string        `json:"entity_category"`
	Icon                string        `json:"icon"`
	Name                string        `json:"name"`
	ObjectId            string        `json:"object_id"`
	Optimistic          bool          `json:"optimistic"`
	PayloadAvailable    string        `json:"payload_available"`
	PayloadLock         string        `json:"payload_lock"`
	PayloadNotAvailable string        `json:"payload_not_available"`
	PayloadOpen         string        `json:"payload_open"`
	PayloadUnlock       string        `json:"payload_unlock"`
	Qos                 int           `json:"qos"`
	Retain              bool          `json:"retain"`
	StateLocked         string        `json:"state_locked"`
	StateTopic          string        `json:"state_topic"`
	StateFunc           func() string `json:"-"`
	StateUnlocked       string        `json:"state_unlocked"`
	UniqueId            string        `json:"unique_id"`
	ValueTemplate       string        `json:"value_template"`
	RawId               string        `json:"-"`
}
type Number struct {
	AvailabilityMode  string                          `json:"availability_mode"`
	AvailabilityTopic string                          `json:"availability_topic"`
	AvailabilityFunc  func() string                   `json:"-"`
	CommandTemplate   string                          `json:"command_template"`
	CommandTopic      string                          `json:"command_topic"`
	CommandFunc       func(mqtt.Message, mqtt.Client) `json:"-"`
	Device            struct {
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
	EnabledByDefault  bool          `json:"enabled_by_default"`
	Encoding          string        `json:"encoding"`
	EntityCategory    string        `json:"entity_category"`
	Icon              string        `json:"icon"`
	Max               float64       `json:"max"`
	Min               float64       `json:"min"`
	Name              string        `json:"name"`
	ObjectId          string        `json:"object_id"`
	Optimistic        bool          `json:"optimistic"`
	PayloadReset      string        `json:"payload_reset"`
	Qos               int           `json:"qos"`
	Retain            bool          `json:"retain"`
	StateTopic        string        `json:"state_topic"`
	StateFunc         func() string `json:"-"`
	Step              float64       `json:"step"`
	UniqueId          string        `json:"unique_id"`
	UnitOfMeasurement string        `json:"unit_of_measurement"`
	ValueTemplate     string        `json:"value_template"`
	RawId             string        `json:"-"`
}
type Scene struct {
	AvailabilityMode     string                          `json:"availability_mode"`
	AvailabilityTemplate string                          `json:"availability_template"`
	AvailabilityTopic    string                          `json:"availability_topic"`
	AvailabilityFunc     func() string                   `json:"-"`
	CommandTopic         string                          `json:"command_topic"`
	CommandFunc          func(mqtt.Message, mqtt.Client) `json:"-"`
	EnabledByDefault     bool                            `json:"enabled_by_default"`
	EntityCategory       string                          `json:"entity_category"`
	Icon                 string                          `json:"icon"`
	Name                 string                          `json:"name"`
	ObjectId             string                          `json:"object_id"`
	PayloadAvailable     string                          `json:"payload_available"`
	PayloadNotAvailable  string                          `json:"payload_not_available"`
	PayloadOn            string                          `json:"payload_on"`
	Qos                  int                             `json:"qos"`
	Retain               bool                            `json:"retain"`
	UniqueId             string                          `json:"unique_id"`
	RawId                string                          `json:"-"`
}
type Select struct {
	AvailabilityMode     string                          `json:"availability_mode"`
	AvailabilityTemplate string                          `json:"availability_template"`
	AvailabilityTopic    string                          `json:"availability_topic"`
	AvailabilityFunc     func() string                   `json:"-"`
	CommandTemplate      string                          `json:"command_template"`
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
	EnabledByDefault bool          `json:"enabled_by_default"`
	Encoding         string        `json:"encoding"`
	EntityCategory   string        `json:"entity_category"`
	Icon             string        `json:"icon"`
	Name             string        `json:"name"`
	ObjectId         string        `json:"object_id"`
	Optimistic       bool          `json:"optimistic"`
	Options          []string      `json:"options"`
	Qos              int           `json:"qos"`
	Retain           bool          `json:"retain"`
	StateTopic       string        `json:"state_topic"`
	StateFunc        func() string `json:"-"`
	UniqueId         string        `json:"unique_id"`
	ValueTemplate    string        `json:"value_template"`
	RawId            string        `json:"-"`
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
}
type Siren struct {
	AvailabilityMode     string                          `json:"availability_mode"`
	AvailabilityTemplate string                          `json:"availability_template"`
	AvailabilityTopic    string                          `json:"availability_topic"`
	AvailabilityFunc     func() string                   `json:"-"`
	AvailableTones       []string                        `json:"available_tones"`
	CommandOffTemplate   string                          `json:"command_off_template"`
	CommandTemplate      string                          `json:"command_template"`
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
	StateValueTemplate  string        `json:"state_value_template"`
	SupportDuration     bool          `json:"support_duration"`
	SupportVolumeSet    bool          `json:"support_volume_set"`
	UniqueId            string        `json:"unique_id"`
	RawId               string        `json:"-"`
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
}
type Tag struct {
	Device struct {
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
	Topic         string `json:"topic"`
	ValueTemplate string `json:"value_template"`
	RawId         string `json:"-"`
}
type Vacuum struct {
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
	Encoding            string                          `json:"encoding"`
	FanSpeedList        []string                        `json:"fan_speed_list"`
	Name                string                          `json:"name"`
	ObjectId            string                          `json:"object_id"`
	PayloadAvailable    string                          `json:"payload_available"`
	PayloadCleanSpot    string                          `json:"payload_clean_spot"`
	PayloadLocate       string                          `json:"payload_locate"`
	PayloadNotAvailable string                          `json:"payload_not_available"`
	PayloadPause        string                          `json:"payload_pause"`
	PayloadReturnToBase string                          `json:"payload_return_to_base"`
	PayloadStart        string                          `json:"payload_start"`
	PayloadStop         string                          `json:"payload_stop"`
	Qos                 int                             `json:"qos"`
	Retain              bool                            `json:"retain"`
	Schema              string                          `json:"schema"`
	SendCommandTopic    string                          `json:"send_command_topic"`
	SendCommandFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	SetFanSpeedTopic    string                          `json:"set_fan_speed_topic"`
	SetFanSpeedFunc     func(mqtt.Message, mqtt.Client) `json:"-"`
	StateTopic          string                          `json:"state_topic"`
	StateFunc           func() string                   `json:"-"`
	SupportedFeatures   []string                        `json:"supported_features"`
	UniqueId            string                          `json:"unique_id"`
	RawId               string                          `json:"-"`
}
type Device interface {
	GetRawId() string
	GetUniqueId() string
	PopulateDevice()
	PopulateTopics()
}
