# Portal

> Multi-repository Git branch management made simple

Portal is a CLI tool that lets you manage Git branches across multiple repositories simultaneously. Perfect for microservices or managing related projects.

## Features

- 🚀 Execute git operations across multiple repos in parallel
- 🌿 Create and switch branches across all repos with one command
- 📦 Organize repos into groups for targeted operations

## Installation

```bash
curl -fsSL https://raw.githubusercontent.com/xenomech/portal/main/install.sh | bash
```

Then add to your PATH (add to `~/.zshrc` or `~/.bashrc`):

```bash
export PATH="$HOME/.local/bin:$PATH"
```

#### Uninstalling Portal

```bash
curl -fsSL https://raw.githubusercontent.com/xenomech/portal/main/install.sh | bash -s uninstall
```

## Quick Start

```bash
# Add repositories
portal add ~/projects/abc/frontend -n frontend
portal add ~/projects/abc/identity -n identity
portal add ~/projects/abc/orders -n orders
portal add ~/projects/abc/checkout -n checkout
portal add ~/projects/abc/api -n api

# List all repos
portal list

# Create a feature branch across all repos
portal checkout -b feature/new-auth

# Switch to main branch in all repos
portal switch main

# Pull latest changes
portal pull

# Push changes
portal push
```

## Commands

### Repository Management

```bash
portal add <path>              # Add a repository
portal add <path> -n <name>    # Add with custom name
portal remove <name>           # Remove a repository
portal list                    # List all repositories
```

### Group Management

```bash
portal group add <name> <repo1> <repo2> ...   # Create group
portal group list                              # List groups
portal group remove <name>                     # Remove group
```

### Branch Operations

```bash
portal checkout <branch>                    # Checkout existing branch
portal checkout -b <branch>                 # Create new branch
portal checkout -b <branch> -g <group>      # Create in specific group
portal checkout -b <branch> -r <repo>       # Create in specific repo

portal switch <branch>                      # Switch to existing branch
portal switch <branch> -g <group>           # Switch in specific group
```

### Git Operations

```bash
portal pull                    # Pull from remote (all repos)
portal pull -g <group>         # Pull in specific group
portal pull -r <repo>          # Pull in specific repo

portal push                    # Push to remote
portal push -u                 # Push and set upstream
portal push -g <group>         # Push specific group
```

## Advanced Usage

### Per-Repo Base Branch checkouts

Create a new branch from different base branches in each repo:

```bash
portal checkout -b feature/auth \
  --from frontend=develop \
  --from backend=main \
  --from api=staging
```

### Sync Before Branch Creation

Fetch latest changes before creating a branch:

```bash
portal checkout -b feature/new-feature --sync
```

### Working with Groups

```bash
# Create a group for your microservices
portal group add microservices api gateway auth-service

# Create a branch only in microservices
portal checkout -b feature/new-endpoint -g microservices

# Pull only microservices
portal pull -g microservices
```

## Examples

### Daily Workflow

```bash
# Morning: sync all repos
portal pull

# Start new feature
portal checkout -b feature/user-profiles

# ... work on changes ...

# Push to all repos
portal push -u
```

### Microservices Development

```bash
# Setup
portal add ~/abc/api-gateway -n gateway
portal add ~/abc/auth-service -n auth
portal add ~/abc/user-service -n users
portal group add services gateway auth users

# Create feature branch from different bases
portal checkout -b feature/oauth \
  --from gateway=develop \
  --from auth=main \
  --from users=develop

# Work on feature...

# Push only services
portal push -g services
```

## Flags Reference

| Flag             | Short | Description                                  |
| ---------------- | ----- | -------------------------------------------- |
| `--create`       | `-b`  | Create new branch                            |
| `--group`        | `-g`  | Target specific group                        |
| `--repo`         | `-r`  | Target specific repo                         |
| `--from`         | `-f`  | Base branch for repo (format: `repo=branch`) |
| `--sync`         |       | Fetch before creating branch                 |
| `--set-upstream` | `-u`  | Set upstream when pushing                    |

## Configuration

Portal stores its configuration in `~/.config/portal/config.yaml`

Example config:

```yaml
repos:
  - path: /Users/you/projects/frontend
    name: frontend
  - path: /Users/you/projects/backend
    name: backend
groups:
  web:
    - frontend
    - backend
```

## Building from Source

```bash
git clone https://github.com/xenomech/portal.git
cd portal
make build
./dist/portal --version
```

## Requirements

- Git
- Go 1.21+ (for building from source)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE file for details

## Support

- 🐛 [Report issues](https://github.com/xenomech/portal/issues)

---

**Portal** - Manage your repos, not your sanity.
