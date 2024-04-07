package main

import (
	"fmt"
)

func main() {
	var grid [16][16]bool
	// var pGrid *[16][16]bool = &grid

	// initialize grid
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = false
		}
	}
	grid = initGrid(grid)

	for i := 1; i > 0; i++ {
		print(grid)
		grid = process(grid)
	}
}

func initGrid(grid [16][16]bool) [16][16]bool {
	var num, x, y uint8

	fmt.Printf("How many cells to set: ")
	fmt.Scan(&num)

	fmt.Printf("What cells should be alive?\n")

	for i := uint8(0); i < num; i++ {
		fmt.Printf("Row: ")
		fmt.Scan(&x)

		fmt.Printf("Column: ")
		fmt.Scan(&y)

		grid[x][y] = true
	}

	return grid
}

func print(grid [16][16]bool) {
	for i := 0; i < len(grid); i++ {
		fmt.Printf("\n")
		for j := 0; j < len(grid[i]); j++ {
			switch grid[i][j] {
			case false:
				fmt.Printf("•")
			case true:
				fmt.Printf("#")
			}
		}
	}
	fmt.Printf("\n")
}

func process(grid [16][16]bool) [16][16]bool {
	var liveNeighbors uint8

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			liveNeighbors = 0

			// Check top neighbor
			if i-1 >= 0 && grid[i-1][j] {
				liveNeighbors++
			}

			// Check bottom neighbor
			if i+1 < len(grid) && grid[i+1][j] {
				liveNeighbors++
			}

			// Check left neighbor
			if j-1 >= 0 && grid[i][j-1] {
				liveNeighbors++
			}

			// Check right neighbor
			if j+1 < len(grid[i]) && grid[i][j+1] {
				liveNeighbors++
			}

			// Check top-left neighbor
			if i-1 >= 0 && j-1 >= 0 && grid[i-1][j-1] {
				liveNeighbors++
			}

			// Check top-right neighbor
			if i-1 >= 0 && j+1 < len(grid[i]) && grid[i-1][j+1] {
				liveNeighbors++
			}

			// Check bottom-left neighbor
			if i+1 < len(grid) && j-1 >= 0 && grid[i+1][j-1] {
				liveNeighbors++
			}

			// Check bottom-right neighbor
			if i+1 < len(grid) && j+1 < len(grid[i]) && grid[i+1][j+1] {
				liveNeighbors++
			}

			switch liveNeighbors {
			case 0, 1:
				grid[i][j] = false
			case 2:
				continue
			case 3:
				grid[i][j] = true
			default:
				grid[i][j] = false
			}
		}
	}

	return grid
}
