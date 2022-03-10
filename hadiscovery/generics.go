package hadiscovery

func GetTopicPrefix[T Device](d T) string {
	return NodeID +d.RawID() + d.UniqueID() + "/"
}