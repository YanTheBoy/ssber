up:
	docker compose -f docker-compose.yml --project-name common up -d
	docker compose -f offering-service/internal/build/docker-compose.yml --project-name offering-service up -d --build

stop:
	docker compose -f docker-compose.yml --project-name common stop
	docker compose -f offering-service/internal/build/docker-compose.yml --project-name offering-service stop
