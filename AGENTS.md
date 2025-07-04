# AGENTS Guidelines

This repository contains the server implementation for **PokeRogue**. It exposes an HTTP API used by the game client for account management, save data handling and daily rankings. Subsequent agents working on this repo should aim to expose game state and interaction methods so that the server can be used to gather data for training a neural network player.

## Project Goals
- Provide endpoints that allow a client to authenticate, retrieve the current save data (system and session), update progress and query player/battle statistics.
- Ensure the API returns information in a structured JSON format that can be easily consumed by external programs.
- Maintain compatibility with the existing Go code style.

## Repository Structure
- `rogueserver.go` – entry point; sets up HTTP server and registers endpoints.
- `api/` – request handlers grouped by feature:
  - `account/` – registration, login and OAuth logic.
  - `savedata/` – retrieve and update system and session save data.
  - `daily/` – logic for daily runs and ranking retrieval.
  - `common.go` and `endpoints.go` – routing and helper functions.
- `db/` – MySQL/S3 persistence layer for accounts and saves.
- `defs/` – type definitions for the JSON structures returned by API calls (e.g. `SessionSaveData`, `SystemSaveData`).

Game state relevant to training a neural network is mostly represented by `defs.SessionSaveData` and `defs.SystemSaveData`. These structures include party composition, modifiers, arena data and various statistics. Clients interact with these via the `/savedata/...` endpoints in `api/endpoints.go`.

### Key Endpoints (see `api/endpoints.go`)
- `GET /account/info` – return username, linked OAuth ids and last save slot.
- `POST /account/login` – returns a base64 token used for authenticated requests.
- `GET|POST /savedata/system/{action}` – fetch or update system save data.
- `GET|POST /savedata/session/{action}` – fetch, update or delete a session save slot.
- `POST /savedata/updateall` – atomically update both system and session data.
- `GET /game/titlestats` – fetch current player and battle counts.
- `GET /daily/rankings` – leaderboard data for daily runs.

These endpoints require an `Authorization` header containing the base64 token from `/account/login`.

## Using the Training API

Several unauthenticated endpoints expose save data for neural network clients.

### `GET /training/data`
Returns a combined object containing the latest session save and the system
save for a player.

- Provide a `username` query parameter to fetch data for a specific account.
  If this parameter is present no `Authorization` header is required.
- Alternatively include an `Authorization` header to fetch data for the
  authenticated user.

Example request:

```bash
curl "http://<host>/training/data?username=Ash"
```

Example response:

```json
{
  "system": { "trainerId": 1, "timestamp": 1234567890 },
  "session": { "seed": "abc", "timestamp": 1234567890 }
}
```

### `GET /training/sessions`
Lists all session save slots with their timestamps. Use the same authentication
rules as `/training/data`.

Example response:

```json
[
  { "slot": 0, "timestamp": 1234567890 },
  { "slot": 1, "timestamp": 1234500000 }
]
```

### `POST /training/actions`
Queues an action that the game should execute. The request body must match the
`TrainingAction` structure:

```json
{
  "name": "UseItem",
  "args": { "itemId": 1 }
}
```

### `GET /training/actions`
Retrieves and clears the queued actions for the user. The response is an array
of `TrainingAction` objects. Authentication works the same way as the other
training endpoints.

## Style and Testing
- Format all Go code with `gofmt -w` before committing.
- After modifications run `go test ./...`. Dependencies may fail to download in this environment; report the failure if it occurs.

## Contribution Workflow
1. Make code or documentation changes.
2. Run `gofmt -w` on changed `.go` files.
3. Run `go test ./...` and capture the output.
4. Commit with a descriptive message.

Following these guidelines will keep the codebase consistent and help future agents extend the API or retrieve data for machine learning.
