package maze

import "testing"

func TestCell(t *testing.T) {
	var c Cell = NewCell(1, 2)
	if c.Coordinates.Row != 1 {
		t.Error("Expected 1, got ", c.Coordinates.Row)
	}
}

func TestCellLinkWithOtherCell(t *testing.T) {

	var c1 Cell = NewCell(1, 2)
	var c2 Cell = NewCell(3, 4)
	c1.Link(c2)

	if !c1.Linked(c2) {
		t.Error("Expected cells to be linked, but they are not")
	}
}
