#!/bin/bash

case "${1}" in
    "go-get")
        go get -u
        go mod tidy
        git add go.mod go.sum
        git commit -m "chore: Go update."
        ;;
    "generate")
        go run ./helpers --pull
        git add devices/ config/
        git commit -m "chore: Regenerate from upstream documentation."
        ;;
    *)
        echo "Unknown command: ${1}"
        ;;
esac
    

exit