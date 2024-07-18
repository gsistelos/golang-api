all: proto run

.PHONY: proto
proto:
	buf generate --template proto/buf.gen.yaml proto

.PHONY: run
run:
	cd cmd/api && go run .
