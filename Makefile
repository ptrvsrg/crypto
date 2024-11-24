.OUTPUT_DIR = build
.CRYPTO_STREAM_APP_DIR = cmd/stream
.CRYPTO_BLOCK_APP_DIR = cmd/block
.CRYPTO_HASH_APP_DIR = cmd/hash

build:
	go mod tidy
	go build -o $(.OUTPUT_DIR)/crypto-stream $(.CRYPTO_STREAM_APP_DIR)/main.go
	go build -o $(.OUTPUT_DIR)/crypto-block $(.CRYPTO_BLOCK_APP_DIR)/main.go
	go build -o $(.OUTPUT_DIR)/crypto-hash $(.CRYPTO_HASH_APP_DIR)/main.go

clean:
	rm -rf $(OUTPUT_DIR)

test:
	go test ./... -v

help:
	@echo "Available commands:"
	@echo "	make help		- print this help"
	@echo "	make build		- build crypto-stream and crypto-block"
	@echo "	make clean		- clean build directory"
	@echo "	make test		- run tests"

.PHONY: build clean test help
.DEFAULT_GOAL := help