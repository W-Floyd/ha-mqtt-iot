#!/bin/bash

case "${1}" in
"get-cpu-temp")
  sensors coretemp-isa-0000 -j -A | jq '."coretemp-isa-0000"."Package id 0"."temp1_input"'
  ;;
esac
