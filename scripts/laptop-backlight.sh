#!/bin/bash

if [ "${#}" -lt 1 ]; then
    echo 'Must pass one or more arguments.'
    exit 1
fi

if [ "${1}" == 'get-state' ]; then
    if [ "$(brightnessctl g)" -gt 0 ]; then
        echo 'ON'
    else
        echo 'OFF'
    fi
elif [ "${1}" == 'set-state' ]; then
    if [ "${#}" -lt 2 ]; then
        echo 'Must pass 2 arguments to set-state'
        exit 1
    fi
    if [ "${2}" == 'OFF' ]; then
        brightnessctl -q g >/var/tmp/backlight_cache
        brightnessctl -q s 0
    elif [ "${2}" == 'ON' ]; then
        if ! [ "$(brightnessctl g)" -gt 0 ]; then
            brightnessctl -q s "$(cat /var/tmp/backlight_cache)"
        fi
    else
        echo "Unknown argument to set-state: ${2}"
        exit 1
    fi
elif [ "${1}" == 'get-brightness' ]; then
    brightnessctl -q g
elif [ "${1}" == 'set-brightness' ]; then
    if [ "${#}" -lt 2 ]; then
        echo 'Must pass 2 arguments to set-brightness'
        exit 1
    fi
    brightnessctl -q s "${2}"
else
    echo "Unknown argument: ${*}"
fi

exit
