#!/bin/zsh

echo "Cleaning docker system..."
docker compose down -v
docker system prune -af
