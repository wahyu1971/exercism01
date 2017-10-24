// Package acronym converts a phrase to its acronym.
package acronym

import (
	"fmt"
	"strings"
)

const testVersion = 3

// Abbreviate receives a string and returns its acronym.
func Abbreviate(s string) string {
	acronym := ""
	// We should tokenize on hiphens as well, so we replace them with white
	// to make our life easier by using split only once with spaces as the
	// delimiter.
	tokens := strings.Split(strings.Replace(s, "-", " ", -1), " ")
	for _, token := range tokens {
		// We keep appending the first letter of each tokens.
		acronym = fmt.Sprintf("%s%c", acronym, token[0])
	}
	// We return an uppercase acronym.
	return strings.ToUpper(acronym)
}
