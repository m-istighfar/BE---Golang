# DRX Backend API

A RESTful API built with Go and Gin framework.

## Setup

1. Set up environment variables:
Create a `.env` file with the following variables:
```
DB_URL=postgresql://user:password@localhost:5432/dbname
APP_ENVIRONMENT=development
```

2. Run database migrations:
```bash
make migrate-up
```

3. Run the application:
```bash
make run
```

## Testing

Run all tests:
```bash
make test
```

Run specific test suites:
- Unit tests: `make test-unit`
- Product handler tests: `make test-product-handler`
- Product usecase tests: `make test-product-usecase`
- Product repository tests: `make test-product-repository`

## Database

Database-related commands:
- Run migrations: `make migrate-up`
- Rollback last migration: `make migrate-down`
- Force migration version: `make force`
- View current migration version: `make version`
- Drop all migrations: `make reset`
- Export schema: `make schema`
