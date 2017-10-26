package pascal

const testVersion = 1

func solve_triangle(arr []int) (result []int) {
	result = make([]int, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		switch i {
		case 0:
			result[0] = arr[0]
			continue
		case 1:
			result[len(result)-1] = arr[len(arr)-1]
			result[i] = arr[i-1] + arr[i]
		default:
			result[i] = arr[i-1] + arr[i]
		}
	}
	return
}

func rec_triangle(n int) []int {
	switch n {
	case 1:
		return []int{1}
	case 2:
		return []int{1, 1}
	default:
		return solve_triangle(rec_triangle(n - 1))
	}
	// should never reach this point..
	//return []int{}
}

func Triangle(n int) (result [][]int) {
	result = make([][]int, n)
	for i := 1; i <= n; i++ {
		tmp := rec_triangle(i)
		result[i-1] = make([]int, len(tmp))
		copy(result[i-1], tmp)
	}
	return
}
