#!/bin/bash

ls -l "/proc/$(xdotool getwindowpid $(xdotool getwindowfocus))/exe" | sed -e 's|.*/||'

exit