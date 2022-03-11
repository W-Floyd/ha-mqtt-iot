package hadiscovery

type Device interface {
	GetRawId() string
	GetUniqueId() string
	PopulateDevice()
	PopulateTopics()
	UpdateState()
	Subscribe()
	AddMessageHandler()
}
