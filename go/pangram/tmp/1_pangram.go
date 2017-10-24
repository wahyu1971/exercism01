package pangram

import (
	"strings"
)

const testVersion = 2

func IsPangram(s string) bool {
	s = strings.ToLower(s)
	// what is struct{}{} ?, somekind of nil value ?
	// for what benefit ?
	x := map[rune]struct{}{}
	for _, v := range s {
		x[v] = struct{}{}
	}
	for i := rune(97); i <= rune(122); i++ {
		if _, ok := x[i]; !ok { // find non existant element
			return false
		}
	}
	return true
}
