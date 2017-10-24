package bob

import "regexp"

const testVersion = 3

func Hey(in string) string {
	re_sure := regexp.MustCompile("^.*?$")
	re_whoa := regexp.MustCompile("^[!@#$%^&*A-Z]+$")
	re_fine := regexp.MustCompile("^[!@#$%^&*A-Z]+$")

	return "WHOA!"
}
