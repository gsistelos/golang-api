.PHONY: proto
proto:
	buf generate --template proto/buf.gen.yaml proto

.PHONY: runapi
runapi:
	cd cmd/api && go run .
