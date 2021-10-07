if common.InstanceName == "" {
	if unchanged {
		logging.LogError("Unable to generate config.ConfigurationTHIS, missing common.InstanceName")
	}
} else {
	n += 1
	config.Configuration.Device.Name = common.StringPointer(common.InstanceName)
}