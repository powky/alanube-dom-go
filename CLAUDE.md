# CLAUDE.md - Project Context for Claude

## Project Overview

This is `alanube-go`, a Go SDK for the Alanube API (Dominican Republic DGII electronic invoicing). It wraps 86 REST API endpoints with typed Go structs and methods.

## Build & Test

```bash
go build ./...    # Build all packages including examples
go vet ./...      # Run static analysis
```

No test files exist yet. No external dependencies (stdlib only).

## Architecture

- **Single package**: `package alanube` (no sub-packages)
- **Client pattern**: `Client` struct with methods, created via `NewClient(token)` or `NewSandboxClient(token)`
- **Options pattern**: `WithHTTPClient`, `WithBaseURL`, `WithTimeout`
- **Internal methods**: `do()` for JSON requests, `doNoAuth()` for public endpoints, `doMultipartFile()` for file uploads
- **Error handling**: `Error` type with custom `UnmarshalJSON` that handles 3 API error formats
- **Typed constants**: Catalogs in `catalogs.go` and `catalogs_provinces.go`, error codes in `error_codes.go`

## Key Conventions

- All endpoint methods take `context.Context` as first param
- ByCompany variants use path pattern `/idCompany/{id}` (not query params)
- GET document endpoints accept `*DocumentQueryOptions` for `pdfType` query param
- List endpoints return paginated responses with `*PaginationMetadata`
- All document creation types share common structs: `IdDoc`, `Sender`, `Buyer`, `Totals`, `ItemDetail`
- Pointer types (`*int`, `*float64`) used for optional numeric fields with `omitempty`
- Webhook parsing via `ParseWebhook(r *http.Request)` with auto-detection

## File Organization

- `*_types.go` files contain structs/types only
- `*.go` files (without `_types`) contain client methods
- `catalogs*.go` contain constants
- `error_codes.go` contains error code constants and `IsErrorCode()` helper
- `examples/` directory has 6 runnable examples

## API Path Mapping (non-obvious paths)

| Document Type | API Path |
|---|---|
| Consumption Invoice (32) | `/invoices` (not `/consumption-invoices`) |
| Special Regimes (44) | `/special-regimes` |
| Governmental (45) | `/gubernamentals` |
| Export (46) | `/export-supports` |
| Payment Abroad (47) | `/payment-abroad-supports` |
| Sign Document | `/sign-document` (multipart/form-data) |
| Set Test Check | `/check-set-tests/{id}` |
| Directory | `/check-directory` |
| DGII Status | `/check-dgii-status` |
| Notify | `/{doctype}/notify-by-email` |
