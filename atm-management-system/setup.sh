#!/bin/bash

set -e
# check if postgres is installed
if ! which psql &> /dev/null; then
    echo "PostgreSQL is not installed or not in your PATH."
    exit 1
fi

if ! error_message=$(PGPASSWORD=1234 psql -d atmdb -U natheer -c "SELECT 1;" 2>&1); then
    echo "Database 'atmdb' does not exist or user 'natheer' does not have access to it."
    # Print the error message
    echo "Error message: $error_message"
    echo -n "Do you want to create the user and database? (y/n): "
    read -r response
    if [ "$response" = "y" ]; then
        # Create the user and database
        echo "Creating user 'natheer'...."
        psql -U postgres -c "CREATE USER natheer WITH PASSWORD '1234';"
        psql -U postgres -c "CREATE DATABASE atmdb"
        psql -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE atmdb TO natheer;"
        echo "adding tables to 'atmdb'...."
        psql -d atmdb -U postgres -f db/create_postgre.sql
        echo "everything created fine"
        exit 0
    fi
    exit 0
fi
echo "nothing is wrong with the database"