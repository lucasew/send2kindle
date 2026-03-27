# Project Conventions and Agent Rules

## Operational Memory (Where things live)
- `cmd/send2kindle/` -> Core application entry points and application logic.
- `cmd/send2kindle/utils.go` -> Centralized error handling and utilities (`ReportError`, `Fatalf`, `Log`).
- `cmd/send2kindle/cleanup.go` -> Temporary file lifecycle and cleanup hooks.
- `.github/workflows/update-deps.yml` -> CI/CD workflows for updating dependencies.

## Key Directives

1. **Strict Tooling**:
   - The project uses Go pinned to version 1.21 via `mise`.
   - Tooling dependencies are managed using `mise.toml`.
   - Wildcard task definitions are used in `mise.toml` (e.g., `lint` depends on `lint:*`).

2. **Error Handling**:
   - **No Swallowed Errors**: The project treats ignored/swallowed errors as 'Medium' severity security issues.
   - All errors must be routed through centralized reporting tools found in `utils.go` (e.g., `Log`, `Fatalf`).

3. **Comments and Documentation**:
   - Docstrings must strictly use block comments (`/* ... */`).
   - Line comments (`//`) are explicitly forbidden outside of function bodies.

4. **Code Reuse**:
   - **Rule of Three**: Do not abstract code unless duplication occurs at least three times.

5. **Commit & Submission**:
   - Do NOT commit generated artifacts or compiled binaries (e.g., `send2kindle`, `install-mise.sh`).
   - Run `mise run ci` locally before submission to enforce quality.
   - Sentinel PR titles must follow: `🛡️ Sentinel: [Severity] [Description]`.
   - Janitor PR titles must follow: `🧹 Janitor: [Description]`.
   - Docs PR titles must follow: `📝 Docs: [Description]`.
   - Refactor PR titles must follow: `🛠️ Refactor: [Description]`.

6. **Agent Journals**:
   - Sentinel must maintain a journal in `.jules/sentinel.md` formatted as `- YYYY-MM-DD: [Severity] [Description]` for each fix.
