package input

import (
	"io/ioutil"
	"path"
)

func GetContents(filename string) string {
	// TODO - See how to get current dir
	data, err := ioutil.ReadFile(path.Join("/Users/toni/Projects/advent-of-code/golang/2022/day-1", filename))

	if err != nil {
		panic(err)
	}

	return string(data)
}
