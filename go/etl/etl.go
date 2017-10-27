package etl

import "strings"

const testVersion = 1

func Transform(inp map[int][]string) (result map[string]int) {
	result = make(map[string]int)
	for key, val := range inp {
		for _, val2 := range val {
			result[strings.ToLower(val2)] = key
		}
	}
	return
}
