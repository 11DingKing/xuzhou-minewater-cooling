# Xuzhou Minewater Cooling

Xuzhou Minewater Cooling coordinates the city's minewater exchange network: abandoned mine wells, heat-exchange stations, district cooling loops, building connections, maintenance work, and resident service requests. The service models the full lifecycle from a monitored minewater source through dispatch, allocation, metering, incident response, and audited recovery.

Operators work with two roles. Regional operators can manage stations and service requests in their assigned region; supervisors can approve capacity changes, close incidents, and review audit history. All state changes are persisted in SQLite migrations and are recoverable after restart.

## Run

```bash
GOTOOLCHAIN=local go run ./cmd/server
```

The service applies versioned SQLite migrations at startup and exposes `/healthz` and `/readyz`. The HTTP API uses request IDs, structured error responses, session expiry/revocation, and graceful worker shutdown.
