package devices

import (
	"github.com/W-Floyd/ha-mqtt-iot/common"
	"github.com/iancoleman/strcase"
)

func (component HADeviceAlarmControlPanel) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/alarm_control_panel/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceAlarmControlPanel) CanGenerateTopic() bool {
	if (component.Device.Name == nil && component.Name == nil) || common.DiscoveryPrefix == "" || common.NodeID == "" {
		return false
	}
	return true
}
func (component HADeviceBinarySensor) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/binary_sensor/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceBinarySensor) CanGenerateTopic() bool {
	if (component.Device.Name == nil && component.Name == nil) || common.DiscoveryPrefix == "" || common.NodeID == "" {
		return false
	}
	return true
}
func (component HADeviceCamera) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/camera/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceCamera) CanGenerateTopic() bool {
	if (component.Device.Name == nil && component.Name == nil) || common.DiscoveryPrefix == "" || common.NodeID == "" {
		return false
	}
	return true
}
func (component HADeviceCover) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/cover/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceCover) CanGenerateTopic() bool {
	if (component.Device.Name == nil && component.Name == nil) || common.DiscoveryPrefix == "" || common.NodeID == "" {
		return false
	}
	return true
}
func (component HADeviceDeviceTracker) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/device_tracker/" + common.NodeID + "/"
	return
}
func (component HADeviceDeviceTracker) CanGenerateTopic() bool {
	if common.NodeID == "" || common.DiscoveryPrefix == "" {
		return false
	}
	return true
}
func (component HADeviceDeviceTrigger) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/device_trigger/" + common.NodeID + "/"
	return
}
func (component HADeviceDeviceTrigger) CanGenerateTopic() bool {
	if common.NodeID == "" || common.DiscoveryPrefix == "" {
		return false
	}
	return true
}
func (component HADeviceFan) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/fan/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceFan) CanGenerateTopic() bool {
	if (component.Device.Name == nil && component.Name == nil) || common.DiscoveryPrefix == "" || common.NodeID == "" {
		return false
	}
	return true
}
func (component HADeviceHumidifier) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/humidifier/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceHumidifier) CanGenerateTopic() bool {
	if (component.Device.Name == nil && component.Name == nil) || common.DiscoveryPrefix == "" || common.NodeID == "" {
		return false
	}
	return true
}
func (component HADeviceClimate) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/climate/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceClimate) CanGenerateTopic() bool {
	if (component.Device.Name == nil && component.Name == nil) || common.DiscoveryPrefix == "" || common.NodeID == "" {
		return false
	}
	return true
}
func (component HADeviceLight) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/light/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceLight) CanGenerateTopic() bool {
	if (component.Device.Name == nil && component.Name == nil) || common.DiscoveryPrefix == "" || common.NodeID == "" {
		return false
	}
	return true
}
func (component HADeviceLock) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/lock/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceLock) CanGenerateTopic() bool {
	if (component.Device.Name == nil && component.Name == nil) || common.DiscoveryPrefix == "" || common.NodeID == "" {
		return false
	}
	return true
}
func (component HADeviceNumber) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/number/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceNumber) CanGenerateTopic() bool {
	if (component.Device.Name == nil && component.Name == nil) || common.DiscoveryPrefix == "" || common.NodeID == "" {
		return false
	}
	return true
}
func (component HADeviceScene) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/scene/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceScene) CanGenerateTopic() bool {
	if component.Name == nil || common.NodeID == "" || common.DiscoveryPrefix == "" {
		return false
	}
	return true
}
func (component HADeviceSelect) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/select/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceSelect) CanGenerateTopic() bool {
	if (component.Device.Name == nil && component.Name == nil) || common.DiscoveryPrefix == "" || common.NodeID == "" {
		return false
	}
	return true
}
func (component HADeviceSensor) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/sensor/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceSensor) CanGenerateTopic() bool {
	if (component.Device.Name == nil && component.Name == nil) || common.DiscoveryPrefix == "" || common.NodeID == "" {
		return false
	}
	return true
}
func (component HADeviceSwitch) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/switch/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceSwitch) CanGenerateTopic() bool {
	if (component.Device.Name == nil && component.Name == nil) || common.DiscoveryPrefix == "" || common.NodeID == "" {
		return false
	}
	return true
}
func (component HADeviceTag) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/tag/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Device.Name) + "/"
	return
}
func (component HADeviceTag) CanGenerateTopic() bool {
	if component.Device.Name == nil || common.NodeID == "" || common.DiscoveryPrefix == "" {
		return false
	}
	return true
}
func (component HADeviceVacuum) GetTopicBase() (out string) {
	out = common.DiscoveryPrefix + "/vacuum/" + common.NodeID + "/"
	out += strcase.ToSnake(*component.Name) + "/"
	return
}
func (component HADeviceVacuum) CanGenerateTopic() bool {
	if component.Name == nil || common.NodeID == "" || common.DiscoveryPrefix == "" {
		return false
	}
	return true
}
