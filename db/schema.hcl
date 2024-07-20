schema "grpc-api" {}

table "users" {
    schema = schema.grpc-api
    column "id" {
        type = varchar(36)
        null = false
        default = "UUID()"
    }
    column "username" {
        type = varchar(255)
        null = false
    }
    column "email" {
        type = varchar(255)
        null = false
    }
    column "password" {
        type = varchar(255)
        null = false
    }
    primary_key {
        columns = [column.id]
    }
    index "email" {
        unique = true
        columns = [column.email]
    }
}
