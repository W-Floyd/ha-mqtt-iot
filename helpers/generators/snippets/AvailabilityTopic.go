err, canGen := output.CanGenerateTopic()

if unchanged && !canGen {
	logging.LogError("Unable to generate outputTHIS, missing " + err.Error())
}

if output.Availability.Topic == nil && output.AvailabilityTopic == nil && canGen {
	n += 1
	output.AvailabilityTopic = common.StringPointer(TopicAvailability(output.GetTopicBase()))
}