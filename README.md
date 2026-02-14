# Projeto Buraco

This repository contains a Go implementation of the card game "Buraco" (a popular rummy-like game). The project provides the core game logic, table handling, card management, and a small HTTP API to interact with the game.

## Contents

- `main.go` - application entry point.
- `api.go` - minimal HTTP API endpoints for interacting with games.
- `cartas.go` - card representations and utilities.
- `mesa.go` - table and game state management.
- `regras.go` - game rules and turn logic.
- `testes.http` - example HTTP requests for testing the API.

## Requirements

- Go 1.18 or newer

## Build and Run

To build the binary:

```bash
go build -o buraco
```

To run directly with `go run`:

```bash
go run .
```

If the project exposes an HTTP API, use `testes.http` or tools like `curl` / `httpie` to exercise endpoints.

## Project Structure

The code is organized around core responsibilities: card handling (`cartas.go`), rules (`regras.go`), table/game state (`mesa.go`), and an optional API layer (`api.go`). The `main.go` file wires these components to start the application.