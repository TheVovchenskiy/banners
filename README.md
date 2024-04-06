# Banner service <!-- omit from toc -->

## Contents <!-- omit from toc -->
- [1. About](#1-about)
  - [1.1. Stack](#11-stack)
- [2. Getting started](#2-getting-started)
  - [2.1. Additioanl tools](#21-additioanl-tools)
  - [2.2. Docker](#22-docker)
  - [2.3. Migrations](#23-migrations)
- [3. Usage](#3-usage)
  - [3.1. Documentation](#31-documentation)
- [4. Additional problems and questions](#4-additional-problems-and-questions)
  - [4.1. Bad naming of endpoints](#41-bad-naming-of-endpoints)
  - [4.2. Additional fields](#42-additional-fields)

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
- [golangci-lint](https://golangci-lint.run/) - linter for golang

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

## 3. Usage

### 3.1. Documentation

To get documentation on Banners API follow `http://localhost:8081/swagger/index.html` url while app is running.

## 4. Additional problems and questions

Here are additional problems and questions that I encountered, and my logic for solving them.

### 4.1. Bad naming of endpoints

According to given [API](https://drive.google.com/file/d/1l4PMTPzsjksRCd_lIm0mVfh4U0Jn-A2R/view) there must be 2 endpoints: `/banner` and `/banner/{id}` but this is a bad example of [API naming](https://medium.com/@nadinCodeHat/rest-api-naming-conventions-and-best-practices-1c4e781eb6a5), that is why I changed its names to plural: `/banners` and `/banners/{id}`.

### 4.2. Additional fields

To improve data readability I added additional fields for the following entities:

- `feature` - added field `description`
- `tag` - added field `name`
