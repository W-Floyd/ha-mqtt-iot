package ExternalDevice

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
type Device interface {
	GetRawId() string
	GetUniqueId() string
	PopulateDevice()
	PopulateTopics()
	UpdateState()
	Subscribe()
	UnSubscribe()
	AddMessageHandler()
	SetMQTTFields(MQTTFields)
	GetMQTTFields() MQTTFields
}
