#!/bin/sh

uname=$(uname)
env=".env"

export $(grep -v '^#' "$env" | xargs)

docker-compose up -d

./gradlew bootRun
