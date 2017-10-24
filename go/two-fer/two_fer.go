// Package twofer implement function ShareWith that return 'twofer' with name in it..
package twofer

// return string 'twofer' with name in it..
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	} else {
		return "One for " + name + ", one for me."
	}
}
