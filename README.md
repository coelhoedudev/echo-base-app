# infra-base-go

A simple Go project base ready for REST APIs, following best practices and a clean architecture.

## Dependencies

This project uses the following main dependencies:

- [Go 1.21+](https://golang.org/dl/)
- [Echo](https://github.com/labstack/echo) - Web framework
- [Go-Playground/Validator](https://github.com/go-playground/validator) - Request validation
- [github.com/golang-migrate/migrate/v4](https://github.com/golang-migrate/migrate) - Database migrations
- [Testify](https://github.com/stretchr/testify) - Testing assertions
- Other dependencies as defined in `go.mod`

To see all dependencies, check the `go.mod` file.

## Getting Started

### Prerequisites

- Go 1.21 or newer installed. Download from [golang.org](https://golang.org/dl/).
- A running database (e.g., PostgreSQL).
- [make](https://www.gnu.org/software/make/) (optional, for easier commands).

### Clone the project

```sh
git clone https://github.com/your-username/infra-base-go.git
cd infra-base-go
```

### Install dependencies

```sh
go mod download
```

### Setting Up Environment Variables

Create a `.env` file or use environment variables for configuration like database connection strings. Example:

```env
DB_URL=postgres://user:password@localhost:5432/yourdb?sslmode=disable
PORT=8080
```

### Running Database Migrations

```sh
migrate -path db/migrations -database "$DB_URL" up
```

Or using Docker:

```sh
docker run --rm -v ${PWD}/db/migrations:/migrations --network host migrate/migrate \
  -path=/migrations -database "$DB_URL" up
```

### Running the Service

```sh
go run cmd/main.go
```

Or with `make` (if a Makefile is provided):

```sh
make run
```

The server will start on the port configured in your environment (default: `8080`).

### Build for Production

```sh
go build -o bin/app cmd/main.go
```

### Testing

```sh
go test ./...
```

## Project Structure

```
cmd/                  # Main application entrypoint
internal/             # Application modules (domain, services, etc.)
  domain/
  pkg/
db/
  migrations/         # SQL migration files
configs/              # Configuration files
```

## Creating Migrations

Generate a new migration file:

```sh
migrate create -ext sql -dir db/migrations -seq create_users_table
```

This creates `xxxx_create_users_table.up.sql` and `.down.sql` files.

## API Endpoints

Typical usage with Echo. See the `internal/domain/*/handler.go` files for endpoint examples.

## Contributing

1. Open issues or pull requests with improvements.
2. Maintain code style and best practices.

## License

MIT

---

// Replace 'your-username' in git URLs with the actual repository location.
