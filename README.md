# Banner service <!-- omit from toc -->

## Table of contents <!-- omit from toc -->
- [1. About](#1-about)
  - [1.1. Stack](#11-stack)
- [2. Getting started](#2-getting-started)
  - [2.1. Additioanl tools](#21-additioanl-tools)
  - [2.2. Docker](#22-docker)
  - [2.3. Migrations](#23-migrations)

## 1. About

This is a repository with a completed test task for Backend internsheep in Avito.

### 1.1. Stack

Here are some main technologies used in this project:

- main language - `Go`
- main database - `Postgresql`
- deploy - `Docker` and `Docker compose`

## 2. Getting started

You will need to create `.env` file in the root directory of this repository with the following format:

```bash
PG_USER=<your_postgres_user>
PG_PASSWORD=<your_postgres_user_password>
PG_DBNAME=<your_postgres_db_name>
PG_PORT=<your_postgres_port>
SECRET_KEY=<your_secret_key_to_generate_jwt>
```

To create this `.env` template simply run:

```bash
make dotenv
```

### 2.1. Additioanl tools

Additional tools you may need to install for this project with commands to install them:

- [tern](https://github.com/jackc/tern) - migration tool for PostgreSQL
```bash
make install-tern
```
- [dotenv-cli](https://www.npmjs.com/package/dotenv-cli) - cli tool to parse `.env` files and load environment variables
```bash
make install-dotenv
```

### 2.2. Docker

To run all necessary Docker containers run:
```bash
make compose-up
```

### 2.3. Migrations

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
