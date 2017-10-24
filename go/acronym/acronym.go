package acronym

import (
	"strings"
)

const testVersion = 3

func abbr2(s string) (result string) {
	for _, w := range strings.Split(s, "-") {
		result = result + strings.ToUpper(w[0:1])
	}
	return
}

func Abbreviate(s string) (result string) {
	for _, w := range strings.Split(s, " ") {
		if strings.Contains(s, "-") {
			result = result + abbr2(w)
		} else {
			result = result + strings.ToUpper(w[0:1])
		}
	}
	return
}
