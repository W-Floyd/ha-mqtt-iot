#!/bin/bash

if [ "${#}" -lt 1 ]; then
  echo 'Must pass one or more arguments.'
  exit
fi

# set dark
# set light
if [ "${1}" == 'set' ]; then
  dconf write /org/gnome/terminal/legacy/theme-variant "'${2}'"
  gsettings set org.gnome.desktop.interface color-scheme "prefer-${2}"
# get
elif [ "${1}" == 'get' ]; then
  __state="$(gsettings get org.gnome.desktop.interface color-scheme)"
  if [ "${__state}" == "'prefer-dark'" ]; then
    echo -n 'ON'
  else
    echo -n 'OFF'
  fi
fi

exit
