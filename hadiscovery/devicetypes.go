package hadiscovery

type DeviceAlarmControlPanel struct {
	UniqueID string
}
type DeviceBinarySensor struct {
	UniqueID string
}
type DeviceCamera struct {
	UniqueID string
}
type DeviceCover struct {
	UniqueID string
}
type DeviceDeviceTracker struct {
	UniqueID string
}
type DeviceDeviceTrigger struct {
	UniqueID string
}
type DeviceFan struct {
	UniqueID string
}
type DeviceHumidifier struct {
	UniqueID string
}
type DeviceClimate struct {
	UniqueID string
}
type DeviceLight struct {
	UniqueID string
}
type DeviceLock struct {
	UniqueID string
}
type DeviceNumber struct {
	UniqueID string
}
type DeviceScene struct {
	UniqueID string
}
type DeviceSelect struct {
	UniqueID string
}
type DeviceSensor struct {
	UniqueID string
}
type DeviceSwitch struct {
	UniqueID string
}
type DeviceTag struct {
	UniqueID string
}
type DeviceVacuum struct {
	UniqueID string
}
type Device interface {
	int
}
