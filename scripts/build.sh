#!/bin/bash

VERSION=${1:?}
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )/.."

env GOOS=linux GOARCH=amd64 go build -o ${SCRIPT_DIR}/builds/harbor-linux-v${VERSION} cmd/harbor/main.go

env GOOS=darwin GOARCH=amd64 go build -o ${SCRIPT_DIR}/builds/harbor-darwin-v${VERSION} cmd/harbor/main.go

env GOOS=windows GOARCH=amd64 go build -o ${SCRIPT_DIR}/builds/harbor-windows-v${VERSION}.exe cmd/harbor/main.go

