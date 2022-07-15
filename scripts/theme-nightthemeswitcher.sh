#!/bin/bash

if [ "${#}" -lt 1 ]; then
  echo 'Must pass one or more arguments.'
  exit 1
fi

if [ "${1}" == 'set' ]; then

  dconf write /org/gnome/shell/extensions/nightthemeswitcher/time/ondemand-time "${2}"

elif [ "${1}" == 'get' ]; then

  dconf read /org/gnome/shell/extensions/nightthemeswitcher/time/ondemand-time

else

  echo "Unkown argument ${1}"
  exit 1

fi

exit
