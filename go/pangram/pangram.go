package pangram

const testVersion = 2

/*
A-65
Z-90
a-97
z-122
*/

// to check if a string contain all characters at least once
func IsPangram(st string) bool {
	var pos byte
	mapBool := make(map[byte]bool)
	// init all to false
	for pos = 65; pos <= 90; pos++ {
		mapBool[pos] = false
	}
	for i := 0; i < len(st); i++ {
		// [a..z] change to [A..Z] -> uppercase
		if st[i] >= 97 && st[i] <= 122 {
			pos = st[i] - 32
		} else {
			pos = st[i]
		}
		// [A..Z] only
		if pos >= 65 && pos <= 90 {
			mapBool[pos] = true
		}
	}
	//fmt.Println(st, mapBool)
	for _, i := range mapBool {
		if i == false {
			return false
		}
	}
	return true
}
