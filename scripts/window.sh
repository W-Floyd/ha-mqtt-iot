#!/bin/bash

file "/proc/$(xdotool getwindowpid "$(xdotool getwindowfocus)")/exe" | sed -e 's|.*/||'

exit