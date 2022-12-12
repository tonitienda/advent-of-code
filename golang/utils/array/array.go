package array

func Map[I any, O any](array []I, fn func(I) O) []O {
	result := make([]O, len(array))
	for i, t := range array {
		result[i] = fn(t)
	}
	return result
}

// Golang does not allow Overloading of methods so we need a different name than Map for this
// Goland does not support variadic functions with different types so we need to have explicit
// number of arguments with the intermediate generics
func Map2[I any, T1 any, O any](array []I, fni func(I) T1, fno func(T1) O) []O {
	result := make([]O, len(array))
	for i, t := range array {
		result[i] = fno(fni(t))
	}
	return result
}

func Map3[I any, T1 any, T2 any, O any](array []I, fni func(I) T1, fn1 func(T1) T2, fno func(T2) O) []O {
	result := make([]O, len(array))
	for i, t := range array {
		result[i] = fno(fn1(fni(t)))
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

func MaxN[T int | int64 | uint64](array []T, n T) []T {
	length := len(array)
	result := make([]T, length)

	for _, item := range array {
		if result[0] < item {
			result[1] = result[0]
			result[0] = item
		} else {
			if result[1] < item {
				result[1] = item
			}
		}

	}

	return result
}
