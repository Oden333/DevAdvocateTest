#!/bin/sh

set -e 

echo "run db migration"
/app/migrate -path /app/migration -database "postgresql://root:qwerty@postgres:5432/users?sslmode=disable" -verbose up

echo "start the app"
exec "$@"