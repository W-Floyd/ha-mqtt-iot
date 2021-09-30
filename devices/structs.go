package devices

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// alarm_control_panel
type HADeviceAlarmControlPanel struct {
	// A list of MQTT topics subscribed to receive availability (online/offline)
	// updates. Must not be used together with `availability_topic`.
	Availability struct {
		// The payload that represents the available state.
		PayloadAvailable string `json:"payload_available,omitempty"`
		// The payload that represents the unavailable state.
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		// An MQTT topic subscribed to receive availability (online/offline) updates.
		Topic string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	// When `availability` is configured, this controls the conditions needed to set
	// the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set
	// to `all`, `payload_available` must be received on all configured availability
	// topics before the entity is marked as online. If set to `any`,
	// `payload_available` must be received on at least one configured availability
	// topic before the entity is marked as online. If set to `latest`, the last
	// `payload_available` or `payload_not_available` received on any configured
	// availability topic controls the availability.
	AvailabilityMode string `json:"availability_mode,omitempty"`
	// The MQTT topic subscribed to receive availability (online/offline) updates.
	// Must not be used together with `availability`.
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	// If defined, specifies a code to enable or disable the alarm in the frontend.
	Code string `json:"code,omitempty"`
	// If true the code is required to arm the alarm. If false the code is not
	// validated.
	CodeArmRequired bool `json:"code_arm_required,omitempty"`
	// If true the code is required to disarm the alarm. If false the code is not
	// validated.
	CodeDisarmRequired bool `json:"code_disarm_required,omitempty"`
	// The [template](/docs/configuration/templating/#processing-incoming-data) used
	// for the command payload. Available variables: `action` and `code`.
	CommandTemplate string `json:"command_template,omitempty"`
	// The MQTT topic to publish commands to change the alarm state.
	CommandTopic string `json:"command_topic,omitempty"`
	// Information about the device this alarm panel is a part of to tie it into the
	// [device
	// registry](https://developers.home-assistant.io/docs/en/device_registry_index.html).
	// Only works through [MQTT discovery](/docs/mqtt/discovery/) and when
	// [`unique_id`](#unique_id) is set. At least one of identifiers or connections
	// must be present to identify the device.
	Device struct {
		// A list of connections of the device to the outside world as a list of tuples
		// `[connection_type, connection_identifier]`. For example the MAC address of a
		// network interface: `\"connections\": [[\"mac\", \"02:5b:26:a8:dc:12\"]]`.
		Connections []string `json:"connections,omitempty"`
		// A list of IDs that uniquely identify the device. For example a serial number.
		Identifiers []string `json:"identifiers,omitempty"`
		// The manufacturer of the device.
		Manufacturer string `json:"manufacturer,omitempty"`
		// The model of the device.
		Model string `json:"model,omitempty"`
		// The name of the device.
		Name string `json:"name,omitempty"`
		// Suggest an area if the device isn’t in one yet.
		SuggestedArea string `json:"suggested_area,omitempty"`
		// The firmware version of the device.
		SwVersion string `json:"sw_version,omitempty"`
		// Identifier of a device that routes messages between this device and Home
		// Assistant. Examples of such devices are hubs, or parent devices of a sub-device.
		// This is used to show device topology in Home Assistant.
		ViaDevice string `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity.
	Icon string `json:"icon,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the JSON dictionary from messages received on the
	// `json_attributes_topic`. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration)
	// documentation.
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as
	// sensor attributes. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration)
	// documentation.
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	// The name of the alarm.
	Name string `json:"name,omitempty"`
	// The payload to set armed-away mode on your Alarm Panel.
	PayloadArmAway string `json:"payload_arm_away,omitempty"`
	// The payload to set armed-custom-bypass mode on your Alarm Panel.
	PayloadArmCustomBypass string `json:"payload_arm_custom_bypass,omitempty"`
	// The payload to set armed-home mode on your Alarm Panel.
	PayloadArmHome string `json:"payload_arm_home,omitempty"`
	// The payload to set armed-night mode on your Alarm Panel.
	PayloadArmNight string `json:"payload_arm_night,omitempty"`
	// The payload to set armed-vacation mode on your Alarm Panel.
	PayloadArmVacation string `json:"payload_arm_vacation,omitempty"`
	// The payload that represents the available state.
	PayloadAvailable string `json:"payload_available,omitempty"`
	// The payload to disarm your Alarm Panel.
	PayloadDisarm string `json:"payload_disarm,omitempty"`
	// The payload that represents the unavailable state.
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	// The maximum QoS level of the state topic.
	Qos int `json:"qos,omitempty"`
	// If the published message should have the retain flag on or not.
	Retain bool `json:"retain,omitempty"`
	// The MQTT topic subscribed to receive state updates.
	StateTopic string `json:"state_topic,omitempty"`
	// An ID that uniquely identifies this alarm panel. If two alarm panels have the
	// same unique ID, Home Assistant will raise an exception.
	UniqueId string `json:"unique_id,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the value.
	ValueTemplate string `json:"value_template,omitempty"`
}

type HADeviceAlarmControlPanelFunctions struct {
	Availability struct {
		State func() string
	}
	Command func(mqtt.Message, mqtt.Client)
	State   func() string
}

type HADeviceAlarmControlPanelFunctionsConfig struct {
	Availability struct {
		State []string `json:"state,omitempty"`
	} `json:"availability,omitempty"`
	Command []string `json:"command,omitempty"`
	State   []string `json:"state,omitempty"`
}

// binary_sensor
type HADeviceBinarySensor struct {
	// A list of MQTT topics subscribed to receive availability (online/offline)
	// updates. Must not be used together with `availability_topic`.
	Availability struct {
		// The payload that represents the available state.
		PayloadAvailable string `json:"payload_available,omitempty"`
		// The payload that represents the unavailable state.
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		// An MQTT topic subscribed to receive availability (online/offline) updates.
		Topic string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	// When `availability` is configured, this controls the conditions needed to set
	// the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set
	// to `all`, `payload_available` must be received on all configured availability
	// topics before the entity is marked as online. If set to `any`,
	// `payload_available` must be received on at least one configured availability
	// topic before the entity is marked as online. If set to `latest`, the last
	// `payload_available` or `payload_not_available` received on any configured
	// availability topic controls the availability.
	AvailabilityMode string `json:"availability_mode,omitempty"`
	// The MQTT topic subscribed to receive birth and LWT messages from the MQTT
	// device. If `availability` is not defined, the binary sensor will always be
	// considered `available` and its state will be `on`, `off` or `unknown`. If
	// `availability` is defined, the binary sensor will be considered as `unavailable`
	// by default and the sensor's initial state will be `unavailable`. Must not be
	// used together with `availability`.
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	// Information about the device this binary sensor is a part of to tie it into the
	// [device
	// registry](https://developers.home-assistant.io/docs/en/device_registry_index.html).
	// Only works through [MQTT discovery](/docs/mqtt/discovery/) and when
	// [`unique_id`](#unique_id) is set. At least one of identifiers or connections
	// must be present to identify the device.
	Device struct {
		// A list of connections of the device to the outside world as a list of tuples
		// `[connection_type, connection_identifier]`. For example the MAC address of a
		// network interface: `'connections': ['mac', '02:5b:26:a8:dc:12']`.
		Connections []string `json:"connections,omitempty"`
		// A list of IDs that uniquely identify the device. For example a serial number.
		Identifiers []string `json:"identifiers,omitempty"`
		// The manufacturer of the device.
		Manufacturer string `json:"manufacturer,omitempty"`
		// The model of the device.
		Model string `json:"model,omitempty"`
		// The name of the device.
		Name string `json:"name,omitempty"`
		// Suggest an area if the device isn’t in one yet.
		SuggestedArea string `json:"suggested_area,omitempty"`
		// The firmware version of the device.
		SwVersion string `json:"sw_version,omitempty"`
		// Identifier of a device that routes messages between this device and Home
		// Assistant. Examples of such devices are hubs, or parent devices of a sub-device.
		// This is used to show device topology in Home Assistant.
		ViaDevice string `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	// Sets the [class of the device](/integrations/binary_sensor/#device-class),
	// changing the device state and icon that is displayed on the frontend.
	DeviceClass string `json:"device_class,omitempty"`
	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	// Defines the number of seconds after the sensor's state expires, if it's not
	// updated. After expiry, the sensor's state becomes `unavailable`.
	ExpireAfter int `json:"expire_after,omitempty"`
	// Sends update events (which results in update of [state
	// object](/docs/configuration/state_object/)'s `last_changed`) even if the
	// sensor's state hasn't changed. Useful if you want to have meaningful value
	// graphs in history or want to create an automation that triggers on *every*
	// incoming state message (not only when the sensor's new state is different to the
	// current one).
	ForceUpdate bool `json:"force_update,omitempty"`
	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity.
	Icon string `json:"icon,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the JSON dictionary from messages received on the
	// `json_attributes_topic`. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration)
	// documentation.
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as
	// sensor attributes. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration)
	// documentation.
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	// The name of the binary sensor.
	Name string `json:"name,omitempty"`
	// For sensors that only send `on` state updates (like PIRs), this variable sets a
	// delay in seconds after which the sensor's state will be updated back to `off`.
	OffDelay int `json:"off_delay,omitempty"`
	// The string that represents the `online` state.
	PayloadAvailable string `json:"payload_available,omitempty"`
	// The string that represents the `offline` state.
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	// The string that represents the `off` state. It will be compared to the message
	// in the `state_topic` (see `value_template` for details)
	PayloadOff string `json:"payload_off,omitempty"`
	// The string that represents the `on` state. It will be compared to the message
	// in the `state_topic` (see `value_template` for details)
	PayloadOn string `json:"payload_on,omitempty"`
	// The maximum QoS level to be used when receiving messages.
	Qos int `json:"qos,omitempty"`
	// The MQTT topic subscribed to receive sensor's state.
	StateTopic string `json:"state_topic,omitempty"`
	// An ID that uniquely identifies this sensor. If two sensors have the same unique
	// ID, Home Assistant will raise an exception.
	UniqueId string `json:"unique_id,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// that returns a string to be compared to `payload_on`/`payload_off` or an empty
	// string, in which case the MQTT message will be removed. Available variables:
	// `entity_id`. Remove this option when 'payload_on' and 'payload_off' are
	// sufficient to match your payloads (i.e no pre-processing of original message is
	// required).
	ValueTemplate string `json:"value_template,omitempty"`
}

type HADeviceBinarySensorFunctions struct {
	Availability struct {
		State func() string
	}
	State func() string
}

type HADeviceBinarySensorFunctionsConfig struct {
	Availability struct {
		State []string `json:"state,omitempty"`
	} `json:"availability,omitempty"`
	State []string `json:"state,omitempty"`
}

// camera
type HADeviceCamera struct {
	// A list of MQTT topics subscribed to receive availability (online/offline)
	// updates. Must not be used together with `availability_topic`.
	Availability struct {
		// The payload that represents the available state.
		PayloadAvailable string `json:"payload_available,omitempty"`
		// The payload that represents the unavailable state.
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		// An MQTT topic subscribed to receive availability (online/offline) updates.
		Topic string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	// When `availability` is configured, this controls the conditions needed to set
	// the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set
	// to `all`, `payload_available` must be received on all configured availability
	// topics before the entity is marked as online. If set to `any`,
	// `payload_available` must be received on at least one configured availability
	// topic before the entity is marked as online. If set to `latest`, the last
	// `payload_available` or `payload_not_available` received on any configured
	// availability topic controls the availability.
	AvailabilityMode string `json:"availability_mode,omitempty"`
	// The MQTT topic subscribed to receive availability (online/offline) updates.
	// Must not be used together with `availability`.
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	// Information about the device this camera is a part of to tie it into the
	// [device
	// registry](https://developers.home-assistant.io/docs/en/device_registry_index.html).
	// Only works through [MQTT discovery](/docs/mqtt/discovery/) and when
	// [`unique_id`](#unique_id) is set. At least one of identifiers or connections
	// must be present to identify the device.
	Device struct {
		// A list of connections of the device to the outside world as a list of tuples
		// `[connection_type, connection_identifier]`. For example the MAC address of a
		// network interface: `\"connections\": [\"mac\", \"02:5b:26:a8:dc:12\"]`.
		Connections []string `json:"connections,omitempty"`
		// A list of IDs that uniquely identify the device. For example a serial number.
		Identifiers []string `json:"identifiers,omitempty"`
		// The manufacturer of the device.
		Manufacturer string `json:"manufacturer,omitempty"`
		// The model of the device.
		Model string `json:"model,omitempty"`
		// The name of the device.
		Name string `json:"name,omitempty"`
		// Suggest an area if the device isn’t in one yet.
		SuggestedArea string `json:"suggested_area,omitempty"`
		// The firmware version of the device.
		SwVersion string `json:"sw_version,omitempty"`
		// Identifier of a device that routes messages between this device and Home
		// Assistant. Examples of such devices are hubs, or parent devices of a sub-device.
		// This is used to show device topology in Home Assistant.
		ViaDevice string `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity.
	Icon string `json:"icon,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the JSON dictionary from messages received on the
	// `json_attributes_topic`.
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as
	// sensor attributes. Implies `force_update` of the current sensor state when a
	// message is received on this topic.
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	// The name of the camera.
	Name string `json:"name,omitempty"`
	// The MQTT topic to subscribe to.
	Topic string `json:"topic,omitempty"`
	// An ID that uniquely identifies this camera. If two cameras have the same unique
	// ID Home Assistant will raise an exception.
	UniqueId string `json:"unique_id,omitempty"`
}

type HADeviceCameraFunctions struct {
	Availability struct {
		State func() string
	}
	State func() string
}

type HADeviceCameraFunctionsConfig struct {
	Availability struct {
		State []string `json:"state,omitempty"`
	} `json:"availability,omitempty"`
	State []string `json:"state,omitempty"`
}

// cover
type HADeviceCover struct {
	// A list of MQTT topics subscribed to receive availability (online/offline)
	// updates. Must not be used together with `availability_topic`.
	Availability struct {
		// The payload that represents the available state.
		PayloadAvailable string `json:"payload_available,omitempty"`
		// The payload that represents the unavailable state.
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		// An MQTT topic subscribed to receive availability (online/offline) updates.
		Topic string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	// When `availability` is configured, this controls the conditions needed to set
	// the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set
	// to `all`, `payload_available` must be received on all configured availability
	// topics before the entity is marked as online. If set to `any`,
	// `payload_available` must be received on at least one configured availability
	// topic before the entity is marked as online. If set to `latest`, the last
	// `payload_available` or `payload_not_available` received on any configured
	// availability topic controls the availability.
	AvailabilityMode string `json:"availability_mode,omitempty"`
	// The MQTT topic subscribed to to receive birth and LWT messages from the MQTT
	// cover device. If an `availability` topic is not defined, the cover availability
	// state will always be `available`. If an `availability` topic is defined, the
	// cover availability state will be `unavailable` by default. Must not be used
	// together with `availability`.
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	// The MQTT topic to publish commands to control the cover.
	CommandTopic string `json:"command_topic,omitempty"`
	// Information about the device this cover is a part of to tie it into the [device
	// registry](https://developers.home-assistant.io/docs/en/device_registry_index.html).
	// Only works through [MQTT discovery](/docs/mqtt/discovery/) and when
	// [`unique_id`](#unique_id) is set. At least one of identifiers or connections
	// must be present to identify the device.
	Device struct {
		// A list of connections of the device to the outside world as a list of tuples
		// `[connection_type, connection_identifier]`. For example the MAC address of a
		// network interface: `\"connections\": [\"mac\", \"02:5b:26:a8:dc:12\"]`.
		Connections []string `json:"connections,omitempty"`
		// A list of IDs that uniquely identify the device. For example a serial number.
		Identifiers []string `json:"identifiers,omitempty"`
		// The manufacturer of the device.
		Manufacturer string `json:"manufacturer,omitempty"`
		// The model of the device.
		Model string `json:"model,omitempty"`
		// The name of the device.
		Name string `json:"name,omitempty"`
		// Suggest an area if the device isn’t in one yet.
		SuggestedArea string `json:"suggested_area,omitempty"`
		// The firmware version of the device.
		SwVersion string `json:"sw_version,omitempty"`
		// Identifier of a device that routes messages between this device and Home
		// Assistant. Examples of such devices are hubs, or parent devices of a sub-device.
		// This is used to show device topology in Home Assistant.
		ViaDevice string `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	// Sets the [class of the device](/integrations/cover/), changing the device state
	// and icon that is displayed on the frontend.
	DeviceClass string `json:"device_class,omitempty"`
	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity.
	Icon string `json:"icon,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the JSON dictionary from messages received on the
	// `json_attributes_topic`. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration)
	// documentation.
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as
	// sensor attributes. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration)
	// documentation.
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	// The name of the cover.
	Name string `json:"name,omitempty"`
	// Flag that defines if switch works in optimistic mode.
	Optimistic bool `json:"optimistic,omitempty"`
	// The payload that represents the online state.
	PayloadAvailable string `json:"payload_available,omitempty"`
	// The command payload that closes the cover.
	PayloadClose string `json:"payload_close,omitempty"`
	// The payload that represents the offline state.
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	// The command payload that opens the cover.
	PayloadOpen string `json:"payload_open,omitempty"`
	// The command payload that stops the cover.
	PayloadStop string `json:"payload_stop,omitempty"`
	// Number which represents closed position.
	PositionClosed int `json:"position_closed,omitempty"`
	// Number which represents open position.
	PositionOpen int `json:"position_open,omitempty"`
	// Defines a [template](/topics/templating/) that can be used to extract the
	// payload for the `position_topic` topic. Within the template the following
	// variables are available: `entity_id`, `position_open`; `position_closed`;
	// `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's
	// attributes with help of the [states](/docs/configuration/templating/#states)
	// template function;
	PositionTemplate string `json:"position_template,omitempty"`
	// The MQTT topic subscribed to receive cover position messages.
	PositionTopic string `json:"position_topic,omitempty"`
	// The maximum QoS level to be used when receiving and publishing messages.
	Qos int `json:"qos,omitempty"`
	// Defines if published messages should have the retain flag set.
	Retain bool `json:"retain,omitempty"`
	// Defines a [template](/topics/templating/) to define the position to be sent to
	// the `set_position_topic` topic. Incoming position value is available for use in
	// the template `{% raw %}{{ position }}{% endraw %}`. Within the template the
	// following variables are available: `entity_id`, `position`, the target position
	// in percent; `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The
	// `entity_id` can be used to reference the entity's attributes with help of the
	// [states](/docs/configuration/templating/#states) template function;
	SetPositionTemplate string `json:"set_position_template,omitempty"`
	// The MQTT topic to publish position commands to. You need to set position_topic
	// as well if you want to use position topic. Use template if position topic wants
	// different values than within range `position_closed` - `position_open`. If
	// template is not defined and `position_closed != 100` and `position_open != 0`
	// then proper position value is calculated from percentage position.
	SetPositionTopic string `json:"set_position_topic,omitempty"`
	// The payload that represents the closed state.
	StateClosed string `json:"state_closed,omitempty"`
	// The payload that represents the closing state.
	StateClosing string `json:"state_closing,omitempty"`
	// The payload that represents the open state.
	StateOpen string `json:"state_open,omitempty"`
	// The payload that represents the opening state.
	StateOpening string `json:"state_opening,omitempty"`
	// The payload that represents the stopped state (for covers that do not report
	// `open`/`closed` state).
	StateStopped string `json:"state_stopped,omitempty"`
	// The MQTT topic subscribed to receive cover state messages. State topic can only
	// read (`open`, `opening`, `closed`, `closing` or `stopped`) state.
	StateTopic string `json:"state_topic,omitempty"`
	// The value that will be sent on a `close_cover_tilt` command.
	TiltClosedValue int `json:"tilt_closed_value,omitempty"`
	// Defines a [template](/topics/templating/) that can be used to extract the
	// payload for the `tilt_command_topic` topic. Within the template the following
	// variables are available: `entity_id`, `tilt_position`, the target tilt position
	// in percent; `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The
	// `entity_id` can be used to reference the entity's attributes with help of the
	// [states](/docs/configuration/templating/#states) template function;
	TiltCommandTemplate string `json:"tilt_command_template,omitempty"`
	// The MQTT topic to publish commands to control the cover tilt.
	TiltCommandTopic string `json:"tilt_command_topic,omitempty"`
	// The maximum tilt value.
	TiltMax int `json:"tilt_max,omitempty"`
	// The minimum tilt value.
	TiltMin int `json:"tilt_min,omitempty"`
	// The value that will be sent on an `open_cover_tilt` command.
	TiltOpenedValue int `json:"tilt_opened_value,omitempty"`
	// Flag that determines if tilt works in optimistic mode.
	TiltOptimistic bool `json:"tilt_optimistic,omitempty"`
	// Defines a [template](/topics/templating/) that can be used to extract the
	// payload for the `tilt_status_topic` topic. Within the template the following
	// variables are available: `entity_id`, `position_open`; `position_closed`;
	// `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's
	// attributes with help of the [states](/docs/configuration/templating/#states)
	// template function;
	TiltStatusTemplate string `json:"tilt_status_template,omitempty"`
	// The MQTT topic subscribed to receive tilt status update values.
	TiltStatusTopic string `json:"tilt_status_topic,omitempty"`
	// An ID that uniquely identifies this cover. If two covers have the same unique
	// ID, Home Assistant will raise an exception.
	UniqueId string `json:"unique_id,omitempty"`
	// Defines a [template](/topics/templating/) that can be used to extract the
	// payload for the `state_topic` topic.
	ValueTemplate string `json:"value_template,omitempty"`
}

type HADeviceCoverFunctions struct {
	Availability struct {
		State func() string
	}
	Command     func(mqtt.Message, mqtt.Client)
	SetPosition func(mqtt.Message, mqtt.Client)
	State       func() string
	TiltCommand func(mqtt.Message, mqtt.Client)
	TiltStatus  func() string
}

type HADeviceCoverFunctionsConfig struct {
	Availability struct {
		State []string `json:"state,omitempty"`
	} `json:"availability,omitempty"`
	Command     []string `json:"command,omitempty"`
	SetPosition []string `json:"set-position,omitempty"`
	State       []string `json:"state,omitempty"`
	TiltCommand []string `json:"tilt-command,omitempty"`
	TiltStatus  []string `json:"tilt-status,omitempty"`
}

// device_tracker
type HADeviceDeviceTracker struct {
	// List of devices with their topic.
	Devices []string `json:"devices,omitempty"`
	// The payload value that represents the 'home' state for the device.
	PayloadHome string `json:"payload_home,omitempty"`
	// The payload value that represents the 'not_home' state for the device.
	PayloadNotHome string `json:"payload_not_home,omitempty"`
	// The QoS level of the topic.
	Qos int `json:"qos,omitempty"`
	// Attribute of a device tracker that affects state when being used to track a
	// [person](/integrations/person/). Valid options are `gps`, `router`, `bluetooth`,
	// or `bluetooth_le`.
	SourceType string `json:"source_type,omitempty"`
}

type HADeviceDeviceTrackerFunctions struct {
}

type HADeviceDeviceTrackerFunctionsConfig struct {
}

// device_trigger
type HADeviceDeviceTrigger struct {
	// The type of automation, must be 'trigger'.
	AutomationType string `json:"automation_type,omitempty"`
	// Information about the device this device trigger is a part of to tie it into
	// the [device
	// registry](https://developers.home-assistant.io/docs/en/device_registry_index.html).
	// At least one of identifiers or connections must be present to identify the
	// device.
	Device struct {
		// A list of connections of the device to the outside world as a list of tuples
		// `[connection_type, connection_identifier]`. For example the MAC address of a
		// network interface: `'connections': ['mac', '02:5b:26:a8:dc:12']`.
		Connections []string `json:"connections,omitempty"`
		// A list of IDs that uniquely identify the device. For example a serial number.
		Identifiers []string `json:"identifiers,omitempty"`
		// The manufacturer of the device.
		Manufacturer string `json:"manufacturer,omitempty"`
		// The model of the device.
		Model string `json:"model,omitempty"`
		// The name of the device.
		Name string `json:"name,omitempty"`
		// Suggest an area if the device isn’t in one yet.
		SuggestedArea string `json:"suggested_area,omitempty"`
		// The firmware version of the device.
		SwVersion string `json:"sw_version,omitempty"`
		// Identifier of a device that routes messages between this device and Home
		// Assistant. Examples of such devices are hubs, or parent devices of a sub-device.
		// This is used to show device topology in Home Assistant.
		ViaDevice string `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	// Optional payload to match the payload being sent over the topic.
	Payload string `json:"payload,omitempty"`
	// The maximum QoS level to be used when receiving messages.
	Qos int `json:"qos,omitempty"`
	// The subtype of the trigger, e.g. `button_1`. Entries supported by the frontend:
	// `turn_on`, `turn_off`, `button_1`, `button_2`, `button_3`, `button_4`,
	// `button_5`, `button_6`. If set to an unsupported value, will render as `subtype
	// type`, e.g. `left_button pressed` with `type` set to `button_short_press` and
	// `subtype` set to `left_button`
	Subtype string `json:"subtype,omitempty"`
	// The MQTT topic subscribed to receive trigger events.
	Topic string `json:"topic,omitempty"`
	// The type of the trigger, e.g. `button_short_press`. Entries supported by the
	// frontend: `button_short_press`, `button_short_release`, `button_long_press`,
	// `button_long_release`, `button_double_press`, `button_triple_press`,
	// `button_quadruple_press`, `button_quintuple_press`. If set to an unsupported
	// value, will render as `subtype type`, e.g. `button_1 spammed` with `type` set to
	// `spammed` and `subtype` set to `button_1`
	Type string `json:"type,omitempty"`
}

type HADeviceDeviceTriggerFunctions struct {
	State func() string
}

type HADeviceDeviceTriggerFunctionsConfig struct {
	State []string `json:"state,omitempty"`
}

// fan
type HADeviceFan struct {
	// A list of MQTT topics subscribed to receive availability (online/offline)
	// updates. Must not be used together with `availability_topic`.
	Availability struct {
		// The payload that represents the available state.
		PayloadAvailable string `json:"payload_available,omitempty"`
		// The payload that represents the unavailable state.
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		// An MQTT topic subscribed to receive availability (online/offline) updates.
		Topic string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	// When `availability` is configured, this controls the conditions needed to set
	// the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set
	// to `all`, `payload_available` must be received on all configured availability
	// topics before the entity is marked as online. If set to `any`,
	// `payload_available` must be received on at least one configured availability
	// topic before the entity is marked as online. If set to `latest`, the last
	// `payload_available` or `payload_not_available` received on any configured
	// availability topic controls the availability.
	AvailabilityMode string `json:"availability_mode,omitempty"`
	// The MQTT topic subscribed to receive availability (online/offline) updates.
	// Must not be used together with `availability`.
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to generate the payload to send to `command_topic`.
	CommandTemplate string `json:"command_template,omitempty"`
	// The MQTT topic to publish commands to change the fan state.
	CommandTopic string `json:"command_topic,omitempty"`
	// Information about the device this fan is a part of to tie it into the [device
	// registry](https://developers.home-assistant.io/docs/en/device_registry_index.html).
	// Only works through [MQTT discovery](/docs/mqtt/discovery/) and when
	// [`unique_id`](#unique_id) is set. At least one of identifiers or connections
	// must be present to identify the device.
	Device struct {
		// A list of connections of the device to the outside world as a list of tuples
		// `[connection_type, connection_identifier]`. For example the MAC address of a
		// network interface: `\"connections\": [[\"mac\", \"02:5b:26:a8:dc:12\"]]`.
		Connections []string `json:"connections,omitempty"`
		// A list of IDs that uniquely identify the device. For example a serial number.
		Identifiers []string `json:"identifiers,omitempty"`
		// The manufacturer of the device.
		Manufacturer string `json:"manufacturer,omitempty"`
		// The model of the device.
		Model string `json:"model,omitempty"`
		// The name of the device.
		Name string `json:"name,omitempty"`
		// Suggest an area if the device isn’t in one yet.
		SuggestedArea string `json:"suggested_area,omitempty"`
		// The firmware version of the device.
		SwVersion string `json:"sw_version,omitempty"`
		// Identifier of a device that routes messages between this device and Home
		// Assistant. Examples of such devices are hubs, or parent devices of a sub-device.
		// This is used to show device topology in Home Assistant.
		ViaDevice string `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity.
	Icon string `json:"icon,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the JSON dictionary from messages received on the
	// `json_attributes_topic`. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration)
	// documentation.
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as
	// sensor attributes. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration)
	// documentation.
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	// The name of the fan.
	Name string `json:"name,omitempty"`
	// Flag that defines if fan works in optimistic mode
	Optimistic bool `json:"optimistic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to generate the payload to send to `oscillation_command_topic`.
	OscillationCommandTemplate string `json:"oscillation_command_template,omitempty"`
	// The MQTT topic to publish commands to change the oscillation state.
	OscillationCommandTopic string `json:"oscillation_command_topic,omitempty"`
	// The MQTT topic subscribed to receive oscillation state updates.
	OscillationStateTopic string `json:"oscillation_state_topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract a value from the oscillation.
	OscillationValueTemplate string `json:"oscillation_value_template,omitempty"`
	// The payload that represents the available state.
	PayloadAvailable string `json:"payload_available,omitempty"`
	// The payload that represents the unavailable state.
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	// The payload that represents the stop state.
	PayloadOff string `json:"payload_off,omitempty"`
	// The payload that represents the running state.
	PayloadOn string `json:"payload_on,omitempty"`
	// The payload that represents the oscillation off state.
	PayloadOscillationOff string `json:"payload_oscillation_off,omitempty"`
	// The payload that represents the oscillation on state.
	PayloadOscillationOn string `json:"payload_oscillation_on,omitempty"`
	// A special payload that resets the `percentage` state attribute to `None` when
	// received at the `percentage_state_topic`.
	PayloadResetPercentage string `json:"payload_reset_percentage,omitempty"`
	// A special payload that resets the `preset_mode` state attribute to `None` when
	// received at the `preset_mode_state_topic`.
	PayloadResetPresetMode string `json:"payload_reset_preset_mode,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to generate the payload to send to `percentage_command_topic`.
	PercentageCommandTemplate string `json:"percentage_command_template,omitempty"`
	// The MQTT topic to publish commands to change the fan speed state based on a
	// percentage.
	PercentageCommandTopic string `json:"percentage_command_topic,omitempty"`
	// The MQTT topic subscribed to receive fan speed based on percentage.
	PercentageStateTopic string `json:"percentage_state_topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the `percentage` value from the payload received on
	// `percentage_state_topic`.
	PercentageValueTemplate string `json:"percentage_value_template,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to generate the payload to send to `preset_mode_command_topic`.
	PresetModeCommandTemplate string `json:"preset_mode_command_template,omitempty"`
	// The MQTT topic to publish commands to change the preset mode.
	PresetModeCommandTopic string `json:"preset_mode_command_topic,omitempty"`
	// The MQTT topic subscribed to receive fan speed based on presets.
	PresetModeStateTopic string `json:"preset_mode_state_topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the `preset_mode` value from the payload received on
	// `preset_mode_state_topic`.
	PresetModeValueTemplate string `json:"preset_mode_value_template,omitempty"`
	// List of preset modes this fan is capable of running at. Common examples include
	// `auto`, `smart`, `whoosh`, `eco` and `breeze`.
	PresetModes []string `json:"preset_modes,omitempty"`
	// The maximum QoS level of the state topic.
	Qos int `json:"qos,omitempty"`
	// If the published message should have the retain flag on or not.
	Retain bool `json:"retain,omitempty"`
	// The maximum of numeric output range (representing 100 %). The number of speeds
	// within the `speed_range` / `100` will determine the `percentage_step`.
	SpeedRangeMax int `json:"speed_range_max,omitempty"`
	// The minimum of numeric output range (`off` not included, so `speed_range_min` -
	// `1` represents 0 %). The number of speeds within the speed_range / 100 will
	// determine the `percentage_step`.
	SpeedRangeMin int `json:"speed_range_min,omitempty"`
	// The MQTT topic subscribed to receive state updates.
	StateTopic string `json:"state_topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract a value from the state.
	StateValueTemplate string `json:"state_value_template,omitempty"`
	// An ID that uniquely identifies this fan. If two fans have the same unique ID,
	// Home Assistant will raise an exception.
	UniqueId string `json:"unique_id,omitempty"`
}

type HADeviceFanFunctions struct {
	Availability struct {
		State func() string
	}
	Command            func(mqtt.Message, mqtt.Client)
	OscillationCommand func(mqtt.Message, mqtt.Client)
	OscillationState   func() string
	PercentageCommand  func(mqtt.Message, mqtt.Client)
	PercentageState    func() string
	PresetModeCommand  func(mqtt.Message, mqtt.Client)
	PresetModeState    func() string
	State              func() string
}

type HADeviceFanFunctionsConfig struct {
	Availability struct {
		State []string `json:"state,omitempty"`
	} `json:"availability,omitempty"`
	Command            []string `json:"command,omitempty"`
	OscillationCommand []string `json:"oscillation-command,omitempty"`
	OscillationState   []string `json:"oscillation-state,omitempty"`
	PercentageCommand  []string `json:"percentage-command,omitempty"`
	PercentageState    []string `json:"percentage-state,omitempty"`
	PresetModeCommand  []string `json:"preset-mode-command,omitempty"`
	PresetModeState    []string `json:"preset-mode-state,omitempty"`
	State              []string `json:"state,omitempty"`
}

// humidifier
type HADeviceHumidifier struct {
	// A list of MQTT topics subscribed to receive availability (online/offline)
	// updates. Must not be used together with `availability_topic`.
	Availability struct {
		// The payload that represents the available state.
		PayloadAvailable string `json:"payload_available,omitempty"`
		// The payload that represents the unavailable state.
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		// An MQTT topic subscribed to receive availability (online/offline) updates.
		Topic string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	// When `availability` is configured, this controls the conditions needed to set
	// the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set
	// to `all`, `payload_available` must be received on all configured availability
	// topics before the entity is marked as online. If set to `any`,
	// `payload_available` must be received on at least one configured availability
	// topic before the entity is marked as online. If set to `latest`, the last
	// `payload_available` or `payload_not_available` received on any configured
	// availability topic controls the availability.
	AvailabilityMode string `json:"availability_mode,omitempty"`
	// The MQTT topic subscribed to receive availability (online/offline) updates.
	// Must not be used together with `availability`.
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to generate the payload to send to `command_topic`.
	CommandTemplate string `json:"command_template,omitempty"`
	// The MQTT topic to publish commands to change the humidifier state.
	CommandTopic string `json:"command_topic,omitempty"`
	// Information about the device this humidifier is a part of to tie it into the
	// [device
	// registry](https://developers.home-assistant.io/docs/en/device_registry_index.html).
	// Only works through [MQTT discovery](/docs/mqtt/discovery/) and when
	// [`unique_id`](#unique_id) is set. At least one of identifiers or connections
	// must be present to identify the device.
	Device struct {
		// A list of connections of the device to the outside world as a list of tuples
		// `[connection_type, connection_identifier]`. For example the MAC address of a
		// network interface: `\"connections\": [[\"mac\", \"02:5b:26:a8:dc:12\"]]`.
		Connections []string `json:"connections,omitempty"`
		// A list of IDs that uniquely identify the device. For example a serial number.
		Identifiers []string `json:"identifiers,omitempty"`
		// The manufacturer of the device.
		Manufacturer string `json:"manufacturer,omitempty"`
		// The model of the device.
		Model string `json:"model,omitempty"`
		// The name of the device.
		Name string `json:"name,omitempty"`
		// Suggest an area if the device isn’t in one yet.
		SuggestedArea string `json:"suggested_area,omitempty"`
		// The firmware version of the device.
		SwVersion string `json:"sw_version,omitempty"`
		// Identifier of a device that routes messages between this device and Home
		// Assistant. Examples of such devices are hubs, or parent devices of a sub-device.
		// This is used to show device topology in Home Assistant.
		ViaDevice string `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	// The device class of the MQTT device. Must be either `humidifier` or
	// `dehumidifier`.
	DeviceClass string `json:"device_class,omitempty"`
	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity.
	Icon string `json:"icon,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the JSON dictionary from messages received on the
	// `json_attributes_topic`. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration)
	// documentation.
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as
	// sensor attributes. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration)
	// documentation.
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	// The minimum target humidity percentage that can be set.
	MaxHumidity int `json:"max_humidity,omitempty"`
	// The maximum target humidity percentage that can be set.
	MinHumidity int `json:"min_humidity,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to generate the payload to send to `mode_command_topic`.
	ModeCommandTemplate string `json:"mode_command_template,omitempty"`
	// The MQTT topic to publish commands to change the `mode` on the humidifier. This
	// attribute ust be configured together with the `modes` attribute.
	ModeCommandTopic string `json:"mode_command_topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract a value for the humidifier `mode` state.
	ModeStateTemplate string `json:"mode_state_template,omitempty"`
	// The MQTT topic subscribed to receive the humidifier `mode`.
	ModeStateTopic string `json:"mode_state_topic,omitempty"`
	// List of available modes this humidifier is capable of running at. Common
	// examples include `normal`, `eco`, `away`, `boost`, `comfort`, `home`, `sleep`,
	// `auto` and `baby`. These examples offer built-in translations but other custom
	// modes are allowed as well. This attribute ust be configured together with the
	// `mode_command_topic` attribute.
	Modes []string `json:"modes,omitempty"`
	// The name of the humidifier.
	Name string `json:"name,omitempty"`
	// Flag that defines if humidifier works in optimistic mode
	Optimistic bool `json:"optimistic,omitempty"`
	// The payload that represents the available state.
	PayloadAvailable string `json:"payload_available,omitempty"`
	// The payload that represents the unavailable state.
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	// The payload that represents the stop state.
	PayloadOff string `json:"payload_off,omitempty"`
	// The payload that represents the running state.
	PayloadOn string `json:"payload_on,omitempty"`
	// A special payload that resets the `target_humidity` state attribute to `None`
	// when received at the `target_humidity_state_topic`.
	PayloadResetHumidity string `json:"payload_reset_humidity,omitempty"`
	// A special payload that resets the `mode` state attribute to `None` when
	// received at the `mode_state_topic`.
	PayloadResetMode string `json:"payload_reset_mode,omitempty"`
	// The maximum QoS level of the state topic.
	Qos int `json:"qos,omitempty"`
	// If the published message should have the retain flag on or not.
	Retain bool `json:"retain,omitempty"`
	// The MQTT topic subscribed to receive state updates.
	StateTopic string `json:"state_topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract a value from the state.
	StateValueTemplate string `json:"state_value_template,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to generate the payload to send to `target_humidity_command_topic`.
	TargetHumidityCommandTemplate string `json:"target_humidity_command_template,omitempty"`
	// The MQTT topic to publish commands to change the humidifier target humidity
	// state based on a percentage.
	TargetHumidityCommandTopic string `json:"target_humidity_command_topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract a value for the humidifier `target_humidity` state.
	TargetHumidityStateTemplate string `json:"target_humidity_state_template,omitempty"`
	// The MQTT topic subscribed to receive humidifier target humidity.
	TargetHumidityStateTopic string `json:"target_humidity_state_topic,omitempty"`
	// An ID that uniquely identifies this humidifier. If two humidifiers have the
	// same unique ID, Home Assistant will raise an exception.
	UniqueId string `json:"unique_id,omitempty"`
}

type HADeviceHumidifierFunctions struct {
	Availability struct {
		State func() string
	}
	Command               func(mqtt.Message, mqtt.Client)
	ModeCommand           func(mqtt.Message, mqtt.Client)
	ModeState             func() string
	State                 func() string
	TargetHumidityCommand func(mqtt.Message, mqtt.Client)
	TargetHumidityState   func() string
}

type HADeviceHumidifierFunctionsConfig struct {
	Availability struct {
		State []string `json:"state,omitempty"`
	} `json:"availability,omitempty"`
	Command               []string `json:"command,omitempty"`
	ModeCommand           []string `json:"mode-command,omitempty"`
	ModeState             []string `json:"mode-state,omitempty"`
	State                 []string `json:"state,omitempty"`
	TargetHumidityCommand []string `json:"target-humidity-command,omitempty"`
	TargetHumidityState   []string `json:"target-humidity-state,omitempty"`
}

// climate
type HADeviceClimate struct {
	// A template to render the value received on the `action_topic` with.
	ActionTemplate string `json:"action_template,omitempty"`
	// The MQTT topic to subscribe for changes of the current action. If this is set,
	// the climate graph uses the value received as data source. Valid values: `off`,
	// `heating`, `cooling`, `drying`, `idle`, `fan`.
	ActionTopic string `json:"action_topic,omitempty"`
	// The MQTT topic to publish commands to switch auxiliary heat.
	AuxCommandTopic string `json:"aux_command_topic,omitempty"`
	// A template to render the value received on the `aux_state_topic` with.
	AuxStateTemplate string `json:"aux_state_template,omitempty"`
	// The MQTT topic to subscribe for changes of the auxiliary heat mode. If this is
	// not set, the auxiliary heat mode works in optimistic mode (see below).
	AuxStateTopic string `json:"aux_state_topic,omitempty"`
	// A list of MQTT topics subscribed to receive availability (online/offline)
	// updates. Must not be used together with `availability_topic`.
	Availability struct {
		// The payload that represents the available state.
		PayloadAvailable string `json:"payload_available,omitempty"`
		// The payload that represents the unavailable state.
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		// An MQTT topic subscribed to receive availability (online/offline) updates.
		Topic string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	// When `availability` is configured, this controls the conditions needed to set
	// the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set
	// to `all`, `payload_available` must be received on all configured availability
	// topics before the entity is marked as online. If set to `any`,
	// `payload_available` must be received on at least one configured availability
	// topic before the entity is marked as online. If set to `latest`, the last
	// `payload_available` or `payload_not_available` received on any configured
	// availability topic controls the availability.
	AvailabilityMode string `json:"availability_mode,omitempty"`
	// The MQTT topic subscribed to receive availability (online/offline) updates.
	// Must not be used together with `availability`.
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	// The MQTT topic to publish commands to change the away mode.
	AwayModeCommandTopic string `json:"away_mode_command_topic,omitempty"`
	// A template to render the value received on the `away_mode_state_topic` with.
	AwayModeStateTemplate string `json:"away_mode_state_template,omitempty"`
	// The MQTT topic to subscribe for changes of the HVAC away mode. If this is not
	// set, the away mode works in optimistic mode (see below).
	AwayModeStateTopic string `json:"away_mode_state_topic,omitempty"`
	// A template with which the value received on `current_temperature_topic` will be
	// rendered.
	CurrentTemperatureTemplate string `json:"current_temperature_template,omitempty"`
	// The MQTT topic on which to listen for the current temperature.
	CurrentTemperatureTopic string `json:"current_temperature_topic,omitempty"`
	// Information about the device this HVAC device is a part of to tie it into the
	// [device
	// registry](https://developers.home-assistant.io/docs/en/device_registry_index.html).
	// Only works through [MQTT discovery](/docs/mqtt/discovery/) and when
	// [`unique_id`](#unique_id) is set. At least one of identifiers or connections
	// must be present to identify the device.
	Device struct {
		// A list of connections of the device to the outside world as a list of tuples
		// `[connection_type, connection_identifier]`. For example the MAC address of a
		// network interface: `\"connections\": [[\"mac\", \"02:5b:26:a8:dc:12\"]]`.
		Connections []string `json:"connections,omitempty"`
		// A list of IDs that uniquely identify the device. For example a serial number.
		Identifiers []string `json:"identifiers,omitempty"`
		// The manufacturer of the device.
		Manufacturer string `json:"manufacturer,omitempty"`
		// The model of the device.
		Model string `json:"model,omitempty"`
		// The name of the device.
		Name string `json:"name,omitempty"`
		// Suggest an area if the device isn’t in one yet.
		SuggestedArea string `json:"suggested_area,omitempty"`
		// The firmware version of the device.
		SwVersion string `json:"sw_version,omitempty"`
		// Identifier of a device that routes messages between this device and Home
		// Assistant. Examples of such devices are hubs, or parent devices of a sub-device.
		// This is used to show device topology in Home Assistant.
		ViaDevice string `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	// A template to render the value sent to the `fan_mode_command_topic` with.
	FanModeCommandTemplate string `json:"fan_mode_command_template,omitempty"`
	// The MQTT topic to publish commands to change the fan mode.
	FanModeCommandTopic string `json:"fan_mode_command_topic,omitempty"`
	// A template to render the value received on the `fan_mode_state_topic` with.
	FanModeStateTemplate string `json:"fan_mode_state_template,omitempty"`
	// The MQTT topic to subscribe for changes of the HVAC fan mode. If this is not
	// set, the fan mode works in optimistic mode (see below).
	FanModeStateTopic string `json:"fan_mode_state_topic,omitempty"`
	// A list of supported fan modes.
	FanModes []string `json:"fan_modes,omitempty"`
	// A template to render the value sent to the `hold_command_topic` with.
	HoldCommandTemplate string `json:"hold_command_template,omitempty"`
	// The MQTT topic to publish commands to change the hold mode.
	HoldCommandTopic string `json:"hold_command_topic,omitempty"`
	// A list of available hold modes.
	HoldModes []string `json:"hold_modes,omitempty"`
	// A template to render the value received on the `hold_state_topic` with.
	HoldStateTemplate string `json:"hold_state_template,omitempty"`
	// The MQTT topic to subscribe for changes of the HVAC hold mode. If this is not
	// set, the hold mode works in optimistic mode (see below).
	HoldStateTopic string `json:"hold_state_topic,omitempty"`
	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity.
	Icon string `json:"icon,omitempty"`
	// Set the initial target temperature.
	Initial int `json:"initial,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the JSON dictionary from messages received on the
	// `json_attributes_topic`. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration)
	// documentation.
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as
	// sensor attributes. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration)
	// documentation.
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	// Maximum set point available.
	MaxTemp float64 `json:"max_temp,omitempty"`
	// Minimum set point available.
	MinTemp float64 `json:"min_temp,omitempty"`
	// A template to render the value sent to the `mode_command_topic` with.
	ModeCommandTemplate string `json:"mode_command_template,omitempty"`
	// The MQTT topic to publish commands to change the HVAC operation mode.
	ModeCommandTopic string `json:"mode_command_topic,omitempty"`
	// A template to render the value received on the `mode_state_topic` with.
	ModeStateTemplate string `json:"mode_state_template,omitempty"`
	// The MQTT topic to subscribe for changes of the HVAC operation mode. If this is
	// not set, the operation mode works in optimistic mode (see below).
	ModeStateTopic string `json:"mode_state_topic,omitempty"`
	// A list of supported modes. Needs to be a subset of the default values.
	Modes []string `json:"modes,omitempty"`
	// The name of the HVAC.
	Name string `json:"name,omitempty"`
	// The payload that represents the available state.
	PayloadAvailable string `json:"payload_available,omitempty"`
	// The payload that represents the unavailable state.
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	// The payload that represents disabled state.
	PayloadOff string `json:"payload_off,omitempty"`
	// The payload that represents enabled state.
	PayloadOn string `json:"payload_on,omitempty"`
	// The MQTT topic to publish commands to change the power state. This is useful if
	// your device has a separate power toggle in addition to mode.
	PowerCommandTopic string `json:"power_command_topic,omitempty"`
	// The desired precision for this device. Can be used to match your actual
	// thermostat's precision. Supported values are `0.1`, `0.5` and `1.0`.
	Precision float64 `json:"precision,omitempty"`
	// The maximum QoS level to be used when receiving and publishing messages.
	Qos int `json:"qos,omitempty"`
	// Defines if published messages should have the retain flag set.
	Retain bool `json:"retain,omitempty"`
	// Set to `false` to suppress sending of all MQTT messages when the current mode
	// is `Off`.
	SendIfOff bool `json:"send_if_off,omitempty"`
	// A template to render the value sent to the `swing_mode_command_topic` with.
	SwingModeCommandTemplate string `json:"swing_mode_command_template,omitempty"`
	// The MQTT topic to publish commands to change the swing mode.
	SwingModeCommandTopic string `json:"swing_mode_command_topic,omitempty"`
	// A template to render the value received on the `swing_mode_state_topic` with.
	SwingModeStateTemplate string `json:"swing_mode_state_template,omitempty"`
	// The MQTT topic to subscribe for changes of the HVAC swing mode. If this is not
	// set, the swing mode works in optimistic mode (see below).
	SwingModeStateTopic string `json:"swing_mode_state_topic,omitempty"`
	// A list of supported swing modes.
	SwingModes []string `json:"swing_modes,omitempty"`
	// Step size for temperature set point.
	TempStep float64 `json:"temp_step,omitempty"`
	// A template to render the value sent to the `temperature_command_topic` with.
	TemperatureCommandTemplate string `json:"temperature_command_template,omitempty"`
	// The MQTT topic to publish commands to change the target temperature.
	TemperatureCommandTopic string `json:"temperature_command_topic,omitempty"`
	// A template to render the value sent to the `temperature_high_command_topic`
	// with.
	TemperatureHighCommandTemplate string `json:"temperature_high_command_template,omitempty"`
	// The MQTT topic to publish commands to change the high target temperature.
	TemperatureHighCommandTopic string `json:"temperature_high_command_topic,omitempty"`
	// A template to render the value received on the `temperature_high_state_topic`
	// with.
	TemperatureHighStateTemplate string `json:"temperature_high_state_template,omitempty"`
	// The MQTT topic to subscribe for changes in the target high temperature. If this
	// is not set, the target high temperature works in optimistic mode (see below).
	TemperatureHighStateTopic string `json:"temperature_high_state_topic,omitempty"`
	// A template to render the value sent to the `temperature_low_command_topic`
	// with.
	TemperatureLowCommandTemplate string `json:"temperature_low_command_template,omitempty"`
	// The MQTT topic to publish commands to change the target low temperature.
	TemperatureLowCommandTopic string `json:"temperature_low_command_topic,omitempty"`
	// A template to render the value received on the `temperature_low_state_topic`
	// with.
	TemperatureLowStateTemplate string `json:"temperature_low_state_template,omitempty"`
	// The MQTT topic to subscribe for changes in the target low temperature. If this
	// is not set, the target low temperature works in optimistic mode (see below).
	TemperatureLowStateTopic string `json:"temperature_low_state_topic,omitempty"`
	// A template to render the value received on the `temperature_state_topic` with.
	TemperatureStateTemplate string `json:"temperature_state_template,omitempty"`
	// The MQTT topic to subscribe for changes in the target temperature. If this is
	// not set, the target temperature works in optimistic mode (see below).
	TemperatureStateTopic string `json:"temperature_state_topic,omitempty"`
	// Defines the temperature unit of the device, `C` or `F`. If this is not set, the
	// temperature unit is set to the system temperature unit.
	TemperatureUnit string `json:"temperature_unit,omitempty"`
	// An ID that uniquely identifies this HVAC device. If two HVAC devices have the
	// same unique ID, Home Assistant will raise an exception.
	UniqueId string `json:"unique_id,omitempty"`
	// Default template to render the payloads on *all* `*_state_topic`s with.
	ValueTemplate string `json:"value_template,omitempty"`
}

type HADeviceClimateFunctions struct {
	AuxCommand   func(mqtt.Message, mqtt.Client)
	AuxState     func() string
	Availability struct {
		State func() string
	}
	AwayModeCommand        func(mqtt.Message, mqtt.Client)
	AwayModeState          func() string
	FanModeCommand         func(mqtt.Message, mqtt.Client)
	FanModeState           func() string
	HoldCommand            func(mqtt.Message, mqtt.Client)
	HoldState              func() string
	ModeCommand            func(mqtt.Message, mqtt.Client)
	ModeState              func() string
	PowerCommand           func(mqtt.Message, mqtt.Client)
	SwingModeCommand       func(mqtt.Message, mqtt.Client)
	SwingModeState         func() string
	TemperatureCommand     func(mqtt.Message, mqtt.Client)
	TemperatureHighCommand func(mqtt.Message, mqtt.Client)
	TemperatureHighState   func() string
	TemperatureLowCommand  func(mqtt.Message, mqtt.Client)
	TemperatureLowState    func() string
	TemperatureState       func() string
}

type HADeviceClimateFunctionsConfig struct {
	AuxCommand   []string `json:"aux-command,omitempty"`
	AuxState     []string `json:"aux-state,omitempty"`
	Availability struct {
		State []string `json:"state,omitempty"`
	} `json:"availability,omitempty"`
	AwayModeCommand        []string `json:"away-mode-command,omitempty"`
	AwayModeState          []string `json:"away-mode-state,omitempty"`
	FanModeCommand         []string `json:"fan-mode-command,omitempty"`
	FanModeState           []string `json:"fan-mode-state,omitempty"`
	HoldCommand            []string `json:"hold-command,omitempty"`
	HoldState              []string `json:"hold-state,omitempty"`
	ModeCommand            []string `json:"mode-command,omitempty"`
	ModeState              []string `json:"mode-state,omitempty"`
	PowerCommand           []string `json:"power-command,omitempty"`
	SwingModeCommand       []string `json:"swing-mode-command,omitempty"`
	SwingModeState         []string `json:"swing-mode-state,omitempty"`
	TemperatureCommand     []string `json:"temperature-command,omitempty"`
	TemperatureHighCommand []string `json:"temperature-high-command,omitempty"`
	TemperatureHighState   []string `json:"temperature-high-state,omitempty"`
	TemperatureLowCommand  []string `json:"temperature-low-command,omitempty"`
	TemperatureLowState    []string `json:"temperature-low-state,omitempty"`
	TemperatureState       []string `json:"temperature-state,omitempty"`
}

// light
type HADeviceLight struct {
	// A list of MQTT topics subscribed to receive availability (online/offline)
	// updates. Must not be used together with `availability_topic`.
	Availability struct {
		// The payload that represents the available state.
		PayloadAvailable string `json:"payload_available,omitempty"`
		// The payload that represents the unavailable state.
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		// An MQTT topic subscribed to receive availability (online/offline) updates.
		Topic string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	// When `availability` is configured, this controls the conditions needed to set
	// the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set
	// to `all`, `payload_available` must be received on all configured availability
	// topics before the entity is marked as online. If set to `any`,
	// `payload_available` must be received on at least one configured availability
	// topic before the entity is marked as online. If set to `latest`, the last
	// `payload_available` or `payload_not_available` received on any configured
	// availability topic controls the availability.
	AvailabilityMode string `json:"availability_mode,omitempty"`
	// The MQTT topic subscribed to receive availability (online/offline) updates.
	// Must not be used together with `availability`.
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	// The MQTT topic to publish commands to change the light’s brightness.
	BrightnessCommandTopic string `json:"brightness_command_topic,omitempty"`
	// Defines the maximum brightness value (i.e., 100%) of the MQTT device.
	BrightnessScale int `json:"brightness_scale,omitempty"`
	// The MQTT topic subscribed to receive brightness state updates.
	BrightnessStateTopic string `json:"brightness_state_topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the brightness value.
	BrightnessValueTemplate string `json:"brightness_value_template,omitempty"`
	// The MQTT topic subscribed to receive color mode updates. If this is not
	// configured, `color_mode` will be automatically set according to the last
	// received valid color or color temperature
	ColorModeStateTopic string `json:"color_mode_state_topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the color mode.
	ColorModeValueTemplate string `json:"color_mode_value_template,omitempty"`
	// Defines a [template](/docs/configuration/templating/) to compose message which
	// will be sent to `color_temp_command_topic`. Available variables: `value`.
	ColorTempCommandTemplate string `json:"color_temp_command_template,omitempty"`
	// The MQTT topic to publish commands to change the light’s color temperature
	// state. The color temperature command slider has a range of 153 to 500 mireds
	// (micro reciprocal degrees).
	ColorTempCommandTopic string `json:"color_temp_command_topic,omitempty"`
	// The MQTT topic subscribed to receive color temperature state updates.
	ColorTempStateTopic string `json:"color_temp_state_topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the color temperature value.
	ColorTempValueTemplate string `json:"color_temp_value_template,omitempty"`
	// The MQTT topic to publish commands to change the switch state.
	CommandTopic string `json:"command_topic,omitempty"`
	// Information about the device this light is a part of to tie it into the [device
	// registry](https://developers.home-assistant.io/docs/en/device_registry_index.html).
	// Only works through [MQTT discovery](/docs/mqtt/discovery/) and when
	// [`unique_id`](#unique_id) is set. At least one of identifiers or connections
	// must be present to identify the device.
	Device struct {
		// A list of connections of the device to the outside world as a list of tuples
		// `[connection_type, connection_identifier]`. For example the MAC address of a
		// network interface: `\"connections\": [[\"mac\", \"02:5b:26:a8:dc:12\"]]`.
		Connections []string `json:"connections,omitempty"`
		// A list of IDs that uniquely identify the device. For example a serial number.
		Identifiers []string `json:"identifiers,omitempty"`
		// The manufacturer of the device.
		Manufacturer string `json:"manufacturer,omitempty"`
		// The model of the device.
		Model string `json:"model,omitempty"`
		// The name of the device.
		Name string `json:"name,omitempty"`
		// Suggest an area if the device isn’t in one yet.
		SuggestedArea string `json:"suggested_area,omitempty"`
		// The firmware version of the device.
		SwVersion string `json:"sw_version,omitempty"`
		// Identifier of a device that routes messages between this device and Home
		// Assistant. Examples of such devices are hubs, or parent devices of a sub-device.
		// This is used to show device topology in Home Assistant.
		ViaDevice string `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	// The MQTT topic to publish commands to change the light's effect state.
	EffectCommandTopic string `json:"effect_command_topic,omitempty"`
	// The list of effects the light supports.
	EffectList []string `json:"effect_list,omitempty"`
	// The MQTT topic subscribed to receive effect state updates.
	EffectStateTopic string `json:"effect_state_topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the effect value.
	EffectValueTemplate string `json:"effect_value_template,omitempty"`
	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	// The MQTT topic to publish commands to change the light's color state in HS
	// format (Hue Saturation). Range for Hue: 0° .. 360°, Range of Saturation:
	// 0..100. Note: Brightness is sent separately in the `brightness_command_topic`.
	HsCommandTopic string `json:"hs_command_topic,omitempty"`
	// The MQTT topic subscribed to receive color state updates in HS format. Note:
	// Brightness is received separately in the `brightness_state_topic`.
	HsStateTopic string `json:"hs_state_topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the HS value.
	HsValueTemplate string `json:"hs_value_template,omitempty"`
	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity.
	Icon string `json:"icon,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the JSON dictionary from messages received on the
	// `json_attributes_topic`. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration)
	// documentation.
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as
	// sensor attributes. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration)
	// documentation.
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	// The maximum color temperature in mireds.
	MaxMireds int `json:"max_mireds,omitempty"`
	// The minimum color temperature in mireds.
	MinMireds int `json:"min_mireds,omitempty"`
	// The name of the light.
	Name string `json:"name,omitempty"`
	// Defines when on the payload_on is sent. Using `last` (the default) will send
	// any style (brightness, color, etc) topics first and then a `payload_on` to the
	// `command_topic`. Using `first` will send the `payload_on` and then any style
	// topics. Using `brightness` will only send brightness commands instead of the
	// `payload_on` to turn the light on.
	OnCommandType string `json:"on_command_type,omitempty"`
	// Flag that defines if switch works in optimistic mode.
	Optimistic bool `json:"optimistic,omitempty"`
	// The payload that represents the available state.
	PayloadAvailable string `json:"payload_available,omitempty"`
	// The payload that represents the unavailable state.
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	// The payload that represents disabled state.
	PayloadOff string `json:"payload_off,omitempty"`
	// The payload that represents enabled state.
	PayloadOn string `json:"payload_on,omitempty"`
	// The maximum QoS level of the state topic.
	Qos int `json:"qos,omitempty"`
	// If the published message should have the retain flag on or not.
	Retain bool `json:"retain,omitempty"`
	// Defines a [template](/docs/configuration/templating/) to compose message which
	// will be sent to `rgb_command_topic`. Available variables: `red`, `green` and
	// `blue`.
	RgbCommandTemplate string `json:"rgb_command_template,omitempty"`
	// The MQTT topic to publish commands to change the light's RGB state. Please note
	// that the color value sent by Home Assistant is normalized to full brightness if
	// `brightness_command_topic` is set. Brightness information is in this case sent
	// separately in the `brightness_command_topic`. This will cause a light that
	// expects an absolute color value (including brightness) to flicker.
	RgbCommandTopic string `json:"rgb_command_topic,omitempty"`
	// The MQTT topic subscribed to receive RGB state updates. The expected payload is
	// the RGB values separated by commas, for example, `255,0,127`. Please note that
	// the color value received by Home Assistant is normalized to full brightness.
	// Brightness information is received separately in the `brightness_state_topic`.
	RgbStateTopic string `json:"rgb_state_topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the RGB value.
	RgbValueTemplate string `json:"rgb_value_template,omitempty"`
	// The schema to use. Must be `default` or omitted to select the default schema\".
	Schema string `json:"schema,omitempty"`
	// The MQTT topic subscribed to receive state updates.
	StateTopic string `json:"state_topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the state value. The template should match the payload `on` and `off`
	// values, so if your light uses `power on` to turn on, your `state_value_template`
	// string should return `power on` when the switch is on. For example if the
	// message is just `on`, your `state_value_template` should be `power {{ value }}`.
	StateValueTemplate string `json:"state_value_template,omitempty"`
	// An ID that uniquely identifies this light. If two lights have the same unique
	// ID, Home Assistant will raise an exception.
	UniqueId string `json:"unique_id,omitempty"`
	// The MQTT topic to publish commands to change the light to white mode with a
	// given brightness.
	WhiteCommandTopic string `json:"white_command_topic,omitempty"`
	// Defines the maximum white level (i.e., 100%) of the MQTT device.
	WhiteScale int `json:"white_scale,omitempty"`
	// The MQTT topic to publish commands to change the light's XY state.
	XyCommandTopic string `json:"xy_command_topic,omitempty"`
	// The MQTT topic subscribed to receive XY state updates.
	XyStateTopic string `json:"xy_state_topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the XY value.
	XyValueTemplate string `json:"xy_value_template,omitempty"`
}

type HADeviceLightFunctions struct {
	Availability struct {
		State func() string
	}
	BrightnessCommand func(mqtt.Message, mqtt.Client)
	BrightnessState   func() string
	ColorModeState    func() string
	ColorTempCommand  func(mqtt.Message, mqtt.Client)
	ColorTempState    func() string
	Command           func(mqtt.Message, mqtt.Client)
	EffectCommand     func(mqtt.Message, mqtt.Client)
	EffectState       func() string
	HsCommand         func(mqtt.Message, mqtt.Client)
	HsState           func() string
	RgbCommand        func(mqtt.Message, mqtt.Client)
	RgbState          func() string
	State             func() string
	WhiteCommand      func(mqtt.Message, mqtt.Client)
	XyCommand         func(mqtt.Message, mqtt.Client)
	XyState           func() string
}

type HADeviceLightFunctionsConfig struct {
	Availability struct {
		State []string `json:"state,omitempty"`
	} `json:"availability,omitempty"`
	BrightnessCommand []string `json:"brightness-command,omitempty"`
	BrightnessState   []string `json:"brightness-state,omitempty"`
	ColorModeState    []string `json:"color-mode-state,omitempty"`
	ColorTempCommand  []string `json:"color-temp-command,omitempty"`
	ColorTempState    []string `json:"color-temp-state,omitempty"`
	Command           []string `json:"command,omitempty"`
	EffectCommand     []string `json:"effect-command,omitempty"`
	EffectState       []string `json:"effect-state,omitempty"`
	HsCommand         []string `json:"hs-command,omitempty"`
	HsState           []string `json:"hs-state,omitempty"`
	RgbCommand        []string `json:"rgb-command,omitempty"`
	RgbState          []string `json:"rgb-state,omitempty"`
	State             []string `json:"state,omitempty"`
	WhiteCommand      []string `json:"white-command,omitempty"`
	XyCommand         []string `json:"xy-command,omitempty"`
	XyState           []string `json:"xy-state,omitempty"`
}

// lock
type HADeviceLock struct {
	// A list of MQTT topics subscribed to receive availability (online/offline)
	// updates. Must not be used together with `availability_topic`.
	Availability struct {
		// The payload that represents the available state.
		PayloadAvailable string `json:"payload_available,omitempty"`
		// The payload that represents the unavailable state.
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		// An MQTT topic subscribed to receive availability (online/offline) updates.
		Topic string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	// When `availability` is configured, this controls the conditions needed to set
	// the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set
	// to `all`, `payload_available` must be received on all configured availability
	// topics before the entity is marked as online. If set to `any`,
	// `payload_available` must be received on at least one configured availability
	// topic before the entity is marked as online. If set to `latest`, the last
	// `payload_available` or `payload_not_available` received on any configured
	// availability topic controls the availability.
	AvailabilityMode string `json:"availability_mode,omitempty"`
	// The MQTT topic subscribed to receive availability (online/offline) updates.
	// Must not be used together with `availability`.
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	// The MQTT topic to publish commands to change the lock state.
	CommandTopic string `json:"command_topic,omitempty"`
	// Information about the device this lock is a part of to tie it into the [device
	// registry](https://developers.home-assistant.io/docs/en/device_registry_index.html).
	// Only works through [MQTT discovery](/docs/mqtt/discovery/) and when
	// [`unique_id`](#unique_id) is set. At least one of identifiers or connections
	// must be present to identify the device.
	Device struct {
		// A list of connections of the device to the outside world as a list of tuples
		// `[connection_type, connection_identifier]`. For example the MAC address of a
		// network interface: `\"connections\": [[\"mac\", \"02:5b:26:a8:dc:12\"]]`.
		Connections []string `json:"connections,omitempty"`
		// A list of IDs that uniquely identify the device. For example a serial number.
		Identifiers []string `json:"identifiers,omitempty"`
		// The manufacturer of the device.
		Manufacturer string `json:"manufacturer,omitempty"`
		// The model of the device.
		Model string `json:"model,omitempty"`
		// The name of the device.
		Name string `json:"name,omitempty"`
		// Suggest an area if the device isn’t in one yet.
		SuggestedArea string `json:"suggested_area,omitempty"`
		// The firmware version of the device.
		SwVersion string `json:"sw_version,omitempty"`
		// Identifier of a device that routes messages between this device and Home
		// Assistant. Examples of such devices are hubs, or parent devices of a sub-device.
		// This is used to show device topology in Home Assistant.
		ViaDevice string `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity.
	Icon string `json:"icon,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the JSON dictionary from messages received on the
	// `json_attributes_topic`. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration)
	// documentation.
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as
	// sensor attributes. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration)
	// documentation.
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	// The name of the lock.
	Name string `json:"name,omitempty"`
	// Flag that defines if lock works in optimistic mode.
	Optimistic bool `json:"optimistic,omitempty"`
	// The payload that represents the available state.
	PayloadAvailable string `json:"payload_available,omitempty"`
	// The payload that represents enabled/locked state.
	PayloadLock string `json:"payload_lock,omitempty"`
	// The payload that represents the unavailable state.
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	// The payload that represents disabled/unlocked state.
	PayloadUnlock string `json:"payload_unlock,omitempty"`
	// The maximum QoS level of the state topic.
	Qos int `json:"qos,omitempty"`
	// If the published message should have the retain flag on or not.
	Retain bool `json:"retain,omitempty"`
	// The value that represents the lock to be in locked state
	StateLocked string `json:"state_locked,omitempty"`
	// The MQTT topic subscribed to receive state updates.
	StateTopic string `json:"state_topic,omitempty"`
	// The value that represents the lock to be in unlocked state
	StateUnlocked string `json:"state_unlocked,omitempty"`
	// An ID that uniquely identifies this lock. If two locks have the same unique ID,
	// Home Assistant will raise an exception.
	UniqueId string `json:"unique_id,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract a value from the payload.
	ValueTemplate string `json:"value_template,omitempty"`
}

type HADeviceLockFunctions struct {
	Availability struct {
		State func() string
	}
	Command func(mqtt.Message, mqtt.Client)
	State   func() string
}

type HADeviceLockFunctionsConfig struct {
	Availability struct {
		State []string `json:"state,omitempty"`
	} `json:"availability,omitempty"`
	Command []string `json:"command,omitempty"`
	State   []string `json:"state,omitempty"`
}

// number
type HADeviceNumber struct {
	// A list of MQTT topics subscribed to receive availability (online/offline)
	// updates. Must not be used together with `availability_topic`.
	Availability struct {
		// The payload that represents the available state.
		PayloadAvailable string `json:"payload_available,omitempty"`
		// The payload that represents the unavailable state.
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		// An MQTT topic subscribed to receive availability (online/offline) updates.
		Topic string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	// When `availability` is configured, this controls the conditions needed to set
	// the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set
	// to `all`, `payload_available` must be received on all configured availability
	// topics before the entity is marked as online. If set to `any`,
	// `payload_available` must be received on at least one configured availability
	// topic before the entity is marked as online. If set to `latest`, the last
	// `payload_available` or `payload_not_available` received on any configured
	// availability topic controls the availability.
	AvailabilityMode string `json:"availability_mode,omitempty"`
	// The MQTT topic subscribed to receive availability (online/offline) updates.
	// Must not be used together with `availability`.
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	// The MQTT topic to publish commands to change the number.
	CommandTopic string `json:"command_topic,omitempty"`
	// Information about the device this Number is a part of to tie it into the
	// [device
	// registry](https://developers.home-assistant.io/docs/en/device_registry_index.html).
	// Only works through [MQTT discovery](/docs/mqtt/discovery/) and when
	// [`unique_id`](#unique_id) is set. At least one of identifiers or connections
	// must be present to identify the device.
	Device struct {
		// A list of connections of the device to the outside world as a list of tuples
		// `[connection_type, connection_identifier]`. For example the MAC address of a
		// network interface: `\"connections\": [\"mac\", \"02:5b:26:a8:dc:12\"]`.
		Connections []string `json:"connections,omitempty"`
		// A list of IDs that uniquely identify the device. For example a serial number.
		Identifiers []string `json:"identifiers,omitempty"`
		// The manufacturer of the device.
		Manufacturer string `json:"manufacturer,omitempty"`
		// The model of the device.
		Model string `json:"model,omitempty"`
		// The name of the device.
		Name string `json:"name,omitempty"`
		// Suggest an area if the device isn’t in one yet.
		SuggestedArea string `json:"suggested_area,omitempty"`
		// The firmware version of the device.
		SwVersion string `json:"sw_version,omitempty"`
		// Identifier of a device that routes messages between this device and Home
		// Assistant. Examples of such devices are hubs, or parent devices of a sub-device.
		// This is used to show device topology in Home Assistant.
		ViaDevice string `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity.
	Icon string `json:"icon,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the JSON dictionary from messages received on the
	// `json_attributes_topic`.
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as
	// number attributes. Implies `force_update` of the current number state when a
	// message is received on this topic.
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	// Maximum value.
	Max float64 `json:"max,omitempty"`
	// Minimum value.
	Min float64 `json:"min,omitempty"`
	// The name of the Number.
	Name string `json:"name,omitempty"`
	// Flag that defines if number works in optimistic mode.
	Optimistic bool `json:"optimistic,omitempty"`
	// The maximum QoS level of the state topic. Default is 0 and will also be used to
	// publishing messages.
	Qos int `json:"qos,omitempty"`
	// If the published message should have the retain flag on or not.
	Retain bool `json:"retain,omitempty"`
	// The MQTT topic subscribed to receive number values.
	StateTopic string `json:"state_topic,omitempty"`
	// Step value. Smallest value `0.001`.
	Step float64 `json:"step,omitempty"`
	// An ID that uniquely identifies this Number. If two Numbers have the same unique
	// ID Home Assistant will raise an exception.
	UniqueId string `json:"unique_id,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the value.
	ValueTemplate string `json:"value_template,omitempty"`
}

type HADeviceNumberFunctions struct {
	Availability struct {
		State func() string
	}
	Command func(mqtt.Message, mqtt.Client)
	State   func() string
}

type HADeviceNumberFunctionsConfig struct {
	Availability struct {
		State []string `json:"state,omitempty"`
	} `json:"availability,omitempty"`
	Command []string `json:"command,omitempty"`
	State   []string `json:"state,omitempty"`
}

// scene
type HADeviceScene struct {
	// A list of MQTT topics subscribed to receive availability (online/offline)
	// updates. Must not be used together with `availability_topic`.
	Availability struct {
		// The payload that represents the available state.
		PayloadAvailable string `json:"payload_available,omitempty"`
		// The payload that represents the unavailable state.
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		// An MQTT topic subscribed to receive availability (online/offline) updates.
		Topic string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	// When `availability` is configured, this controls the conditions needed to set
	// the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set
	// to `all`, `payload_available` must be received on all configured availability
	// topics before the entity is marked as online. If set to `any`,
	// `payload_available` must be received on at least one configured availability
	// topic before the entity is marked as online. If set to `latest`, the last
	// `payload_available` or `payload_not_available` received on any configured
	// availability topic controls the availability.
	AvailabilityMode string `json:"availability_mode,omitempty"`
	// The MQTT topic subscribed to receive availability (online/offline) updates.
	// Must not be used together with `availability`.
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	// The MQTT topic to publish commands to change the scene state.
	CommandTopic string `json:"command_topic,omitempty"`
	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	// Icon for the scene.
	Icon string `json:"icon,omitempty"`
	// The name to use when displaying this scene.
	Name string `json:"name,omitempty"`
	// The payload that represents the available state.
	PayloadAvailable string `json:"payload_available,omitempty"`
	// The payload that represents the unavailable state.
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	// The payload that represents `on` state. If specified, will be used for both
	// comparing to the value in the `state_topic` (see `value_template` and `state_on`
	// for details) and sending as `on` command to the `command_topic`.
	PayloadOn string `json:"payload_on,omitempty"`
	// The maximum QoS level of the state topic. Default is 0 and will also be used to
	// publishing messages.
	Qos int `json:"qos,omitempty"`
	// If the published message should have the retain flag on or not.
	Retain bool `json:"retain,omitempty"`
	// An ID that uniquely identifies this scene entity. If two scenes have the same
	// unique ID, Home Assistant will raise an exception.
	UniqueId string `json:"unique_id,omitempty"`
}

type HADeviceSceneFunctions struct {
	Availability struct {
		State func() string
	}
	Command func(mqtt.Message, mqtt.Client)
}

type HADeviceSceneFunctionsConfig struct {
	Availability struct {
		State []string `json:"state,omitempty"`
	} `json:"availability,omitempty"`
	Command []string `json:"command,omitempty"`
}

// select
type HADeviceSelect struct {
	// A list of MQTT topics subscribed to receive availability (online/offline)
	// updates. Must not be used together with `availability_topic`.
	Availability struct {
		// The payload that represents the available state.
		PayloadAvailable string `json:"payload_available,omitempty"`
		// The payload that represents the unavailable state.
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		// An MQTT topic subscribed to receive availability (online/offline) updates.
		Topic string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	// When `availability` is configured, this controls the conditions needed to set
	// the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set
	// to `all`, `payload_available` must be received on all configured availability
	// topics before the entity is marked as online. If set to `any`,
	// `payload_available` must be received on at least one configured availability
	// topic before the entity is marked as online. If set to `latest`, the last
	// `payload_available` or `payload_not_available` received on any configured
	// availability topic controls the availability.
	AvailabilityMode string `json:"availability_mode,omitempty"`
	// The MQTT topic subscribed to receive availability (online/offline) updates.
	// Must not be used together with `availability`.
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	// The MQTT topic to publish commands to change the selected option.
	CommandTopic string `json:"command_topic,omitempty"`
	// Information about the device this Select is a part of to tie it into the
	// [device
	// registry](https://developers.home-assistant.io/docs/en/device_registry_index.html).
	// Only works through [MQTT discovery](/docs/mqtt/discovery/) and when
	// [`unique_id`](#unique_id) is set. At least one of identifiers or connections
	// must be present to identify the device.
	Device struct {
		// A list of connections of the device to the outside world as a list of tuples
		// `[connection_type, connection_identifier]`. For example the MAC address of a
		// network interface: `\"connections\": [\"mac\", \"02:5b:26:a8:dc:12\"]`.
		Connections []string `json:"connections,omitempty"`
		// A list of IDs that uniquely identify the device. For example a serial number.
		Identifiers []string `json:"identifiers,omitempty"`
		// The manufacturer of the device.
		Manufacturer string `json:"manufacturer,omitempty"`
		// The model of the device.
		Model string `json:"model,omitempty"`
		// The name of the device.
		Name string `json:"name,omitempty"`
		// Suggest an area if the device isn’t in one yet.
		SuggestedArea string `json:"suggested_area,omitempty"`
		// The firmware version of the device.
		SwVersion string `json:"sw_version,omitempty"`
		// Identifier of a device that routes messages between this device and Home
		// Assistant. Examples of such devices are hubs, or parent devices of a sub-device.
		// This is used to show device topology in Home Assistant.
		ViaDevice string `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity.
	Icon string `json:"icon,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the JSON dictionary from messages received on the
	// `json_attributes_topic`.
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as
	// entity attributes. Implies `force_update` of the current select state when a
	// message is received on this topic.
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	// The name of the Select.
	Name string `json:"name,omitempty"`
	// Flag that defines if the select works in optimistic mode.
	Optimistic bool `json:"optimistic,omitempty"`
	// List of options that can be selected.
	Options []string `json:"options,omitempty"`
	// The maximum QoS level of the state topic. Default is 0 and will also be used to
	// publishing messages.
	Qos int `json:"qos,omitempty"`
	// If the published message should have the retain flag on or not.
	Retain bool `json:"retain,omitempty"`
	// The MQTT topic subscribed to receive update of the selected option.
	StateTopic string `json:"state_topic,omitempty"`
	// An ID that uniquely identifies this Select. If two Selects have the same unique
	// ID Home Assistant will raise an exception.
	UniqueId string `json:"unique_id,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the value.
	ValueTemplate string `json:"value_template,omitempty"`
}

type HADeviceSelectFunctions struct {
	Availability struct {
		State func() string
	}
	Command func(mqtt.Message, mqtt.Client)
	State   func() string
}

type HADeviceSelectFunctionsConfig struct {
	Availability struct {
		State []string `json:"state,omitempty"`
	} `json:"availability,omitempty"`
	Command []string `json:"command,omitempty"`
	State   []string `json:"state,omitempty"`
}

// sensor
type HADeviceSensor struct {
	// A list of MQTT topics subscribed to receive availability (online/offline)
	// updates. Must not be used together with `availability_topic`.
	Availability struct {
		// The payload that represents the available state.
		PayloadAvailable string `json:"payload_available,omitempty"`
		// The payload that represents the unavailable state.
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		// An MQTT topic subscribed to receive availability (online/offline) updates.
		Topic string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	// When `availability` is configured, this controls the conditions needed to set
	// the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set
	// to `all`, `payload_available` must be received on all configured availability
	// topics before the entity is marked as online. If set to `any`,
	// `payload_available` must be received on at least one configured availability
	// topic before the entity is marked as online. If set to `latest`, the last
	// `payload_available` or `payload_not_available` received on any configured
	// availability topic controls the availability.
	AvailabilityMode string `json:"availability_mode,omitempty"`
	// The MQTT topic subscribed to receive availability (online/offline) updates.
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	// Information about the device this sensor is a part of to tie it into the
	// [device
	// registry](https://developers.home-assistant.io/docs/en/device_registry_index.html).
	// Only works through [MQTT discovery](/docs/mqtt/discovery/) and when
	// [`unique_id`](#unique_id) is set. At least one of identifiers or connections
	// must be present to identify the device.
	Device struct {
		// A list of connections of the device to the outside world as a list of tuples
		// `[connection_type, connection_identifier]`. For example the MAC address of a
		// network interface: `\"connections\": [[\"mac\", \"02:5b:26:a8:dc:12\"]]`.
		Connections []string `json:"connections,omitempty"`
		// A list of IDs that uniquely identify the device. For example a serial number.
		Identifiers []string `json:"identifiers,omitempty"`
		// The manufacturer of the device.
		Manufacturer string `json:"manufacturer,omitempty"`
		// The model of the device.
		Model string `json:"model,omitempty"`
		// The name of the device.
		Name string `json:"name,omitempty"`
		// Suggest an area if the device isn’t in one yet.
		SuggestedArea string `json:"suggested_area,omitempty"`
		// The firmware version of the device.
		SwVersion string `json:"sw_version,omitempty"`
		// Identifier of a device that routes messages between this device and Home
		// Assistant. Examples of such devices are hubs, or parent devices of a sub-device.
		// This is used to show device topology in Home Assistant.
		ViaDevice string `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	// The [type/class](/integrations/sensor/#device-class) of the sensor to set the
	// icon in the frontend.
	DeviceClass string `json:"device_class,omitempty"`
	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	// Defines the number of seconds after the sensor's state expires, if it's not
	// updated. After expiry, the sensor's state becomes `unavailable`.
	ExpireAfter int `json:"expire_after,omitempty"`
	// Sends update events even if the value hasn't changed. Useful if you want to
	// have meaningful value graphs in history.
	ForceUpdate bool `json:"force_update,omitempty"`
	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity.
	Icon string `json:"icon,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the JSON dictionary from messages received on the
	// `json_attributes_topic`.
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as
	// sensor attributes. Implies `force_update` of the current sensor state when a
	// message is received on this topic.
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	// The MQTT topic subscribed to receive timestamps for when an accumulating sensor
	// such as an energy meter was reset. If the sensor never resets, set
	// `last_reset_topic` to same as `state_topic` and set the
	// `last_reset_value_template` to a constant valid timstamp, for example UNIX epoch
	// 0: `1970-01-01T00:00:00+00:00`.
	LastResetTopic string `json:"last_reset_topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the last_reset. Available variables: `entity_id`. The `entity_id` can
	// be used to reference the entity's attributes.
	LastResetValueTemplate string `json:"last_reset_value_template,omitempty"`
	// The name of the MQTT sensor.
	Name string `json:"name,omitempty"`
	// The payload that represents the available state.
	PayloadAvailable string `json:"payload_available,omitempty"`
	// The payload that represents the unavailable state.
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	// The maximum QoS level of the state topic.
	Qos int `json:"qos,omitempty"`
	// The
	// [state_class](https://developers.home-assistant.io/docs/core/entity/sensor#available-state-classes)
	// of the sensor.
	StateClass string `json:"state_class,omitempty"`
	// The MQTT topic subscribed to receive sensor values.
	StateTopic string `json:"state_topic,omitempty"`
	// An ID that uniquely identifies this sensor. If two sensors have the same unique
	// ID, Home Assistant will raise an exception.
	UniqueId string `json:"unique_id,omitempty"`
	// Defines the units of measurement of the sensor, if any.
	UnitOfMeasurement string `json:"unit_of_measurement,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the value. Available variables: `entity_id`. The `entity_id` can be
	// used to reference the entity's attributes.
	ValueTemplate string `json:"value_template,omitempty"`
}

type HADeviceSensorFunctions struct {
	Availability struct {
		State func() string
	}
	State func() string
}

type HADeviceSensorFunctionsConfig struct {
	Availability struct {
		State []string `json:"state,omitempty"`
	} `json:"availability,omitempty"`
	State []string `json:"state,omitempty"`
}

// switch
type HADeviceSwitch struct {
	// A list of MQTT topics subscribed to receive availability (online/offline)
	// updates. Must not be used together with `availability_topic`.
	Availability struct {
		// The payload that represents the available state.
		PayloadAvailable string `json:"payload_available,omitempty"`
		// The payload that represents the unavailable state.
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		// An MQTT topic subscribed to receive availability (online/offline) updates.
		Topic string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	// When `availability` is configured, this controls the conditions needed to set
	// the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set
	// to `all`, `payload_available` must be received on all configured availability
	// topics before the entity is marked as online. If set to `any`,
	// `payload_available` must be received on at least one configured availability
	// topic before the entity is marked as online. If set to `latest`, the last
	// `payload_available` or `payload_not_available` received on any configured
	// availability topic controls the availability.
	AvailabilityMode string `json:"availability_mode,omitempty"`
	// The MQTT topic subscribed to receive availability (online/offline) updates.
	// Must not be used together with `availability`.
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	// The MQTT topic to publish commands to change the switch state.
	CommandTopic string `json:"command_topic,omitempty"`
	// Information about the device this switch is a part of to tie it into the
	// [device
	// registry](https://developers.home-assistant.io/docs/en/device_registry_index.html).
	// Only works through [MQTT discovery](/docs/mqtt/discovery/) and when
	// [`unique_id`](#unique_id) is set. At least one of identifiers or connections
	// must be present to identify the device.
	Device struct {
		// A list of connections of the device to the outside world as a list of tuples
		// `[connection_type, connection_identifier]`. For example the MAC address of a
		// network interface: `\"connections\": [[\"mac\", \"02:5b:26:a8:dc:12\"]]`.
		Connections []string `json:"connections,omitempty"`
		// A list of IDs that uniquely identify the device. For example a serial number.
		Identifiers []string `json:"identifiers,omitempty"`
		// The manufacturer of the device.
		Manufacturer string `json:"manufacturer,omitempty"`
		// The model of the device.
		Model string `json:"model,omitempty"`
		// The name of the device.
		Name string `json:"name,omitempty"`
		// Suggest an area if the device isn’t in one yet.
		SuggestedArea string `json:"suggested_area,omitempty"`
		// The firmware version of the device.
		SwVersion string `json:"sw_version,omitempty"`
		// Identifier of a device that routes messages between this device and Home
		// Assistant. Examples of such devices are hubs, or parent devices of a sub-device.
		// This is used to show device topology in Home Assistant.
		ViaDevice string `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity.
	Icon string `json:"icon,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the JSON dictionary from messages received on the
	// `json_attributes_topic`. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration)
	// documentation.
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as
	// sensor attributes. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration)
	// documentation.
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	// The name to use when displaying this switch.
	Name string `json:"name,omitempty"`
	// Flag that defines if switch works in optimistic mode.
	Optimistic bool `json:"optimistic,omitempty"`
	// The payload that represents the available state.
	PayloadAvailable string `json:"payload_available,omitempty"`
	// The payload that represents the unavailable state.
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	// The payload that represents `off` state. If specified, will be used for both
	// comparing to the value in the `state_topic` (see `value_template` and
	// `state_off` for details) and sending as `off` command to the `command_topic`.
	PayloadOff string `json:"payload_off,omitempty"`
	// The payload that represents `on` state. If specified, will be used for both
	// comparing to the value in the `state_topic` (see `value_template` and `state_on`
	// for details) and sending as `on` command to the `command_topic`.
	PayloadOn string `json:"payload_on,omitempty"`
	// The maximum QoS level of the state topic. Default is 0 and will also be used to
	// publishing messages.
	Qos int `json:"qos,omitempty"`
	// If the published message should have the retain flag on or not.
	Retain bool `json:"retain,omitempty"`
	// The payload that represents the `off` state. Used when value that represents
	// `off` state in the `state_topic` is different from value that should be sent to
	// the `command_topic` to turn the device `off`.
	StateOff string `json:"state_off,omitempty"`
	// The payload that represents the `on` state. Used when value that represents
	// `on` state in the `state_topic` is different from value that should be sent to
	// the `command_topic` to turn the device `on`.
	StateOn string `json:"state_on,omitempty"`
	// The MQTT topic subscribed to receive state updates.
	StateTopic string `json:"state_topic,omitempty"`
	// An ID that uniquely identifies this switch device. If two switches have the
	// same unique ID, Home Assistant will raise an exception.
	UniqueId string `json:"unique_id,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract device's state from the `state_topic`. To determine the switches's
	// state result of this template will be compared to `state_on` and `state_off`.
	ValueTemplate string `json:"value_template,omitempty"`
}

type HADeviceSwitchFunctions struct {
	Availability struct {
		State func() string
	}
	Command func(mqtt.Message, mqtt.Client)
	State   func() string
}

type HADeviceSwitchFunctionsConfig struct {
	Availability struct {
		State []string `json:"state,omitempty"`
	} `json:"availability,omitempty"`
	Command []string `json:"command,omitempty"`
	State   []string `json:"state,omitempty"`
}

// tag
type HADeviceTag struct {
	// Information about the device this device trigger is a part of to tie it into
	// the [device
	// registry](https://developers.home-assistant.io/docs/en/device_registry_index.html).
	// At least one of identifiers or connections must be present to identify the
	// device.
	Device struct {
		// A list of connections of the device to the outside world as a list of tuples
		// `[connection_type, connection_identifier]`. For example the MAC address of a
		// network interface: `'connections': ['mac', '02:5b:26:a8:dc:12']`.
		Connections []string `json:"connections,omitempty"`
		// A list of IDs that uniquely identify the device. For example a serial number.
		Identifiers []string `json:"identifiers,omitempty"`
		// The manufacturer of the device.
		Manufacturer string `json:"manufacturer,omitempty"`
		// The model of the device.
		Model string `json:"model,omitempty"`
		// The name of the device.
		Name string `json:"name,omitempty"`
		// Suggest an area if the device isn’t in one yet.
		SuggestedArea string `json:"suggested_area,omitempty"`
		// The firmware version of the device.
		SwVersion string `json:"sw_version,omitempty"`
		// Identifier of a device that routes messages between this device and Home
		// Assistant. Examples of such devices are hubs, or parent devices of a sub-device.
		// This is used to show device topology in Home Assistant.
		ViaDevice string `json:"via_device,omitempty"`
	} `json:"device,omitempty"`
	// The MQTT topic subscribed to receive tag scanned events.
	Topic string `json:"topic,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// that returns a tag ID.
	ValueTemplate string `json:"value_template,omitempty"`
}

type HADeviceTagFunctions struct {
	State func() string
}

type HADeviceTagFunctionsConfig struct {
	State []string `json:"state,omitempty"`
}

// vacuum
type HADeviceVacuum struct {
	// A list of MQTT topics subscribed to receive availability (online/offline)
	// updates. Must not be used together with `availability_topic`.
	Availability struct {
		// The payload that represents the available state.
		PayloadAvailable string `json:"payload_available,omitempty"`
		// The payload that represents the unavailable state.
		PayloadNotAvailable string `json:"payload_not_available,omitempty"`
		// An MQTT topic subscribed to receive availability (online/offline) updates.
		Topic string `json:"topic,omitempty"`
	} `json:"availability,omitempty"`
	// When `availability` is configured, this controls the conditions needed to set
	// the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set
	// to `all`, `payload_available` must be received on all configured availability
	// topics before the entity is marked as online. If set to `any`,
	// `payload_available` must be received on at least one configured availability
	// topic before the entity is marked as online. If set to `latest`, the last
	// `payload_available` or `payload_not_available` received on any configured
	// availability topic controls the availability.
	AvailabilityMode string `json:"availability_mode,omitempty"`
	// The MQTT topic subscribed to receive availability (online/offline) updates.
	// Must not be used together with `availability`.
	AvailabilityTopic string `json:"availability_topic,omitempty"`
	// Defines a [template](/topics/templating/) to define the battery level of the
	// vacuum. This is required if `battery_level_topic` is set.
	BatteryLevelTemplate string `json:"battery_level_template,omitempty"`
	// The MQTT topic subscribed to receive battery level values from the vacuum.
	BatteryLevelTopic string `json:"battery_level_topic,omitempty"`
	// Defines a [template](/topics/templating/) to define the charging state of the
	// vacuum. This is required if `charging_topic` is set.
	ChargingTemplate string `json:"charging_template,omitempty"`
	// The MQTT topic subscribed to receive charging state values from the vacuum.
	ChargingTopic string `json:"charging_topic,omitempty"`
	// Defines a [template](/topics/templating/) to define the cleaning state of the
	// vacuum. This is required if `cleaning_topic` is set.
	CleaningTemplate string `json:"cleaning_template,omitempty"`
	// The MQTT topic subscribed to receive cleaning state values from the vacuum.
	CleaningTopic string `json:"cleaning_topic,omitempty"`
	// The MQTT topic to publish commands to control the vacuum.
	CommandTopic string `json:"command_topic,omitempty"`
	// Defines a [template](/topics/templating/) to define the docked state of the
	// vacuum. This is required if `docked_topic` is set.
	DockedTemplate string `json:"docked_template,omitempty"`
	// The MQTT topic subscribed to receive docked state values from the vacuum.
	DockedTopic string `json:"docked_topic,omitempty"`
	// Flag which defines if the entity should be enabled when first added.
	EnabledByDefault bool `json:"enabled_by_default,omitempty"`
	// Defines a [template](/topics/templating/) to define potential error messages
	// emitted by the vacuum. This is required if `error_topic` is set.
	ErrorTemplate string `json:"error_template,omitempty"`
	// The MQTT topic subscribed to receive error messages from the vacuum.
	ErrorTopic string `json:"error_topic,omitempty"`
	// List of possible fan speeds for the vacuum.
	FanSpeedList []string `json:"fan_speed_list,omitempty"`
	// Defines a [template](/topics/templating/) to define the fan speed of the
	// vacuum. This is required if `fan_speed_topic` is set.
	FanSpeedTemplate string `json:"fan_speed_template,omitempty"`
	// The MQTT topic subscribed to receive fan speed values from the vacuum.
	FanSpeedTopic string `json:"fan_speed_topic,omitempty"`
	// [Icon](/docs/configuration/customizing-devices/#icon) for the entity.
	Icon string `json:"icon,omitempty"`
	// Defines a [template](/docs/configuration/templating/#processing-incoming-data)
	// to extract the JSON dictionary from messages received on the
	// `json_attributes_topic`. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-template-configuration)
	// documentation.
	JsonAttributesTemplate string `json:"json_attributes_template,omitempty"`
	// The MQTT topic subscribed to receive a JSON dictionary payload and then set as
	// sensor attributes. Usage example can be found in [MQTT
	// sensor](/integrations/sensor.mqtt/#json-attributes-topic-configuration)
	// documentation.
	JsonAttributesTopic string `json:"json_attributes_topic,omitempty"`
	// The name of the vacuum.
	Name string `json:"name,omitempty"`
	// The payload that represents the available state.
	PayloadAvailable string `json:"payload_available,omitempty"`
	// The payload to send to the `command_topic` to begin a spot cleaning cycle.
	PayloadCleanSpot string `json:"payload_clean_spot,omitempty"`
	// The payload to send to the `command_topic` to locate the vacuum (typically
	// plays a song).
	PayloadLocate string `json:"payload_locate,omitempty"`
	// The payload that represents the unavailable state.
	PayloadNotAvailable string `json:"payload_not_available,omitempty"`
	// The payload to send to the `command_topic` to tell the vacuum to return to
	// base.
	PayloadReturnToBase string `json:"payload_return_to_base,omitempty"`
	// The payload to send to the `command_topic` to start or pause the vacuum.
	PayloadStartPause string `json:"payload_start_pause,omitempty"`
	// The payload to send to the `command_topic` to stop the vacuum.
	PayloadStop string `json:"payload_stop,omitempty"`
	// The payload to send to the `command_topic` to turn the vacuum off.
	PayloadTurnOff string `json:"payload_turn_off,omitempty"`
	// The payload to send to the `command_topic` to begin the cleaning cycle.
	PayloadTurnOn string `json:"payload_turn_on,omitempty"`
	// The maximum QoS level of the state topic.
	Qos int `json:"qos,omitempty"`
	// If the published message should have the retain flag on or not.
	Retain bool `json:"retain,omitempty"`
	// The schema to use. Must be `legacy` or omitted to select the legacy schema.
	Schema string `json:"schema,omitempty"`
	// The MQTT topic to publish custom commands to the vacuum.
	SendCommandTopic string `json:"send_command_topic,omitempty"`
	// The MQTT topic to publish commands to control the vacuum's fan speed.
	SetFanSpeedTopic string `json:"set_fan_speed_topic,omitempty"`
	// List of features that the vacuum supports (possible values are `turn_on`,
	// `turn_off`, `pause`, `stop`, `return_home`, `battery`, `status`, `locate`,
	// `clean_spot`, `fan_speed`, `send_command`).
	SupportedFeatures []string `json:"supported_features,omitempty"`
	// An ID that uniquely identifies this vacuum. If two vacuums have the same unique
	// ID, Home Assistant will raise an exception.
	UniqueId string `json:"unique_id,omitempty"`
}

type HADeviceVacuumFunctions struct {
	Availability struct {
		State func() string
	}
	Command     func(mqtt.Message, mqtt.Client)
	SendCommand func(mqtt.Message, mqtt.Client)
	SetFanSpeed func(mqtt.Message, mqtt.Client)
}

type HADeviceVacuumFunctionsConfig struct {
	Availability struct {
		State []string `json:"state,omitempty"`
	} `json:"availability,omitempty"`
	Command     []string `json:"command,omitempty"`
	SendCommand []string `json:"send-command,omitempty"`
	SetFanSpeed []string `json:"set-fan-speed,omitempty"`
}
