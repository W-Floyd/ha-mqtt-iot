#!/bin/bash

case "${1}" in
  "get-cpu-temp")
    temp=$(cat /sys/class/thermal/thermal_zone0/temp)
    echo "print(round($temp/1000,1))" | python3
    ;;
esac

