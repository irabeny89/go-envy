# Go Envy 🚀

`go-envy` is a zero-dependency Go package that makes it easy to load environment variables from `.env` files into your system environment. It supports comments, quoted strings, and even multiline values, making it flexible for complex configurations.

## Features

- ✅ **Automatic Loading**: Optionally load your `.env` file automatically at startup.
- ✅ **Multiline Support**: Handle complex strings spanning multiple lines.
- ✅ **Quoted Values**: Support for both single (`'`) and double (`"`) quotes.
- ✅ **Zero Dependencies**: Pure Go implementation using only the standard library.
- ✅ **Flexible Manual Loading**: Load and merge multiple environment files in sequence.

## Installation

```bash
go get github.com/irabeny89/go-envy
```

## How To Use

### 1. Create your `.env` file

Create a `.env` file in your project's root directory:

```sh
# This is a comment
DATABASE_URL=postgres://user:pass@localhost:5432/db
API_KEY="my-secret-key"
APP_ENV='development'

# Multiline values are supported
CERTIFICATE="-----BEGIN CERTIFICATE-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA75...
-----END CERTIFICATE-----"
```

### 2. Load environment variables

#### Automatic Loading (Recommended)

You can enable automatic loading of the default `.env` file by setting the `GOENVY_AUTO_LOAD` environment variable to `true` and importing the package with a blank identifier:

```go
package main

import (
 "os"
 // Automatically loads .env if GOENVY_AUTO_LOAD=true
 _ "github.com/irabeny89/go-envy"
)

func main() {
 dbURL := os.Getenv("DATABASE_URL")
 // Use your variables...
}
```

#### Manual Loading

If you prefer explicit control or need to load specific files:

```go
package main

import (
 "os"
 goenvy "github.com/irabeny89/go-envy"
)

func main() {
 // Load the default .env file manually
 err := goenvy.LoadEnvPath(".env")
 if err != nil {
  // Handle error (e.g., file not found)
 }

 // Load and override with environment-specific configurations
 goenvy.LoadEnvPath(".env.development")

 apiKey := os.Getenv("API_KEY")
}
```

## Configuration

You can control the behavior of the package using the following environment variables:

| Variable | Default | Description |
| :--- | :--- | :--- |
| `GOENVY_DEFAULT_PATH` | `.env` | The default path to load environment variables from. |
| `GOENVY_AUTO_LOAD` | `false` | Set to `true` or `1` to automatically load `.env` on startup. |
| `GOENVY_SHOW_LOGS` | `false` | Set to `true` or `1` to enable success/error logging to console. |

## Running Tests

```bash
go test -v .
```

## CI/CD and Releasing

This project uses an automated CI/CD pipeline powered by GitHub Actions and [Cocogitto](https://github.com/cocogitto/cocogitto).

### 1. Conventional Commits

Releases are triggered based on [Conventional Commits](https://www.conventionalcommits.org/). Every commit message must follow this format:

- `feat: ...` → Triggers a **Minor** version bump (e.g., v0.3.0 -> v0.4.0).
- `fix: ...` → Triggers a **Patch** version bump (e.g., v0.3.0 -> v0.3.1).
- `chore:`, `docs:`, `refactor:`, `test:`, `style:` → Do not trigger a version bump by default.
- `feat!:` or `fix!:` (or a `BREAKING CHANGE` footer) → Triggers a **Major** version bump.

### 2. Automated Workflow

When you push to the `main` branch:

1. **Tests**: All tests (Unit, Race Detector, Coverage) are executed.
2. **Versioning**: If tests pass, Cocogitto analyzes the new commits to calculate the next SemVer version.
3. **Tagging**: A new Git tag (with the `v` prefix) is automatically created and pushed.
4. **Release**: A GitHub Release is created with an automatically generated changelog based on your commit messages.

### 3. Manual Release

You can dry-run the version bump locally if you have `cog` installed:

```bash
cog bump --auto --dry-run
```

### 4. Git Hooks

To ensure your commit messages are always compliant, you can set up a Git hook that runs every time you commit:

```bash
chmod +x scripts/setup-hooks.sh
./scripts/setup-hooks.sh
```

This will create a `commit-msg` hook in your `.git` directory that validates your messages before allowing a commit.

> [!TIP]
> **Recommended**: Install [Cocogitto](https://github.com/cocogitto/cocogitto) locally. It provides a better experience for validating commits and handles versioning/changelogs automatically.

## License

This project is licensed under the MIT License.
