#!/bin/bash
### DEPENDENCIES_BEGIN
# systemctl
### DEPENDENCIES_END

### CONFIG_TEMPLATE_BEGIN
# {
#    "binary_sensor": [
#        {
#            "name": "___HOSTNAME___ ___SERVICE___ Service",
#            "object_id": "___HOSTNAME___-___SERVICE___-service",
#            "unique_id": "___HOSTNAME___-___SERVICE___-service",
#            "icon": "mdi:crosshairs",
#            "payload_on": "ON",
#            "payload_off": "OFF",
#            "state": [
#               "___SCRIPTS_DIR______SCRIPT_NAME___",
#               "___SERVICE___" 
#            ],
#            "mqtt": {
#                "update_interval": 5
#            }            
#        }
#     ]
# }
### CONFIG_TEMPLATE_END

### HELPER_BEGIN
# { "___SERVICE___": [ "sshd", "firewalld", "bluetooth" ]}
### HELPER_END

# changed: https://www.freedesktop.org/software/systemd/man/systemctl.html#Exit%20status
# added exit codes of systemd 

systemctl status ${1} &>/dev/null
RET=$(echo $?)
case "${RET}" in
    "0") echo ON ;;
    "3") echo OFF ;;
    *) echo unavailable ;;
esac

