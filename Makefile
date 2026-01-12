.PHONY: test
test: ## Run unit tests
	go test -race -count=1 -shuffle=on -timeout=12m ./...

.PHONY: lint
lint: ## Run golangci linters
	golangci-lint run --output.text.print-issued-lines=false

.PHONY: fmt
fmt: ## Run golangci formatters
	golangci-lint fmt

.PHONY: help
help: ## Show available targets
	@grep -E '^[a-z]+(-[a-z]+)*:.*?## .+$$' $(MAKEFILE_LIST) | awk \
		'BEGIN {FS=":.*?## "} {printf "%-8s %s\n", $$1, $$2}'
