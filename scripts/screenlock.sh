#!/bin/bash

### DEPENDENCIES_BEGIN
# loginctl dbus-send whoami
### DEPENDENCIES_END

### CONFIG_TEMPLATE_BEGIN
# {
#    "lock": [
#        {
#            "name": "___HOSTNAME___ Screen Lock",
#            "state": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "state"
#            ],
#            "command": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___",
#                "command"
#            ],
#            "payload_lock": "lock",
#            "payload_unlock": "unlock",
#            "state_unlocked": "false",
#            "state_locked": "true",
#            "mqtt": {
#                "update_interval": 1
#            }
#        }
#    ]
# }
# ### CONFIG_TEMPLATE_END


__args="${@}"
__target_user="$(whoami)"
__current_session="$(loginctl list-sessions --no-legend | grep " ${__target_user} " | grep -Eo '^[^ ]*')"

if [ "$(wc -l <<<"${__current_session}")" -gt 1 ]; then
    echo "More than one session logged in for target user ${__target_user}"
    exit 1
fi

__error() {
    if [ -z "${1}" ]; then
        echo "Error parsing ${__args}"
    else
        echo "${1}"
    fi
    exit 1
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
        case "${XDG_CURRENT_DESKTOP}" in
            "GNOME")
                dbus-send --print-reply --session --type=method_call --dest=org.gnome.ScreenSaver /org/gnome/ScreenSaver org.gnome.ScreenSaver.GetActive | tail -n 1 | grep -Eo '[^ ]*$'
                ;;
            "KDE")
                dbus-send --print-reply --session --type=method_call --dest=org.kde.screensaver /ScreenSaver org.freedesktop.ScreenSaver.GetActive | tail -n 1 | grep -Eo '[^ ]*$'
                ;;
            *)
                __error "Unknown desktop environment ${XDG_CURRENT_DESKTOP}"
                ;;
            esac
            ;;
    *)
        __error
        ;;
esac

exit
