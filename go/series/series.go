package series

const testVersion = 2

func All(n int, s string) (result []string) {
	for i := 0; i < len(s); i++ {
		if i+n > len(s) {
			return
		}
		result = append(result, s[i:i+n])
	}
	return
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return s[0:len(s)], false
	}
	return s[0:n], true
}

func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return s[0:len(s)]
	}
	return s[0:n]
}
