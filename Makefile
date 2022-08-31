DB_URL=postgresql://root:secret@localhost:5432/goBank?sslmode=disable

network:
	docker network create bank-network

postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

clean:
	docker-compose down
	docker system prune --all -f

createdb:
	docker exec -it postgres createdb --username=root --owner=root goBank

dropdb:
	docker exec -it postgres dropdb goBank

migrateup:
	migrate -path postgres/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path postgres/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path postgres/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path postgres/migration -database "$(DB_URL)" -verbose down 1

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination postgres/mock/store.go github.com/BMogetta/goBank_go-course/postgres/sqlc Store

proto:
	rm -f proto_go/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=proto_go --go_opt=paths=source_relative \
	--go-grpc_out=proto_go --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=proto_go --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=goBank \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

.PHONY: network postgres clean createdb dropdb migrateup migratedown migrateup1 migratedown1 db_docs db_schema sqlc test server mock proto evans