# Commands Reference

This document provides detailed reference for all GRPM commands.

## Table of Contents

- [Global Options](#global-options)
- [search](#search)
- [config](#config)
- [info](#info)
- [release](#release)
- [install](#install)
- [list](#list)
- [update](#update)
- [cache](#cache)
- [remove](#remove)

## Global Options

| Flag | Short | Description |
|------|-------|-------------|
| `--define` | `-d` | Initialize GRPM configuration (requires admin) |

## search

Search for GitHub repositories.

```bash
grpm search [options]
```

### Options

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--repo` | `-r` | string | Repository name to search for |
| `--page` | `-p` | int | Page number of results (default: 1) |
| `--sort` | `-s` | string | Sort by: stars, forks, help-wanted-issues, updated |
| `--order` | `-o` | string | Order: asc, desc |

### Examples

```bash
# Search for repositories
grpm search --repo fzf

# Search sorted by stars (descending)
grpm search --repo fzf --sort stars --order desc

# Paginated search
grpm search --repo go --page 2
```

---

## config

Display or edit GRPM configuration.

```bash
grpm config [options]
```

### Options

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--show` | `-s` | bool | Show current configuration |
| `--open` | `-o` | bool | Open configuration file in editor |

### Examples

```bash
# View current configuration
grpm config --show

# Open configuration file
grpm config --open
```

---

## info

Display repository information including description, stars, forks, and owner details.

```bash
grpm info [options]
```

### Options

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--owner` | `-o` | string | Repository owner |
| `--repo` | `-r` | string | Repository name |

### Examples

```bash
# Full repository info
grpm info --owner junegunn --repo fzf

# Partial info (either owner or repo)
grpm info --owner junegunn
grpm info --repo fzf
```

---

## release

View release information for a repository.

```bash
grpm release [options]
```

### Options

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--repo` | `-r` | string | Repository in format owner/repo |
| `--latest` | `-a` | bool | Show 5 latest releases |
| `--latest-release` | `-l` | bool | Show only the latest release |
| `--tag` | `-t` | string | Show specific release by tag |

### Examples

```bash
# View latest release
grpm release --repo junegunn/fzf --latest-release

# View 5 latest releases
grpm release --repo junegunn/fzf --latest

# View specific release by tag
grpm release --repo junegunn/fzf --tag v0.17.0
```

---

## install

Install a release asset from a GitHub repository.

```bash
grpm install [options]
```

### Options

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--repo` | `-r` | string | Repository in format owner/repo |
| `--tag` | `-t` | string | Specific release tag to install |
| `--match` | `-m` | bool | Only show assets matching config (os/arch) |
| `--setup` | `-s` | bool | Auto-setup installed asset |
| `--force` | `-f` | bool | Auto-confirm all prompts |

### Examples

```bash
# Install latest release
grpm install --repo junegunn/fzf

# Install specific version
grpm install --repo junegunn/fzf --tag v0.17.0

# Install with auto-setup
grpm install --repo junegunn/fzf --setup

# Only show matching assets
grpm install --repo junegunn/fzf --match

# Force install (skip confirmations)
grpm install --repo junegunn/fzf --force
```

### Installation Process

1. Fetches release information from GitHub
2. Displays available assets with index numbers
3. Prompts user to select an asset by index
4. Downloads asset to the downloads directory
5. Extracts/setup the asset to the library directory
6. Tracks installation in the database

---

## list

List all installed assets.

```bash
grpm list
```

### Examples

```bash
# List all installed assets
grpm list
```

---

## update

Update installed assets to newer versions.

```bash
grpm update [options]
```

### Options

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--repo` | `-r` | string | Repository to update (owner/repo) |
| `--latest` | `-l` | bool | Update to latest version |
| `--force` | `-f` | bool | Auto-confirm all prompts |
| `--check` | `-c` | bool | Check if update available |

### Examples

```bash
# Update specific repository
grpm update --repo junegunn/fzf --latest

# Check for updates
grpm update --repo junegunn/fzf --check
```

---

## cache

Manage GRPM cache.

```bash
grpm cache [options]
```

### Options

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--clear` | `-c` | bool | Clear all cached data |

### Examples

```bash
# Clear cache
grpm cache --clear
```

### Cache Information

- Cache location: `~/.cache/grpm/` (Linux) or `C:\Users\<user>\.cache\grpm\` (Windows)
- Cache expiration: 24 hours
- Cache is used for GitHub API responses

---

## remove

Remove an installed asset.

```bash
grpm remove [options]
```

### Options

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--repo` | `-r` | string | Repository to remove (owner/repo) |

### Examples

```bash
# Remove installed asset
grpm remove --repo junegunn/fzf
```

### Notes

- Removes the asset from the library directory
- Removes tracking entry from the database
- Does not remove downloaded files

---

## Additional Commands

### new

Check for new versions of installed packages.

```bash
grpm new
```

### upgrade

Upgrade all installed packages to their latest versions.

```bash
grpm upgrade
```
