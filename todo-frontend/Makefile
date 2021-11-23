GREEN := "\033[0;32m"
NC := "\033[0;0m"

LOG := @sh -c '\
	   printf ${GREEN}; \
	   echo -e "\n> $$1\n"; \
	   printf ${NC}' VALUE

lint:
	$(LOG) "Linting..."
	@yarn lint


test:
	$(LOG) "Testing..."
	@yarn test --ci --collect-coverage --verbose


build:
	$(LOG) "Building..."
	@yarn build

