package main

import (
	"fmt"
)

func main() {
	var grid [16][16]bool
	// var pGrid *[16][16]bool = &grid

	// initialize grid to false
	for i := 0; i <= len(grid); i++ {
		for j := 0; j <= len(grid[i]); j++ {
			grid[i][j] = false
		}
	}

	for i := 1; i > 0; i++ {
		print(grid)
	}
}

func print(grid [16][16]bool) {
	for i := 0; i <= len(grid); i++ {
		fmt.Printf("\n")
		for j := 0; j <= len(grid[i]); j++ {
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

func process(grid [16][16]bool)
