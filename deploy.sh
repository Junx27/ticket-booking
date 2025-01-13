#!/bin/bash
cd /home/junxgurit/app/ticket-booking
git pull origin main
if [[ -z $(git diff --name-only HEAD@{1}..HEAD | grep -E "Dockerfile|\.go|\.js|\.py|\.env") ]]; then
  echo "Tidak ada perubahan yang relevan untuk rebuild Docker. Proses selesai."
else
  echo "Ada perubahan, membangun ulang Docker..."
  docker-compose down --volumes --remove-orphans
  docker-compose up --build -d
fi
