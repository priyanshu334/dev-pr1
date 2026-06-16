#!/bin/zsh

echo "Cleaning old backups..."

ls -tp backups/*.sql | grep -v '/$' | tail -n +8 | xargs -I {} rm -- {}
