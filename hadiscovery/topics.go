package hadiscovery

func (d AlarmControlPanel) GetTopicPrefix() string {
	return NodeID + "/alarm_control_panel/" + d.UniqueID() + "/"
}
func (d AlarmControlPanel) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/alarm_control_panel/" + NodeID + "/" + d.UniqueID() + "/"
}
func (d BinarySensor) GetTopicPrefix() string {
	return NodeID + "/binary_sensor/" + d.UniqueID() + "/"
}
func (d BinarySensor) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/binary_sensor/" + NodeID + "/" + d.UniqueID() + "/"
}
func (d Camera) GetTopicPrefix() string {
	return NodeID + "/camera/" + d.UniqueID() + "/"
}
func (d Camera) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/camera/" + NodeID + "/" + d.UniqueID() + "/"
}
func (d Cover) GetTopicPrefix() string {
	return NodeID + "/cover/" + d.UniqueID() + "/"
}
func (d Cover) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/cover/" + NodeID + "/" + d.UniqueID() + "/"
}
func (d DeviceTracker) GetTopicPrefix() string {
	return NodeID + "/device_tracker/" + d.UniqueID() + "/"
}
func (d DeviceTracker) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/device_tracker/" + NodeID + "/" + d.UniqueID() + "/"
}
func (d DeviceTrigger) GetTopicPrefix() string {
	return NodeID + "/device_trigger/" + d.UniqueID() + "/"
}
func (d DeviceTrigger) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/device_trigger/" + NodeID + "/" + d.UniqueID() + "/"
}
func (d Fan) GetTopicPrefix() string {
	return NodeID + "/fan/" + d.UniqueID() + "/"
}
func (d Fan) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/fan/" + NodeID + "/" + d.UniqueID() + "/"
}
func (d Humidifier) GetTopicPrefix() string {
	return NodeID + "/humidifier/" + d.UniqueID() + "/"
}
func (d Humidifier) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/humidifier/" + NodeID + "/" + d.UniqueID() + "/"
}
func (d Climate) GetTopicPrefix() string {
	return NodeID + "/climate/" + d.UniqueID() + "/"
}
func (d Climate) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/climate/" + NodeID + "/" + d.UniqueID() + "/"
}
func (d Light) GetTopicPrefix() string {
	return NodeID + "/light/" + d.UniqueID() + "/"
}
func (d Light) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/light/" + NodeID + "/" + d.UniqueID() + "/"
}
func (d Lock) GetTopicPrefix() string {
	return NodeID + "/lock/" + d.UniqueID() + "/"
}
func (d Lock) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/lock/" + NodeID + "/" + d.UniqueID() + "/"
}
func (d Number) GetTopicPrefix() string {
	return NodeID + "/number/" + d.UniqueID() + "/"
}
func (d Number) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/number/" + NodeID + "/" + d.UniqueID() + "/"
}
func (d Scene) GetTopicPrefix() string {
	return NodeID + "/scene/" + d.UniqueID() + "/"
}
func (d Scene) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/scene/" + NodeID + "/" + d.UniqueID() + "/"
}
func (d Select) GetTopicPrefix() string {
	return NodeID + "/select/" + d.UniqueID() + "/"
}
func (d Select) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/select/" + NodeID + "/" + d.UniqueID() + "/"
}
func (d Sensor) GetTopicPrefix() string {
	return NodeID + "/sensor/" + d.UniqueID() + "/"
}
func (d Sensor) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/sensor/" + NodeID + "/" + d.UniqueID() + "/"
}
func (d Switch) GetTopicPrefix() string {
	return NodeID + "/switch/" + d.UniqueID() + "/"
}
func (d Switch) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/switch/" + NodeID + "/" + d.UniqueID() + "/"
}
func (d Tag) GetTopicPrefix() string {
	return NodeID + "/tag/" + d.UniqueID() + "/"
}
func (d Tag) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/tag/" + NodeID + "/" + d.UniqueID() + "/"
}
func (d Vacuum) GetTopicPrefix() string {
	return NodeID + "/vacuum/" + d.UniqueID() + "/"
}
func (d Vacuum) GetDiscoveryTopic() string {
	return DiscoveryPrefix + "/vacuum/" + NodeID + "/" + d.UniqueID() + "/"
}
