# Librarian

## Project Overview
Librarian is a multi-media library management application built with Go (1.25+). It uses DDD/Clean Architecture and Kratos framework patterns.

## Core Commands
- **Setup**: `make init` (Install tools)
- **Generate**: `make generate` (Update Wire, Protobuf, mocks. Run after `wire.go` or proto changes)
- **Lint**: `make lint` (Auto-fix formatting, strict checks)
- **Test**: `make test-unit` (Race detection enabled)
- **Build**: `make build`

## Architecture & Tech Stack
- **Structure**:
  - `cmd/`: Entry points (`admin`, `serve`).
  - `internal/biz`: Domain logic (Pure Go).
  - `internal/data`: Persistence (DB implementations).
  - `internal/service`: Transport (gRPC/HTTP).
  - `internal/service/angelaweb`: Embedded UI (Tailwind/DaisyUI).
  - `pkg/`: Public libs. **Prohibited**: Importing `internal/`.
- **DI**: [Google Wire](https://github.com/google/wire). See `cmd/wire_gen.go`.
- **Framework**: Kratos-based.

## Coding Standards
- **Style**: Strict `golangci-lint` (revive, funlen, cyclop).
- **Formatting**: `gci` + `golines` (Standard Go imports).
- **Error Handling**: Explicit check required (`errcheck`).
  - *Good*: `if err := do(); err != nil { return fmt.Errorf("do: %w", err) }`
  - *Bad*: `do()` (Ignoring error)
- **Tests**: Table-driven tests preferred for unit tests.

## Git Workflow
- **Commits**: [Conventional Commits](https://www.conventionalcommits.org/) (e.g., `feat: add scraper`, `fix: db timeout`).
- **Pre-PR Checklist**:
  1. `make generate`
  2. `make lint`
  3. `make test-unit`
