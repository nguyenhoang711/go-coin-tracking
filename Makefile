run:
	go run cmd/main.go 
migrateup:
	migrate -path database/migration -database "postgresql://root:123@abc@localhost:5433/trackingcoin?sslmode=disable" -verbose up
migratedown:
	migrate -path database/migration -database "postgresql://root:mysecurepassword@localhost:5433/trackingcoin?sslmode=disable" -verbose down
migrateforce:
	migrate -path database/migration -database "postgresql://root:mysecurepassword@localhost:5433/trackingcoin?sslmode=disable" force 1
sqlc:
	sqlc generate --file=database/sqlc.yaml
compose:
	docker compose up -d
.PHONY: run migrateup migratedown sqlc compose