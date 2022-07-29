package unidecode

// Error represents an error that occurred during transliteration.
type Error struct {
	character rune
	index     int
	message   string
}

// Error returns the formatted error message.
func (e *Error) Error() string {
	return e.message
}
