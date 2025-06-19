#!/bin/bash

TARGET="src"
BUILD_CMD='CGO_ENABLED=0 go build -ldflags "-s -w" -o ./tmp/main -o ./tmp/main ./src'
RUN_CMD="./tmp/main"
POLLING_INTERVAL=1
PID=""

mkdir -p ./tmp

restart_command() {
  if [ -n "$PID" ]; then
    kill $PID 2>/dev/null
  fi

  echo "please wait..."
  eval "$BUILD_CMD"

  bash -c "$RUN_CMD" &
  PID=$!
}

create_hash() {
  find "$TARGET" -type f -name "*.go" -exec md5sum {} + 2>/dev/null | awk '{print $1}' | sort | md5sum | awk '{print $1}'
}

last_hash=$(create_hash)
restart_command

while true; do
  sleep $POLLING_INTERVAL
  current_hash=$(create_hash)
  if [ "$current_hash" != "$last_hash" ]; then
    restart_command
    last_hash=$current_hash
  fi
done