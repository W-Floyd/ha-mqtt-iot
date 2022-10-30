#!/bin/bash

# sysstat required for cpu
# acpi required for battery


com="${1}"
arg="${2}"

case "${com}" in
  "get-cpu")
    mpstat | tail -n 1 | sed 's#,#.#g' | awk '{print $3+$4+$5+$6+$7+$8}'
    ;;

  "get-ram")
    LANG= LC_ALL= free | grep Mem | awk '{$x=sprintf("%.2f",100-($7/$2*100))} {print $x}'
    ;;
  
  "get-battery-status")
    cat /sys/class/power_supply/BAT0/capacity
    ;;

  "get-charging-status")
    charging=$(acpi -a | awk '{print $3}')
    if [ "${charging}" == 'on-line' ]; then
        echo "ON"
    else
        echo "OFF"
    fi
    ;;
  "get-board")
    if [ "${arg}" == 'pi' ]; then
      cat /sys/firmware/devicetree/base/model
    else
      echo "ToDo"
    fi
    ;;

esac













