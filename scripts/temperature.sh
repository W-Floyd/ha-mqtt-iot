#!/bin/bash

### DEPENDENCIES_BEGIN
# python3
### DEPENDENCIES_END

### CONFIG_TEMPLATE_BEGIN
# {
#     "sensor": [
#         {
#             "name": "___HOSTNAME___ CPU Temperature",
#             "object_id": "___HOSTNAME___-cpu-temperature",
#             "unique_id": "___HOSTNAME___-cpu-temperature",
#             "icon": "mdi:thermometer",
#             "device_class": "temperature",
#             "unit_of_measurement": "°C",
#             "state": [
#                 "___SCRIPTS_DIR______SCRIPT_NAME___",
#                 "get-cpu-temp"
#             ],
#             "mqtt": {
#                 "update_interval": 10
#             }
#       }
#     ]
# }
# ### CONFIG_TEMPLATE_END

case "${1}" in
"get-cpu-temp")
  temp=$(cat /sys/class/thermal/thermal_zone0/temp)
  echo "print(round($temp/1000,1))" | python3
  ;;
esac
