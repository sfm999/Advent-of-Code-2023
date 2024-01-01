package models_imp

import (
	"advent_of_code/2023/models"
)

func NewGameLine(line string) *models.GameLine {
	gL := models.GameLine{Line: line}
	return &gL
}
