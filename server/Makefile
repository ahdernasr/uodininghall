PROTO_DIR := proto
GEN_DIR := proto/gen

PROTOC ?= protoc

PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)

.PHONY: proto
proto: ## Generate Go gRPC stubs
	@mkdir -p $(GEN_DIR)
	$(PROTOC) \
		-I=$(PROTO_DIR) \
		--go_out=$(GEN_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)

.PHONY: run-scraper
run-scraper:
	go run ./services/scraper

.PHONY: run-subscribers
run-subscribers:
	go run ./services/subscribers

.PHONY: run-mailer
run-mailer:
	go run ./services/mailer

.PHONY: run-api
run-api:
	go run ./services/api

.PHONY: run-scheduler
run-scheduler:
	go run ./services/scheduler


