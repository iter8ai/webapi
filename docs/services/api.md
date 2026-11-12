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

Answers `GET /healthz` with 200.

## Configuration

| Variable | Required | Default | Meaning |
| --- | --- | --- | --- |
| `PORT` | no | `8080` | Port it listens on; a number from 1 to 65535. |
| `ALLOWED_ORIGINS` | yes | none | Comma-separated website origins. |

`.env` is loaded when present; `.env.example` lists the same set. Start-up fails with an error naming a missing or invalid variable.

## Depends on

None.

## Gaps

- No weather endpoints yet; `/healthz` is its only route.
- `ALLOWED_ORIGINS` is parsed but no CORS middleware uses it yet.
