package array

import (
	"sort"
)

func SayHello() string {
	return "Hi from package dir1"
}

func Map[I any, O any](array []I, fn func(I) O) []O {
	result := make([]O, len(array))
	for i, t := range array {
		result[i] = fn(t)
	}
	return result
}

func Sum[I int | float32](array []I) I {
	result := I(0)
	for _, t := range array {
		result += t
	}
	return result
}

func Max[I int | float32](array []I) I {
	result := I(0)
	for _, t := range array {
		if result < t {
			result = t
		}
	}
	return result
}

func MaxN(array []int, n int) []int {
	result := make([]int, n)
	copy(result, array[:n])
	sort.Ints(result)

	for _, t := range array {
		if result[0] < t {
			result[0] = t
			sort.Ints(result[:])
		}
	}
	return result
}
