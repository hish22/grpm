# System Overview

GRPM (GitHub Releases Packet Manager) is a command-line package manager for GitHub releases.

## What is GRPM?

GRPM allows users to:
- Search for GitHub repositories
- View repository information and releases
- Download and install release assets
- Track installed packages in a local database
- Check for updates and upgrade packages

## System Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                     CLI Layer (Cobra)                        │
│  search | config | info | release | install | list | update │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                    Business Logic Layer                      │
│   search | release | info | install | asset | config        │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                      Data Layer                               │
│         SQLite Database | File System | Cache                │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                   External Services                          │
│                   GitHub API                                  │
└─────────────────────────────────────────────────────────────┘
```

## Key Features

### 1. Search
- Query GitHub repositories by name
- Sort by stars, forks, or update time
- Paginated results

### 2. Release Management
- Fetch latest releases
- Browse all releases
- View specific release by tag

### 3. Installation
- Download release assets
- Support for multiple archive formats:
  - ZIP
  - tar.gz
  - tar.zst
- Platform-specific setup:
  - Binary installation
  - Symlink creation (Linux)
  - Environment variables (Windows)

### 4. Tracking
- SQLite database for installed packages
- Track version, installation path, and metadata
- Update checking

### 5. Caching
- Cache GitHub API responses
- 24-hour expiration
- BLAKE3 hash-based cache keys

## Data Storage

### Configuration
- Location: `~/.config/grpm/`
- File: `config.toml`

### Database
- Location: `~/.config/grpm/metadata.db`
- Type: SQLite
- Tables: releases, cache

### Cache
- Location: `~/.cache/grpm/`
- Format: JSON files
- Expiration: 24 hours

### Installation Directories
- Base: `/opt/grpm/` (Linux) or `C:\Tools\grpm\` (Windows)
- Library: `/opt/grpm/lib/`
- Downloads: `/opt/grpm/Downloads/`

## Supported Platforms

- **Operating Systems:** Linux, Windows
- **Architectures:** amd64, arm64, 386

## Dependencies

- Go 1.25+
- SQLite (via modernc.org/sqlite)
- Cobra (CLI framework)
- BurntSushi/toml (configuration)

## License

MIT License
