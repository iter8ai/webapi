# api

An Api that returns mocked weather data

`main.go` and a config package sit beside `go.mod`; Chi answers `/healthz`, and godotenv loads `.env` when present.

| | |
| --- | --- |
| Type | API |
| Language | Go |
| Framework | None stated |
| Packages | `github.com/go-chi/chi/v5`, `github.com/joho/godotenv` |
| Path | `apps/services/api` |
| Port | 8080 |

## Interfaces

Answers `GET /healthz` with 200; other methods on it get 405, unknown paths 404.

## Configuration

Read once at start-up from the environment; a `.env` file is loaded first when present, and real environment variables win. The same set is in `.env.example`.

- `PORT`: optional, default `8080`; must be a number from 1 to 65535.
- `ALLOWED_ORIGINS`: required, a comma-separated list of website origins. Start-up fails with an error naming it when it is missing or blank.

## Depends on

Nothing.

## Gaps

No weather endpoints yet, and no CORS middleware: `ALLOWED_ORIGINS` is parsed but not yet enforced.
