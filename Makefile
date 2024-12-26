COVERAGE_FILE ?= coverage.out

TARGET ?= run # CHANGE THIS TO YOUR BINARY NAME/NAMES

.PHONY: build
build:
	@echo "Выполняется go build для таргета ${TARGET}"
	@mkdir -p .bin
	@go build -o ./bin/${TARGET} ./cmd/${TARGET}

## test: run all tests
.PHONY: test
test:
	@go test -coverpkg='github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/...' --race -count=1 -coverprofile='$(COVERAGE_FILE)' ./...
	@go tool cover -func='$(COVERAGE_FILE)' | grep ^total | tr -s '\t'
