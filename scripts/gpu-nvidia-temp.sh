#!/bin/bash

nvidia-smi -q -x | xq -r '."nvidia_smi_log"."gpu"."temperature"."gpu_temp"' | sed -e 's/ .*//'

exit
