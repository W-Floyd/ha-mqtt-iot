#!/bin/bash

case "${1}" in
    "go")
        go get -u
        go mod tidy
        git add go.mod go.sum
        git commit -m "chore: Go update."
        ;;
    "gen")
        go run ./helpers --pull
        git add devices/ config/
        git commit -m "chore: Regenerate from upstream documentation."
        ;;
    *)
        echo "Unknown command: ${1}"
        ;;
esac
    

exit