// Package maze contains various types and functions for creating and manipulating mazes.
package maze

import "fmt"

type coordinate struct {
	row, col int
}

// Grid is a 2D grid of cells that is use to represent a maze.
type Grid struct {
	rows, cols int
	cells      map[coordinate]*Cell
}

// NewGrid creates a new grid with the given number of rows and columns.
// It also creates the cells for the grid, and assigns the correct neigbours.
func NewGrid(rows, cols int) *Grid {
	cells := make(map[coordinate]*Cell)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			cells[coordinate{row, col}] = NewCell(row, col)
		}
	}

	for index, cell := range cells {
		cell.north = cells[coordinate{index.row - 1, index.col}]
		cell.south = cells[coordinate{index.row + 1, index.col}]
		cell.east = cells[coordinate{index.row, index.col + 1}]
		cell.west = cells[coordinate{index.row, index.col - 1}]
	}

	return &Grid{
		rows:  rows,
		cols:  cols,
		cells: cells,
	}
}

func (g Grid) String() string {
	return fmt.Sprintf("Grid{rows: %d, cols: %d, cells: %v}", g.rows, g.cols, g.cells)
}
