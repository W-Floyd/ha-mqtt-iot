#!/bin/bash

### DEPENDENCIES_BEGIN
# iw iwconfig
### DEPENDENCIES_END

### CONFIG_TEMPLATE_BEGIN
# {
#     "sensor": 
#     [
#         {
#             "name": "___HOSTNAME___ ___INTERFACE___Wifi Signal",
#             "object_id": "___HOSTNAME___-___INTERFACE___-wifi-signal",
#             "unique_id": "___HOSTNAME___-___INTERFACE___-wifi-signal",
#             "icon": "mdi:wifi-strength-2",
#             "unit_of_measurement": "dBm",
#             "device_class": "signal_strength",
#             "state": [
#                 "___SCRIPTS_DIR______SCRIPT_NAME___",
#                 "get-signal",
#                 "___INTERFACE___"
#             ],
#             "mqtt": {
#                 "update_interval": 10
#             }
#         },
#         {
#             "name": "___HOSTNAME___ ___INTERFACE___Wifi Quality",
#             "object_id": "___HOSTNAME___-___INTERFACE___-wifi-quality",
#             "unique_id": "___HOSTNAME___-___INTERFACE___-wifi-quality",
#             "icon": "mdi:quality-medium",
#             "unit_of_measurement": "dBm",
#             "device_class": "signal_strength",
#             "state": [
#                 "___SCRIPTS_DIR______SCRIPT_NAME___",
#                 "get-quality",
#                 "___INTERFACE___"
#             ],
#             "mqtt": {
#                 "update_interval": 10
#             }
#         }
#     ]
# }
### CONFIG_TEMPLATE_END

### HELPER_BEGIN
# { "___INTERFACE___": {"shell": "iw dev | grep Interface | awk '{print $2}'" }}
### HELPER_END


com="${1}"
INTERFACE="${2}"

STATE=$(cat /sys/class/net/$INTERFACE/operstate)
if [ "${STATE}" = 'down' ]; then
  echo ${STATE}
  exit 0
fi

case "${com}" in
"get-signal")
  INTERFACE=$(iw dev | grep Interface | awk '{print $2}')
  /sbin/iwconfig $INTERFACE | grep Signal | awk -F " " '{print $4}' | cut -d "=" -f2
  ;;

"get-quality")
  INTERFACE=$(iw dev | grep Interface | awk '{print $2}')
  /sbin/iwconfig $INTERFACE 2>&1 | grep Link | cut -d "=" -f2 | cut -d "/" -f1
  ;;

esac
