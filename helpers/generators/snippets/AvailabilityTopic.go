if unchanged && output.CanGenerateTopic() == false {
	logging.LogError("Unable to generate availability topics.")
}

if output.Availability.Topic == nil && output.AvailabilityTopic == nil && output.CanGenerateTopic() {
	n += 1
	outputTHIS = common.StringPointer(TopicAvailability(output.GetTopicBase()))
}