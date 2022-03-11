package hadiscovery

import mqtt "github.com/eclipse/paho.mqtt.golang"

func (d Light) GetRawId() string {
	return "light"
}
func (d Light) GetMQTTClient() mqtt.Client {
	return *d.MQTT.Client
}
func (d Light) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Light) GetUniqueId() string {
	return d.UniqueId
}
func (d Light) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d Light) UpdateState()       {}
func (d Light) Subscribe()         {}
func (d Light) UnSubscribe()       {}
func (d Light) AnnounceAvailable() {}
