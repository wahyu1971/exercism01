// Package acronym provides the ability to create an acronym from a string.
package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 3

// Abbreviate returns an acronym for a given string. The acronym is formed from the first character of each grouping of word characters in the string.
func Abbreviate(phrase string) string {
	words := regexp.MustCompile(`\W+`).Split(phrase, -1)

	result := make([]string, len(words))

	for _, word := range words {
		result = append(result, strings.ToUpper(string(word[0])))
	}

	return strings.Join(result, "")
}
