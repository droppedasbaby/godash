package strutil

import (
	"strings"
	"unicode"
)

// CamelCase converts a string to camel case.
func CamelCase(s string) string {
	r := strings.Builder{}
	nextUpper := false

	for _, c := range s {
		if unicode.IsSpace(c) || c == '_' || c == '-' {
			nextUpper = true
			continue
		}

		if nextUpper {
			r.WriteRune(unicode.ToUpper(c))
			nextUpper = false
		} else {
			r.WriteRune(unicode.ToLower(c))
		}
	}

	return r.String()
}

// convertToCase is a utility function used by public case functions.
func convertToCase(s string, separator rune) string {
	r := strings.Builder{}

	for i, c := range s {
		if i > 0 && unicode.IsUpper(c) {
			r.WriteRune(separator)
		}
		r.WriteRune(unicode.ToLower(c))
	}

	return r.String()
}

// SnakeCase converts a string to snake case.
func SnakeCase(s string) string {
	return convertToCase(s, '_')
}

// KebabCase converts a string to kebab case.
func KebabCase(s string) string {
	return convertToCase(s, '-')
}
