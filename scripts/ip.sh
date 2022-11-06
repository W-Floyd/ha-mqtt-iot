#!/bin/bash

### DEPENDENCIES_BEGIN
# ip
### DEPENDENCIES_END

### CONFIG_TEMPLATE_BEGIN
# {
#     "sensor": 
#     [
#         {
#             "name": "___HOSTNAME___ ___INTERFACE___ IP Address",
#             "object_id": "___HOSTNAME___-___INTERFACE___-ip-address",
#             "unique_id": "___HOSTNAME___-___INTERFACE___-ip-address",
#             "icon": "mdi:ip-network",
#             "state": [
#                 "___SCRIPTS_DIR______SCRIPT_NAME___",
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
# { "___INTERFACE___": {"shell": "ls /sys/class/net/ | grep -v docker0 | grep -v lo | grep -v br-b83d983040f6 | grep -v bond0 | grep -v bonding_masters" }}
### HELPER_END



__interface="${1}"
__state=$(cat "/sys/class/net/${__interface}/operstate")

if [ "${__state}" = 'down' ]; then
    echo "${__state}"
else
    ip -j address show "${__interface}" | jq -r '.[0].addr_info[0].local'
fi
