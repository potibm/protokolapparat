# protokolapparat

[![go Reference](https://pkg.go.dev/badge/github.com/potibm/protokolapparat.svg)](https://pkg.go.dev/github.com/potibm/protokolapparat)
[![CI](https://github.com/potibm/protokolapparat/actions/workflows/ci.yml/badge.svg)](https://github.com/potibm/protokolapparat/actions/workflows/ci.yml)
[![Semantic Release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

**protokolapparat** is the central "contract" library for the Apparat ecosystem. It defines the shared Go types, validation logic, and versioned schemas used for seamless data exchange between distributed services via **Redis**.

By centralizing these definitions and utilizing **Go Generics**, we ensure that producers (like `tidsapparat`) and consumers (like `billedapparat`) speak the exact same language with absolute type safety and zero code duplication.

## Features

**Domain-Driven Schemas:** Cleanly separated packages for `schedule` and `news`

- **Generic Envelope (`common.Event[T]`):** A single, type-safe event wrapper for all actions (Create, Update, Delete, Sync).

* **Strict Validation:** Every payload implements a `Validate()` method to ensure data integrity before it hits the wire.

- **Redis-Ready:** Designed for easy serialization into JSON payloads for Redis Pub/Sub or Streams.

## Repository Structure

```text
pkg/
├-- common/      # Generic Event[T] envelope and ActionTypes
├-- news/        # News domain models (Entry)
└-- schedule/    # Timetable domain models (Entry, Category, Location)
```

## Installation

```bash
go get github.com/potibm/protokolapparat
```

## Usage

### Producer

```go
import (
    "github.com/potibm/protokolapparat/pkg/common"
    "github.com/potibm/protokolapparat/pkg/schedule"
)

// 1. Create a valid domain entry
entry := schedule.Entry{
    ID:    123,
    Title: "Workshop: Building Apparats",
}

// 2. Wrap it in a type-safe generic event
event := common.NewCreateEvent(entry)

// 3. Validate before publishing
if err := event.Validate(); err != nil {
    log.Fatalf("Invalid protocol message: %v", err)
}
```

### Consumer

```gor
// Tell Go exactly what type of payload to expect
var event common.Event[schedule.Entry]

// Unmarshal incoming JSON from Redis
if err := json.Unmarshal(redisData, &event); err != nil {
    log.Fatal(err)
}

// Process the type-safe payload
for _, item := range event.Payload {
    fmt.Println("Received item:", item.Title)
}
```

## Development

This repository uses **mise** for tool management and **Conventional Commits** for automated releases.

### Linting & Formatting

```bash
mise run fmt       # Auto-format Markdown and YAML using dprint
mise run lint      # RCheck formatting
```

### Commits

Please use the following prefixes for your commits to trigger the automated release process:

- `feat:` for new features (Minor release)
- `fix:` for bugfixes (Patch release)
- `feat!:` or `fix!:` for breaking changes (Major release)

## License

MIT – Created by [potibm](https://github.com/potibm)
