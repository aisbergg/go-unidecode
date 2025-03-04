package unidecode

// Error represents an error that occurred during transliteration.
type Error struct {
	message   string
	character rune
	offset    int
}

// Error returns the formatted error message.
func (e *Error) Error() string {
	return e.message
}
