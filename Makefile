create-migrate:
	migrate create -ext=sql -dir=migrations -seq init
server-up-dev:
	docker-compose  -f docker-compose.dev.yaml up --build
server-up-prod:
	docker-compose  -f docker-compose.prod.yaml up --build
server-down:
	docker-compose down
install-migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
generate:
	go run github.com/99designs/gqlgen generate
