package secret

import (
	"container/list"
)

const testVersion = 2

func Handshake(code uint) (result []string) {
	reverse := false
	ll := list.New()
	for i := 0; i < 5; i++ {
		if code == 0 {
			break
		}
		if code&1 == 1 {
			switch i {
			case 0:
				ll.PushBack("wink")
			case 1:
				ll.PushBack("double blink")
			case 2:
				ll.PushBack("close your eyes")
			case 3:
				ll.PushBack("jump")
			case 4:
				reverse = true
			}
		}
		code = code >> 1
	}
	i := 0
	result = make([]string, ll.Len())
	if reverse == false {
		for e := ll.Front(); e != nil; e = e.Next() {
			result[i] = e.Value.(string)
			i++
		}
		return

	}
	// reverse == true
	for e := ll.Back(); e != nil; e = e.Prev() {
		result[i] = e.Value.(string)
		i++
	}
	return
}
