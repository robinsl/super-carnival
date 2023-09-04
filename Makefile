DSN=edgedb://edgedb:password@localhost:5656
DSN_DB=$(DSN)/super-carnival


.PHONY: start-dev
start-dev:
	@echo "Starting development environment..."
	docker-compose -f deployments/docker-compose.yml up -d

.PHONY: setup
setup:
	@echo "Setting up development environment..."
	docker-compose -f deployments/docker-compose.yml up -d
	docker-compose -f deployments/docker-compose.yml exec edgedb edgedb database create --dsn=$(DSN) --tls-security=insecure super-carnival

.PHONY: clean
clean:
	docker-compose -f deployments/docker-compose.yml down -v

.PHONY: generate-migration
generate-migration:
	@echo "Generating migration..."
	docker-compose -f deployments/docker-compose.yml exec edgedb edgedb --dsn=$(DSN_DB) --tls-security=insecure migration create

.PHONY: run-migration
run-migration:
	@echo "Running migration..."
	docker-compose -f deployments/docker-compose.yml exec edgedb edgedb --dsn=$(DSN_DB) --tls-security=insecure migration apply