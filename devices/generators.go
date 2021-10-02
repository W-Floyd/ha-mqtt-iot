package devices

import (
	"github.com/W-Floyd/ha-mqtt-iot/common"
	"github.com/W-Floyd/ha-mqtt-iot/logging"
)

func (component *HADeviceAlarmControlPanel) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if component.Availability.Topic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.Availability.Topic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.AvailabilityTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.AvailabilityTopic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.CommandTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.CommandTopic, missing " + err.Error())
			}

			if component.CommandTopic == nil && canGen {
				n += 1
				component.CommandTopic = common.StringPointer(TopicCommand(component.GetTopicBase()))
			}
		}
		if component.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate component.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				component.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if component.StateTopic == nil {
			if unchanged {
				logging.LogError("component.StateTopic generator not found, but field is required!")
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
func (component *HADeviceBinarySensor) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if component.Availability.Topic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.Availability.Topic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.AvailabilityTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.AvailabilityTopic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate component.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				component.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if component.StateTopic == nil {
			if unchanged {
				logging.LogError("component.StateTopic generator not found, but field is required!")
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
func (component *HADeviceCamera) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if component.Availability.Topic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.Availability.Topic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.AvailabilityTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.AvailabilityTopic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate component.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				component.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if component.Topic == nil {
			if unchanged {
				logging.LogError("component.Topic generator not found, but field is required!")
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
func (component *HADeviceCover) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if component.Availability.Topic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.Availability.Topic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.AvailabilityTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.AvailabilityTopic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.CommandTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.CommandTopic, missing " + err.Error())
			}

			if component.CommandTopic == nil && canGen {
				n += 1
				component.CommandTopic = common.StringPointer(TopicCommand(component.GetTopicBase()))
			}
		}
		if component.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate component.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				component.Device.Name = common.StringPointer(common.InstanceName)
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
func (component *HADeviceDeviceTracker) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if component.Devices == nil {
			if unchanged {
				logging.LogError("component.Devices generator not found, but field is required!")
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
func (component *HADeviceDeviceTrigger) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if component.AutomationType == nil {
			if unchanged {
				logging.LogError("component.AutomationType generator not found, but field is required!")
			}
		}
		if component.Device.Connections == nil {
			if unchanged {
				logging.LogError("component.Device.Connections generator not found, but field is required!")
			}
		}
		if component.Device.Identifiers == nil {
			if unchanged {
				logging.LogError("component.Device.Identifiers generator not found, but field is required!")
			}
		}
		if component.Device.Manufacturer == nil {
			if unchanged {
				logging.LogError("component.Device.Manufacturer generator not found, but field is required!")
			}
		}
		if component.Device.Model == nil {
			if unchanged {
				logging.LogError("component.Device.Model generator not found, but field is required!")
			}
		}
		if component.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate component.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				component.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if component.Device.SuggestedArea == nil {
			if unchanged {
				logging.LogError("component.Device.SuggestedArea generator not found, but field is required!")
			}
		}
		if component.Device.SwVersion == nil {
			if unchanged {
				logging.LogError("component.Device.SwVersion generator not found, but field is required!")
			}
		}
		if component.Device.ViaDevice == nil {
			if unchanged {
				logging.LogError("component.Device.ViaDevice generator not found, but field is required!")
			}
		}
		if component.Subtype == nil {
			if unchanged {
				logging.LogError("component.Subtype generator not found, but field is required!")
			}
		}
		if component.Topic == nil {
			if unchanged {
				logging.LogError("component.Topic generator not found, but field is required!")
			}
		}
		if component.Type == nil {
			if unchanged {
				logging.LogError("component.Type generator not found, but field is required!")
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
func (component *HADeviceFan) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if component.Availability.Topic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.Availability.Topic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.AvailabilityTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.AvailabilityTopic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.CommandTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.CommandTopic, missing " + err.Error())
			}

			if component.CommandTopic == nil && canGen {
				n += 1
				component.CommandTopic = common.StringPointer(TopicCommand(component.GetTopicBase()))
			}
		}
		if component.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate component.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				component.Device.Name = common.StringPointer(common.InstanceName)
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
func (component *HADeviceHumidifier) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if component.Availability.Topic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.Availability.Topic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.AvailabilityTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.AvailabilityTopic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.CommandTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.CommandTopic, missing " + err.Error())
			}

			if component.CommandTopic == nil && canGen {
				n += 1
				component.CommandTopic = common.StringPointer(TopicCommand(component.GetTopicBase()))
			}
		}
		if component.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate component.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				component.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if component.TargetHumidityCommandTopic == nil {
			if unchanged {
				logging.LogError("component.TargetHumidityCommandTopic generator not found, but field is required!")
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
func (component *HADeviceClimate) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if component.Availability.Topic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.Availability.Topic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.AvailabilityTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.AvailabilityTopic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate component.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				component.Device.Name = common.StringPointer(common.InstanceName)
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
func (component *HADeviceLight) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if component.Availability.Topic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.Availability.Topic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.AvailabilityTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.AvailabilityTopic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.CommandTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.CommandTopic, missing " + err.Error())
			}

			if component.CommandTopic == nil && canGen {
				n += 1
				component.CommandTopic = common.StringPointer(TopicCommand(component.GetTopicBase()))
			}
		}
		if component.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate component.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				component.Device.Name = common.StringPointer(common.InstanceName)
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
func (component *HADeviceLock) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if component.Availability.Topic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.Availability.Topic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.AvailabilityTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.AvailabilityTopic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.CommandTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.CommandTopic, missing " + err.Error())
			}

			if component.CommandTopic == nil && canGen {
				n += 1
				component.CommandTopic = common.StringPointer(TopicCommand(component.GetTopicBase()))
			}
		}
		if component.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate component.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				component.Device.Name = common.StringPointer(common.InstanceName)
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
func (component *HADeviceNumber) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if component.Availability.Topic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.Availability.Topic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.AvailabilityTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.AvailabilityTopic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.CommandTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.CommandTopic, missing " + err.Error())
			}

			if component.CommandTopic == nil && canGen {
				n += 1
				component.CommandTopic = common.StringPointer(TopicCommand(component.GetTopicBase()))
			}
		}
		if component.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate component.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				component.Device.Name = common.StringPointer(common.InstanceName)
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
func (component *HADeviceScene) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if component.Availability.Topic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.Availability.Topic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.AvailabilityTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.AvailabilityTopic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.CommandTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.CommandTopic, missing " + err.Error())
			}

			if component.CommandTopic == nil && canGen {
				n += 1
				component.CommandTopic = common.StringPointer(TopicCommand(component.GetTopicBase()))
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
func (component *HADeviceSelect) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if component.Availability.Topic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.Availability.Topic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.AvailabilityTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.AvailabilityTopic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.CommandTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.CommandTopic, missing " + err.Error())
			}

			if component.CommandTopic == nil && canGen {
				n += 1
				component.CommandTopic = common.StringPointer(TopicCommand(component.GetTopicBase()))
			}
		}
		if component.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate component.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				component.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if component.Options == nil {
			if unchanged {
				logging.LogError("component.Options generator not found, but field is required!")
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
func (component *HADeviceSensor) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if component.Availability.Topic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.Availability.Topic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.AvailabilityTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.AvailabilityTopic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate component.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				component.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if component.StateTopic == nil {
			if unchanged {
				logging.LogError("component.StateTopic generator not found, but field is required!")
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
func (component *HADeviceSwitch) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if component.Availability.Topic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.Availability.Topic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.AvailabilityTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.AvailabilityTopic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.CommandTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.CommandTopic, missing " + err.Error())
			}

			if component.CommandTopic == nil && canGen {
				n += 1
				component.CommandTopic = common.StringPointer(TopicCommand(component.GetTopicBase()))
			}
		}
		if component.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate component.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				component.Device.Name = common.StringPointer(common.InstanceName)
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
func (component *HADeviceTag) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if component.Device.Connections == nil {
			if unchanged {
				logging.LogError("component.Device.Connections generator not found, but field is required!")
			}
		}
		if component.Device.Identifiers == nil {
			if unchanged {
				logging.LogError("component.Device.Identifiers generator not found, but field is required!")
			}
		}
		if component.Device.Manufacturer == nil {
			if unchanged {
				logging.LogError("component.Device.Manufacturer generator not found, but field is required!")
			}
		}
		if component.Device.Model == nil {
			if unchanged {
				logging.LogError("component.Device.Model generator not found, but field is required!")
			}
		}
		if component.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate component.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				component.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if component.Device.SuggestedArea == nil {
			if unchanged {
				logging.LogError("component.Device.SuggestedArea generator not found, but field is required!")
			}
		}
		if component.Device.SwVersion == nil {
			if unchanged {
				logging.LogError("component.Device.SwVersion generator not found, but field is required!")
			}
		}
		if component.Device.ViaDevice == nil {
			if unchanged {
				logging.LogError("component.Device.ViaDevice generator not found, but field is required!")
			}
		}
		if component.Topic == nil {
			if unchanged {
				logging.LogError("component.Topic generator not found, but field is required!")
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
func (component *HADeviceVacuum) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if component.Availability.Topic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.Availability.Topic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.AvailabilityTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.AvailabilityTopic, missing " + err.Error())
			}

			if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
				n += 1
				component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
			}
		}
		if component.CommandTopic == nil {
			err, canGen := component.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate component.CommandTopic, missing " + err.Error())
			}

			if component.CommandTopic == nil && canGen {
				n += 1
				component.CommandTopic = common.StringPointer(TopicCommand(component.GetTopicBase()))
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
