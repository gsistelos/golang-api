all: proto sqlc run

.PHONY: proto
proto:
	buf generate --template proto/buf.gen.yaml proto

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: run
run:
	cd cmd/api && go run .
