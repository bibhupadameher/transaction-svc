# Transaction Service

This project is a **transaction microservice** built with **Golang** , following clean architecture principles.  
It provides APIs for creating accounts, retrieving account information, and processing transactions.  

---

## 📂 Project Structure

- **docker-compose.yml** → Docker Compose file for local development (Postgres, PGAdmin, service).  
- **tx-api/** → Main Go application source code.  
  - **config/** → Loading Enviorment based Config
  - **constant/** → Storing Application level constants
  - **core/**
    - **aperror/** → Custom Error to maintain apllication level error 
    - **kit/** → Common Go-kit helpers (error encoder, response encoder, etc.).  
    - **logging/** → Centralized structured logging (zap-based), trace IDs.  
    - **postgres/** → Low level functions of GO ORM  
  - **dao/** → Data Access Layer (DB operations).  
  - **dto/** → Request/Response DTOs for accounts and transactions (with validation).  
  - **endpoint/** → Business endpoints and middleware.  
  - **errors/** → Centralized error handling and custom error definitions.  
  - **http/** → HTTP handlers, decoders/encoders, routing with Gorilla Mux.  
  - **service/** → Business logic implementation (transaction service handler).
  - **config.yaml** → Application configuration file 
  - **Dockerfile** → Docker file for Go app
  - **enum.yaml** → Enum List file 
  - **main.go** → Application entrypoint (logging, service, endpoints, HTTP server).
  - **go.mod** → GO Mod file
- **pgadmin/** → Configurations data for PGAdmin (database UI).  
- **postman/** → Postman collections and environment files for API testing.  

---

## 📦 Packages Overview

### `config`
- Handles application configuration loading and access.
- Supports multiple environments (`local`, `prod`).
- Key functions:
  - `Load()` – load configuration from `config.yaml`.
  - `Get()` – retrieve loaded configuration.



### `constant`
- Defines application-wide constants and enums.
- Key constants:
  - `UNIQUE_VIOLATION`, `FOREIGN_KEY_VIOLATION` – common DB constraints.



### `core/postgres`
- Database service layer using GORM for PostgreSQL.
- Handles schema creation, migrations, and batch writes.
- Key types/functions:
  - `DBService` – interface for DB operations.
  - `MigrateTables(tables ...interface{})` – auto-migrate tables.
  - `FindRows`, `FindFirst`, `BatchWriteData` – generic DB operations.


### `core/apperror`
- Structured application errors with codes, messages, types, and details.
- Key types/functions:
  - `AppError` – structured error object.
  - `New(code, type, message, httpStatus, details...)` – create new error.

### `core/logging`
- Provides structured logging with **zap**.
- Attaches `traceID` to each request for debugging.

### `core/kit`
- Standard utilities for **error encoding** and **generic response encoding**.



### `dto`
- Data transfer objects for:
  - CreateAccountRequest / Response
  - GetAccountRequest / Response
  - CreateTransactionRequest / Response
- Each DTO includes **validation logic** using `govalidator`.

### `errors`
- Helper package to create application-level errors.
- Key functions:
  - `NewErrFieldMissingorInvalid(field string)`
  - `NewErrDatabaseError()`

### `dao`
- Data access layer for `Account` and `Transaction` models.
- Uses `postgres.DBService` for database interactions.
- Handles DB-specific errors (unique violations, foreign key violations).
- Key types/functions:
  - `SaveAccount`, `GetAccountByID`, `SaveTransaction` – CRUD operations.







### `service`
- Implements the business logic for transactions and accounts.
- Uses DAO layer (`TransactionDAOInterface`) to interact with the database.
- Key functions:
  - `CreateAccount(ctx, request)` – validates and saves a new account.
  - `GetAccount(ctx, request)` – retrieves account by ID.
  - `CreateTransaction(ctx, request)` – validates and saves a new transaction.




### `endpoint`
- Pure business logic entrypoints for Go-kit.
- Implements:
  - `MakeCreateAccountEndpoint`
  - `MakeGetAccountEndpoint`
  - `MakeCreateTransactionEndpoint`
- Middleware:
  - `LoggingMiddleware` – traces execution and logs request duration.



### `http`
- Handles HTTP transport layer using Go-kit and Gorilla Mux.
- Encodes/decodes requests and responses between HTTP and service endpoints.
- Key functions:
  - `decodeCreateAccountRequest`
  - `decodeGetAccountRequest`
  - `decodeCreateTransactionRequest`
  - `NewHTTPHandler` – sets up router and HTTP endpoints.






---

## 🚀 Running the Service


#### Run Services
```bash
docker-compose up --build -d

```

- Service will run on **localhost:8080**

- Postgres DB will be available on configured port.

- PGAdmin UI available at **localhost:5050** [Check docker-compose.yml for credentials]

- Import the postman collection and enviroment and test the apis

#### Stop Services
```bash
docker-compose down

```
---
## 🛠️ Devlopment

This section is only for devlopment mode, where we will use database and pgadmin in docker but run the GO app from laptop
#### Start Database
```bash
docker compose up -d postgres pgadmin

```
#### Start Go app
```bash
cd .\tx-api\
go mod tidy
go run main.go
```

#### Unit Testing
```bash
go test ./...
```