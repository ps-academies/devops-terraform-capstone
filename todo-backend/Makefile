GREEN := "\033[0;32m"
NC := "\033[0;0m"

LOG := @sh -c '\
	   printf ${GREEN}; \
	   echo -e "\n> $$1\n"; \
	   printf ${NC}' VALUE

VERSION ?= 0.1.0

generate:
	$(LOG) "Regenerating from schema..."
	@go run github.com/99designs/gqlgen generate


build:
	$(LOG) "Building backend..."
	@go build ./...



binary:
	$(LOG) "Building backend with version ${VERSION}..."
	@go build -o dist/todo -ldflags="-X main.version=${VERSION}" .


test: test-acceptance test-unit

test-unit:
	$(LOG) "Running backend unit tests.."
	@go test -v ./...

test-acceptance:
	$(LOG) "Running acceptance tests"
	@go test -v -race -tags=acceptance ./testing/acceptance/...


run:
	$(LOG) "Running backend..."
	@go run server.go

