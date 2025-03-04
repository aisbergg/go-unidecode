package unidecode

import (
	"unsafe"
)

type stringBuilder struct {
	buf []byte
}

// newStringBuilder returns a new stringBuilder with the given initial capacity.
func newStringBuilder(n int) *stringBuilder {
	return &stringBuilder{buf: make([]byte, 0, n)}
}

// String returns the accumulated string.
func (b *stringBuilder) String() string {
	return unsafe.String(unsafe.SliceData(b.buf), len(b.buf))
}

// WriteByte appends the byte c to b's buffer.
// The returned error is always nil.
func (b *stringBuilder) WriteByte(c byte) error {
	b.buf = append(b.buf, c)
	return nil
}

// WriteString appends the contents of s to b's buffer.
// It returns the length of s and a nil error.
func (b *stringBuilder) WriteString(s string) (int, error) {
	b.buf = append(b.buf, s...)
	return len(s), nil
}
