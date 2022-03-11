package hadiscovery

import mqtt "github.com/eclipse/paho.mqtt.golang"

func (d BinarySensor) GetRawId() string {
	return "binary_sensor"
}
func (d BinarySensor) GetUniqueId() string {
	return d.UniqueId
}
func (d *BinarySensor) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d *BinarySensor) UpdateState(*mqtt.Client)     {}
func (d *BinarySensor) Subscribe(client mqtt.Client) {}
func (d Light) GetRawId() string {
	return "light"
}
func (d Light) GetUniqueId() string {
	return d.UniqueId
}
func (d *Light) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d *Light) UpdateState(*mqtt.Client)     {}
func (d *Light) Subscribe(client mqtt.Client) {}
func (d Sensor) GetRawId() string {
	return "sensor"
}
func (d Sensor) GetUniqueId() string {
	return d.UniqueId
}
func (d *Sensor) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d *Sensor) UpdateState(*mqtt.Client)     {}
func (d *Sensor) Subscribe(client mqtt.Client) {}
func (d Switch) GetRawId() string {
	return "switch"
}
func (d Switch) GetUniqueId() string {
	return d.UniqueId
}
func (d *Switch) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d *Switch) UpdateState(*mqtt.Client)     {}
func (d *Switch) Subscribe(client mqtt.Client) {}
