package sensordevice

import (
	"log"
	"os/exec"

	"github.com/W-Floyd/ha-mqtt-iot/hadiscovery"
	"github.com/W-Floyd/ha-mqtt-iot/iotconfig/common"
)

type SensorHA struct {
	Info              common.InfoIcon `json:"info"`
	CommandState      []string        `json:"command_state"`
	UnitOfMeasurement string          `json:"unit_of_measurement,omitempty"`
	UpdateInterval    float64         `json:"update_interval"`
	ForceUpdateMQTT   bool            `json:"force_update"`
}

func (sw SensorHA) constructStateFunc() (f func() string) {
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

func (se SensorHA) Translate() hadiscovery.Sensor {
	nse := hadiscovery.Sensor{}
	nse.UpdateInterval = se.UpdateInterval
	nse.ExpireAfter = int(nse.UpdateInterval + 1)
	nse.ForceUpdateMQTT = se.ForceUpdateMQTT
	nse.UnitOfMeasurement = se.UnitOfMeasurement
	if !se.ForceUpdateMQTT {
		nse.ExpireAfter = 0
	}
	nse.Name = se.Info.Name
	nse.UniqueID = se.Info.ID + "_" + hadiscovery.NodeID
	if se.Info.Icon != "" {
		nse.Icon = se.Info.Icon
	}

	if len(se.CommandState) > 0 {
		nse.StateFunc = se.constructStateFunc()
	} else {
		log.Fatalln("Missing state func, needed for sensors!")
	}

	nse.Initialize()

	return nse
}
