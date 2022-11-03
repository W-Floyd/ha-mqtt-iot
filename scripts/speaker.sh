#!/bin/bash

function __set-speaker-state {
  case "${1}" in
  "ON")
    pactl set-sink-mute @DEFAULT_SINK@ 0
    ;;
  "OFF")
    pactl set-sink-mute @DEFAULT_SINK@ 1
    ;;
  *)
    __unknown "${arg}"
    ;;
  esac
}

function __get-speaker-state {
  yesno=$(LANG= LC_ALL= pactl get-sink-mute @DEFAULT_SINK@ | awk '{print $2}')
  if [ "${yesno}" == 'yes' ]; then
    echo "OFF"
  else
    echo "ON"
  fi

}

function __set-speaker-volume {
  pactl set-sink-volume @DEFAULT_SINK@ "${1}%"
}

function __get-speaker-volume {
  percent=$(LANG= LC_ALL= pactl get-sink-volume @DEFAULT_SINK@ | awk '{print $5}')
  val=${percent%?}
  echo $val
}

com="${1}"
arg="${2}"

case "${com}" in
"set-state")
  __set-speaker-state ${arg}
  ;;

"get-state")
  mystate=$(__get-speaker-state)
  echo $mystate
  ;;

"set-volume")
  __set-speaker-volume ${arg}
  ;;

"get-volume")
  myloudness=$(__get-speaker-volume)
  echo $myloudness
  ;;

esac
