DEFAULT_GOAL := help

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

start-infra: ## Run infrastructure locally
	docker-compose -f docker-compose.yml up -d --build

stop-infra: ## Stop infrastructure locally
	docker-compose -f docker-compose.yml down

cleanup-infra: ## remove all data from docker volume
	docker stop postgres_coinwallet \
    & docker rm postgres_coinwallet \
    & docker volume rm postgres_coinwallett