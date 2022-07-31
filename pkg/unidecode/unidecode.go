package unidecode

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

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
	b.Grow(len(s) + len(s)/3) // 33% extra capacity
	for _, r := range s {
		trl, ok := transliterateRune(r)
		if !ok {
			switch errors {
			case Ignore:
				continue
			case Strict:
				i := strings.IndexRune(s, r)
				return "", &Error{r, i, fmt.Sprintf("no replacement found for character %c in position %d", r, i)}
			case Replace:
				repl := ""
				if len(replacement) > 0 {
					repl = replacement[0]
				}
				trl = repl
			case Preserve:
				trl = runeToString(r)
			default:
				panic("invalid value for errors parameter")
			}
		}
		b.WriteString(trl)
	}
	return b.String(), nil
}

// transliterateRune converts the given rune into a latin representation.
func transliterateRune(r rune) (string, bool) {
	// keep ASCII characters as is
	if r < unicode.MaxASCII {
		return runeToString(r), true
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
	if trl == "" {
		return "", false
	}

	return trl, true
}

// runeToString converts the given rune into a string without allocating memory
// on the HEAP.
func runeToString(r rune) string {
	buf := make([]byte, utf8.UTFMax)
	w := utf8.EncodeRune(buf[:utf8.UTFMax], r)
	return string(buf[:w])
}
