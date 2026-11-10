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

Read once at start-up from the environment; a `.env` file is loaded first when present. Listed in `.env.example`.

- `PORT`: optional, default `8080`; must be a number from 1 to 65535.
- `ALLOWED_ORIGINS`: required, comma-separated browser origins. The API exits at start-up naming it when it is missing or blank.

## Depends on

None.

## Gaps

- No weather endpoints yet.
- `ALLOWED_ORIGINS` is parsed but no CORS handling uses it yet.
