err, canGen := config.Configuration.CanGenerateTopic()

if unchanged && !canGen {
	logging.LogError("Unable to generate config.ConfigurationTHIS, missing " + err.Error())
}

if config.Configuration.Availability.Topic == nil && config.Configuration.AvailabilityTopic == nil && canGen {
	n += 1
	config.Configuration.AvailabilityTopic = common.StringPointer(TopicAvailability(config.Configuration.GetTopicBase()))
}