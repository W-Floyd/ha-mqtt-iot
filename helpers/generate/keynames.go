package main

import (
	"log"
	"sort"
)

var keyNames = []string{}

var blacklistKeys = []string{
	"device",
	"availability",
}

func loadKeyNames() {
	for _, devicename := range DeviceNames {
		g, err := JsonExtractor(devicename)
		if err != nil {
			log.Println(err)
			continue
		}

		for k := range g.ChildrenMap() {
			doAdd := true
			for _, b := range blacklistKeys {
				if k == b {
					doAdd = false
				}
			}
			if doAdd {
				keyNames = append(keyNames, k)
			}
		}

	}

	singles := make(map[string]bool)

	for _, n := range keyNames {
		singles[n] = true
	}

	keyNames = []string{}

	for n := range singles {
		keyNames = append(keyNames, n)
	}

	sort.Strings(keyNames)

}
