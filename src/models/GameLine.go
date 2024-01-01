package models

type GameLine struct {
	Line       string
	GameNumber int
	Subsets    []Set
}

func (gL GameLine) String() string {
	return gL.Line
}
