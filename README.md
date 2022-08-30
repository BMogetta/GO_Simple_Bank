
https://www.youtube.com/c/TECHSCHOOLGURU

## Graphical view of the DB with [DBeaver](https://dbeaver.io/)

# [migrate package & install instructions](https://github.com/golang-migrate/migrate)


## Install MAKE
 * `sudo apt install make`

### how to create the schemas for the initial migration
`migrate create -ext <extension> -dir <directory> -seq <migration_name>`

How to migrate
 * up all versions:     `make migrateup`
 * down all versions:   `make migratedown`
 * up 1 version:        `make migrateup1`
 * down 1 version:      `make migratedown1`

## Install [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html#):
* `sudo snap install sqlc`  
* `go install github.com/kyleconroy/sqlc/cmd/sqlc@latest`

### Manual Start of proyect:
  * `make network`
  * `make postgres`
  * `make createdb`
  * `make migrateup`
  * `make server`

## [Gin Web Framework:](https://github.com/gin-gonic/gin) 
 * `go get -u github.com/gin-gonic/gin`    


## Handle environment variables with [Viper](https://github.com/spf13/viper):
 * `go get github.com/spf13/viper`

## Mock-db with [gomock](https://github.com/golang/mock):
 * `go install github.com/golang/mock/mockgen@v1.6.0`

### Install [vscode-dbml](https://marketplace.visualstudio.com/items?itemName=matt-meyers.vscode-dbml) for .DBML sintax highlights
  * Ctrl+p
  * `ext install matt-meyers.vscode-dbml`

## Build the database documentation page with [dbdocs](dbdocs.io/docs):
  * `sudo npm install -g dbdocs`
  * `dbdocs login` and follow the instructions
  * `make db_docs`

Make will call -> `dbdocs build doc/db.dbml`

You could secure the docs with a password with:
  * `dbdocs password --set <your-password> --project Simple_bank`

##  Generate the POSTGRES SQL code with [dbml](dbml.org/cli/):
  * `sudo npm install -g @dbml/cli`
  * `make db_schema`

Make will call -> `dbml2sql --postgres -o doc/schema.sql doc/db.dbml`

## To modify the DB Schema:
  * update the code of doc/db.dbml
  * `make db_docs`
  * `make db_schema`

## [gRPC](https://grpc.io/docs/)

### [proto compiler](https://grpc.io/docs/protoc-installation/)
  * `sudo apt install -y protobuf-compiler`
  * `protoc --version`

### [Go plugins](https://grpc.io/docs/languages/go/quickstart/)
  * go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
  * go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

## gRPC client for testing: [evans](https://github.com/ktr0731/evans)

  * `curl -Lo evans.tgz https://github.com/ktr0731/evans/releases/download/v0.10.9/evans_linux_amd64.tar.gz`
  * `tar -xf evans.tgz`
  * `sudo install evans /usr/local/bin/`
  * run with `evans`

## [gRPC-Gateway](https://grpc-ecosystem.github.io/grpc-gateway/)
    `go install \`
    `github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \`
    `github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \`
    `google.golang.org/protobuf/cmd/protoc-gen-go \`
    `google.golang.org/grpc/cmd/protoc-gen-go-grpc`