# web

Website that displayed weather data

A Vite single-page site with React; it has no runtime config. Where the repository holds an API, its dev server proxies `/api` to the first one, so no API address is baked into the build.

| | |
| --- | --- |
| Type | Website |
| Language | TypeScript |
| Framework | None stated |
| Packages | `react`, `react-dom`, `vite`, `typescript`, `vitest`, `eslint`, `@eslint/js`, `typescript-eslint`, `@types/react`, `@types/react-dom` |
| Path | `apps/clients/web` |
| Port | 5173 |

## Interfaces

A single page at `/` with a heading. Scripts: `dev` (Vite on 5173, strict port), `build`, `lint`, `typecheck`, `test`.

## Configuration

None. Nothing environment-specific is baked into the build.

## Depends on

The dev server proxies `/api` to the api at `http://localhost:8080`, stripping the `/api` prefix.

## Gaps

- No API calls yet; the page does not use the proxy.
- No weather display yet.
