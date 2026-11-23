# MyAsdos Backend

## Academic Portal Information System — Backend API

## Table of Contents

1. [Overview](#1-overview)
2. [Tech Stack](#2-tech-stack)
3. [Architecture](#3-architecture)
4. [Implemented Design Patterns](#4-implemented-design-patterns)
5. [Folder Structure](#5-folder-structure)
6. [How to Run](#6-how-to-run)
7. [Environment Variables](#7-environment-variables)
8. [API Response Format](#8-api-response-format)
9. [References](#9-references)

---

## 1. Overview

MyAsdos Backend is a REST API built with **Go (Golang)** using the **Gin** framework. This backend manages the entire administrative cycle of Teaching Assistants (Asdos) for Informatics Engineering at Institut Teknologi Sepuluh Nopember (ITS) — from recruitment, reporting, digital SK issuance, to honorarium disbursement.

This backend adopts **Layered Architecture** with a clear separation between:

- **Application Layer** — HTTP handlers, middleware, DTO
- **Domain Layer** — Business logic, entities, services, design patterns
- **Infrastructure Layer** — Database, repository, file storage, cache

---

## 2. Tech Stack

| Technology     | Version | Purpose                           |
| -------------- | ------- | --------------------------------- |
| **Go**         | 1.22+   | Backend programming language      |
| **Gin**        | 1.9+    | HTTP web framework                |
| **GORM**       | 1.25+   | ORM for PostgreSQL                |
| **PostgreSQL** | 15+     | Relational database               |
| **UUID**       | —       | Generate unique identifiers       |
| **JWT**        | —       | JSON Web Token for authentication |

---

## 3. Architecture

### Layered Architecture

```
┌─────────────────────────────────────────────┐
│  APPLICATION LAYER (internal/app/)          │
│  - HTTP Handlers (Gin Controllers)          │
│  - Middleware (Auth, CORS, Rate Limit)      │
│  - DTO (Request / Response structs)         │
└──────────────────┬──────────────────────────┘
                   │ depends on
                   ▼
┌─────────────────────────────────────────────┐
│  DOMAIN LAYER (internal/domain/)            │
│  - Entity (User, Application, Notification) │
│  - Service (Business Logic / Use Cases)     │
│  - Factory Pattern (Notification)           │
│  - State Pattern (Application Status)       │
└──────────────────┬──────────────────────────┘
                   │ depends on
                   ▼
┌─────────────────────────────────────────────┐
│  INFRASTRUCTURE LAYER (internal/infra/)     │
│  - Repository (GORM queries)                │
│  - Singleton (DB Connection)                │
│  - Storage (File Upload)                    │
│  - Cache (Redis)                            │
└─────────────────────────────────────────────┘
```

### Dependency Rules

> **Dependencies only flow downward.**  
> Application Layer → Domain Layer → Infrastructure Layer

- Domain Layer **does not know** about HTTP framework (Gin) or ORM (GORM)
- Domain Layer only knows **interfaces** it defines itself (e.g., `ApplicationRepository`)
- Infrastructure Layer provides **concrete implementations** of Domain Layer interfaces

---

## 4. Implemented Design Patterns

| #   | Pattern            | Category   | File                                                           | Description                                                                      |
| --- | ------------------ | ---------- | -------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| 1   | **Singleton**      | Creational | `internal/infrastructure/database/connection.go`               | Ensures a single DB connection pool instance, thread-safe via `sync.Once`        |
| 2   | **Factory Method** | Creational | `internal/domain/factory/notification_factory.go`              | Creates various notification types without lengthy `switch-case` blocks          |
| 3   | **Repository**     | Structural | `internal/infrastructure/repository/application_repository.go` | Abstracts GORM/PostgreSQL queries, facilitates testing & DB engine replacement   |
| 4   | **State**          | Behavioral | `internal/domain/state/application_state.go`                   | Encapsulates application status transition rules (PENDING → REVIEWED → ACCEPTED) |

---

## 5. Folder Structure

```
myasdos-backend/
├── cmd/
│   └── main.go                              # Application entry point
│
├── internal/                                # Private package (Go convention)
│   ├── app/                                 # APPLICATION LAYER
│   │   ├── handlers/                        # HTTP Controllers (Gin)
│   │   ├── middleware/                      # JWT, CORS, Rate Limit, RBAC
│   │   └── dto/                             # Data Transfer Objects
│   │
│   ├── domain/                              # DOMAIN LAYER
│   │   ├── entity/                          # Domain Models
│   │   │   ├── user.go
│   │   │   ├── application.go
│   │   │   ├── notification.go
│   │   │   └── job_posting.go
│   │   ├── service/                         # Business Logic / Use Cases
│   │   ├── factory/                         # FACTORY METHOD PATTERN
│   │   │   └── notification_factory.go
│   │   └── state/                           # STATE PATTERN
│   │       └── application_state.go
│   │
│   └── infrastructure/                      # INFRASTRUCTURE LAYER
│       ├── database/
│       │   └── connection.go                # SINGLETON PATTERN
│       ├── repository/                      # REPOSITORY PATTERN
│       │   ├── application_repository.go
│       │   ├── user_repository.go
│       │   └── job_repository.go
│       ├── storage/                         # File Storage (MinIO/S3)
│       └── cache/                           # Redis Client
│
├── go.mod
├── go.sum
└── README.md                                # This documentation
```

### Folder Descriptions

| Folder                                | Purpose                                                                 |
| ------------------------------------- | ----------------------------------------------------------------------- |
| `cmd/`                                | Application entry point. Contains only `main.go`.                       |
| `internal/app/handlers/`              | Gin HTTP handlers. Receive requests, call domain services, return DTOs. |
| `internal/app/middleware/`            | Middleware chains: JWT validation, CORS, rate limiting, RBAC.           |
| `internal/app/dto/`                   | Request/Response structs. Separate API format from database entities.   |
| `internal/domain/entity/`             | Pure domain models. Structs with minimal fields and methods.            |
| `internal/domain/service/`            | Business logic. Orchestrates repository, state, and factory.            |
| `internal/domain/factory/`            | Factory Method pattern. Creates notifications based on event type.      |
| `internal/domain/state/`              | State pattern. State machine for application status.                    |
| `internal/infrastructure/database/`   | Singleton DB connection manager.                                        |
| `internal/infrastructure/repository/` | Repository implementations. GORM queries to PostgreSQL.                 |
| `internal/infrastructure/storage/`    | File upload/download abstraction (S3/MinIO/Local).                      |
| `internal/infrastructure/cache/`      | Redis client abstraction.                                               |

---

## 6. How to Run

### Prerequisites

- Go 1.22+ installed
- PostgreSQL 15+ (optional for demo — uses Mock Repository)

### Step 1: Clone & Install Dependencies

```bash
cd myasdos-backend
go mod tidy
```

### Step 2: Setup Environment Variables (Optional)

If you want to use PostgreSQL, create a `.env` file:

```bash
export DB_HOST=localhost
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=myasdos
export DB_PORT=5432
```

If not set, the application will use **Mock Repository** (in-memory) for demonstration.

### Step 3: Build

```bash
go build ./cmd/main.go
```

### Step 4: Run

```bash
go run ./cmd/main.go
```

**Expected output:**

```
=== MyAsdos Backend - Design Pattern Demonstration ===

[Singleton] Initializing database connection...
Note: Database not configured. Using mock repositories for demo.

[Repository] Creating ApplicationRepository...
Application created: ID=33333333-3333-3333-3333-333333333333, Status=PENDING

[State] Initializing Application State Machine...
Initial status: PENDING
Action: Lecturer reviews the application...
Status after review: REVIEWED
Action: Lecturer decides ACCEPTED...
Status after decision: ACCEPTED
Action: Student tries to withdraw (should fail)...
Expected error: cannot withdraw application in ACCEPTED state

[Factory Method] Creating notifications...
Notification 1 created:
  Title: Lamaran Diterima
  Message: Selamat! Lamaran kamu untuk asdos Pemrograman Web telah diterima oleh Dr. Budi Santoso.
  Type: APPLICATION_ACCEPTED
...
=== All 4 Design Patterns demonstrated successfully ===
```

---

## 7. Environment Variables

| Variable      | Default          | Description                |
| ------------- | ---------------- | -------------------------- |
| `DB_HOST`     | `localhost`      | PostgreSQL host            |
| `DB_USER`     | `postgres`       | PostgreSQL username        |
| `DB_PASSWORD` | `postgres`       | PostgreSQL password        |
| `DB_NAME`     | `myasdos`        | Database name              |
| `DB_PORT`     | `5432`           | PostgreSQL port            |
| `JWT_SECRET`  | —                | Secret key for JWT signing |
| `REDIS_HOST`  | `localhost:6379` | Redis host                 |

---

## 8. API Response Format

All API endpoints return responses with a consistent format:

```json
{
  "success": true,
  "message": "Operation successful",
  "data": { ... },
  "meta": {
    "page": 1,
    "limit": 20,
    "total": 100
  }
}
```

### Error Response

```json
{
  "success": false,
  "message": "Anda sudah mendaftar di lowongan ini",
  "code": "DUPLICATE_APPLICATION"
}
```
