.PHONY: install-tern
install-tern:
	go install github.com/jackc/tern/v2@latest

.PHONY: install-dotenv
install-dotenv:
	sudo npm install -g dotenv-cli

.PHONY: compose-up
compose-up: 
	docker compose up -d

.PHONY: create-migration
create-migration:
	tern new -m migrations/ $(name)

.PHONY: migrate
migrate:
	dotenv -- tern migrate -m migrations/

.PHONY: rollback
rollback:
	dotenv -- tern migrate -m migrations/ -d -1
