# golang-api

A simple API to learn [go](https://go.dev/) and tools.

### Tools:

- [gRPC](https://grpc.io/): A high performance, open source universal RPC framework.
- [proto](https://protobuf.dev/): Language-neutral, platform-neutral extensible mechanisms for serializing structured data.
- [sqlc](https://sqlc.dev/): Generates fully type-safe idiomatic Go code from SQL.
- [atlas](https://atlasgo.io/): Language-independent tool for managing and migrating database schemas.

## Project structure

`db/schema.hcl`: Atlas schema file. Used to generate migration files.

`db/migrations/`: Atlas migration files output directory. Sqlc generates models from these files.

`db/*.sql`: Sql code that is compiled by sqlc.

`sqlc.yaml`: Sqlc configuration file.

`proto/`: Proto configuration files and code.

`gen/`: Go code generated by sqlc and proto.

`server/`: API services implementation.

## Makefile rules

`all`: Calls `build`

`build`: Builds binaries from `cmd/*` in `bin/`.

`gen`: Calls `proto` and `sqlc`.

`proto`: Generates go code in `gen/` from proto code in `proto/`.

`sqlc`: Generates go code in `gen/` from sql `db/*.sql` files.

`docker`: Starts all containers from `docker-compose.yaml`.

`dclean`: Stops and cleans all containers, networks and volumes.

`mdiff`: Use atlas to generate migration files.

`mapply`: Push migrations to the database.

Environment variables:

These variables are used by `atlas` to apply migrations.

`MYSQL_PASSWORD`: Required. Password for the mysql database.

`MYSQL_USER`: User of the mysql database. Default: `root`.

`MYSQL_ADDR`: Address of the mysql database. Default: `mysql:3306`.

`MYSQL_DATABASE`: Database name. Default: `mysql`.
