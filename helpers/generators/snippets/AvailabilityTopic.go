err, canGen := component.CanGenerateTopic()

if unchanged && !canGen {
	logging.LogError("Unable to generate componentTHIS, missing " + err.Error())
}

if component.Availability.Topic == nil && component.AvailabilityTopic == nil && canGen {
	n += 1
	component.AvailabilityTopic = common.StringPointer(TopicAvailability(component.GetTopicBase()))
}