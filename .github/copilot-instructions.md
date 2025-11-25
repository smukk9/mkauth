# Copilot / AI Agent Instructions for mkauth

Purpose: quickly orient an AI coding agent to the codebase so it can be productive with minimal back-and-forth.

- **Runbook (how to run & test locally):**
  - **Start the server:** this project reads its config via the `MKAUTH_FILE` environment variable. Example:
    `MKAUTH_FILE=config.yaml go run ./cmd/server`
  - **Build all packages:** `go build ./...`
  - **Run tests:** `go test ./...`

- **Big picture / architecture:**
  - Single Go service with a small web UI and HTTP API.
  - Entrypoint: `cmd/server/main.go` — it loads config (`internal/config.Load()`), opens the SQLite DB (`internal/db.New`) and creates the HTTP server (`internal/server.New`).
  - HTTP routing is centralized in `internal/server/server.go` via `registerRoutes()` which calls per-package `RegisterRoutes` functions.
  - Web UI templates are generated and placed under `internal/web/templates` (generated code) and the editable template is `internal/web/templates/home.templ`.

- **Key directories / files to reference:**
  - `cmd/server/main.go` — application main; how the app boots
  - `internal/config/config.go` — config loader; requires `MKAUTH_FILE` env var
  - `internal/db/db.go` — SQLite initialization & migrations
  - `internal/server/server.go` — HTTP server and route registration
  - `internal/health/` and `internal/web/handlers/` — examples of feature packages and how they register routes
  - `internal/web/templates/` — generated templates (`*_templ.go`) and source `home.templ`

- **Project conventions & patterns (explicit, discoverable):**
  - Package pattern: many features use a trio of files: `route.go` (or `routes.go`) exports `RegisterRoutes(mux, db, cfg)`, `service.go` contains business logic, and `store.go` contains DB interactions. Example: `internal/health/routes.go`, `internal/health/service.go`.
  - Handlers use a base `Handler` struct in `internal/web/handlers/handler.go` that carries common dependencies (`cfg`, `db`) and page handlers embed it (see `NewHomeHandler`).
  - Templates are generated. Do not edit files under `internal/web/templates/*_templ.go` directly — edit `.templ` and re-run the generator used in the project (the repo uses `github.com/a-h/templ` in generated code).

- **Integration points & external dependencies:**
  - SQLite via `github.com/mattn/go-sqlite3` — DB path is configured in the YAML file loaded by `MKAUTH_FILE`.
  - `github.com/spf13/viper` used for config file parsing.
  - `github.com/a-h/templ` used to generate web templates.

- **Important gotchas the agent should know (observed patterns):**
  - `internal/config.Load()` *requires* `MKAUTH_FILE` env var. If missing, the program returns an error (`ErrEnvVarNotSet`). Always set it when running or testing.
  - Route registration in code currently uses strings like `mux.HandleFunc("GET /health", ...)`. The standard library `http.ServeMux` expects pattern paths (e.g. `"/health"`). This *is a discoverable behaviour in the source* — be cautious when changing it and run tests / try the running server to validate routing.
  - Templates render by calling `templates.Home(service, version).Render(ctx, w)` (see `internal/web/handlers/home.go`). Frontend JS expects `/health` to return JSON/plain text.

- **Examples of common edits & where to implement them:**
  - Add a new API/feature: create `internal/<feature>/routes.go` with `RegisterRoutes`, `service.go` for logic, `store.go` for DB access. Register the feature in `internal/server/server.go` by calling `<feature>.RegisterRoutes(...)` from `registerRoutes()`.
  - Update home page UI: edit `internal/web/templates/home.templ` and regenerate templates (do not modify `*_templ.go` directly).
  - Change DB schema: update migrations in `internal/db/db.go` and verify `migrate()` changes do not break existing data; run `go test` and start the server to validate.

- **Developer workflow notes (practical commands & checks):**
  - Developer startup: `export MKAUTH_FILE=$PWD/config.yaml && go run ./cmd/server`
  - Quick health check: with server running, `curl -v http://localhost:<port>/health` (port is in `config.yaml` — see `Server.Port`).
  - Run unit tests: `go test ./...` — tests live next to packages (e.g., `internal/config/config_test.go`).

- **When you modify generated code:**
  - Prefer changing source `.templ` files and re-generating. If you must change generated Go, add a comment explaining why and ideally update the generator source so changes are reproducible.

If anything in this file looks incomplete or you want additional examples (e.g., how to add a new `user` route end-to-end), tell me which area to expand and I will iterate.
