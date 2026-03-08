# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**meaning-of-songs** is an interactive app to discover the meaning of songs. Monolith with Go backend and React frontend.

## Tech Stack

- **Backend**: Go 1.20, chi v5.0.12 (HTTP router)
- **Frontend**: React + TypeScript + Vite
- **Architecture**: Hexagonal (ports & adapters)

## Commands

```bash
./run.sh install-web   # Install frontend dependencies
./run.sh dev           # Run backend (port 8080)
./run.sh dev-web       # Run frontend with hot reload (port 5173)
./run.sh build         # Build frontend + backend
./run.sh run           # Build and run production
./run.sh clean         # Remove build artifacts
```

## Architecture

```
cmd/server/main.go                    # Entry point, dependency injection
internal/
  domain/                             # Entities + ports (interfaces)
    song.go                           # Song entity
    ports.go                          # SongRepository interface
  application/                        # Use cases (depends only on domain)
    song_service.go                   # SongService
  infrastructure/                     # Adapters (implements ports)
    http/                             # HTTP adapter (router, handlers, middleware)
    repository/memory/                # In-memory adapter for SongRepository
web/                                  # React frontend (Vite + TypeScript)
```

**Dependency rule**: `infrastructure → application → domain`. Domain has zero external dependencies.

## API

- `GET /api/health` - Health check
- `GET /api/songs` - List songs
- `GET /api/songs/{id}` - Get song by ID

## Adding a Database

Create a new adapter in `internal/infrastructure/repository/postgres/` implementing `domain.SongRepository`. Swap in `main.go`. No changes to domain or application layers.
