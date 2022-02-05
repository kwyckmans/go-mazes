package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

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

	maze.BinaryTree(*g)
	fmt.Print(g)

	testGrid := maze.NewGrid(3, 3)

	origin := testGrid.Get(0, 0)
	connected := testGrid.Get(1, 0)
	origin.LinkBidirectional(connected)

	fmt.Println(testGrid)

	sidewinderGrid := maze.NewGrid(50, 50)
	maze.Sidewinder(*sidewinderGrid)
	fmt.Println(sidewinderGrid)

	f, err := os.Create("test-image.png")
	if err != nil {
		log.Fatalln("Failed to create image")
	}

	// TODO: read up on defer
	defer f.Close()

	img := maze.ToPNG(10, 500, 500, 1, *sidewinderGrid)

	err = png.Encode(f, img)
	if err != nil {
		log.Fatalln("Failed to convert image to png")
	}
	log.Println("Generated image")

	// cells := g.GetCells()

	// for k, v := range cells {
	// 	fmt.Println(k, v)
	// }
	// var otherC maze.Cell = maze.Cell{
	// 	Coordinates: maze.Coordinates{Row: 1, Col: 2},
	// 	Links:       make(map[maze.Coordinates]struct{}),
	// }

	// fmt.Println(otherC)
}
