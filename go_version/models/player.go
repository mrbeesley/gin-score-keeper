package models

type Player struct {
	ID      int
	Name    string
	Score   int
	History map[int]int
}

func NewPlayer(id int, name string) Player {
	return Player{
		ID:      id,
		Name:    name,
		Score:   0,
		History: map[int]int{},
	}
}
