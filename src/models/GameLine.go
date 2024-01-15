package models

type GameLine struct {
	Line       string
	GameNumber int
	Subsets    []*Set
	// [red, green, blue]
	Maximums [3]int
	Minimums [3]int
}

func (gL GameLine) String() string {
	return gL.Line
}
