package models_imp

import (
	"advent_of_code/2023/models"
)

func NewColour(count int, colour string) *models.Colour {
	c := models.Colour{Count: count, Colour: colour}
	return &c
}
