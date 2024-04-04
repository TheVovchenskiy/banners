# Banners service <!-- omit from toc -->

This is a repository with a completed test task for Backend internsheep in Avito.

## Table of contents <!-- omit from toc -->
- [1. Starting](#1-starting)
  - [1.1. Additioanl tools](#11-additioanl-tools)
  - [1.2. Docker](#12-docker)
  - [1.3. Migrations](#13-migrations)


## 1. Starting

You will need to create `.env` file in the root directory of this repository with the following format:

```bash
PG_USER=<your_postgres_user>
PG_PASSWORD=<your_postgres_user_password>
PG_DBNAME=<your_postgres_db_name>
PG_PORT=<your_postgres_port>
SECRET_KEY=<your_secret_key_to_generate_jwt>
```

### 1.1. Additioanl tools

Additional tools you may need to run this project with commands to install them:

- [Tern](https://github.com/jackc/tern) - migration tool for PostgreSQL
```bash
make install-tern
```
- [dotenv-cli](https://www.npmjs.com/package/dotenv-cli) - cli tool to parse `.env` files and load environment variables
```bash
make install-dotenv
```

### 1.2. Docker

To run all necessary docker containers run:
```bash
make compose-up
```

### 1.3. Migrations

To update db to the latest migration  run:
```bash
make migrate
```

To rollback one migration run:
```bash
make rollback
```

To create a new migration run:
```bash
make create-migration name=<name_of_migration>
```
