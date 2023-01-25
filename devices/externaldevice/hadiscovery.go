package ExternalDevice

// DiscoveryPrefix is the prefix that HA uses to discover on.
// Does not usually need changing.
const DiscoveryPrefix = "homeassistant"

// SWVersion is the software version.
// TODO - Move this elsewhere maybe?
var SWVersion = "0.6.0"

var SoftwareName = "Homeassistant MQTT IOT"

// InstanceName is the instance name, helpful for identifying a given client
var InstanceName = "Homeassistant MQTT IOT"

// NodeID is the Node ID, that is, what that node connects under.
var NodeID = "ha-mqtt-iot"

var Manufacturer = "William Floyd"

///////////////////

var stateStore = initStore()
