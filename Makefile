
run:
	@go run cmd/main.go

migrate_import:
	go get github.com/golang-migrate/migrate/v4/database/postgres

migrate_create:
	migrate create -ext sql -dir migrations -seq fitness

migrate_up:
	migrate -database postgres://postgres:3115@localhost:5432/fitness_tracking_db?sslmode=disable -path ./migrations up

migrate_down:
	migrate -database postgres://postgres:3115@localhost:5432/fitness_tracking_db?sslmode=disable -path ./migrations down

migrate_force:
	migrate -database postgres://postgres:3115@localhost:5432/fitness_tracking_db?sslmode=disable -path ./migrations force

sqlc-generate:
	sqlc vet ; sqlc generate