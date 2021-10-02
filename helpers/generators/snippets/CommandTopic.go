err, canGen := component.CanGenerateTopic()

if unchanged && !canGen {
	logging.LogError("Unable to generate componentTHIS, missing " + err.Error())
}

if component.CommandTopic == nil && canGen {
	n += 1
	componentTHIS = common.StringPointer(TopicCommand(component.GetTopicBase()))
}