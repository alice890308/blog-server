PATH := $(CURDIR)/bin:$(PATH)

DOCKER_COMPOSE := $(or $(DOCKER_COMPOSE),$(DOCKER_COMPOSE),docker-compose)

# lint

.PHONY: dc.lint
dc.lint:
	$(DOCKER_COMPOSE) run --rm lint

lint:
	golangci-lint run ./...
