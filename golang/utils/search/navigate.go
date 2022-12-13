package search

func Get4Neighbours[T any](current [2]int, grid [][]T) [][2]int {
	candidates := [][2]int{}

	if current[0] > 0 {
		candidates = append(candidates, [2]int{current[0] - 1, current[1]})
	}

	if current[0] < len(grid)-1 {
		candidates = append(candidates, [2]int{current[0] + 1, current[1]})
	}

	if current[1] > 0 {
		candidates = append(candidates, [2]int{current[0], current[1] - 1})
	}

	if current[1] < len(grid[current[0]])-1 {
		candidates = append(candidates, [2]int{current[0], current[1] + 1})
	}

	return candidates
}
