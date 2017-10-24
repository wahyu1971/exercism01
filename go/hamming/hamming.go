package hamming

const testVersion = 6

type HammingError struct {
}

func (e HammingError) Error() string {
	return "Hamming error - two strings have different size.."
}

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		e := new(HammingError)
		return -1, e
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
