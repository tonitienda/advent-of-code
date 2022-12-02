package funct

func GetValue[I any](value I, err error) I {
	if err != nil {
		panic((err))
	}

	return value
}
