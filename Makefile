PATH := $(CURDIR)/bin:$(PATH)

MODULES := api

DOCKER_COMPOSE := $(or $(DOCKER_COMPOSE),$(DOCKER_COMPOSE),docker-compose)

# lint

.PHONY: dc.lint
dc.lint:
	$(DOCKER_COMPOSE) run --rm lint

lint:
	golangci-lint run ./...



# generate
define make-dc-generate-rules

.PHONY: dc.$1.generate

dc.$1.generate:
	$(DOCKER_COMPOSE) run --rm generate make $1.generate

endef
$(foreach module,$(MODULES),$(eval $(call make-dc-generate-rules,$(module))))

.PHONY: dc.pkg.generate
dc.pkg.generate:
	$(DOCKER_COMPOSE) run --rm generate make pkg.generate

.PHONY: dc.generate
dc.generate:
	$(DOCKER_COMPOSE) run --rm generate

define make-generate-rules

$1.generate: bin/protoc-gen-go bin/protoc-gen-go-grpc bin/protoc-gen-grpc-gateway
	protoc \
		-I . \
		-I ./pkg/pb \
		--proto_path=./modules/$1/proto --go_out=./modules/$1 \
		--proto_path=./modules/$1/proto --go-grpc_out=./modules/$1 \
		--proto_path=./modules/$1/proto --grpc-gateway_out=./modules/$1 \
		./modules/$1/proto/*.proto
	
	go generate ./modules/$1/...

endef
$(foreach module,$(MODULES),$(eval $(call make-generate-rules,$(module))))

pkg.generate: bin/protoc-gen-go bin/protoc-gen-go-grpc bin/protoc-gen-grpc-gateway
	go generate ./pkg/...

generate: pkg.generate $(addsuffix .generate,$(MODULES))

bin/protoc-gen-go: go.mod
	go build -o $@ google.golang.org/protobuf/cmd/protoc-gen-go

bin/protoc-gen-go-grpc: go.mod
	go build -o $@ google.golang.org/grpc/cmd/protoc-gen-go-grpc

bin/protoc-gen-grpc-gateway: go.mod
	go build -o $@ github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
