package hadiscovery

func (d Light) GetRawId() string {
	return "light"
}
func (d Light) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Light) GetUniqueId() string {
	return d.UniqueId
}
func (d Light) UnSubscribe()       {}
func (d Light) AnnounceAvailable() {}
