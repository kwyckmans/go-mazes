package maze

import (
	"github.com/martinlindhe/base36"
)

type DistanceMap struct {
	distances map[*Cell]int
	root      *Cell
}

func NewDistanceMap(root *Cell, distanceMap map[*Cell]int) DistanceMap {
	return DistanceMap{
		distances: distanceMap,
		root:      root,
	}
}

// PathTo finds a path from the root to the given cell.
func (d DistanceMap) PathTo(goal *Cell) DistanceMap {
	current := goal
	breadcrumbs := make(map[*Cell]int)
	// breadcrumbs[d.root] = 0
	breadcrumbs[current] = d.distances[current]

	for current != d.root {
		for neighbor := range current.links {
			if d.distances[neighbor] < d.distances[current] {
				breadcrumbs[neighbor] = d.distances[current]
				current = neighbor
				break
			}
		}
	}

	return DistanceMap{
		distances: breadcrumbs,
		root:      d.root,
	}
}

// DistanceRenderer is a Renderer that takes a DistanceMap and renders it as an ASCII grid.
//
// TODO: See if we can clean-up the connection between DistanceMap, Cell, Grid and renderer. It's a bit odd that
//    the renderer just prints a distancemap. So, we have cell to generate a distancemap, that map we attach to a renderer
//    and finally we assign the renderer to a grid. It feels a bit convoluted.
//
//    Maybe I can just pass in a renderer as as a parameter a render method in grid. Or I can have renderer accept a grid,
//    instead of the other way around. Ideally the grid doesn't know about the distancemap, since for now it doesn't care.
//    I'd like to just print arbitrary distance maps. The thing is, obviously it's the grids responsability to print itself,
//    the issue just being that we don't actually print the grid, we just print the distancemap.
type DistanceRenderer struct {
	distanceMap DistanceMap
}

func NewDistanceRenderer(distanceMap DistanceMap) DistanceRenderer {
	return DistanceRenderer{
		distanceMap: distanceMap,
	}
}

func (d DistanceRenderer) ContentsOf(c *Cell) string {
	distance, ok := d.distanceMap.distances[c]
	if ok {
		return base36.Encode(uint64(distance))
	} else {
		return " "
	}
}
