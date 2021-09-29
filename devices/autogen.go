package devices

type HADeviceAlarmControlPanel struct {
	PayloadArmHome         string `json:"payload_arm_home,omitempty"`
	CommandTemplate        string `json:"command_template,omitempty"`
	PayloadArmVacation     string `json:"payload_arm_vacation,omitempty"`
	CodeDisarmRequired     bool   `json:"code_disarm_required,omitempty"`
	PayloadArmCustomBypass string `json:"payload_arm_custom_bypass,omitempty"`
	StateTopic             string `json:"state_topic,omitempty"`
	Device                 struct {
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
	} `json:"device,omitempty"`
	AvailabilityTopic      string `json:"availability_topic,omitempty"`
	EnabledByDefault       bool   `json:"enabled_by_default,omitempty"`
	PayloadNotAvailable    string `json:"payload_not_available,omitempty"`
	Code                   string `json:"code,omitempty"`
	UniqueId               string `json:"unique_id,omitempty"`
	ValueTemplate          string `json:"value_template,omitempty"`
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	CommandTopic           string `json:"command_topic,omitempty"`
	PayloadDisarm          string `json:"payload_disarm,omitempty"`
	Name                   string `json:"name,omitempty"`
	CodeArmRequired        bool   `json:"code_arm_required,omitempty"`
	PayloadArmAway         string `json:"payload_arm_away,omitempty"`
	PayloadAvailable       string `json:"payload_available,omitempty"`
	Availability           struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	AvailabilityMode    string `json:"availability_mode,omitempty"`
	Retain              bool   `json:"retain,omitempty"`
	Icon                string `json:"icon,omitempty"`
	Qos                 int    `json:"qos,omitempty"`
	PayloadArmNight     string `json:"payload_arm_night,omitempty"`
}

