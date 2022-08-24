postgres:
	docker compose up --build postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb: docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path postgres/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path postgres/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test