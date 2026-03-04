# Getting Started with GRPM

GRPM (GitHub Releases Packet Manager) is a command-line tool written in Go that helps you discover, manage, and install GitHub releases. It functions as a package manager specifically designed for GitHub releases.

## Prerequisites

- Go 1.25 or later
- Git
- Administrator/root privileges (required for initial configuration)

## Installation

### Build from Source

```bash
git clone https://github.com/yourusername/grpm.git
cd grpm
go build -o grpm ./cmd/grpm.go
```

### Initialize Configuration

Before using GRPM, you must initialize the configuration file. This requires administrator privileges:

```bash
./grpm --define
# or
./grpm -d
```

This creates:
- `~/.config/grpm/config.toml` - User configuration
- `/Tools/grpm/` (Windows) or `/opt/grpm/` (Linux) - Installation directories
- `metadata.db` - SQLite database for tracking installed releases

## Quick Start

### 1. Search for a Repository

```bash
grpm search --repo <repository-name>
```

Sort by stars:
```bash
grpm search --repo <name> --stars desc    # Most stars first
grpm search --repo <name> --stars asc     # Fewest stars first
```

### 2. View Repository Information

```bash
grpm info --owner <owner> --repo <name>
```

### 3. Browse Releases

View latest release:
```bash
grpm release --repo <owner>/<repo> --latest
```

View specific release:
```bash
grpm release --repo <owner>/<repo> --tag v1.0.0
```

### 4. Install a Release

Install latest release:
```bash
grpm install --repo <owner>/<repo>
```

Install specific version:
```bash
grpm install --repo <owner>/<repo> --tag v1.0.0
```

Auto-setup installed asset:
```bash
grpm install --repo <owner>/<repo> --setup
```

### 5. List Installed Packages

```bash
grpm list
```

### 6. Check for Updates

```bash
grpm new
```

### 7. Upgrade Releases

Update specific release:
```bash
grpm update --repo <owner>/<repo>
```

Upgrade all:
```bash
grpm upgrade
```

## Common Workflows

### Finding and Installing a Tool

1. Search for the repository:
   ```bash
   grpm search --repo fzf
   ```

2. View repository details:
   ```bash
   grpm info --owner junegunn --repo fzf
   ```

3. View available releases:
   ```bash
   grpm release --repo junegunn/fzf --latest
   ```

4. Install:
   ```bash
   grpm install --repo junegunn/fzf
   ```

5. Select an asset from the list by entering its index number

### Keeping Packages Updated

1. Check for new versions:
   ```bash
   grpm new
   ```

2. Upgrade all packages:
   ```bash
   grpm upgrade
   ```

### Managing Cache

Clear search cache:
```bash
grpm cache --clear
```

Cache expires after 24 hours automatically.

## Next Steps

- Read the [Configuration Guide](configuration.md) to customize your installation
- Read the [Commands Reference](commands.md) for detailed command documentation
- Read the [Architecture Guide](architecture.md) for developers
