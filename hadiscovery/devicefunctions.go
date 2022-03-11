package hadiscovery

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
func (d DeviceTracker) GetRawId() string {
	return "device_tracker"
}
func (d DeviceTracker) GetUniqueId() string {
	return ""
}
func (d *DeviceTracker) PopulateDevice() {}
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
func (d Scene) GetRawId() string {
	return "scene"
}
func (d Scene) GetUniqueId() string {
	return d.UniqueId
}
func (d *Scene) PopulateDevice() {}
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
