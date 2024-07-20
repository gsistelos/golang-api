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
	buf generate --template proto/buf.gen.yaml proto

.PHONY: sqlc
sqlc:
	sqlc generate -f sqlc/sqlc.yaml

.PHONY: run
run:
	cd cmd/api && go run .

.PHONY: docker
docker:
	docker compose up -d

.PHONY: dclean
dclean:
	-docker compose down
	-docker rm $$(docker ps -aq)
	-docker network rm -f $$(docker network ls -q)
	-docker volume rm -f $$(docker volume ls -q)

.PHONY: migrations
migrations:
	atlas migrate diff \
		--dir "file://migrations" \
		--to "file://schema.hcl" \
		--dev-url "docker://mysql/8/example"

.PHONY: migrate
migrate:
	atlas migrate apply \
		-u $(MYSQL_URL)
