#!/bin/bash

__interface="${1}"
__state=$(cat "/sys/class/net/${__interface}/operstate")

if [ "${__state}" = 'down' ]; then
    echo "${__state}"
else
    ip -j address show "${__interface}" | jq -r '.[0].addr_info[0].local'
fi
