package main

import (
	"strings"
)

func IsCommand(name string, d Device) bool {

	retval := (!strings.Contains(name, "state")) &&
		(!strings.Contains(name, "availability")) &&
		(!d.JSONContainer.Exists("set_" + name)) &&
		(!strings.Contains(name, "status")) &&
		(!strings.Contains(name, "current")) &&
		(name != "topic")

	return retval
}
