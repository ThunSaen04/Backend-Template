# 🚀 Go Fiber v3 + GORM Enterprise Backend Template

A production-ready RESTful backend template built with Go (Golang). Designed following **Clean Architecture** principles and optimized for building scalable microservices and APIs.

## 🌟 Key Features

- ⚡ **High Performance**: Built on `fiber/v3`, an Express-inspired web framework for Go.
- 🏗️ **Clean Architecture**: Decoupled layers (Handler → Service → Repository → Model).
- 🗄️ **Multi-Database Support**: Utilizes `GORM` with dynamic connection loading. Supports **PostgreSQL**, **MySQL**, **SQL Server**, and **SQLite** out of the box via the `DB_TYPE` environment variable. Includes Auto-Migration capability.
- 🔐 **Advanced Authentication**:
  - Dual Token Approach: **Access Token** (6h) & **Refresh Token** (24h) via JWT.
  - Token Rotation & DB Revocation system for high security.
- 🛡️ **Role-Based Access Control (RBAC)**: Level-based Role verification (`member`, `admin`).
- 🚥 **Security & Protection**: Built-in **Rate Limiting** via Fiber middleware to prevent brute-force attacks and abuse.
- ✅ **Data Validation**: Automatic struct validation via `go-playground/validator/v10` with human-readable error messages.
- 📦 **Standard Response Wrapper**: Consistent JSON response structure across all endpoints (`success`, `message`, `data`, `meta`, `errors`).
- 📄 **Pagination Utility**: Reusable pagination with GORM Scope and metadata calculation.
- 🐳 **Docker Ready**: Multi-stage Dockerfile & Docker Compose with PostgreSQL, MySQL, and SQL Server support.
- 📚 **Auto-Generated Swagger Docs**: Powered by `swaggo`, supporting both Global and Module-specific API documentation UI.
- ⚙️ **Environment configuration**: Clean setup using `.env` (`joho/godotenv`).

---

## 📂 Project Structure

```text
Backend_Template/
├── docs/                                  # Global Swagger Docs (Auto-generated)
├── internal/
│   ├── config/                            # Environment variable loader
│   ├── database/                          # DB Connection (Multi-driver) & AutoMigrate
│   ├── models/                            # GORM Entities (User, RefreshToken)
│   ├── utils/                             # Shared utilities (Validator, Response, Pagination)
│   └── modules/
│       └── auth/                          # Auth feature module
│           ├── dto/                       # Request/Response data shapes
│           ├── handler/                   # HTTP Handlers (Fiber controllers)
│           ├── repository/                # Database interactions
│           ├── service/                   # Business logic operations
│           ├── utils/                     # Helpers (JWT, Bcrypt, Roles)
│           └── docs/                      # Module-specific Swagger (Auto-generated)
├── middleware/                            # Global & Route middlewares (Auth, Rate Limit, Role)
├── routes/                                # API Routing & Grouping (/api/v1)
├── scripts/                               # Automation bash/bat scripts
├── Dockerfile                             # Multi-stage Docker build
├── docker-compose.yml                     # App + DB containers
├── .env                                   # Environment configuration
└── main.go                                # Entry point
```

---

## 🛠️ Prerequisites

