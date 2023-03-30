package module01

import "strings"

// This will return the provided word in reverse order
func Reverse(word string) string {
	var sb strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		sb.WriteByte(word[i])
	}
	return sb.String()
}
