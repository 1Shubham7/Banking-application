#!/bin/sh

# exits the script if any command fails
set -e

echo "run db migration"
# source /app/app.env

# run database migrations.
/app/migrate -path /app/migration --database "$DB_SOURCE" --verbose up

echo "start the app"

# execute the command passed to the script
exec "$@"