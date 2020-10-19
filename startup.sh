#!/bin/bash
set -e
set -x

echo "Starting API-BACKEND..."
api-backend \
    --log-level=$API_LOG_LEVEL \
    --server-hostname=$API_SERVER_HOSTNAME \
    --server-port=$API_SERVER_PORT \
    --database-hostname=$API_DATABASE_HOSTNAME \
    --database-port=$API_DATABASE_PORT \
    --database-username=$API_DATABASE_USERNAME \
    --database-password=$API_DATABASE_PASSWORD \
    --database-name=$API_DATABASE_NAME \