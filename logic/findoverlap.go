package logic

import "../db"

type cordination []int

//FindOverlap finds the overlap between the main rectangle and input rectangles.
func FindOverlap(main db.Rectangle, input []db.Rectangle) []db.Rectangle {
	var results []db.Rectangle
	mainX := cordination{main.X, main.X + main.Width}
	mainY := cordination{main.Y, main.Y + main.Height}
	for _, rect := range input {
		secondryX := cordination{rect.X, rect.X + rect.Width}
		secondryY := cordination{rect.Y, rect.Y + rect.Height}
		if mainX[1] <= secondryX[0] || mainX[0] >= secondryX[1] {
			continue
		} else if mainY[1] <= secondryY[0] || mainY[0] >= secondryY[1] {
			continue
		} else {
			results = append(results, rect)
		}

	}
	return results

}
