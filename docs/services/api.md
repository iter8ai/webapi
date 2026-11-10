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

Not written yet: the variables it reads.

## Depends on

Not written yet: other services and outside systems.

## Gaps

Not written yet: what it does not do yet.
