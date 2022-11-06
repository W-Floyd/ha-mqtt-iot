#!/bin/bash

### DEPENDENCIES_BEGIN
# fuser
### DEPENDENCIES_END

### CONFIG_TEMPLATE_BEGIN
# {
#    "binary_sensor": [
#        {
#            "name": "___HOSTNAME___ Camera Used",
#            "object_id": "___HOSTNAME___-camera-used",
#            "unique_id": "___HOSTNAME___-camera-used",
#            "icon": "mdi:webcam",
#            "payload_on": "1",
#            "payload_off": "0",
#            "state": [
#               "___SCRIPTS_DIR______SCRIPT_NAME___",
#               "get" 
#            ],
#            "mqtt": {
#                "update_interval": 5
#            }            
#        }
#     ]
# }
### CONFIG_TEMPLATE_END




com="${1}"

case "${com}" in
"get")
    WEBCAM_USE=$(for i in $(echo /dev/video*); do fuser $i 2>&1; done)
    if [[ -z $WEBCAM_USE ]]; then
        echo 0
    else
        echo 1
    fi
    ;;

esac
