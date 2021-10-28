package devices

import (
	"github.com/W-Floyd/ha-mqtt-iot/common"
	"github.com/W-Floyd/ha-mqtt-iot/logging"
)

func (config *HAConfigAlarmControlPanel) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if config.Configuration.Availability.Topic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.Availability.Topic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.AvailabilityTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.AvailabilityTopic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.CommandTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.CommandTopic, missing " + err.Error())
			}

			if config.Configuration.CommandTopic == nil && canGen {
				n += 1
				config.Configuration.CommandTopic = common.StringPointer(TopicCommand(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate config.Configuration.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				config.Configuration.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if config.Configuration.StateTopic == nil {
			if unchanged {
				logging.LogError("config.Configuration.StateTopic generator not found, but field is required!")
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
func (config *HAConfigBinarySensor) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if config.Configuration.Availability.Topic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.Availability.Topic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.AvailabilityTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.AvailabilityTopic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate config.Configuration.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				config.Configuration.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if config.Configuration.StateTopic == nil {
			if unchanged {
				logging.LogError("config.Configuration.StateTopic generator not found, but field is required!")
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
func (config *HAConfigCamera) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if config.Configuration.Availability.Topic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.Availability.Topic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.AvailabilityTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.AvailabilityTopic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate config.Configuration.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				config.Configuration.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if config.Configuration.Topic == nil {
			if unchanged {
				logging.LogError("config.Configuration.Topic generator not found, but field is required!")
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
func (config *HAConfigCover) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if config.Configuration.Availability.Topic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.Availability.Topic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.AvailabilityTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.AvailabilityTopic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.CommandTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.CommandTopic, missing " + err.Error())
			}

			if config.Configuration.CommandTopic == nil && canGen {
				n += 1
				config.Configuration.CommandTopic = common.StringPointer(TopicCommand(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate config.Configuration.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				config.Configuration.Device.Name = common.StringPointer(common.InstanceName)
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
func (config *HAConfigDeviceTracker) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if config.Configuration.Devices == nil {
			if unchanged {
				logging.LogError("config.Configuration.Devices generator not found, but field is required!")
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
func (config *HAConfigDeviceTrigger) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if config.Configuration.AutomationType == nil {
			if unchanged {
				logging.LogError("config.Configuration.AutomationType generator not found, but field is required!")
			}
		}
		if config.Configuration.Device.ConfigurationUrl == nil {
			if unchanged {
				logging.LogError("config.Configuration.Device.ConfigurationUrl generator not found, but field is required!")
			}
		}
		if config.Configuration.Device.Connections == nil {
			if unchanged {
				logging.LogError("config.Configuration.Device.Connections generator not found, but field is required!")
			}
		}
		if config.Configuration.Device.Identifiers == nil {
			if unchanged {
				logging.LogError("config.Configuration.Device.Identifiers generator not found, but field is required!")
			}
		}
		if config.Configuration.Device.Manufacturer == nil {
			if unchanged {
				logging.LogError("config.Configuration.Device.Manufacturer generator not found, but field is required!")
			}
		}
		if config.Configuration.Device.Model == nil {
			if unchanged {
				logging.LogError("config.Configuration.Device.Model generator not found, but field is required!")
			}
		}
		if config.Configuration.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate config.Configuration.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				config.Configuration.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if config.Configuration.Device.SuggestedArea == nil {
			if unchanged {
				logging.LogError("config.Configuration.Device.SuggestedArea generator not found, but field is required!")
			}
		}
		if config.Configuration.Device.SwVersion == nil {
			if unchanged {
				logging.LogError("config.Configuration.Device.SwVersion generator not found, but field is required!")
			}
		}
		if config.Configuration.Device.ViaDevice == nil {
			if unchanged {
				logging.LogError("config.Configuration.Device.ViaDevice generator not found, but field is required!")
			}
		}
		if config.Configuration.Subtype == nil {
			if unchanged {
				logging.LogError("config.Configuration.Subtype generator not found, but field is required!")
			}
		}
		if config.Configuration.Topic == nil {
			if unchanged {
				logging.LogError("config.Configuration.Topic generator not found, but field is required!")
			}
		}
		if config.Configuration.Type == nil {
			if unchanged {
				logging.LogError("config.Configuration.Type generator not found, but field is required!")
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
func (config *HAConfigFan) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if config.Configuration.Availability.Topic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.Availability.Topic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.AvailabilityTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.AvailabilityTopic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.CommandTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.CommandTopic, missing " + err.Error())
			}

			if config.Configuration.CommandTopic == nil && canGen {
				n += 1
				config.Configuration.CommandTopic = common.StringPointer(TopicCommand(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate config.Configuration.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				config.Configuration.Device.Name = common.StringPointer(common.InstanceName)
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
func (config *HAConfigHumidifier) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if config.Configuration.Availability.Topic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.Availability.Topic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.AvailabilityTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.AvailabilityTopic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.CommandTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.CommandTopic, missing " + err.Error())
			}

			if config.Configuration.CommandTopic == nil && canGen {
				n += 1
				config.Configuration.CommandTopic = common.StringPointer(TopicCommand(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate config.Configuration.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				config.Configuration.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if config.Configuration.TargetHumidityCommandTopic == nil {
			if unchanged {
				logging.LogError("config.Configuration.TargetHumidityCommandTopic generator not found, but field is required!")
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
func (config *HAConfigClimate) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if config.Configuration.Availability.Topic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.Availability.Topic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.AvailabilityTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.AvailabilityTopic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate config.Configuration.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				config.Configuration.Device.Name = common.StringPointer(common.InstanceName)
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
func (config *HAConfigLight) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if config.Configuration.Availability.Topic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.Availability.Topic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.AvailabilityTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.AvailabilityTopic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.CommandTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.CommandTopic, missing " + err.Error())
			}

			if config.Configuration.CommandTopic == nil && canGen {
				n += 1
				config.Configuration.CommandTopic = common.StringPointer(TopicCommand(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate config.Configuration.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				config.Configuration.Device.Name = common.StringPointer(common.InstanceName)
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
func (config *HAConfigLock) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if config.Configuration.Availability.Topic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.Availability.Topic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.AvailabilityTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.AvailabilityTopic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.CommandTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.CommandTopic, missing " + err.Error())
			}

			if config.Configuration.CommandTopic == nil && canGen {
				n += 1
				config.Configuration.CommandTopic = common.StringPointer(TopicCommand(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate config.Configuration.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				config.Configuration.Device.Name = common.StringPointer(common.InstanceName)
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
func (config *HAConfigNumber) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if config.Configuration.Availability.Topic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.Availability.Topic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.AvailabilityTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.AvailabilityTopic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.CommandTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.CommandTopic, missing " + err.Error())
			}

			if config.Configuration.CommandTopic == nil && canGen {
				n += 1
				config.Configuration.CommandTopic = common.StringPointer(TopicCommand(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate config.Configuration.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				config.Configuration.Device.Name = common.StringPointer(common.InstanceName)
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
func (config *HAConfigScene) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if config.Configuration.Availability.Topic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.Availability.Topic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.AvailabilityTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.AvailabilityTopic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.CommandTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.CommandTopic, missing " + err.Error())
			}

			if config.Configuration.CommandTopic == nil && canGen {
				n += 1
				config.Configuration.CommandTopic = common.StringPointer(TopicCommand(config.Configuration.GetTopicBase()))
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
func (config *HAConfigSelect) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if config.Configuration.Availability.Topic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.Availability.Topic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.AvailabilityTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.AvailabilityTopic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.CommandTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.CommandTopic, missing " + err.Error())
			}

			if config.Configuration.CommandTopic == nil && canGen {
				n += 1
				config.Configuration.CommandTopic = common.StringPointer(TopicCommand(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate config.Configuration.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				config.Configuration.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if config.Configuration.Options == nil {
			if unchanged {
				logging.LogError("config.Configuration.Options generator not found, but field is required!")
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
func (config *HAConfigSensor) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if config.Configuration.Availability.Topic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.Availability.Topic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.AvailabilityTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.AvailabilityTopic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate config.Configuration.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				config.Configuration.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if config.Configuration.StateTopic == nil {
			if unchanged {
				logging.LogError("config.Configuration.StateTopic generator not found, but field is required!")
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
func (config *HAConfigSwitch) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if config.Configuration.Availability.Topic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.Availability.Topic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.AvailabilityTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.AvailabilityTopic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.CommandTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.CommandTopic, missing " + err.Error())
			}

			if config.Configuration.CommandTopic == nil && canGen {
				n += 1
				config.Configuration.CommandTopic = common.StringPointer(TopicCommand(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate config.Configuration.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				config.Configuration.Device.Name = common.StringPointer(common.InstanceName)
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
func (config *HAConfigTag) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if config.Configuration.Device.ConfigurationUrl == nil {
			if unchanged {
				logging.LogError("config.Configuration.Device.ConfigurationUrl generator not found, but field is required!")
			}
		}
		if config.Configuration.Device.Connections == nil {
			if unchanged {
				logging.LogError("config.Configuration.Device.Connections generator not found, but field is required!")
			}
		}
		if config.Configuration.Device.Identifiers == nil {
			if unchanged {
				logging.LogError("config.Configuration.Device.Identifiers generator not found, but field is required!")
			}
		}
		if config.Configuration.Device.Manufacturer == nil {
			if unchanged {
				logging.LogError("config.Configuration.Device.Manufacturer generator not found, but field is required!")
			}
		}
		if config.Configuration.Device.Model == nil {
			if unchanged {
				logging.LogError("config.Configuration.Device.Model generator not found, but field is required!")
			}
		}
		if config.Configuration.Device.Name == nil {
			if common.InstanceName == "" {
				if unchanged {
					logging.LogError("Unable to generate config.Configuration.Device.Name, missing common.InstanceName")
				}
			} else {
				n += 1
				config.Configuration.Device.Name = common.StringPointer(common.InstanceName)
			}
		}
		if config.Configuration.Device.SuggestedArea == nil {
			if unchanged {
				logging.LogError("config.Configuration.Device.SuggestedArea generator not found, but field is required!")
			}
		}
		if config.Configuration.Device.SwVersion == nil {
			if unchanged {
				logging.LogError("config.Configuration.Device.SwVersion generator not found, but field is required!")
			}
		}
		if config.Configuration.Device.ViaDevice == nil {
			if unchanged {
				logging.LogError("config.Configuration.Device.ViaDevice generator not found, but field is required!")
			}
		}
		if config.Configuration.Topic == nil {
			if unchanged {
				logging.LogError("config.Configuration.Topic generator not found, but field is required!")
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
func (config *HAConfigVacuum) Generate() {
	n := 0
	oldN := 0
	unchanged := false
	for {
		if config.Configuration.Availability.Topic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.Availability.Topic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.AvailabilityTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.AvailabilityTopic, missing " + err.Error())
			}

			if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
				n += 1
				config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
			}
		}
		if config.Configuration.CommandTopic == nil {
			err, canGen := config.Configuration.CanGenerateTopic()

			if unchanged && !canGen {
				logging.LogError("Unable to generate config.Configuration.CommandTopic, missing " + err.Error())
			}

			if config.Configuration.CommandTopic == nil && canGen {
				n += 1
				config.Configuration.CommandTopic = common.StringPointer(TopicCommand(config.Configuration.GetTopicBase()))
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
