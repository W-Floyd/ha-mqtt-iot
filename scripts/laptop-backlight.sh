#!/bin/bash

### DEPENDENCIES_BEGIN
# brightnessctl gsettings bc
### DEPENDENCIES_END

### CONFIG_TEMPLATE_BEGIN
# {
#    "light": [
#        {
#            "name": "___HOSTNAME___ Display",
#            "brightness_command": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "set-brightness"
#            ],
#            "brightness_state": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "get-brightness"
#            ],
#            "command": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "set-state"
#            ],
#            "state": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "get-state"
#            ],
#            "color_temp_command": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "color-temp"
#            ],
#            "color_temp_state": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "color-temp-state"
#            ],
#            "payload_off": "OFF",
#            "payload_on": "ON",
#            "brightness_scale": 120000,
#            "mqtt": {
#                "update_interval": 1
#            }
#        }
#    ]
# }
# ### CONFIG_TEMPLATE_END


function f2i() {
    awk 'BEGIN{for (i=1; i<ARGC;i++)
    printf "%.0f\n", ARGV[i]}' "$@"
}

function __get_state
{
    if [ "$(brightnessctl g)" -gt 0 ]; then
        echo 'ON'
    else
        echo 'OFF'
    fi
}

function __set_state
{
    __state=${1}
    [ -f /var/tmp/backlight_cache ] || echo "1000" > /var/tmp/backlight_cache
    
    case "${__state}" in
        "ON")
            if ! [ "$(brightnessctl g)" -gt 0 ]; then
                brightnessctl -q s "$(cat /var/tmp/backlight_cache)"
            fi
            ;;
        "OFF")
            brightnessctl -q g >/var/tmp/backlight_cache
            brightnessctl -q s 0
            ;;
        *)
            echo "Unknown argument to set-state: ${__state}"
            exit 1
            ;;
    esac
}

function __color-temp
{
    __value=${1}
    __mired="$(bc -l <<<"${__mired_scaler}*${__value}+(${__mired_offset})")"
    if [ "$(f2i "${__mired}")" -gt 420 ]; then
        __mired=420
    fi
    if [ "$(f2i "${__mired}")" -lt 0 ]; then
        __mired=0
    fi
    __kelvin="$(f2i "$(bc -l <<<"1000000/${__mired}")")"
    gsettings set org.gnome.settings-daemon.plugins.color night-light-temperature "${__kelvin}"
}

function __color-temp-state
{
    __kelvin="$(gsettings get org.gnome.settings-daemon.plugins.color night-light-temperature)"
    __mired="$(bc <<<"(1000000/${__kelvin/* /})/${__mired_scaler}-(${__mired_offset})")"
    echo "${__mired}"
}


__mired_offset='20'
__mired_scaler='0.8'

__command="${1}"
__argument="${2}"


if [ "${#}" -lt 1 ]; then
    echo 'Must pass one or more arguments.'
    exit 1
fi


case ${__command} in
    "get-state")
        __get_state
        ;;
    "set-state")
        __set_state "${__argument}"
        ;;
    "get-brightness")
        brightnessctl -q g
        ;;
    "set-brightness")
        brightnessctl -q s "${__argument}"
        ;;
    "color-temp")
        __color-temp ${__argument}
        ;;
    "color-temp-state")
        __color-temp-state
        ;;
    *)
        echo "Unknown command"
        exit 1
esac


exit