Daily Dining Hall is a simple service that emails subscribers the uOttawa dining hall menu every morning.

Website: Suspended for now to reduce operational expenses

## Architecture (brief)
- **Client (Next.js)**: Subscription UI.
- **Server (Go, microservices via gRPC)**:
  - `scraper`: fetches the daily menu.
  - `subscribers`: manages mailing list in Postgres.
  - `mailer`: sends emails (Mailgun).
  - `api`: HTTP gateway (Fiber) that the client calls; proxies to gRPC services.
  - `scheduler`: cron job that orchestrates scrape → list → email at 6 AM (America/Toronto).

Proto contracts live in `server/proto`. Run `make proto` in `server/` to generate stubs.
