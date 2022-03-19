#!/usr/bin/env bash
set -x
set -e pipefail

MONGO_USERNAME=${MONGO_USER:=mongo}
# Check if a custom password has been set, otherwise default to 'password'
MONGO_PASSWORD="${MONGO_PASSWORD:=password}"
# Check if a custom port has been set
DB_PORT="${MONGO_PORT:=8000}"

# Launch mongo using Docker
docker run \
-e MONGO_INITDB_ROOT_PASSWORD=${MONGO_PASSWORD} \
-e MONGO_INITDB_ROOT_USERNAME=${MONGO_USERNAME} \
-p "${DB_PORT}":27017 \
--name mongodb \
mongo