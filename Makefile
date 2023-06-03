postgres-run:
	@docker run --name postgreslaw -p 5432:5432 -e POSTGRES_USER=lawyerinyou -e POSTGRES_PASSWORD=asdqwe123 -e POSTGRES_DB=lawyerinyou2023 -d postgres

postgres:
	@docker start postgreslaw

postgres-stop:
	@docker stop postgreslaw

redis-run:
	@docker run -d -p 6379:6379 --name redislaw redis

redis:
	@docker start redislaw

redis-stop:
	@docker stop redislaw

setup: postgres-run redis-run postgres-stop redis-stop

stop: postgres-stop redis-stop

container: postgres redis

main: 
	@go run ./cmd/main.go

run: postgres redis main

test:
	@go test -v -cover ./...

run-test: postgres redis test