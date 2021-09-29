package commonHomeAssistant

// Opinionation: we will use base level ' availability_topic' and so on, instead of `availability` list and subobject.

type Device struct {
	Device struct {
		Connections
		Identifiers
		Manufacturer
		Model
		Name
		SuggestedArea
		SWVersion
		ViaDevice
	} `json:"device,omitempty"`
}
type Connections struct {
	Connections map[string]string `json:"connections,omitempty"`
}
type Identifiers struct {
	Identifiers []string `json:"identifiers,omitempty"`
}
type Manufacturer struct {
	Manufacturer string `json:"manufacturer,omitempty"`
}
type Model struct {
	Model string `json:"model,omitempty"`
}
type SWVersion struct {
	SWVersion string `json:"sw_version,omitempty"`
}
type ViaDevice struct {
	ViaDevice string `json:"via_device,omitempty"`
}
type Name struct {
	Name string `json:"name,omitempty"`
}
type SuggestedArea struct {
	SuggestedArea string `json:"suggested_area,omitempty"`
}
type StateTopic struct {
	StateTopic string `json:"state_topic,omitempty"`
}
type CommandTemplate struct {
	CommandTemplate string `json:"command_template,omitempty"`
}
type EnabledByDefault struct {
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
}
type ValueTemplate struct {
	ValueTemplate string `json:"value_template,omitempty"`
}
type CodeDisarmRequired struct {
	CodeDisarmRequired bool `json:"code_disarm_required,omitempty"`
}
type Icon struct {
	Icon string `json:"icon,omitempty"`
}
type JsonAttributesTemplate struct {
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
}
type PayloadAvailable struct {
	PayloadAvailable string `json:"payload_available,omitempty"`
}
type AvailabilityTopic struct {
	AvailabilityTopic string `json:"availability_topic,omitempty"`
}
type Retain struct {
	Retain bool `json:"retain,omitempty"`
}
type AvailabilityMode struct {
	AvailabilityMode string `json:"availability_mode,omitempty"`
}
type CommandTopic struct {
	CommandTopic string `json:"command_topic,omitempty"`
}
type PayloadArmHome struct {
	PayloadArmHome string `json:"payload_arm_home,omitempty"`
}
type PayloadArmNight struct {
	PayloadArmNight string `json:"payload_arm_night,omitempty"`
}
type PayloadArmVacation struct {
	PayloadArmVacation string `json:"payload_arm_vacation,omitempty"`
}
type UniqueId struct {
	UniqueId string `json:"unique_id,omitempty"`
}
type Code struct {
	Code string `json:"code,omitempty"`
}
type JsonAttributesTopic struct {
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
}
type PayloadArmAway struct {
	PayloadArmAway string `json:"payload_arm_away,omitempty"`
}
type Qos struct {
	Qos int `json:"qos,omitempty"`
}
type PayloadDisarm struct {
	PayloadDisarm string `json:"payload_disarm,omitempty"`
}
type PayloadNotAvailable struct {
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
}
type CodeArmRequired struct {
	CodeArmRequired bool `json:"code_arm_required,omitempty"`
}
type PayloadArmCustomBypass struct {
	PayloadArmCustomBypass string `json:"payload_arm_custom_bypass,omitempty"`
}
type PayloadOn struct {
	PayloadOn string `json:"payload_on,omitempty"`
}
type OffDelay struct {
	OffDelay int `json:"off_delay,omitempty"`
}
type ExpireAfter struct {
	ExpireAfter int `json:"expire_after,omitempty"`
}
type PayloadOff struct {
	PayloadOff string `json:"payload_off,omitempty"`
}
type DeviceClass struct {
	DeviceClass string `json:"device_class,omitempty"`
}
type ForceUpdate struct {
	ForceUpdate bool `json:"force_update,omitempty"`
}
type Topic struct {
	Topic string `json:"topic,omitempty"`
}
type PositionTopic struct {
	PositionTopic string `json:"position_topic,omitempty"`
}
type TiltMin struct {
	TiltMin int `json:"tilt_min,omitempty"`
}
type TiltOptimistic struct {
	TiltOptimistic bool `json:"tilt_optimistic,omitempty"`
}
type TiltStatusTopic struct {
	TiltStatusTopic string `json:"tilt_status_topic,omitempty"`
}
type PayloadStop struct {
	PayloadStop string `json:"payload_stop,omitempty"`
}
type PositionClosed struct {
	PositionClosed int `json:"position_closed,omitempty"`
}
type StateOpen struct {
	StateOpen string `json:"state_open,omitempty"`
}
type TiltClosedValue struct {
	TiltClosedValue int `json:"tilt_closed_value,omitempty"`
}
type TiltCommandTopic struct {
	TiltCommandTopic string `json:"tilt_command_topic,omitempty"`
}
type TiltCommandTemplate struct {
	TiltCommandTemplate string `json:"tilt_command_template,omitempty"`
}
type TiltMax struct {
	TiltMax int `json:"tilt_max,omitempty"`
}
type PayloadOpen struct {
	PayloadOpen string `json:"payload_open,omitempty"`
}
type StateStopped struct {
	StateStopped string `json:"state_stopped,omitempty"`
}
type Optimistic struct {
	Optimistic bool `json:"optimistic,omitempty"`
}
type PayloadClose struct {
	PayloadClose string `json:"payload_close,omitempty"`
}
type PositionTemplate struct {
	PositionTemplate string `json:"position_template,omitempty"`
}
type StateClosing struct {
	StateClosing string `json:"state_closing,omitempty"`
}
type TiltOpenedValue struct {
	TiltOpenedValue int `json:"tilt_opened_value,omitempty"`
}
type SetPositionTopic struct {
	SetPositionTopic string `json:"set_position_topic,omitempty"`
}
type StateOpening struct {
	StateOpening string `json:"state_opening,omitempty"`
}
type SetPositionTemplate struct {
	SetPositionTemplate string `json:"set_position_template,omitempty"`
}
type PositionOpen struct {
	PositionOpen int `json:"position_open,omitempty"`
}
type StateClosed struct {
	StateClosed string `json:"state_closed,omitempty"`
}
type TiltStatusTemplate struct {
	TiltStatusTemplate string `json:"tilt_status_template,omitempty"`
}
type SourceType struct {
	SourceType string `json:"source_type,omitempty"`
}
type PayloadHome struct {
	PayloadHome string `json:"payload_home,omitempty"`
}
type PayloadNotHome struct {
	PayloadNotHome string `json:"payload_not_home,omitempty"`
}
type Type struct {
	Type string `json:"type,omitempty"`
}
type Subtype struct {
	Subtype string `json:"subtype,omitempty"`
}
type AutomationType struct {
	AutomationType string `json:"automation_type,omitempty"`
}
type Payload struct {
	Payload string `json:"payload,omitempty"`
}

