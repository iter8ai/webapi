# webapi

A basic website and API

This file is the map of the repository. Iter8 wrote it when the project was created.

## Layout

- [api](docs/services/api.md): API in Go, at `apps/services/api`, port 8080.
- [web](docs/services/web.md): Website in TypeScript, at `apps/clients/web`, port 5173.

## Pipeline

`.github/workflows/ci.yml` runs lint, the tests and the build on every push, and deploys nothing.

## Commands

Not written yet: the first build fills in how to install, run, test, lint and build.

## Docs

Read the one that matches the work before changing it; none is loaded automatically.

- `docs/conventions/configuration.md`: Where a service's settings come from and how it reads them.
- `docs/conventions/testing.md`: What a change must prove before it merges.
- `docs/conventions/pipeline.md`: What runs on every push, and what never does.
- `docs/conventions/tooling.md`: The formatter, linter and test runner for each language.
- `docs/conventions/layout.md`: Where each service lives in a repository that holds several, and what the root holds.
- `docs/specs/technical-specs.md`: The technical specs written with the project; where it disagrees with a convention, it wins.

## Working rules

- Keep this file current. When a command, a service or a doc changes, change its line here in the same commit.
- A change to a service updates its `docs/services` file in the same commit.
