# TODO for Exposing Game State

This file outlines incremental steps to extend the server so it can provide data for neural network training, as suggested in `AGENTS.md`.

## 1. Review Existing Structures
- Inspect `defs/SessionSaveData` and `defs/SystemSaveData` to understand available game state fields. These are defined around lines 22-120 in `defs/savedata.go`【F:defs/savedata.go†L18-L120】.

## 2. Identify Current Endpoints
- Current API routes are registered in `api/common.go` lines 29-62 which include `/savedata/system/{action}` and `/savedata/session/{action}`【F:api/common.go†L29-L62】.
- Handler logic for these endpoints lives in `api/endpoints.go` starting at `handleSession` on line 173 and `handleSystem` on line 292【F:api/endpoints.go†L168-L211】【F:api/endpoints.go†L292-L324】.

## 3. Plan Training Data Export Endpoint
1. Add a new HTTP handler in `api/endpoints.go` (e.g. `handleTrainingData`) that returns combined `SessionSaveData` and `SystemSaveData` for a user. This should output structured JSON.
2. Register the route in `api/common.go` with a path like `GET /training/data` that requires authorization similar to other endpoints.
3. Use existing `savedata.GetSession` and `savedata.GetSystem` functions (files `api/savedata/session.go` and `api/savedata/system.go`) to fetch the data.

## 4. Update Database Layer if Needed
- If additional queries are required (for example retrieving the latest session slot), modify or add functions in `db/savedata.go` around the session retrieval functions (`ReadSessionSaveData`, `GetLatestSessionSaveDataSlot`) defined near lines 91-120 and 163-184【F:db/savedata.go†L91-L129】【F:db/savedata.go†L160-L188】.

## 5. Extend Data Definitions
- If the neural network requires additional fields not currently present, extend the structs in `defs/savedata.go` accordingly. Keep compatibility in mind and ensure the JSON tags remain consistent.

## 6. Follow Contribution Workflow
- Format all Go changes with `gofmt -w`.
- Run `go test ./...` and capture results before committing as required by `AGENTS.md` lines 33-41【F:AGENTS.md†L33-L41】.

