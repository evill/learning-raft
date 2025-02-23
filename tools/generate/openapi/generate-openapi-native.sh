#!/bin/bash
set -e

# Set variables
BIN_DIR="bin"
HTTP_NATIVE_GENERATED_TRANSPORT_PATH=internal/transports/http/inbound/native

$BIN_DIR/openapi-generator generate \
    -i api/openapi/api.yml \
    -g go-server \
    -o $HTTP_NATIVE_GENERATED_TRANSPORT_PATH \
    --additional-properties=outputAsLibrary=true,sourceFolder=,serverPort=8081,addResponseHeaders=false \
    --skip-validate-spec \
    --skip-operation-example \
    --enable-post-process-file \