type HADeviceBinarySensor struct {
	PayloadOn              string `json:"payload_on,omitempty"`
	ExpireAfter            int    `json:"expire_after,omitempty"`
	StateTopic             string `json:"state_topic,omitempty"`
	PayloadAvailable       string `json:"payload_available,omitempty"`
	PayloadOff             string `json:"payload_off,omitempty"`
	Qos                    int    `json:"qos,omitempty"`
	DeviceClass            string `json:"device_class,omitempty"`
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	Availability           struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	AvailabilityMode    string `json:"availability_mode,omitempty"`
	ValueTemplate       string `json:"value_template,omitempty"`
	Icon                string `json:"icon,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	UniqueId            string `json:"unique_id,omitempty"`
	Name                string `json:"name,omitempty"`
	EnabledByDefault    bool   `json:"enabled_by_default,omitempty"`
	OffDelay            int    `json:"off_delay,omitempty"`
	Device              struct {
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
		Connections   []string `json:"connections,omitempty"`
	} `json:"device,omitempty"`
	ForceUpdate         bool   `json:"force_update,omitempty"`
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	AvailabilityTopic   string `json:"availability_topic,omitempty"`
}

type HADeviceCamera struct {
	AvailabilityTopic      string `json:"availability_topic,omitempty"`
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	Topic                  string `json:"topic,omitempty"`
	EnabledByDefault       bool   `json:"enabled_by_default,omitempty"`
	Name                   string `json:"name,omitempty"`
	Icon                   string `json:"icon,omitempty"`
	Availability           struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	Device struct {
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
	} `json:"device,omitempty"`
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	UniqueId            string `json:"unique_id,omitempty"`
	AvailabilityMode    string `json:"availability_mode,omitempty"`
}

type HADeviceCover struct {
	ValueTemplate    string `json:"value_template,omitempty"`
	PositionOpen     int    `json:"position_open,omitempty"`
	SetPositionTopic string `json:"set_position_topic,omitempty"`
	Availability     struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	TiltMax                int    `json:"tilt_max,omitempty"`
	JsonAttributesTopic    string `json:"json_attributes_topic,omitempty"`
	Qos                    int    `json:"qos,omitempty"`
	PositionClosed         int    `json:"position_closed,omitempty"`
	PositionTemplate       string `json:"position_template,omitempty"`
	TiltOptimistic         bool   `json:"tilt_optimistic,omitempty"`
	TiltCommandTemplate    string `json:"tilt_command_template,omitempty"`
	PositionTopic          string `json:"position_topic,omitempty"`
	PayloadClose           string `json:"payload_close,omitempty"`
	UniqueId               string `json:"unique_id,omitempty"`
	Name                   string `json:"name,omitempty"`
	TiltCommandTopic       string `json:"tilt_command_topic,omitempty"`
	TiltStatusTemplate     string `json:"tilt_status_template,omitempty"`
	TiltMin                int    `json:"tilt_min,omitempty"`
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	PayloadOpen            string `json:"payload_open,omitempty"`
	EnabledByDefault       bool   `json:"enabled_by_default,omitempty"`
	SetPositionTemplate    string `json:"set_position_template,omitempty"`
	AvailabilityMode       string `json:"availability_mode,omitempty"`
	PayloadNotAvailable    string `json:"payload_not_available,omitempty"`
	Optimistic             bool   `json:"optimistic,omitempty"`
	DeviceClass            string `json:"device_class,omitempty"`
	TiltClosedValue        int    `json:"tilt_closed_value,omitempty"`
	CommandTopic           string `json:"command_topic,omitempty"`
	StateClosing           string `json:"state_closing,omitempty"`
	PayloadStop            string `json:"payload_stop,omitempty"`
	StateClosed            string `json:"state_closed,omitempty"`
	Device                 struct {
		ViaDevice     string   `json:"via_device,omitempty"`
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
	} `json:"device,omitempty"`
	StateOpen         string `json:"state_open,omitempty"`
	TiltOpenedValue   int    `json:"tilt_opened_value,omitempty"`
	Retain            bool   `json:"retain,omitempty"`
	StateTopic        string `json:"state_topic,omitempty"`
	Icon              string `json:"icon,omitempty"`
	StateOpening      string `json:"state_opening,omitempty"`
	PayloadAvailable  string `json:"payload_available,omitempty"`
	StateStopped      string `json:"state_stopped,omitempty"`
	TiltStatusTopic   string `json:"tilt_status_topic,omitempty"`
	AvailabilityTopic string `json:"availability_topic,omitempty"`
}

type HADeviceDeviceTracker struct {
	Devices        []string `json:"devices,omitempty"`
	PayloadHome    string   `json:"payload_home,omitempty"`
	PayloadNotHome string   `json:"payload_not_home,omitempty"`
	Qos            int      `json:"qos,omitempty"`
	SourceType     string   `json:"source_type,omitempty"`
}

type HADeviceDeviceTrigger struct {
	Subtype        string `json:"subtype,omitempty"`
	Topic          string `json:"topic,omitempty"`
	Type           string `json:"type,omitempty"`
	AutomationType string `json:"automation_type,omitempty"`
	Device         struct {
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
	} `json:"device,omitempty"`
	Payload string `json:"payload,omitempty"`
	Qos     int    `json:"qos,omitempty"`
}

type HADeviceFan struct {
	PayloadResetPercentage string `json:"payload_reset_percentage,omitempty"`
	Availability           struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	Icon                       string `json:"icon,omitempty"`
	CommandTopic               string `json:"command_topic,omitempty"`
	PayloadResetPresetMode     string `json:"payload_reset_preset_mode,omitempty"`
	Retain                     bool   `json:"retain,omitempty"`
	StateTopic                 string `json:"state_topic,omitempty"`
	PayloadNotAvailable        string `json:"payload_not_available,omitempty"`
	PayloadOff                 string `json:"payload_off,omitempty"`
	OscillationCommandTemplate string `json:"oscillation_command_template,omitempty"`
	CommandTemplate            string `json:"command_template,omitempty"`
	PayloadOscillationOn       string `json:"payload_oscillation_on,omitempty"`
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
	PresetModeStateTopic      string   `json:"preset_mode_state_topic,omitempty"`
	UniqueId                  string   `json:"unique_id,omitempty"`
	OscillationCommandTopic   string   `json:"oscillation_command_topic,omitempty"`
	PercentageCommandTopic    string   `json:"percentage_command_topic,omitempty"`
	PresetModeValueTemplate   string   `json:"preset_mode_value_template,omitempty"`
	PresetModeCommandTopic    string   `json:"preset_mode_command_topic,omitempty"`
	EnabledByDefault          bool     `json:"enabled_by_default,omitempty"`
	OscillationValueTemplate  string   `json:"oscillation_value_template,omitempty"`
	PercentageStateTopic      string   `json:"percentage_state_topic,omitempty"`
	SpeedRangeMin             int      `json:"speed_range_min,omitempty"`
	PercentageValueTemplate   string   `json:"percentage_value_template,omitempty"`
	PayloadOscillationOff     string   `json:"payload_oscillation_off,omitempty"`
	Name                      string   `json:"name,omitempty"`
	PayloadAvailable          string   `json:"payload_available,omitempty"`
	AvailabilityTopic         string   `json:"availability_topic,omitempty"`
	SpeedRangeMax             int      `json:"speed_range_max,omitempty"`
	PayloadOn                 string   `json:"payload_on,omitempty"`
	Qos                       int      `json:"qos,omitempty"`
	PresetModes               []string `json:"preset_modes,omitempty"`
	JsonAttributesTopic       string   `json:"json_attributes_topic,omitempty"`
	JsonAttributesTemplate    string   `json:"json_attributes_template,omitempty"`
	OscillationStateTopic     string   `json:"oscillation_state_topic,omitempty"`
	AvailabilityMode          string   `json:"availability_mode,omitempty"`
	PresetModeCommandTemplate string   `json:"preset_mode_command_template,omitempty"`
	PercentageCommandTemplate string   `json:"percentage_command_template,omitempty"`
	StateValueTemplate        string   `json:"state_value_template,omitempty"`
	Optimistic                bool     `json:"optimistic,omitempty"`
}

type HADeviceHumidifier struct {
	StateValueTemplate string `json:"state_value_template,omitempty"`
	Device             struct {
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
		Connections   []string `json:"connections,omitempty"`
	} `json:"device,omitempty"`
	ModeCommandTopic       string   `json:"mode_command_topic,omitempty"`
	MaxHumidity            int      `json:"max_humidity,omitempty"`
	MinHumidity            int      `json:"min_humidity,omitempty"`
	JsonAttributesTopic    string   `json:"json_attributes_topic,omitempty"`
	EnabledByDefault       bool     `json:"enabled_by_default,omitempty"`
	Name                   string   `json:"name,omitempty"`
	AvailabilityMode       string   `json:"availability_mode,omitempty"`
	Modes                  []string `json:"modes,omitempty"`
	PayloadResetHumidity   string   `json:"payload_reset_humidity,omitempty"`
	CommandTemplate        string   `json:"command_template,omitempty"`
	PayloadAvailable       string   `json:"payload_available,omitempty"`
	PayloadOn              string   `json:"payload_on,omitempty"`
	ModeStateTopic         string   `json:"mode_state_topic,omitempty"`
	StateTopic             string   `json:"state_topic,omitempty"`
	CommandTopic           string   `json:"command_topic,omitempty"`
	DeviceClass            string   `json:"device_class,omitempty"`
	JsonAttributesTemplate string   `json:"json_attributes_template,omitempty"`
	Availability           struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	TargetHumidityCommandTemplate string `json:"target_humidity_command_template,omitempty"`
	UniqueId                      string `json:"unique_id,omitempty"`
	Icon                          string `json:"icon,omitempty"`
	TargetHumidityStateTopic      string `json:"target_humidity_state_topic,omitempty"`
	Retain                        bool   `json:"retain,omitempty"`
	Optimistic                    bool   `json:"optimistic,omitempty"`
	Qos                           int    `json:"qos,omitempty"`
	PayloadResetMode              string `json:"payload_reset_mode,omitempty"`
	ModeCommandTemplate           string `json:"mode_command_template,omitempty"`
	AvailabilityTopic             string `json:"availability_topic,omitempty"`
	PayloadNotAvailable           string `json:"payload_not_available,omitempty"`
	TargetHumidityCommandTopic    string `json:"target_humidity_command_topic,omitempty"`
	TargetHumidityStateTemplate   string `json:"target_humidity_state_template,omitempty"`
	ModeStateTemplate             string `json:"mode_state_template,omitempty"`
	PayloadOff                    string `json:"payload_off,omitempty"`
}

type HADeviceClimate struct {
	MinTemp                        float64  `json:"min_temp,omitempty"`
	EnabledByDefault               bool     `json:"enabled_by_default,omitempty"`
	TemperatureLowStateTopic       string   `json:"temperature_low_state_topic,omitempty"`
	CurrentTemperatureTopic        string   `json:"current_temperature_topic,omitempty"`
	SwingModes                     []string `json:"swing_modes,omitempty"`
	SwingModeCommandTemplate       string   `json:"swing_mode_command_template,omitempty"`
	AwayModeCommandTopic           string   `json:"away_mode_command_topic,omitempty"`
	FanModeCommandTemplate         string   `json:"fan_mode_command_template,omitempty"`
	AuxCommandTopic                string   `json:"aux_command_topic,omitempty"`
	AuxStateTemplate               string   `json:"aux_state_template,omitempty"`
	TemperatureHighCommandTemplate string   `json:"temperature_high_command_template,omitempty"`
	ModeCommandTemplate            string   `json:"mode_command_template,omitempty"`
	UniqueId                       string   `json:"unique_id,omitempty"`
	SwingModeStateTemplate         string   `json:"swing_mode_state_template,omitempty"`
	PayloadOff                     string   `json:"payload_off,omitempty"`
	FanModeStateTemplate           string   `json:"fan_mode_state_template,omitempty"`
	HoldModes                      []string `json:"hold_modes,omitempty"`
	HoldStateTemplate              string   `json:"hold_state_template,omitempty"`
	Modes                          []string `json:"modes,omitempty"`
	JsonAttributesTemplate         string   `json:"json_attributes_template,omitempty"`
	AvailabilityMode               string   `json:"availability_mode,omitempty"`
	Qos                            int      `json:"qos,omitempty"`
	AuxStateTopic                  string   `json:"aux_state_topic,omitempty"`
	SwingModeStateTopic            string   `json:"swing_mode_state_topic,omitempty"`
	Retain                         bool     `json:"retain,omitempty"`
	TemperatureHighStateTopic      string   `json:"temperature_high_state_topic,omitempty"`
	Name                           string   `json:"name,omitempty"`
	HoldCommandTopic               string   `json:"hold_command_topic,omitempty"`
	TemperatureCommandTopic        string   `json:"temperature_command_topic,omitempty"`
	TemperatureCommandTemplate     string   `json:"temperature_command_template,omitempty"`
	Initial                        int      `json:"initial,omitempty"`
	TemperatureStateTopic          string   `json:"temperature_state_topic,omitempty"`
	Device                         struct {
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
		Connections   []string `json:"connections,omitempty"`
	} `json:"device,omitempty"`
	ActionTemplate               string   `json:"action_template,omitempty"`
	CurrentTemperatureTemplate   string   `json:"current_temperature_template,omitempty"`
	ActionTopic                  string   `json:"action_topic,omitempty"`
	TemperatureHighStateTemplate string   `json:"temperature_high_state_template,omitempty"`
	PayloadAvailable             string   `json:"payload_available,omitempty"`
	Icon                         string   `json:"icon,omitempty"`
	AwayModeStateTopic           string   `json:"away_mode_state_topic,omitempty"`
	JsonAttributesTopic          string   `json:"json_attributes_topic,omitempty"`
	TemperatureHighCommandTopic  string   `json:"temperature_high_command_topic,omitempty"`
	SendIfOff                    bool     `json:"send_if_off,omitempty"`
	SwingModeCommandTopic        string   `json:"swing_mode_command_topic,omitempty"`
	ModeStateTopic               string   `json:"mode_state_topic,omitempty"`
	PowerCommandTopic            string   `json:"power_command_topic,omitempty"`
	TemperatureUnit              string   `json:"temperature_unit,omitempty"`
	FanModes                     []string `json:"fan_modes,omitempty"`
	PayloadOn                    string   `json:"payload_on,omitempty"`
	ModeCommandTopic             string   `json:"mode_command_topic,omitempty"`
	FanModeCommandTopic          string   `json:"fan_mode_command_topic,omitempty"`
	PayloadNotAvailable          string   `json:"payload_not_available,omitempty"`
	MaxTemp                      float64  `json:"max_temp,omitempty"`
	HoldStateTopic               string   `json:"hold_state_topic,omitempty"`
	Precision                    float64  `json:"precision,omitempty"`
	HoldCommandTemplate          string   `json:"hold_command_template,omitempty"`
	TemperatureLowCommandTopic   string   `json:"temperature_low_command_topic,omitempty"`
	AvailabilityTopic            string   `json:"availability_topic,omitempty"`
	AwayModeStateTemplate        string   `json:"away_mode_state_template,omitempty"`
	TemperatureStateTemplate     string   `json:"temperature_state_template,omitempty"`
	TemperatureLowStateTemplate  string   `json:"temperature_low_state_template,omitempty"`
	Availability                 struct {
		Topic               string `json:"topic,omitempty"`
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	} `json:"availability,omitempty"`
	ValueTemplate                 string  `json:"value_template,omitempty"`
	FanModeStateTopic             string  `json:"fan_mode_state_topic,omitempty"`
	TemperatureLowCommandTemplate string  `json:"temperature_low_command_template,omitempty"`
	ModeStateTemplate             string  `json:"mode_state_template,omitempty"`
	TempStep                      float64 `json:"temp_step,omitempty"`
}

type HADeviceLight struct {
	AvailabilityMode       string   `json:"availability_mode,omitempty"`
	Optimistic             bool     `json:"optimistic,omitempty"`
	Qos                    int      `json:"qos,omitempty"`
	BrightnessScale        int      `json:"brightness_scale,omitempty"`
	EnabledByDefault       bool     `json:"enabled_by_default,omitempty"`
	Schema                 string   `json:"schema,omitempty"`
	BrightnessStateTopic   string   `json:"brightness_state_topic,omitempty"`
	EffectList             []string `json:"effect_list,omitempty"`
	RgbCommandTemplate     string   `json:"rgb_command_template,omitempty"`
	PayloadNotAvailable    string   `json:"payload_not_available,omitempty"`
	RgbValueTemplate       string   `json:"rgb_value_template,omitempty"`
	HsStateTopic           string   `json:"hs_state_topic,omitempty"`
	Retain                 bool     `json:"retain,omitempty"`
	ColorTempValueTemplate string   `json:"color_temp_value_template,omitempty"`
	WhiteCommandTopic      string   `json:"white_command_topic,omitempty"`
	ColorModeStateTopic    string   `json:"color_mode_state_topic,omitempty"`
	StateValueTemplate     string   `json:"state_value_template,omitempty"`
	ColorTempStateTopic    string   `json:"color_temp_state_topic,omitempty"`
	Availability           struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	Icon                   string `json:"icon,omitempty"`
	CommandTopic           string `json:"command_topic,omitempty"`
	MinMireds              int    `json:"min_mireds,omitempty"`
	UniqueId               string `json:"unique_id,omitempty"`
	WhiteScale             int    `json:"white_scale,omitempty"`
	HsValueTemplate        string `json:"hs_value_template,omitempty"`
	PayloadOn              string `json:"payload_on,omitempty"`
	PayloadAvailable       string `json:"payload_available,omitempty"`
	XyStateTopic           string `json:"xy_state_topic,omitempty"`
	ColorModeValueTemplate string `json:"color_mode_value_template,omitempty"`
	PayloadOff             string `json:"payload_off,omitempty"`
	RgbStateTopic          string `json:"rgb_state_topic,omitempty"`
	MaxMireds              int    `json:"max_mireds,omitempty"`
	BrightnessCommandTopic string `json:"brightness_command_topic,omitempty"`
	Device                 struct {
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
	} `json:"device,omitempty"`
	EffectValueTemplate      string `json:"effect_value_template,omitempty"`
	JsonAttributesTopic      string `json:"json_attributes_topic,omitempty"`
	XyValueTemplate          string `json:"xy_value_template,omitempty"`
	ColorTempCommandTemplate string `json:"color_temp_command_template,omitempty"`
	JsonAttributesTemplate   string `json:"json_attributes_template,omitempty"`
	RgbCommandTopic          string `json:"rgb_command_topic,omitempty"`
	Name                     string `json:"name,omitempty"`
	ColorTempCommandTopic    string `json:"color_temp_command_topic,omitempty"`
	HsCommandTopic           string `json:"hs_command_topic,omitempty"`
	XyCommandTopic           string `json:"xy_command_topic,omitempty"`
	AvailabilityTopic        string `json:"availability_topic,omitempty"`
	OnCommandType            string `json:"on_command_type,omitempty"`
	BrightnessValueTemplate  string `json:"brightness_value_template,omitempty"`
	EffectCommandTopic       string `json:"effect_command_topic,omitempty"`
	StateTopic               string `json:"state_topic,omitempty"`
	EffectStateTopic         string `json:"effect_state_topic,omitempty"`
}

type HADeviceLock struct {
	Availability struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	StateTopic          string `json:"state_topic,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	UniqueId            string `json:"unique_id,omitempty"`
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	PayloadLock         string `json:"payload_lock,omitempty"`
	Retain              bool   `json:"retain,omitempty"`
	AvailabilityTopic   string `json:"availability_topic,omitempty"`
	Icon                string `json:"icon,omitempty"`
	ValueTemplate       string `json:"value_template,omitempty"`
	Device              struct {
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
	} `json:"device,omitempty"`
	Optimistic             bool   `json:"optimistic,omitempty"`
	PayloadAvailable       string `json:"payload_available,omitempty"`
	StateUnlocked          string `json:"state_unlocked,omitempty"`
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	Name                   string `json:"name,omitempty"`
	Qos                    int    `json:"qos,omitempty"`
	AvailabilityMode       string `json:"availability_mode,omitempty"`
	PayloadUnlock          string `json:"payload_unlock,omitempty"`
	CommandTopic           string `json:"command_topic,omitempty"`
	EnabledByDefault       bool   `json:"enabled_by_default,omitempty"`
	StateLocked            string `json:"state_locked,omitempty"`
}

