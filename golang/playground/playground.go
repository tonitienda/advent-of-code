package main

import (
	"fmt"
	"strconv"
)

func Map[I any, O any](array []I, fn func(I) O) []O {
	result := make([]O, len(array))
	for i, t := range array {
		result[i] = fn(t)
	}
	return result
}

func Increment(i int) int {
	return i + 1
}

func Decrement(i int) int {
	return i - 1
}

func MapN[I any, X any, O any](array []I, fni func(I) X, fns []func(X) X, fno func(X) O) []O {
	result := make([]O, len(array))
	for i, t := range array {
		t2 := fni(t)
		for _, fn := range fns {
			t2 = fn(t2)
		}

		result[i] = fno(t2)
	}
	return result
}

func IncrementStr(str string) string {
	return str + "1"
}

func StrToInt(str string) int {

	value, err := strconv.Atoi(str)
	if err != nil {
		panic((err))
	}

	return value
}

func main() {
	data := []int{1, 2, 3, 4, 5}

	fmt.Println(data)

	fmt.Println(Map(data, Increment))
	fmt.Println(MapN(data, Increment, []func(int) int{Increment}, Decrement))

	fmt.Println(MapN(data, strconv.Itoa, []func(string) string{IncrementStr}, StrToInt))

}
