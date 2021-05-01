#!/bin/bash
set -e

MQTT_HOST=$(bashio::services mqtt "host")
MQTT_USER=$(bashio::services mqtt "username")
MQTT_PASSWORD=$(bashio::services mqtt "password")

echo '
{
    "mqtt": {
        "broker": "tcp://'${MQTT_HOST}':1883",
        "username": "'${MQTT_USER}'",
        "password": "'${MQTT_PASSWORD}'",
        "node_id": "ha-mqtt-iot-addon",
        "instance_name": "Home Assistant MQTT IOT Addon"
    }
}' > '/app/secrets.json'

/app/out