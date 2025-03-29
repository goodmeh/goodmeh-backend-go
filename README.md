# GoodMeh? Backend (Go)

Backend service for the GoodMeh? application built with Go.

## Prerequisites

- Go 1.23+
- PostgreSQL database
- `.env` file with required environment variables

## Setup

1. Clone the repository
2. Download and install Go from https://golang.org/dl/
3. Set the `GOPRIVATE` environment variable for private module:
  ```bash
  export GOPRIVATE=github.com/goodmeh/backend-private
  ```
4. Download Wire:
  ```bash
  go install github.com/google/wire/cmd/wire@latest
  ```
5. Download Goose:
  ```bash
  go install github.com/pressly/goose/v3/cmd/goose@latest
  ```
6. Set up your environment variables in a `.env` file:
  ```bash
  cp .env.example .env
  ```
7. Run `go mod download` to download all dependencies

## Running the Application

Start the server:
```bash
go run main.go
```

Or build and run:
```bash
go build -o goodmeh-app
./goodmeh-app
```

## Project Structure

```
/app
  /controller  - HTTP request handlers
  /dto         - Data Transfer Objects for API responses
  /events      - Event Bus and event handlers
  /repository  - Database access layer. This is where sqlc generated files are stored
  /mapper      - Data mapping layer
  /router      - API routes configuration
  /service     - Business logic layer
  /socket      - WebSocket handling
  /summarizer  - OpenAI summarization service
  /utils       - Utility functions
/db            - Database migration files and queries for sqlc
/deps          - Dependency injection setup (Wire)
```

## Development

- **Code Generation**: After modifying SQL queries, regenerate:
  ```bash
  sqlc generate
  ```
  
- **Dependency Injection**: After modifying dependencies, regenerate:
  ```bash
  wire ./deps
  ```
  Ensure that `$GOPATH/bin` is added to your `$PATH`.

- **Database Migrations**: Run migrations using Goose:
  ```bash
  goose up
  ```
  To create a new migration:
  ```bash
  goose -s create <migration_name> sql
  ```

- **Testing**: Run tests (TODO: Add tests):
  ```bash
  go test ./...
  ```

## Project Stack

- **Web Framework**: Gin
- **Database**: PostgreSQL with pgx driver
- **SQL Tools**: sqlc for type-safe SQL
- **Dependency Injection**: Google Wire
- **Migration Tool**: Goose