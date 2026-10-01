# Payment Processing & Reconciliation System

A simulated fintech backend built with **Go, Gin, GORM, and PostgreSQL**. It supports payment requests, simulated provider outcomes, report imports, and reconciliation.

> Demo project only; no real money is processed.

## Features

- JWT authentication and role-aware middleware
- Payment creation, amount validation, and idempotency
- Simulated payment success/failure
- Database transactions with commit/rollback
- Provider report import and reconciliation
- Background workers, goroutines, and context cancellation
- Swagger API documentation
- Dockerized local setup

## Tech Stack

Go · Gin · GORM · PostgreSQL · JWT · Docker · Docker Compose · Swagger · Render · Neon

## Project Structure

```text
.
├── config/       # Database and app configuration
├── docs/         # Generated Swagger docs
├── dto/          # Request/response DTOs
├── handlers/     # HTTP handlers
├── middleware/   # Authentication middleware
├── models/       # Database models
├── routes/       # API routes
├── services/     # Business logic
├── workers/      # Background reconciliation workers
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── main.go
```

## API Endpoints

| Method | Endpoint | Purpose |
|---|---|---|
| GET | `/api/auth/` | Welcome |
| POST | `/api/auth/register` | Register |
| POST | `/api/auth/login` | Login / JWT |
| POST | `/api/payment/create` | Create payment |
| POST | `/api/transaction/{id}/{change}` | Simulate outcome (`success` / `fail`) |
| POST | `/api/provider/report` | Import provider report |
| POST | `/api/reconcil/{report_id}` | Start reconciliation |

Protected endpoints require a JWT. See Swagger for exact request/response schemas.

## Run with Docker

1. Clone the repository.
2. Create `.env` using the required variables (do not commit secrets).
3. Start the app:

```bash
docker compose up --build
```

API: `https://payment-reconciliatio-api.onrender.com/api/auth`  
Swagger: `https://payment-reconciliatio-api.onrender.com/swagger/index.html#/`

## How It Works

Users create payments, submit simulated provider outcomes, import provider reports, and start reconciliation. The service compares report records with internal data. GORM transactions help keep related database changes consistent; background workers process reconciliation jobs.

## Deployment

API: Render · Database: Neon PostgreSQL. Configure production secrets in the hosting provider's environment settings.

## Author

**Pavan Sonawane**
