#!/bin/bash

case "${1}" in
    "go-get")
        go get -u
        go mod tidy
        git add go.mod go.sum
        git commit -m "chore: Go update."
        ;;
    *)
        echo "Unknown command: ${1}"
        ;;
esac
    

exit