#!/bin/bash

INTERFACE=${1}

STATE=$(cat /sys/class/net/$INTERFACE/operstate)
if [ "${STATE}" = 'down' ]; then
    echo ${STATE}
else
    echo $(ip -j address show $INTERFACE | jq -r '.[0].addr_info[0].local')
fi
