package unidecode

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/aisbergg/go-unidecode/pkg/unidecode/table"
)

// ErrorHandling specifies the behavior of Unidecode in case of an error.
type ErrorHandling uint8

const (
	// Ignore specifies that untransliteratable characters should be ignored.
	Ignore ErrorHandling = iota
	// Strict specifies that untransliteratable characters should cause an
	// error.
	Strict
	// Replace specifies that untransliteratable characters should be replaced
	// with a given replacement value.
	Replace
	// Preserve specifies that untransliteratable characters should be
	// preserved.
	Preserve
)

// Unidecode transliterates Unicode text into plain 7-bit ASCII. For best
// results you might use NFC or NFKC normalization prior to transliteration.
//
// Examples:
//   unidecode.Unidecode("北京kožušček", unidecode.Ignore) // Output: Bei Jing kozuscek
//   unidecode.Unidecode("⁐", unidecode.Strict) // Error: no replacement found for character ⁐ in position 0
//   unidecode.Unidecode("⁐", unidecode.Preserve) // Output: ⁐
//   unidecode.Unidecode("⁐", unidecode.Replace, "?") // Output: ?
func Unidecode(s string, errors ErrorHandling, replacement ...string) (string, error) {
	var b strings.Builder
	b.Grow(len(s) + len(s)/10)
	for i, r := range []rune(s) {
		tr, ok := transliterateRune(r)
		if !ok {
			switch errors {
			case Ignore:
				continue
			case Strict:
				return "", &Error{r, i, fmt.Sprintf("no replacement found for character %c in position %d", r, i)}
			case Replace:
				repl := ""
				if len(replacement) > 0 {
					repl = replacement[0]
				}
				tr = repl
			case Preserve:
				tr = string(r)
			default:
				panic("invalid value for errors parameter")
			}
		}
		b.WriteString(tr)
	}
	return b.String(), nil
}

// transliterateRune converts the given rune into a latin representation.
func transliterateRune(r rune) (string, bool) {
	// keep ASCII characters as is
	if r < unicode.MaxASCII {
		return string(r), true
	}

	// cannot transliterate private use characters
	if unicode.In(r, unicode.Co, unicode.Noncharacter_Code_Point) {
		return "", false
	}

	// transliterate rune using lookup table
	section := r >> 8   // Chop off the last two hex digits
	position := r % 256 // Last two hex digits
	tb, ok := table.Tables[section]
	if !(ok && int(position) < len(tb)) {
		return "", false
	}
	trl := tb[position]
	if trl == "\uffff" {
		return "", false
	}

	return trl, true
}
