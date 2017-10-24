package raindrops

import "fmt"

const testVersion = 3

func Convert(n int) (st string) {
	st = ""
	if n%3 == 0 {
		st = "Pling"
	}
	if n%5 == 0 {
		st = st + "Plang"
	}
	if n%7 == 0 {
		st = st + "Plong"
	}
	if st == "" {
		st = fmt.Sprintf("%d", n)
	}
	return
}

// Don't forget the test program has a benchmark too.
// How fast does your Convert convert?
