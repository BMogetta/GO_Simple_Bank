
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
