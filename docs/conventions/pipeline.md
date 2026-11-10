# Pipeline

What runs on every push, and what never does.

- One pipeline file: `.github/workflows/ci.yml` on GitHub, `azure-pipelines.yml` on Azure DevOps.
- On every push it installs, lints, typechecks where the language has a typechecker, tests and builds.
- It publishes nothing and deploys nothing.
- No Dockerfile: containers are not part of the repository yet.
