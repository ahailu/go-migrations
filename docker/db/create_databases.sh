#!/usr/bin/env bash

set -e

export PGUSER=postgres
psql <<- EOSQL
    CREATE USER scheduler;
    CREATE DATABASE scheduler;
    GRANT ALL PRIVILEGES ON DATABASE scheduler TO scheduler;
EOSQL
