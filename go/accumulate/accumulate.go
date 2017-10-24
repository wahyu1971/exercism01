package accumulate

const testVersion = 1

func Accumulate(lst []string, mfunc func(string) string) (result []string) {
	for _, i := range lst {
		result = append(result, mfunc(i))
	}
	return
}
