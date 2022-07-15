package common

import (
	"log"
	"os"
	"time"
)

var (
	Retain           bool          = false
	QoS              byte          = 0
	HADiscoveryDelay time.Duration = 500 * time.Millisecond
	MachineID        string
)

var DebugLog = log.New(os.Stdout, "DEBUG   ", 0)
var ErrorLog = log.New(os.Stdout, "ERROR   ", 0)
var WarnLog = log.New(os.Stdout, "WARN    ", 0)
var CriticalLog = log.New(os.Stdout, "CRITICAL", 0)

const logPrefix = "[ha-mqtt]  "

func LogError(message ...interface{}) {
	if len(message) > 1 {
		for _, mes := range message[:len(message)-2] {
			ErrorLog.Printf("%v%v\n", logPrefix, mes)
		}
	}
	ErrorLog.Fatalf("%v%v\n", logPrefix, message[len(message)-1])
}

func LogDebug(message ...interface{}) {
	for _, mes := range message {
		DebugLog.Printf("%v%v\n", logPrefix, mes)
	}
}
