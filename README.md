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
grpm -d
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
- `install_dir`: Where assets are installed.
- `download_dir`: Where files are downloaded.
- `os`: Operating system filter.
- `arch`: Architecture filter.

#### Flags

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--show` | `-s` | bool | Show TOML configuration information |
| `--open` | `-o` | bool | Open TOML configuration file |

---

### Search Repositories

Search for GitHub repositories:

```bash
grpm search --repo <name>
```

Sort results by stars:

```bash
grpm search --repo <name> --sort stars --order desc    # Most stars first
grpm search --repo <name> --sort stars --order asc     # Fewest stars first
```

#### Flags

| Flag | Short | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--repo` | `-r` | string | | Repository name to search |
| `--page` | `-p` | int | 1 | Page number of results |
| `--sort` | `-s` | string | | Sort by: stars, forks, help-wanted-issues, updated |
| `--order` | `-o` | string | | Order: asc, desc |

---

### View Repository Information

Get repository details and README:

```bash
grpm info --owner <owner> --repo <name>
```

You can also provide just one of them:

```bash
grpm info --owner <owner>    # Uses most starred repo by owner
grpm info --repo <name>      # Uses most starred repo with that name
```

#### Flags

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--owner` | `-o` | string | Repository owner |
| `--repo` | `-r` | string | Repository name |

---

### Browse Releases

View latest releases (limited to 5):

```bash
grpm release --repo <owner>/<repo> --latest
```

View the latest release:

```bash
grpm release --repo <owner>/<repo> --latest-release
```

View a specific release by tag:

```bash
grpm release --repo <owner>/<repo> --tag v1.0.0
```

#### Flags

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--repo` | `-r` | string | Repository name (owner/repo) |
| `--latest` | `-a` | bool | Show 5 latest releases |
| `--latest-release` | `-l` | bool | Show only the latest release |
| `--tag` | `-t` | string | Show specific release by tag |

---

### Install Releases

Install the latest release:

```bash
grpm install --repo <owner>/<repo>
```

Install a specific version:

```bash
grpm install --repo <owner>/<repo> --tag v1.0.0
```

Auto-extract and setup the binary:

```bash
grpm install --repo <owner>/<repo> --setup
```

During installation, select an asset from the available release assets.

#### Flags

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--repo` | `-r` | string | Repository name (owner/repo) |
| `--tag` | `-t` | string | Specific release tag to install |
| `--match` | `-m` | bool | Show only assets matching config OS/arch |
| `--setup` | `-s` | bool | Auto extract and setup the asset |
| `--force` | `-f` | bool | Skip confirmation prompts |

---

### List Installed Packages

View all installed releases:

```bash
grpm list
```

---

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

#### Flags

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--repo` | `-r` | string | Repository name (owner/repo) |
| `--latest` | `-l` | bool | Update to the latest version |
| `--check` | `-c` | bool | Check if update is available |
| `--force` | `-f` | bool | Skip confirmation prompts |

---

### Remove Installed Packages

Remove an installed release:

```bash
grpm remove --repo <owner>/<repo>
```

#### Flags

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--repo` | `-r` | string | Repository name (owner/repo) |

---

### Clear Cache

Clear the search cache:

```bash
grpm cache --clear
```

#### Flags

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--clear` | `-c` | bool | Clear all cached data |

---

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## License

MIT License
