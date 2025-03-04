package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aisbergg/go-unidecode/pkg/unidecode"
)

func printUsage(w io.Writer) {
	flag.CommandLine.SetOutput(w)
	fmt.Fprintln(w, `Usage: unidecode [OPTIONS] STRING...

Options:`)
	flag.PrintDefaults()
	fmt.Fprintln(w, `
Examples:
  unidecode "北京kožušček" -e ignore
  cat file.txt | unidecode -e replace -r "?" -`)
}

func main() {
	errorsFlag := flag.String("e", "ignore", "How to handle errors, accepted values: ignore, strict, replace, preserve")
	replacementFlag := flag.String("r", "?", "Replacement character for -e replace")
	flag.Parse()
	if len(flag.Args()) == 0 {
		printUsage(os.Stderr)
		os.Exit(1)
	}
	errHnd := unidecode.ErrorHandling(0)
	*errorsFlag = strings.TrimSpace(strings.ToLower(*errorsFlag))
	switch *errorsFlag {
	case "ignore":
		errHnd = unidecode.Ignore
	case "strict":
		errHnd = unidecode.Strict
	case "replace":
		errHnd = unidecode.Replace
	case "preserve":
		errHnd = unidecode.Preserve
	default:
		fmt.Fprintln(os.Stderr, "invalid value for -e option")
		os.Exit(1)
	}

	input := strings.Join(flag.Args(), " ")
	var reader io.Reader
	if input == "-" {
		reader = os.Stdin
	} else {
		reader = strings.NewReader(input)
	}

	writer := unidecode.NewWriter(os.Stdout, errHnd, *replacementFlag)
	if _, err := io.Copy(writer, reader); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
