#!/bin/bash

set -euo pipefail
#set -x

version_gte() {
  [ "$1" = "$(echo -e "$1\n$2" | sort -V | tail -n1)" ]
}

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"


hash protoc 2>/dev/null || { echo >&2 -e 'error: protoc binary not found\nDownload it from https://github.com/google/protobuf/releases and put it in your path.\nMake sure you get the one starting with `protoc`, not `protobuf`.'; exit 1; }


PROTOC="$(which protoc)"
VERSION="$($PROTOC --version | cut -d' ' -f2)"
MIN_VERSION="3.0"

version_gte "$VERSION" "$MIN_VERSION" || { echo >&2 "error: protoc version must be >= $MIN_VERSION (your $PROTOC is $VERSION)"; exit 1; }


hash protoc-gen-go 2>/dev/null || go get -u github.com/golang/protobuf/protoc-gen-go
hash protoc-gen-go 2>/dev/null || { echo >&2 'error: Make sure $GOPATH/bin is in your $PATH'; exit 1; }


find $DIR/go $DIR/python $DIR/js -type f -delete


protoc --proto_path="$DIR/proto" --python_out="$DIR/python" --go_out="$DIR/go" --js_out="import_style=commonjs,binary:$DIR/js" $DIR/proto/*.proto
