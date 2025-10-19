package goenvy

import (
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	LoadEnv()
	testKey := "TEST_KEY"
	testVal := "test-value"

	envVal := os.Getenv(testKey)
	if len(envVal) == 0 {
		t.Errorf("LoadEnv failed to get value for env key %q", testKey)
	} else if os.Getenv(testKey) != testVal {
		t.Errorf("LoadEnv value for key %q did not match value %q", testKey, testVal)
	}
}
