#!/bin/bash

if [ $# -ne 2 ]; then
    echo "Usage: $0 <BASE_URL> <AUTH_TOKEN>"
    exit 1
fi

export BASE_URL="$1"
export AUTH_TOKEN="$2"

go test
