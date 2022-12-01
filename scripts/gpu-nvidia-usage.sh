#!/bin/bash

nvidia-smi -q -x | xq -r '."nvidia_smi_log"."gpu"."utilization"."gpu_util"' | sed -e 's/ .*//'

exit
