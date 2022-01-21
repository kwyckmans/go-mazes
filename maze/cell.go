package maze

import "fmt"

type Coordinates struct {
	Row, Col int
}

func (c Coordinates) String() string {
	return fmt.Sprintf("(%d, %d)", c.Row, c.Col)
}

// Cell is a single cell in a maze.
type Cell struct {
	// Row, Col    int
	Coordinates Coordinates
	// North, East, South, West *Cell
	// Links                    map[*Cell]bool
	Links map[Coordinates]struct{}
}

func NewCell(row int, col int) Cell {
	coords := Coordinates{row, col}

	return Cell{
		Coordinates: coords,
		Links:       make(map[Coordinates]struct{}),
	}
}

func (c Cell) String() string {
	return fmt.Sprintf("%d", c.Coordinates)
}

// Link adds a link between the cell c and cell c2 to c.
func (c Cell) Link(c2 Cell) {
	c.Links[c2.Coordinates] = struct{}{}
}

// Link adds a link between the cell c and cell c2 to both c and c2.
func (c Cell) LinkBidirectional(c2 Cell) {
	c.Link(c2)
	c2.Link(c)
}

// Unlink removes the link between the cell c and cell c2.
// It does not remove the link from cell c2 to cell c, if present.
func (c Cell) Unlink(c2 Cell) {
	delete(c.Links, c2.Coordinates)
}

// UnlinkBidirectional removes the link between the two cells.
func (c Cell) UnlinkBidirectional(c2 Cell) {
	c.Unlink(c2)
	c2.Unlink(c)
}

// Linked returns true if the cell is linked to the given cell.
func (c Cell) Linked(c2 Cell) bool {
	_, ok := c.Links[c2.Coordinates]
	return ok
}