type PresetModeCommandTopic struct {
	PresetModeCommandTopic string `json:"preset_mode_command_topic,omitempty"`
}
type StateValueTemplate struct {
	StateValueTemplate string `json:"state_value_template,omitempty"`
}
type PercentageStateTopic struct {
	PercentageStateTopic string `json:"percentage_state_topic,omitempty"`
}
type PercentageValueTemplate struct {
	PercentageValueTemplate string `json:"percentage_value_template,omitempty"`
}
type PresetModeStateTopic struct {
	PresetModeStateTopic string `json:"preset_mode_state_topic,omitempty"`
}
type PresetModes struct {
	PresetModes []string `json:"preset_modes,omitempty"`
}
type PayloadOscillationOff struct {
	PayloadOscillationOff string `json:"payload_oscillation_off,omitempty"`
}
type PayloadOscillationOn struct {
	PayloadOscillationOn string `json:"payload_oscillation_on,omitempty"`
}
type SpeedRangeMin struct {
	SpeedRangeMin int `json:"speed_range_min,omitempty"`
}
type OscillationCommandTopic struct {
	OscillationCommandTopic string `json:"oscillation_command_topic,omitempty"`
}
type OscillationStateTopic struct {
	OscillationStateTopic string `json:"oscillation_state_topic,omitempty"`
}
type PresetModeCommandTemplate struct {
	PresetModeCommandTemplate string `json:"preset_mode_command_template,omitempty"`
}
type PayloadResetPresetMode struct {
	PayloadResetPresetMode string `json:"payload_reset_preset_mode,omitempty"`
}
type OscillationValueTemplate struct {
	OscillationValueTemplate string `json:"oscillation_value_template,omitempty"`
}
type PayloadResetPercentage struct {
	PayloadResetPercentage string `json:"payload_reset_percentage,omitempty"`
}
type PercentageCommandTemplate struct {
	PercentageCommandTemplate string `json:"percentage_command_template,omitempty"`
}
type PercentageCommandTopic struct {
	PercentageCommandTopic string `json:"percentage_command_topic,omitempty"`
}
type PresetModeValueTemplate struct {
	PresetModeValueTemplate string `json:"preset_mode_value_template,omitempty"`
}
type SpeedRangeMax struct {
	SpeedRangeMax int `json:"speed_range_max,omitempty"`
}
type OscillationCommandTemplate struct {
	OscillationCommandTemplate string `json:"oscillation_command_template,omitempty"`
}

