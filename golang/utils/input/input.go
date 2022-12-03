package input

import (
	"fmt"
	"io/ioutil"
	"path"
)

func GetContents(year int, day int, filename string) string {
	// TODO - See how to get current dir
	data, err := ioutil.ReadFile(path.Join(fmt.Sprintf("/Users/toni/Projects/advent-of-code/golang/%d/day-%d", year, day), filename))

	if err != nil {
		panic(err)
	}

	return string(data)
}
