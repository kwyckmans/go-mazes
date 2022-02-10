package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/go-mazes/maze"
)

func emptyContents(maze.Cell) string {
	return " "
}

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

	start := g.Get(0, 0)
	distanceMap := maze.NewDistanceMap(start, start.GenerateDistanceMap())

	renderer := maze.NewDistanceRenderer(distanceMap)
	g.ContentRenderer = renderer
	fmt.Println(g)

	path := distanceMap.PathTo(g.Get(9, 9))
	renderer = maze.NewDistanceRenderer(path)
	g.ContentRenderer = renderer
	fmt.Println(g)

	fmt.Println("Calculating longest path")
	new_start, _ := distanceMap.Max()
	longestDistancesMap := maze.NewDistanceMap(new_start, new_start.GenerateDistanceMap())
	goal, _ := longestDistancesMap.Max()

	path = distanceMap.PathTo(goal)
	// TODO: Keep the renderer immutable, or just should I just allow to modify the distancemap?
	//   It's not a particularly heavy object
	renderer = maze.NewDistanceRenderer(path)
	g.ContentRenderer = renderer
	fmt.Println(g)
}
