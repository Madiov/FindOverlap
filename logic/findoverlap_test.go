package logic

import (
	"testing"

	"../db"
)

func TestFindOverlap(t *testing.T) {
	answers := []db.Rectangle{{X: 2, Y: 18, Width: 5, Height: 4}, {X: -1, Y: -1, Width: 5, Height: 4}}
	mainRect := db.Rectangle{X: 0, Y: 0, Width: 10, Height: 20}
	inputRect := []db.Rectangle{{X: 2, Y: 18, Width: 5, Height: 4}, {X: 12, Y: 18, Width: 5, Height: 4}, {X: -1, Y: -1, Width: 5, Height: 4}}
	rects := FindOverlap(mainRect, inputRect)
	if len(rects) != len(answers) {
		t.Errorf("Expected rects lenght in 2 but got %v", len(rects))
	}
	counter := 0
	for _, rect := range rects {
		for _, ans := range answers {
			if rect == ans {
				counter++
				continue
			}

		}
	}
	if counter != len(answers) {
		t.Errorf("wrong answers")
	}
}
