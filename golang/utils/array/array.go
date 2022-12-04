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

func FlatMap[I any, O any](array []I, fn func(I) []O) []O {
	result := make([]O, 0)
	for _, t := range array {
		result = append(result, fn(t)...)
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
	length := len(array)
	result := make([]int, length)
	copy(result, array)
	sort.Ints(result)

	return result[length-n:]
}
