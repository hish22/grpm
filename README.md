# Github Releases Packet Manager (GRPM)

GRPM is a command-line tool written in Go that helps you discover, manage, and install GitHub releases. It functions as a package manager specifically designed for GitHub releases, allowing you to search repositories, view release information, and download assets directly from your terminal.

## Features

- **Search**: Find GitHub repositories by name with sorting options
- **Info**: View repository details and README content
- **Releases**: Browse release information for any repository
- **Install**: Download and install release assets with tracking
- **Configure**: Customize installation paths and preferences
- **Track**: SQLite database tracks all installed releases

## Installation

### Build from Source

```bash
git clone https://github.com/yourusername/grpm.git
cd grpm
go build -o grpm ./cmd/grpm.go
```

### Initialize Configuration

Before using GRPM, initialize the configuration file:

```bash
./grpm --define
```

This creates `~/.config/grpm/config.toml` with default settings.

## Usage

### Configuration

Show current configuration:

```bash
grpm config --show
```

Edit configuration directly:

```bash
nano ~/.config/grpm/config.toml
```

Configuration options include:
- `install_dir`: Where assets are installed
- `download_dir`: Where files are downloaded
- `os`: Operating system filter
- `arch`: Architecture filter

### Search Repositories

Search for GitHub repositories:

```bash
grpm search --repo <name>
```

Sort results by stars:

```bash
grpm search --repo <name> --stars desc    # Most stars first
grpm search --repo <name> --stars asc     # Fewest stars first
```

### View Repository Information

Get repository details and README:

```bash
grpm info --owner <owner> --repo <name>
```

### Browse Releases

View latest releases (limited to 5):

```bash
grpm release --repo <owner>/<repo> --latest
```

View a specific release by tag:

```bash
grpm release --repo <owner>/<repo> --tag v1.0.0
```

### Install Releases

Install the latest release:

```bash
grpm install --repo <owner>/<repo>
```

Install a specific version:

```bash
grpm install --repo <owner>/<repo> --tag v1.0.0
```

During installation, select an asset from the available release assets.

### List Installed Packages

View all installed releases:

```bash
grpm list
```

### Check for Updates

Check if installed releases have newer versions available:

```bash
grpm new
```

### Upgrade Releases

Update a specific release:

```bash
grpm update --repo <owner>/<repo>
```

Upgrade all installed releases:

```bash
grpm upgrade
```

### Clear Cache

Clear the search cache:

```bash
grpm cache --clear
```

## Architecture

```
/home/pling77/Desktop/REAL_PROJECTS/grpm/
├── cmd/                           # CLI commands (Cobra framework)
│   ├── grpm.go                   # Main entry point
│   ├── searchc/                  # Search command
│   ├── installc/                 # Install command
│   ├── releasec/                 # Release info command
│   ├── infoc/                    # Repository info command
│   └── configc/                  # Configuration command
├── internal/                      # Core business logic
│   ├── config/                   # TOML configuration
│   ├── core/                     # HTTP and URL handling
│   ├── search/                   # GitHub search
│   ├── release/                  # Release fetching
│   ├── info/                     # Repository info
│   ├── install/                  # Asset installation
│   ├── persistance/              # SQLite database
│   ├── serialization/            # JSON handling
│   └── util/                     # Utilities
```

## Data Storage

GRPM stores data in `~/.config/grpm/`:
- `config.toml`: User configuration
- `metadata.db`: SQLite database tracking installed releases
- `cache/`: Cached API responses (24-hour expiration)

## Dependencies

- Go 1.25+
- spf13/cobra - CLI framework
- BurntSushi/toml - Configuration
- mattn/go-sqlite3 - Database
- dustin/go-humanize - Human-readable formatting
- lukechampine.com/blake3 - Cache hashing

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## License

MIT License
