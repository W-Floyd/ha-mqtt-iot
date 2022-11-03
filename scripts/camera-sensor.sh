#!/bin/bash

com="${1}"

case "${com}" in
"get")
    WEBCAM_USE=$(for i in $(echo /dev/video*); do fuser $i 2>&1; done)
    if [[ -z $WEBCAM_USE ]]; then
        echo 0
    else
        echo 1
    fi
    ;;

esac
