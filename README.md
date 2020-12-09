This is a generic contact service with a RESTful API.

# Features:
* You can get list of contacts
* You can add a contact
* You can edit a contact
* You can delete a contact

# Requirements
* [docker](https://www.docker.com/) - a container platform
* [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) - a tool for database migrations

# Quickstart
```
# get deps
make init
# run services
make compose
# apply migrations˚
make migrate
# run the app
make run
```
# Services hosts
|Address|App|Description|
|---|---|---|
|localhost:8080|Contact app|The test application|
|localhost:5432|Postgres|Database|

# GNU Make commands
|Command|Description|
|---|---|
|init|create a new .env file from template and download deps|
|compose|run services in docker|
|migration|apply all up migrations|
|migratie name=my_cool_migration|create the new migration with the name my_cool_migration|
|migrate-down|apply all down migrations|
|remigrate|apply all down and up migrations|
|run|run the app|
