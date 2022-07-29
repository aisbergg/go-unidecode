# go-unidecode

[![Go Report Card](https://goreportcard.com/badge/github.com/aisbergg/go-unidecode)](https://goreportcard.com/report/github.com/aisbergg/go-unidecode)
[![GoDoc](https://godoc.org/github.com/aisbergg/go-unidecode?status.svg)](https://godoc.org/github.com/aisbergg/go-unidecode)

ASCII transliterations of Unicode text for Go. Inspired by [python-unidecode](https://github.com/avian2/unidecode).


## Installation

```
go get -u github.com/aisbergg/go-unidecode
```

Install CLI tool:

```
$ go get -u github.com/aisbergg/go-unidecode/unidecode

$ unidecode 北京kožušček
Bei Jing kozuscek
```


## Documentation

API documentation can be found here: https://godoc.org/github.com/aisbergg/go-unidecode


## Usage

```go
package main

import (
	"fmt"

	"github.com/aisbergg/go-unidecode/pkg/unidecode"
)

func main() {
	s := "abc"
	d, _ := unidecode.Unidecode(s, unidecode.Ignore)
	fmt.Println(d)
	// Output: abc

	s = "北京"
	d, _ = unidecode.Unidecode(s, unidecode.Ignore)
	fmt.Println(d)
	// Output: Bei Jing

	s = "kožušček"
	d, _ = unidecode.Unidecode(s, unidecode.Ignore)
	fmt.Println(d)
	// Output: kozuscek

	// return an error if an untransliteratable character is found
	s = "⁐"
	_, err := unidecode.Unidecode(s, unidecode.Strict)
	fmt.Println(err)
	// Output: no replacement found for character ⁐ in position 0

	// preserve untransliteratable characters
	d, _ = unidecode.Unidecode(s, unidecode.Preserve)
	fmt.Println(d)
	// Output: ⁐

	// replace untransliteratable characters with specified replacement text.
	d, _ = unidecode.Unidecode(s, unidecode.Replace, "?")
	fmt.Println(d)
	// Output: ?
}
```
