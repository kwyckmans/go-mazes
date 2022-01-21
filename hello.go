package main

import (
	"fmt"

	"github.com/go-mazes/maze"
)

func main() {
	var c maze.Cell = maze.NewCell(1, 2)
	fmt.Println(c)

	var otherC maze.Cell = maze.Cell{
		Coordinates: maze.Coordinates{Row: 1, Col: 2},
		Links:       make(map[maze.Coordinates]struct{}),
	}

	fmt.Println(otherC)
}
