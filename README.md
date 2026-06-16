# travel-audio-guide-go

---

## Introduction

Go backend for the travel audio guide system, providing RESTful API for user profile management with Supabase Auth JWT verification (ES256/JWKS), built using chi, pgx, sqlc, and Docker.

This project is for learning and technical practice.

---

## Related App Client

This backend is part of a full-stack travel audio guide system and is designed to work together with the Flutter app:

- [travel-audio-guide-flutter](https://github.com/tenSunFree/travel-audio-guide-flutter)

The Flutter app provides a cross-platform mobile client built with Flutter, Riverpod, Drift, Clean Architecture, and Supabase Auth.
It handles user authentication via Supabase, retrieves the JWT access token, and calls this Go backend API for profile management and user data operations.

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
- Swagger UI for interactive API documentation

---

## Tech Stack

- **Go 1.25**
  Statically typed, compiled language — high performance, simple deployment, minimal runtime overhead
- **chi v5**
  Lightweight HTTP router with composable middleware and route grouping, more flexible than standard `net/http`
- **pgx v5 + pgxpool**
  PostgreSQL driver with connection pooling via pgxpool, outperforms `database/sql` for concurrent workloads
- **sqlc**
  Generates type-safe Go code from SQL — no manual row scanning, SQL managed in `.sql` files
- **Supabase Auth + ES256/JWKS**
  JWTs are signed by Supabase using a private key (ES256); the backend fetches the public key from the JWKS endpoint for verification — more secure than symmetric HS256
- **slog (Go built-in)**
  Structured logger introduced in Go 1.21, outputs JSON format suitable for production observability
- **OpenAPI 3.0 + Swagger UI**
  API contract defined in `docs/openapi.yaml`, served via Docker as an interactive Swagger UI
- **Docker + Docker Compose**
  Multi-stage build (builder + alpine), Compose orchestrates PostgreSQL, Go API, and Swagger UI
- **Layered Architecture (Feature-first)**
  Each feature is grouped under `internal/<feature>/` with three internal layers:
  `handler` → `service` → `repository`.
  Each layer has a single responsibility — handlers never write SQL, services never touch HTTP, repositories never handle JWT.

---

## Security Design

- The client only sends `Authorization: Bearer <token>` — never a `user_id`
- The backend extracts `user_id` from JWT `claims.sub`
- All SQL `WHERE id = $1` conditions are derived from the JWT — client input is never trusted
- Supabase uses ES256 (asymmetric signing) — the backend holds only the public key; the private key never leaves Supabase

---

## Environment

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

This repository is intended for learning and demonstration.

If you plan to open-source it, please choose a license and confirm third-party asset usage rights.

---

## Project Structure

```
travel-audio-guide-go
├─ .dockerignore
├─ .idea
│  ├─ caches
│  │  └─ deviceStreaming.xml
│  ├─ copilot.data.migration.ask2agent.xml
│  ├─ markdown.xml
│  ├─ misc.xml
│  ├─ modules.xml
│  ├─ travel-audio-guide-go.iml
│  ├─ vcs.xml
│  └─ workspace.xml
├─ cmd
│  ├─ server
│  │  └─ main.go
│  └─ tmp
│     └─ build-errors.log
├─ docker-compose.yml
├─ Dockerfile
├─ docs
│  └─ openapi.yaml
├─ go.mod
├─ go.sum
├─ internal
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
│  └─ {auth,me,middleware,server,config,database,db}
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
