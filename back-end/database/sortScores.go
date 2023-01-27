package backend

import (
	"fmt"
	"sort"
)

func (s Scores) Less(other Scores) bool {
	return s.Pointsnumber > other.Pointsnumber
}

func SortScore(Scorestosort []Scores) []Scores {
	sort.Slice(Scorestosort, func(i, j int) bool {
		return Scorestosort[i].Pointsnumber > Scorestosort[j].Pointsnumber
	})
	fmt.Println(Scorestosort)
	return Scorestosort
}
