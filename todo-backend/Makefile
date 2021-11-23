GREEN := "\033[0;32m"
NC := "\033[0;0m"

LOG := @sh -c '\
	   printf ${GREEN}; \
	   echo -e "\n> $$1\n"; \
	   printf ${NC}' VALUE

generate:
	$(LOG) "Regenerating from schema..."	
	@go run github.com/99designs/gqlgen generate

test: test-acceptance test-unit

test-unit:
	$(LOG) "Running unit tests"	
	@go test -v ./...


test-acceptance:
	$(LOG) "Running acceptance tests"	
	@go test -v -race -tags=acceptance ./testing/acceptance/...

