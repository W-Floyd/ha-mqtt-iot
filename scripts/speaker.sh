#!/bin/bash

### DEPENDENCIES_BEGIN
# pactl
### DEPENDENCIES_END

### CONFIG_TEMPLATE_BEGIN
# {
#     "light": 
#     [
#         {
#            "name": "___HOSTNAME___ Volume",
#            "icon": "mdi:volume-high",
#            "brightness_command": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "set-volume"
#            ],
#            "brightness_state": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "get-volume"
#            ],
#           "command": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "set-state"
#            ],
#            "state": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "get-state"
#            ],
#            "payload_off": "OFF",
#            "payload_on": "ON",
#            "brightness_scale": 100,
#            "mqtt": {
#                "update_interval": 1
#            }
#        }
#     ]
# }
### CONFIG_TEMPLATE_END


function __set_speaker_state {
  case "${1}" in
  "ON")
    pactl set-sink-mute @DEFAULT_SINK@ 0
    ;;
  "OFF")
    pactl set-sink-mute @DEFAULT_SINK@ 1
    ;;
  *)
    __unknown "${__argument}"
    ;;
  esac
}

function __get_speaker_state {
  yesno=$(LANG='' LC_ALL='' pactl get-sink-mute @DEFAULT_SINK@ | awk '{print $2}')
  if [ "${yesno}" == 'yes' ]; then
    echo "OFF"
  else
    echo "ON"
  fi

}

function __set_speaker_volume {
  pactl set-sink-volume @DEFAULT_SINK@ "${1}%"
}

function __get_speaker_volume {
  percent=$(LANG='' LC_ALL='' pactl get-sink-volume @DEFAULT_SINK@ | awk '{print $5}')
  echo "${percent%?}"
}

__command="${1}"
__argument="${2}"

case "${__command}" in
"set-state")
  __set_speaker_state "${__argument}"
  ;;

"get-state")
  __get_speaker_state
  ;;

"set-volume")
  __set_speaker_volume "${__argument}"
  ;;

"get-volume")
  __get_speaker_volume
  ;;

esac
