err, canGen := output.CanGenerateTopic()

if unchanged && !canGen {
	logging.LogError("Unable to generate outputTHIS, missing " + err.Error())
}

if output.CommandTopic == nil && canGen {
	n += 1
	outputTHIS = common.StringPointer(TopicCommand(output.GetTopicBase()))
}