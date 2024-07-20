BUF_GENERATE = buf generate --template proto/buf.gen.yaml

SQLC_GENERATE = sqlc generate -f sqlc.yaml

DOCKER_COMPOSE = docker compose -f docker-compose.yaml

# Set variables if not set
ifeq ($(MYSQL_USER),)
	MYSQL_USER = root
endif

ifeq ($(MYSQL_DATABASE), mysql)
	MYSQL_DATABASE = mysql
endif

MYSQL_URL = "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@localhost:3306/$(MYSQL_DATABASE)"


all: proto sqlc run

.PHONY: proto
proto:
	$(BUF_GENERATE) proto

.PHONY: sqlc
sqlc:
	$(SQLC_GENERATE)

.PHONY: run
run:
	cd cmd/api && go run .

.PHONY: docker
docker:
	$(DOCKER_COMPOSE) up -d

.PHONY: dclean
dclean:
	-$(DOCKER_COMPOSE) down
	-docker rm $$(docker ps -aq)
	-docker network rm -f $$(docker network ls -q)
	-docker volume rm -f $$(docker volume ls -q)

.PHONY: migrations
migrations:
	atlas migrate diff \
		--dir "file://db/migrations" \
		--to "file://db/schema.hcl" \
		--dev-url "docker://mysql/8/example"

.PHONY: migrate
migrate:
	atlas migrate apply \
		--dir "file://db/migrations" \
		--url $(MYSQL_URL)
