package iotconfig

import (
	"log"
	"os/exec"

	"../hadiscovery"
)

func (sw BinarySensorsHA) constructStateFunc() (f func() string) {
	var err error
	return func() string {
		var out []byte
		if len(sw.CommandState) > 1 {
			out, err = exec.Command(sw.CommandState[0], sw.CommandState[1:]...).Output()
		} else {

			out, err = exec.Command(sw.CommandState[0]).Output()
		}
		if err != nil {
			log.Printf("%s", err)
		}
		return string(out)
	}
}

func (bse BinarySensorsHA) Translate() hadiscovery.BinarySensor {
	nse := hadiscovery.BinarySensor{}
	nse.UpdateInterval = bse.UpdateInterval
	nse.ExpireAfter = int(nse.UpdateInterval + 1)
	nse.ForceUpdateMQTT = bse.ForceUpdateMQTT
	if !bse.ForceUpdateMQTT {
		nse.ExpireAfter = 0
	}
	nse.Name = bse.Info.Name
	nse.UniqueID = bse.Info.ID + "_" + hadiscovery.NodeID
	if bse.Info.DeviceClass != "" {
		nse.DeviceClass = bse.Info.DeviceClass
	}

	if len(bse.CommandState) > 0 {
		nse.StateFunc = bse.constructStateFunc()
	} else {
		log.Fatalln("Missing state func, needed for binary sensors!")
	}

	nse.Initialize()

	return nse
}
