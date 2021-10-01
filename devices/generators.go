package devices

import (
	"github.com/W-Floyd/ha-mqtt-iot/logging"
	"github.com/clarketm/json"
)

func DeepCopy(a, b interface{}) {
	byt, _ := json.Marshal(a)
	json.Unmarshal(byt, b)
}
func (entity HADeviceAlarmControlPanel) Generate() (output HADeviceAlarmControlPanel) {
	DeepCopy(&entity, &output)
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
	DeepCopy(&entity, &output)
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.StateTopic == nil {
		logging.LogError("entity.StateTopic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceCamera) Generate() (output HADeviceCamera) {
	DeepCopy(&entity, &output)
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.Topic == nil {
		logging.LogError("entity.Topic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceCover) Generate() (output HADeviceCover) {
	DeepCopy(&entity, &output)
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceDeviceTracker) Generate() (output HADeviceDeviceTracker) {
	DeepCopy(&entity, &output)
	if entity.Devices == nil {
		logging.LogError("entity.Devices generator not found, but field is required!")
	}
	return
}
func (entity HADeviceDeviceTrigger) Generate() (output HADeviceDeviceTrigger) {
	DeepCopy(&entity, &output)
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
	DeepCopy(&entity, &output)
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.CommandTopic == nil {
		logging.LogError("entity.CommandTopic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceHumidifier) Generate() (output HADeviceHumidifier) {
	DeepCopy(&entity, &output)
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
	DeepCopy(&entity, &output)
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceLight) Generate() (output HADeviceLight) {
	DeepCopy(&entity, &output)
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.CommandTopic == nil {
		logging.LogError("entity.CommandTopic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceLock) Generate() (output HADeviceLock) {
	DeepCopy(&entity, &output)
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.CommandTopic == nil {
		logging.LogError("entity.CommandTopic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceNumber) Generate() (output HADeviceNumber) {
	DeepCopy(&entity, &output)
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceScene) Generate() (output HADeviceScene) {
	DeepCopy(&entity, &output)
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceSelect) Generate() (output HADeviceSelect) {
	DeepCopy(&entity, &output)
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
	DeepCopy(&entity, &output)
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	if entity.StateTopic == nil {
		logging.LogError("entity.StateTopic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceSwitch) Generate() (output HADeviceSwitch) {
	DeepCopy(&entity, &output)
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceTag) Generate() (output HADeviceTag) {
	DeepCopy(&entity, &output)
	if entity.Topic == nil {
		logging.LogError("entity.Topic generator not found, but field is required!")
	}
	return
}
func (entity HADeviceVacuum) Generate() (output HADeviceVacuum) {
	DeepCopy(&entity, &output)
	if entity.Availability.Topic == nil {
		logging.LogError("entity.Availability.Topic generator not found, but field is required!")
	}
	return
}
