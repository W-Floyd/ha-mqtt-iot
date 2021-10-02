if unchanged && output.CanGenerateTopic() == false {
	logging.LogError("Unable to generate command topics.")
}

if output.CommandTopic == nil && output.CanGenerateTopic() {
	n += 1
	outputTHIS = common.StringPointer(TopicCommand(output.GetTopicBase()))
}