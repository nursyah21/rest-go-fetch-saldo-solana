#!/bin/bash

TARGET="src"
COMMAND="go build -o ./tmp/main.exe ./src && ./main.exe"
POLLING_INTERVAL=1

PID=""

restart_command() {
    if [ ! -z "$PID" ]; then
        kill $PID 2>/dev/null
    fi
    eval "$COMMAND" &
    PID=$!
}

get_modified_time() {
    if [ -f "$1" ]; then
        stat -c %Y "$1"
    elif [ -d "$1" ]; then
        find "$1" -type f -exec stat -c %Y {} + | sort -n | tail -1
    else
        echo 0
    fi
}

last_mod=$(get_modified_time "$TARGET")
restart_command

while true; do
    sleep $POLLING_INTERVAL
    current_mod=$(get_modified_time "$TARGET")
    if [ "$current_mod" != "$last_mod" ]; then
        restart_command
        last_mod=$current_mod
    fi
done