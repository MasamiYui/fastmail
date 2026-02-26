# Tasks

- [x] Task 1: Project Initialization
  - [x] Initialize Go module (`go mod init fastmail`)
  - [x] Create project structure (`cmd/`, `internal/config`, `internal/api`, `internal/service`)
  - [x] Add `.gitignore` and `README.md` updates.

- [x] Task 2: Configuration Management
  - [x] Implement `internal/config` package using `viper` to read from `config.yaml`, environment variables, and flags.
  - [x] Define `Config` struct (Server Port, Auth Token, SMTP Host/Port/User/Pass).
  - [x] Add support for command-line flags (e.g., `-config`, `-port`).

- [x] Task 3: SMTP Service Implementation
  - [x] Implement `internal/service/email.go` using `gopkg.in/gomail.v2` (or similar library).
  - [x] Create `EmailService` struct with `SendEmail` method.
  - [x] Support HTML body and file attachments.
  - [x] Implement connection pooling/reuse if possible, or create new connection per request (simple start).

- [x] Task 4: HTTP API Implementation
  - [x] Implement `internal/api/handler.go` using `gin` framework.
  - [x] Create `POST /api/v1/send` handler.
  - [x] Parse `multipart/form-data` for: `to`, `subject`, `body`, `attachments`.
  - [x] Validate required fields.
  - [x] Call `EmailService.SendEmail`.
  - [x] Return JSON response (Success/Error).

- [ ] Task 5: Authentication & Middleware
  - [x] Implement `AuthMiddleware` to check `Authorization` header against configured token.
  - [ ] Implement `RequestLogger` middleware for structured logging.
  - [x] Add `GET /health` endpoint for readiness checks.

- [x] Task 6: Main Application Entry Point
  - [x] Wire `config`, `service`, and `api` in `cmd/server/main.go`.
  - [x] Implement graceful shutdown handling (context cancellation on SIGINT/SIGTERM).

- [x] Task 7: Build & Documentation
  - [x] Create `Makefile` for building the binary (`make build`).
  - [x] Create example `config.yaml`.
  - [x] Document usage in `README.md`.
