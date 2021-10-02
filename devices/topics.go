package devices

import (
	"errors"
	"github.com/W-Floyd/ha-mqtt-iot/common"
	"github.com/iancoleman/strcase"
)

func (component HADeviceAlarmControlPanel) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/alarm_control_panel/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceAlarmControlPanel) CanGenerateTopic() (error, bool) {
	if component.Name == nil {
		return errors.New("component.Name"), false
	}
	if common.NodeID == "" {
		return errors.New("common.NodeID"), false
	}
	if common.DiscoveryPrefix == "" {
		return errors.New("common.DiscoveryPrefix"), false
	}
	return nil, true
}
func (component HADeviceBinarySensor) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/binary_sensor/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceBinarySensor) CanGenerateTopic() (error, bool) {
	if component.Name == nil {
		return errors.New("component.Name"), false
	}
	if common.NodeID == "" {
		return errors.New("common.NodeID"), false
	}
	if common.DiscoveryPrefix == "" {
		return errors.New("common.DiscoveryPrefix"), false
	}
	return nil, true
}
func (component HADeviceCamera) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/camera/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceCamera) CanGenerateTopic() (error, bool) {
	if component.Name == nil {
		return errors.New("component.Name"), false
	}
	if common.NodeID == "" {
		return errors.New("common.NodeID"), false
	}
	if common.DiscoveryPrefix == "" {
		return errors.New("common.DiscoveryPrefix"), false
	}
	return nil, true
}
func (component HADeviceCover) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/cover/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceCover) CanGenerateTopic() (error, bool) {
	if component.Name == nil {
		return errors.New("component.Name"), false
	}
	if common.NodeID == "" {
		return errors.New("common.NodeID"), false
	}
	if common.DiscoveryPrefix == "" {
		return errors.New("common.DiscoveryPrefix"), false
	}
	return nil, true
}
func (component HADeviceDeviceTracker) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/device_tracker/" + common.NodeID + "/"
	return
}
func (component HADeviceDeviceTracker) CanGenerateTopic() (error, bool) {
	if common.NodeID == "" {
		return errors.New("common.NodeID"), false
	}
	if common.DiscoveryPrefix == "" {
		return errors.New("common.DiscoveryPrefix"), false
	}
	return nil, true
}
func (component HADeviceDeviceTrigger) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/device_trigger/" + common.NodeID + "/"
	return
}
func (component HADeviceDeviceTrigger) CanGenerateTopic() (error, bool) {
	if common.NodeID == "" {
		return errors.New("common.NodeID"), false
	}
	if common.DiscoveryPrefix == "" {
		return errors.New("common.DiscoveryPrefix"), false
	}
	return nil, true
}
func (component HADeviceFan) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/fan/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceFan) CanGenerateTopic() (error, bool) {
	if component.Name == nil {
		return errors.New("component.Name"), false
	}
	if common.NodeID == "" {
		return errors.New("common.NodeID"), false
	}
	if common.DiscoveryPrefix == "" {
		return errors.New("common.DiscoveryPrefix"), false
	}
	return nil, true
}
func (component HADeviceHumidifier) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/humidifier/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceHumidifier) CanGenerateTopic() (error, bool) {
	if component.Name == nil {
		return errors.New("component.Name"), false
	}
	if common.NodeID == "" {
		return errors.New("common.NodeID"), false
	}
	if common.DiscoveryPrefix == "" {
		return errors.New("common.DiscoveryPrefix"), false
	}
	return nil, true
}
func (component HADeviceClimate) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/climate/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceClimate) CanGenerateTopic() (error, bool) {
	if component.Name == nil {
		return errors.New("component.Name"), false
	}
	if common.NodeID == "" {
		return errors.New("common.NodeID"), false
	}
	if common.DiscoveryPrefix == "" {
		return errors.New("common.DiscoveryPrefix"), false
	}
	return nil, true
}
func (component HADeviceLight) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/light/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceLight) CanGenerateTopic() (error, bool) {
	if component.Name == nil {
		return errors.New("component.Name"), false
	}
	if common.NodeID == "" {
		return errors.New("common.NodeID"), false
	}
	if common.DiscoveryPrefix == "" {
		return errors.New("common.DiscoveryPrefix"), false
	}
	return nil, true
}
func (component HADeviceLock) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/lock/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceLock) CanGenerateTopic() (error, bool) {
	if component.Name == nil {
		return errors.New("component.Name"), false
	}
	if common.NodeID == "" {
		return errors.New("common.NodeID"), false
	}
	if common.DiscoveryPrefix == "" {
		return errors.New("common.DiscoveryPrefix"), false
	}
	return nil, true
}
func (component HADeviceNumber) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/number/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceNumber) CanGenerateTopic() (error, bool) {
	if component.Name == nil {
		return errors.New("component.Name"), false
	}
	if common.NodeID == "" {
		return errors.New("common.NodeID"), false
	}
	if common.DiscoveryPrefix == "" {
		return errors.New("common.DiscoveryPrefix"), false
	}
	return nil, true
}
func (component HADeviceScene) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/scene/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceScene) CanGenerateTopic() (error, bool) {
	if component.Name == nil {
		return errors.New("component.Name"), false
	}
	if common.NodeID == "" {
		return errors.New("common.NodeID"), false
	}
	if common.DiscoveryPrefix == "" {
		return errors.New("common.DiscoveryPrefix"), false
	}
	return nil, true
}
func (component HADeviceSelect) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/select/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceSelect) CanGenerateTopic() (error, bool) {
	if component.Name == nil {
		return errors.New("component.Name"), false
	}
	if common.NodeID == "" {
		return errors.New("common.NodeID"), false
	}
	if common.DiscoveryPrefix == "" {
		return errors.New("common.DiscoveryPrefix"), false
	}
	return nil, true
}
func (component HADeviceSensor) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/sensor/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceSensor) CanGenerateTopic() (error, bool) {
	if component.Name == nil {
		return errors.New("component.Name"), false
	}
	if common.NodeID == "" {
		return errors.New("common.NodeID"), false
	}
	if common.DiscoveryPrefix == "" {
		return errors.New("common.DiscoveryPrefix"), false
	}
	return nil, true
}
func (component HADeviceSwitch) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/switch/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceSwitch) CanGenerateTopic() (error, bool) {
	if component.Name == nil {
		return errors.New("component.Name"), false
	}
	if common.NodeID == "" {
		return errors.New("common.NodeID"), false
	}
	if common.DiscoveryPrefix == "" {
		return errors.New("common.DiscoveryPrefix"), false
	}
	return nil, true
}
func (component HADeviceTag) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/tag/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Device.Name) + "/"
	return
}
func (component HADeviceTag) CanGenerateTopic() (error, bool) {
	if component.Device.Name == nil {
		return errors.New("component.Device.Name"), false
	}
	if common.NodeID == "" {
		return errors.New("common.NodeID"), false
	}
	if common.DiscoveryPrefix == "" {
		return errors.New("common.DiscoveryPrefix"), false
	}
	return nil, true
}
func (component HADeviceVacuum) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/vacuum/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceVacuum) CanGenerateTopic() (error, bool) {
	if component.Name == nil {
		return errors.New("component.Name"), false
	}
	if common.NodeID == "" {
		return errors.New("common.NodeID"), false
	}
	if common.DiscoveryPrefix == "" {
		return errors.New("common.DiscoveryPrefix"), false
	}
	return nil, true
}
