# GoodMeh? Backend (Go)

Backend service for the GoodMeh? application built with Go.

## Prerequisites

- Go 1.23+
- PostgreSQL database
- `.env` file with required environment variables

## Setup

1. Clone the repository
2. Download and install Go from https://golang.org/dl/
3. Set up your environment variables in a `.env` file:
   ```bash
   cp .env.example .env
   ```
4. Generate Wire dependencies:
   ```bash
   go install github.com/google/wire/cmd/wire@latest
   wire ./deps
   ```
5. Run `go mod download` to download all dependencies

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
  /repository  - Database access layer. This is where sqlc generated files are stored
  /mapper      - Data mapping layer
  /router      - API routes configuration
  /service     - Business logic layer
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

- **Testing**: Run tests (TODO: Add tests):
  ```bash
  go test ./...
  ```

## Project Stack

- **Web Framework**: Gin
- **Database**: PostgreSQL with pgx driver
- **SQL Tools**: sqlc for type-safe SQL
- **Dependency Injection**: Google Wire