type HADeviceNumber struct {
	EnabledByDefault  bool    `json:"enabled_by_default,omitempty"`
	Icon              string  `json:"icon,omitempty"`
	Optimistic        bool    `json:"optimistic,omitempty"`
	Max               float64 `json:"max,omitempty"`
	AvailabilityTopic string  `json:"availability_topic,omitempty"`
	Device            struct {
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
	} `json:"device,omitempty"`
	CommandTopic           string `json:"command_topic,omitempty"`
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	Name                   string `json:"name,omitempty"`
	Availability           struct {
		Topic               string `json:"topic,omitempty"`
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	} `json:"availability,omitempty"`
	StateTopic          string  `json:"state_topic,omitempty"`
	JsonAttributesTopic string  `json:"json_attributes_topic,omitempty"`
	Retain              bool    `json:"retain,omitempty"`
	UniqueId            string  `json:"unique_id,omitempty"`
	Step                float64 `json:"step,omitempty"`
	ValueTemplate       string  `json:"value_template,omitempty"`
	AvailabilityMode    string  `json:"availability_mode,omitempty"`
	Qos                 int     `json:"qos,omitempty"`
	Min                 float64 `json:"min,omitempty"`
}

type HADeviceScene struct {
	Qos          int `json:"qos,omitempty"`
	Availability struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	CommandTopic        string `json:"command_topic,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	AvailabilityTopic   string `json:"availability_topic,omitempty"`
	Icon                string `json:"icon,omitempty"`
	Name                string `json:"name,omitempty"`
	PayloadOn           string `json:"payload_on,omitempty"`
	AvailabilityMode    string `json:"availability_mode,omitempty"`
	EnabledByDefault    bool   `json:"enabled_by_default,omitempty"`
	PayloadAvailable    string `json:"payload_available,omitempty"`
	UniqueId            string `json:"unique_id,omitempty"`
	Retain              bool   `json:"retain,omitempty"`
}

type HADeviceSelect struct {
	UniqueId               string `json:"unique_id,omitempty"`
	ValueTemplate          string `json:"value_template,omitempty"`
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	AvailabilityMode       string `json:"availability_mode,omitempty"`
	Optimistic             bool   `json:"optimistic,omitempty"`
	CommandTopic           string `json:"command_topic,omitempty"`
	Availability           struct {
		Topic               string `json:"topic,omitempty"`
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	} `json:"availability,omitempty"`
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	StateTopic          string `json:"state_topic,omitempty"`
	EnabledByDefault    bool   `json:"enabled_by_default,omitempty"`
	AvailabilityTopic   string `json:"availability_topic,omitempty"`
	Device              struct {
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
	} `json:"device,omitempty"`
	Icon    string   `json:"icon,omitempty"`
	Options []string `json:"options,omitempty"`
	Retain  bool     `json:"retain,omitempty"`
	Name    string   `json:"name,omitempty"`
	Qos     int      `json:"qos,omitempty"`
}

type HADeviceSensor struct {
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	LastResetValueTemplate string `json:"last_reset_value_template,omitempty"`
	EnabledByDefault       bool   `json:"enabled_by_default,omitempty"`
	ValueTemplate          string `json:"value_template,omitempty"`
	AvailabilityMode       string `json:"availability_mode,omitempty"`
	UnitOfMeasurement      string `json:"unit_of_measurement,omitempty"`
	StateClass             string `json:"state_class,omitempty"`
	Icon                   string `json:"icon,omitempty"`
	UniqueId               string `json:"unique_id,omitempty"`
	LastResetTopic         string `json:"last_reset_topic,omitempty"`
	Availability           struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	DeviceClass string `json:"device_class,omitempty"`
	StateTopic  string `json:"state_topic,omitempty"`
	ForceUpdate bool   `json:"force_update,omitempty"`
	Qos         int    `json:"qos,omitempty"`
	Device      struct {
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
	} `json:"device,omitempty"`
	Name                string `json:"name,omitempty"`
	ExpireAfter         int    `json:"expire_after,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	AvailabilityTopic   string `json:"availability_topic,omitempty"`
	PayloadAvailable    string `json:"payload_available,omitempty"`
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
}

