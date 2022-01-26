package maze

import "math/rand"

func Sidewinder(g Grid) {
	for r := 0; r < g.rows; r++ {
		row := g.GetRow(r)

		run := make([]*Cell, 0, g.cols)

		for c := 0; c < g.cols; c++ {
			cell := row[c]
			run = append(run, cell)

			at_eastern_boundary := cell.east == nil
			at_northern_boundary := cell.north == nil

			should_close_out := at_eastern_boundary || (!at_northern_boundary && rand.Intn(2) == 0)

			if should_close_out {
				member := run[rand.Intn(len(run))]
				if member.north != nil {
					member.LinkBidirectional(member.north)
				}
				run = make([]*Cell, 0, g.cols)

			} else {
				cell.LinkBidirectional(cell.east)
			}
		}
	}
}
