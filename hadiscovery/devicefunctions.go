package hadiscovery

import mqtt "github.com/eclipse/paho.mqtt.golang"

func (d AlarmControlPanel) GetRawId() string {
	return "alarm_control_panel"
}
func (d AlarmControlPanel) GetUniqueId() string {
	return d.UniqueId
}
func (d *AlarmControlPanel) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d *AlarmControlPanel) UpdateState()                 {}
func (d *AlarmControlPanel) Subscribe(client mqtt.Client) {}
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
func (d *BinarySensor) UpdateState()                 {}
func (d *BinarySensor) Subscribe(client mqtt.Client) {}
func (d Button) GetRawId() string {
	return "button"
}
func (d Button) GetUniqueId() string {
	return d.UniqueId
}
func (d *Button) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d *Button) UpdateState()                 {}
func (d *Button) Subscribe(client mqtt.Client) {}
func (d Camera) GetRawId() string {
	return "camera"
}
func (d Camera) GetUniqueId() string {
	return d.UniqueId
}
func (d *Camera) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d *Camera) UpdateState()                 {}
func (d *Camera) Subscribe(client mqtt.Client) {}
func (d Cover) GetRawId() string {
	return "cover"
}
func (d Cover) GetUniqueId() string {
	return d.UniqueId
}
func (d *Cover) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d *Cover) UpdateState()                 {}
func (d *Cover) Subscribe(client mqtt.Client) {}
func (d DeviceTracker) GetRawId() string {
	return "device_tracker"
}
func (d DeviceTracker) GetUniqueId() string {
	return ""
}
func (d *DeviceTracker) PopulateDevice()              {}
func (d *DeviceTracker) UpdateState()                 {}
func (d *DeviceTracker) Subscribe(client mqtt.Client) {}
func (d DeviceTrigger) GetRawId() string {
	return "device_trigger"
}
func (d DeviceTrigger) GetUniqueId() string {
	return ""
}
func (d *DeviceTrigger) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d *DeviceTrigger) UpdateState()                 {}
func (d *DeviceTrigger) Subscribe(client mqtt.Client) {}
func (d Fan) GetRawId() string {
	return "fan"
}
func (d Fan) GetUniqueId() string {
	return d.UniqueId
}
func (d *Fan) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d *Fan) UpdateState()                 {}
func (d *Fan) Subscribe(client mqtt.Client) {}
func (d Humidifier) GetRawId() string {
	return "humidifier"
}
func (d Humidifier) GetUniqueId() string {
	return d.UniqueId
}
func (d *Humidifier) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d *Humidifier) UpdateState()                 {}
func (d *Humidifier) Subscribe(client mqtt.Client) {}
func (d Climate) GetRawId() string {
	return "climate"
}
func (d Climate) GetUniqueId() string {
	return d.UniqueId
}
func (d *Climate) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d *Climate) UpdateState()                 {}
func (d *Climate) Subscribe(client mqtt.Client) {}
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
func (d *Light) UpdateState()                 {}
func (d *Light) Subscribe(client mqtt.Client) {}
func (d Lock) GetRawId() string {
	return "lock"
}
func (d Lock) GetUniqueId() string {
	return d.UniqueId
}
func (d *Lock) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d *Lock) UpdateState()                 {}
func (d *Lock) Subscribe(client mqtt.Client) {}
func (d Number) GetRawId() string {
	return "number"
}
func (d Number) GetUniqueId() string {
	return d.UniqueId
}
func (d *Number) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d *Number) UpdateState()                 {}
func (d *Number) Subscribe(client mqtt.Client) {}
func (d Scene) GetRawId() string {
	return "scene"
}
func (d Scene) GetUniqueId() string {
	return d.UniqueId
}
func (d *Scene) PopulateDevice()              {}
func (d *Scene) UpdateState()                 {}
func (d *Scene) Subscribe(client mqtt.Client) {}
func (d Select) GetRawId() string {
	return "select"
}
func (d Select) GetUniqueId() string {
	return d.UniqueId
}
func (d *Select) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d *Select) UpdateState()                 {}
func (d *Select) Subscribe(client mqtt.Client) {}
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
func (d *Sensor) UpdateState()                 {}
func (d *Sensor) Subscribe(client mqtt.Client) {}
func (d Siren) GetRawId() string {
	return "siren"
}
func (d Siren) GetUniqueId() string {
	return d.UniqueId
}
func (d *Siren) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d *Siren) UpdateState()                 {}
func (d *Siren) Subscribe(client mqtt.Client) {}
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
func (d *Switch) UpdateState()                 {}
func (d *Switch) Subscribe(client mqtt.Client) {}
func (d Tag) GetRawId() string {
	return "tag"
}
func (d Tag) GetUniqueId() string {
	return ""
}
func (d *Tag) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d *Tag) UpdateState()                 {}
func (d *Tag) Subscribe(client mqtt.Client) {}
func (d Vacuum) GetRawId() string {
	return "vacuum"
}
func (d Vacuum) GetUniqueId() string {
	return d.UniqueId
}
func (d *Vacuum) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}
func (d *Vacuum) UpdateState()                 {}
func (d *Vacuum) Subscribe(client mqtt.Client) {}
