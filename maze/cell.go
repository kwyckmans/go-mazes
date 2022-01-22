package maze

import "fmt"

// Cell is a single cell in a maze.
type Cell struct {
	row, col                 int
	north, east, south, west *Cell
	links                    map[*Cell]struct{}
}

func NewCell(row, col int) *Cell {
	return &Cell{
		row:   row,
		col:   col,
		links: make(map[*Cell]struct{}),
	}
}

func (c Cell) String() string {
	return fmt.Sprintf("%d, %d", c.row, c.col)
}

// Link adds a link between the cell c and cell c2 to c.
func (c *Cell) Link(c2 *Cell) {
	c.links[c2] = struct{}{}
}

// Link adds a link between the cell c and cell c2 to both c and c2.
func (c *Cell) LinkBidirectional(c2 *Cell) {
	c.Link(c2)
	c2.Link(c)
}

// Unlink removes the link between the cell c and cell c2.
// It does not remove the link from cell c2 to cell c, if present.
func (c *Cell) Unlink(c2 *Cell) {
	delete(c.links, c2)
}

// UnlinkBidirectional removes the link between the two cells.
func (c *Cell) UnlinkBidirectional(c2 *Cell) {
	c.Unlink(c2)
	c2.Unlink(c)
}

// Linked returns true if the cell is linked to the given cell.
func (c *Cell) Linked(c2 *Cell) bool {
	_, ok := c.links[c2]
	return ok
}

func (c *Cell) Links() []*Cell {
	var links []*Cell
	for link := range c.links {
		links = append(links, link)
	}
	return links
}

func (c *Cell) Neighbours() map[*Cell]struct{} {
	var neighbours map[*Cell]struct{} = make(map[*Cell]struct{})

	if c.north != nil {
		neighbours[c.north] = struct{}{}
	}
	if c.east != nil {
		neighbours[c.east] = struct{}{}
	}
	if c.south != nil {
		neighbours[c.south] = struct{}{}
	}
	if c.west != nil {
		neighbours[c.west] = struct{}{}
	}

	return neighbours
}
