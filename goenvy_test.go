// Package goenvy tests for environment variable loading and parsing.
package goenvy

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadEnvPath(t *testing.T) {
	content := `
BASIC=basic_value
SPACED = spaced_value
QUOTED="quoted_value"
SINGLE_QUOTED='single_quoted_value'
# This is a comment
EMPTY_LINE=

MULTILINE="line1
line2
line3"

WITH_EQUALS=key=value
`
	tmpDir := t.TempDir()
	envPath := filepath.Join(tmpDir, ".env.test")
	err := os.WriteFile(envPath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create temp env file: %v", err)
	}

	// Clean up environment variables before and after
	keys := []string{"BASIC", "SPACED", "QUOTED", "SINGLE_QUOTED", "EMPTY_LINE", "MULTILINE", "WITH_EQUALS"}
	clearEnv(keys)
	defer clearEnv(keys)

	LoadEnvPath(envPath)

	tests := []struct {
		key      string
		expected string
	}{
		{"BASIC", "basic_value"},
		{"SPACED", "spaced_value"},
		{"QUOTED", "quoted_value"},
		{"SINGLE_QUOTED", "single_quoted_value"},
		{"EMPTY_LINE", ""},
		{"MULTILINE", "line1\nline2\nline3"},
		{"WITH_EQUALS", "key=value"},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			actual := os.Getenv(tt.key)
			if actual != tt.expected {
				t.Errorf("expected %s=%q, got %q", tt.key, tt.expected, actual)
			}
		})
	}
}

func TestOverwriting(t *testing.T) {
	tmpDir := t.TempDir()
	env1 := filepath.Join(tmpDir, ".env1")
	env2 := filepath.Join(tmpDir, ".env2")

	os.WriteFile(env1, []byte("KEY=value1\nSHARED=orig"), 0644)
	os.WriteFile(env2, []byte("KEY=value2\nNEW=value3"), 0644)

	keys := []string{"KEY", "SHARED", "NEW"}
	clearEnv(keys)
	defer clearEnv(keys)

	LoadEnvPath(env1)
	if os.Getenv("KEY") != "value1" {
		t.Errorf("expected KEY=value1, got %s", os.Getenv("KEY"))
	}

	LoadEnvPath(env2)
	if os.Getenv("KEY") != "value2" {
		t.Errorf("expected KEY=value2 (overwritten), got %s", os.Getenv("KEY"))
	}
	if os.Getenv("SHARED") != "orig" {
		t.Errorf("expected SHARED=orig, got %s", os.Getenv("SHARED"))
	}
	if os.Getenv("NEW") != "value3" {
		t.Errorf("expected NEW=value3, got %s", os.Getenv("NEW"))
	}
}

func TestConfigureAndLoad(t *testing.T) {
	// Backup original state
	origAutoLoad := autoLoad
	origShowLogs := showLogs
	defer func() {
		autoLoad = origAutoLoad
		showLogs = origShowLogs
	}()

	t.Run("Disabled", func(t *testing.T) {
		os.Setenv("GOENVY_AUTO_LOAD", "false")
		configureAndLoad()
		if autoLoad != false {
			t.Errorf("expected autoLoad to be false")
		}
	})

	t.Run("Enabled", func(t *testing.T) {
		os.Setenv("GOENVY_AUTO_LOAD", "true")
		os.Setenv("GOENVY_SHOW_LOGS", "true")
		
		// Create a temporary .env file in the CURRENT directory because 
		// defaultPath is hardcoded to ".env"
		err := os.WriteFile(".env", []byte("AUTO_KEY=auto_val"), 0644)
		if err != nil {
			t.Fatalf("failed to create .env: %v", err)
		}
		defer os.Remove(".env")

		configureAndLoad()
		
		if autoLoad != true {
			t.Errorf("expected autoLoad to be true")
		}
		if showLogs != true {
			t.Errorf("expected showLogs to be true")
		}
		if os.Getenv("AUTO_KEY") != "auto_val" {
			t.Errorf("expected AUTO_KEY=auto_val, got %s", os.Getenv("AUTO_KEY"))
		}
		os.Unsetenv("AUTO_KEY")
	})
}

// TestInitBehavior tests the actual init() function behavior by running a subprocess.
// This is necessary because init() only runs once per process.
func TestInitBehavior(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping subprocess test in short mode")
	}

	// Create a temporary directory for the test project
	tmpDir := t.TempDir()
	
	// Create a .env file
	err := os.WriteFile(filepath.Join(tmpDir, ".env"), []byte("INIT_TEST=success"), 0644)
	if err != nil {
		t.Fatalf("failed to create .env: %v", err)
	}

	// Create a small go program that imports goenvy and prints an env var
	goCode := `
package main
import (
	"fmt"
	"os"
	_ "github.com/irabeny89/go-envy"
)
func main() {
	fmt.Print(os.Getenv("INIT_TEST"))
}
`
	err = os.WriteFile(filepath.Join(tmpDir, "main.go"), []byte(goCode), 0644)
	if err != nil {
		t.Fatalf("failed to create main.go: %v", err)
	}

	// Initialize go mod in the temp dir
	cmdMod := exec.Command("go", "mod", "init", "testapp")
	cmdMod.Dir = tmpDir
	if err := cmdMod.Run(); err != nil {
		t.Fatalf("go mod init failed: %v", err)
	}

	// Use go mod edit to point to the local version of goenvy
	absPath, _ := filepath.Abs(".")
	cmdReplace := exec.Command("go", "mod", "edit", "-replace", "github.com/irabeny89/go-envy="+absPath)
	cmdReplace.Dir = tmpDir
	if err := cmdReplace.Run(); err != nil {
		t.Fatalf("go mod edit replace failed: %v", err)
	}

	cmdRequire := exec.Command("go", "mod", "edit", "-require", "github.com/irabeny89/go-envy@v0.0.0")
	cmdRequire.Dir = tmpDir
	if err := cmdRequire.Run(); err != nil {
		t.Fatalf("go mod edit require failed: %v", err)
	}

	// Run the program with GOENVY_AUTO_LOAD=true and GOENVY_SHOW_LOGS=false
	cmdRun := exec.Command("go", "run", "main.go")
	cmdRun.Dir = tmpDir
	cmdRun.Env = append(os.Environ(), "GOENVY_AUTO_LOAD=true", "GOENVY_SHOW_LOGS=false", "GOPROXY=off")
	
	output, err := cmdRun.CombinedOutput()
	if err != nil {
		t.Fatalf("go run failed: %v\nOutput: %s", err, string(output))
	}

	if !strings.Contains(string(output), "success") {
		t.Errorf("expected output to contain 'success', got %q", string(output))
	}
}

func clearEnv(keys []string) {
	for _, k := range keys {
		os.Unsetenv(k)
	}
}
