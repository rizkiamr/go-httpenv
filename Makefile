build:
	go build -o main main.go

docker-build:
	docker build -t go-httpenv:latest .

docker-compose-up:
	docker compose up -d

docker-compose-down:
	docker compose down -v

.PHONY: docker-build build docker-compose-up docker-compose-down