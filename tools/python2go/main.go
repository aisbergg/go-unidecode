package main

import (
	"bufio"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var singleQuotePattern = regexp.MustCompile(`'(?:[^'\\]|\\.)*',`)

func replaceSingleQuote(s string) string {
	s = s[1 : len(s)-2]
	s = strings.ReplaceAll(s, `\'`, `'`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	return `"` + s + `",`
}

func main() {
	files, _ := filepath.Glob("x*.py")
	if len(files) == 0 {
		die("No files found\n")
	}

	// go through all files and convert them
	for _, path := range files {
		filename := strings.TrimSuffix(path, ".py")

		file, err := os.Open(path)
		if err != nil {
			die("Error opening file %s: %s", path, err)
		}
		defer file.Close()

		content := make([]string, 0)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			switch line {
			case "data = (":
				content = append(content, "package table")
				content = append(content, fmt.Sprintf("var %s = []string{", filename))
				continue
			case ")":
				content = append(content, "}")
				continue
			}
			line = strings.ReplaceAll(line, "#", "//")
			line = strings.ReplaceAll(line, "None", `"\uffff"`)
			// replace single quotes with double quotes
			line = singleQuotePattern.ReplaceAllStringFunc(line, replaceSingleQuote)
			content = append(content, line)
		}

		if err := scanner.Err(); err != nil {
			die("Error reading file %s: %s", path, err)
		}

		converted := strings.Join(content, "\n")
		formatted, err := format.Source([]byte(converted))
		if err != nil {
			die("Error formatting %s: %s", filename, err)
		}

		// write to file
		outpath := fmt.Sprintf("../../pkg/unidecode/table/%s.go", filename)
		outfile, err := os.OpenFile(outpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			die("Failed to create file %s: %s", outpath, err)
		}
		defer outfile.Close()
		_, err = outfile.Write(formatted)
		if err != nil {
			die("Failed to write to file %s: %s", outpath, err)
		}
	}
}

func die(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
