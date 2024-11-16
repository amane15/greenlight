## help: print this help message
help:
	@echo 'Usage: '
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

## run/api: run the cmd/api application
run/api:
	go run ./cmd/api

## db/migrations/new name=$1: create a new database migration
db/migrations/new:
	@echo 'Create migration files for ${name}'
	migrate create -seq -ext=".sql" -dir="./migrations" ${name}

## db/migrations/up: apply all up database migrations
db/migrations/up: confirm
	@echo "Running up migrations"
	migrate -path ./migrations -database postgres://greenlight:pass@localhost:5433/greenlight?sslmode=disable up

audit:
	@echo "Formatting code..."
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck 
	@echo 'Running tests...'
	go test -race -vet=off ./...
	@echo 'Vendoring dependancies'
	go mod vendor

build/api:
	@echo 'Building cmd/api'
	go build -ldflags='-s' -o=./bin/api ./cmd/api