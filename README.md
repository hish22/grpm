# Grab Releases Package Manager (GRPM)

<img src="assets/grpm_w.png" />

GRPM is a command-line tool written in Go that helps you discover, manage, and install GitHub releases. It functions as a package manager specifically designed for GitHub releases, allowing you to search repositories, view release information, and download assets directly from your terminal.

## Installation

### Download a binary

Download the tool from the <a href="https://github.com/hish22/grpm/releases">releases</a>.

### Build from Source

```bash
git clone https://github.com/hish22/grpm.git
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
grpm config --open
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

Check if a specific installed release has a newer version available:

```bash
grpm update --repo <owner>/<repo> --check
```

### Update Releases

Update a specific release to the latest version:

```bash
grpm update --repo <owner>/<repo> --latest
```

### Remove Installed Packages

Remove an installed release:

```bash
grpm remove --repo <owner>/<repo>
```

### Clear Cache

Clear the search cache:

```bash
grpm cache --clear
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## License

MIT License
