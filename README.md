# travel-audio-guide-go

[![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)
[![chi](https://img.shields.io/badge/Router-chi_v5-blue)](https://github.com/go-chi/chi)
[![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL-4169E1?logo=postgresql&logoColor=white)](https://www.postgresql.org)
[![sqlc](https://img.shields.io/badge/Codegen-sqlc-00ADD8)](https://sqlc.dev)
[![Docker](https://img.shields.io/badge/Container-Docker-2496ED?logo=docker&logoColor=white)](https://www.docker.com)
[![Auth](https://img.shields.io/badge/Auth-Supabase_JWT-3FCF8E?logo=supabase&logoColor=white)](#security-design)
[![Testing](https://img.shields.io/badge/Testing-Go_test-FF9800)](#testing)
[![Deploy](https://img.shields.io/badge/Deploy-Cloud_Run-4285F4?logo=googlecloud&logoColor=white)](#deployment-google-cloud)
[![Database](https://img.shields.io/badge/Cloud_SQL-PostgreSQL_16-4285F4?logo=googlecloud&logoColor=white)](#deployment-google-cloud)
[![Secrets](https://img.shields.io/badge/Secrets-Secret_Manager-4285F4?logo=googlecloud&logoColor=white)](#deployment-google-cloud)
[![OpenAPI](https://img.shields.io/badge/API-OpenAPI_3.0-6BA539?logo=openapiinitiative&logoColor=white)](./docs/openapi.yaml)
[![Architecture](https://img.shields.io/badge/Architecture-Feature--First-4CAF50)](#tech-stack)

---

## Introduction

Go backend for the travel audio guide system, providing RESTful API for user profile management with
Supabase Auth JWT verification (ES256/JWKS), built using chi, pgx, sqlc, and Docker. It also acts as
a compatibility proxy in front of the Taipei Travel open API, so the Flutter client can move its
attractions, events, media, tours, and category data through this backend without changing its
response models.

This project is for learning and technical practice.

---

## Related App Client

This backend is part of a full-stack travel audio guide system and is designed to work together with
the Flutter app:

- [travel-audio-guide-flutter](https://github.com/tenSunFree/travel-audio-guide-flutter)

The Flutter app provides a cross-platform mobile client built with Flutter, Riverpod, Drift, Clean
Architecture, and Supabase Auth.
It handles user authentication via Supabase, retrieves the JWT access token, and calls this Go
backend API for profile management, user data operations, and Taipei Travel data (attractions,
news, activities, event calendar, audio guides, themed tours, and category lookups).

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
  kept compatible with the upstream for the documented fields and JSON types, so the Flutter
  client only needs to change its base URL
- `GET /open-api/{lang}/Events/News` — proxy for Taipei Travel tourism news
- `GET /open-api/{lang}/Events/Activity` — proxy for Taipei Travel activities and exhibitions
- `GET /open-api/{lang}/Events/Calendar` — proxy for Taipei Travel event calendar entries
- `GET /open-api/{lang}/Media/Audio` — proxy for Taipei Travel audio guide entries
- `GET /open-api/{lang}/Tours/Theme` — proxy for Taipei Travel themed tours
- `GET /open-api/{lang}/Miscellaneous/Categories` — proxy for category lookups by resource type
- `GET /health` / `GET /ready` — liveness and readiness checks (see [Health Checks](#health-checks))
- Swagger UI for interactive API documentation
- Deployed on Google Cloud Run with Cloud SQL and Secret Manager (see
  [Deployment (Google Cloud)](#deployment-google-cloud))
- Unit tests for configuration loading and HTTP routing (see [Testing](#testing))

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
- **Google Cloud Run + Cloud SQL + Secret Manager**
  Stateless container deployment. The runtime service account is separate from any deploy identity;
  `DATABASE_URL` comes from Secret Manager (pinned to a secret version), and Cloud SQL is reached
  over the Cloud SQL Unix socket.
- **Go `testing` + `httptest`**
  Standard-library tests, no extra framework. Router tests run the real `NewRouter` against
  `httptest.NewRecorder`.
- **Layered Architecture (Feature-first)**
  Each feature is grouped under `internal/<feature>/` with three internal layers:
  `handler` → `service` → `repository`.
  Each layer has a single responsibility — handlers never write SQL, services never touch HTTP,
  repositories never handle JWT.
- **Anti-Corruption Layer for third-party APIs**
  `internal/taipeitravel` holds a single shared HTTP client (`getJSON`) plus the upstream raw DTOs
  for every Taipei Travel endpoint; `internal/attractions`, `internal/events`, `internal/media`,
  `internal/tours`, and `internal/miscellaneous` each hold their own app-facing DTO, mapper,
  service, and handler. If the upstream schema changes, only the client/DTO/mapper need updating
  — the contract exposed to Flutter stays stable.

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

| Variable                 | Required                                                        | Description                                                                                                                             | Example                                                                       |
|--------------------------|------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------|
| `APP_ENV`                | No                                                               | Application environment, defaults to `local`                                                                                              | `local`                                                                         |
| `HTTP_ADDR`               | No                                                               | Address/port the HTTP server listens on, defaults to `:8080`. If the `PORT` env var is set (e.g. on Cloud Run), `PORT` takes precedence over `HTTP_ADDR`. | `:8080`                                                                         |
| `PORT`                   | No                                                               | Injected automatically by Cloud Run (and similar platforms) — not meant to be set manually in local development. When present, it overrides `HTTP_ADDR`. | `8080`                                                                          |
| `DATABASE_URL`            | **Yes**                                                          | PostgreSQL connection string                                                                                                              | `postgres://user:password@localhost:5432/travel_audio_guide?sslmode=disable`   |
| `SUPABASE_JWKS_URL`       | One of `SUPABASE_JWKS_URL` / `SUPABASE_JWT_SECRET` is required  | Supabase JWKS endpoint, used for ES256 asymmetric verification (recommended)                                                               | `https://<project-ref>.supabase.co/auth/v1/.well-known/jwks.json`              |
| `SUPABASE_JWT_SECRET`     | Same as above                                                    | Supabase JWT Secret, used for HS256 symmetric verification (fallback)                                                                      | `your-supabase-jwt-secret`                                                     |
| `TAIPEI_TRAVEL_BASE_URL`  | No                                                               | Base URL of the upstream Taipei Travel open API, defaults to the official endpoint                                                         | `https://www.travel.taipei/open-api`                                           |

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

On success, this also prints the ready-to-use Swagger UI and health check URLs.

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
curl http://localhost:8080/health
# {"success":true,"data":{"status":"ok"}}

curl http://localhost:8080/ready
# {"success":true,"data":{"status":"ready"}}
```

---

## Health Checks

Two endpoints are exposed for operational checks, and they answer different questions:

| Endpoint    | Checks                                   | Typical use                                                                 |
|-------------|-------------------------------------------|-------------------------------------------------------------------------------|
| `GET /health` | The Go process / HTTP server is running | Liveness — "is the container alive at all?"                                   |
| `GET /ready` | The above, plus a live ping to the PostgreSQL connection pool | Readiness — "is the service actually able to serve real traffic?" |

```bash
curl http://localhost:8080/health
# {"success":true,"data":{"status":"ok"}}

curl http://localhost:8080/ready
# 200 → {"success":true,"data":{"status":"ready"}}
# 503 → {"success":false,"error":"database not ready"}
```

`/ready` returns `503` if the database is unreachable, instead of `200`. Neither endpoint requires
authentication. These endpoints exist primarily for container orchestration platforms (e.g. Cloud
Run health checks, Kubernetes probes) — they are not part of the versioned `/api/v1` or `/open-api`
contract.

---

## API Documentation (Swagger UI)

The OpenAPI contract lives in [`docs/openapi.yaml`](./docs/openapi.yaml) and is served as an
interactive Swagger UI via Docker Compose.

After running `make docker-up`, the terminal prints the ready URLs, or you can open it directly:

```text
http://localhost:8081/
```

> This matches the port mapped to the `swagger-ui` service in `docker-compose.yml`. If you've
> changed that port mapping, use your updated port instead.

---

## API Endpoints

### `GET /open-api/{lang}/Attractions/All`

Compatibility proxy for the [Taipei Travel open API](https://www.travel.taipei/open-api/swagger/docs/V1).
**No authentication required** — this is public tourism data.

The response schema and client-facing contract are kept compatible with the upstream API (no
`success`/`data` wrapper), so existing Flutter clients only need to change their base URL — no
model changes required.

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

### `GET /open-api/{lang}/Events/News`

Compatibility proxy for Taipei Travel tourism news. **No authentication required.**

```http
GET /open-api/zh-tw/Events/News?page=1
```

| Param   | In    | Description                                        |
|---------|-------|-----------------------------------------------------|
| `lang`  | path  | Language code: `zh-tw`, `zh-cn`, `en`, `ja`, `ko`    |
| `begin` | query | Start date, format `yyyy-MM-dd`                      |
| `end`   | query | End date, format `yyyy-MM-dd`                        |
| `page`  | query | Page number, 30 results per page, defaults to `1`    |

**200 OK**

```json
{
  "total": 35,
  "data": [
    {
      "id": 45417,
      "title": "水門何時關？北市府設多管道　民眾防汛資訊不漏接",
      "begin": null,
      "end": null,
      "posted": "2023-04-04 09:06:00 +08:00"
    }
  ]
}
```

> `begin` and `end` are nullable and stay `null` when the upstream does not provide a date range —
> they are never coerced into empty strings.

---

### `GET /open-api/{lang}/Events/Activity`

Compatibility proxy for Taipei Travel activities and exhibitions. **No authentication required.**

```http
GET /open-api/zh-tw/Events/Activity?page=1
```

| Param   | In    | Description                                                                         |
|---------|-------|---------------------------------------------------------------------------------------|
| `lang`  | path  | Language code: `zh-tw`, `zh-cn`, `en`, `ja`, `ko`, `es`, `id`, `th`, `vi`               |
| `begin` | query | Start date, format `yyyy-MM-dd`                                                        |
| `end`   | query | End date, format `yyyy-MM-dd`                                                          |
| `page`  | query | Page number, 30 results per page, defaults to `1`                                      |

**200 OK**

```json
{
  "total": 24,
  "data": [
    {
      "id": 67755,
      "title": "存在的輕聲 ─ 對話 吳偉谷陶塑個展",
      "distric": "",
      "nlat": "25.0346",
      "elong": "121.522",
      "co_rganizer": ""
    }
  ]
}
```

> **Type note:** unlike Attractions, `nlat` and `elong` here are returned as **strings** by the
> upstream API, not numbers — this is preserved as-is. The upstream field is also spelled
> `co_rganizer` (not `co_organizer`); this is kept unchanged for compatibility.

---

### `GET /open-api/{lang}/Events/Calendar`

Compatibility proxy for the Taipei Travel event calendar. **No authentication required.**

```http
GET /open-api/zh-tw/Events/Calendar?page=1
```

| Param        | In    | Description                                        |
|--------------|-------|-----------------------------------------------------|
| `lang`       | path  | Language code: `zh-tw`, `zh-cn`, `en` (only three)   |
| `categoryId` | query | Category ID                                          |
| `begin`      | query | Start date, format `yyyy-MM-dd`                      |
| `end`        | query | End date, format `yyyy-MM-dd`                        |
| `page`       | query | Page number, 30 results per page, defaults to `1`    |

**200 OK**

```json
{
  "total": 23,
  "data": [
    {
      "id": 66376,
      "title": "2026臺北夜市打牙祭",
      "nlat": "25.0886",
      "elong": "121.524",
      "is_major": false
    }
  ]
}
```

> This endpoint only supports three languages (`zh-tw`, `zh-cn`, `en`) — a wider language code such
> as `ja` or `ko` returns `400`, even though those are valid for Attractions and Events/Activity.

---

### `GET /open-api/{lang}/Media/Audio`

Compatibility proxy for Taipei Travel audio guide entries. **No authentication required.**

```http
GET /open-api/zh-tw/Media/Audio?page=1
```

| Param  | In    | Description                                        |
|--------|-------|-----------------------------------------------------|
| `lang` | path  | Language code: `zh-tw`, `zh-cn`, `en`, `ja`, `ko`    |
| `page` | query | Page number, 30 results per page, defaults to `1`    |

**200 OK**

```json
{
  "total": 140,
  "data": [
    {
      "id": 28,
      "title": "北投圖書館",
      "summary": null,
      "url": "https://www.travel.taipei/audio/28",
      "file_ext": null,
      "modified": "2025-12-10 15:55:41 +08:00"
    }
  ]
}
```

---

### `GET /open-api/{lang}/Tours/Theme`

Compatibility proxy for Taipei Travel themed tours. **No authentication required.**

```http
GET /open-api/zh-tw/Tours/Theme?categoryIds=532&page=1
```

| Param         | In    | Description                                                                         |
|---------------|-------|---------------------------------------------------------------------------------------|
| `lang`        | path  | Language code: `zh-tw`, `zh-cn`, `en`, `ja`, `ko`                                       |
| `categoryIds` | query | Comma-separated category IDs, e.g. `532,228`. Look these up via `Miscellaneous/Categories?type=Tours` |
| `page`        | query | Page number, 30 results per page, defaults to `1`                                      |

**200 OK**

```json
{
  "total": 87,
  "data": [
    {
      "id": 1409,
      "seasons": ["1", "2", "3", "4"],
      "months": ["01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"],
      "days": 1,
      "title": "鐵道觀光主題遊程",
      "category": null,
      "transport": null,
      "users": null
    }
  ]
}
```

> **Schema note:** `category`, `transport`, and `users` are always `null` in every observed
> sample. Their non-null shape is undocumented by the upstream, so these fields are passed
> through as-is (`json.RawMessage` on the Go side) rather than guessed at.

---

### `GET /open-api/{lang}/Miscellaneous/Categories`

Compatibility proxy for category lookups by resource type. **No authentication required.**
Unlike other proxy endpoints, the `type` query parameter is **required**; this backend validates
it before forwarding the request upstream.

```http
GET /open-api/zh-tw/Miscellaneous/Categories?type=Tours
```

| Param  | In    | Description                                                                                    |
|--------|-------|--------------------------------------------------------------------------------------------------|
| `lang` | path  | Language code: `zh-tw`, `zh-cn`, `en`, `ja`, `ko`                                                  |
| `type` | query | **Required.** One of `Activity`, `Calendar`, `Pictorial`, `Attractions`, `Accommodation`, `Tours` |

**200 OK**

```json
{
  "total": 20,
  "data": {
    "Category": [
      { "id": 532, "name": "樂遊臺北" },
      { "id": 228, "name": "經典遊程" }
    ]
  }
}
```

**400 Bad Request** — returned when `type` is missing or not one of the allowed values:

```json
{
  "error": "invalid parameter: type ()"
}
```

> **Schema note:** the `Category` wrapper key inside `data` has only been confirmed for
> `type=Tours`. If other `type` values turn out to use a different key, the DTO will need
> per-type handling — see `internal/taipeitravel/dto.go`.

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
make test        # or: go test ./...
```

Current coverage is intentionally small and focused:

| Package           | What is tested                                                                                  |
|-------------------|---------------------------------------------------------------------------------------------------|
| `internal/config` | `PORT` takes precedence over `HTTP_ADDR` (default `:8080`); `Load()` fails without `DATABASE_URL` or a Supabase verifier setting; defaults for `APP_ENV` and the Taipei Travel base URL |
| `internal/server` | `GET /health` returns 200; legacy `/healthz` and `/readyz` stay unregistered (404); `/api/v1/me` returns 401 for a missing, non-Bearer, or empty token; unknown routes return a JSON 404 |

The router tests build the real router with `nil` dependencies, so they never need a database, a
JWKS endpoint, or network access. `/ready` is not covered yet because it needs a real connection
pool; it can be tested later by extracting a small interface that exposes only `Ping`.

Before opening a pull request:

```bash
gofmt -l .        # should print nothing
go vet ./...
go test ./...
go build ./...
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

## Deployment (Google Cloud)

The API runs on Google Cloud in the `asia-east1` region:

| Component         | Role                                                                                     |
|-------------------|-------------------------------------------------------------------------------------------|
| Cloud Run         | Runs the container. Listens on the `PORT` injected by Cloud Run                           |
| Artifact Registry | Stores images, tagged with the Git commit SHA (`api:<git-sha>`)                           |
| Cloud SQL         | PostgreSQL 16, reached through the Cloud SQL Unix socket; deletion protection, automated backups and point-in-time recovery enabled |
| Secret Manager    | Holds `DATABASE_URL`, pinned to a secret version (not `latest`)                           |
| Service accounts  | A dedicated runtime service account with only Secret Accessor and Cloud SQL Client roles  |

`SUPABASE_JWKS_URL` is a plain environment variable. Production uses JWKS-only verification, so no
JWT secret is stored on the service.

Health endpoints are `/health` and `/ready` rather than `/healthz` and `/readyz`, because Cloud Run
reserves some URL paths ending in `z` (see Cloud Run "Known issues → Reserved URL paths").

Manual deploy of a new image:

```bash
IMAGE=asia-east1-docker.pkg.dev/<project-id>/travel-audio-guide/api:$(git rev-parse --short HEAD)
docker build -t "$IMAGE" .
docker push "$IMAGE"
gcloud run deploy travel-audio-guide-api --image="$IMAGE" --region=asia-east1
```

Only `--image` and `--region` are passed on redeploy; the service account, secret, Cloud SQL
connection and environment variables carry over from the previous revision.

Smoke test after a deploy:

```bash
curl -i "$SERVICE_URL/health"   # 200
curl -i "$SERVICE_URL/ready"    # 200 when Cloud SQL is reachable, 503 otherwise
curl -i "$SERVICE_URL/api/v1/me" # 401 without a token
```

Rollback by sending traffic back to a previous revision:

```bash
gcloud run services update-traffic travel-audio-guide-api --region=asia-east1 \
  --to-revisions=<previous-revision>=100
```

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
│  │  ├─ config.go
│  │  └─ config_test.go
│  ├─ database
│  │  └─ postgres.go
│  ├─ db
│  │  ├─ db.go
│  │  ├─ models.go
│  │  └─ profiles.sql.go
│  ├─ events
│  │  ├─ dto.go
│  │  ├─ handler.go
│  │  ├─ mapper.go
│  │  ├─ repository.go
│  │  └─ service.go
│  ├─ media
│  │  ├─ dto.go
│  │  ├─ handler.go
│  │  ├─ mapper.go
│  │  ├─ repository.go
│  │  └─ service.go
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
│  ├─ miscellaneous
│  │  ├─ dto.go
│  │  ├─ handler.go
│  │  ├─ mapper.go
│  │  ├─ repository.go
│  │  └─ service.go
│  ├─ server
│  │  ├─ router.go
│  │  ├─ router_test.go
│  │  └─ server.go
│  ├─ taipeitravel
│  │  ├─ client.go
│  │  └─ dto.go
│  └─ tours
│     ├─ dto.go
│     ├─ handler.go
│     ├─ mapper.go
│     ├─ repository.go
│     └─ service.go
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