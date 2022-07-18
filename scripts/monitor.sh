#!/bin/bash

__monitor_i2c='dev:/dev/i2c-3'
__monitor_dpms='0xd6'
__monitor_brightness='0x10'
__monitor_audio='0x62'
__monitor_standby='5'
__monitor_off='5'
__monitor_on='1'

__mired_offset='20'
__mired_scaler='0.8'

__unknown() {
    echo "Unknown ${1}"
}

f2i() {
    awk 'BEGIN{for (i=1; i<ARGC;i++)
   printf "%.0f\n", ARGV[i]}' "$@"
}

com="${1}"
arg="${2}"

case "${com}" in
"command")
    case "${arg}" in
    "ON")
        ddccontrol -r "${__monitor_dpms}" -w "${__monitor_on}" "${__monitor_i2c}" -f
        xset dpms force on
        ;;
    "OFF")
        (
            #xset dpms force off
            #sleep 0.5s
            #ddccontrol -r "${__monitor_dpms}" -w "${__monitor_standby}" "${__monitor_i2c}" -f
            #sleep 2s
            ddccontrol -r "${__monitor_dpms}" -w "${__monitor_standby}" "${__monitor_i2c}" -f
        ) &
        ;;
    *)
        __unknown "${arg}"
        ;;
    esac
    ;;
"command-state")
    if (ddccontrol -r "${__monitor_dpms}" "${__monitor_i2c}" 2>/dev/null | grep 'DPMS Control - On' -q); then
        echo -n 'ON'
    else
        echo -n 'OFF'
    fi
    ;;
"color-temp")
    __mired="$(bc -l <<<"${__mired_scaler}*${arg}+(${__mired_offset})")"
    if [ "$(f2i "${__mired}")" -gt 420 ]; then
        __mired=420
    fi
    if [ "$(f2i "${__mired}")" -lt 0 ]; then
        __mired=0
    fi
    __kelvin="$(f2i "$(bc -l <<<"1000000/${__mired}")")"
    gsettings set org.gnome.settings-daemon.plugins.color night-light-temperature "${__kelvin}"
    ;;
"color-temp-state")
    __kelvin="$(gsettings get org.gnome.settings-daemon.plugins.color night-light-temperature)"
    __mired="$(bc -l <<<"(1000000/${__kelvin/* /})/${__mired_scaler}-(${__mired_offset})")"
    echo "${__mired}"
    ;;
"brightness")
    ddccontrol -r "${__monitor_brightness}" "${__monitor_i2c}" -w "${arg}"
    ;;
"brightness-state")
    echo -n "$(ddccontrol 2>/dev/null -r "${__monitor_brightness}" "${__monitor_i2c}" | tail -n 1 | grep -o '/[0-9]*/100' | sed -e 's|^/||' -e 's|/.*||')"
    ;;
"audio")
    ddccontrol -r "${__monitor_audio}" "${__monitor_i2c}" -w "${arg}"
    ;;
"audio-state")
    echo -n "$(ddccontrol 2>/dev/null -r "${__monitor_audio}" "${__monitor_i2c}" | tail -n 1 | grep -o '/[0-9]*/100' | sed -e 's|^/||' -e 's|/.*||')"
    ;;
*)
    __unknown "root command ${com}"
    ;;
esac

exit
