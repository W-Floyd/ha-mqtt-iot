package hadiscovery

import mqtt "github.com/eclipse/paho.mqtt.golang"

func (d Light) GetRawId() string {
	return "light"
}
func (d Light) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Light) GetUniqueId() string {
	return d.UniqueId
}
func (d Light) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

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
	MQTT                MQTTFields    `json:"-"`
}

func (d Light) UpdateState()       {}
func (d Light) Subscribe()         {}
func (d Light) UnSubscribe()       {}
func (d Light) AnnounceAvailable() {}
func (d Light) Initialize() {
	d.Retain = false
	d.PopulateDevice()
	d.PopulateTopics()
	d.AddMessageHandler()
}
func (d Light) PopulateTopics() {
	if d.AvailabilityFunc != nil {
		d.AvailabilityTopic = GetTopic(d, "availability_topic")
	}
	if d.CommandFunc != nil {
		d.CommandTopic = GetTopic(d, "command_topic")
		topicStore[d.CommandTopic] = &d.CommandFunc
	}
	if d.StateFunc != nil {
		d.StateTopic = GetTopic(d, "state_topic")
	}
}