type HADeviceSwitch struct {
	ValueTemplate       string `json:"value_template,omitempty"`
	EnabledByDefault    bool   `json:"enabled_by_default,omitempty"`
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	PayloadOn           string `json:"payload_on,omitempty"`
	AvailabilityTopic   string `json:"availability_topic,omitempty"`
	Name                string `json:"name,omitempty"`
	PayloadAvailable    string `json:"payload_available,omitempty"`
	StateTopic          string `json:"state_topic,omitempty"`
	Device              struct {
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
	} `json:"device,omitempty"`
	JsonAttributesTopic    string `json:"json_attributes_topic,omitempty"`
	UniqueId               string `json:"unique_id,omitempty"`
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	Qos                    int    `json:"qos,omitempty"`
	StateOff               string `json:"state_off,omitempty"`
	CommandTopic           string `json:"command_topic,omitempty"`
	Availability           struct {
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		Topic               string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	AvailabilityMode string `json:"availability_mode,omitempty"`
	Icon             string `json:"icon,omitempty"`
	PayloadOff       string `json:"payload_off,omitempty"`
	Retain           bool   `json:"retain,omitempty"`
	StateOn          string `json:"state_on,omitempty"`
	Optimistic       bool   `json:"optimistic,omitempty"`
}

type HADeviceTag struct {
	ValueTemplate string `json:"value_template,omitempty"`
	Device        struct {
		SuggestedArea string   `json:"suggested_area,omitempty"`
		SwVersion     string   `json:"sw_version,omitempty"`
		ViaDevice     string   `json:"via_device,omitempty"`
		Connections   []string `json:"connections,omitempty"`
		Identifiers   []string `json:"identifiers,omitempty"`
		Manufacturer  string   `json:"manufacturer,omitempty"`
		Model         string   `json:"model,omitempty"`
		Name          string   `json:"name,omitempty"`
	} `json:"device,omitempty"`
	Topic string `json:"topic,omitempty"`
}

type HADeviceVacuum struct {
	ChargingTemplate     string   `json:"charging_template,omitempty"`
	FanSpeedList         []string `json:"fan_speed_list,omitempty"`
	BatteryLevelTemplate string   `json:"battery_level_template,omitempty"`
	PayloadStartPause    string   `json:"payload_start_pause,omitempty"`
	BatteryLevelTopic    string   `json:"battery_level_topic,omitempty"`
	DockedTemplate       string   `json:"docked_template,omitempty"`
	ErrorTopic           string   `json:"error_topic,omitempty"`
	AvailabilityTopic    string   `json:"availability_topic,omitempty"`
	AvailabilityMode     string   `json:"availability_mode,omitempty"`
	SupportedFeatures    []string `json:"supported_features,omitempty"`
	PayloadAvailable     string   `json:"payload_available,omitempty"`
	Retain               bool     `json:"retain,omitempty"`
	Schema               string   `json:"schema,omitempty"`
	PayloadCleanSpot     string   `json:"payload_clean_spot,omitempty"`
	Qos                  int      `json:"qos,omitempty"`
	SetFanSpeedTopic     string   `json:"set_fan_speed_topic,omitempty"`
	Availability         struct {
		Topic               string `json:"topic,omitempty"`
		PayloadAvailable    string `json:"payload_available,omitempty"`
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	} `json:"availability,omitempty"`
	DockedTopic            string `json:"docked_topic,omitempty"`
	Icon                   string `json:"icon,omitempty"`
	SendCommandTopic       string `json:"send_command_topic,omitempty"`
	UniqueId               string `json:"unique_id,omitempty"`
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	CommandTopic           string `json:"command_topic,omitempty"`
	PayloadTurnOff         string `json:"payload_turn_off,omitempty"`
	CleaningTopic          string `json:"cleaning_topic,omitempty"`
	ChargingTopic          string `json:"charging_topic,omitempty"`
	PayloadNotAvailable    string `json:"payload_not_available,omitempty"`
	PayloadTurnOn          string `json:"payload_turn_on,omitempty"`
	ErrorTemplate          string `json:"error_template,omitempty"`
	PayloadStop            string `json:"payload_stop,omitempty"`
	Name                   string `json:"name,omitempty"`
	PayloadLocate          string `json:"payload_locate,omitempty"`
	FanSpeedTemplate       string `json:"fan_speed_template,omitempty"`
	EnabledByDefault       bool   `json:"enabled_by_default,omitempty"`
	FanSpeedTopic          string `json:"fan_speed_topic,omitempty"`
	JsonAttributesTopic    string `json:"json_attributes_topic,omitempty"`
	CleaningTemplate       string `json:"cleaning_template,omitempty"`
	PayloadReturnToBase    string `json:"payload_return_to_base,omitempty"`
}
