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
| `GOENVY_AUTO_LOAD` | `false` | Set to `true` or `1` to automatically load `.env` on startup. |
| `GOENVY_SHOW_LOGS` | `false` | Set to `true` or `1` to enable success/error logging to console. |

## Running Tests

To run the test suite:

```bash
go test -v .
```

## License

This project is licensed under the MIT License.
