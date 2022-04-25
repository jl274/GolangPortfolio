package main

import "fmt"

func main() {
	fmt.Println("Test")
	racket1 := NewRacket("Player1", "Ping!")
	racket2 := NewRacket("Player2", "Pong!")
	racket1.setOpponent(racket2)
	racket2.setOpponent(racket1)
}

// Ball message passed between two players
type Ball string

// Racket Struct defining player
type Racket struct {
	name     string
	motto    string
	receive  *chan Ball
	opponent *Racket
}

func NewRacket(name, motto string) *Racket {
	receive := make(chan Ball, 1)
	racket := Racket{name, motto, &receive, nil}
	return &racket
}

func (thisRacket *Racket) setOpponent(opponent *Racket) {
	thisRacket.opponent = opponent
}

//func (*Racket) letsPlay(c chan <- Ball){
//
//}
