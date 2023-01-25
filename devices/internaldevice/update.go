package InternalDevice

import (
	common "github.com/W-Floyd/ha-mqtt-iot/devices/common"
	externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Update struct {
	AvailabilityMode       *string     `json:"availability_mode,omitempty"`        // "When `availability` is configured, this controls the conditions needed to set the entity to `available`. Valid entries are `all`, `any`, and `latest`. If set to `all`, `payload_available` must be received on all configured availability topics before the entity is marked as online. If set to `any`, `payload_available` must be received on at least one configured availability topic before the entity is marked as online. If set to `latest`, the last `payload_available` or `payload_not_available` received on any configured availability topic controls the availability."
	AvailabilityTemplate   *string     `json:"availability_template,omitempty"`    // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract device's availability from the `availability_topic`. To determine the devices's availability result of this template will be compared to `payload_available` and `payload_not_available`."
	Availability           *([]string) `json:"availability,omitempty"`             // Availability for the Update
	Command                *([]string) `json:"command,omitempty"`                  // Command for the Update
	DeviceClass            *string     `json:"device_class,omitempty"`             // "The [type/class](/integrations/update/#device-classes) of the update to set the icon in the frontend."
	EnabledByDefault       *bool       `json:"enabled_by_default,omitempty"`       // "Flag which defines if the entity should be enabled when first added."
	Encoding               *string     `json:"encoding,omitempty"`                 // "The encoding of the payloads received and published messages. Set to `\"\"` to disable decoding of incoming payload."
	EntityCategory         *string     `json:"entity_category,omitempty"`          // "The [category](https://developers.home-assistant.io/docs/core/entity#generic-properties) of the entity."
	EntityPicture          *string     `json:"entity_picture,omitempty"`           // "Picture URL for the entity."
	Icon                   *string     `json:"icon,omitempty"`                     // "[Icon](/docs/configuration/customizing-devices/#icon) for the entity."
	JsonAttributesTemplate *string     `json:"json_attributes_template,omitempty"` // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the JSON dictionary from messages received on the `json_attributes_topic`."
	JsonAttributes         *([]string) `json:"json_attributes,omitempty"`          // JsonAttributes for the Update
	LatestVersionTemplate  *string     `json:"latest_version_template,omitempty"`  // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the latest version value."
	LatestVersion          *([]string) `json:"latest_version,omitempty"`           // LatestVersion for the Update
	Name                   *string     `json:"name,omitempty"`                     // "The name of the Select."
	ObjectId               *string     `json:"object_id,omitempty"`                // "Used instead of `name` for automatic generation of `entity_id`"
	PayloadInstall         *string     `json:"payload_install,omitempty"`          // "The MQTT payload to start installing process."
	Qos                    *int        `json:"qos,omitempty"`                      // "The maximum QoS level of the state topic. Default is 0 and will also be used to publishing messages."
	ReleaseSummary         *string     `json:"release_summary,omitempty"`          // "Summary of the release notes or changelog. This is suitable a brief update description of max 255 characters."
	ReleaseUrl             *string     `json:"release_url,omitempty"`              // "URL to the full release notes of the latest version available."
	Retain                 *bool       `json:"retain,omitempty"`                   // "If the published message should have the retain flag on or not."
	State                  *([]string) `json:"state,omitempty"`                    // State for the Update
	Title                  *string     `json:"title,omitempty"`                    // "Title of the software, or firmware update. This helps to differentiate between the device or entity name versus the title of the software installed."
	UniqueId               *string     `json:"unique_id,omitempty"`                // "An ID that uniquely identifies this Select. If two Selects have the same unique ID Home Assistant will raise an exception."
	ValueTemplate          *string     `json:"value_template,omitempty"`           // "Defines a [template](/docs/configuration/templating/#using-templates-with-the-mqtt-integration) to extract the `installed_version` state value or to render to a valid JSON payload on from the payload received on `state_topic`."
	MQTT                   struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}

func (iDevice Update) Translate() externaldevice.Update {
	eDevice := externaldevice.Update{}
	eDevice.MQTT = new(externaldevice.MQTTFields)
	if iDevice.MQTT.ForceUpdate != nil {
		eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	}
	if iDevice.MQTT.UpdateInterval != nil {
		eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	}
	if iDevice.AvailabilityMode != nil {
		eDevice.AvailabilityMode = iDevice.AvailabilityMode
	}
	if iDevice.AvailabilityTemplate != nil {
		eDevice.AvailabilityTemplate = iDevice.AvailabilityTemplate
	}
	if iDevice.Availability != nil {
		eDevice.AvailabilityFunc = common.ConstructStateFunc(*iDevice.Availability)
	}
	if iDevice.Command != nil {
		eDevice.CommandFunc = common.ConstructCommandFunc(*iDevice.Command)
	}
	if iDevice.DeviceClass != nil {
		eDevice.DeviceClass = iDevice.DeviceClass
	}
	if iDevice.EnabledByDefault != nil {
		eDevice.EnabledByDefault = iDevice.EnabledByDefault
	}
	if iDevice.Encoding != nil {
		eDevice.Encoding = iDevice.Encoding
	}
	if iDevice.EntityCategory != nil {
		eDevice.EntityCategory = iDevice.EntityCategory
	}
	if iDevice.EntityPicture != nil {
		eDevice.EntityPicture = iDevice.EntityPicture
	}
	if iDevice.Icon != nil {
		eDevice.Icon = iDevice.Icon
	}
	if iDevice.JsonAttributesTemplate != nil {
		eDevice.JsonAttributesTemplate = iDevice.JsonAttributesTemplate
	}
	if iDevice.JsonAttributes != nil {
		eDevice.JsonAttributesFunc = common.ConstructCommandFunc(*iDevice.JsonAttributes)
	}
	if iDevice.LatestVersionTemplate != nil {
		eDevice.LatestVersionTemplate = iDevice.LatestVersionTemplate
	}
	if iDevice.LatestVersion != nil {
		eDevice.LatestVersionFunc = common.ConstructCommandFunc(*iDevice.LatestVersion)
	}
	if iDevice.Name != nil {
		eDevice.Name = iDevice.Name
	}
	if iDevice.ObjectId != nil {
		eDevice.ObjectId = iDevice.ObjectId
	}
	if iDevice.PayloadInstall != nil {
		eDevice.PayloadInstall = iDevice.PayloadInstall
	}
	if iDevice.Qos != nil {
		eDevice.Qos = iDevice.Qos
	}
	if iDevice.ReleaseSummary != nil {
		eDevice.ReleaseSummary = iDevice.ReleaseSummary
	}
	if iDevice.ReleaseUrl != nil {
		eDevice.ReleaseUrl = iDevice.ReleaseUrl
	}
	if iDevice.Retain != nil {
		eDevice.Retain = iDevice.Retain
	}
	if iDevice.State != nil {
		eDevice.StateFunc = common.ConstructStateFunc(*iDevice.State)
	}
	if iDevice.Title != nil {
		eDevice.Title = iDevice.Title
	}
	if iDevice.UniqueId != nil {
		eDevice.UniqueId = iDevice.UniqueId
	}
	if iDevice.ValueTemplate != nil {
		eDevice.ValueTemplate = iDevice.ValueTemplate
	}
	if iDevice.Availability == nil {
		eDevice.AvailabilityFunc = common.AvailabilityFunc
	}
	eDevice.Initialize()
	return eDevice
}
