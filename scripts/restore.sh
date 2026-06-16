#!/bin/zsh

FILE=$1

if [ -z "$FILE" ]; then
    echo "Usage: ./scripts/restore.sh backups/filename.sql"
    exit 1
fi

echo "Restoring from $FILE..."

cat $FILE | docker exec -i task-db psql -U postgres taskdb

if [ $? -eq 0 ]; then
    echo "Restore completed"
else
    echo "Restore failed"
fi
