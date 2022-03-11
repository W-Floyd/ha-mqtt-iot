package hadiscovery

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Light struct {
	AvailabilityMode     string                          `json:"availability_mode"`
	AvailabilityTemplate string                          `json:"availability_template"`
	AvailabilityTopic    string                          `json:"availability_topic"`
	AvailabilityFunc     func() string                   `json:"-"`
	Brightness           bool                            `json:"brightness"`
	BrightnessScale      int                             `json:"brightness_scale"`
	ColorMode            bool                            `json:"color_mode"`
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
	Effect              bool          `json:"effect"`
	EffectList          []string      `json:"effect_list"`
	EnabledByDefault    bool          `json:"enabled_by_default"`
	Encoding            string        `json:"encoding"`
	EntityCategory      string        `json:"entity_category"`
	FlashTimeLong       int           `json:"flash_time_long"`
	FlashTimeShort      int           `json:"flash_time_short"`
	Icon                string        `json:"icon"`
	MaxMireds           int           `json:"max_mireds"`
	MinMireds           int           `json:"min_mireds"`
	Name                string        `json:"name"`
	ObjectId            string        `json:"object_id"`
	Optimistic          bool          `json:"optimistic"`
	PayloadAvailable    string        `json:"payload_available"`
	PayloadNotAvailable string        `json:"payload_not_available"`
	Qos                 int           `json:"qos"`
	Retain              bool          `json:"retain"`
	Schema              string        `json:"schema"`
	StateTopic          string        `json:"state_topic"`
	StateFunc           func() string `json:"-"`
	SupportedColorModes []string      `json:"supported_color_modes"`
	UniqueId            string        `json:"unique_id"`
	RawId               string        `json:"-"`
	MQTT                MQTTFields    `json:"-"`
}
type Device interface {
	GetRawId() string
	GetUniqueId() string
	PopulateDevice()
	PopulateTopics()
	UpdateState()
	Subscribe()
	AddMessageHandler()
}
