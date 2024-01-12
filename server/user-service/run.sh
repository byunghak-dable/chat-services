#!/bin/sh

uname=$(uname)
env=".env"

if [ "$uname" = 'Linux' ]; then
	export $(grep -v '^#' "$env" | xargs -d '\n')
fi

./gradlew bootRun
