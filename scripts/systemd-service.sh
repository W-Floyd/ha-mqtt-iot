#!/bin/bash

systemctl status ${1} > /dev/null
RET=$(echo $?)

if [ "${RET}" == '0' ]; then
    echo ON
else
    echo OFF
fi
