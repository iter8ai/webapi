# web

Website that displayed weather data

A Vite single-page site with React; it has no runtime config. Where the repository holds an API, its dev server proxies `/api` to the first one, so no API address is baked into the build.

| | |
| --- | --- |
| Type | Website |
| Language | TypeScript |
| Framework | None stated |
| Packages | `react`, `react-dom`, `vite`, `@vitejs/plugin-react`, `vitest`, `typescript`, `@types/react`, `@types/react-dom`, `eslint`, `@eslint/js`, `typescript-eslint`, `globals` |
| Path | `apps/clients/web` |
| Port | 5173 |

## Interfaces

A single page at `/`. Commands: `pnpm --filter web dev`, `pnpm exec vitest run`, `pnpm exec vite build`.

## Configuration

None.

## Depends on

The API, only through the dev server's `/api` proxy to `http://localhost:8080`.

## Gaps

No weather display yet.
