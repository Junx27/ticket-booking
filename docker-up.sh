read -p "Are you sure you want to prune all unused Docker resources? This will remove stopped containers, networks, and dangling images. (y/n): " confirm
if [[ "$confirm" == "y" || "$confirm" == "Y" ]]; then
    echo "Cleaning up unused Docker resources..."
    docker system prune
    echo "Docker system prune completed."
else
    echo "Skipping Docker system prune."
fi
echo "Starting Docker Compose with build..."
docker-compose up --build
