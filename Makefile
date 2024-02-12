build:
	go build -o main main.go

docker-build:
	docker build -t go-httpenv:latest .

docker-compose:
	docker compose up -d

.PHONY: docker-build build