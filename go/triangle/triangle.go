package triangle

import (
	"math"
	"sort"
)

const testVersion = 3

type Kind int

const (
	NaT  Kind = iota // not a triangle
	Equ              // equilateral
	Iso              // isosceles
	Sca              // scalene
	ErrT             // should not happening
)

// return kind of triangle
func KindFromSides(a, b, c float64) Kind {
	arr := make([]float64, 3)
	arr[0] = a
	arr[1] = b
	arr[2] = c

	sort.Float64s(arr)

	// check triangular rule
	if (arr[2] > arr[1]+arr[0]) || arr[0] <= 0 || math.IsNaN(arr[0]) || math.IsInf(arr[1], 1) || math.IsInf(arr[1], -1) {
		return NaT
	}

	switch {
	case arr[2] == arr[1] && arr[1] == arr[0]:
		return Equ
	case arr[2] == arr[1] || arr[1] == arr[0]:
		return Iso
	case arr[2] != arr[1] && arr[1] != arr[0]:
		return Sca
	}

	// should never pass this
	return ErrT
}
