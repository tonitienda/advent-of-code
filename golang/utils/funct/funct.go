package funct

func GetValue[I any](value I, err error) I {
	if err != nil {
		panic((err))
	}

	return value
}

func GetValueB[I any](value I, isOk bool) I {
	// if !isOk {
	// 	panic("Value is not ok")
	// }

	return value
}
