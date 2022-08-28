
https://dbeaver.io/

[migrate package & install instructions](https://github.com/golang-migrate/migrate)


how to install MAKE
 *`sudo apt install make`

how to create the schemas for the initial [or subsecuential] migration
`migrate create -ext <extension> -dir <directory> -seq <migration_name>`

How to migrate
 * up all versions:     `make migrateup`
 * down all versions:   `make migratedown`
 * up 1 version:        `make migrateup1`
 * down 1 version:      `make migratedown1`

[how to install sqlc](https://docs.sqlc.dev/en/latest/overview/install.html#):
* `sudo snap install sqlc` or `go install github.com/kyleconroy/sqlc/cmd/sqlc@latest`

Manual Start of proyect:
    * `make network`
    * `make postgres`
    * `make createdb`
    * `make migrateup`
    * `make server`

(Gin Web Framework)[https://github.com/gin-gonic/gin]: 
 * `go get -u github.com/gin-gonic/gin`    
t
(Handle environment variables with Viper)[https://github.com/spf13/viper]:
 * `go get github.com/spf13/viper`

(Mock-db with gomock)[https://github.com/golang/mock]:
 * `go install github.com/golang/mock/mockgen@v1.6.0`

Install vscode-dbml for (.DBML sintax highlights)[https://marketplace.visualstudio.com/items?itemName=matt-meyers.vscode-dbml]
    * Ctrl+p
    * `ext install matt-meyers.vscode-dbml`

To build the database documentation page: (dbdocs)[dbdocs.io/docs]
  * `sudo npm install -g dbdocs`
  * `dbdocs login` and follow the instructions
  * `make db_docs`

Make will call: `dbdocs build postgres/doc/db.dbml`

You could secure the docs with a password with:
  * `dbdocs password --set <your-password> --project Simple_bank`

To generate the POSTGRES SQL code: (dbml)[dbml.org/cli/]
  * `sudo npm install -g @dbml/cli`
  * `make db_schema`

Make will call: `dbml2sql --postgres -o postgres/doc/schema.sql postgres/doc/db.dbml`

To modify the DB Schema:
    * update the code of postgres/doc/db.dbml
    * `make db_docs`
    * `make db_schema`