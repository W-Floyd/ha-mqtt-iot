#!/bin/bash

# sysstat required for cpu

### DEPENDENCIES_BEGIN
# mpstat
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
#             "name": "___HOSTNAME___ Mainboard",
#             "object_id": "___HOSTNAME___-mainboard-usage",
#             "unique_id": "___HOSTNAME___-mainboard-usage",
#             "icon": "mdi:expansion-card-variant",
#             "state": [
#                 "___SCRIPTS_DIR______SCRIPT_NAME___",
#                 "get-board"
#             ],
#             "mqtt": {
#                 "update_interval": 60
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
"get-board")
  if [ -f /sys/firmware/devicetree/base/model ]; then
    cat /sys/firmware/devicetree/base/model
  else
    VENDOR="unknown"
    BOARD="unknown"
    [ -f /sys/class/dmi/id/sys_vendor ] && VENDOR=$(cat /sys/class/dmi/id/sys_vendor)
    [ -f /sys/class/dmi/id/board_name ] && BOARD=$(cat /sys/class/dmi/id/board_name)
    echo "$VENDOR $BOARD"
  fi
  ;;

esac
