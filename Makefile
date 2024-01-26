# assets
C_GREEN=\033[0;32m
C_RED=\033[0;31m
C_BLUE=\033[0;34m
C_END=\033[0m

# environment
PROTO_DIR	:= protobuf
PROTO_FILE	:= zkproof.proto
BIN_FOLDER	:= bin
CLIENT_NAME	:= zkproof-client
SERVER_NAME	:= zkproof-server

# optimize built binaries
LDFLAGS		:= "-w -s"

######
# Go #
######
.PHONY: all
all: clean generate-proto test server client

.PHONY: client
client:
	@echo "\n\t$(C_GREEN)# Build binary $(BINARY)$(C_END)"
	go build -trimpath -ldflags $(LDFLAGS) -o $(BIN_FOLDER)/$(CLIENT_NAME) client/client.go

.PHONY: server
server:
	@echo "\n\t$(C_GREEN)# Build binary $(BINARY)$(C_END)"
	go build -trimpath -ldflags $(LDFLAGS) -o $(BIN_FOLDER)/$(SERVER_NAME) server/server.go

.PHONY: clean
clean:
	@echo "\n\t$(C_GREEN)# Cleaning environment$(C_END)"
	go clean -x
	rm -rf $(BIN_FOLDER)

############
# Protobuf #
############
.PHONY: generate-proto
generate-proto:
	@echo "\n\t$(C_GREEN)# Generating Go protobuf bindings$(C_END)"
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(PROTO_DIR)/$(PROTO_FILE)

########
# Test #
########
.PHONY: test
test:
	@echo "\n\t$(C_GREEN)# Run zk test$(C_END)"
	go test -v ./...

.PHONY: e2e
e2e:
	@echo "\n\t$(C_GREEN)# Run e2e zk test$(C_END)"
	@docker rm --force zkproof-server
	@docker rm --force zkproof-client-success
	@docker rm --force zkproof-client-failure
	docker-compose up

##############
# containers #
##############

.PHONY: server-img
server-img:
	@echo "\n\t$(C_GREEN)# Build zkproog-server container image$(C_END)"
	docker build -f builds/Dockerfile.server -t zkproof-server:latest .

.PHONY: client-img
client-img:
	@echo "\n\t$(C_GREEN)# Build zkproog-client container image$(C_END)"
	docker build -f builds/Dockerfile.client -t zkproof-client:latest .

.PHONY: server-img-run
server-img-run:
	@docker rm --force zkproof-server
	docker run --rm -ti --network host -p 50051:50051/tcp --name zkproof-server -d zkproof-server:latest