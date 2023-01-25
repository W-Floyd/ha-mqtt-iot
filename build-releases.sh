#!/bin/bash

__build_combos="aix, ppc64
android, 386
android, amd64
android, arm
android, arm64
darwin, amd64
darwin, arm64
dragonfly, amd64
freebsd, 386
freebsd, amd64
freebsd, arm
illumos, amd64
ios, arm64
js, wasm
linux, 386
linux, amd64
linux, arm
linux, arm64
linux, loong64
linux, mips
linux, mipsle
linux, mips64
linux, mips64le
linux, ppc64
linux, ppc64le
linux, riscv64
linux, s390x
netbsd, 386
netbsd, amd64
netbsd, arm
openbsd, 386
openbsd, amd64
openbsd, arm
openbsd, arm64
plan9, 386
plan9, amd64
plan9, arm
solaris, amd64
windows, 386
windows, amd64
windows, arm
windows, arm64"

set -x

while read -r __line; do
    __goos="${__line/,*/}"
    __goarch="${__line/*, /}"
    GOOS="${__goos}" GOARCH="${__goarch}" go build -o "ha-mqtt-iot-${__goos}-${__goarch}"
done <<<"${__build_combos}"

exit
