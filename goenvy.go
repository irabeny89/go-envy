// Package to load and set environment variables from a file e.g .env file.
package goenvy

import (
	"bufio"
	"os"
	"strings"
)

// This function reads the environment variables in a file into the system environment(e.g. ".env") at runtime.
//
// Note: ensure to call this function before accessing the environment variables from the file. It is best placed as the first line of code in your root program e.g "main.go" or [package].go
// func main() {
//  // invoke early to load and set variables in env file
//  LoadEnv()
//  // then you can access variables like usual
//  env := os.GetEnv("KEY")
// }
func LoadEnv() {
	file, err := os.Open(".env")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var multilineKey string
	var multilineVal string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		kvSlice := strings.Split(line, "=")

		isComment := strings.HasPrefix(line, "#")
		hasValidPattern := len(kvSlice) == 2
		isMultiline := len(multilineVal) != 0
		hasLine := len(line) != 0

		if !isComment && hasValidPattern {
			k, v := kvSlice[0], kvSlice[1]
			k, v = strings.Trim(k, " "), strings.Trim(v, " ")
			// multiline begins
			if strings.HasPrefix(v, "\"") || strings.HasPrefix(v, "'") {
				multilineKey = k
				multilineVal = v[1:] + "\n" // ignore closing quote eg `'` or `"`
				continue
			}
			os.Setenv(k, v)
		}
		if isMultiline && hasLine {
			// multiline ends
			if strings.HasSuffix(line, "\"") || strings.HasSuffix(line, "'") {
				multilineVal += line[:len(line)-1] // ignore closing quote eg `'` or `"`
				os.Setenv(multilineKey, multilineVal)
			} else {
				multilineVal += line + "\n"
			}
			continue
		}
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}
}
