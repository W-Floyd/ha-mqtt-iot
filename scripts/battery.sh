#!/bin/bash

### DEPENDENCIES_BEGIN
# acpi
### DEPENDENCIES_END

### CONFIG_TEMPLATE_BEGIN
# {
#       "sensor": [
#         {
#             "name": "___HOSTNAME___ Battery Load",
#             "object_id": "___HOSTNAME___-battery-load",
#             "unique_id": "___HOSTNAME___-battery-load",
#             "device_class" : "battery",
#             "unit_of_measurement": "%",
#             "icon": "mdi:chip",
#             "state": [
#                 "___SCRIPTS_DIR______SCRIPT_NAME___",
#                 "get-battery-status"
#             ],
#             "mqtt": {
#                 "update_interval": 10
#             }
#         }
#     ],
#     "binary_sensor": [
#         {
#             "name": "___HOSTNAME___ Battery Charging",
#             "object_id": "___HOSTNAME___-battery-charging",
#             "unique_id": "___HOSTNAME___-battery-charging",
#             "icon": "mdi:battery-charging-80",
#             "device_class": "battery_charging",
#             "payload_on": "ON",
#             "payload_off": "OFF",
#             "state": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "get-charging-status" 
#             ],
#             "mqtt": {
#                 "update_interval": 5
#             }            
#         }
#     ]
# }
### CONFIG_TEMPLATE_END

com="${1}"

case "${com}" in
  "get-battery-status")
    cat /sys/class/power_supply/BAT0/capacity
    ;;

  "get-charging-status")
    charging=$(acpi -a | awk '{print $3}')
    if [ "${charging}" == 'on-line' ]; then
      echo "ON"
    else
      echo "OFF"
    fi
    ;;
esac
