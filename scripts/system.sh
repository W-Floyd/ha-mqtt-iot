#!/bin/bash

# sysstat required for cpu
# acpi required for battery

### DEPENDENCIES_BEGIN
# mpstat acpi
### DEPENDENCIES_END

### CONFIG_TEMPLATE_BEGIN
# {
#       "sensor": [
#         {
#             "name": "___HOSTNAME___ CPU Usage",
#             "object_id": "___HOSTNAME___-cpu-usage",
#             "unique_id": "___HOSTNAME___-cpu-usage",
#             "icon": "mdi:memory",
#             "unit_of_measurement": "%",
#             "state": [
#                 "___SCRIPTS_DIR______SCRIPT_NAME___",
#                 "get-cpu"
#             ],
#             "mqtt": {
#                 "update_interval": 10
#             }
#         },
#         {
#             "name": "___HOSTNAME___ RAM Usage",
#             "object_id": "___HOSTNAME___-ram-usage",
#             "unique_id": "___HOSTNAME___-ram-usage",
#             "unit_of_measurement": "%",
#             "icon": "mdi:chip",
#             "state": [
#                 "___SCRIPTS_DIR______SCRIPT_NAME___",
#                 "get-ram"
#             ],
#             "mqtt": {
#                 "update_interval": 10
#             }
#         },
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
# ### CONFIG_TEMPLATE_END








com="${1}"
arg="${2}"

case "${com}" in
"get-cpu")
  mpstat 1 1 | tail -n 1 | sed 's#,#.#g' | awk '{print $3+$4+$5+$6+$7+$8}'
  ;;

"get-ram")
  LANG= LC_ALL= free | grep Mem | awk '{$x=sprintf("%.2f",100-($7/$2*100))} {print $x}'
  ;;

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
"get-board")
  if [ "${arg}" == 'pi' ]; then
    cat /sys/firmware/devicetree/base/model
  else
    echo "ToDo"
  fi
  ;;

esac
