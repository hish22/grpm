# Architecture Guide

This document provides an overview of GRPM's internal architecture for developers.

## Project Structure

```
grpm/
├── cmd/                    # CLI commands (Cobra framework)
│   ├── grpm.go            # Main entry point
│   ├── searchc/           # Search command
│   ├── installc/         # Install command
│   ├── releasec/         # Release info command
│   ├── infoc/            # Repository info command
│   ├── configc/          # Configuration command
│   ├── listc/            # List installed command
│   ├── updatec/          # Update command
│   ├── cachec/           # Cache command
│   └── removec/          # Remove command
├── internal/              # Core business logic
│   ├── config/           # TOML configuration management
│   ├── coreHttp/         # HTTP and URL handling
│   ├── search/           # GitHub search functionality
│   ├── release/          # Release fetching
│   ├── info/              # Repository info
│   ├── install/          # Asset installation
│   ├── asset/            # Asset handling
│   ├── persistance/      # SQLite database and caching
│   ├── middlewares/      # Database middleware
│   ├── serialization/    # JSON handling
│   ├── structures/       # Data structures
│   ├── util/             # Utilities
│   ├── setup/            # Setup/extraction utilities
│   ├── remove/           # Asset removal
│   └── update/           # Update functionality
└── docs/                 # Documentation
```

## Core Components

### 1. CLI Layer (cmd/)

The CLI layer uses the Cobra framework to define commands. Each command is in its own package under `cmd/`.

**Main Entry Point (`cmd/grpm.go`)**
- Initializes all subcommands
- Handles the `--define` flag for configuration initialization

**Command Pattern**
Each command follows this pattern:
```go
func CommandNameC() *cobra.Command {
    c := cobra.Command{
        Use:   "command",
        Short: "Description",
        Run:   commandFunc,
        PreRun: func(cmd *cobra.Command, args []string) {
            // Check config exists
        },
    }
    // Add flags
    return &c
}
```

### 2. Configuration (`internal/config`)

Manages TOML configuration files and directory paths.

**Key Files:**
- `configuration.go` - Config struct definition
- `generate.go` - Config file generation
- `decode.go` - Config file reading
- `check.go` - Config existence checking

**Default Paths:**
- Linux: `~/.config/grpm/config.toml`
- Windows: `C:\Users\<user>\.config\grpm\config.toml`

### 3. HTTP Layer (`internal/coreHttp`)

Handles all GitHub API requests.

**Components:**
- `request.go` - HTTP request handling with timeout
- `handle_links.go` - URL building for API endpoints

**API Endpoints:**
- Search: `GET /search/repositories`
- Releases: `GET /repos/{owner}/{repo}/releases`
- Latest: `GET /repos/{owner}/{repo}/releases/latest`
- Tags: `GET /repos/{owner}/{repo}/releases/tags/{tag}`

### 4. Data Structures (`internal/structures`)

Defines data types for GitHub API responses:

- `Repository` - GitHub repository information
- `Release` - Release metadata
- `Assets` - Release asset information
- `Owner` - Repository owner information

### 5. Persistence (`internal/persistance`)

Manages SQLite database and caching.

**Database Tables:**
- `releases` - Tracks installed releases
- `cache` - Cache metadata

**Cache System:**
- Uses BLAKE3 hashing for cache keys
- 24-hour expiration
- Stored in `~/.cache/grpm/`

### 6. Installation (`internal/install`)

Handles asset download and installation:
- Downloads release assets
- Extracts archives (ZIP, tar.gz, tar.zst)
- Creates symlinks (Linux)
- Sets up environment variables (Windows)

### 7. Asset Management (`internal/asset`)

- Lists available assets
- Tracks installed assets
- Deletes assets

## Data Flow

### Search Command Flow
```
User Input → searchc.SearchC() 
  → search.SearchRepositories() 
  → coreHttp.ApiRequest 
  → GitHub API 
  → Cache (if cached) 
  → Display Results
```

### Install Command Flow
```
User Input → installc.InstallC()
  → release.FetchLatestRelease() / FetchSpecificRelease()
  → Display Assets
  → User selects asset
  → install.InstallSelectedAsset()
  → Download → Extract → Setup
  → Track in database
```

## Dependencies

**Direct Dependencies:**
- `github.com/spf13/cobra` - CLI framework
- `github.com/BurntSushi/toml` - TOML parsing
- `modernc.org/sqlite` - SQLite database
- `github.com/charmbracelet/log` - Logging
- `github.com/dustin/go-humanize` - Human-readable formatting
- `lukechampine.com/blake3` - Cache hashing

**Build Requirements:**
- Go 1.25+

## Extending GRPM

### Adding a New Command

1. Create directory: `cmd/newcommandc/`
2. Create `newcommand_command.go`:
```go
package newcommandc

import (
    "hish22/grpm/internal/config"
    "github.com/spf13/cobra"
)

var flagExample string

func NewCommandC() *cobra.Command {
    c := &cobra.Command{
        Use:   "newcommand",
        Short: "Description",
        Run:   newCommandFunc,
        PreRun: func(cmd *cobra.Command, args []string) {
            if !config.CheckConfig() {
                // Handle missing config
            }
        },
    }
    c.Flags().StringVarP(&flagExample, "example", "e", "", "Example flag")
    return c
}

func newCommandFunc(cmd *cobra.Command, args []string) {
    // Implementation
}
```

3. Add to main in `cmd/grpm.go`:
```go
r.AddCommand(newcommandc.NewCommandC())
```

### Adding New Asset Extraction

Add extraction logic in `internal/setup/`:
- `unzip_zip_*.go` - ZIP extraction
- `unzip_tar_*.go` - Tar extraction
- Platform-specific implementations in files ending with OS name

## Testing

The project includes unit tests for utilities:
- `internal/util/extractor_test.go`
- `internal/util/matcher_test.go`

Run tests:
```bash
go test ./...
```
