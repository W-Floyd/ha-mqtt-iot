package store

import mqtt "github.com/eclipse/paho.mqtt.golang"

var TopicStore = make(map[string]*func(mqtt.Message, mqtt.Client))
