package unidecode_test

import (
	"fmt"
	"strings"

	"github.com/aisbergg/go-unidecode/pkg/unidecode"
)

func ExampleUnidecode() {
	s := "北京kožušček"
	d, _ := unidecode.Unidecode(s, unidecode.Ignore)
	fmt.Println(d)
	// Output: Bei Jing kozuscek
}

func ExampleUnidecode_errorStrict() {
	s := "北京⁐"
	_, err := unidecode.Unidecode(s, unidecode.Strict)
	fmt.Println(err)
	// Output: no replacement found for character '⁐' at offset 6
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

func ExampleAppend() {
	s := "北京kožušček"
	buf := make([]byte, 0, len(s)+len(s)/3)
	b, err := unidecode.Append(buf, s, unidecode.Ignore)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	// Output: Bei Jing kozuscek
}

func ExampleNewWriter() {
	s := "北京kožušček"
	bld := strings.Builder{}
	w := unidecode.NewWriter(&bld, unidecode.Ignore)
	w.Write([]byte(s))
	fmt.Println(bld.String())
	// Output: Bei Jing kozuscek
}
