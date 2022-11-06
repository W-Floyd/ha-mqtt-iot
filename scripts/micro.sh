#!/bin/bash

### DEPENDENCIES_BEGIN
# pactl
### DEPENDENCIES_END

### CONFIG_TEMPLATE_BEGIN
# {
#     "light": 
#     [
#        {
#            "name": "___HOSTNAME___ Microphone",
#            "icon": "mdi:microphone",
#            "brightness_command": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "set-sensitivity"
#            ],
#            "brightness_state": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "get-sensitivity"
#            ],
#            "command": [
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

function __set-microphone-state {
  case "${1}" in
  "ON")
    pactl set-source-mute @DEFAULT_SOURCE@ 0
    ;;
  "OFF")
    pactl set-source-mute @DEFAULT_SOURCE@ 1
    ;;
  *)
    __unknown "${arg}"
    ;;
  esac
}

function __get-microphone-state {
  yesno=$(LANG= LC_ALL= pactl get-source-mute @DEFAULT_SOURCE@ | awk '{print $2}')
  if [ "${yesno}" == 'yes' ]; then
    echo "OFF"
  else
    echo "ON"
  fi

}

function __set-microphone-sensitivity {
  pactl set-source-volume @DEFAULT_SOURCE@ "${1}%"
}

function __get-microphone-sensitivity {
  percent=$(LANG= LC_ALL= pactl get-source-volume @DEFAULT_SOURCE@ | awk '{print $5}')
  val=${percent%?}
  echo $val
}

com="${1}"
arg="${2}"

case "${com}" in
"set-state")
  __set-microphone-state ${arg}
  ;;

"get-state")
  mystate=$(__get-microphone-state)
  echo $mystate
  ;;

"set-sensitivity")
  __set-microphone-sensitivity ${arg}
  ;;

"get-sensitivity")
  mysensi=$(__get-microphone-sensitivity)
  echo $mysensi
  ;;

esac
