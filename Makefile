postgres-run:
	@docker run --name postgreslaw -p 5432:5432 -e POSTGRES_USER=lawyerinyou -e POSTGRES_PASSWORD=asdqwe123 -e POSTGRES_DB=lawyerinyou2023 -d postgres

postgres:
	@docker start postgreslaw

postgres-stop:
	@docker stop postgreslaw

setup: postgres-run postgres-stop

stop: postgres-stop

container: postgres

main: 
	@go run ./cmd/main.go

run: postgres main