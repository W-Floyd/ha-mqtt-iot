if common.InstanceName == "" {
	if unchanged {
		logging.LogError("Unable to generate outputTHIS, missing common.InstanceName")
	}
} else {
	n += 1
	output.Device.Name = common.StringPointer(common.InstanceName)
}