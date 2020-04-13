package models

import "fmt"

type Game struct {
	Round   int
	Wilds   map[int]string
	Players []Player
}

func (gm *Game) BuildScoreCard() string {
	people := ""
	scores := ""
	first := true
	for _, p := range gm.Players {
		// people = fmt.Sprint(people, p.Name, "\t| ")
		if first {
			people = fmt.Sprint(people, "|    ", p.Name, "\t| ")
			scores = fmt.Sprint(scores, "|\t", p.Score, "\t| ")
			first = false
		} else {
			people = fmt.Sprint(people, "   ", p.Name, "\t| ")
			scores = fmt.Sprint(scores, "\t", p.Score, "\t| ")
		}
	}
	return fmt.Sprint(people, "\n", scores)
}

func NewGame() *Game {
	return &Game{
		Round:   0,
		Wilds:   GetWildsMap(),
		Players: []Player{},
	}
}

func GetWildsMap() map[int]string {
	return map[int]string{
		1:  "Kings",
		2:  "Queens",
		3:  "Jacks",
		4:  "Tens",
		5:  "Nines",
		6:  "Eights",
		7:  "Sevens",
		8:  "Sixes",
		9:  "Fives",
		10: "Fours",
		11: "Threes",
		12: "Twos",
		13: "Aces",
	}
}
