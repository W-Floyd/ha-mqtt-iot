package devices

import (
	"github.com/W-Floyd/ha-mqtt-iot/logging"
	"github.com/jinzhu/copier"
)

func (entity HADeviceAlarmControlPanel) Generate() (output HADeviceAlarmControlPanel) {
	copier.CopyWithOption(&output, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.CommandTopic == nil {
		logging.LogError("entity.CommandTopic generator not found, but field is required!")
	}
	if entity.StateTopic == nil {
		logging.LogError("entity.StateTopic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceBinarySensor) Generate() (output HADeviceBinarySensor) {
	copier.CopyWithOption(&output, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.StateTopic == nil {
		logging.LogError("entity.StateTopic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceCamera) Generate() (output HADeviceCamera) {
	copier.CopyWithOption(&output, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.Topic == nil {
		logging.LogError("entity.Topic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceCover) Generate() (output HADeviceCover) {
	copier.CopyWithOption(&output, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceDeviceTracker) Generate() (output HADeviceDeviceTracker) {
	copier.CopyWithOption(&output, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if entity.Devices == nil {
		logging.LogError("entity.Devices generator not found, but field is required!")
	}
	return
}
func (entity HADeviceDeviceTrigger) Generate() (output HADeviceDeviceTrigger) {
	copier.CopyWithOption(&output, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if entity.AutomationType == nil {
		logging.LogError("entity.AutomationType generator not found, but field is required!")
	}
	if entity.Subtype == nil {
		logging.LogError("entity.Subtype generator not found, but field is required!")
	}
	if entity.Topic == nil {
		logging.LogError("entity.Topic generator not found, but field is required!")
	}
	if entity.Type == nil {
		logging.LogError("entity.Type generator not found, but field is required!")
	}
	return
}
func (entity HADeviceFan) Generate() (output HADeviceFan) {
	copier.CopyWithOption(&output, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.CommandTopic == nil {
		logging.LogError("entity.CommandTopic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceHumidifier) Generate() (output HADeviceHumidifier) {
	copier.CopyWithOption(&output, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.CommandTopic == nil {
		logging.LogError("entity.CommandTopic generator not found, but field is required!")
	}
	if entity.TargetHumidityCommandTopic == nil {
		logging.LogError("entity.TargetHumidityCommandTopic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceClimate) Generate() (output HADeviceClimate) {
	copier.CopyWithOption(&output, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceLight) Generate() (output HADeviceLight) {
	copier.CopyWithOption(&output, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.CommandTopic == nil {
		logging.LogError("entity.CommandTopic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceLock) Generate() (output HADeviceLock) {
	copier.CopyWithOption(&output, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.CommandTopic == nil {
		logging.LogError("entity.CommandTopic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceNumber) Generate() (output HADeviceNumber) {
	copier.CopyWithOption(&output, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceScene) Generate() (output HADeviceScene) {
	copier.CopyWithOption(&output, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceSelect) Generate() (output HADeviceSelect) {
	copier.CopyWithOption(&output, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.CommandTopic == nil {
		logging.LogError("entity.CommandTopic generator not found, but field is required!")
	}
	if entity.Options == nil {
		logging.LogError("entity.Options generator not found, but field is required!")
	}
	return
}
func (entity HADeviceSensor) Generate() (output HADeviceSensor) {
	copier.CopyWithOption(&output, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.StateTopic == nil {
		logging.LogError("entity.StateTopic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceSwitch) Generate() (output HADeviceSwitch) {
	copier.CopyWithOption(&output, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceTag) Generate() (output HADeviceTag) {
	copier.CopyWithOption(&output, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if entity.Topic == nil {
		logging.LogError("entity.Topic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceVacuum) Generate() (output HADeviceVacuum) {
	copier.CopyWithOption(&output, &entity, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	return
}
