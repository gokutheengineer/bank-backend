#!/bin/sh

set -e 

echo "migrate down db"
echo Y | /app/migrate -path /app/migration -database "$DB_SOURCE" -verbose down
echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up
echo $DB_SOURCE
echo "db migration done"

echo "start app"
exec "$@"