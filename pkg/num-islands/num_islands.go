package num_islands

// IslandsCounter provides islandCount interface
type IslandsCounter interface {
	traverseIsland(grid [][]byte, row int, col int)
	NumIslands(grid [][]byte) int
}

// islandsCount implements NumIslands
type islandsCount struct {
}

// NumIslands returns number of islands on the map
func (i *islandsCount) NumIslands(grid [][]byte) int {
	count := 0
	for k := 0; k < len(grid); k++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[k][j] == '1' {
				count++
				i.traverseIsland(grid, k, j)
			}
		}
	}
	return count
}

// traverseIsland traverse island and set all island bytes to `0`
func (i *islandsCount) traverseIsland(grid [][]byte, row int, col int) {
	grid[row][col] = 0
	if row+1 < len(grid) && grid[row+1][col] == '1' {
		i.traverseIsland(grid, row+1, col)
	}
	if col+1 < len(grid[0]) && grid[row][col+1] == '1' {
		i.traverseIsland(grid, row, col+1)
	}
	if row-1 >= 0 && grid[row-1][col] == '1' {
		i.traverseIsland(grid, row-1, col)
	}
	if col-1 >= 0 && grid[row][col-1] == '1' {
		i.traverseIsland(grid, row, col-1)
	}
}

// NewIslandCount returns new islandCount
func NewIslandsCounter() IslandsCounter {
	return &islandsCount{}
}
