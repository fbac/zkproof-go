# assets
C_GREEN=\033[0;32m
C_RED=\033[0;31m
C_BLUE=\033[0;34m
C_END=\033[0m

# environment
PROTO_DIR	:= protobuf
PROTO_FILE	:= zkproof.proto

.PHONY: gen-grpc
generate-proto:
	@echo "\n\t$(C_GREEN)# Generating Go protobuf bindings$(C_END)"
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(PROTO_DIR)/$(PROTO_FILE)