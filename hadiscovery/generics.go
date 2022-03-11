package hadiscovery

import "strings"

func GetTopicPrefix(d Device) string {
	uID := d.GetUniqueId()
	if uID != "" {
		uID = uID + "/"
	}
	return NodeID + "/" + d.GetRawId() + "/" + uID
}

func GetDiscoveryTopic(d Device) string {
	uID := d.GetUniqueId()
	if uID != "" {
		uID = uID + "/"
	}
	return DiscoveryPrefix + "/" + d.GetRawId() + "/" + NodeID + "/" + uID
}

func GetTopic(d Device, rawTopicString string) string {
	return GetTopicPrefix(d) + strings.TrimSuffix(rawTopicString, "_topic")
}
