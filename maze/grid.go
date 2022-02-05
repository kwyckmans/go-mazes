// Package maze contains various types and functions for creating and manipulating mazes.
package maze

import (
	"math/rand"
	"strings"
)

type ContentRenderer interface {
	ContentsOf(c *Cell) string
}

type EmptyRenderer struct {
}

func (e EmptyRenderer) ContentsOf(c *Cell) string {
	return " "
}

type coordinate struct {
	row, col int
}

// Grid is a 2D grid of cells that is used to represent a maze.
type Grid struct {
	rows, cols      int
	cells           map[coordinate]*Cell
	ContentRenderer ContentRenderer
}

// NewGrid creates a new grid with the given number of rows and columns.
// It also creates the cells for the grid, and assigns the correct neigbours.
//
// In the book the cells are initialized and connected in dedicated methods so
// they can be overridden. We don't have inheritance in Go, so no need to do that
// here. All grids should just implement a specific interface, but since wee have to
// reeimplement everything per class anyhow, no need to do that in three different mtehods.
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
		rows:            rows,
		cols:            cols,
		cells:           cells,
		ContentRenderer: EmptyRenderer{},
	}
}

// func (g Grid) contentsOf(cell Cell) string {
// 	return " "
// }

func (g Grid) String() string {
	output := "+" + strings.Repeat("---+", g.cols) + "\n"
	for row := 0; row < g.rows; row++ {
		top := "|"
		bottom := "+"

		for col := 0; col < g.cols; col++ {
			cell := g.cells[coordinate{row, col}]
			body := " " + g.ContentRenderer.ContentsOf(cell) + " "
			east_boundary := ""

			if cell.Linked(cell.east) {
				east_boundary = " "
			} else {
				east_boundary = "|"
			}

			top = top + body + east_boundary

			south_boundary := ""

			if cell.Linked(cell.south) {
				south_boundary = "   "
			} else {
				south_boundary = "---"
			}

			corner := "+"
			bottom = bottom + south_boundary + corner
		}
		output = output + top + "\n"
		output = output + bottom + "\n"
	}

	return output
}

func (g Grid) Get(row, col int) *Cell {
	return g.cells[coordinate{row, col}]
}

func (g Grid) RandomCell() *Cell {
	return g.cells[coordinate{rand.Intn(g.rows), rand.Intn(g.cols)}]
}

func (g Grid) Size() int {
	return g.rows * g.cols
}

func (g Grid) Rows() int {
	return g.rows
}

// GetRow returns all cells in the given row of the grid.
func (g Grid) GetRow(row int) map[int]*Cell {
	cells := make(map[int]*Cell)

	for col := 0; col < g.cols; col++ {
		cells[col] = g.Get(row, col)
	}

	return cells
}

func (g Grid) GetCells() map[coordinate]*Cell {
	return g.cells
}
