package maze

import (
	"image"
	"image/color"
	"image/draw"
)

func ToPNG(cell_size, width, height, thickness int, grid Grid) *image.RGBA {

	background := color.RGBA{
		255, 255, 255, 255,
	}

	wall := color.RGBA{
		0, 0, 0, 255,
	}

	m := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(m, m.Bounds(), &image.Uniform{background}, image.Point{}, draw.Src)

	for _, cell := range grid.GetCells() {
		// TODO: Double check if this is correct.
		// 		I think in my case that x represents row and y col
		//      though while writing I realise I may be mixing up
		//      concepts. The x and y here are unrelated to that?
		x1 := cell.col * cell_size       // 0
		y1 := cell.row * cell_size       // 0
		x2 := (cell.col + 1) * cell_size // 10
		y2 := (cell.row + 1) * cell_size // 10
		// log.Printf("Checking %d, %d", coord.col, coord.row)
		if cell.north == nil {
			line := image.Rect(x1, y1, x2, y1+thickness)
			// log.Printf("Drawing from %d, %d to %d, %d", x1, y1, x2, y1+thickness)
			draw.Draw(m, line, &image.Uniform{wall}, image.Point{}, draw.Src)
		}

		if cell.west == nil {
			line := image.Rect(x1, y1, x1+thickness, y2)
			draw.Draw(m, line, &image.Uniform{wall}, image.Point{}, draw.Src)
		}

		if !cell.Linked(cell.east) {
			line := image.Rect(x2-thickness, y1-thickness, x2, y2)
			draw.Draw(m, line, &image.Uniform{wall}, image.Point{}, draw.Src)
		}

		if !cell.Linked(cell.south) {
			line := image.Rect(x1, y2, x2, y2-thickness)
			// log.Printf("Drawing from %d, %d to %d, %d", x1, y1, x2, y1+thickness)
			draw.Draw(m, line, &image.Uniform{wall}, image.Point{}, draw.Src)
		}
	}

	return m
}
