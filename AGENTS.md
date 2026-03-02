# Agent Conventions

This file captures project conventions for AI assistants working in this repository.

## Operational Memory

- `cmd/send2kindle/` -> Main application logic for sending files to kindle via email
- `.github/workflows/` -> CI workflows, including `autorelease.yml` for automated releases

## Tools and Tasks

This project uses `mise` for task management. All tasks should be executed through `mise` via the configuration in `mise.toml`.
- Run formatting via `mise run fmt`
- Run linting via `mise run lint`
- Run tests via `mise run test`

## Code Guidelines
- Follow standard go formatting guidelines (`gofmt`).
- Do not check in build artifacts or installer scripts like `install-mise.sh`.