type HoldCommandTopic struct {
	HoldCommandTopic string `json:"hold_command_topic,omitempty"`
}
type SwingModeCommandTemplate struct {
	SwingModeCommandTemplate string `json:"swing_mode_command_template,omitempty"`
}
type ModeCommandTemplate struct {
	ModeCommandTemplate string `json:"mode_command_template,omitempty"`
}
type TemperatureCommandTopic struct {
	TemperatureCommandTopic string `json:"temperature_command_topic,omitempty"`
}
type TemperatureLowStateTemplate struct {
	TemperatureLowStateTemplate string `json:"temperature_low_state_template,omitempty"`
}
type Initial struct {
	Initial int `json:"initial,omitempty"`
}
type TemperatureUnit struct {
	TemperatureUnit string `json:"temperature_unit,omitempty"`
}
type AwayModeStateTopic struct {
	AwayModeStateTopic string `json:"away_mode_state_topic,omitempty"`
}
type MinTemp struct {
	MinTemp float64 `json:"min_temp,omitempty"`
}
type PowerCommandTopic struct {
	PowerCommandTopic string `json:"power_command_topic,omitempty"`
}
type TemperatureStateTopic struct {
	TemperatureStateTopic string `json:"temperature_state_topic,omitempty"`
}
type AuxCommandTopic struct {
	AuxCommandTopic string `json:"aux_command_topic,omitempty"`
}
type FanModeCommandTemplate struct {
	FanModeCommandTemplate string `json:"fan_mode_command_template,omitempty"`
}
type SwingModeCommandTopic struct {
	SwingModeCommandTopic string `json:"swing_mode_command_topic,omitempty"`
}
type SwingModeStateTemplate struct {
	SwingModeStateTemplate string `json:"swing_mode_state_template,omitempty"`
}
type SwingModeStateTopic struct {
	SwingModeStateTopic string `json:"swing_mode_state_topic,omitempty"`
}
type AuxStateTemplate struct {
	AuxStateTemplate string `json:"aux_state_template,omitempty"`
}
type FanModes struct {
	FanModes []string `json:"fan_modes,omitempty"`
}
type HoldStateTemplate struct {
	HoldStateTemplate string `json:"hold_state_template,omitempty"`
}
type Modes struct {
	Modes []string `json:"modes,omitempty"`
}
type TemperatureHighCommandTemplate struct {
	TemperatureHighCommandTemplate string `json:"temperature_high_command_template,omitempty"`
}
type TemperatureStateTemplate struct {
	TemperatureStateTemplate string `json:"temperature_state_template,omitempty"`
}
type AuxStateTopic struct {
	AuxStateTopic string `json:"aux_state_topic,omitempty"`
}
type FanModeStateTopic struct {
	FanModeStateTopic string `json:"fan_mode_state_topic,omitempty"`
}
type SwingModes struct {
	SwingModes []string `json:"swing_modes,omitempty"`
}
type TemperatureCommandTemplate struct {
	TemperatureCommandTemplate string `json:"temperature_command_template,omitempty"`
}
type TemperatureHighStateTopic struct {
	TemperatureHighStateTopic string `json:"temperature_high_state_topic,omitempty"`
}
type TempStep struct {
	TempStep float64 `json:"temp_step,omitempty"`
}
type AwayModeCommandTopic struct {
	AwayModeCommandTopic string `json:"away_mode_command_topic,omitempty"`
}
type HoldStateTopic struct {
	HoldStateTopic string `json:"hold_state_topic,omitempty"`
}
type ModeStateTemplate struct {
	ModeStateTemplate string `json:"mode_state_template,omitempty"`
}
type Precision struct {
	Precision float64 `json:"precision,omitempty"`
}
type TemperatureHighStateTemplate struct {
	TemperatureHighStateTemplate string `json:"temperature_high_state_template,omitempty"`
}
type ActionTopic struct {
	ActionTopic string `json:"action_topic,omitempty"`
}
type FanModeCommandTopic struct {
	FanModeCommandTopic string `json:"fan_mode_command_topic,omitempty"`
}
type TemperatureLowCommandTemplate struct {
	TemperatureLowCommandTemplate string `json:"temperature_low_command_template,omitempty"`
}
type ModeCommandTopic struct {
	ModeCommandTopic string `json:"mode_command_topic,omitempty"`
}
type TemperatureLowCommandTopic struct {
	TemperatureLowCommandTopic string `json:"temperature_low_command_topic,omitempty"`
}
type ModeStateTopic struct {
	ModeStateTopic string `json:"mode_state_topic,omitempty"`
}
type TemperatureHighCommandTopic struct {
	TemperatureHighCommandTopic string `json:"temperature_high_command_topic,omitempty"`
}
type TemperatureLowStateTopic struct {
	TemperatureLowStateTopic string `json:"temperature_low_state_topic,omitempty"`
}
type ActionTemplate struct {
	ActionTemplate string `json:"action_template,omitempty"`
}
type HoldModes struct {
	HoldModes []string `json:"hold_modes,omitempty"`
}
type CurrentTemperatureTopic struct {
	CurrentTemperatureTopic string `json:"current_temperature_topic,omitempty"`
}
type MaxTemp struct {
	MaxTemp float64 `json:"max_temp,omitempty"`
}
type SendIfOff struct {
	SendIfOff bool `json:"send_if_off,omitempty"`
}
type AwayModeStateTemplate struct {
	AwayModeStateTemplate string `json:"away_mode_state_template,omitempty"`
}
type CurrentTemperatureTemplate struct {
	CurrentTemperatureTemplate string `json:"current_temperature_template,omitempty"`
}
type FanModeStateTemplate struct {
	FanModeStateTemplate string `json:"fan_mode_state_template,omitempty"`
}
type HoldCommandTemplate struct {
	HoldCommandTemplate string `json:"hold_command_template,omitempty"`
}

