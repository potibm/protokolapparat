# protokolapparat 🇩🇰

[![Go Reference](https://pkg.go.dev/badge/github.com/potibm/protokolapparat.svg)](https://pkg.go.dev/github.com/potibm/protokolapparat)
[![CI](https://github.com/potibm/protokolapparat/actions/workflows/ci.yml/badge.svg)](https://github.com/potibm/protokolapparat/actions/workflows/ci.yml)
[![Semantic Release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

**protokolapparat** is the central "contract" library for the Apparat ecosystem. It defines the shared Go types, validation logic, and versioned schemas used for seamless data exchange between distributed services via **Redis**.

By centralizing these definitions, we ensure that producers (like `tidsapparat` or `funkapparat`) and consumers (like `billedapparat`) speak the exact same language without code duplication.

## ✨ Features

*   **Domain-Driven Schemas:** Cleanly separated packages for `schedule`, `news`, and `common` types.
*   **Schema Evolution:** Built-in versioning (e.g., `v: 1`) in every payload to handle rolling updates and breaking changes.
*   **Strict Validation:** Every event type implements a `Validate()` method to ensure data integrity before it hits the wire.
*   **Redis-Ready:** Designed for easy serialization into JSON payloads for Redis Pub/Sub or Streams.
*   **Industrial Grade:** Automated versioning via Semantic Release and a multi-version Go matrix CI.

## 📁 Repository Structure

```text
pkg/
├── common/      # Shared types (ActionTypes, Envelopes)
├── news/        # News and announcement schemas
└── schedule/    # Timetable and agenda schemas