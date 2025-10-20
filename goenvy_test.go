package goenvy

import (
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	LoadEnv()

	noSpaceExpect := os.Getenv("NO_SPACE")
	noSpaceActual := "value1"
	spacedEntryExpect := os.Getenv("SPACED_ENTRY")
	spacedEntryActual := "value2"
	singleQuotedExpect := os.Getenv("SINGLE_QUOTED")
	singleQuotedActual := "single quote text allowed"
	doubleQuotedExpect := os.Getenv("DOUBLE_QUOTED")
	doubleQuotedActual := "double quote text allowed"
	singleQuotesMultilineExpect := os.Getenv("SINGLE_QUOTES_MULTILINE")
	singleQuotesMultilineActual := "single quotes first line\nanother line but ensure they are wrapped in quotes\n"
	doubleQuotesMultilineExpect := os.Getenv("DOUBLE_QUOTES_MULTILINE")
	doubleQuotesMultilineActual := "double quotes first line\nanother line but ensure they are wrapped in quotes"

	if noSpaceExpect != noSpaceActual {
		t.Errorf("LoadEnv failed to get environment variable with no space around '=' assignment operator")
	}
	if spacedEntryExpect != spacedEntryActual {
		t.Errorf("LoadEnv failed to get environment variable with space around the '=' assignment operator")
	}
	if singleQuotedExpect != singleQuotedActual {
		t.Errorf("LoadEnv failed to get environment variable with single quotes value")
	}
	if doubleQuotedExpect != doubleQuotedActual {
		t.Errorf("LoadEnv failed to get environment variable with double quotes value")
	}
	if singleQuotesMultilineExpect != singleQuotesMultilineActual {
		t.Errorf("LoadEnv failed to get environment variable with multiline values wrapped in single quotes")
	}
	if doubleQuotesMultilineExpect != doubleQuotesMultilineActual {
		t.Errorf("LoadEnv failed to get environment variable with multiline values wrapped in double quotes")
	}
}
