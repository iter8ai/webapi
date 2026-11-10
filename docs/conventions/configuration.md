# Configuration

Where a service's settings come from and how it reads them.

- Every value comes from environment variables, read once at start-up.
- Each runnable service has one config struct or module, holding only what the code reads. A library has none: whoever calls it passes its settings.
- A `.env` file is loaded first when present. Never commit `.env`.
- A missing required variable stops start-up with an error that names it.
- An API reads its port from `PORT`, with a default, and its allowed origins from `ALLOWED_ORIGINS` where a browser calls it.
- Bake nothing environment-specific at build time (no `VITE_*`, no `NEXT_PUBLIC_*`), so the same build runs in every environment.
- `.env.example` lists every variable, and the Configuration section of the service's file under `docs/services` lists the same set.
