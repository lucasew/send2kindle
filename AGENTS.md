# Agents Operational Memory & Project Conventions

This file provides historical context, codebase navigation, and guidelines for Agents working on the project.

## Architecture & Code Map
* Entry point -> `cmd/send2kindle/main.go`
* External binary calls -> `MustBinary` and `Command` in `cmd/send2kindle/utils.go`
* Centralized error handling / logging -> `Fatalf` / `Log` in `cmd/send2kindle/utils.go`
* Email logic -> `cmd/send2kindle/mail.go`
* EPUB Conversion -> `cmd/send2kindle/convert.go`
* Temporary files and cleanup -> `AddCleanupHook` and `Cleanup()` in `cmd/send2kindle/cleanup.go`
* Dependency management -> Go modules + Nix (`package.nix` / `shell.nix`)
* Task runner -> `mise` via `mise.toml`

## Project Rules
1. **Never ignore/swallow errors.** Empty `catch` blocks or ignored return errors are retroactive violations and treated as medium-priority security issues.
2. Route all unexpected errors through centralized reporting (e.g. `Fatalf` or `Log` in `cmd/send2kindle/utils.go`).
3. Always pin tool versions (e.g. `mise.toml` requires a specific Go version like `1.15.0`).
4. Avoid line comments (`//`) when writing documentation docstrings (`/* ... */`).
5. Ensure external dependencies (like `ebook-convert`) are validated through `MustBinary` before use.
