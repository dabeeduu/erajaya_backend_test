# Erajaya Backend Test

[![Go Version](https://img.shields.io/badge/Go-1.24.5-blue)](https://golang.org)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-blue)](https://www.postgresql.org/)
[![Redis](https://img.shields.io/badge/Redis-7-red)](https://redis.io/)
[![Docker](https://img.shields.io/badge/Docker-Compose-blue)](https://docs.docker.com/compose/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A simple product API built in Go for Erajaya backend test, using PostgreSQL, Redis, and Docker.

## Architecture Overview

```text
                        +-----------------+
                        |     Client      |
                        +--------+--------+
                                 |
                                 v
                        +--------+--------+
                        |   Go Application |
                        |  (Gin, Clean Arch) |
                        +--------+--------+
                                 |
                 +---------------+----------------+
                 |                                |
                 v                                v
        +--------+--------+              +--------+--------+
        |    PostgreSQL    |              |      Redis      |
        |   (Persistent DB)|              | (Caching Layer) |
        +------------------+              +-----------------+
```


*Built with the tools and technologies:*

![Markdown](https://img.shields.io/badge/Markdown-000000.svg?style=flat&logo=Markdown&logoColor=white)
![Redis](https://img.shields.io/badge/Redis-FF4438.svg?style=flat&logo=Redis&logoColor=white)
![Go](https://img.shields.io/badge/Go-00ADD8.svg?style=flat&logo=Go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-008ECF.svg?style=flat&logo=Gin&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED.svg?style=flat&logo=Docker&logoColor=white)
![YAML](https://img.shields.io/badge/YAML-CB171E.svg?style=flat&logo=YAML&logoColor=white)


---

## Overview

**Erajaya_backend_test** is a backend service built in **Go**, developed as part of the technical test for a position at **Erajaya**. The project uses a clean architecture style with Redis caching, PostgreSQL database integration, and is fully containerized using Docker.

## 📁 Project Structure


```
.
├── app/                        # Application bootstrap logic
├── config/                     # Environment config loader
├── database/                   # PostgreSQL and Redis setup
├── internal/                   # Main business logic
│   ├── cache/                  # Redis cache operations
│   ├── dto/                    # Request/response DTOs
│   ├── entity/                 # Domain entities
│   ├── handler/                # HTTP route handlers
│   ├── repository/             # Data persistence layer
│   └── usecase/                # Business use cases
├── logger/                     # Central logging utility
├── middleware/                 # Custom middlewares (error, logger)
├── migrations/                 # SQL schema definitions
├── utils/                      # Utilities and helpers
├── .env.example                # Example environment variables
├── Dockerfile                  # App Docker build
├── docker-compose.yaml         # Docker Compose config
├── go.mod, go.sum              # Go dependencies
├── main.go                     # App entry point
├── LICENSE
└── README.md
```


## ⚙️ Tech Stack

- **Language**: Go 1.24.5
- **Web Framework**: `net/http` (standard library)
- **Database**: PostgreSQL 15 (via Docker)
- **Cache**: Redis 7 (via Docker)
- **Migration Tool**: [migrate/migrate](https://hub.docker.com/r/migrate/migrate)
- **Containerization**: Docker & Docker Compose
- **Service Health**: PostgreSQL health check using `pg_isready`
- **Environment Management**: `.env` file support via Docker Compose
- **Auto Migration**: Migration container auto-applies SQL files at startup

## 🚀 Features

- RESTful Product CRUD API
- Redis caching for response optimization
- Centralized error and logging middleware
- Graceful shutdown with timeout
- SQL migration for schema version control
- Modular architecture that's easy to maintain and scale

## 🧱 Design and Architecture

The project uses a **Clean Architecture** approach with a **layered** structure that is divided into several main layers:
### 📌 Layering:

```
Handler (Controller)
    ↓
Usecase (Business Logic)
    ↓
Repository (Data Access)
```
### ✨Reasons for Design Selection:

- **Separation of Concerns (SoC)**: Each layer has its own responsibilities, making code easier to read and maintain.
- **Testability**: Since the business logic is separated from the transport and database layers, unit tests can be performed more easily.
- **Scalability & Extensibility**: Adding new features does not break the old code.
- **Flexibility**: For example, replacing Redis with another cache system or changing the database, just change it in the repository layer.

## 🔐 Environment Configuration

Use `.env.example` as a template:

```env
APP_PORT=8000
DATABASE_URL=postgres://user:pass@postgres:5432/mydb?sslmode=disable
REDIS_ADDR=redis:6379
SHUTDOWN_TIME_SEC=3
```
## 📦 Getting Started

### 1. Clone Repository

```bash
git clone https://github.com/your-username/erajaya_backend_test.git
cd erajaya_backend_test
```

### 2. Create`.env` file

```bash
cp .env.example .env
```

### 3. Run with Docker

```bash
docker compose up --build
```

API is accessible at: [http://localhost:8000](http://localhost:8000)

---

## 🧪 API Endpoints

| Method | Endpoint               | Description           |
|--------|------------------------|---------------------|
| GET    | `/api/products`        | List all product  |
| POST   | `/api/products`        | Add new product  |

### 🔎 Dokumentasi API (Postman)

👉 [API Documentation (Postman)](https://documenter.getpostman.com/view/46907087/2sB34kFfCU#a6d1b779-eed1-404f-b0c3-6b990be04159)
