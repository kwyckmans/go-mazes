package main

import (
	"fmt"

	"github.com/go-mazes/maze"
)

func main() {
	c := maze.NewCell(1, 2)
	fmt.Println(c)
	// carray := [10][10]*maze.Cell{}

	// for i := range carray {
	// 	for j := range carray[i] {
	// 		fmt.Println(carray[i][j])
	// 	}
	// }

	g := maze.NewGrid(10, 10)

	fmt.Println(g)

	// var otherC maze.Cell = maze.Cell{
	// 	Coordinates: maze.Coordinates{Row: 1, Col: 2},
	// 	Links:       make(map[maze.Coordinates]struct{}),
	// }

	// fmt.Println(otherC)
}
