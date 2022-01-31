ROOT_DIR=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

docker-sqlc-generate: ## Generate sqlc repository
	docker run --rm -v $(ROOT_DIR):/src -w /src kjconroy/sqlc:1.11.0 generate