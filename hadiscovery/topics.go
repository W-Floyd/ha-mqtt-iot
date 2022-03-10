package hadiscovery

func (d DeviceAlarmControlPanel) GetTopicPrefix() string {
	return NodeID + "/alarm_control_panel/" + d.UniqueID + "/"
}
func (d DeviceAlarmControlPanel) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/alarm_control_panel/" + NodeID + "/" + d.UniqueID + "/"
}
func (d DeviceBinarySensor) GetTopicPrefix() string {
	return NodeID + "/binary_sensor/" + d.UniqueID + "/"
}
func (d DeviceBinarySensor) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/binary_sensor/" + NodeID + "/" + d.UniqueID + "/"
}
func (d DeviceCamera) GetTopicPrefix() string {
	return NodeID + "/camera/" + d.UniqueID + "/"
}
func (d DeviceCamera) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/camera/" + NodeID + "/" + d.UniqueID + "/"
}
func (d DeviceCover) GetTopicPrefix() string {
	return NodeID + "/cover/" + d.UniqueID + "/"
}
func (d DeviceCover) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/cover/" + NodeID + "/" + d.UniqueID + "/"
}
func (d DeviceDeviceTracker) GetTopicPrefix() string {
	return NodeID + "/device_tracker/" + d.UniqueID + "/"
}
func (d DeviceDeviceTracker) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/device_tracker/" + NodeID + "/" + d.UniqueID + "/"
}
func (d DeviceDeviceTrigger) GetTopicPrefix() string {
	return NodeID + "/device_trigger/" + d.UniqueID + "/"
}
func (d DeviceDeviceTrigger) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/device_trigger/" + NodeID + "/" + d.UniqueID + "/"
}
func (d DeviceFan) GetTopicPrefix() string {
	return NodeID + "/fan/" + d.UniqueID + "/"
}
func (d DeviceFan) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/fan/" + NodeID + "/" + d.UniqueID + "/"
}
func (d DeviceHumidifier) GetTopicPrefix() string {
	return NodeID + "/humidifier/" + d.UniqueID + "/"
}
func (d DeviceHumidifier) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/humidifier/" + NodeID + "/" + d.UniqueID + "/"
}
func (d DeviceClimate) GetTopicPrefix() string {
	return NodeID + "/climate/" + d.UniqueID + "/"
}
func (d DeviceClimate) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/climate/" + NodeID + "/" + d.UniqueID + "/"
}
func (d DeviceLight) GetTopicPrefix() string {
	return NodeID + "/light/" + d.UniqueID + "/"
}
func (d DeviceLight) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/light/" + NodeID + "/" + d.UniqueID + "/"
}
func (d DeviceLock) GetTopicPrefix() string {
	return NodeID + "/lock/" + d.UniqueID + "/"
}
func (d DeviceLock) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/lock/" + NodeID + "/" + d.UniqueID + "/"
}
func (d DeviceNumber) GetTopicPrefix() string {
	return NodeID + "/number/" + d.UniqueID + "/"
}
func (d DeviceNumber) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/number/" + NodeID + "/" + d.UniqueID + "/"
}
func (d DeviceScene) GetTopicPrefix() string {
	return NodeID + "/scene/" + d.UniqueID + "/"
}
func (d DeviceScene) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/scene/" + NodeID + "/" + d.UniqueID + "/"
}
func (d DeviceSelect) GetTopicPrefix() string {
	return NodeID + "/select/" + d.UniqueID + "/"
}
func (d DeviceSelect) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/select/" + NodeID + "/" + d.UniqueID + "/"
}
func (d DeviceSensor) GetTopicPrefix() string {
	return NodeID + "/sensor/" + d.UniqueID + "/"
}
func (d DeviceSensor) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/sensor/" + NodeID + "/" + d.UniqueID + "/"
}
func (d DeviceSwitch) GetTopicPrefix() string {
	return NodeID + "/switch/" + d.UniqueID + "/"
}
func (d DeviceSwitch) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/switch/" + NodeID + "/" + d.UniqueID + "/"
}
func (d DeviceTag) GetTopicPrefix() string {
	return NodeID + "/tag/" + d.UniqueID + "/"
}
func (d DeviceTag) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/tag/" + NodeID + "/" + d.UniqueID + "/"
}
func (d DeviceVacuum) GetTopicPrefix() string {
	return NodeID + "/vacuum/" + d.UniqueID + "/"
}
func (d DeviceVacuum) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/vacuum/" + NodeID + "/" + d.UniqueID + "/"
}
