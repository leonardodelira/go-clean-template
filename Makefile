include .env
export

.PHOYNY: run
run: migration-up
	GIN_MODE=debug go run main.go

.PHOYNY: unit test
unit-test:
	go test ./...

.PHOYNY: migration-create
migration-create:
	migrate create -ext sql -dir migrations $(name)

.PHOYNY: migration-up
migration-up:
	migrate -path migrations -database 'postgres://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DATABASE)?sslmode=disable' up
