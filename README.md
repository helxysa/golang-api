# Go API Catalog

This is a simple REST API built with Go, Gin, and GORM for managing solutions catalog.

## Prerequisites

- Go (version 1.16 or higher)
- PostgreSQL
- Git

## Installation

### 1. Clone the repository

```bash
git clone <repository-url>
cd go-api-catalog
```

### 2. Install dependencies

```bash
go mod tidy
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

## Database Setup

### 1. Create PostgreSQL database

Make sure PostgreSQL is installed and running on your system. Then create a new database:

```bash
psql -U postgres
```

In the PostgreSQL shell:

```sql
CREATE DATABASE "teste-go";
\q
```

### 2. Configure database connection

The database connection string is located in `database/postgres.go`. Update it if needed:

```go
dsn := "host=localhost port=5432 user=postgres password=123@senha dbname=teste-go sslmode=disable"
```

Replace the values with your PostgreSQL configuration.

## Project Structure

```
go-api-catalog/
├── cmd/
│   ├── main.go           # Main application entry point
│   └── migrate/
│       └── main.go       # Database migration tool
├── controller/
│   └── solucoes_controller.go  # API controllers
├── database/
│   ├── migrations/
│   │   ├── create_solucoes.sql # SQL migration file
│   │   └── migrate.go          # Migration runner
│   └── postgres.go      # Database connection
├── models/
│   └── solucao.go       # Data models
└── README.md            # This file
```

## Running Migrations

Before running the application, you need to run the migrations to create the necessary tables:

```bash
go run cmd/migrate/main.go
```

## Running the Application

To start the API server:

```bash
go run cmd/main.go
```

The server will start on port 8000 by default.

## API Endpoints

### GET /teste

A simple test endpoint that returns a "Hello World" message.

Example:
```bash
curl http://localhost:8000/teste
```

### GET /solucoes

Retrieves all solutions from the database.

Example:
```bash
curl http://localhost:8000/solucoes
```

### POST /solucoes

Creates a new solution.

Example:
```bash
curl -X POST http://localhost:8000/solucoes \
  -H "Content-Type: application/json" \
  -d '{
    "demanda_id": 1,
    "nome": "Solution Name",
    "sigla": "SN",
    "versao": "1.0",
    "link": "https://example.com",
    "andamento": "In Progress",
    "repositorio": "https://github.com/example/repo",
    "criticidade": "High",
    "data_status": "2023-04-09"
  }'
```

## Model Structure

The `Solucao` model has the following fields:

```go
type Solucao struct {
    ID          int    `json:"id_solucoes" gorm:"column:id;primaryKey"`
    DemandaID   int    `json:"demanda_id" gorm:"column:demanda_id"`
    Nome        string `json:"nome" gorm:"column:nome"`
    Sigla       string `json:"sigla" gorm:"column:sigla"`
    Versao      string `json:"versao" gorm:"column:versao"`
    Link        string `json:"link" gorm:"column:link"`
    Andamento   string `json:"andamento" gorm:"column:andamento"`
    Repositorio string `json:"repositorio" gorm:"column:repositorio"`
    Criticidade string `json:"criticidade" gorm:"column:criticidade"`
    DataStatus  string `json:"data_status" gorm:"column:data_status"`
}
```

## Using GORM

This project uses GORM as the ORM (Object-Relational Mapping) library. Here are some examples of how GORM is used:

### Database Connection

```go
import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
    dsn := "host=localhost port=5432 user=postgres password=123@senha dbname=teste-go sslmode=disable"
    
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    return db
}
```

### Querying Data

```go
// Get all solutions
var solucoes []model.Solucao
db.Find(&solucoes)

// Get solution by ID
var solucao model.Solucao
db.First(&solucao, id)

// Create a new solution
db.Create(&solucao)
```

## Troubleshooting

### Database Connection Issues

If you encounter database connection issues, check:

1. PostgreSQL is running
2. The database "teste-go" exists
3. The connection string in `database/postgres.go` is correct
4. The PostgreSQL user has the correct permissions

### API Errors

If you receive a 404 error when making a POST request to `/solucoes`, ensure:

1. The server is running
2. The JSON payload is correctly formatted
3. All required fields are included in the request

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
