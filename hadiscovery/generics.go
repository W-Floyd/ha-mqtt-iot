package hadiscovery

func GetTopicPrefix[T Device](d T) string {
	return NodeID + d.RawID() + d.UniqueID() + "/"
}

func GetDiscoveryTopic[T Device](d T) string {
	return DiscoveryPrefix + "/" + d.RawID() + "/" + NodeID + "/" + d.UniqueID() + "/"
}

func GetCommandTopic[T Device](d T) string {
	return GetTopicPrefix(d) + "command"
}

func GetAvailabilityTopic[T Device](d T) string {
	return GetTopicPrefix(d) + "availability"
}
