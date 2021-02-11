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

  gsettings set org.gnome.Terminal.ProfilesList default $UUID

  echo "user_pref(\"devtools.theme\", \"${1}\");
user_pref(\"extensions.activeThemeID\", \"firefox-compact-${1}@mozilla.org\");
user_pref(\"toolkit.legacyUserProfileCustomizations.stylesheets\", true);
user_pref(\"browser.tabs.drawInTitlebar\", true);
user_pref(\"browser.uidensity\", 0);" >~/.mozilla/firefox/pcmd5xvl.default-release/user.js

  if [ $1 == "dark" ]; then
    THEME='WhiteSur-dark-solid-purple'
    ICON='Numix-Square'
    WALLPAPER='/home/william/Pictures/Wallpapers/dark.jpg'
    gsettings set com.solus-project.budgie-panel dark-theme true
    dconf write /org/gnome/terminal/legacy/theme-variant "'dark'"
    gsettings set org.gnome.desktop.background picture-uri "file://${WALLPAPER}"
    spicetify -q config current_theme Adapta-Nokto
    spicetify -q update
  elif [ $1 == "light" ]; then
    THEME='WhiteSur-light-solid-purple'
    ICON='Numix-Square-Light'
    WALLPAPER='/home/william/Pictures/Wallpapers/light.png'
    gsettings set com.solus-project.budgie-panel dark-theme false
    dconf write /org/gnome/terminal/legacy/theme-variant "'light'"
    gsettings set org.gnome.desktop.background picture-uri "file://${WALLPAPER}"
    spicetify -q config current_theme Midnight-Light
    spicetify -q update
  fi

  # while read -r __window; do
  #   if [ "${__window}" != "" ]; then
  #     xdotool windowactivate --sync "${__window}" sleep 0.1 key "ctrl+shift+r" &
  #   fi
  # done <<<"$(xdotool search --class Spotify | tail -n 1)"

  gsettings set org.gnome.desktop.interface gtk-theme "${THEME}"
  gsettings set org.gnome.desktop.interface icon-theme "${ICON}"

  gdbus introspect -e -x -d org.gnome.Terminal -o /org/gnome/Terminal/window |
    xmllint --xpath '/node/node/@name' - |
    sed -e 's/name="\([0-9]*\)"/\1\n/g' -e 's/ //g' |
    while read -r __node; do
      gdbus call -e -d org.gnome.Terminal -o /org/gnome/Terminal/window/${__node} -m org.gtk.Actions.SetState profile "<'${UUID}'>" "{}" >/dev/null
    done

elif [ "${1}" == 'get' ]; then
  __state="$(gsettings get com.solus-project.budgie-panel dark-theme)"
  if [ "${__state}" == 'false' ]; then
    echo -n 'OFF'
  else
    echo -n 'ON'
  fi

fi

exit
