if common.InstanceName == "" {
	if unchanged {
		logging.LogError("Unable to generate componentTHIS, missing common.InstanceName")
	}
} else {
	n += 1
	component.Device.Name = common.StringPointer(common.InstanceName)
}