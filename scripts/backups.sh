#!/bin/zsh

TIMESTAMP=$(date +"%Y-%m-%d_%H-%M-%S")
BACKUP_FILE="backups/taskdb_$TIMESTAMP.sql"

echo "Creating backup..."

docker exec task-db pg_dump -U postgres taskdb > $BACKUP_FILE

if [ $? -eq 0 ]; then
    echo "Backup created: $BACKUP_FILE"
else
    echo "Backup failed"
fi
