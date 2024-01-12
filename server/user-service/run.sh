#!/bin/sh

unamestr=$(uname)
envfile=".env"

if [ "$unamestr" = 'Linux' ]; then
	export $(grep -v '^#' "$envfile" | xargs -d '\n')
fi

./gradlew bootRun
