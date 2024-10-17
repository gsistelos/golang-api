BUF_GENERATE = buf generate --template proto/buf.gen.yaml proto

SQLC_GENERATE = sqlc generate -f sqlc.yaml

DOCKER_COMPOSE = docker compose -f docker-compose.yaml

CMD_DIR = cmd

CMDS = $(wildcard $(CMD_DIR)/*)

BIN_DIR = bin

BINARIES = $(patsubst $(CMD_DIR)/%, $(BIN_DIR)/%, $(CMDS))

# Set variables if not set
ifeq ($(MYSQL_USER),)
	MYSQL_USER = root
endif

ifeq ($(MYSQL_ADDRESS),)
	MYSQL_ADDRESS = mysql:3306
endif

ifeq ($(MYSQL_DATABASE), mysql)
	MYSQL_DATABASE = mysql
endif

MYSQL_URL = "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@$(MYSQL_ADDRESS)/$(MYSQL_DATABASE)"


all: proto sqlc build

.PHONY: proto
proto:
	$(BUF_GENERATE)

.PHONY: sqlc
sqlc:
	$(SQLC_GENERATE)

.PHONY: build
build: $(BINARIES)

$(BIN_DIR)/%: $(CMD_DIR)/%
	go build -o $@ ./$<

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

.PHONY: docker
docker:
	$(DOCKER_COMPOSE) up -d

.PHONY: dclean
dclean:
	-$(DOCKER_COMPOSE) down
	-docker rm $$(docker ps -aq)
	-docker network rm -f $$(docker network ls -q)
	-docker volume rm -f $$(docker volume ls -q)

.PHONY: mdiff
mdiff:
	atlas migrate diff \
		--dir "file://db/migrations" \
		--to "file://db/schema.hcl" \
		--dev-url "docker://mysql/8/example"

.PHONY: mapply
mapply:
	atlas migrate apply \
		--dir "file://db/migrations" \
		--url $(MYSQL_URL)
