package unidecode_test

import (
	"fmt"

	"github.com/aisbergg/go-unidecode/pkg/unidecode"
)

func ExampleUnidecode() {
	s := "北京kožušček"
	d, _ := unidecode.Unidecode(s, unidecode.Ignore)
	fmt.Println(d)
	// Output: Bei Jing kozuscek
}

func ExampleUnidecode_errorStrict() {
	s := "⁐"
	_, err := unidecode.Unidecode(s, unidecode.Strict)
	fmt.Println(err)
	// Output: no replacement found for character ⁐ in position 0
}

func ExampleUnidecode_errorPreserve() {
	s := "⁐"
	d, _ := unidecode.Unidecode(s, unidecode.Preserve)
	fmt.Println(d)
	// Output: ⁐
}

func ExampleUnidecode_errorReplace() {
	s := "⁐"
	d, _ := unidecode.Unidecode(s, unidecode.Replace, "?")
	fmt.Println(d)
	// Output: ?
}
