package common

import "time"

var (
	Retain           bool          = true
	QoS              byte          = 0
	HADiscoveryDelay time.Duration = 500 * time.Millisecond
)
