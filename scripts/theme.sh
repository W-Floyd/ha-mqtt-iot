#!/bin/bash

if [ "${#}" -lt 1 ]; then
  echo 'Must pass one or more arguments.'
  exit
fi

get_uuid() {
  # Print the UUID linked to the profile name sent in parameter
  local profile_name=$1
  profiles=($(gsettings get org.gnome.Terminal.ProfilesList list | tr -d "[]\',"))
  for i in ${!profiles[*]}; do
    local uuid="$(dconf read /org/gnome/terminal/legacy/profiles:/:${profiles[i]}/visible-name)"
    if [[ "${uuid,,}" = "'${profile_name,,}'" ]]; then
      echo "${profiles[i]}"
      return 0
    fi
  done
  echo "$profile_name"
}

if [ "${1}" == 'set' ]; then

  shift

  UUID=$(get_uuid $1)

  gsettings set org.gnome.Terminal.ProfilesList default "${UUID}"

  dconf write /org/gnome/terminal/legacy/theme-variant "'${1}'"
  gsettings set org.gnome.desktop.interface color-scheme "prefer-${1}"

  gdbus introspect -e -x -d org.gnome.Terminal -o /org/gnome/Terminal/window |
    xmllint --xpath '/node/node/@name' - |
    sed -e 's/name="\([0-9]*\)"/\1\n/g' -e 's/ //g' |
    while read -r __node; do
      gdbus call -e -d org.gnome.Terminal -o /org/gnome/Terminal/window/${__node} -m org.gtk.Actions.SetState profile "<'${UUID}'>" "{}" >/dev/null
    done

elif [ "${1}" == 'get' ]; then
  __state="$(gsettings get org.gnome.desktop.interface color-scheme)"
  if [ "${__state}" == "'prefer-dark'" ]; then
    echo -n 'ON'
  else
    echo -n 'OFF'
  fi

fi

exit
