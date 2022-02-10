package maze

import "testing"

func TestGridNew(t *testing.T) {
	g := NewGrid(10, 10)

	if g.rows != 10 {
		t.Error("Expected 10, got ", g.rows)
	}
	if g.cols != 10 {
		t.Error("Expected 10, got ", g.cols)
	}

	cell := g.cells[coordinate{row: 0, col: 0}]

	if cell.north != nil {
		t.Error("Expected nil, got ", cell.north)
	}

	if cell.west != nil {
		t.Error("Expected nil, got ", cell.west)
	}

	if cell.east == nil {
		t.Error("Expected not nil, got ", cell.east)
	}

	if cell.south == nil {
		t.Error("Expected not nil, got ", cell.south)
	}

}