type RgbCommandTopic struct {
	RgbCommandTopic string `json:"rgb_command_topic,omitempty"`
}
type Schema struct {
	Schema string `json:"schema,omitempty"`
}
type BrightnessScale struct {
	BrightnessScale int `json:"brightness_scale,omitempty"`
}
type ColorTempValueTemplate struct {
	ColorTempValueTemplate string `json:"color_temp_value_template,omitempty"`
}
type HsCommandTopic struct {
	HsCommandTopic string `json:"hs_command_topic,omitempty"`
}
type WhiteCommandTopic struct {
	WhiteCommandTopic string `json:"white_command_topic,omitempty"`
}
type BrightnessStateTopic struct {
	BrightnessStateTopic string `json:"brightness_state_topic,omitempty"`
}
type MaxMireds struct {
	MaxMireds int `json:"max_mireds,omitempty"`
}
type RgbCommandTemplate struct {
	RgbCommandTemplate string `json:"rgb_command_template,omitempty"`
}
type RgbValueTemplate struct {
	RgbValueTemplate string `json:"rgb_value_template,omitempty"`
}
type ColorModeValueTemplate struct {
	ColorModeValueTemplate string `json:"color_mode_value_template,omitempty"`
}
type ColorTempCommandTopic struct {
	ColorTempCommandTopic string `json:"color_temp_command_topic,omitempty"`
}
type ColorTempStateTopic struct {
	ColorTempStateTopic string `json:"color_temp_state_topic,omitempty"`
}
type HsValueTemplate struct {
	HsValueTemplate string `json:"hs_value_template,omitempty"`
}
type OnCommandType struct {
	OnCommandType string `json:"on_command_type,omitempty"`
}
type RgbStateTopic struct {
	RgbStateTopic string `json:"rgb_state_topic,omitempty"`
}
type ColorTempCommandTemplate struct {
	ColorTempCommandTemplate string `json:"color_temp_command_template,omitempty"`
}
type EffectStateTopic struct {
	EffectStateTopic string `json:"effect_state_topic,omitempty"`
}
type EffectValueTemplate struct {
	EffectValueTemplate string `json:"effect_value_template,omitempty"`
}
type BrightnessCommandTopic struct {
	BrightnessCommandTopic string `json:"brightness_command_topic,omitempty"`
}
type EffectList struct {
	EffectList string `json:"effect_list,omitempty"`
}
type MinMireds struct {
	MinMireds int `json:"min_mireds,omitempty"`
}
type XyCommandTopic struct {
	XyCommandTopic string `json:"xy_command_topic,omitempty"`
}
type XyValueTemplate struct {
	XyValueTemplate string `json:"xy_value_template,omitempty"`
}
type HsStateTopic struct {
	HsStateTopic string `json:"hs_state_topic,omitempty"`
}
type WhiteScale struct {
	WhiteScale int `json:"white_scale,omitempty"`
}
type XyStateTopic struct {
	XyStateTopic string `json:"xy_state_topic,omitempty"`
}
type BrightnessValueTemplate struct {
	BrightnessValueTemplate string `json:"brightness_value_template,omitempty"`
}
type ColorModeStateTopic struct {
	ColorModeStateTopic string `json:"color_mode_state_topic,omitempty"`
}
type EffectCommandTopic struct {
	EffectCommandTopic string `json:"effect_command_topic,omitempty"`
}
type StateUnlocked struct {
	StateUnlocked string `json:"state_unlocked,omitempty"`
}
type PayloadUnlock struct {
	PayloadUnlock string `json:"payload_unlock,omitempty"`
}
type PayloadLock struct {
	PayloadLock string `json:"payload_lock,omitempty"`
}
type StateLocked struct {
	StateLocked string `json:"state_locked,omitempty"`
}
type Step struct {
	Step float64 `json:"step,omitempty"`
}
type Max struct {
	Max float64 `json:"max,omitempty"`
}
type Min struct {
	Min float64 `json:"min,omitempty"`
}
type Options struct {
	Options []string `json:"options,omitempty"`
}
type StateClass struct {
	StateClass string `json:"state_class,omitempty"`
}
type LastResetTopic struct {
	LastResetTopic string `json:"last_reset_topic,omitempty"`
}
type LastResetValueTemplate struct {
	LastResetValueTemplate string `json:"last_reset_value_template,omitempty"`
}
type UnitOfMeasurement struct {
	UnitOfMeasurement string `json:"unit_of_measurement,omitempty"`
}
type StateOff struct {
	StateOff string `json:"state_off,omitempty"`
}
type StateOn struct {
	StateOn string `json:"state_on,omitempty"`
}
type FanSpeedList struct {
	FanSpeedList string `json:"fan_speed_list,omitempty"`
}
type PayloadCleanSpot struct {
	PayloadCleanSpot string `json:"payload_clean_spot,omitempty"`
}
type SetFanSpeedTopic struct {
	SetFanSpeedTopic string `json:"set_fan_speed_topic,omitempty"`
}
type SupportedFeatures struct {
	SupportedFeatures string `json:"supported_features,omitempty"`
}
type SendCommandTopic struct {
	SendCommandTopic string `json:"send_command_topic,omitempty"`
}
type PayloadLocate struct {
	PayloadLocate string `json:"payload_locate,omitempty"`
}
type PayloadPause struct {
	PayloadPause string `json:"payload_pause,omitempty"`
}
type PayloadReturnToBase struct {
	PayloadReturnToBase string `json:"payload_return_to_base,omitempty"`
}
type PayloadStart struct {
	PayloadStart string `json:"payload_start,omitempty"`
}
type BatteryLevelTopic struct {
	BatteryLevelTopic string `json:"battery_level_topic,omitempty"`
}
type DockedTemplate struct {
	DockedTemplate string `json:"docked_template,omitempty"`
}
type ErrorTopic struct {
	ErrorTopic string `json:"error_topic,omitempty"`
}
type PayloadTurnOn struct {
	PayloadTurnOn string `json:"payload_turn_on,omitempty"`
}
type DockedTopic struct {
	DockedTopic string `json:"docked_topic,omitempty"`
}
type FanSpeedTopic struct {
	FanSpeedTopic string `json:"fan_speed_topic,omitempty"`
}
type PayloadTurnOff struct {
	PayloadTurnOff string `json:"payload_turn_off,omitempty"`
}
type PayloadStartPause struct {
	PayloadStartPause string `json:"payload_start_pause,omitempty"`
}
type BatteryLevelTemplate struct {
	BatteryLevelTemplate string `json:"battery_level_template,omitempty"`
}
type CleaningTemplate struct {
	CleaningTemplate string `json:"cleaning_template,omitempty"`
}
type FanSpeedTemplate struct {
	FanSpeedTemplate string `json:"fan_speed_template,omitempty"`
}
type ChargingTemplate struct {
	ChargingTemplate string `json:"charging_template,omitempty"`
}
type ChargingTopic struct {
	ChargingTopic string `json:"charging_topic,omitempty"`
}
type CleaningTopic struct {
	CleaningTopic string `json:"cleaning_topic,omitempty"`
}
type ErrorTemplate struct {
	ErrorTemplate string `json:"error_template,omitempty"`
}
