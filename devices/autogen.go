package devices

type HADeviceAlarmControlPanel struct {
	Availability struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	AvailabilityMode   string `json:"availability_mode,omitempty"`
	AvailabilityTopic  string `json:"availability_topic,omitempty"`
	Code               string `json:"code,omitempty"`
	CodeArmRequired    bool   `json:"code_arm_required,omitempty"`
	CodeDisarmRequired bool   `json:"code_disarm_required,omitempty"`
	CommandTemplate    string `json:"command_template,omitempty"`
	CommandTopic       string `json:"command_topic,omitempty"`
	Device             struct {
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	EnabledByDefault       bool   `json:"enabled_by_default,omitempty"`
	Icon                   string `json:"icon,omitempty"`
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JsonAttributesTopic    string `json:"json_attributes_topic,omitempty"`
	Name                   string `json:"name,omitempty"`
	PayloadArmAway         string `json:"payload_arm_away,omitempty"`
	PayloadArmCustomBypass string `json:"payload_arm_custom_bypass,omitempty"`
	PayloadArmHome         string `json:"payload_arm_home,omitempty"`
	PayloadArmNight        string `json:"payload_arm_night,omitempty"`
	PayloadArmVacation     string `json:"payload_arm_vacation,omitempty"`
	PayloadAvailable       string `json:"payload_available,omitempty"`
	PayloadDisarm          string `json:"payload_disarm,omitempty"`
	PayloadNotAvailable    string `json:"payload_not_available,omitempty"`
	Qos                    int    `json:"qos,omitempty"`
	Retain                 bool   `json:"retain,omitempty"`
	StateTopic             string `json:"state_topic,omitempty"`
	UniqueId               string `json:"unique_id,omitempty"`
	ValueTemplate          string `json:"value_template,omitempty"`
}

type HADeviceBinarySensor struct {
	Availability struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	AvailabilityMode  string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	Device            struct {
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	DeviceClass            string `json:"device_class,omitempty"`
	EnabledByDefault       bool   `json:"enabled_by_default,omitempty"`
	ExpireAfter            int    `json:"expire_after,omitempty"`
	ForceUpdate            bool   `json:"force_update,omitempty"`
	Icon                   string `json:"icon,omitempty"`
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JsonAttributesTopic    string `json:"json_attributes_topic,omitempty"`
	Name                   string `json:"name,omitempty"`
	OffDelay               int    `json:"off_delay,omitempty"`
	PayloadAvailable       string `json:"payload_available,omitempty"`
	PayloadNotAvailable    string `json:"payload_not_available,omitempty"`
	PayloadOff             string `json:"payload_off,omitempty"`
	PayloadOn              string `json:"payload_on,omitempty"`
	Qos                    int    `json:"qos,omitempty"`
	StateTopic             string `json:"state_topic,omitempty"`
	UniqueId               string `json:"unique_id,omitempty"`
	ValueTemplate          string `json:"value_template,omitempty"`
}

type HADeviceCamera struct {
	Availability struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	AvailabilityMode  string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	Device            struct {
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	EnabledByDefault       bool   `json:"enabled_by_default,omitempty"`
	Icon                   string `json:"icon,omitempty"`
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JsonAttributesTopic    string `json:"json_attributes_topic,omitempty"`
	Name                   string `json:"name,omitempty"`
	Topic                  string `json:"topic,omitempty"`
	UniqueId               string `json:"unique_id,omitempty"`
}

type HADeviceCover struct {
	Availability struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	AvailabilityMode  string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	CommandTopic      string `json:"command_topic,omitempty"`
	Device            struct {
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	DeviceClass            string `json:"device_class,omitempty"`
	EnabledByDefault       bool   `json:"enabled_by_default,omitempty"`
	Icon                   string `json:"icon,omitempty"`
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JsonAttributesTopic    string `json:"json_attributes_topic,omitempty"`
	Name                   string `json:"name,omitempty"`
	Optimistic             bool   `json:"optimistic,omitempty"`
	PayloadAvailable       string `json:"payload_available,omitempty"`
	PayloadClose           string `json:"payload_close,omitempty"`
	PayloadNotAvailable    string `json:"payload_not_available,omitempty"`
	PayloadOpen            string `json:"payload_open,omitempty"`
	PayloadStop            string `json:"payload_stop,omitempty"`
	PositionClosed         int    `json:"position_closed,omitempty"`
	PositionOpen           int    `json:"position_open,omitempty"`
	PositionTemplate       string `json:"position_template,omitempty"`
	PositionTopic          string `json:"position_topic,omitempty"`
	Qos                    int    `json:"qos,omitempty"`
	Retain                 bool   `json:"retain,omitempty"`
	SetPositionTemplate    string `json:"set_position_template,omitempty"`
	SetPositionTopic       string `json:"set_position_topic,omitempty"`
	StateClosed            string `json:"state_closed,omitempty"`
	StateClosing           string `json:"state_closing,omitempty"`
	StateOpen              string `json:"state_open,omitempty"`
	StateOpening           string `json:"state_opening,omitempty"`
	StateStopped           string `json:"state_stopped,omitempty"`
	StateTopic             string `json:"state_topic,omitempty"`
	TiltClosedValue        int    `json:"tilt_closed_value,omitempty"`
	TiltCommandTemplate    string `json:"tilt_command_template,omitempty"`
	TiltCommandTopic       string `json:"tilt_command_topic,omitempty"`
	TiltMax                int    `json:"tilt_max,omitempty"`
	TiltMin                int    `json:"tilt_min,omitempty"`
	TiltOpenedValue        int    `json:"tilt_opened_value,omitempty"`
	TiltOptimistic         bool   `json:"tilt_optimistic,omitempty"`
	TiltStatusTemplate     string `json:"tilt_status_template,omitempty"`
	TiltStatusTopic        string `json:"tilt_status_topic,omitempty"`
	UniqueId               string `json:"unique_id,omitempty"`
	ValueTemplate          string `json:"value_template,omitempty"`
}

type HADeviceDeviceTracker struct {
	Devices        []string `json:"devices,omitempty"`
	PayloadHome    string   `json:"payload_home,omitempty"`
	PayloadNotHome string   `json:"payload_not_home,omitempty"`
	Qos            int      `json:"qos,omitempty"`
	SourceType     string   `json:"source_type,omitempty"`
}

type HADeviceDeviceTrigger struct {
	AutomationType string `json:"automation_type,omitempty"`
	Device         struct {
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	Payload string `json:"payload,omitempty"`
	Qos     int    `json:"qos,omitempty"`
	Subtype string `json:"subtype,omitempty"`
	Topic   string `json:"topic,omitempty"`
	Type    string `json:"type,omitempty"`
}

type HADeviceFan struct {
	Availability struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	AvailabilityMode  string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	CommandTemplate   string `json:"command_template,omitempty"`
	CommandTopic      string `json:"command_topic,omitempty"`
	Device            struct {
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	EnabledByDefault           bool     `json:"enabled_by_default,omitempty"`
	Icon                       string   `json:"icon,omitempty"`
	JsonAttributesTemplate     string   `json:"json_attributes_template,omitempty"`
	JsonAttributesTopic        string   `json:"json_attributes_topic,omitempty"`
	Name                       string   `json:"name,omitempty"`
	Optimistic                 bool     `json:"optimistic,omitempty"`
	OscillationCommandTemplate string   `json:"oscillation_command_template,omitempty"`
	OscillationCommandTopic    string   `json:"oscillation_command_topic,omitempty"`
	OscillationStateTopic      string   `json:"oscillation_state_topic,omitempty"`
	OscillationValueTemplate   string   `json:"oscillation_value_template,omitempty"`
	PayloadAvailable           string   `json:"payload_available,omitempty"`
	PayloadNotAvailable        string   `json:"payload_not_available,omitempty"`
	PayloadOff                 string   `json:"payload_off,omitempty"`
	PayloadOn                  string   `json:"payload_on,omitempty"`
	PayloadOscillationOff      string   `json:"payload_oscillation_off,omitempty"`
	PayloadOscillationOn       string   `json:"payload_oscillation_on,omitempty"`
	PayloadResetPercentage     string   `json:"payload_reset_percentage,omitempty"`
	PayloadResetPresetMode     string   `json:"payload_reset_preset_mode,omitempty"`
	PercentageCommandTemplate  string   `json:"percentage_command_template,omitempty"`
	PercentageCommandTopic     string   `json:"percentage_command_topic,omitempty"`
	PercentageStateTopic       string   `json:"percentage_state_topic,omitempty"`
	PercentageValueTemplate    string   `json:"percentage_value_template,omitempty"`
	PresetModeCommandTemplate  string   `json:"preset_mode_command_template,omitempty"`
	PresetModeCommandTopic     string   `json:"preset_mode_command_topic,omitempty"`
	PresetModeStateTopic       string   `json:"preset_mode_state_topic,omitempty"`
	PresetModeValueTemplate    string   `json:"preset_mode_value_template,omitempty"`
	PresetModes                []string `json:"preset_modes,omitempty"`
	Qos                        int      `json:"qos,omitempty"`
	Retain                     bool     `json:"retain,omitempty"`
	SpeedRangeMax              int      `json:"speed_range_max,omitempty"`
	SpeedRangeMin              int      `json:"speed_range_min,omitempty"`
	StateTopic                 string   `json:"state_topic,omitempty"`
	StateValueTemplate         string   `json:"state_value_template,omitempty"`
	UniqueId                   string   `json:"unique_id,omitempty"`
}

type HADeviceHumidifier struct {
	Availability struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	AvailabilityMode  string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	CommandTemplate   string `json:"command_template,omitempty"`
	CommandTopic      string `json:"command_topic,omitempty"`
	Device            struct {
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	DeviceClass                   string   `json:"device_class,omitempty"`
	EnabledByDefault              bool     `json:"enabled_by_default,omitempty"`
	Icon                          string   `json:"icon,omitempty"`
	JsonAttributesTemplate        string   `json:"json_attributes_template,omitempty"`
	JsonAttributesTopic           string   `json:"json_attributes_topic,omitempty"`
	MaxHumidity                   int      `json:"max_humidity,omitempty"`
	MinHumidity                   int      `json:"min_humidity,omitempty"`
	ModeCommandTemplate           string   `json:"mode_command_template,omitempty"`
	ModeCommandTopic              string   `json:"mode_command_topic,omitempty"`
	ModeStateTemplate             string   `json:"mode_state_template,omitempty"`
	ModeStateTopic                string   `json:"mode_state_topic,omitempty"`
	Modes                         []string `json:"modes,omitempty"`
	Name                          string   `json:"name,omitempty"`
	Optimistic                    bool     `json:"optimistic,omitempty"`
	PayloadAvailable              string   `json:"payload_available,omitempty"`
	PayloadNotAvailable           string   `json:"payload_not_available,omitempty"`
	PayloadOff                    string   `json:"payload_off,omitempty"`
	PayloadOn                     string   `json:"payload_on,omitempty"`
	PayloadResetHumidity          string   `json:"payload_reset_humidity,omitempty"`
	PayloadResetMode              string   `json:"payload_reset_mode,omitempty"`
	Qos                           int      `json:"qos,omitempty"`
	Retain                        bool     `json:"retain,omitempty"`
	StateTopic                    string   `json:"state_topic,omitempty"`
	StateValueTemplate            string   `json:"state_value_template,omitempty"`
	TargetHumidityCommandTemplate string   `json:"target_humidity_command_template,omitempty"`
	TargetHumidityCommandTopic    string   `json:"target_humidity_command_topic,omitempty"`
	TargetHumidityStateTemplate   string   `json:"target_humidity_state_template,omitempty"`
	TargetHumidityStateTopic      string   `json:"target_humidity_state_topic,omitempty"`
	UniqueId                      string   `json:"unique_id,omitempty"`
}

type HADeviceClimate struct {
	ActionTemplate   string `json:"action_template,omitempty"`
	ActionTopic      string `json:"action_topic,omitempty"`
	AuxCommandTopic  string `json:"aux_command_topic,omitempty"`
	AuxStateTemplate string `json:"aux_state_template,omitempty"`
	AuxStateTopic    string `json:"aux_state_topic,omitempty"`
	Availability     struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	AvailabilityMode           string `json:"availability_mode,omitempty"`
	AvailabilityTopic          string `json:"availability_topic,omitempty"`
	AwayModeCommandTopic       string `json:"away_mode_command_topic,omitempty"`
	AwayModeStateTemplate      string `json:"away_mode_state_template,omitempty"`
	AwayModeStateTopic         string `json:"away_mode_state_topic,omitempty"`
	CurrentTemperatureTemplate string `json:"current_temperature_template,omitempty"`
	CurrentTemperatureTopic    string `json:"current_temperature_topic,omitempty"`
	Device                     struct {
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	EnabledByDefault               bool     `json:"enabled_by_default,omitempty"`
	FanModeCommandTemplate         string   `json:"fan_mode_command_template,omitempty"`
	FanModeCommandTopic            string   `json:"fan_mode_command_topic,omitempty"`
	FanModeStateTemplate           string   `json:"fan_mode_state_template,omitempty"`
	FanModeStateTopic              string   `json:"fan_mode_state_topic,omitempty"`
	FanModes                       []string `json:"fan_modes,omitempty"`
	HoldCommandTemplate            string   `json:"hold_command_template,omitempty"`
	HoldCommandTopic               string   `json:"hold_command_topic,omitempty"`
	HoldModes                      []string `json:"hold_modes,omitempty"`
	HoldStateTemplate              string   `json:"hold_state_template,omitempty"`
	HoldStateTopic                 string   `json:"hold_state_topic,omitempty"`
	Icon                           string   `json:"icon,omitempty"`
	Initial                        int      `json:"initial,omitempty"`
	JsonAttributesTemplate         string   `json:"json_attributes_template,omitempty"`
	JsonAttributesTopic            string   `json:"json_attributes_topic,omitempty"`
	MaxTemp                        float64  `json:"max_temp,omitempty"`
	MinTemp                        float64  `json:"min_temp,omitempty"`
	ModeCommandTemplate            string   `json:"mode_command_template,omitempty"`
	ModeCommandTopic               string   `json:"mode_command_topic,omitempty"`
	ModeStateTemplate              string   `json:"mode_state_template,omitempty"`
	ModeStateTopic                 string   `json:"mode_state_topic,omitempty"`
	Modes                          []string `json:"modes,omitempty"`
	Name                           string   `json:"name,omitempty"`
	PayloadAvailable               string   `json:"payload_available,omitempty"`
	PayloadNotAvailable            string   `json:"payload_not_available,omitempty"`
	PayloadOff                     string   `json:"payload_off,omitempty"`
	PayloadOn                      string   `json:"payload_on,omitempty"`
	PowerCommandTopic              string   `json:"power_command_topic,omitempty"`
	Precision                      float64  `json:"precision,omitempty"`
	Qos                            int      `json:"qos,omitempty"`
	Retain                         bool     `json:"retain,omitempty"`
	SendIfOff                      bool     `json:"send_if_off,omitempty"`
	SwingModeCommandTemplate       string   `json:"swing_mode_command_template,omitempty"`
	SwingModeCommandTopic          string   `json:"swing_mode_command_topic,omitempty"`
	SwingModeStateTemplate         string   `json:"swing_mode_state_template,omitempty"`
	SwingModeStateTopic            string   `json:"swing_mode_state_topic,omitempty"`
	SwingModes                     []string `json:"swing_modes,omitempty"`
	TempStep                       float64  `json:"temp_step,omitempty"`
	TemperatureCommandTemplate     string   `json:"temperature_command_template,omitempty"`
	TemperatureCommandTopic        string   `json:"temperature_command_topic,omitempty"`
	TemperatureHighCommandTemplate string   `json:"temperature_high_command_template,omitempty"`
	TemperatureHighCommandTopic    string   `json:"temperature_high_command_topic,omitempty"`
	TemperatureHighStateTemplate   string   `json:"temperature_high_state_template,omitempty"`
	TemperatureHighStateTopic      string   `json:"temperature_high_state_topic,omitempty"`
	TemperatureLowCommandTemplate  string   `json:"temperature_low_command_template,omitempty"`
	TemperatureLowCommandTopic     string   `json:"temperature_low_command_topic,omitempty"`
	TemperatureLowStateTemplate    string   `json:"temperature_low_state_template,omitempty"`
	TemperatureLowStateTopic       string   `json:"temperature_low_state_topic,omitempty"`
	TemperatureStateTemplate       string   `json:"temperature_state_template,omitempty"`
	TemperatureStateTopic          string   `json:"temperature_state_topic,omitempty"`
	TemperatureUnit                string   `json:"temperature_unit,omitempty"`
	UniqueId                       string   `json:"unique_id,omitempty"`
	ValueTemplate                  string   `json:"value_template,omitempty"`
}

type HADeviceLight struct {
	Availability struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	AvailabilityMode         string `json:"availability_mode,omitempty"`
	AvailabilityTopic        string `json:"availability_topic,omitempty"`
	BrightnessCommandTopic   string `json:"brightness_command_topic,omitempty"`
	BrightnessScale          int    `json:"brightness_scale,omitempty"`
	BrightnessStateTopic     string `json:"brightness_state_topic,omitempty"`
	BrightnessValueTemplate  string `json:"brightness_value_template,omitempty"`
	ColorModeStateTopic      string `json:"color_mode_state_topic,omitempty"`
	ColorModeValueTemplate   string `json:"color_mode_value_template,omitempty"`
	ColorTempCommandTemplate string `json:"color_temp_command_template,omitempty"`
	ColorTempCommandTopic    string `json:"color_temp_command_topic,omitempty"`
	ColorTempStateTopic      string `json:"color_temp_state_topic,omitempty"`
	ColorTempValueTemplate   string `json:"color_temp_value_template,omitempty"`
	CommandTopic             string `json:"command_topic,omitempty"`
	Device                   struct {
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	EffectCommandTopic     string   `json:"effect_command_topic,omitempty"`
	EffectList             []string `json:"effect_list,omitempty"`
	EffectStateTopic       string   `json:"effect_state_topic,omitempty"`
	EffectValueTemplate    string   `json:"effect_value_template,omitempty"`
	EnabledByDefault       bool     `json:"enabled_by_default,omitempty"`
	HsCommandTopic         string   `json:"hs_command_topic,omitempty"`
	HsStateTopic           string   `json:"hs_state_topic,omitempty"`
	HsValueTemplate        string   `json:"hs_value_template,omitempty"`
	Icon                   string   `json:"icon,omitempty"`
	JsonAttributesTemplate string   `json:"json_attributes_template,omitempty"`
	JsonAttributesTopic    string   `json:"json_attributes_topic,omitempty"`
	MaxMireds              int      `json:"max_mireds,omitempty"`
	MinMireds              int      `json:"min_mireds,omitempty"`
	Name                   string   `json:"name,omitempty"`
	OnCommandType          string   `json:"on_command_type,omitempty"`
	Optimistic             bool     `json:"optimistic,omitempty"`
	PayloadAvailable       string   `json:"payload_available,omitempty"`
	PayloadNotAvailable    string   `json:"payload_not_available,omitempty"`
	PayloadOff             string   `json:"payload_off,omitempty"`
	PayloadOn              string   `json:"payload_on,omitempty"`
	Qos                    int      `json:"qos,omitempty"`
	Retain                 bool     `json:"retain,omitempty"`
	RgbCommandTemplate     string   `json:"rgb_command_template,omitempty"`
	RgbCommandTopic        string   `json:"rgb_command_topic,omitempty"`
	RgbStateTopic          string   `json:"rgb_state_topic,omitempty"`
	RgbValueTemplate       string   `json:"rgb_value_template,omitempty"`
	Schema                 string   `json:"schema,omitempty"`
	StateTopic             string   `json:"state_topic,omitempty"`
	StateValueTemplate     string   `json:"state_value_template,omitempty"`
	UniqueId               string   `json:"unique_id,omitempty"`
	WhiteCommandTopic      string   `json:"white_command_topic,omitempty"`
	WhiteScale             int      `json:"white_scale,omitempty"`
	XyCommandTopic         string   `json:"xy_command_topic,omitempty"`
	XyStateTopic           string   `json:"xy_state_topic,omitempty"`
	XyValueTemplate        string   `json:"xy_value_template,omitempty"`
}

type HADeviceLock struct {
	Availability struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	AvailabilityMode  string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	CommandTopic      string `json:"command_topic,omitempty"`
	Device            struct {
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	EnabledByDefault       bool   `json:"enabled_by_default,omitempty"`
	Icon                   string `json:"icon,omitempty"`
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JsonAttributesTopic    string `json:"json_attributes_topic,omitempty"`
	Name                   string `json:"name,omitempty"`
	Optimistic             bool   `json:"optimistic,omitempty"`
	PayloadAvailable       string `json:"payload_available,omitempty"`
	PayloadLock            string `json:"payload_lock,omitempty"`
	PayloadNotAvailable    string `json:"payload_not_available,omitempty"`
	PayloadUnlock          string `json:"payload_unlock,omitempty"`
	Qos                    int    `json:"qos,omitempty"`
	Retain                 bool   `json:"retain,omitempty"`
	StateLocked            string `json:"state_locked,omitempty"`
	StateTopic             string `json:"state_topic,omitempty"`
	StateUnlocked          string `json:"state_unlocked,omitempty"`
	UniqueId               string `json:"unique_id,omitempty"`
	ValueTemplate          string `json:"value_template,omitempty"`
}

type HADeviceNumber struct {
	Availability struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	AvailabilityMode  string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	CommandTopic      string `json:"command_topic,omitempty"`
	Device            struct {
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	EnabledByDefault       bool    `json:"enabled_by_default,omitempty"`
	Icon                   string  `json:"icon,omitempty"`
	JsonAttributesTemplate string  `json:"json_attributes_template,omitempty"`
	JsonAttributesTopic    string  `json:"json_attributes_topic,omitempty"`
	Max                    float64 `json:"max,omitempty"`
	Min                    float64 `json:"min,omitempty"`
	Name                   string  `json:"name,omitempty"`
	Optimistic             bool    `json:"optimistic,omitempty"`
	Qos                    int     `json:"qos,omitempty"`
	Retain                 bool    `json:"retain,omitempty"`
	StateTopic             string  `json:"state_topic,omitempty"`
	Step                   float64 `json:"step,omitempty"`
	UniqueId               string  `json:"unique_id,omitempty"`
	ValueTemplate          string  `json:"value_template,omitempty"`
}

type HADeviceScene struct {
	Availability struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	AvailabilityMode    string `json:"availability_mode,omitempty"`
	AvailabilityTopic   string `json:"availability_topic,omitempty"`
	CommandTopic        string `json:"command_topic,omitempty"`
	EnabledByDefault    bool   `json:"enabled_by_default,omitempty"`
	Icon                string `json:"icon,omitempty"`
	Name                string `json:"name,omitempty"`
	PayloadAvailable    string `json:"payload_available,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	PayloadOn           string `json:"payload_on,omitempty"`
	Qos                 int    `json:"qos,omitempty"`
	Retain              bool   `json:"retain,omitempty"`
	UniqueId            string `json:"unique_id,omitempty"`
}

type HADeviceSelect struct {
	Availability struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	AvailabilityMode  string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	CommandTopic      string `json:"command_topic,omitempty"`
	Device            struct {
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	EnabledByDefault       bool     `json:"enabled_by_default,omitempty"`
	Icon                   string   `json:"icon,omitempty"`
	JsonAttributesTemplate string   `json:"json_attributes_template,omitempty"`
	JsonAttributesTopic    string   `json:"json_attributes_topic,omitempty"`
	Name                   string   `json:"name,omitempty"`
	Optimistic             bool     `json:"optimistic,omitempty"`
	Options                []string `json:"options,omitempty"`
	Qos                    int      `json:"qos,omitempty"`
	Retain                 bool     `json:"retain,omitempty"`
	StateTopic             string   `json:"state_topic,omitempty"`
	UniqueId               string   `json:"unique_id,omitempty"`
	ValueTemplate          string   `json:"value_template,omitempty"`
}

type HADeviceSensor struct {
	Availability struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	AvailabilityMode  string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	Device            struct {
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	DeviceClass            string `json:"device_class,omitempty"`
	EnabledByDefault       bool   `json:"enabled_by_default,omitempty"`
	ExpireAfter            int    `json:"expire_after,omitempty"`
	ForceUpdate            bool   `json:"force_update,omitempty"`
	Icon                   string `json:"icon,omitempty"`
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JsonAttributesTopic    string `json:"json_attributes_topic,omitempty"`
	LastResetTopic         string `json:"last_reset_topic,omitempty"`
	LastResetValueTemplate string `json:"last_reset_value_template,omitempty"`
	Name                   string `json:"name,omitempty"`
	PayloadAvailable       string `json:"payload_available,omitempty"`
	PayloadNotAvailable    string `json:"payload_not_available,omitempty"`
	Qos                    int    `json:"qos,omitempty"`
	StateClass             string `json:"state_class,omitempty"`
	StateTopic             string `json:"state_topic,omitempty"`
	UniqueId               string `json:"unique_id,omitempty"`
	UnitOfMeasurement      string `json:"unit_of_measurement,omitempty"`
	ValueTemplate          string `json:"value_template,omitempty"`
}

type HADeviceSwitch struct {
	Availability struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	AvailabilityMode  string `json:"availability_mode,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	CommandTopic      string `json:"command_topic,omitempty"`
	Device            struct {
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	EnabledByDefault       bool   `json:"enabled_by_default,omitempty"`
	Icon                   string `json:"icon,omitempty"`
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	JsonAttributesTopic    string `json:"json_attributes_topic,omitempty"`
	Name                   string `json:"name,omitempty"`
	Optimistic             bool   `json:"optimistic,omitempty"`
	PayloadAvailable       string `json:"payload_available,omitempty"`
	PayloadNotAvailable    string `json:"payload_not_available,omitempty"`
	PayloadOff             string `json:"payload_off,omitempty"`
	PayloadOn              string `json:"payload_on,omitempty"`
	Qos                    int    `json:"qos,omitempty"`
	Retain                 bool   `json:"retain,omitempty"`
	StateOff               string `json:"state_off,omitempty"`
	StateOn                string `json:"state_on,omitempty"`
	StateTopic             string `json:"state_topic,omitempty"`
	UniqueId               string `json:"unique_id,omitempty"`
	ValueTemplate          string `json:"value_template,omitempty"`
}

type HADeviceTag struct {
	Device struct {
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	Topic         string `json:"topic,omitempty"`
	ValueTemplate string `json:"value_template,omitempty"`
}

type HADeviceVacuum struct {
	Availability struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	AvailabilityMode       string   `json:"availability_mode,omitempty"`
	AvailabilityTopic      string   `json:"availability_topic,omitempty"`
	BatteryLevelTemplate   string   `json:"battery_level_template,omitempty"`
	BatteryLevelTopic      string   `json:"battery_level_topic,omitempty"`
	ChargingTemplate       string   `json:"charging_template,omitempty"`
	ChargingTopic          string   `json:"charging_topic,omitempty"`
	CleaningTemplate       string   `json:"cleaning_template,omitempty"`
	CleaningTopic          string   `json:"cleaning_topic,omitempty"`
	CommandTopic           string   `json:"command_topic,omitempty"`
	DockedTemplate         string   `json:"docked_template,omitempty"`
	DockedTopic            string   `json:"docked_topic,omitempty"`
	EnabledByDefault       bool     `json:"enabled_by_default,omitempty"`
	ErrorTemplate          string   `json:"error_template,omitempty"`
	ErrorTopic             string   `json:"error_topic,omitempty"`
	FanSpeedList           []string `json:"fan_speed_list,omitempty"`
	FanSpeedTemplate       string   `json:"fan_speed_template,omitempty"`
	FanSpeedTopic          string   `json:"fan_speed_topic,omitempty"`
	Icon                   string   `json:"icon,omitempty"`
	JsonAttributesTemplate string   `json:"json_attributes_template,omitempty"`
	JsonAttributesTopic    string   `json:"json_attributes_topic,omitempty"`
	Name                   string   `json:"name,omitempty"`
	PayloadAvailable       string   `json:"payload_available,omitempty"`
	PayloadCleanSpot       string   `json:"payload_clean_spot,omitempty"`
	PayloadLocate          string   `json:"payload_locate,omitempty"`
	PayloadNotAvailable    string   `json:"payload_not_available,omitempty"`
	PayloadReturnToBase    string   `json:"payload_return_to_base,omitempty"`
	PayloadStartPause      string   `json:"payload_start_pause,omitempty"`
	PayloadStop            string   `json:"payload_stop,omitempty"`
	PayloadTurnOff         string   `json:"payload_turn_off,omitempty"`
	PayloadTurnOn          string   `json:"payload_turn_on,omitempty"`
	Qos                    int      `json:"qos,omitempty"`
	Retain                 bool     `json:"retain,omitempty"`
	Schema                 string   `json:"schema,omitempty"`
	SendCommandTopic       string   `json:"send_command_topic,omitempty"`
	SetFanSpeedTopic       string   `json:"set_fan_speed_topic,omitempty"`
	SupportedFeatures      []string `json:"supported_features,omitempty"`
	UniqueId               string   `json:"unique_id,omitempty"`
}
