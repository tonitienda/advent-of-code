package output

import (
	"fmt"
	"io/ioutil"
	"path"
)

func WriteContents(year int, day int, filename string, contents string) {
	filepath := path.Join(fmt.Sprintf("/Users/toni/Projects/advent-of-code/golang/%d/day-%d", year, day), filename)
	err := ioutil.WriteFile(filepath, []byte(contents), 0644)

	if err != nil {
		panic(err)
	}

}
