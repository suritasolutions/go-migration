.PHONY: up down bash test test-bench

up:
	docker compose up -d

down:
	docker compose down -v

bash:
	docker compose exec cli bash

test:
	docker compose exec cli bash -c "go test -v ./..."

test-bench:
	docker compose exec cli bash -c "go test -bench=. -benchmem ./..."