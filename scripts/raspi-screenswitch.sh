 #!/bin/bash

# The command vcgencmd is may be unique to raspberry pis


### DEPENDENCIES_BEGIN
# vcgencmd
### DEPENDENCIES_END

### CONFIG_TEMPLATE_BEGIN
# {
#    "switch": [
#        {
#            "name": "___HOSTNAME___ Pi Display",
#            "object_id": "___HOSTNAME___-pi-display",
#            "unique_id": "___HOSTNAME___-pi-display",
#            "icon": "mdi:cellphone-screenshot",
#            "payload_on": "ON",
#            "payload_off": "OFF",
#            "command": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "set"
#            ],
#            "state": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "get"
#            ],
#            "mqtt": {
#                "update_interval": 10
#            }
#        }
#    ]
# }
# ### CONFIG_TEMPLATE_END}

__command="${1}"
__argument="${2}"

if [ "${__command}" == 'set' ]; then
    
    case ${__argument} in
        "ON")
            RAWRES=$(/usr/bin/vcgencmd display_power 1)
            ;;
        "OFF")
            RAWRES=$(/usr/bin/vcgencmd display_power 0)
            ;;
        *)
            echo "argument unknown"
            ;;
    esac
else
    RAWRES=$(/usr/bin/vcgencmd display_power)
fi
RES=$(echo $RAWRES | cut -d "=" -f2)
if [ $RES -eq 1 ]; then
    echo "ON"
else
    echo "OFF"
fi








