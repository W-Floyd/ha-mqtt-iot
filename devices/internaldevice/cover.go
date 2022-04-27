package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
func (iDevice Cover) Translate() externaldevice.Cover {
	eDevice := externaldevice.Cover{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.CommandFunc = common.ConstructCommandFunc(iDevice.Command)
	eDevice.DeviceClass = iDevice.DeviceClass
	eDevice.EnabledByDefault = iDevice.EnabledByDefault
	eDevice.Encoding = iDevice.Encoding
	eDevice.EntityCategory = iDevice.EntityCategory
	eDevice.Icon = iDevice.Icon
	eDevice.Name = iDevice.Name
	eDevice.ObjectId = iDevice.ObjectId
	eDevice.Optimistic = iDevice.Optimistic
	eDevice.PayloadAvailable = iDevice.PayloadAvailable
	eDevice.PayloadClose = iDevice.PayloadClose
	eDevice.PayloadNotAvailable = iDevice.PayloadNotAvailable
	eDevice.PayloadOpen = iDevice.PayloadOpen
	eDevice.PayloadStop = iDevice.PayloadStop
	eDevice.PositionClosed = iDevice.PositionClosed
	eDevice.PositionOpen = iDevice.PositionOpen
	eDevice.PositionTemplate = iDevice.PositionTemplate
	eDevice.PositionFunc = common.ConstructStateFunc(iDevice.Position)
	eDevice.Qos = iDevice.Qos
	eDevice.Retain = iDevice.Retain
	eDevice.SetPositionTemplate = iDevice.SetPositionTemplate
	eDevice.SetPositionFunc = common.ConstructCommandFunc(iDevice.SetPosition)
	eDevice.StateClosed = iDevice.StateClosed
	eDevice.StateClosing = iDevice.StateClosing
	eDevice.StateOpen = iDevice.StateOpen
	eDevice.StateOpening = iDevice.StateOpening
	eDevice.StateStopped = iDevice.StateStopped
	eDevice.StateFunc = common.ConstructStateFunc(iDevice.State)
	eDevice.TiltClosedValue = iDevice.TiltClosedValue
	eDevice.TiltCommandTemplate = iDevice.TiltCommandTemplate
	eDevice.TiltCommandFunc = common.ConstructCommandFunc(iDevice.TiltCommand)
	eDevice.TiltMax = iDevice.TiltMax
	eDevice.TiltMin = iDevice.TiltMin
	eDevice.TiltOpenedValue = iDevice.TiltOpenedValue
	eDevice.TiltOptimistic = iDevice.TiltOptimistic
	eDevice.TiltStatusTemplate = iDevice.TiltStatusTemplate
	eDevice.TiltStatusFunc = common.ConstructStateFunc(iDevice.TiltStatus)
	eDevice.UniqueId = iDevice.UniqueId
	eDevice.ValueTemplate = iDevice.ValueTemplate
	eDevice.Initialize()
	return eDevice
}

type Cover struct {
	Command             []string `json:"command"`
	DeviceClass         string   `json:"device_class"`          // "Sets the [class of the device](/integrations/cover/), changing the device state and icon that is displayed on the frontend."
	EnabledByDefault    bool     `json:"enabled_by_default"`    // "Flag which defines if the entity should be enabled when first added."
	Encoding            string   `json:"encoding"`              // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory      string   `json:"entity_category"`       // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	Icon                string   `json:"icon"`                  // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	Name                string   `json:"name"`                  // "The name of the cover."
	ObjectId            string   `json:"object_id"`             // "Used instead of `name` for automatic generation of `entity_id`"
	Optimistic          bool     `json:"optimistic"`            // "Flag that defines if switch works in optimistic mode."
	PayloadAvailable    string   `json:"payload_available"`     // "The payload that represents the online state."
	PayloadClose        string   `json:"payload_close"`         // "The command payload that closes the cover."
	PayloadNotAvailable string   `json:"payload_not_available"` // "The payload that represents the offline state."
	PayloadOpen         string   `json:"payload_open"`          // "The command payload that opens the cover."
	PayloadStop         string   `json:"payload_stop"`          // "The command payload that stops the cover."
	PositionClosed      int      `json:"position_closed"`       // "Number which represents closed position."
	PositionOpen        int      `json:"position_open"`         // "Number which represents open position."
	PositionTemplate    string   `json:"position_template"`     // "Defines a [template](/topics/templating/) that can be used to extract the payload for the `position_topic` topic. Within the template the following variables are available: `entity_id`, `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function;"
	Position            []string `json:"position"`
	Qos                 int      `json:"qos"`                   // "The maximum QoS level to be used when receiving and publishing messages."
	Retain              bool     `json:"retain"`                // "Defines if published messages should have the retain flag set."
	SetPositionTemplate string   `json:"set_position_template"` // "Defines a [template](/topics/templating/) to define the position to be sent to the `set_position_topic` topic. Incoming position value is available for use in the template `{% raw %}{{ position }}{% endraw %}`. Within the template the following variables are available: `entity_id`, `position`, the target position in percent; `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function;"
	SetPosition         []string `json:"set_position"`
	StateClosed         string   `json:"state_closed"`  // "The payload that represents the closed state."
	StateClosing        string   `json:"state_closing"` // "The payload that represents the closing state."
	StateOpen           string   `json:"state_open"`    // "The payload that represents the open state."
	StateOpening        string   `json:"state_opening"` // "The payload that represents the opening state."
	StateStopped        string   `json:"state_stopped"` // "The payload that represents the stopped state (for covers that do not report `open`/`closed` state)."
	State               []string `json:"state"`
	TiltClosedValue     int      `json:"tilt_closed_value"`     // "The value that will be sent on a `close_cover_tilt` command."
	TiltCommandTemplate string   `json:"tilt_command_template"` // "Defines a [template](/topics/templating/) that can be used to extract the payload for the `tilt_command_topic` topic. Within the template the following variables are available: `entity_id`, `tilt_position`, the target tilt position in percent; `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function;"
	TiltCommand         []string `json:"tilt_command"`
	TiltMax             int      `json:"tilt_max"`             // "The maximum tilt value."
	TiltMin             int      `json:"tilt_min"`             // "The minimum tilt value."
	TiltOpenedValue     int      `json:"tilt_opened_value"`    // "The value that will be sent on an `open_cover_tilt` command."
	TiltOptimistic      bool     `json:"tilt_optimistic"`      // "Flag that determines if tilt works in optimistic mode."
	TiltStatusTemplate  string   `json:"tilt_status_template"` // "Defines a [template](/topics/templating/) that can be used to extract the payload for the `tilt_status_topic` topic. Within the template the following variables are available: `entity_id`, `position_open`; `position_closed`; `tilt_min`; `tilt_max`. The `entity_id` can be used to reference the entity's attributes with help of the [states](/docs/configuration/templating/#states) template function;"
	TiltStatus          []string `json:"tilt_status"`
	UniqueId            string   `json:"unique_id"`      // "An ID that uniquely identifies this cover. If two covers have the same unique ID, Home Assistant will raise an exception."
	ValueTemplate       string   `json:"value_template"` // "Defines a [template](/topics/templating/) that can be used to extract the payload for the `state_topic` topic."
	MQTT                struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
