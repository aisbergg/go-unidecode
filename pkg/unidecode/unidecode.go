package unidecode

import (
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"
	"unsafe"

	"github.com/aisbergg/go-unidecode/internal/table"
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
	w      io.Writer
	buf    []byte
	repl   string
	errors ErrorHandling
}

// NewWriter returns a new Writer that transliterates Unicode text written to w
// into plain 7-bit ASCII. The behavior in case of untransliteratable characters
// is specified by the errors parameter. An optional replacement string can be
// provided to be used when errors is set to Replace.
func NewWriter(w io.Writer, errors ErrorHandling, replacement ...string) Writer {
	repl := ""
	if len(replacement) > 0 {
		repl = replacement[0]
	}
	// buffer for at least 32 characters
	buf := make([]byte, 32*utf8.UTFMax) // buffer for copying strings when needed
	return Writer{
		w:      w,
		buf:    buf,
		repl:   repl,
		errors: errors,
	}
}

// Write writes the contents of p to the underlying writer after transliterating
// it into plain 7-bit ASCII.
func (uw Writer) Write(p []byte) (n int, err error) {
	return uw.writeString(*(*string)(unsafe.Pointer(&p)))
}

// WriteString writes the contents of s to the underlying writer after
// transliterating it into plain 7-bit ASCII.
func (uw Writer) WriteString(s string) (n int, err error) {
	return uw.writeString(s)
}

func (uw Writer) writeString(s string) (n int, err error) { //nolint:revive
	buf := uw.buf[:0]
	writtenCount := 0
	for len(s) > 0 {
		// decode next rune
		r, size := utf8.DecodeRuneInString(s)
		s = s[size:]

		// keep ASCII characters as is
		if r < unicode.MaxASCII {
			// flush buffer if it is full
			if len(buf) >= cap(buf) {
				if n, err = uw.w.Write(buf); err != nil {
					return writtenCount + n, err
				}
				buf = buf[:0]
				writtenCount += n
			}
			buf = append(buf, byte(r))
			continue
		}

		trl := transliterateRune(r)
		if trl == "" {
			switch uw.errors {
			case Ignore:
				continue
			case Strict:
				return writtenCount, &Error{
					fmt.Sprintf("no replacement found for character '%c' at offset %d", r, writtenCount),
					r,
					writtenCount,
				}
			case Replace:
				trl = uw.repl
			case Preserve:
				trl = runeToString(r)
			default:
				panic("invalid value for errors parameter")
			}
		}
		// if the transliteration is larger than the buffer capacity, write directly
		if len(trl) > cap(buf) {
			if len(buf) > 0 {
				if n, err = uw.w.Write(buf); err != nil {
					return writtenCount + n, err
				}
				buf = buf[:0]
				writtenCount += n
			}
			if sw, ok := uw.w.(io.StringWriter); ok {
				if n, err = sw.WriteString(trl); err != nil {
					return writtenCount + n, err
				}
				writtenCount += n
				continue
			}
			// we have no choice, we have to allocate new bytes and copy the
			// transliteration string into it without making it unsafe
			if n, err = uw.w.Write([]byte(trl)); err != nil {
				return writtenCount + n, err
			}
			writtenCount += n
			continue

		} else if len(buf)+len(trl) > cap(buf) {
			// if the transliteration does not fit into the buffer, flush buffer
			if n, err = uw.w.Write(buf); err != nil {
				return writtenCount + n, err
			}
			buf = buf[:0]
			writtenCount += n
		}

		buf = append(buf, trl...)
	}

	// flush buffer one last time
	if len(buf) > 0 {
		if n, err = uw.w.Write(buf); err != nil {
			return writtenCount - len(buf) + n, err
		}
	}
	writtenCount += n
	return writtenCount, nil
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
