#!/bin/sh

uname=$(uname)
env=".env"

export $(grep -v '^#' "$env" | xargs)

./gradlew bootRun
