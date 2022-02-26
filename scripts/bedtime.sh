#!/bin/bash

__unknown() {
    echo "Unknown ${1}"
}

com="${1}"
arg="${2}"

case "${com}" in
"command")
    case "${arg}" in
    "ON")
        dconf write /org/gnome/shell/extensions/bedtime-mode/bedtime-mode-active true
        ;;
    "OFF")
        dconf write /org/gnome/shell/extensions/bedtime-mode/bedtime-mode-active false
        ;;
    *)
        __unknown "${arg}"
        ;;
    esac
    ;;
"command-state")
    if [ "$(dconf read /org/gnome/shell/extensions/bedtime-mode/bedtime-mode-active)" == "false" ]; then
        echo -n 'OFF'
    else
        echo -n 'ON'
    fi
    ;;
*)
    __unknown "root command ${com}"
    ;;
esac

exit
