#!/bin/bash

__monitor_i2c='dev:/dev/i2c-3'
__monitor_dpms='0xd6'
__monitor_brightness='0x10'
__monitor_off='4'
__monitor_on='1'

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
        xset dpms force on
        ;;
    "OFF")
        (
            xset dpms force off
            sleep 0.5s
            ddccontrol -r "${__monitor_dpms}" -w "${__monitor_off}" "${__monitor_i2c}" -f
        ) &
        ;;
    *)
        __unknown "${arg}"
        ;;
    esac
    ;;
"command-state")
    echo -n "$(xset q | grep 'Monitor is' | sed -e 's/.* //' | tr '[:lower:]' '[:upper:]')"
    ;;
"color-temp")
    v="$(f2i "$(bc -l <<<"1000000/${arg}")")"
    ./scripts/run-in-user-session.sh gsettings set org.gnome.settings-daemon.plugins.color night-light-temperature "${v}"
    ;;
"color-temp-state")
    v="$(./scripts/run-in-user-session.sh gsettings get org.gnome.settings-daemon.plugins.color night-light-temperature)"
    echo -n "$(f2i "$(bc -l <<<"1000000/${v/* /}")")"
    ;;
"brightness")
    ddccontrol -r "${__monitor_brightness}" "${__monitor_i2c}" -w "${arg}"
    ;;
"brightness-state")
    echo -n "$(ddccontrol 2>/dev/null -r "${__monitor_brightness}" "${__monitor_i2c}" | tail -n 1 | grep -o '/[0-9]*/100' | sed -e 's|^/||' -e 's|/.*||')"
    ;;

*)
    __unknown "root command ${com}"
    ;;
esac

exit
