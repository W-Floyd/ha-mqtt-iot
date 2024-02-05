package ExternalDevice

import (
	"encoding/json"
	common "github.com/W-Floyd/ha-mqtt-iot/common"
	"log"
	"time"
)

// //////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
// //////////////////////////////////////////////////////////////////////////////
type Vacuum struct {
	MQTT *MQTTFields `json:"-"` // MQTT configuration parameters
}

func (d *Vacuum) GetRawId() string {
	return "vacuum"
}
func (d *Vacuum) AddMessageHandler() {
	d.MQTT.MessageHandler = MakeMessageHandler(d)
}
func (d Vacuum) GetUniqueId() string {
	return ""
}
func (d *Vacuum) PopulateDevice() {}
func (d *Vacuum) UpdateState()    {}
func (d *Vacuum) Subscribe() {
	c := *d.MQTT.Client
	message, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	token := c.Publish(GetDiscoveryTopic(d), 0, true, message)
	token.Wait()
	time.Sleep(common.HADiscoveryDelay)
	d.UpdateState()
}
func (d *Vacuum) UnSubscribe()       {}
func (d *Vacuum) AnnounceAvailable() {}
func (d *Vacuum) Initialize() {
	d.PopulateDevice()
	d.AddMessageHandler()
	d.PopulateTopics()
}
func (d *Vacuum) PopulateTopics() {}
func (d *Vacuum) SetMQTTFields(fields MQTTFields) {
	*d.MQTT = fields
}
func (d *Vacuum) GetMQTTFields() (fields MQTTFields) {
	return *d.MQTT
}
