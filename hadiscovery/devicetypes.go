package hadiscovery

type AlarmControlPanel struct {
	uniqueID string
	retain   bool
}

func (d AlarmControlPanel) RawID() string {
	return "alarm_control_panel"
}
func (d AlarmControlPanel) UniqueID() string {
	return d.uniqueID
}

type BinarySensor struct {
	uniqueID string
	retain   bool
}

func (d BinarySensor) RawID() string {
	return "binary_sensor"
}
func (d BinarySensor) UniqueID() string {
	return d.uniqueID
}

type Camera struct {
	uniqueID string
	retain   bool
}

func (d Camera) RawID() string {
	return "camera"
}
func (d Camera) UniqueID() string {
	return d.uniqueID
}

type Cover struct {
	uniqueID string
	retain   bool
}

func (d Cover) RawID() string {
	return "cover"
}
func (d Cover) UniqueID() string {
	return d.uniqueID
}

type DeviceTracker struct {
	uniqueID string
	retain   bool
}

func (d DeviceTracker) RawID() string {
	return "device_tracker"
}
func (d DeviceTracker) UniqueID() string {
	return d.uniqueID
}

type DeviceTrigger struct {
	uniqueID string
	retain   bool
}

func (d DeviceTrigger) RawID() string {
	return "device_trigger"
}
func (d DeviceTrigger) UniqueID() string {
	return d.uniqueID
}

type Fan struct {
	uniqueID string
	retain   bool
}

func (d Fan) RawID() string {
	return "fan"
}
func (d Fan) UniqueID() string {
	return d.uniqueID
}

type Humidifier struct {
	uniqueID string
	retain   bool
}

func (d Humidifier) RawID() string {
	return "humidifier"
}
func (d Humidifier) UniqueID() string {
	return d.uniqueID
}

type Climate struct {
	uniqueID string
	retain   bool
}

func (d Climate) RawID() string {
	return "climate"
}
func (d Climate) UniqueID() string {
	return d.uniqueID
}

type Light struct {
	uniqueID string
	retain   bool
}

func (d Light) RawID() string {
	return "light"
}
func (d Light) UniqueID() string {
	return d.uniqueID
}

type Lock struct {
	uniqueID string
	retain   bool
}

func (d Lock) RawID() string {
	return "lock"
}
func (d Lock) UniqueID() string {
	return d.uniqueID
}

type Number struct {
	uniqueID string
	retain   bool
}

func (d Number) RawID() string {
	return "number"
}
func (d Number) UniqueID() string {
	return d.uniqueID
}

type Scene struct {
	uniqueID string
	retain   bool
}

func (d Scene) RawID() string {
	return "scene"
}
func (d Scene) UniqueID() string {
	return d.uniqueID
}

type Select struct {
	uniqueID string
	retain   bool
}

func (d Select) RawID() string {
	return "select"
}
func (d Select) UniqueID() string {
	return d.uniqueID
}

type Sensor struct {
	uniqueID string
	retain   bool
}

func (d Sensor) RawID() string {
	return "sensor"
}
func (d Sensor) UniqueID() string {
	return d.uniqueID
}

type Switch struct {
	uniqueID string
	retain   bool
}

func (d Switch) RawID() string {
	return "switch"
}
func (d Switch) UniqueID() string {
	return d.uniqueID
}

type Tag struct {
	uniqueID string
	retain   bool
}

func (d Tag) RawID() string {
	return "tag"
}
func (d Tag) UniqueID() string {
	return d.uniqueID
}

type Vacuum struct {
	uniqueID string
	retain   bool
}

func (d Vacuum) RawID() string {
	return "vacuum"
}
func (d Vacuum) UniqueID() string {
	return d.uniqueID
}

type Device interface {
	AlarmControlPanel | BinarySensor | Camera | Cover | DeviceTracker | DeviceTrigger | Fan | Humidifier | Climate | Light | Lock | Number | Scene | Select | Sensor | Switch | Tag | Vacuum
	GetDiscoveryTopic() string
	RawID() string
	UniqueID() string
}
type WithState interface {
	AlarmControlPanel | BinarySensor | Cover | Fan | Humidifier | Light | Lock | Number | Select | Sensor | Switch
}
