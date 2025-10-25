// Package to load and set environment variables from a file e.g .env file.
//
// You can load your environment variables from multiple files by calling any of the load functions (LoadEnv, LoadEnvPath etc) multiple times.
//
// Note: The last load function overwrites any previous environment variable value
//
// Usage:
//
//	package main
//	import (
//	  "os"
//	  goenvy "github.com/irabeny89/go-envy"
//	)
//	func main() {
//	 // invoke early to load and set variables in env file
//	 goenvy.LoadEnv() // 1: loads from default .env file
//	 // 2: (optional) load & append from `.env.development` file
//	 goenvy.LoadEnvPath(".env.development") // this will overwrite same key values assigned from step 1
//	 // then you can access variables like usual
//	 env := os.GetEnv("KEY")
//	}
package goenvy

import (
	"bufio"
	"os"
	"strings"
)

const defaultPath string = ".env"

func processEnvFile(file *os.File) {
	var multilineKey string
	var multilineVal string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		kvSlice := strings.Split(lineText, "=")

		isComment := strings.HasPrefix(lineText, "#")
		hasValidPattern := len(kvSlice) == 2
		isMultiline := len(multilineVal) != 0
		hasLineText := len(lineText) > 0

		if !isComment && hasValidPattern {
			k, v := kvSlice[0], kvSlice[1]
			k, v = strings.Trim(k, " "), strings.Trim(v, " ")

			isSingleLineDoubleQuotes := strings.HasPrefix(v, "\"") && strings.HasSuffix(v, "\"")
			isSingleLineSingleQuotes := strings.HasPrefix(v, "'") && strings.HasSuffix(v, "'")

			// single line with quotes e.g KEY='val'
			if isSingleLineDoubleQuotes || isSingleLineSingleQuotes {
				v = v[1 : len(v)-1] // remove enclosing quotes
				os.Setenv(k, v)
				continue
			}
			// multiline with quotes begins
			if strings.HasPrefix(v, "\"") || strings.HasPrefix(v, "'") {
				multilineKey = k
				multilineVal = v[1:] + "\n" // ignore starting quote eg `'` or `"`
				continue
			}
			// single line no quotes
			os.Setenv(k, v)
		}
		if isMultiline && hasLineText {
			// multiline ends with closing quote on a new line alone
			if len(lineText) == 1 {
				// do nothing & reset multiline variables
				multilineKey, multilineVal = "", ""
				continue
			}
			// multiline ends with text e.g abc'
			if strings.HasSuffix(lineText, "\"") || strings.HasSuffix(lineText, "'") {
				// ignore closing quote eg `'` or `"`
				multilineVal += lineText[:len(lineText)-1]
				os.Setenv(multilineKey, multilineVal)
				// reset multiline variables
				multilineKey, multilineVal = "", ""
				continue
			}
			// regular line, so add new line to look multiline on print
			multilineVal += lineText + "\n"
			os.Setenv(multilineKey, multilineVal)
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// This function reads the environment variables in a file (e.g. ".env") into the system environment at runtime.
//
// Note: ensure to call this function before accessing the environment variables from the file. It is best placed as the first line of code in your root program e.g "main.go" or [package].go.
//
// Usage:
//
//	package main
//	import (
//	  "os"
//	  goenvy "github.com/irabeny89/go-envy"
//	)
//	func main() {
//	 // invoke early to load and set variables in env file
//	 goenvy.LoadEnv()
//	 // then you can access variables like usual
//	 env := os.GetEnv("KEY")
//	}
func LoadEnv(path ...string) {
	var file *os.File
	var err error
	if len(path) == 0 {
		file, err = os.Open(defaultPath)
	} else {
		file, err = os.Open(path[0])
	}
	if err != nil {
		panic(err)
	}
	defer file.Close()
	processEnvFile(file)
}

// This function reads the environment variables from the file path (e.g. ".env.development") argument provided on execution.
//
// Note: ensure to call this function before accessing the environment variables from the file. It is best placed as the first line of code in your root program e.g "main.go" or [package].go.
//
// Usage:
//
//	package main
//	import (
//	  "os"
//	  goenvy "github.com/irabeny89/go-envy"
//	)
//	func main() {
//	 // invoke early to load and set variables in env file
//	 goenvy.LoadEnvPath()
//	 // then you can access variables like usual
//	 env := os.GetEnv("KEY")
//	}
func LoadEnvPath(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	processEnvFile(file)
}
