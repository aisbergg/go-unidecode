package unidecode

import (
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"
	"unsafe"

	"github.com/aisbergg/go-unidecode/internal/table"
)

type Buffer interface {
	io.StringWriter
	fmt.Stringer
}

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

// Unidecode transliterates Unicode text into plain 7-bit ASCII.
func Unidecode(s string, errors ErrorHandling, replacement ...string) (string, error) {
	buf := newStringBuilder(len(s) + len(s)/3)
	if err := unidecode(buf, s, errors, replacement...); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// UnidecodeBytes transliterates Unicode text into plain 7-bit ASCII.
func UnidecodeBytes(b []byte, errors ErrorHandling, replacement ...string) ([]byte, error) { //nolint:revive
	buf := newStringBuilder(len(b) + len(b)/3)
	if err := unidecode(buf, *(*string)(unsafe.Pointer(&b)), errors, replacement...); err != nil {
		return nil, err
	}
	return buf.buf, nil
}

// unidecode is the internal implementation of Unidecode and UnidecodeBytes.
func unidecode(buf *stringBuilder, s string, errors ErrorHandling, replacement ...string) error { //nolint:revive
	repl := ""
	if len(replacement) > 0 {
		repl = replacement[0]
	}

	for pos, r := range s {
		// keep ASCII characters as is
		if r < unicode.MaxASCII {
			buf.WriteByte(byte(r)) //nolint:all
			continue
		}
		trl := transliterateRune(r)
		if trl == "" {
			switch errors {
			case Ignore:
				continue
			case Strict:
				return &Error{fmt.Sprintf("no replacement found for character '%c' at offset %d", r, pos), r, pos}
			case Replace:
				trl = repl
			case Preserve:
				trl = runeToString(r)
			default:
				panic("invalid value for errors parameter")
			}
		}
		buf.WriteString(trl) //nolint:all
	}
	return nil
}

// Append transliterates Unicode text into plain 7-bit ASCII, appends the
// result to the byte slice, and returns the updated slice.
func Append(b []byte, s string, errors ErrorHandling, replacement ...string) ([]byte, error) {
	buf := &stringBuilder{buf: b}
	if err := unidecode(buf, s, errors, replacement...); err != nil {
		return b, err
	}
	return buf.buf, nil
}

// AppendBytes transliterates Unicode text into plain 7-bit ASCII, appends the
// result to the byte slice, and returns the updated slice.
func AppendBytes(b, s []byte, errors ErrorHandling, replacement ...string) ([]byte, error) {
	buf := &stringBuilder{buf: b}
	if err := unidecode(buf, *(*string)(unsafe.Pointer(&s)), errors, replacement...); err != nil {
		return b, err
	}
	return buf.buf, nil
}

// Writer is an io.Writer that transliterates Unicode text into plain 7-bit ASCII.
type Writer struct {
	w  io.Writer
	sw io.StringWriter
	// when the underlying writer is not a StringWriter, we use a buffer to safely copy strings
	buf    []byte
	repl   string
	errors ErrorHandling
}

// NewWriter returns a new [Writer].
func NewWriter(w io.Writer, errors ErrorHandling, replacement ...string) Writer {
	repl := ""
	if len(replacement) > 0 {
		repl = replacement[0]
	}
	sw, _ := w.(io.StringWriter) //nolint:all
	buf := make([]byte, 96)
	return Writer{
		w:      w,
		sw:     sw,
		buf:    buf,
		repl:   repl,
		errors: errors,
	}
}

func (uw Writer) Write(p []byte) (n int, err error) {
	return uw.writeString(*(*string)(unsafe.Pointer(&p)), false)
}

func (uw Writer) WriteString(s string) (n int, err error) {
	return uw.writeString(s, true)
}

func (uw Writer) writeString(s string, mustCopy bool) (n int, err error) { //nolint:revive
	b := unsafe.Slice(unsafe.StringData(s), len(s))
	for pos, r := range s {
		// keep ASCII characters as is
		if r < unicode.MaxASCII {
			if uw.sw != nil { //nolint:all
				if _, err = uw.sw.WriteString(s[pos : pos+1]); err != nil {
					return pos, err
				}
			} else if mustCopy {
				// to avoid allocations while converting the string to a byte slice
				// we copy the string in a preallocated buffer
				cn := copy(uw.buf, s[pos:pos+1])
				if _, err = uw.w.Write(uw.buf[:cn]); err != nil {
					return pos, err
				}
			} else {
				// we can safely cast the string to a byte slice
				if _, err = uw.w.Write(b[pos : pos+1]); err != nil {
					return pos, err
				}
			}
			continue
		}

		trl := transliterateRune(r)
		if trl == "" {
			switch uw.errors {
			case Ignore:
				continue
			case Strict:
				return pos, &Error{fmt.Sprintf("no replacement found for character '%c' at offset %d", r, pos), r, pos}
			case Replace:
				trl = uw.repl
			case Preserve:
				trl = runeToString(r)
			default:
				panic("invalid value for errors parameter")
			}
		}
		if uw.sw != nil {
			if _, err = uw.sw.WriteString(trl); err != nil {
				return pos, err
			}
		} else {
			// to avoid allocations while converting the string to a byte slice
			// we copy the string in a preallocated buffer
			cn := copy(uw.buf, trl)
			if _, err = uw.w.Write(uw.buf[:cn]); err != nil {
				return pos, err
			}
		}
	}
	return len(s), nil
}

// transliterateRune converts the given rune into a Latin representation. An
// empty return value means that the rune cannot be transliterated.
func transliterateRune(r rune) string {
	// cannot transliterate private use characters
	if unicode.In(r, unicode.Co, unicode.Noncharacter_Code_Point) {
		return ""
	}
	return table.Lookup(r)
}

// runeToString converts the given rune into a string without allocating memory
// on the heap.
func runeToString(r rune) string {
	buf := make([]byte, utf8.UTFMax)
	w := utf8.EncodeRune(buf[:utf8.UTFMax], r)
	return string(buf[:w])
}
