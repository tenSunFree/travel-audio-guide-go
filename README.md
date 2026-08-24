# travel-audio-guide-go

[![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)
[![chi](https://img.shields.io/badge/Router-chi_v5-blue)](https://github.com/go-chi/chi)
[![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL-4169E1?logo=postgresql&logoColor=white)](https://www.postgresql.org)
[![sqlc](https://img.shields.io/badge/Codegen-sqlc-00ADD8)](https://sqlc.dev)
[![Docker](https://img.shields.io/badge/Container-Docker-2496ED?logo=docker&logoColor=white)](https://www.docker.com)
[![Architecture](https://img.shields.io/badge/Architecture-Feature--First-4CAF50)](#tech-stack)

---

## Introduction

Go backend for the travel audio guide system, providing RESTful API for user profile management with
Supabase Auth JWT verification (ES256/JWKS), built using chi, pgx, sqlc, and Docker. It also acts as
a compatibility proxy in front of the Taipei Travel open API, so the Flutter client can move its
attractions data through this backend without changing its response models.

This project is for learning and technical practice.

---

## Related App Client

This backend is part of a full-stack travel audio guide system and is designed to work together with
the Flutter app:

- [travel-audio-guide-flutter](https://github.com/tenSunFree/travel-audio-guide-flutter)

The Flutter app provides a cross-platform mobile client built with Flutter, Riverpod, Drift, Clean
Architecture, and Supabase Auth.
It handles user authentication via Supabase, retrieves the JWT access token, and calls this Go
backend API for profile management, user data operations, and attractions data.

---

## Preview

<p align="left">
  <img src="https://i.postimg.cc/KvZxVXzf/2026-06-17-044154.png" width="500"/>
</p> 
<p align="left">
  <img src="https://i.postimg.cc/MKWq4kTP/2026-06-17-044316.png" width="500"/>
  <img src="https://i.postimg.cc/Wb2TKR3Y/2026-06-17-044255.png" width="500"/>
</p> 

---

## Features

- Supabase Auth JWT verification (ES256/JWKS)
- `GET /api/v1/me` — fetch current user profile, auto-created on first login
- `PUT /api/v1/me` — partial update profile (only fields passed in are updated)
- `GET /open-api/{lang}/Attractions/All` — proxy to the Taipei Travel open API, response schema
  kept fully compatible with the upstream so the Flutter client only needs to change its base URL
- Swagger UI for interactive API documentation

---

## Tech Stack

- **Go 1.25**
  Statically typed, compiled language — high performance, simple deployment, minimal runtime
  overhead
- **chi v5**
  Lightweight HTTP router with composable middleware and route grouping, more flexible than standard
  `net/http`
- **pgx v5 + pgxpool**
  PostgreSQL driver with connection pooling via pgxpool, outperforms `database/sql` for concurrent
  workloads
- **sqlc**
  Generates type-safe Go code from SQL — no manual row scanning, SQL managed in `.sql` files
- **Supabase Auth + ES256/JWKS**
  JWTs are signed by Supabase using a private key (ES256); the backend fetches the public key from
  the JWKS endpoint for verification — more secure than symmetric HS256
- **slog (Go built-in)**
  Structured logger introduced in Go 1.21, outputs JSON format suitable for production observability
- **OpenAPI 3.0 + Swagger UI**
  API contract defined in `docs/openapi.yaml`, served via Docker as an interactive Swagger UI
- **Docker + Docker Compose**
  Multi-stage build (builder + alpine), Compose orchestrates PostgreSQL, Go API, and Swagger UI
- **Layered Architecture (Feature-first)**
  Each feature is grouped under `internal/<feature>/` with three internal layers:
  `handler` → `service` → `repository`.
  Each layer has a single responsibility — handlers never write SQL, services never touch HTTP,
  repositories never handle JWT.
- **Anti-Corruption Layer for third-party APIs**
  `internal/taipeitravel` holds the upstream client and raw DTOs; `internal/attractions` holds the
  app-facing DTO, mapper, service, and handler. If the upstream schema changes, only the client/DTO/
  mapper need updating — the contract exposed to Flutter stays stable.

---

## Security Design

- The client only sends `Authorization: Bearer <token>` — never a `user_id`
- The backend extracts `user_id` from JWT `claims.sub`
- All SQL `WHERE id = $1` conditions are derived from the JWT — client input is never trusted
- Supabase uses ES256 (asymmetric signing) — the backend holds only the public key; the private key
  never leaves Supabase
- `/open-api/*` routes are public tourism data and intentionally do not require authentication;
  `/api/v1/*` routes always require a valid Supabase JWT

---

## Prerequisites

- [Go 1.25+](https://go.dev/dl/)
- [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)
- A [Supabase](https://supabase.com/) project (used to obtain the JWKS URL or JWT Secret for auth
  verification)
- (Optional) [sqlc](https://docs.sqlc.dev/) — only needed if you plan to regenerate DB code from
  `.sql` files

---

## Environment

Copy `.env.example` to `.env` and fill in the values:

```bash
cp .env.example .env
```

| Variable                 | Required                                                        | Description                                                                         | Example                                                                       |
|--------------------------|------------------------------------------------------------------|--------------------------------------------------------------------------------------|--------------------------------------------------------------------------------|
| `APP_ENV`                | No                                                               | Application environment, defaults to `local`                                         | `local`                                                                         |
| `HTTP_ADDR`               | No                                                               | Address/port the HTTP server listens on, defaults to `:8080`                         | `:8080`                                                                         |
| `DATABASE_URL`            | **Yes**                                                          | PostgreSQL connection string                                                          | `postgres://user:password@localhost:5432/travel_audio_guide?sslmode=disable`   |
| `SUPABASE_JWKS_URL`       | One of `SUPABASE_JWKS_URL` / `SUPABASE_JWT_SECRET` is required  | Supabase JWKS endpoint, used for ES256 asymmetric verification (recommended)          | `https://<project-ref>.supabase.co/auth/v1/.well-known/jwks.json`              |
| `SUPABASE_JWT_SECRET`     | Same as above                                                    | Supabase JWT Secret, used for HS256 symmetric verification (fallback)                 | `your-supabase-jwt-secret`                                                     |
| `TAIPEI_TRAVEL_BASE_URL`  | No                                                               | Base URL of the upstream Taipei Travel open API, defaults to the official endpoint    | `https://www.travel.taipei/open-api`                                           |

> You can find `SUPABASE_JWKS_URL` / `SUPABASE_JWT_SECRET` in your Supabase project dashboard under
**Project Settings → API**.
> If both are provided, the backend prefers `SUPABASE_JWKS_URL` (ES256/JWKS mode).

---

## Getting Started

### Option 1: Run everything with Docker Compose (recommended)

Starts PostgreSQL, the Go API, and Swagger UI together:

```bash
make docker-up
```

Check logs:

```bash
make docker-logs
```

Stop everything:

```bash
make docker-down
```

### Option 2: Run locally with Go

Start only PostgreSQL via Docker:

```bash
make docker-db
```

Run the API directly on your machine:

```bash
make run
```

### Verify it's running

```bash
curl http://localhost:8080/healthz
# {"status":"ok"}
```

---

## API Documentation (Swagger UI)

The OpenAPI contract lives in [`docs/openapi.yaml`](./docs/openapi.yaml) and is served as an
interactive Swagger UI via Docker Compose.

After running `make docker-up`, open:

```
http://localhost:<SWAGGER_PORT>
```

> Replace `<SWAGGER_PORT>` with the port mapped to the `swagger-ui` service in `docker-compose.yml`.

---

## API Endpoints

### `GET /open-api/{lang}/Attractions/All`

Compatibility proxy for the [Taipei Travel open API](https://www.travel.taipei/open-api/swagger/docs/V1).
**No authentication required** — this is public tourism data.

The response schema is kept byte-for-byte compatible with the upstream API (no `success`/`data`
wrapper), so existing Flutter clients only need to change their base URL — no model changes
required.

```http
GET /open-api/zh-tw/Attractions/All?page=1
```

| Param         | In    | Description                                                       |
|---------------|-------|---------------------------------------------------------------------|
| `lang`        | path  | Language code: `zh-tw`, `zh-cn`, `en`, `ja`, `ko`                    |
| `categoryIds` | query | Comma-separated category IDs to filter by, e.g. `13,15`              |
| `nlat`        | query | Latitude (WGS84), used for nearby search                             |
| `elong`       | query | Longitude (WGS84), used for nearby search                            |
| `page`        | query | Page number, 30 results per page, defaults to `1`                    |

**200 OK**

```json
{
  "total": 471,
  "data": [
    {
      "id": 257,
      "name": "National Taiwan Museum (Natural History Branch)",
      "nlat": 25.04356,
      "elong": 121.51436
    }
  ]
}
```

**502 Bad Gateway** — returned when the upstream Taipei Travel API is unavailable or blocks the
request:

```json
{
  "error": "upstream service unavailable"
}
```

> Internally this endpoint calls the upstream API through an anti-corruption layer
> (`internal/taipeitravel` → `internal/attractions`): a raw upstream DTO is fetched by the client,
> mapped into an app-facing DTO, and returned unchanged. If the upstream schema changes in the
> future, only the client/DTO/mapper need updating.
>
> The upstream is behind Cloudflare, which blocks requests without a standard browser
> `User-Agent`. The client always sends one — see `internal/taipeitravel/client.go` if this ever
> needs adjusting.

---

All `/api/v1/*` routes require an `Authorization: Bearer <token>` header (a Supabase Auth JWT). The
`user_id` is always derived from the token — it is never accepted as client input.

### `GET /api/v1/me`

Fetch the current user's profile. The profile is auto-created on first login.

```http
GET /api/v1/me
Authorization: Bearer <token>
```

**200 OK**

```json
{
  "id": "uuid",
  "display_name": "string",
  "avatar_url": "string",
  "created_at": "2026-01-01T00:00:00Z",
  "updated_at": "2026-01-01T00:00:00Z"
}
```

### `PUT /api/v1/me`

Partially update the current user's profile. Only the fields included in the request body are
updated.

```http
PUT /api/v1/me
Authorization: Bearer <token>
Content-Type: application/json

{
  "display_name": "New Name"
}
```

**200 OK**

```json
{
  "id": "uuid",
  "display_name": "New Name",
  "avatar_url": "string",
  "created_at": "2026-01-01T00:00:00Z",
  "updated_at": "2026-01-01T00:00:00Z"
}
```

> See [`docs/openapi.yaml`](./docs/openapi.yaml) or the Swagger UI above for the full schema.

---

## Testing

```bash
make test
```

---

## Useful Make Commands

| Command            | Description                                              |
|--------------------|------------------------------------------------------------|
| `make run`         | Run the server locally (requires PostgreSQL running)       |
| `make build`       | Compile a static binary into `./bin/server`                |
| `make docker-up`   | Start PostgreSQL, API, and Swagger UI via Docker Compose    |
| `make docker-db`   | Start only PostgreSQL (for local development)               |
| `make docker-down` | Stop all Docker Compose services                             |
| `make docker-logs` | Tail API container logs                                      |
| `make sqlc-gen`    | Regenerate type-safe Go code from SQL (requires `sqlc`)      |
| `make tidy`        | Run `go mod tidy`                                             |
| `make test`        | Run all tests                                                 |

---

## Credits

This project is created for independent learning and demonstration purposes.
Special thanks to the original author for their open-source contribution.

---

## Notes

Image resources are for learning and purposes only. Please do not use them for commercial purposes.

If there is any infringement, please contact me for removal. Thank you.

---

## License

This repository is intended for learning and demonstration purposes.

If you plan to open-source this project, please add a `LICENSE` file (e.g. MIT) and confirm the
usage rights of any third-party assets before distribution.

---

## Project Structure

```
travel-audio-guide-go
├─ cmd
│  └─ server
│     └─ main.go
├─ docker-compose.yml
├─ Dockerfile
├─ docs
│  └─ openapi.yaml
├─ go.mod
├─ go.sum
├─ internal
│  ├─ attractions
│  │  ├─ dto.go
│  │  ├─ handler.go
│  │  ├─ mapper.go
│  │  ├─ repository.go
│  │  └─ service.go
│  ├─ auth
│  │  ├─ claims.go
│  │  ├─ context.go
│  │  └─ jwt_verifier.go
│  ├─ config
│  │  └─ config.go
│  ├─ database
│  │  └─ postgres.go
│  ├─ db
│  │  ├─ db.go
│  │  ├─ models.go
│  │  └─ profiles.sql.go
│  ├─ me
│  │  ├─ dto.go
│  │  ├─ handler.go
│  │  ├─ model.go
│  │  ├─ repository.go
│  │  └─ service.go
│  ├─ middleware
│  │  ├─ auth.go
│  │  ├─ cors.go
│  │  ├─ logger.go
│  │  └─ recovery.go
│  ├─ server
│  │  ├─ router.go
│  │  └─ server.go
│  └─ taipeitravel
│     ├─ client.go
│     └─ dto.go
├─ Makefile
├─ pkg
│  └─ response
│     └─ response.go
├─ README.md
├─ sql
│  ├─ queries
│  │  └─ profiles.sql
│  └─ schema
│     └─ 001_create_profiles.sql
└─ sqlc.yaml
```