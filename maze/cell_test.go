package maze

import "testing"

func TestCell(t *testing.T) {
	var c *Cell = NewCell(1, 2)
	if c.row != 1 {
		t.Error("Expected 1, got ", c.row)
	}
}

func TestCellLinkWithOtherCell(t *testing.T) {

	var c1 *Cell = NewCell(1, 2)
	var c2 *Cell = NewCell(3, 4)

	c1.Link(c2)

	if !c1.Linked(c2) {
		t.Error("Expected cells to be linked, but they are not")
	}
}

func TestCellLinkWithOtherCellInBothDirections(t *testing.T) {

	var c1 *Cell = NewCell(1, 2)
	var c2 *Cell = NewCell(3, 4)

	c1.LinkBidirectional(c2)

	if !c1.Linked(c2) {
		t.Error("Expected cells to be linked, but they are not")
	}

	if !c2.Linked(c1) {
		t.Error("Expected cells to be linked in both directions, but they are not")
	}
}

func TestCellUnlinkWithOtherCell(t *testing.T) {

	var c1 *Cell = NewCell(1, 2)
	var c2 *Cell = NewCell(3, 4)

	c1.Link(c2)
	c1.Unlink(c2)

	if c1.Linked(c2) {
		t.Error("Expected cells to be unlinked, but they are not")
	}
}

func TestCellUnlinkWithOtherCellBidirectional(t *testing.T) {

	var c1 *Cell = NewCell(1, 2)
	var c2 *Cell = NewCell(3, 4)

	c1.LinkBidirectional(c2)
	c1.UnlinkBidirectional(c2)

	if c1.Linked(c2) {
		t.Error("Expected cells to be unlinked, but they are not")
	}
	if c2.Linked(c1) {
		t.Error("Expected cells to be unlinked in both directions, but they are not")
	}
}

func TestCellLinkedTrueWhenLinked(t *testing.T) {

	var c1 *Cell = NewCell(1, 2)
	var c2 *Cell = NewCell(3, 4)

	c1.Link(c2)
	result := c1.Linked(c2)

	if !result {
		t.Error("Expected linked to be true, but it is not")
	}
}

func TestCellNeighboursReturnsNeighbours(t *testing.T) {
	var c1 *Cell = NewCell(2, 2)
	var north *Cell = NewCell(2, 1)

	c1.north = north

	var neighbours map[*Cell]struct{} = c1.Neighbours()
	_, ok := neighbours[north]
	if !ok {
		t.Error("Expected neighbours to contain the north cell, but it does not")
	}
}
