#!/bin/bash

### DEPENDENCIES_BEGIN
# powerprofilesctl
### DEPENDENCIES_END

### CONFIG_TEMPLATE_BEGIN
# {
#     "select": [
#         {
#             "name": "___HOSTNAME___ Power Mode",
#             "options": [
#                 "performance",
#                 "balanced",
#                 "power-saver"
#             ],
#             "state": [
#                 "powerprofilesctl",
#                 "get"
#             ],
#             "command": [
#                 "powerprofilesctl",
#                 "set"
#             ],
#             "mqtt": {
#                 "update_interval": 1
#             }
#         }
#     ]
# }
### CONFIG_TEMPLATE_END