#!/bin/bash
set -e
set -u

function create_database() {
    local database=$1
    echo "  Creating user and database '$database' and setting ownership of it to $POSTGRES_USER"
    psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
        CREATE USER $database;
        CREATE DATABASE $database;
        GRANT ALL PRIVILEGES ON DATABASE $database TO $POSTGRES_USER;
EOSQL
}

if [ -n "$POSTGRES_ADDITIONAL_DATABASES" ]; then
    echo "Additional database creation requested: $POSTGRES_ADDITIONAL_DATABASES"
    for db in $(echo $POSTGRES_ADDITIONAL_DATABASES | tr ',' ' '); do
        create_database $db
    done
    echo "Additional databases successfully created"
fi