#!/bin/bash

__tmpdir="$(mktemp -d)"
__image_path="${__tmpdir}/00000030.jpg"
__image_path_2="${__tmpdir}/00000001_rescaled.jpg"

cd "${__tmpdir}" || exit 1

mplayer -vo jpeg -frames 30 tv:// &>/dev/null

# convert "${__image_path}" -scale '50%' "${__image_path_2}"

base64 <"${__image_path}"

rm -r "${__tmpdir}"

exit
