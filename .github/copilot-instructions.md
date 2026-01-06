# GitHub Copilot Instructions for Librarian

## Project Overview
Go (1.25+) multi-media library manager using DDD/Clean Architecture, Kratos, and Google Wire.

## Tech Stack
- **Core**: Go, Wire (DI), Kratos.
- **Data**: SQL, Ent/GORM (check context).
- **Web UI**: `internal/service/angelaweb` (Tailwind CSS, DaisyUI).

## Coding Standards
- **General**:
  - Always handle errors explicitly. Wrap errors with context (`fmt.Errorf("...: %w", err)`).
  - Use `gci` standard imports order.
- **Testing**:
  - Generate table-driven tests for unit logic.
  - Mock interfaces in `biz/data` layers.
- **Architecture Constraints**:
  - `biz` layer: Pure Go, no external deps.
  - `pkg` layer: No imports from `internal`.

## Workflow
- If modifying `wire.go` or Protobufs, suggest running `make generate`.
- Ensure code passes `golangci-lint` (strict rules enabled).
