package store

import mqtt "tinygo.org/x/drivers/net/mqtt"

var TopicStore = make(map[string]*func(mqtt.Message, mqtt.Client))
