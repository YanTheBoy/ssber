PROJECTNAME=offering-service
API_PATH=api/sbercloud/${PROJECTNAME}

generate:
	protoc --go_out=./pkg/ \
	--go-grpc_out=./pkg/ \
	./api/proto/*.proto

up:
	docker-compose up -d
	go run cmd/main.go


run-migrate-development:
	sql-migrate up -env="development"

stop:
	docker compose stop