# AGENTS.md — protokolapparat

## What this repo is

A pure Go **library** (no `main` package). It is the central contract/schema module for the Apparat ecosystem. Producers and consumers share these types for Redis-backed event exchange.

## Project layout

```
pkg/
  common/    # ActionType, NowUnix()
  news/      # News model + NewsEvent / NewsSyncEvent
  schedule/  # ScheduleEntry model + ScheduleEvent / ScheduleSyncEvent
```

Each domain package follows the same file split:

- `model.go` — structs with JSON tags and a `Validate() error` method
- `events.go` — versioned event wrappers, constructors, and `Validate()`

## Commands

Use `mise` tasks (defined in `.mise.toml`):

| Task      | Command                                     |
| --------- | ------------------------------------------- |
| Run tests | `mise run test` (`go test -v -race ./...`)  |
| Lint      | `mise run lint` (`golangci-lint run ./...`) |
| Tidy      | `mise run tidy` (`go mod tidy`)             |
| Docs      | `mise run doc` (`pkgsite`)                  |

There are **no test files** yet; `go test` currently reports `no test files`.

## Linting / formatting

- Formatter is **`gofumpt`** with `extra-rules: true`, not `gofmt`.
- Key enabled linters: `staticcheck`, `revive`, `gocritic`, `godot`, `misspell`, `whitespace`, `tagalign`.
- `tagalign` requires struct tags to be aligned and sorted.
- `godot` requires comments to end with a period.
- `revive` enforces `package-comments`, `exported`, and standard error-naming rules.

## Code conventions

- **Schema versioning:** each domain package defines `const SchemaVersion = 1`. Event structs embed it in the `v` JSON field.
- **Every model and event must implement `Validate() error`.**
- **JSON tags:** use `snake_case`.
- **Time fields:** model times are RFC3339 strings; event metadata timestamps are Unix `int64` from `common.NowUnix()`.
- **Constructor pattern:** provide `New{Domain}{Action}Event` and `New{Domain}SyncEvent` helpers that set `Version`, `Action`, and `Timestamp`.
- **Action types:** use `common.ActionType` — values are `create`, `update`, `delete`, `sync`.

## Adding a new domain

1. Create `pkg/<domain>/model.go` with structs and `Validate()`.
2. Create `pkg/<domain>/events.go` with event wrappers, constructors, and `Validate()`.
3. Follow the `news` / `schedule` packages as templates.
4. Run `mise run lint` before committing — tag alignment and godot are easy to miss.
