#!/bin/bash

### DEPENDENCIES_BEGIN
# mplayer base64 convert
### DEPENDENCIES_END

### CONFIG_TEMPLATE_BEGIN
# {
#    "camera": [
#        {
#            "name": "___HOSTNAME___ Camera",
#            "object_id": "___HOSTNAME___-camera",
#            "unique_id": "___HOSTNAME___-camera",
#            "state": [
#                "___SCRIPTS_DIR______SCRIPT_NAME___"
#            ],
#            "encoding": "b64",
#            "mqtt": {
#                "update_interval": 10
#            }
#        }
#    ]
# }
# ### CONFIG_TEMPLATE_END



__tmpdir="$(mktemp -d)"
__image_path="${__tmpdir}/00000030.jpg"
__image_path_2="${__tmpdir}/00000001_rescaled.jpg"

cd "${__tmpdir}" || exit 1

mplayer -vo jpeg -frames 30 tv:// &>/dev/null

# convert "${__image_path}" -scale '50%' "${__image_path_2}"

base64 <"${__image_path}"

rm -r "${__tmpdir}"

exit
