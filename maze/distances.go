package maze

import "github.com/martinlindhe/base36"

type DistanceRenderer struct {
	Distances map[*Cell]int
}

func (d DistanceRenderer) ContentsOf(c *Cell) string {
	distance, ok := d.Distances[c]
	if ok {
		return base36.Encode(uint64(distance))
	} else {
		return " "
	}
}
