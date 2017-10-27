package grains

import "errors"

const testVersion = 1

func Square(n int) (result uint64, ok error) {
	err_msg := errors.New("cannot solve the problem")
	result = 0
	ok = nil
	if n < 1 || n > 64 {
		ok = err_msg
		return
	}
	result = 1 << uint8(n-1)
	return
}

func Total() (result uint64) {
	result = 0
	for i := 1; i <= 64; i++ {
		sq, _ := Square(i)
		result = result + sq
	}
	return
}
