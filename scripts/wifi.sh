#!/bin/bash

com="${1}"
INTERFACE="${2}"

STATE=$(cat /sys/class/net/$INTERFACE/operstate)
if [ "${STATE}" = 'down' ]; then
  echo ${STATE}
  exit 0
fi

case "${com}" in
"get-signal")
  INTERFACE=$(iw dev | grep Interface | awk '{print $2}')
  /sbin/iwconfig $INTERFACE | grep Signal | awk -F " " '{print $4}' | cut -d "=" -f2
  ;;

"get-quality")
  INTERFACE=$(iw dev | grep Interface | awk '{print $2}')
  /sbin/iwconfig $INTERFACE 2>&1 | grep Link | cut -d "=" -f2 | cut -d "/" -f1
  ;;

esac
