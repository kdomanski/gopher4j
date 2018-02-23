#!/usr/bin/env bash

set -eu
set -o pipefail

TARGET_PORT=$1
LISTEN_PORT="${2:-2222}"

socat -v "tcp-l:${LISTEN_PORT},fork,reuseaddr" "tcp:127.0.0.1:${TARGET_PORT}"
