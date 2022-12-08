package input

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

func GetContents(year int, day int, filename string) string {
	// TODO - See how to get current dir
	data, err := ioutil.ReadFile(path.Join(fmt.Sprintf("/Users/toni/Projects/advent-of-code/golang/%d/day-%d", year, day), filename))

	if err != nil {
		panic(err)
	}

	return string(data)
}

func GetLines(year int, day int, filename string) []string {
	return strings.Split(GetContents(year, day, filename), "\n")
}

func Get2DMatrix[O int | float32](data string, splitRow, splitCol string, fn func(string) O) [][]O {
	rows := strings.Split(data, splitRow)

	result := make([][]O, len(rows))

	for i, row := range rows {
		cols := strings.Split(row, splitCol)
		result[i] = make([]O, len(cols))
		for j, col := range cols {
			result[i][j] = fn(col)
		}
	}
	return result
}
