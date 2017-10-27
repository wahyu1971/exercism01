package queenattack

import "errors"

const testVersion = 2

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

// first check board validity & can queen attack ?
func CanQueenAttack(w string, b string) (attack bool, ok error) {
	err_msg := errors.New("board is invalid...")
	ok = nil
	attack = false
	// check board is valid ?
	switch {
	case w == b:
		ok = err_msg
	case w[1] > '8' || b[1] > '8' || w[1] < '1' || b[1] < '1':
		ok = err_msg
	case w[0] < 'a' || w[0] > 'h' || b[0] < 'a' || b[0] > 'h':
		ok = err_msg
	}
	if ok != nil {
		return
	}
	switch {
	case w[0] == b[0]:
		attack = true
	case w[1] == b[1]:
		attack = true
	case abs(int(w[1])-int(b[1])) == abs(int(w[0])-int(b[0])):
		attack = true
	}
	return
}
