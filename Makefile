.PHONY: up down bash

up:
	docker-compose up -d

down:
	docker-compose down -v

bash:
	docker-compose exec cli bash