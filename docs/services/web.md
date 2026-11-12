# web

Website that displayed weather data

A Vite single-page site with React; it has no runtime config. Where the repository holds an API, its dev server proxies `/api` to the first one, so no API address is baked into the build.

| | |
| --- | --- |
| Type | Website |
| Language | TypeScript |
| Framework | None stated |
| Packages | `vite`, `react`, `react-dom`, `typescript`, `@types/react`, `@types/react-dom`, `vitest`, `eslint`, `@eslint/js`, `typescript-eslint` |
| Path | `apps/clients/web` |
| Port | 5173 |

## Interfaces

A single page at `/`. Scripts: `dev`, `build`, `lint`, `typecheck`, `test`.

## Configuration

None. The dev server proxies `/api` to `http://localhost:8080`.

## Depends on

The api, through the dev proxy only.

## Gaps

No weather display yet; the page is a heading only.
