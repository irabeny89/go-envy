# Go Envy

Package to load and set environment variables from a file e.g `.env` file.
This function reads the environment variables in a file(e.g. ".env") into the system environment at runtime.

## How To Use

First, ensure a `.env` file contains environment variables you want within your code:

```sh
# comment
NO_SPACE=value1
SPACED_ENTRY = value2
SINGLE_QUOTED='single quote text allowed'
DOUBLE_QUOTED="double quote text allowed"
SINGLE_QUOTES_MULTILINE='single quotes first line
another line but ensure they are wrapped in quotes
'
DOUBLE_QUOTES_MULTILINE="double quotes first line
another line but ensure they are wrapped in quotes"
```

Ensure to call the `LoadEnv()` function before accessing the environment variables from the file.
It is best placed as the first line of code in your root program e.g `main.go` or `package.go`

```go
func main() {
  // invoke early to load and set variables in env file
  LoadEnv()
  // then you can access variables like usual
  env := os.GetEnv("KEY")
}
```

## Test

If you have access to the source code, you can run test.

```bash
go test
```
