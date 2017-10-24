// Package pangram provides functions for testing if strings are pangrams.
package pangram

import "strings"

const testVersion = 2

// IsPangram test whether s is a pangram.
func IsPangram(s string) bool {
	seen := make([]bool, 26) // array of boolean, init to false
	for _, c := range strings.ToLower(s) {
		if c >= 'a' && c <= 'z' {
			seen[c-'a'] = true
		}
	}
	for _, v := range seen {
		if !v {
			return false
		}
	}
	return true
}
