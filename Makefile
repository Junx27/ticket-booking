start:
	@bash docker-up.sh

stop:
	@docker compose down -v
