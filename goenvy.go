// Package goenvy provides a simple way to load environment variables from .env files.
// It supports basic key-value pairs, comments, quoted strings, and multiline values.
//
// Behavior:
// By default, goenvy can automatically load environment variables from a ".env" file in the
// root directory at startup if the environment variable GOENVY_AUTO_LOAD is set to "true" or "1".
//
// Configuration:
//   - GOENVY_DEFAULT_PATH: The default path to load environment variables from. Defaults to ".env".
//   - GOENVY_AUTO_LOAD: Set to "true" or "1" to enable automatic loading of the .env file on init.
//   - GOENVY_SHOW_LOGS: Set to "true" or "1" to show loading status and error messages in the console.
//
// Usage:
//
// Automatic loading:
//
//	import _ "github.com/irabeny89/go-envy" // Ensure GOENVY_AUTO_LOAD=true is set in your environment
//
// Manual loading:
//
//	import goenvy "github.com/irabeny89/go-envy"
//
//	func main() {
//	    // Load from default .env
//	    goenvy.LoadEnvPath(".env")
//
//	    // Load and override with environment-specific file
//	    goenvy.LoadEnvPath(".env.development")
//	}
package goenvy

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	defaultPath = ".env"
	autoLoad    = false
	showLogs    = false
)

// init automatically attempts to load the .env file if GOENVY_AUTO_LOAD is enabled.
func init() {
	configureAndLoad()
}

// configureAndLoad reads environment configuration and attempts to load the default .env file.
// It is separated from init() to allow for easier unit testing.
func configureAndLoad() {
	envDefaultPath := os.Getenv("GOENVY_DEFAULT_PATH")
	if envDefaultPath != "" {
		defaultPath = envDefaultPath
	}
	envAutoLoad := os.Getenv("GOENVY_AUTO_LOAD")
	if envAutoLoad != "" {
		autoLoad = envAutoLoad == "true" || envAutoLoad == "1"
	}
	if !autoLoad {
		return
	}
	envShowLogs := os.Getenv("GOENVY_SHOW_LOGS")
	if envShowLogs != "" {
		showLogs = envShowLogs == "true" || envShowLogs == "1"
	}
	if showLogs {
		fmt.Printf("[GOENVY.INFO] Auto-loading environment variables from default path: %s\n", defaultPath)
	}
	err := LoadEnvPath(defaultPath)
	if err != nil {
		if showLogs {
			fmt.Printf("[GOENVY.ERROR] Error loading %s file: %v\n", defaultPath, err)
			fmt.Println("[GOENVY.INFO] Ensure that the file exists and is in the root directory of your project.")
			fmt.Println("[GOENVY.INFO] Or load it yourself by calling LoadEnvPath(pathToFile) function in main function. e.g goenvy.LoadEnvPath(\".env.development\") in main function.")
		}
		return
	}
	if showLogs {
		fmt.Println("[GOENVY.INFO] You can override and load more with LoadEnvPath() function e.g goenvy.LoadEnvPath(\".env.development\") in main function.")
	}
}

// processEnvFile reads the provided reader line by line and sets environment variables.
// It handles comments (#), single/double quoted values, and multiline strings.
func processEnvFile(r io.Reader) error {
	var multilineKey string
	var multilineVal string
	var quoteChar string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lineText := scanner.Text()

		if multilineKey != "" {
			// In multiline mode
			if strings.HasSuffix(lineText, quoteChar) {
				// End of multiline
				multilineVal += lineText[:len(lineText)-1]
				os.Setenv(multilineKey, multilineVal)
				multilineKey = ""
				multilineVal = ""
				quoteChar = ""
			} else {
				multilineVal += lineText + "\n"
				// Update environment variable as we go (matching original behavior but potentially more robust)
				os.Setenv(multilineKey, multilineVal)
			}
			continue
		}

		if strings.HasPrefix(strings.TrimSpace(lineText), "#") || strings.TrimSpace(lineText) == "" {
			continue
		}

		kvSlice := strings.SplitN(lineText, "=", 2)
		if len(kvSlice) != 2 {
			continue
		}

		k := strings.TrimSpace(kvSlice[0])
		v := strings.TrimSpace(kvSlice[1])

		if (strings.HasPrefix(v, "\"") && strings.HasSuffix(v, "\"")) || (strings.HasPrefix(v, "'") && strings.HasSuffix(v, "'")) {
			// Single line with quotes
			os.Setenv(k, v[1:len(v)-1])
			continue
		}

		if strings.HasPrefix(v, "\"") || strings.HasPrefix(v, "'") {
			// Start multiline
			multilineKey = k
			quoteChar = v[:1]
			multilineVal = v[1:] + "\n"
			continue
		}

		// Regular single line
		os.Setenv(k, v)
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// LoadEnvPath reads the environment variables from the specified file path and sets them.
// If the file cannot be opened, it prints an error message to stdout (if logging is enabled or relevant).
// Existing environment variables with the same keys will be overwritten.
func LoadEnvPath(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := processEnvFile(file); err != nil {
		return err
	}
	if showLogs {
		fmt.Printf("[GOENVY.SUCCESS] Environment variables loaded from %s successfully.\n", path)
	}
	return nil
}

// Get returns the value of the provided key(k) from the
// environment and if not found it rerurns the default(d)
func Get(k, d string) string {
	v, exist := os.LookupEnv(k)
	if !exist {
		return d
	}
	return v
}
