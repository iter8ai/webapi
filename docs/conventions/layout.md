# Layout

Where each service lives in a repository that holds several, and what the root holds.

- Each service lives at the path `CLAUDE.md` gives it, with everything it needs in that folder: its config, `.env.example`, tests, linter and build.
- The root holds the Nx workspace and nothing that runs: a `package.json` whose only dependency is `nx`, `nx.json`, `pnpm-workspace.yaml`, the pipeline file, `CLAUDE.md`, `README.md` and `docs`.
- `nx.json` sets `defaultBase` to the repository's default branch and caches `lint`, `typecheck`, `test` and `build`.
- Each service has a `project.json` whose targets are `nx:run-commands` calling the service's own tools: `lint`, `test` and `build`, and `typecheck` where the language has one. No Nx plugin, and no root `go.work`.
- `pnpm-workspace.yaml` lists `apps/clients/*`, `apps/services/*`, `apps/tools/*` and `libs/js/*`, so every TypeScript or JavaScript service installs from the one root lockfile.
- The one pipeline installs, then runs `nx affected -t lint typecheck test build`. On GitHub it checks out with `fetch-depth: 0` and sets the base with `nrwl/nx-set-shas`. On Azure DevOps it checks out with `fetchDepth: 0` and computes the base itself: the previous commit on the default branch, the merge base with it anywhere else.
- `pnpm exec nx run-many -t lint typecheck test build` checks the whole repository.
- Each runnable service's default port is the one `CLAUDE.md` gives it, so every service can run at once.
- A website built with Vite reaches an API through its dev server's proxy: `/api` goes to the first API in `CLAUDE.md`'s Layout. An API's `ALLOWED_ORIGINS` in `.env.example` lists every website's origin.
- Beyond the proxy above, no service depends on another until a ticket needs it. `libs/<lang>/README.md` says how a library there is consumed.
