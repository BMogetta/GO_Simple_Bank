migrate package & install instructions
https://github.com/golang-migrate/migrate

how to install MAKE in windows
choco install make

how to create the schemas for the initial migration
migrate create -ext sql -dir /postgres/migrations -seq init_schema


how to migrate
 * make migrateup
 * make migratedown

how to install sqlc:
https://docs.sqlc.dev/en/latest/overview/install.html#

go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

if "cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%"
https://jmeubank.github.io/tdm-gcc/download/