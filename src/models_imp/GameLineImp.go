package models_imp

import (
	"advent_of_code/2023/models"
	"math"
)

func NewGameLine(line string) *models.GameLine {
	gL := models.GameLine{Line: line, Maximums: [3]int{0, 0, 0}, Minimums: [3]int{math.MaxInt, math.MaxInt, math.MaxInt}}
	return &gL
}
