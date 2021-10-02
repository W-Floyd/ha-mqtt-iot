package devices

import (
	"github.com/W-Floyd/ha-mqtt-iot/common"
	"github.com/W-Floyd/ha-mqtt-iot/logging"
	"github.com/jinzhu/copier"
)

func (component HADeviceAlarmControlPanel) Generate() (output HADeviceAlarmControlPanel) {
	copier.CopyWithOption(&output, &component, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	n := 0
	oldN := 0
	unchanged := false
	for {
		if output.Availability.Topic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.Availability.Topic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.AvailabilityTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.AvailabilityTopic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.CommandTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.CommandTopic, missing " + err.Error())
			}

			if output.CommandTopic == nil && canGen {
				n += 1
				output.CommandTopic = common.StringPointer(TopicCommand(output.GetTopicBase()))
			}
		}
		if output.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate output.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				output.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if output.StateTopic == nil {
			if unchanged {
				logging.LogError("output.StateTopic generator not found, but field is required!")
			}
		}
		if unchanged {
			break
		}
		if n == oldN {
			unchanged = true
		}
		oldN = n
	}
	return
}
func (component HADeviceBinarySensor) Generate() (output HADeviceBinarySensor) {
	copier.CopyWithOption(&output, &component, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	n := 0
	oldN := 0
	unchanged := false
	for {
		if output.Availability.Topic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.Availability.Topic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.AvailabilityTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.AvailabilityTopic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate output.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				output.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if output.StateTopic == nil {
			if unchanged {
				logging.LogError("output.StateTopic generator not found, but field is required!")
			}
		}
		if unchanged {
			break
		}
		if n == oldN {
			unchanged = true
		}
		oldN = n
	}
	return
}
func (component HADeviceCamera) Generate() (output HADeviceCamera) {
	copier.CopyWithOption(&output, &component, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	n := 0
	oldN := 0
	unchanged := false
	for {
		if output.Availability.Topic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.Availability.Topic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.AvailabilityTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.AvailabilityTopic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate output.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				output.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if output.Topic == nil {
			if unchanged {
				logging.LogError("output.Topic generator not found, but field is required!")
			}
		}
		if unchanged {
			break
		}
		if n == oldN {
			unchanged = true
		}
		oldN = n
	}
	return
}
func (component HADeviceCover) Generate() (output HADeviceCover) {
	copier.CopyWithOption(&output, &component, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	n := 0
	oldN := 0
	unchanged := false
	for {
		if output.Availability.Topic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.Availability.Topic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.AvailabilityTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.AvailabilityTopic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.CommandTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.CommandTopic, missing " + err.Error())
			}

			if output.CommandTopic == nil && canGen {
				n += 1
				output.CommandTopic = common.StringPointer(TopicCommand(output.GetTopicBase()))
			}
		}
		if output.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate output.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				output.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if unchanged {
			break
		}
		if n == oldN {
			unchanged = true
		}
		oldN = n
	}
	return
}
func (component HADeviceDeviceTracker) Generate() (output HADeviceDeviceTracker) {
	copier.CopyWithOption(&output, &component, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	n := 0
	oldN := 0
	unchanged := false
	for {
		if output.Devices == nil {
			if unchanged {
				logging.LogError("output.Devices generator not found, but field is required!")
			}
		}
		if unchanged {
			break
		}
		if n == oldN {
			unchanged = true
		}
		oldN = n
	}
	return
}
func (component HADeviceDeviceTrigger) Generate() (output HADeviceDeviceTrigger) {
	copier.CopyWithOption(&output, &component, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	n := 0
	oldN := 0
	unchanged := false
	for {
		if output.AutomationType == nil {
			if unchanged {
				logging.LogError("output.AutomationType generator not found, but field is required!")
			}
		}
		if output.Device.Connections == nil {
			if unchanged {
				logging.LogError("output.Device.Connections generator not found, but field is required!")
			}
		}
		if output.Device.Identifiers == nil {
			if unchanged {
				logging.LogError("output.Device.Identifiers generator not found, but field is required!")
			}
		}
		if output.Device.Manufacturer == nil {
			if unchanged {
				logging.LogError("output.Device.Manufacturer generator not found, but field is required!")
			}
		}
		if output.Device.Model == nil {
			if unchanged {
				logging.LogError("output.Device.Model generator not found, but field is required!")
			}
		}
		if output.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate output.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				output.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if output.Device.SuggestedArea == nil {
			if unchanged {
				logging.LogError("output.Device.SuggestedArea generator not found, but field is required!")
			}
		}
		if output.Device.SwVersion == nil {
			if unchanged {
				logging.LogError("output.Device.SwVersion generator not found, but field is required!")
			}
		}
		if output.Device.ViaDevice == nil {
			if unchanged {
				logging.LogError("output.Device.ViaDevice generator not found, but field is required!")
			}
		}
		if output.Subtype == nil {
			if unchanged {
				logging.LogError("output.Subtype generator not found, but field is required!")
			}
		}
		if output.Topic == nil {
			if unchanged {
				logging.LogError("output.Topic generator not found, but field is required!")
			}
		}
		if output.Type == nil {
			if unchanged {
				logging.LogError("output.Type generator not found, but field is required!")
			}
		}
		if unchanged {
			break
		}
		if n == oldN {
			unchanged = true
		}
		oldN = n
	}
	return
}
func (component HADeviceFan) Generate() (output HADeviceFan) {
	copier.CopyWithOption(&output, &component, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	n := 0
	oldN := 0
	unchanged := false
	for {
		if output.Availability.Topic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.Availability.Topic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.AvailabilityTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.AvailabilityTopic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.CommandTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.CommandTopic, missing " + err.Error())
			}

			if output.CommandTopic == nil && canGen {
				n += 1
				output.CommandTopic = common.StringPointer(TopicCommand(output.GetTopicBase()))
			}
		}
		if output.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate output.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				output.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if unchanged {
			break
		}
		if n == oldN {
			unchanged = true
		}
		oldN = n
	}
	return
}
func (component HADeviceHumidifier) Generate() (output HADeviceHumidifier) {
	copier.CopyWithOption(&output, &component, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	n := 0
	oldN := 0
	unchanged := false
	for {
		if output.Availability.Topic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.Availability.Topic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.AvailabilityTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.AvailabilityTopic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.CommandTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.CommandTopic, missing " + err.Error())
			}

			if output.CommandTopic == nil && canGen {
				n += 1
				output.CommandTopic = common.StringPointer(TopicCommand(output.GetTopicBase()))
			}
		}
		if output.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate output.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				output.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if output.TargetHumidityCommandTopic == nil {
			if unchanged {
				logging.LogError("output.TargetHumidityCommandTopic generator not found, but field is required!")
			}
		}
		if unchanged {
			break
		}
		if n == oldN {
			unchanged = true
		}
		oldN = n
	}
	return
}
func (component HADeviceClimate) Generate() (output HADeviceClimate) {
	copier.CopyWithOption(&output, &component, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	n := 0
	oldN := 0
	unchanged := false
	for {
		if output.Availability.Topic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.Availability.Topic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.AvailabilityTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.AvailabilityTopic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate output.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				output.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if unchanged {
			break
		}
		if n == oldN {
			unchanged = true
		}
		oldN = n
	}
	return
}
func (component HADeviceLight) Generate() (output HADeviceLight) {
	copier.CopyWithOption(&output, &component, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	n := 0
	oldN := 0
	unchanged := false
	for {
		if output.Availability.Topic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.Availability.Topic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.AvailabilityTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.AvailabilityTopic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.CommandTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.CommandTopic, missing " + err.Error())
			}

			if output.CommandTopic == nil && canGen {
				n += 1
				output.CommandTopic = common.StringPointer(TopicCommand(output.GetTopicBase()))
			}
		}
		if output.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate output.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				output.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if unchanged {
			break
		}
		if n == oldN {
			unchanged = true
		}
		oldN = n
	}
	return
}
func (component HADeviceLock) Generate() (output HADeviceLock) {
	copier.CopyWithOption(&output, &component, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	n := 0
	oldN := 0
	unchanged := false
	for {
		if output.Availability.Topic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.Availability.Topic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.AvailabilityTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.AvailabilityTopic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.CommandTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.CommandTopic, missing " + err.Error())
			}

			if output.CommandTopic == nil && canGen {
				n += 1
				output.CommandTopic = common.StringPointer(TopicCommand(output.GetTopicBase()))
			}
		}
		if output.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate output.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				output.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if unchanged {
			break
		}
		if n == oldN {
			unchanged = true
		}
		oldN = n
	}
	return
}
func (component HADeviceNumber) Generate() (output HADeviceNumber) {
	copier.CopyWithOption(&output, &component, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	n := 0
	oldN := 0
	unchanged := false
	for {
		if output.Availability.Topic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.Availability.Topic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.AvailabilityTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.AvailabilityTopic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.CommandTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.CommandTopic, missing " + err.Error())
			}

			if output.CommandTopic == nil && canGen {
				n += 1
				output.CommandTopic = common.StringPointer(TopicCommand(output.GetTopicBase()))
			}
		}
		if output.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate output.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				output.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if unchanged {
			break
		}
		if n == oldN {
			unchanged = true
		}
		oldN = n
	}
	return
}
func (component HADeviceScene) Generate() (output HADeviceScene) {
	copier.CopyWithOption(&output, &component, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	n := 0
	oldN := 0
	unchanged := false
	for {
		if output.Availability.Topic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.Availability.Topic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.AvailabilityTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.AvailabilityTopic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.CommandTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.CommandTopic, missing " + err.Error())
			}

			if output.CommandTopic == nil && canGen {
				n += 1
				output.CommandTopic = common.StringPointer(TopicCommand(output.GetTopicBase()))
			}
		}
		if unchanged {
			break
		}
		if n == oldN {
			unchanged = true
		}
		oldN = n
	}
	return
}
func (component HADeviceSelect) Generate() (output HADeviceSelect) {
	copier.CopyWithOption(&output, &component, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	n := 0
	oldN := 0
	unchanged := false
	for {
		if output.Availability.Topic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.Availability.Topic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.AvailabilityTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.AvailabilityTopic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.CommandTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.CommandTopic, missing " + err.Error())
			}

			if output.CommandTopic == nil && canGen {
				n += 1
				output.CommandTopic = common.StringPointer(TopicCommand(output.GetTopicBase()))
			}
		}
		if output.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate output.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				output.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if output.Options == nil {
			if unchanged {
				logging.LogError("output.Options generator not found, but field is required!")
			}
		}
		if unchanged {
			break
		}
		if n == oldN {
			unchanged = true
		}
		oldN = n
	}
	return
}
func (component HADeviceSensor) Generate() (output HADeviceSensor) {
	copier.CopyWithOption(&output, &component, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	n := 0
	oldN := 0
	unchanged := false
	for {
		if output.Availability.Topic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.Availability.Topic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.AvailabilityTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.AvailabilityTopic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate output.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				output.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if output.StateTopic == nil {
			if unchanged {
				logging.LogError("output.StateTopic generator not found, but field is required!")
			}
		}
		if unchanged {
			break
		}
		if n == oldN {
			unchanged = true
		}
		oldN = n
	}
	return
}
func (component HADeviceSwitch) Generate() (output HADeviceSwitch) {
	copier.CopyWithOption(&output, &component, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	n := 0
	oldN := 0
	unchanged := false
	for {
		if output.Availability.Topic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.Availability.Topic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.AvailabilityTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.AvailabilityTopic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.CommandTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.CommandTopic, missing " + err.Error())
			}

			if output.CommandTopic == nil && canGen {
				n += 1
				output.CommandTopic = common.StringPointer(TopicCommand(output.GetTopicBase()))
			}
		}
		if output.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate output.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				output.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if unchanged {
			break
		}
		if n == oldN {
			unchanged = true
		}
		oldN = n
	}
	return
}
func (component HADeviceTag) Generate() (output HADeviceTag) {
	copier.CopyWithOption(&output, &component, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	n := 0
	oldN := 0
	unchanged := false
	for {
		if output.Device.Connections == nil {
			if unchanged {
				logging.LogError("output.Device.Connections generator not found, but field is required!")
			}
		}
		if output.Device.Identifiers == nil {
			if unchanged {
				logging.LogError("output.Device.Identifiers generator not found, but field is required!")
			}
		}
		if output.Device.Manufacturer == nil {
			if unchanged {
				logging.LogError("output.Device.Manufacturer generator not found, but field is required!")
			}
		}
		if output.Device.Model == nil {
			if unchanged {
				logging.LogError("output.Device.Model generator not found, but field is required!")
			}
		}
		if output.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate output.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				output.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if output.Device.SuggestedArea == nil {
			if unchanged {
				logging.LogError("output.Device.SuggestedArea generator not found, but field is required!")
			}
		}
		if output.Device.SwVersion == nil {
			if unchanged {
				logging.LogError("output.Device.SwVersion generator not found, but field is required!")
			}
		}
		if output.Device.ViaDevice == nil {
			if unchanged {
				logging.LogError("output.Device.ViaDevice generator not found, but field is required!")
			}
		}
		if output.Topic == nil {
			if unchanged {
				logging.LogError("output.Topic generator not found, but field is required!")
			}
		}
		if unchanged {
			break
		}
		if n == oldN {
			unchanged = true
		}
		oldN = n
	}
	return
}
func (component HADeviceVacuum) Generate() (output HADeviceVacuum) {
	copier.CopyWithOption(&output, &component, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	n := 0
	oldN := 0
	unchanged := false
	for {
		if output.Availability.Topic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.Availability.Topic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.AvailabilityTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.AvailabilityTopic, missing " + err.Error())
			}

			if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
				n += 1
				output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
			}
		}
		if output.CommandTopic == nil {
			err, canGen := output.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate output.CommandTopic, missing " + err.Error())
			}

			if output.CommandTopic == nil && canGen {
				n += 1
				output.CommandTopic = common.StringPointer(TopicCommand(output.GetTopicBase()))
			}
		}
		if unchanged {
			break
		}
		if n == oldN {
			unchanged = true
		}
		oldN = n
	}
	return
}
