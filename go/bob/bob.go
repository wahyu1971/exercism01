package bob

import (
	"regexp"
	"strings"
)

const testVersion = 3

func Hey(in string) string {
	re_has_upper := regexp.MustCompile("[[:upper:]]")
	re_sure := regexp.MustCompile("^.*\\?$") // ?
	re_whoa := regexp.MustCompile("^[[:upper:][:punct:][[:blank:][:digit:]]+$")
	re_fine := regexp.MustCompile("^[[:space:]]*$")

	in = strings.Trim(in, " \t")
	switch {
	case re_has_upper.MatchString(in) && re_whoa.MatchString(in):
		return "Whoa, chill out!"
	case re_sure.MatchString(in):
		return "Sure."
	case re_fine.MatchString(in):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}

}