- [Go 1.25+](https://go.dev/dl/)
- A Database (PostgreSQL, MySQL, SQL Server, or SQLite)
- `swag` CLI (for generating Swagger docs):
  ```bash
  go install github.com/swaggo/swag/cmd/swag@latest
  ```
- [Docker](https://www.docker.com/) (optional, for containerized deployment)

---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/ThunSaen04/Backend-Template.git
cd Backend-Template
```

### 2. Configure Environment

Copy the example environment file and adjust your database credentials:

```bash
cp .env.example .env
```

Ensure your Database instance is running and set the `DB_TYPE` in your `.env` to match your target (`postgres`, `mysql`, `sqlserver`, `sqlite`).

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Run the Application

```bash
go run .
```

The server should now be running on `http://localhost:8080`.

---

## 🐳 Docker Deployment

This template uses Docker Compose profiles to manage multiple database environments easily.

### 1. Default (PostgreSQL)
By default, running `docker compose up` will start the application along with a PostgreSQL database container.
**Update `.env`**: `DB_TYPE=postgres`, `DB_HOST=postgres`, `DB_PORT=5432`

```bash
docker compose up -d
```

### 2. Using MySQL
To use MySQL, you need to specify the `mysql` profile.
**Update `.env`**: `DB_TYPE=mysql`, `DB_HOST=mysql`, `DB_PORT=3306`

```bash
docker compose --profile mysql up -d
```

### 3. Using SQL Server
To use SQL Server, you need to specify the `sqlserver` profile. **Note:** SQL Server requires a strong password (e.g., `YourStrong!Passw0rd`).
**Update `.env`**: `DB_TYPE=sqlserver`, `DB_HOST=sqlserver`, `DB_PORT=1433`, `DB_PASSWORD=YourStrong!Passw0rd`

```bash
docker compose --profile sqlserver up -d
```

### 4. Running Only the App (External Database)
If you already have a database hosted elsewhere (e.g., AWS RDS, Supabase, or on your host machine) and just want to run the Go application container:
**Update `.env`**: Set `DB_TYPE`, `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, and `DB_NAME` to match your external database credentials. (If the DB is on your local machine, use `host.docker.internal` for `DB_HOST`).

```bash
docker compose up -d app
```

### 🛑 Stop & Cleanup
To stop the containers and remove all database volumes (wiping data):

```bash
docker compose down -v
```

> ⚠️ **SQLite Limitation**: The Docker image is built with `CGO_ENABLED=0` (static binary) for a minimal image size (~15MB). This means the **SQLite driver is NOT supported** inside Docker containers, as it requires CGO. When deploying with Docker, use **PostgreSQL**, **MySQL**, or **SQL Server** instead. SQLite remains fully functional when running the application directly with `go run .`.

---

## 📦 Standard API Response Format

All endpoints return a consistent JSON structure:

```json
{
  "success": true,
  "message": "Operation successful",
  "data": {},
  "meta": {
    "current_page": 1,
    "per_page": 10,
    "total_rows": 100,
    "total_pages": 10
  },
  "errors": null
}
```

| Field     | Type     | Description                                |
| :-------- | :------- | :----------------------------------------- |
| `success` | `bool`   | Whether the request was successful         |
| `message` | `string` | Human-readable message                     |
| `data`    | `any`    | Response payload (omitted if empty)        |
| `meta`    | `object` | Pagination metadata (omitted if empty)     |
| `errors`  | `array`  | Validation error details (omitted if null) |

### Validation Error Example

```json
{
  "success": false,
  "message": "Validation failed",
  "errors": [
    { "field": "email", "message": "Email must be a valid email address" },
    { "field": "password", "message": "Password must be at least 6 characters" }
  ]
}
```

---

## 📖 API Documentation (Swagger)

This template offers advanced Swagger multi-instance configuration.

**Generate Documentation:**
Use the provided scripts to regenerate Swagger docs whenever you update Godoc annotations:

- Windows: `scripts\swagger-gen.bat`
- Mac/Linux: `./scripts/swagger-gen.sh`

**Available UIs:**

- 🌐 **Global Swagger API**: `http://localhost:8080/swagger/index.html`
- 🔑 **Auth Module API**: `http://localhost:8080/swagger/auth/index.html`

---

## 🛣️ Included Built-in Endpoints

All endpoints are versioned under `/api/v1`.

| Method | Endpoint                | Auth Required |  Role   | Rate Limit | Description                 |
| :----- | :---------------------- | :-----------: | :-----: | :--------: | :-------------------------- |
| `GET`  | `/`                     |      ❌       |    -    |  100/min   | Root API status             |
| `GET`  | `/api/v1/auth/health`   |      ❌       |    -    |  100/min   | DB & Module Health check    |
| `POST` | `/api/v1/auth/register` |      ❌       |    -    |   5/min    | Sign up a new user          |
| `POST` | `/api/v1/auth/login`    |      ❌       |    -    |   5/min    | Authenticate & Fetch Tokens |
| `POST` | `/api/v1/auth/refresh`  |      ❌       |    -    |  100/min   | Refresh Access Token        |
| `POST` | `/api/v1/auth/logout`   |      ✅       |   Any   |  100/min   | Revoke Session              |
| `GET`  | `/api/v1/auth/profile`  |      ✅       |   Any   |  100/min   | Retrieve current User data  |
| `GET`  | `/api/v1/auth/users`    |      ✅       | `admin` |  100/min   | Admin: List all Users       |

---

## 🧰 Shared Utilities (`internal/utils/`)

### Validator (`validator.go`)

```go
// Add validate tags to your DTO structs
type CreateItemRequest struct {
    Name  string `json:"name" validate:"required,min=3,max=100"`
    Price float64 `json:"price" validate:"required,gte=0"`
    Email string `json:"email" validate:"required,email"`
}

// In your handler
if errs := utils.ValidateStruct(&req); errs != nil {
    return utils.ValidationErrorResponse(c, errs)
}
```

### Pagination (`pagination.go`)

```go
// In your handler
params := utils.ParsePagination(c)  // reads ?page=1&limit=10

// In your repository
var items []models.Item
var total int64

db.Model(&models.Item{}).Count(&total)
db.Scopes(utils.Paginate(params)).Find(&items)

meta := utils.CalculateMeta(params, total)
return utils.PaginatedResponse(c, "Items fetched", items, meta)
```

### API Response Helpers (`response.go`)

```go
// In your handler
// Success response with data
return utils.SuccessResponse(c, fiber.StatusOK, "User fetched", user)

// Message-only response
return utils.MessageResponse(c, fiber.StatusOK, "Logged out successfully")

// Error response
return utils.ErrorResponse(c, fiber.StatusNotFound, "User not found")
```

---

## 🛡️ License

Created and maintained by [Thun Saen](https://thunsaen.online/).
