#!/bin/bash

__tmp_file="$(mktemp --suffix='.jpg')"
__tmp_file2="$(mktemp --suffix='.jpg')"
gnome-screenshot -f "${__tmp_file}"
convert "${__tmp_file}" -scale '10%' "${__tmp_file2}"
cat "${__tmp_file2}" | base64
rm "${__tmp_file}" "${__tmp_file2}"

exit
