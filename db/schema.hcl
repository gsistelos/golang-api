schema "golang-api" {}

table "users" {
    schema = schema.golang-api
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

table "posts" {
    schema = schema.golang-api
    column "id" {
        type = varchar(36)
        null = false
        default = "UUID()"
    }
    column "content" {
        type = varchar(255)
        null = false
    }
    column "visibility" {
        type = varchar(255)
        null = false
        default = "public"
    }
    column "user_id" {
        type = varchar(36)
        null = false
    }
    primary_key {
        columns = [column.id]
    }
    foreign_key "user_id" {
        columns = [column.user_id]
        ref_columns = [table.users.column.id]
        on_delete = "CASCADE"
    }
}
