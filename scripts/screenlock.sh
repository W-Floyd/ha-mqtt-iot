#!/bin/bash

__args="${@}"

__target_user="$(whoami)"

__current_session="$(loginctl list-sessions --no-legend | grep " ${__target_user} " | grep -Eo '^[^ ]*')"

if [ "$(wc -l <<<"${__current_session}")" -gt 1 ]; then
    echo "More than one session logged in for target user ${__target_user}"
    exit 1
fi

__error() {
    echo "Error parsing ${__args}"
}

com="${1}"
arg="${2}"

case "${com}" in

"command")

    case "${arg}" in

    "lock")
        loginctl lock-session "${__current_session}"
        ;;

    "unlock")
        loginctl unlock-session "${__current_session}"
        ;;

    *)
        __error
        ;;

    esac

    ;;

"state")
    dbus-send --print-reply --session --type=method_call --dest=org.gnome.ScreenSaver /org/gnome/ScreenSaver org.gnome.ScreenSaver.GetActive | tail -n 1 | grep -Eo '[^ ]*$'
    ;;

*)
    __error
    ;;

esac

exit
