package maze

import (
	"fmt"
	"math/rand"
)

func BinaryTree(g Grid) {
	for _, cell := range g.cells {
		fmt.Println(cell)
		neighbours := make([]*Cell, 0, 2)

		if cell.north != nil {
			neighbours = append(neighbours, cell.north)
		}

		if cell.east != nil {
			neighbours = append(neighbours, cell.east)
		}

		fmt.Println(neighbours)
		if len(neighbours) > 0 {
			index := rand.Intn(len(neighbours))
			neighbour := neighbours[index]

			if neighbour != nil {
				cell.LinkBidirectional(neighbour)
			}
		}
	}
}
