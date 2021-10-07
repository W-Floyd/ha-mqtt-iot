err, canGen := config.Configuration.CanGenerateTopic()

if unchanged && !canGen {
	logging.LogError("Unable to generate config.ConfigurationTHIS, missing " + err.Error())
}

if config.Configuration.CommandTopic == nil && canGen {
	n += 1
	config.ConfigurationTHIS = common.StringPointer(TopicCommand(config.Configuration.GetTopicBase()))
}