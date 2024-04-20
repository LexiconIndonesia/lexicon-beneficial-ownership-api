.EXPORT_ALL_VARIABLES:
OUT_DIR := ./_output
BIN_DIR := ./bin

$(shell mkdir -p $(OUT_DIR) $(BIN_DIR))

# Main Test Targets (without docker)
.PHONY: test
test:
	go test -race -coverprofile=$(OUT_DIR)/coverage.out ./...

.PHONY: integration-test
integration-test:
	go test -race -tags=integration -coverprofile=$(OUT_DIR)/coverage.out ./...

.PHONY: install
install:
	go mod tidy && go mod vendor

.PHONY: run-dev
run-dev:
	docker-compose up --build
.PHONY: seed
seed:
	go run database/seeders/seeders.go
