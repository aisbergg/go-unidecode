package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aisbergg/go-unidecode/pkg/unidecode"
)

var usage = `Usage: unidecode [OPTIONS] STRING...`

func main() {

	errors := flag.String("e", "ignore", "How to handle errors, accepted values: ignore, strict, replace, preserve")
	flag.Parse()
	if len(flag.Args()) == 0 {
		die(usage)
	}
	errHnd := unidecode.ErrorHandling(0)
	*errors = strings.TrimSpace(strings.ToLower(*errors))
	switch *errors {
	case "ignore":
		errHnd = unidecode.Ignore
	case "strict":
		errHnd = unidecode.Strict
	case "replace":
		errHnd = unidecode.Replace
	case "preserve":
		errHnd = unidecode.Preserve
	default:
		die("invalid value for -e parameter")
	}

	input := strings.Join(flag.Args(), " ")
	ret, err := unidecode.Unidecode(input, errHnd)
	if err != nil {
		die("%v", err)
	}
	fmt.Println(ret)
}

func die(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
