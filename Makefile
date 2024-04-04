.PHONY: dotenv
dotenv:
	echo "PG_USER=<your_postgres_user>\n\
	PG_PASSWORD=<your_postgres_user_password>\n\
	PG_DBNAME=<your_postgres_db_name>\n\
	PG_PORT=<your_postgres_port>\n\
	SECRET_KEY=<your_secret_key_to_generate_jwt>" > .env

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


.PHONY: swag-init
swag-init:
	swag init -d cmd/banners/

.PHONY: run
run:
	go run cmd/banners/main.go
