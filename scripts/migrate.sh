#!/bin/bash

# PostgreSQL connection details
DB_URL="postgres://username:password@localhost:5432/database_name?sslmode=disable"

# Apply all pending migrations
migrate -path ./scripts/migrations -database $DB_URL up

# Check for any errors
if [ $? -eq 0 ]; then
  echo "Migrations applied successfully!"
else
  echo "Migration failed."
  exit 1
fi
