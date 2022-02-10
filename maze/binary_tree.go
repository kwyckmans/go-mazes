package maze

import (
	"math/rand"
)

func BinaryTree(g Grid) {
	for _, cell := range g.cells {
		neighbours := make([]*Cell, 0, 2)

		if cell.north != nil {
			neighbours = append(neighbours, cell.north)
		}

		if cell.east != nil {
			neighbours = append(neighbours, cell.east)
		}

		if len(neighbours) > 0 {
			index := rand.Intn(len(neighbours))
			neighbour := neighbours[index]

			if neighbour != nil {
				cell.LinkBidirectional(neighbour)
			}
		}
	}
}
