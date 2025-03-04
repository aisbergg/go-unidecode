// Package unidecode provides ASCII transliterations of Unicode text. Unicode
// characters are mapped to ASCII characters based on their phonetic
// representation.
//
// The package provides three ways to transliterate Unicode text:
//
//  1. The [Unidecode] function for transliterates a string into plain 7-bit ASCII.
//  2. The [Append] function transliterates a string into plain 7-bit ASCII and appends the result to a byte slice.
//  3. The [NewWriter] function creates a writer that transliterates Unicode text into plain 7-bit ASCII and writes the result to an [io.Writer].
//
// The package also provides an [ErrorHandling] type that specifies how to handle errors during transliteration.
//
// The best results can be achieved by first applying NFC or NFKC normalizing to the input text:
//
//	import (
//		"golang.org/x/text/unicode/norm"
//		"github.com/aisbergg/go-unidecode/pkg/unidecode"
//	)
//
//	s := "北京kožušček"
//	n := norm.NFKC.String(s)
//	d, _ := unidecode.Unidecode(n, unidecode.Ignore)
//	fmt.Println(d)
//	// Output: Bei Jing kozuscek
package unidecode
