package helpers

import (
	"unicode"
	"unicode/utf8"
)

// LCFirst returns the string with the first letter lower cased.
func LCFirst(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size <= 1 {
		return s
	}
	lc := unicode.ToLower(r)
	if r == lc {
		return s
	}
	return string(lc) + s[size:]
}

// UCFirst returns the string with the first letter upper cased.
func UCFirst(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size <= 1 {
		return s
	}
	c := unicode.ToUpper(r)
	if r == c {
		return s
	}
	return string(c) + s[size:]
}
