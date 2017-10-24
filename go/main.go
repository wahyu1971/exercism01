package main

import (
	//"exercism01/go/pangram"
	"fmt"
)

func main() {
	aa := make(map[byte]bool)
	aa['A'] = true
	aa['B'] = true
	bb := "CDEF"
	for i := 0; i < len(bb); i++ {
		aa[bb[i]] = true
	}
	fmt.Println(aa)
	fmt.Println('A')
	fmt.Println('Z')
	fmt.Println('a')
	fmt.Println('z')
	for v, i := range aa {
		fmt.Println(v, i)
	}
}

/*
import (
	"exercism01/go/acronym"
	"fmt"
)

func main() {
	fmt.Println(acronym.Abbreviate("Complementary metal-oxide semiconductor"))
}


Complementary metal-oxide semiconductor
01234567890123456789012345678901234567890

*/
