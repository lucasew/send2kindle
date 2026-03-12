# Project Conventions & Operational Memory

This file serves as the source of truth for repository conventions and operational memory.

## Codebase Layout
- `cmd/send2kindle/` -> Main application entrypoint and logic for converting and emailing files.
- `cmd/send2kindle/utils.go` -> Shared utilities, logging, and centralized error handling functions (`Fatalf`, `MustSucess`).
- `cmd/send2kindle/cleanup.go` -> Temporary file management and cleanup hooks.
- `cmd/send2kindle/mail.go` -> Email composition and sending logic (SMTP).
- `cmd/send2kindle/convert.go` -> Wraps `ebook-convert` for format conversions.
- `cmd/send2kindle/main.go` -> CLI flag parsing and application coordination.
- `package.nix`, `shell.nix`, `default.nix` -> Nix packaging definitions.

## Conventions
- **Error Handling:** Use the centralized error handling functions in `cmd/send2kindle/utils.go` (`Fatalf`, `MustSucess`). Do not silently ignore errors or call `panic` directly unless strictly necessary. Ensure errors provide adequate context before program termination.
- **Task Management:** The project uses `mise` for task execution. Always use `mise run` for commands. Ensure all required tool versions are pinned in `.mise.toml` or `mise.toml`.
- **Code Quality:** Use standard Go tooling (`go vet ./...`, `gofmt -s -w .`). These are mapped to `mise run lint`.
- **CI/CD:** The repository uses a single GitHub Actions workflow (`.github/workflows/autorelease.yml`) for installation, codegen, CI, and releases.

## Memory / Context
- The project is a Go 1.15 application.
- Uses `mise` for task management; wildcard dependencies are used for group tasks.
- Ensure dependencies are never downgraded without explicit instruction.
