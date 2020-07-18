package proto

import (
	"fmt"
)

// Byte-safe string
func String(s string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(s), s)
}

// Inline string
func Inline(s string) string {
	return inline('+', s)
}

// Error
func Error(s string) string {
	return inline('-', s)
}

func inline(r rune, s string) string {
	return fmt.Sprintf("%s%s\r\n", string(r), s)
}