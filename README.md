# URL Shortener Service

A high-performance URL shortener service built with Go and Clean Architecture principles. This service provides RESTful APIs for creating shortened URLs and redirecting users to original URLs.

## 🚀 Features

- **Fast URL Shortening**: Generate short URLs from long URLs using Snowflake ID generation
- **Redirect Service**: Redirect users from short URLs to original URLs
- **Click Tracking**: Track click counts for shortened URLs
- **Clean Architecture**: Follows Clean Architecture principles for maintainability
- **MySQL Database**: Uses MySQL with GORM for data persistence
- **RESTful API**: Clean and intuitive REST API endpoints
- **Environment Configuration**: Configurable via environment variables

## 📋 Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Database Setup](#database-setup)
- [API Documentation](#api-documentation)
- [Project Structure](#project-structure)

## 🛠 Installation

### Prerequisites

- Go 1.24.5 or later
- MySQL 5.7 or later
- Git

### Steps

1. **Clone the repository**
   ```bash
   git clone https://github.com/harmancioglue/url-shortener.git
   cd url-shortener
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   ```bash
   cp .env.dist .env
   # Edit .env with your configuration
   ```

4. **Set up the database**
   ```bash
   # Create database and run migrations (see Database Setup section)
   ```

5. **Run the application**
   ```bash
   go run cmd/main.go
   ```

The service will start on `http://localhost:8080` by default.

## ⚙️ Configuration

The service uses environment variables for configuration. Create a `.env` file:

```env
# Server Configuration
SERVER_HOST=localhost
SERVER_PORT=8080

# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=url_shortener
DB_SSLMODE=disable

# Snowflake ID Generator
WORKER_ID=1
```

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `SERVER_HOST` | Server host | `localhost` |
| `SERVER_PORT` | Server port | `8080` |
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `3306` |
| `DB_USER` | Database username | `root` |
| `DB_PASSWORD` | Database password | `""` |
| `DB_NAME` | Database name | `url_shortener` |
| `DB_SSLMODE` | Database SSL mode | `disable` |
| `WORKER_ID` | Snowflake worker ID | `1` |

## 🗄️ Database Setup

### Create Database

```sql
CREATE DATABASE url_shortener CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### Create Table

```sql
USE url_shortener;

CREATE TABLE urls (
    id           bigint auto_increment primary key,
    short_code   varchar(255) not null,
    original_url text not null,
    created_at   timestamp default CURRENT_TIMESTAMP not null,
    updated_at   timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    expires_at   timestamp null,
    click_count  bigint default 0 null,
    is_active    tinyint(1) default 1 null,
    constraint short_code unique (short_code)
) collate = utf8mb4_unicode_ci;

CREATE INDEX idx_short_code ON urls (short_code);
```

### Table Structure

| Column | Type | Description |
|--------|------|-------------|
| `id` | bigint | Primary key, auto-increment |
| `short_code` | varchar(255) | Unique short code for URL |
| `original_url` | text | Original long URL |
| `created_at` | timestamp | Record creation time |
| `updated_at` | timestamp | Last update time |
| `expires_at` | timestamp | URL expiration time (optional) |
| `click_count` | bigint | Number of times URL was accessed |
| `is_active` | tinyint(1) | Whether URL is active (1=active, 0=inactive) |



## 🏗️ Project Structure

```
url-shortener/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── api/
│   │   └── http/
│   │       ├── api.go          # API setup and routing
│   │       └── handlers/
│   │           └── url_shortener.go  # HTTP handlers
│   ├── app/
│   │   └── application.go      # Application initialization
│   ├── common/
│   │   └── utils/
│   │       └── base62.go      # Base62 encoding utilities
│   ├── config/
│   │   └── config.go          # Configuration management
│   ├── domain/
│   │   ├── entity/
│   │   │   └── url.go          # URL entity
│   │   ├── repository/
│   │   │   └── url_repository.go  # Repository interface
│   │   └── service/
│   │       ├── id_generator.go     # ID generator interface
│   │       └── url_service.go      # URL service interface
│   ├── dto/
│   │   ├── request/
│   │   │   └── url_request.go      # Request DTOs
│   │   └── response/
│   │       ├── base_response.go    # Base response structure
│   │       └── url_response.go     # URL response DTOs
│   ├── infrastructure/
│   │   └── repository/
│   │       └── mysql/
│   │           └── url_repository.go  # MySQL repository implementation
│   └── services/
│       ├── snowflake_id_generator.go  # Snowflake ID generator
│       ├── snowflake_id_generator_test.go  # ID generator tests
│       └── url_service.go         # URL service implementation
├── .env                         # Environment variables
├── .env.dist                    # Environment variables template
├── go.mod                       # Go module file
├── go.sum                       # Go module checksums
└── README.md                    # This file
```

## 🧪 Testing

### Run All Tests

```bash
go test ./...
```

### Run Specific Test

```bash
go test ./internal/services -v
```

### Run Tests with Coverage

```bash
go test -cover ./...
```

### Test Coverage Report

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 🏛️ Architecture

This project follows **Clean Architecture** principles with clear separation of concerns:

### Layers

1. **Presentation Layer** (`internal/api/`): HTTP handlers and routing
2. **Application Layer** (`internal/app/`): Application coordination
3. **Domain Layer** (`internal/domain/`): Business logic and entities
4. **Infrastructure Layer** (`internal/infrastructure/`): External concerns

### Key Patterns

- **Repository Pattern**: Abstracts data access
- **Service Layer Pattern**: Encapsulates business logic
- **Dependency Injection**: Promotes loose coupling
- **DTO Pattern**: Separates internal/external representations


## 🆘 Support

If you have any questions or issues, please open an issue on GitHub.

## 🙏 Acknowledgments

- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) by Robert C. Martin
- [Snowflake](https://developer.twitter.com/en/docs/twitter-ids) ID generation algorithm
- [GORM](https://gorm.io/) for Go ORM
- [Fiber](https://gofiber.io/) for web framework