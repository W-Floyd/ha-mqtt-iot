#!/bin/bash

### DEPENDENCIES_BEGIN
# dconf 
### DEPENDENCIES_END

### CONFIG_TEMPLATE_BEGIN
# {
#    "switch": [
#        {
#            "name": "___HOSTNAME___ Dark Mode",
#            "object_id": "___HOSTNAME___-dark-mode",
#            "unique_id": "___HOSTNAME___-dark-mode",
#            "icon": "mdi:theme-light-dark",
#            "payload_on": "'night'",
#            "payload_off": "'day'",
#            "command": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "set"
#            ],
#            "state": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "get"
#            ],
#            "mqtt": {
#                "update_interval": 1
#            }
#        }
#    ]
# }
# ### CONFIG_TEMPLATE_END

if [ "${#}" -lt 1 ]; then
  echo 'Must pass one or more arguments.'
  exit 1
fi

if [ "${1}" == 'set' ]; then

  dconf write /org/gnome/shell/extensions/nightthemeswitcher/time/ondemand-time "${2}"

elif [ "${1}" == 'get' ]; then

  dconf read /org/gnome/shell/extensions/nightthemeswitcher/time/ondemand-time

else

  echo "Unkown argument ${1}"
  exit 1

fi

exit
