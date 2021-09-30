package devices

import "github.com/W-Floyd/ha-mqtt-iot/logging"

func (entity HADeviceAlarmControlPanel) GenerateAlarmControlPanel() (output HADeviceAlarmControlPanel) {
	if entity.Availability.PayloadAvailable != "" {
		output.Availability.PayloadAvailable = entity.Availability.PayloadAvailable
	}
	if entity.Availability.PayloadNotAvailable != "" {
		output.Availability.PayloadNotAvailable = entity.Availability.PayloadNotAvailable
	}
	if entity.Availability.Topic != "" {
		output.Availability.Topic = entity.Availability.Topic
	} else {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.AvailabilityMode != "" {
		output.AvailabilityMode = entity.AvailabilityMode
	}
	if entity.AvailabilityTopic != "" {
		output.AvailabilityTopic = entity.AvailabilityTopic
	}
	if entity.Code != "" {
		output.Code = entity.Code
	}
	if entity.CodeArmRequired != false {
		output.CodeArmRequired = entity.CodeArmRequired
	}
	if entity.CodeDisarmRequired != false {
		output.CodeDisarmRequired = entity.CodeDisarmRequired
	}
	if entity.CommandTemplate != "" {
		output.CommandTemplate = entity.CommandTemplate
	}
	if entity.CommandTopic != "" {
		output.CommandTopic = entity.CommandTopic
	} else {
		logging.LogError("entity.CommandTopic generator not found, but field is required!")
	}
	if entity.Device.Connections != nil {
		output.Device.Connections = entity.Device.Connections
	}
	if entity.Device.Identifiers != nil {
		output.Device.Identifiers = entity.Device.Identifiers
	}
	if entity.Device.Manufacturer != "" {
		output.Device.Manufacturer = entity.Device.Manufacturer
	}
	if entity.Device.Model != "" {
		output.Device.Model = entity.Device.Model
	}
	if entity.Device.Name != "" {
		output.Device.Name = entity.Device.Name
	}
	if entity.Device.SuggestedArea != "" {
		output.Device.SuggestedArea = entity.Device.SuggestedArea
	}
	if entity.Device.SwVersion != "" {
		output.Device.SwVersion = entity.Device.SwVersion
	}
	if entity.Device.ViaDevice != "" {
		output.Device.ViaDevice = entity.Device.ViaDevice
	}
	if entity.EnabledByDefault != false {
		output.EnabledByDefault = entity.EnabledByDefault
	}
	if entity.Icon != "" {
		output.Icon = entity.Icon
	}
	if entity.JsonAttributesTemplate != "" {
		output.JsonAttributesTemplate = entity.JsonAttributesTemplate
	}
	if entity.JsonAttributesTopic != "" {
		output.JsonAttributesTopic = entity.JsonAttributesTopic
	}
	if entity.Name != "" {
		output.Name = entity.Name
	}
	if entity.PayloadArmAway != "" {
		output.PayloadArmAway = entity.PayloadArmAway
	}
	if entity.PayloadArmCustomBypass != "" {
		output.PayloadArmCustomBypass = entity.PayloadArmCustomBypass
	}
	if entity.PayloadArmHome != "" {
		output.PayloadArmHome = entity.PayloadArmHome
	}
	if entity.PayloadArmNight != "" {
		output.PayloadArmNight = entity.PayloadArmNight
	}
	if entity.PayloadArmVacation != "" {
		output.PayloadArmVacation = entity.PayloadArmVacation
	}
	if entity.PayloadAvailable != "" {
		output.PayloadAvailable = entity.PayloadAvailable
	}
	if entity.PayloadDisarm != "" {
		output.PayloadDisarm = entity.PayloadDisarm
	}
	if entity.PayloadNotAvailable != "" {
		output.PayloadNotAvailable = entity.PayloadNotAvailable
	}
	if entity.Qos != 0 {
		output.Qos = entity.Qos
	}
	if entity.Retain != false {
		output.Retain = entity.Retain
	}
	if entity.StateTopic != "" {
		output.StateTopic = entity.StateTopic
	} else {
		logging.LogError("entity.StateTopic generator not found, but field is required!")
	}
	if entity.UniqueId != "" {
		output.UniqueId = entity.UniqueId
	}
	if entity.ValueTemplate != "" {
		output.ValueTemplate = entity.ValueTemplate
	}
	return
}
func (entity HADeviceBinarySensor) GenerateBinarySensor() (output HADeviceBinarySensor) {
	if entity.Availability.PayloadAvailable != "" {
		output.Availability.PayloadAvailable = entity.Availability.PayloadAvailable
	}
	if entity.Availability.PayloadNotAvailable != "" {
		output.Availability.PayloadNotAvailable = entity.Availability.PayloadNotAvailable
	}
	if entity.Availability.Topic != "" {
		output.Availability.Topic = entity.Availability.Topic
	} else {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.AvailabilityMode != "" {
		output.AvailabilityMode = entity.AvailabilityMode
	}
	if entity.AvailabilityTopic != "" {
		output.AvailabilityTopic = entity.AvailabilityTopic
	}
	if entity.Device.Connections != nil {
		output.Device.Connections = entity.Device.Connections
	}
	if entity.Device.Identifiers != nil {
		output.Device.Identifiers = entity.Device.Identifiers
	}
	if entity.Device.Manufacturer != "" {
		output.Device.Manufacturer = entity.Device.Manufacturer
	}
	if entity.Device.Model != "" {
		output.Device.Model = entity.Device.Model
	}
	if entity.Device.Name != "" {
		output.Device.Name = entity.Device.Name
	}
	if entity.Device.SuggestedArea != "" {
		output.Device.SuggestedArea = entity.Device.SuggestedArea
	}
	if entity.Device.SwVersion != "" {
		output.Device.SwVersion = entity.Device.SwVersion
	}
	if entity.Device.ViaDevice != "" {
		output.Device.ViaDevice = entity.Device.ViaDevice
	}
	if entity.DeviceClass != "" {
		output.DeviceClass = entity.DeviceClass
	}
	if entity.EnabledByDefault != false {
		output.EnabledByDefault = entity.EnabledByDefault
	}
	if entity.ExpireAfter != 0 {
		output.ExpireAfter = entity.ExpireAfter
	}
	if entity.ForceUpdate != false {
		output.ForceUpdate = entity.ForceUpdate
	}
	if entity.Icon != "" {
		output.Icon = entity.Icon
	}
	if entity.JsonAttributesTemplate != "" {
		output.JsonAttributesTemplate = entity.JsonAttributesTemplate
	}
	if entity.JsonAttributesTopic != "" {
		output.JsonAttributesTopic = entity.JsonAttributesTopic
	}
	if entity.Name != "" {
		output.Name = entity.Name
	}
	if entity.OffDelay != 0 {
		output.OffDelay = entity.OffDelay
	}
	if entity.PayloadAvailable != "" {
		output.PayloadAvailable = entity.PayloadAvailable
	}
	if entity.PayloadNotAvailable != "" {
		output.PayloadNotAvailable = entity.PayloadNotAvailable
	}
	if entity.PayloadOff != "" {
		output.PayloadOff = entity.PayloadOff
	}
	if entity.PayloadOn != "" {
		output.PayloadOn = entity.PayloadOn
	}
	if entity.Qos != 0 {
		output.Qos = entity.Qos
	}
	if entity.StateTopic != "" {
		output.StateTopic = entity.StateTopic
	} else {
		logging.LogError("entity.StateTopic generator not found, but field is required!")
	}
	if entity.UniqueId != "" {
		output.UniqueId = entity.UniqueId
	}
	if entity.ValueTemplate != "" {
		output.ValueTemplate = entity.ValueTemplate
	}
	return
}
func (entity HADeviceCamera) GenerateCamera() (output HADeviceCamera) {
	if entity.Availability.PayloadAvailable != "" {
		output.Availability.PayloadAvailable = entity.Availability.PayloadAvailable
	}
	if entity.Availability.PayloadNotAvailable != "" {
		output.Availability.PayloadNotAvailable = entity.Availability.PayloadNotAvailable
	}
	if entity.Availability.Topic != "" {
		output.Availability.Topic = entity.Availability.Topic
	} else {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.AvailabilityMode != "" {
		output.AvailabilityMode = entity.AvailabilityMode
	}
	if entity.AvailabilityTopic != "" {
		output.AvailabilityTopic = entity.AvailabilityTopic
	}
	if entity.Device.Connections != nil {
		output.Device.Connections = entity.Device.Connections
	}
	if entity.Device.Identifiers != nil {
		output.Device.Identifiers = entity.Device.Identifiers
	}
	if entity.Device.Manufacturer != "" {
		output.Device.Manufacturer = entity.Device.Manufacturer
	}
	if entity.Device.Model != "" {
		output.Device.Model = entity.Device.Model
	}
	if entity.Device.Name != "" {
		output.Device.Name = entity.Device.Name
	}
	if entity.Device.SuggestedArea != "" {
		output.Device.SuggestedArea = entity.Device.SuggestedArea
	}
	if entity.Device.SwVersion != "" {
		output.Device.SwVersion = entity.Device.SwVersion
	}
	if entity.Device.ViaDevice != "" {
		output.Device.ViaDevice = entity.Device.ViaDevice
	}
	if entity.EnabledByDefault != false {
		output.EnabledByDefault = entity.EnabledByDefault
	}
	if entity.Icon != "" {
		output.Icon = entity.Icon
	}
	if entity.JsonAttributesTemplate != "" {
		output.JsonAttributesTemplate = entity.JsonAttributesTemplate
	}
	if entity.JsonAttributesTopic != "" {
		output.JsonAttributesTopic = entity.JsonAttributesTopic
	}
	if entity.Name != "" {
		output.Name = entity.Name
	}
	if entity.Topic != "" {
		output.Topic = entity.Topic
	} else {
		logging.LogError("entity.Topic generator not found, but field is required!")
	}
	if entity.UniqueId != "" {
		output.UniqueId = entity.UniqueId
	}
	return
}
func (entity HADeviceCover) GenerateCover() (output HADeviceCover) {
	if entity.Availability.PayloadAvailable != "" {
		output.Availability.PayloadAvailable = entity.Availability.PayloadAvailable
	}
	if entity.Availability.PayloadNotAvailable != "" {
		output.Availability.PayloadNotAvailable = entity.Availability.PayloadNotAvailable
	}
	if entity.Availability.Topic != "" {
		output.Availability.Topic = entity.Availability.Topic
	} else {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.AvailabilityMode != "" {
		output.AvailabilityMode = entity.AvailabilityMode
	}
	if entity.AvailabilityTopic != "" {
		output.AvailabilityTopic = entity.AvailabilityTopic
	}
	if entity.CommandTopic != "" {
		output.CommandTopic = entity.CommandTopic
	}
	if entity.Device.Connections != nil {
		output.Device.Connections = entity.Device.Connections
	}
	if entity.Device.Identifiers != nil {
		output.Device.Identifiers = entity.Device.Identifiers
	}
	if entity.Device.Manufacturer != "" {
		output.Device.Manufacturer = entity.Device.Manufacturer
	}
	if entity.Device.Model != "" {
		output.Device.Model = entity.Device.Model
	}
	if entity.Device.Name != "" {
		output.Device.Name = entity.Device.Name
	}
	if entity.Device.SuggestedArea != "" {
		output.Device.SuggestedArea = entity.Device.SuggestedArea
	}
	if entity.Device.SwVersion != "" {
		output.Device.SwVersion = entity.Device.SwVersion
	}
	if entity.Device.ViaDevice != "" {
		output.Device.ViaDevice = entity.Device.ViaDevice
	}
	if entity.DeviceClass != "" {
		output.DeviceClass = entity.DeviceClass
	}
	if entity.EnabledByDefault != false {
		output.EnabledByDefault = entity.EnabledByDefault
	}
	if entity.Icon != "" {
		output.Icon = entity.Icon
	}
	if entity.JsonAttributesTemplate != "" {
		output.JsonAttributesTemplate = entity.JsonAttributesTemplate
	}
	if entity.JsonAttributesTopic != "" {
		output.JsonAttributesTopic = entity.JsonAttributesTopic
	}
	if entity.Name != "" {
		output.Name = entity.Name
	}
	if entity.Optimistic != false {
		output.Optimistic = entity.Optimistic
	}
	if entity.PayloadAvailable != "" {
		output.PayloadAvailable = entity.PayloadAvailable
	}
	if entity.PayloadClose != "" {
		output.PayloadClose = entity.PayloadClose
	}
	if entity.PayloadNotAvailable != "" {
		output.PayloadNotAvailable = entity.PayloadNotAvailable
	}
	if entity.PayloadOpen != "" {
		output.PayloadOpen = entity.PayloadOpen
	}
	if entity.PayloadStop != "" {
		output.PayloadStop = entity.PayloadStop
	}
	if entity.PositionClosed != 0 {
		output.PositionClosed = entity.PositionClosed
	}
	if entity.PositionOpen != 0 {
		output.PositionOpen = entity.PositionOpen
	}
	if entity.PositionTemplate != "" {
		output.PositionTemplate = entity.PositionTemplate
	}
	if entity.PositionTopic != "" {
		output.PositionTopic = entity.PositionTopic
	}
	if entity.Qos != 0 {
		output.Qos = entity.Qos
	}
	if entity.Retain != false {
		output.Retain = entity.Retain
	}
	if entity.SetPositionTemplate != "" {
		output.SetPositionTemplate = entity.SetPositionTemplate
	}
	if entity.SetPositionTopic != "" {
		output.SetPositionTopic = entity.SetPositionTopic
	}
	if entity.StateClosed != "" {
		output.StateClosed = entity.StateClosed
	}
	if entity.StateClosing != "" {
		output.StateClosing = entity.StateClosing
	}
	if entity.StateOpen != "" {
		output.StateOpen = entity.StateOpen
	}
	if entity.StateOpening != "" {
		output.StateOpening = entity.StateOpening
	}
	if entity.StateStopped != "" {
		output.StateStopped = entity.StateStopped
	}
	if entity.StateTopic != "" {
		output.StateTopic = entity.StateTopic
	}
	if entity.TiltClosedValue != 0 {
		output.TiltClosedValue = entity.TiltClosedValue
	}
	if entity.TiltCommandTemplate != "" {
		output.TiltCommandTemplate = entity.TiltCommandTemplate
	}
	if entity.TiltCommandTopic != "" {
		output.TiltCommandTopic = entity.TiltCommandTopic
	}
	if entity.TiltMax != 0 {
		output.TiltMax = entity.TiltMax
	}
	if entity.TiltMin != 0 {
		output.TiltMin = entity.TiltMin
	}
	if entity.TiltOpenedValue != 0 {
		output.TiltOpenedValue = entity.TiltOpenedValue
	}
	if entity.TiltOptimistic != false {
		output.TiltOptimistic = entity.TiltOptimistic
	}
	if entity.TiltStatusTemplate != "" {
		output.TiltStatusTemplate = entity.TiltStatusTemplate
	}
	if entity.TiltStatusTopic != "" {
		output.TiltStatusTopic = entity.TiltStatusTopic
	}
	if entity.UniqueId != "" {
		output.UniqueId = entity.UniqueId
	}
	if entity.ValueTemplate != "" {
		output.ValueTemplate = entity.ValueTemplate
	}
	return
}
func (entity HADeviceDeviceTracker) GenerateDeviceTracker() (output HADeviceDeviceTracker) {
	if entity.Devices != nil {
		output.Devices = entity.Devices
	} else {
		logging.LogError("entity.Devices generator not found, but field is required!")
	}
	if entity.PayloadHome != "" {
		output.PayloadHome = entity.PayloadHome
	}
	if entity.PayloadNotHome != "" {
		output.PayloadNotHome = entity.PayloadNotHome
	}
	if entity.Qos != 0 {
		output.Qos = entity.Qos
	}
	if entity.SourceType != "" {
		output.SourceType = entity.SourceType
	}
	return
}
func (entity HADeviceDeviceTrigger) GenerateDeviceTrigger() (output HADeviceDeviceTrigger) {
	if entity.AutomationType != "" {
		output.AutomationType = entity.AutomationType
	} else {
		logging.LogError("entity.AutomationType generator not found, but field is required!")
	}
	if entity.Device.Connections != nil {
		output.Device.Connections = entity.Device.Connections
	}
	if entity.Device.Identifiers != nil {
		output.Device.Identifiers = entity.Device.Identifiers
	}
	if entity.Device.Manufacturer != "" {
		output.Device.Manufacturer = entity.Device.Manufacturer
	}
	if entity.Device.Model != "" {
		output.Device.Model = entity.Device.Model
	}
	if entity.Device.Name != "" {
		output.Device.Name = entity.Device.Name
	}
	if entity.Device.SuggestedArea != "" {
		output.Device.SuggestedArea = entity.Device.SuggestedArea
	}
	if entity.Device.SwVersion != "" {
		output.Device.SwVersion = entity.Device.SwVersion
	}
	if entity.Device.ViaDevice != "" {
		output.Device.ViaDevice = entity.Device.ViaDevice
	}
	if entity.Payload != "" {
		output.Payload = entity.Payload
	}
	if entity.Qos != 0 {
		output.Qos = entity.Qos
	}
	if entity.Subtype != "" {
		output.Subtype = entity.Subtype
	} else {
		logging.LogError("entity.Subtype generator not found, but field is required!")
	}
	if entity.Topic != "" {
		output.Topic = entity.Topic
	} else {
		logging.LogError("entity.Topic generator not found, but field is required!")
	}
	if entity.Type != "" {
		output.Type = entity.Type
	} else {
		logging.LogError("entity.Type generator not found, but field is required!")
	}
	return
}
func (entity HADeviceFan) GenerateFan() (output HADeviceFan) {
	if entity.Availability.PayloadAvailable != "" {
		output.Availability.PayloadAvailable = entity.Availability.PayloadAvailable
	}
	if entity.Availability.PayloadNotAvailable != "" {
		output.Availability.PayloadNotAvailable = entity.Availability.PayloadNotAvailable
	}
	if entity.Availability.Topic != "" {
		output.Availability.Topic = entity.Availability.Topic
	} else {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.AvailabilityMode != "" {
		output.AvailabilityMode = entity.AvailabilityMode
	}
	if entity.AvailabilityTopic != "" {
		output.AvailabilityTopic = entity.AvailabilityTopic
	}
	if entity.CommandTemplate != "" {
		output.CommandTemplate = entity.CommandTemplate
	}
	if entity.CommandTopic != "" {
		output.CommandTopic = entity.CommandTopic
	} else {
		logging.LogError("entity.CommandTopic generator not found, but field is required!")
	}
	if entity.Device.Connections != nil {
		output.Device.Connections = entity.Device.Connections
	}
	if entity.Device.Identifiers != nil {
		output.Device.Identifiers = entity.Device.Identifiers
	}
	if entity.Device.Manufacturer != "" {
		output.Device.Manufacturer = entity.Device.Manufacturer
	}
	if entity.Device.Model != "" {
		output.Device.Model = entity.Device.Model
	}
	if entity.Device.Name != "" {
		output.Device.Name = entity.Device.Name
	}
	if entity.Device.SuggestedArea != "" {
		output.Device.SuggestedArea = entity.Device.SuggestedArea
	}
	if entity.Device.SwVersion != "" {
		output.Device.SwVersion = entity.Device.SwVersion
	}
	if entity.Device.ViaDevice != "" {
		output.Device.ViaDevice = entity.Device.ViaDevice
	}
	if entity.EnabledByDefault != false {
		output.EnabledByDefault = entity.EnabledByDefault
	}
	if entity.Icon != "" {
		output.Icon = entity.Icon
	}
	if entity.JsonAttributesTemplate != "" {
		output.JsonAttributesTemplate = entity.JsonAttributesTemplate
	}
	if entity.JsonAttributesTopic != "" {
		output.JsonAttributesTopic = entity.JsonAttributesTopic
	}
	if entity.Name != "" {
		output.Name = entity.Name
	}
	if entity.Optimistic != false {
		output.Optimistic = entity.Optimistic
	}
	if entity.OscillationCommandTemplate != "" {
		output.OscillationCommandTemplate = entity.OscillationCommandTemplate
	}
	if entity.OscillationCommandTopic != "" {
		output.OscillationCommandTopic = entity.OscillationCommandTopic
	}
	if entity.OscillationStateTopic != "" {
		output.OscillationStateTopic = entity.OscillationStateTopic
	}
	if entity.OscillationValueTemplate != "" {
		output.OscillationValueTemplate = entity.OscillationValueTemplate
	}
	if entity.PayloadAvailable != "" {
		output.PayloadAvailable = entity.PayloadAvailable
	}
	if entity.PayloadNotAvailable != "" {
		output.PayloadNotAvailable = entity.PayloadNotAvailable
	}
	if entity.PayloadOff != "" {
		output.PayloadOff = entity.PayloadOff
	}
	if entity.PayloadOn != "" {
		output.PayloadOn = entity.PayloadOn
	}
	if entity.PayloadOscillationOff != "" {
		output.PayloadOscillationOff = entity.PayloadOscillationOff
	}
	if entity.PayloadOscillationOn != "" {
		output.PayloadOscillationOn = entity.PayloadOscillationOn
	}
	if entity.PayloadResetPercentage != "" {
		output.PayloadResetPercentage = entity.PayloadResetPercentage
	}
	if entity.PayloadResetPresetMode != "" {
		output.PayloadResetPresetMode = entity.PayloadResetPresetMode
	}
	if entity.PercentageCommandTemplate != "" {
		output.PercentageCommandTemplate = entity.PercentageCommandTemplate
	}
	if entity.PercentageCommandTopic != "" {
		output.PercentageCommandTopic = entity.PercentageCommandTopic
	}
	if entity.PercentageStateTopic != "" {
		output.PercentageStateTopic = entity.PercentageStateTopic
	}
	if entity.PercentageValueTemplate != "" {
		output.PercentageValueTemplate = entity.PercentageValueTemplate
	}
	if entity.PresetModeCommandTemplate != "" {
		output.PresetModeCommandTemplate = entity.PresetModeCommandTemplate
	}
	if entity.PresetModeCommandTopic != "" {
		output.PresetModeCommandTopic = entity.PresetModeCommandTopic
	}
	if entity.PresetModeStateTopic != "" {
		output.PresetModeStateTopic = entity.PresetModeStateTopic
	}
	if entity.PresetModeValueTemplate != "" {
		output.PresetModeValueTemplate = entity.PresetModeValueTemplate
	}
	if entity.PresetModes != nil {
		output.PresetModes = entity.PresetModes
	}
	if entity.Qos != 0 {
		output.Qos = entity.Qos
	}
	if entity.Retain != false {
		output.Retain = entity.Retain
	}
	if entity.SpeedRangeMax != 0 {
		output.SpeedRangeMax = entity.SpeedRangeMax
	}
	if entity.SpeedRangeMin != 0 {
		output.SpeedRangeMin = entity.SpeedRangeMin
	}
	if entity.StateTopic != "" {
		output.StateTopic = entity.StateTopic
	}
	if entity.StateValueTemplate != "" {
		output.StateValueTemplate = entity.StateValueTemplate
	}
	if entity.UniqueId != "" {
		output.UniqueId = entity.UniqueId
	}
	return
}
func (entity HADeviceHumidifier) GenerateHumidifier() (output HADeviceHumidifier) {
	if entity.Availability.PayloadAvailable != "" {
		output.Availability.PayloadAvailable = entity.Availability.PayloadAvailable
	}
	if entity.Availability.PayloadNotAvailable != "" {
		output.Availability.PayloadNotAvailable = entity.Availability.PayloadNotAvailable
	}
	if entity.Availability.Topic != "" {
		output.Availability.Topic = entity.Availability.Topic
	} else {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.AvailabilityMode != "" {
		output.AvailabilityMode = entity.AvailabilityMode
	}
	if entity.AvailabilityTopic != "" {
		output.AvailabilityTopic = entity.AvailabilityTopic
	}
	if entity.CommandTemplate != "" {
		output.CommandTemplate = entity.CommandTemplate
	}
	if entity.CommandTopic != "" {
		output.CommandTopic = entity.CommandTopic
	} else {
		logging.LogError("entity.CommandTopic generator not found, but field is required!")
	}
	if entity.Device.Connections != nil {
		output.Device.Connections = entity.Device.Connections
	}
	if entity.Device.Identifiers != nil {
		output.Device.Identifiers = entity.Device.Identifiers
	}
	if entity.Device.Manufacturer != "" {
		output.Device.Manufacturer = entity.Device.Manufacturer
	}
	if entity.Device.Model != "" {
		output.Device.Model = entity.Device.Model
	}
	if entity.Device.Name != "" {
		output.Device.Name = entity.Device.Name
	}
	if entity.Device.SuggestedArea != "" {
		output.Device.SuggestedArea = entity.Device.SuggestedArea
	}
	if entity.Device.SwVersion != "" {
		output.Device.SwVersion = entity.Device.SwVersion
	}
	if entity.Device.ViaDevice != "" {
		output.Device.ViaDevice = entity.Device.ViaDevice
	}
	if entity.DeviceClass != "" {
		output.DeviceClass = entity.DeviceClass
	}
	if entity.EnabledByDefault != false {
		output.EnabledByDefault = entity.EnabledByDefault
	}
	if entity.Icon != "" {
		output.Icon = entity.Icon
	}
	if entity.JsonAttributesTemplate != "" {
		output.JsonAttributesTemplate = entity.JsonAttributesTemplate
	}
	if entity.JsonAttributesTopic != "" {
		output.JsonAttributesTopic = entity.JsonAttributesTopic
	}
	if entity.MaxHumidity != 0 {
		output.MaxHumidity = entity.MaxHumidity
	}
	if entity.MinHumidity != 0 {
		output.MinHumidity = entity.MinHumidity
	}
	if entity.ModeCommandTemplate != "" {
		output.ModeCommandTemplate = entity.ModeCommandTemplate
	}
	if entity.ModeCommandTopic != "" {
		output.ModeCommandTopic = entity.ModeCommandTopic
	}
	if entity.ModeStateTemplate != "" {
		output.ModeStateTemplate = entity.ModeStateTemplate
	}
	if entity.ModeStateTopic != "" {
		output.ModeStateTopic = entity.ModeStateTopic
	}
	if entity.Modes != nil {
		output.Modes = entity.Modes
	}
	if entity.Name != "" {
		output.Name = entity.Name
	}
	if entity.Optimistic != false {
		output.Optimistic = entity.Optimistic
	}
	if entity.PayloadAvailable != "" {
		output.PayloadAvailable = entity.PayloadAvailable
	}
	if entity.PayloadNotAvailable != "" {
		output.PayloadNotAvailable = entity.PayloadNotAvailable
	}
	if entity.PayloadOff != "" {
		output.PayloadOff = entity.PayloadOff
	}
	if entity.PayloadOn != "" {
		output.PayloadOn = entity.PayloadOn
	}
	if entity.PayloadResetHumidity != "" {
		output.PayloadResetHumidity = entity.PayloadResetHumidity
	}
	if entity.PayloadResetMode != "" {
		output.PayloadResetMode = entity.PayloadResetMode
	}
	if entity.Qos != 0 {
		output.Qos = entity.Qos
	}
	if entity.Retain != false {
		output.Retain = entity.Retain
	}
	if entity.StateTopic != "" {
		output.StateTopic = entity.StateTopic
	}
	if entity.StateValueTemplate != "" {
		output.StateValueTemplate = entity.StateValueTemplate
	}
	if entity.TargetHumidityCommandTemplate != "" {
		output.TargetHumidityCommandTemplate = entity.TargetHumidityCommandTemplate
	}
	if entity.TargetHumidityCommandTopic != "" {
		output.TargetHumidityCommandTopic = entity.TargetHumidityCommandTopic
	} else {
		logging.LogError("entity.TargetHumidityCommandTopic generator not found, but field is required!")
	}
	if entity.TargetHumidityStateTemplate != "" {
		output.TargetHumidityStateTemplate = entity.TargetHumidityStateTemplate
	}
	if entity.TargetHumidityStateTopic != "" {
		output.TargetHumidityStateTopic = entity.TargetHumidityStateTopic
	}
	if entity.UniqueId != "" {
		output.UniqueId = entity.UniqueId
	}
	return
}
func (entity HADeviceClimate) GenerateClimate() (output HADeviceClimate) {
	if entity.ActionTemplate != "" {
		output.ActionTemplate = entity.ActionTemplate
	}
	if entity.ActionTopic != "" {
		output.ActionTopic = entity.ActionTopic
	}
	if entity.AuxCommandTopic != "" {
		output.AuxCommandTopic = entity.AuxCommandTopic
	}
	if entity.AuxStateTemplate != "" {
		output.AuxStateTemplate = entity.AuxStateTemplate
	}
	if entity.AuxStateTopic != "" {
		output.AuxStateTopic = entity.AuxStateTopic
	}
	if entity.Availability.PayloadAvailable != "" {
		output.Availability.PayloadAvailable = entity.Availability.PayloadAvailable
	}
	if entity.Availability.PayloadNotAvailable != "" {
		output.Availability.PayloadNotAvailable = entity.Availability.PayloadNotAvailable
	}
	if entity.Availability.Topic != "" {
		output.Availability.Topic = entity.Availability.Topic
	} else {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.AvailabilityMode != "" {
		output.AvailabilityMode = entity.AvailabilityMode
	}
	if entity.AvailabilityTopic != "" {
		output.AvailabilityTopic = entity.AvailabilityTopic
	}
	if entity.AwayModeCommandTopic != "" {
		output.AwayModeCommandTopic = entity.AwayModeCommandTopic
	}
	if entity.AwayModeStateTemplate != "" {
		output.AwayModeStateTemplate = entity.AwayModeStateTemplate
	}
	if entity.AwayModeStateTopic != "" {
		output.AwayModeStateTopic = entity.AwayModeStateTopic
	}
	if entity.CurrentTemperatureTemplate != "" {
		output.CurrentTemperatureTemplate = entity.CurrentTemperatureTemplate
	}
	if entity.CurrentTemperatureTopic != "" {
		output.CurrentTemperatureTopic = entity.CurrentTemperatureTopic
	}
	if entity.Device.Connections != nil {
		output.Device.Connections = entity.Device.Connections
	}
	if entity.Device.Identifiers != nil {
		output.Device.Identifiers = entity.Device.Identifiers
	}
	if entity.Device.Manufacturer != "" {
		output.Device.Manufacturer = entity.Device.Manufacturer
	}
	if entity.Device.Model != "" {
		output.Device.Model = entity.Device.Model
	}
	if entity.Device.Name != "" {
		output.Device.Name = entity.Device.Name
	}
	if entity.Device.SuggestedArea != "" {
		output.Device.SuggestedArea = entity.Device.SuggestedArea
	}
	if entity.Device.SwVersion != "" {
		output.Device.SwVersion = entity.Device.SwVersion
	}
	if entity.Device.ViaDevice != "" {
		output.Device.ViaDevice = entity.Device.ViaDevice
	}
	if entity.EnabledByDefault != false {
		output.EnabledByDefault = entity.EnabledByDefault
	}
	if entity.FanModeCommandTemplate != "" {
		output.FanModeCommandTemplate = entity.FanModeCommandTemplate
	}
	if entity.FanModeCommandTopic != "" {
		output.FanModeCommandTopic = entity.FanModeCommandTopic
	}
	if entity.FanModeStateTemplate != "" {
		output.FanModeStateTemplate = entity.FanModeStateTemplate
	}
	if entity.FanModeStateTopic != "" {
		output.FanModeStateTopic = entity.FanModeStateTopic
	}
	if entity.FanModes != nil {
		output.FanModes = entity.FanModes
	}
	if entity.HoldCommandTemplate != "" {
		output.HoldCommandTemplate = entity.HoldCommandTemplate
	}
	if entity.HoldCommandTopic != "" {
		output.HoldCommandTopic = entity.HoldCommandTopic
	}
	if entity.HoldModes != nil {
		output.HoldModes = entity.HoldModes
	}
	if entity.HoldStateTemplate != "" {
		output.HoldStateTemplate = entity.HoldStateTemplate
	}
	if entity.HoldStateTopic != "" {
		output.HoldStateTopic = entity.HoldStateTopic
	}
	if entity.Icon != "" {
		output.Icon = entity.Icon
	}
	if entity.Initial != 0 {
		output.Initial = entity.Initial
	}
	if entity.JsonAttributesTemplate != "" {
		output.JsonAttributesTemplate = entity.JsonAttributesTemplate
	}
	if entity.JsonAttributesTopic != "" {
		output.JsonAttributesTopic = entity.JsonAttributesTopic
	}
	if entity.MaxTemp != 0 {
		output.MaxTemp = entity.MaxTemp
	}
	if entity.MinTemp != 0 {
		output.MinTemp = entity.MinTemp
	}
	if entity.ModeCommandTemplate != "" {
		output.ModeCommandTemplate = entity.ModeCommandTemplate
	}
	if entity.ModeCommandTopic != "" {
		output.ModeCommandTopic = entity.ModeCommandTopic
	}
	if entity.ModeStateTemplate != "" {
		output.ModeStateTemplate = entity.ModeStateTemplate
	}
	if entity.ModeStateTopic != "" {
		output.ModeStateTopic = entity.ModeStateTopic
	}
	if entity.Modes != nil {
		output.Modes = entity.Modes
	}
	if entity.Name != "" {
		output.Name = entity.Name
	}
	if entity.PayloadAvailable != "" {
		output.PayloadAvailable = entity.PayloadAvailable
	}
	if entity.PayloadNotAvailable != "" {
		output.PayloadNotAvailable = entity.PayloadNotAvailable
	}
	if entity.PayloadOff != "" {
		output.PayloadOff = entity.PayloadOff
	}
	if entity.PayloadOn != "" {
		output.PayloadOn = entity.PayloadOn
	}
	if entity.PowerCommandTopic != "" {
		output.PowerCommandTopic = entity.PowerCommandTopic
	}
	if entity.Precision != 0 {
		output.Precision = entity.Precision
	}
	if entity.Qos != 0 {
		output.Qos = entity.Qos
	}
	if entity.Retain != false {
		output.Retain = entity.Retain
	}
	if entity.SendIfOff != false {
		output.SendIfOff = entity.SendIfOff
	}
	if entity.SwingModeCommandTemplate != "" {
		output.SwingModeCommandTemplate = entity.SwingModeCommandTemplate
	}
	if entity.SwingModeCommandTopic != "" {
		output.SwingModeCommandTopic = entity.SwingModeCommandTopic
	}
	if entity.SwingModeStateTemplate != "" {
		output.SwingModeStateTemplate = entity.SwingModeStateTemplate
	}
	if entity.SwingModeStateTopic != "" {
		output.SwingModeStateTopic = entity.SwingModeStateTopic
	}
	if entity.SwingModes != nil {
		output.SwingModes = entity.SwingModes
	}
	if entity.TempStep != 0 {
		output.TempStep = entity.TempStep
	}
	if entity.TemperatureCommandTemplate != "" {
		output.TemperatureCommandTemplate = entity.TemperatureCommandTemplate
	}
	if entity.TemperatureCommandTopic != "" {
		output.TemperatureCommandTopic = entity.TemperatureCommandTopic
	}
	if entity.TemperatureHighCommandTemplate != "" {
		output.TemperatureHighCommandTemplate = entity.TemperatureHighCommandTemplate
	}
	if entity.TemperatureHighCommandTopic != "" {
		output.TemperatureHighCommandTopic = entity.TemperatureHighCommandTopic
	}
	if entity.TemperatureHighStateTemplate != "" {
		output.TemperatureHighStateTemplate = entity.TemperatureHighStateTemplate
	}
	if entity.TemperatureHighStateTopic != "" {
		output.TemperatureHighStateTopic = entity.TemperatureHighStateTopic
	}
	if entity.TemperatureLowCommandTemplate != "" {
		output.TemperatureLowCommandTemplate = entity.TemperatureLowCommandTemplate
	}
	if entity.TemperatureLowCommandTopic != "" {
		output.TemperatureLowCommandTopic = entity.TemperatureLowCommandTopic
	}
	if entity.TemperatureLowStateTemplate != "" {
		output.TemperatureLowStateTemplate = entity.TemperatureLowStateTemplate
	}
	if entity.TemperatureLowStateTopic != "" {
		output.TemperatureLowStateTopic = entity.TemperatureLowStateTopic
	}
	if entity.TemperatureStateTemplate != "" {
		output.TemperatureStateTemplate = entity.TemperatureStateTemplate
	}
	if entity.TemperatureStateTopic != "" {
		output.TemperatureStateTopic = entity.TemperatureStateTopic
	}
	if entity.TemperatureUnit != "" {
		output.TemperatureUnit = entity.TemperatureUnit
	}
	if entity.UniqueId != "" {
		output.UniqueId = entity.UniqueId
	}
	if entity.ValueTemplate != "" {
		output.ValueTemplate = entity.ValueTemplate
	}
	return
}
func (entity HADeviceLight) GenerateLight() (output HADeviceLight) {
	if entity.Availability.PayloadAvailable != "" {
		output.Availability.PayloadAvailable = entity.Availability.PayloadAvailable
	}
	if entity.Availability.PayloadNotAvailable != "" {
		output.Availability.PayloadNotAvailable = entity.Availability.PayloadNotAvailable
	}
	if entity.Availability.Topic != "" {
		output.Availability.Topic = entity.Availability.Topic
	} else {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.AvailabilityMode != "" {
		output.AvailabilityMode = entity.AvailabilityMode
	}
	if entity.AvailabilityTopic != "" {
		output.AvailabilityTopic = entity.AvailabilityTopic
	}
	if entity.BrightnessCommandTopic != "" {
		output.BrightnessCommandTopic = entity.BrightnessCommandTopic
	}
	if entity.BrightnessScale != 0 {
		output.BrightnessScale = entity.BrightnessScale
	}
	if entity.BrightnessStateTopic != "" {
		output.BrightnessStateTopic = entity.BrightnessStateTopic
	}
	if entity.BrightnessValueTemplate != "" {
		output.BrightnessValueTemplate = entity.BrightnessValueTemplate
	}
	if entity.ColorModeStateTopic != "" {
		output.ColorModeStateTopic = entity.ColorModeStateTopic
	}
	if entity.ColorModeValueTemplate != "" {
		output.ColorModeValueTemplate = entity.ColorModeValueTemplate
	}
	if entity.ColorTempCommandTemplate != "" {
		output.ColorTempCommandTemplate = entity.ColorTempCommandTemplate
	}
	if entity.ColorTempCommandTopic != "" {
		output.ColorTempCommandTopic = entity.ColorTempCommandTopic
	}
	if entity.ColorTempStateTopic != "" {
		output.ColorTempStateTopic = entity.ColorTempStateTopic
	}
	if entity.ColorTempValueTemplate != "" {
		output.ColorTempValueTemplate = entity.ColorTempValueTemplate
	}
	if entity.CommandTopic != "" {
		output.CommandTopic = entity.CommandTopic
	} else {
		logging.LogError("entity.CommandTopic generator not found, but field is required!")
	}
	if entity.Device.Connections != nil {
		output.Device.Connections = entity.Device.Connections
	}
	if entity.Device.Identifiers != nil {
		output.Device.Identifiers = entity.Device.Identifiers
	}
	if entity.Device.Manufacturer != "" {
		output.Device.Manufacturer = entity.Device.Manufacturer
	}
	if entity.Device.Model != "" {
		output.Device.Model = entity.Device.Model
	}
	if entity.Device.Name != "" {
		output.Device.Name = entity.Device.Name
	}
	if entity.Device.SuggestedArea != "" {
		output.Device.SuggestedArea = entity.Device.SuggestedArea
	}
	if entity.Device.SwVersion != "" {
		output.Device.SwVersion = entity.Device.SwVersion
	}
	if entity.Device.ViaDevice != "" {
		output.Device.ViaDevice = entity.Device.ViaDevice
	}
	if entity.EffectCommandTopic != "" {
		output.EffectCommandTopic = entity.EffectCommandTopic
	}
	if entity.EffectList != nil {
		output.EffectList = entity.EffectList
	}
	if entity.EffectStateTopic != "" {
		output.EffectStateTopic = entity.EffectStateTopic
	}
	if entity.EffectValueTemplate != "" {
		output.EffectValueTemplate = entity.EffectValueTemplate
	}
	if entity.EnabledByDefault != false {
		output.EnabledByDefault = entity.EnabledByDefault
	}
	if entity.HsCommandTopic != "" {
		output.HsCommandTopic = entity.HsCommandTopic
	}
	if entity.HsStateTopic != "" {
		output.HsStateTopic = entity.HsStateTopic
	}
	if entity.HsValueTemplate != "" {
		output.HsValueTemplate = entity.HsValueTemplate
	}
	if entity.Icon != "" {
		output.Icon = entity.Icon
	}
	if entity.JsonAttributesTemplate != "" {
		output.JsonAttributesTemplate = entity.JsonAttributesTemplate
	}
	if entity.JsonAttributesTopic != "" {
		output.JsonAttributesTopic = entity.JsonAttributesTopic
	}
	if entity.MaxMireds != 0 {
		output.MaxMireds = entity.MaxMireds
	}
	if entity.MinMireds != 0 {
		output.MinMireds = entity.MinMireds
	}
	if entity.Name != "" {
		output.Name = entity.Name
	}
	if entity.OnCommandType != "" {
		output.OnCommandType = entity.OnCommandType
	}
	if entity.Optimistic != false {
		output.Optimistic = entity.Optimistic
	}
	if entity.PayloadAvailable != "" {
		output.PayloadAvailable = entity.PayloadAvailable
	}
	if entity.PayloadNotAvailable != "" {
		output.PayloadNotAvailable = entity.PayloadNotAvailable
	}
	if entity.PayloadOff != "" {
		output.PayloadOff = entity.PayloadOff
	}
	if entity.PayloadOn != "" {
		output.PayloadOn = entity.PayloadOn
	}
	if entity.Qos != 0 {
		output.Qos = entity.Qos
	}
	if entity.Retain != false {
		output.Retain = entity.Retain
	}
	if entity.RgbCommandTemplate != "" {
		output.RgbCommandTemplate = entity.RgbCommandTemplate
	}
	if entity.RgbCommandTopic != "" {
		output.RgbCommandTopic = entity.RgbCommandTopic
	}
	if entity.RgbStateTopic != "" {
		output.RgbStateTopic = entity.RgbStateTopic
	}
	if entity.RgbValueTemplate != "" {
		output.RgbValueTemplate = entity.RgbValueTemplate
	}
	if entity.Schema != "" {
		output.Schema = entity.Schema
	}
	if entity.StateTopic != "" {
		output.StateTopic = entity.StateTopic
	}
	if entity.StateValueTemplate != "" {
		output.StateValueTemplate = entity.StateValueTemplate
	}
	if entity.UniqueId != "" {
		output.UniqueId = entity.UniqueId
	}
	if entity.WhiteCommandTopic != "" {
		output.WhiteCommandTopic = entity.WhiteCommandTopic
	}
	if entity.WhiteScale != 0 {
		output.WhiteScale = entity.WhiteScale
	}
	if entity.XyCommandTopic != "" {
		output.XyCommandTopic = entity.XyCommandTopic
	}
	if entity.XyStateTopic != "" {
		output.XyStateTopic = entity.XyStateTopic
	}
	if entity.XyValueTemplate != "" {
		output.XyValueTemplate = entity.XyValueTemplate
	}
	return
}
func (entity HADeviceLock) GenerateLock() (output HADeviceLock) {
	if entity.Availability.PayloadAvailable != "" {
		output.Availability.PayloadAvailable = entity.Availability.PayloadAvailable
	}
	if entity.Availability.PayloadNotAvailable != "" {
		output.Availability.PayloadNotAvailable = entity.Availability.PayloadNotAvailable
	}
	if entity.Availability.Topic != "" {
		output.Availability.Topic = entity.Availability.Topic
	} else {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.AvailabilityMode != "" {
		output.AvailabilityMode = entity.AvailabilityMode
	}
	if entity.AvailabilityTopic != "" {
		output.AvailabilityTopic = entity.AvailabilityTopic
	}
	if entity.CommandTopic != "" {
		output.CommandTopic = entity.CommandTopic
	} else {
		logging.LogError("entity.CommandTopic generator not found, but field is required!")
	}
	if entity.Device.Connections != nil {
		output.Device.Connections = entity.Device.Connections
	}
	if entity.Device.Identifiers != nil {
		output.Device.Identifiers = entity.Device.Identifiers
	}
	if entity.Device.Manufacturer != "" {
		output.Device.Manufacturer = entity.Device.Manufacturer
	}
	if entity.Device.Model != "" {
		output.Device.Model = entity.Device.Model
	}
	if entity.Device.Name != "" {
		output.Device.Name = entity.Device.Name
	}
	if entity.Device.SuggestedArea != "" {
		output.Device.SuggestedArea = entity.Device.SuggestedArea
	}
	if entity.Device.SwVersion != "" {
		output.Device.SwVersion = entity.Device.SwVersion
	}
	if entity.Device.ViaDevice != "" {
		output.Device.ViaDevice = entity.Device.ViaDevice
	}
	if entity.EnabledByDefault != false {
		output.EnabledByDefault = entity.EnabledByDefault
	}
	if entity.Icon != "" {
		output.Icon = entity.Icon
	}
	if entity.JsonAttributesTemplate != "" {
		output.JsonAttributesTemplate = entity.JsonAttributesTemplate
	}
	if entity.JsonAttributesTopic != "" {
		output.JsonAttributesTopic = entity.JsonAttributesTopic
	}
	if entity.Name != "" {
		output.Name = entity.Name
	}
	if entity.Optimistic != false {
		output.Optimistic = entity.Optimistic
	}
	if entity.PayloadAvailable != "" {
		output.PayloadAvailable = entity.PayloadAvailable
	}
	if entity.PayloadLock != "" {
		output.PayloadLock = entity.PayloadLock
	}
	if entity.PayloadNotAvailable != "" {
		output.PayloadNotAvailable = entity.PayloadNotAvailable
	}
	if entity.PayloadUnlock != "" {
		output.PayloadUnlock = entity.PayloadUnlock
	}
	if entity.Qos != 0 {
		output.Qos = entity.Qos
	}
	if entity.Retain != false {
		output.Retain = entity.Retain
	}
	if entity.StateLocked != "" {
		output.StateLocked = entity.StateLocked
	}
	if entity.StateTopic != "" {
		output.StateTopic = entity.StateTopic
	}
	if entity.StateUnlocked != "" {
		output.StateUnlocked = entity.StateUnlocked
	}
	if entity.UniqueId != "" {
		output.UniqueId = entity.UniqueId
	}
	if entity.ValueTemplate != "" {
		output.ValueTemplate = entity.ValueTemplate
	}
	return
}
func (entity HADeviceNumber) GenerateNumber() (output HADeviceNumber) {
	if entity.Availability.PayloadAvailable != "" {
		output.Availability.PayloadAvailable = entity.Availability.PayloadAvailable
	}
	if entity.Availability.PayloadNotAvailable != "" {
		output.Availability.PayloadNotAvailable = entity.Availability.PayloadNotAvailable
	}
	if entity.Availability.Topic != "" {
		output.Availability.Topic = entity.Availability.Topic
	} else {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.AvailabilityMode != "" {
		output.AvailabilityMode = entity.AvailabilityMode
	}
	if entity.AvailabilityTopic != "" {
		output.AvailabilityTopic = entity.AvailabilityTopic
	}
	if entity.CommandTopic != "" {
		output.CommandTopic = entity.CommandTopic
	}
	if entity.Device.Connections != nil {
		output.Device.Connections = entity.Device.Connections
	}
	if entity.Device.Identifiers != nil {
		output.Device.Identifiers = entity.Device.Identifiers
	}
	if entity.Device.Manufacturer != "" {
		output.Device.Manufacturer = entity.Device.Manufacturer
	}
	if entity.Device.Model != "" {
		output.Device.Model = entity.Device.Model
	}
	if entity.Device.Name != "" {
		output.Device.Name = entity.Device.Name
	}
	if entity.Device.SuggestedArea != "" {
		output.Device.SuggestedArea = entity.Device.SuggestedArea
	}
	if entity.Device.SwVersion != "" {
		output.Device.SwVersion = entity.Device.SwVersion
	}
	if entity.Device.ViaDevice != "" {
		output.Device.ViaDevice = entity.Device.ViaDevice
	}
	if entity.EnabledByDefault != false {
		output.EnabledByDefault = entity.EnabledByDefault
	}
	if entity.Icon != "" {
		output.Icon = entity.Icon
	}
	if entity.JsonAttributesTemplate != "" {
		output.JsonAttributesTemplate = entity.JsonAttributesTemplate
	}
	if entity.JsonAttributesTopic != "" {
		output.JsonAttributesTopic = entity.JsonAttributesTopic
	}
	if entity.Max != 0 {
		output.Max = entity.Max
	}
	if entity.Min != 0 {
		output.Min = entity.Min
	}
	if entity.Name != "" {
		output.Name = entity.Name
	}
	if entity.Optimistic != false {
		output.Optimistic = entity.Optimistic
	}
	if entity.Qos != 0 {
		output.Qos = entity.Qos
	}
	if entity.Retain != false {
		output.Retain = entity.Retain
	}
	if entity.StateTopic != "" {
		output.StateTopic = entity.StateTopic
	}
	if entity.Step != 0 {
		output.Step = entity.Step
	}
	if entity.UniqueId != "" {
		output.UniqueId = entity.UniqueId
	}
	if entity.ValueTemplate != "" {
		output.ValueTemplate = entity.ValueTemplate
	}
	return
}
func (entity HADeviceScene) GenerateScene() (output HADeviceScene) {
	if entity.Availability.PayloadAvailable != "" {
		output.Availability.PayloadAvailable = entity.Availability.PayloadAvailable
	}
	if entity.Availability.PayloadNotAvailable != "" {
		output.Availability.PayloadNotAvailable = entity.Availability.PayloadNotAvailable
	}
	if entity.Availability.Topic != "" {
		output.Availability.Topic = entity.Availability.Topic
	} else {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.AvailabilityMode != "" {
		output.AvailabilityMode = entity.AvailabilityMode
	}
	if entity.AvailabilityTopic != "" {
		output.AvailabilityTopic = entity.AvailabilityTopic
	}
	if entity.CommandTopic != "" {
		output.CommandTopic = entity.CommandTopic
	}
	if entity.EnabledByDefault != false {
		output.EnabledByDefault = entity.EnabledByDefault
	}
	if entity.Icon != "" {
		output.Icon = entity.Icon
	}
	if entity.Name != "" {
		output.Name = entity.Name
	}
	if entity.PayloadAvailable != "" {
		output.PayloadAvailable = entity.PayloadAvailable
	}
	if entity.PayloadNotAvailable != "" {
		output.PayloadNotAvailable = entity.PayloadNotAvailable
	}
	if entity.PayloadOn != "" {
		output.PayloadOn = entity.PayloadOn
	}
	if entity.Qos != 0 {
		output.Qos = entity.Qos
	}
	if entity.Retain != false {
		output.Retain = entity.Retain
	}
	if entity.UniqueId != "" {
		output.UniqueId = entity.UniqueId
	}
	return
}
func (entity HADeviceSelect) GenerateSelect() (output HADeviceSelect) {
	if entity.Availability.PayloadAvailable != "" {
		output.Availability.PayloadAvailable = entity.Availability.PayloadAvailable
	}
	if entity.Availability.PayloadNotAvailable != "" {
		output.Availability.PayloadNotAvailable = entity.Availability.PayloadNotAvailable
	}
	if entity.Availability.Topic != "" {
		output.Availability.Topic = entity.Availability.Topic
	} else {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.AvailabilityMode != "" {
		output.AvailabilityMode = entity.AvailabilityMode
	}
	if entity.AvailabilityTopic != "" {
		output.AvailabilityTopic = entity.AvailabilityTopic
	}
	if entity.CommandTopic != "" {
		output.CommandTopic = entity.CommandTopic
	} else {
		logging.LogError("entity.CommandTopic generator not found, but field is required!")
	}
	if entity.Device.Connections != nil {
		output.Device.Connections = entity.Device.Connections
	}
	if entity.Device.Identifiers != nil {
		output.Device.Identifiers = entity.Device.Identifiers
	}
	if entity.Device.Manufacturer != "" {
		output.Device.Manufacturer = entity.Device.Manufacturer
	}
	if entity.Device.Model != "" {
		output.Device.Model = entity.Device.Model
	}
	if entity.Device.Name != "" {
		output.Device.Name = entity.Device.Name
	}
	if entity.Device.SuggestedArea != "" {
		output.Device.SuggestedArea = entity.Device.SuggestedArea
	}
	if entity.Device.SwVersion != "" {
		output.Device.SwVersion = entity.Device.SwVersion
	}
	if entity.Device.ViaDevice != "" {
		output.Device.ViaDevice = entity.Device.ViaDevice
	}
	if entity.EnabledByDefault != false {
		output.EnabledByDefault = entity.EnabledByDefault
	}
	if entity.Icon != "" {
		output.Icon = entity.Icon
	}
	if entity.JsonAttributesTemplate != "" {
		output.JsonAttributesTemplate = entity.JsonAttributesTemplate
	}
	if entity.JsonAttributesTopic != "" {
		output.JsonAttributesTopic = entity.JsonAttributesTopic
	}
	if entity.Name != "" {
		output.Name = entity.Name
	}
	if entity.Optimistic != false {
		output.Optimistic = entity.Optimistic
	}
	if entity.Options != nil {
		output.Options = entity.Options
	} else {
		logging.LogError("entity.Options generator not found, but field is required!")
	}
	if entity.Qos != 0 {
		output.Qos = entity.Qos
	}
	if entity.Retain != false {
		output.Retain = entity.Retain
	}
	if entity.StateTopic != "" {
		output.StateTopic = entity.StateTopic
	}
	if entity.UniqueId != "" {
		output.UniqueId = entity.UniqueId
	}
	if entity.ValueTemplate != "" {
		output.ValueTemplate = entity.ValueTemplate
	}
	return
}
func (entity HADeviceSensor) GenerateSensor() (output HADeviceSensor) {
	if entity.Availability.PayloadAvailable != "" {
		output.Availability.PayloadAvailable = entity.Availability.PayloadAvailable
	}
	if entity.Availability.PayloadNotAvailable != "" {
		output.Availability.PayloadNotAvailable = entity.Availability.PayloadNotAvailable
	}
	if entity.Availability.Topic != "" {
		output.Availability.Topic = entity.Availability.Topic
	} else {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.AvailabilityMode != "" {
		output.AvailabilityMode = entity.AvailabilityMode
	}
	if entity.AvailabilityTopic != "" {
		output.AvailabilityTopic = entity.AvailabilityTopic
	}
	if entity.Device.Connections != nil {
		output.Device.Connections = entity.Device.Connections
	}
	if entity.Device.Identifiers != nil {
		output.Device.Identifiers = entity.Device.Identifiers
	}
	if entity.Device.Manufacturer != "" {
		output.Device.Manufacturer = entity.Device.Manufacturer
	}
	if entity.Device.Model != "" {
		output.Device.Model = entity.Device.Model
	}
	if entity.Device.Name != "" {
		output.Device.Name = entity.Device.Name
	}
	if entity.Device.SuggestedArea != "" {
		output.Device.SuggestedArea = entity.Device.SuggestedArea
	}
	if entity.Device.SwVersion != "" {
		output.Device.SwVersion = entity.Device.SwVersion
	}
	if entity.Device.ViaDevice != "" {
		output.Device.ViaDevice = entity.Device.ViaDevice
	}
	if entity.DeviceClass != "" {
		output.DeviceClass = entity.DeviceClass
	}
	if entity.EnabledByDefault != false {
		output.EnabledByDefault = entity.EnabledByDefault
	}
	if entity.ExpireAfter != 0 {
		output.ExpireAfter = entity.ExpireAfter
	}
	if entity.ForceUpdate != false {
		output.ForceUpdate = entity.ForceUpdate
	}
	if entity.Icon != "" {
		output.Icon = entity.Icon
	}
	if entity.JsonAttributesTemplate != "" {
		output.JsonAttributesTemplate = entity.JsonAttributesTemplate
	}
	if entity.JsonAttributesTopic != "" {
		output.JsonAttributesTopic = entity.JsonAttributesTopic
	}
	if entity.LastResetTopic != "" {
		output.LastResetTopic = entity.LastResetTopic
	}
	if entity.LastResetValueTemplate != "" {
		output.LastResetValueTemplate = entity.LastResetValueTemplate
	}
	if entity.Name != "" {
		output.Name = entity.Name
	}
	if entity.PayloadAvailable != "" {
		output.PayloadAvailable = entity.PayloadAvailable
	}
	if entity.PayloadNotAvailable != "" {
		output.PayloadNotAvailable = entity.PayloadNotAvailable
	}
	if entity.Qos != 0 {
		output.Qos = entity.Qos
	}
	if entity.StateClass != "" {
		output.StateClass = entity.StateClass
	}
	if entity.StateTopic != "" {
		output.StateTopic = entity.StateTopic
	} else {
		logging.LogError("entity.StateTopic generator not found, but field is required!")
	}
	if entity.UniqueId != "" {
		output.UniqueId = entity.UniqueId
	}
	if entity.UnitOfMeasurement != "" {
		output.UnitOfMeasurement = entity.UnitOfMeasurement
	}
	if entity.ValueTemplate != "" {
		output.ValueTemplate = entity.ValueTemplate
	}
	return
}
func (entity HADeviceSwitch) GenerateSwitch() (output HADeviceSwitch) {
	if entity.Availability.PayloadAvailable != "" {
		output.Availability.PayloadAvailable = entity.Availability.PayloadAvailable
	}
	if entity.Availability.PayloadNotAvailable != "" {
		output.Availability.PayloadNotAvailable = entity.Availability.PayloadNotAvailable
	}
	if entity.Availability.Topic != "" {
		output.Availability.Topic = entity.Availability.Topic
	} else {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.AvailabilityMode != "" {
		output.AvailabilityMode = entity.AvailabilityMode
	}
	if entity.AvailabilityTopic != "" {
		output.AvailabilityTopic = entity.AvailabilityTopic
	}
	if entity.CommandTopic != "" {
		output.CommandTopic = entity.CommandTopic
	}
	if entity.Device.Connections != nil {
		output.Device.Connections = entity.Device.Connections
	}
	if entity.Device.Identifiers != nil {
		output.Device.Identifiers = entity.Device.Identifiers
	}
	if entity.Device.Manufacturer != "" {
		output.Device.Manufacturer = entity.Device.Manufacturer
	}
	if entity.Device.Model != "" {
		output.Device.Model = entity.Device.Model
	}
	if entity.Device.Name != "" {
		output.Device.Name = entity.Device.Name
	}
	if entity.Device.SuggestedArea != "" {
		output.Device.SuggestedArea = entity.Device.SuggestedArea
	}
	if entity.Device.SwVersion != "" {
		output.Device.SwVersion = entity.Device.SwVersion
	}
	if entity.Device.ViaDevice != "" {
		output.Device.ViaDevice = entity.Device.ViaDevice
	}
	if entity.EnabledByDefault != false {
		output.EnabledByDefault = entity.EnabledByDefault
	}
	if entity.Icon != "" {
		output.Icon = entity.Icon
	}
	if entity.JsonAttributesTemplate != "" {
		output.JsonAttributesTemplate = entity.JsonAttributesTemplate
	}
	if entity.JsonAttributesTopic != "" {
		output.JsonAttributesTopic = entity.JsonAttributesTopic
	}
	if entity.Name != "" {
		output.Name = entity.Name
	}
	if entity.Optimistic != false {
		output.Optimistic = entity.Optimistic
	}
	if entity.PayloadAvailable != "" {
		output.PayloadAvailable = entity.PayloadAvailable
	}
	if entity.PayloadNotAvailable != "" {
		output.PayloadNotAvailable = entity.PayloadNotAvailable
	}
	if entity.PayloadOff != "" {
		output.PayloadOff = entity.PayloadOff
	}
	if entity.PayloadOn != "" {
		output.PayloadOn = entity.PayloadOn
	}
	if entity.Qos != 0 {
		output.Qos = entity.Qos
	}
	if entity.Retain != false {
		output.Retain = entity.Retain
	}
	if entity.StateOff != "" {
		output.StateOff = entity.StateOff
	}
	if entity.StateOn != "" {
		output.StateOn = entity.StateOn
	}
	if entity.StateTopic != "" {
		output.StateTopic = entity.StateTopic
	}
	if entity.UniqueId != "" {
		output.UniqueId = entity.UniqueId
	}
	if entity.ValueTemplate != "" {
		output.ValueTemplate = entity.ValueTemplate
	}
	return
}
func (entity HADeviceTag) GenerateTag() (output HADeviceTag) {
	if entity.Device.Connections != nil {
		output.Device.Connections = entity.Device.Connections
	}
	if entity.Device.Identifiers != nil {
		output.Device.Identifiers = entity.Device.Identifiers
	}
	if entity.Device.Manufacturer != "" {
		output.Device.Manufacturer = entity.Device.Manufacturer
	}
	if entity.Device.Model != "" {
		output.Device.Model = entity.Device.Model
	}
	if entity.Device.Name != "" {
		output.Device.Name = entity.Device.Name
	}
	if entity.Device.SuggestedArea != "" {
		output.Device.SuggestedArea = entity.Device.SuggestedArea
	}
	if entity.Device.SwVersion != "" {
		output.Device.SwVersion = entity.Device.SwVersion
	}
	if entity.Device.ViaDevice != "" {
		output.Device.ViaDevice = entity.Device.ViaDevice
	}
	if entity.Topic != "" {
		output.Topic = entity.Topic
	} else {
		logging.LogError("entity.Topic generator not found, but field is required!")
	}
	if entity.ValueTemplate != "" {
		output.ValueTemplate = entity.ValueTemplate
	}
	return
}
func (entity HADeviceVacuum) GenerateVacuum() (output HADeviceVacuum) {
	if entity.Availability.PayloadAvailable != "" {
		output.Availability.PayloadAvailable = entity.Availability.PayloadAvailable
	}
	if entity.Availability.PayloadNotAvailable != "" {
		output.Availability.PayloadNotAvailable = entity.Availability.PayloadNotAvailable
	}
	if entity.Availability.Topic != "" {
		output.Availability.Topic = entity.Availability.Topic
	} else {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.AvailabilityMode != "" {
		output.AvailabilityMode = entity.AvailabilityMode
	}
	if entity.AvailabilityTopic != "" {
		output.AvailabilityTopic = entity.AvailabilityTopic
	}
	if entity.BatteryLevelTemplate != "" {
		output.BatteryLevelTemplate = entity.BatteryLevelTemplate
	}
	if entity.BatteryLevelTopic != "" {
		output.BatteryLevelTopic = entity.BatteryLevelTopic
	}
	if entity.ChargingTemplate != "" {
		output.ChargingTemplate = entity.ChargingTemplate
	}
	if entity.ChargingTopic != "" {
		output.ChargingTopic = entity.ChargingTopic
	}
	if entity.CleaningTemplate != "" {
		output.CleaningTemplate = entity.CleaningTemplate
	}
	if entity.CleaningTopic != "" {
		output.CleaningTopic = entity.CleaningTopic
	}
	if entity.CommandTopic != "" {
		output.CommandTopic = entity.CommandTopic
	}
	if entity.DockedTemplate != "" {
		output.DockedTemplate = entity.DockedTemplate
	}
	if entity.DockedTopic != "" {
		output.DockedTopic = entity.DockedTopic
	}
	if entity.EnabledByDefault != false {
		output.EnabledByDefault = entity.EnabledByDefault
	}
	if entity.ErrorTemplate != "" {
		output.ErrorTemplate = entity.ErrorTemplate
	}
	if entity.ErrorTopic != "" {
		output.ErrorTopic = entity.ErrorTopic
	}
	if entity.FanSpeedList != nil {
		output.FanSpeedList = entity.FanSpeedList
	}
	if entity.FanSpeedTemplate != "" {
		output.FanSpeedTemplate = entity.FanSpeedTemplate
	}
	if entity.FanSpeedTopic != "" {
		output.FanSpeedTopic = entity.FanSpeedTopic
	}
	if entity.Icon != "" {
		output.Icon = entity.Icon
	}
	if entity.JsonAttributesTemplate != "" {
		output.JsonAttributesTemplate = entity.JsonAttributesTemplate
	}
	if entity.JsonAttributesTopic != "" {
		output.JsonAttributesTopic = entity.JsonAttributesTopic
	}
	if entity.Name != "" {
		output.Name = entity.Name
	}
	if entity.PayloadAvailable != "" {
		output.PayloadAvailable = entity.PayloadAvailable
	}
	if entity.PayloadCleanSpot != "" {
		output.PayloadCleanSpot = entity.PayloadCleanSpot
	}
	if entity.PayloadLocate != "" {
		output.PayloadLocate = entity.PayloadLocate
	}
	if entity.PayloadNotAvailable != "" {
		output.PayloadNotAvailable = entity.PayloadNotAvailable
	}
	if entity.PayloadReturnToBase != "" {
		output.PayloadReturnToBase = entity.PayloadReturnToBase
	}
	if entity.PayloadStartPause != "" {
		output.PayloadStartPause = entity.PayloadStartPause
	}
	if entity.PayloadStop != "" {
		output.PayloadStop = entity.PayloadStop
	}
	if entity.PayloadTurnOff != "" {
		output.PayloadTurnOff = entity.PayloadTurnOff
	}
	if entity.PayloadTurnOn != "" {
		output.PayloadTurnOn = entity.PayloadTurnOn
	}
	if entity.Qos != 0 {
		output.Qos = entity.Qos
	}
	if entity.Retain != false {
		output.Retain = entity.Retain
	}
	if entity.Schema != "" {
		output.Schema = entity.Schema
	}
	if entity.SendCommandTopic != "" {
		output.SendCommandTopic = entity.SendCommandTopic
	}
	if entity.SetFanSpeedTopic != "" {
		output.SetFanSpeedTopic = entity.SetFanSpeedTopic
	}
	if entity.SupportedFeatures != nil {
		output.SupportedFeatures = entity.SupportedFeatures
	}
	if entity.UniqueId != "" {
		output.UniqueId = entity.UniqueId
	}
	return
}
