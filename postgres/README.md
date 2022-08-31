# This project run a [PostGresSQL docker](https://hub.docker.com/_/postgres/) container as database

[PostgreSQL](https://www.postgresql.org/)  is an object-relational database management system (ORDBMS) with an emphasis on extensibility and standards-compliance. As a database server, its primary function is to store data, securely and supporting best practices, and retrieve it later, as requested by other software applications, be it those on the same computer or those running on another computer across a network (including the Internet). It can handle workloads ranging from small single-machine applications to large Internet-facing applications with many concurrent users. Recent versions also provide replication of the database itself for security and scalability.

---

## **Graphical view of the DB with [DBeaver](https://dbeaver.io/)**

DBeaver is free and open source universal database tool for developers and database administrators.

---

## **[Migrate](https://github.com/golang-migrate/migrate) package & install instructions**


Database migrations written in Go. Use as CLI or import as library.

Migrate reads migrations from sources and applies them in correct order to a database.
Drivers are "dumb", migrate glues everything together and makes sure the logic is bulletproof. (Keeps the drivers lightweight, too.)
Database drivers don't assume things or try to correct user input. When in doubt, fail.

### **how to create the schemas for the initial migration**
`migrate create -ext <extension> -dir <directory> -seq <migration_name>`

**How to migrate**
 * up all versions:     `make migrateup`
 * down all versions:   `make migratedown`
 * up 1 version:        `make migrateup1`
 * down 1 version:      `make migratedown1`

---
## **Install [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html#):**

sqlc generates fully type-safe idiomatic Go code from SQL. Here's how it works:

1. You write SQL queries
2. You run sqlc to generate Go code that presents type-safe interfaces to those queries
3. You write application code that calls the methods sqlc generated

`sudo snap install sqlc`  

`sqlc`

---
## **Build the database documentation page with [dbdocs](dbdocs.io/docs):**
`sudo npm install -g dbdocs`

`dbdocs login` and follow the instructions

`make db_docs`

You could secure the docs with a password with:

`dbdocs password --set <your-password> --project goBank_go-course`

---
##  **Generate the POSTGRES SQL code with [dbml](https://www.dbml.org/cli/):**

DBML (Database Markup Language) is an open-source DSL language designed to define and document database schemas and structures. It is designed to be simple, consistent and highly-readable.

DBML comes with a built-in CLI which can be used to convert between different formats from the command line

`sudo npm install -g @dbml/cli`

`make db_schema`

---
## **To modify the DB Schema:**
  1. update the code of doc/db.dbml
  2. `make db_docs`
  3. `make db_schema`

---

## To review the DB docs

Run the project (with docker-compose or manually), open your brownser and paste:

`http://localhost:8080/swagger/`

