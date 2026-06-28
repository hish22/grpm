# Configuration Guide

This guide explains how to configure GRPM for your needs.

## Configuration File Location

The configuration file is located at:
- Linux: `~/.config/grpm/config.toml`
- Windows: `C:\Users\<username>\.config\grpm\config.toml`

## Initializing Configuration

Before using GRPM, you must initialize the configuration. This requires administrator/root privileges:

```bash
grpm --define
# or
grpm -d
```

This creates the configuration file and necessary directories.

## Configuration Options

### Main Configuration

| Option | Description | Default (Linux) | Default (Windows) |
|--------|-------------|-----------------|-------------------|
| `location` | Installation base directory | `/opt/grpm` | `C:\Tools\grpm` |
| `library` | Library directory for installed assets | `/opt/grpm/lib` | `C:\Tools\grpm\lib` |
| `downloaded` | Download directory | `/opt/grpm/Downloads` | `C:\Tools\grpm\Downloads` |
| `os` | Operating system filter | Current OS | Current OS |
| `arch` | Architecture filter | Current architecture | Current architecture |

### Example Configuration

```toml
location = "/opt/grpm"
library = "/opt/grpm/lib"
downloaded = "/opt/grpm/Downloads"
arch = "amd64"
os = "linux"
```

## Viewing Configuration

Show current configuration:

```bash
grpm config --show
```

## Editing Configuration

You can manually edit the configuration file:

```bash
# Linux
nano ~/.config/grpm/config.toml

# Windows
notepad %USERPROFILE%\.config\grpm\config.toml
```

## Configuration for Asset Filtering

### OS Filter

Set the `os` option to filter assets by operating system:
- `linux`
- `windows`
- `darwin`

### Architecture Filter

Set the `arch` option to filter assets by architecture:
- `amd64`
- `arm64`
- `386`

### Using Match Mode

When installing, use the `--match` or `-m` flag to only show assets matching your configuration:

```bash
grpm install --repo owner/repo --match
```

## Data Storage

GRPM stores data in the following locations:

### Configuration Directory
- Path: `~/.config/grpm/`
- Files:
  - `config.toml` - User configuration
  - `metadata.db` - SQLite database tracking installed releases

### Cache Directory
- Linux: `~/.cache/grpm/`
- Windows: `C:\Users\<username>\.cache\grpm\`
- Contains cached API responses
- Cache expires after 24 hours

### Installation Directories
- Base: `/opt/grpm/` (Linux) or `C:\Tools\grpm\` (Windows)
- Library: `/opt/grpm/lib/` (Linux) or `C:\Tools\grpm\lib\` (Windows)
- Downloads: `/opt/grpm/Downloads/` (Linux) or `C:\Tools\grpm\Downloads\` (Windows)

## Troubleshooting

### Configuration Not Found

If you see an error about missing configuration:
```bash
Please Run (grpm -d) to define grpm configuration files
```

Run the initialization command with admin privileges:
```bash
grpm -d
```

### Permission Denied

Ensure you run the `--define` flag with administrator/root privileges. On Windows, run Command Prompt or PowerShell as Administrator. On Linux, use `sudo`.

### Cache Issues

Clear the cache if you encounter stale data:
```bash
grpm cache --clear
```
