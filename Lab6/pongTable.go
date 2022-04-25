package main

import (
	"fmt"
	"time"
)

func main() {
	racket1 := NewRacket("Player1", "Ping!")
	racket2 := NewRacket("Player2", "Pong!")
	racket1.setOpponent(racket2)
	racket2.setOpponent(racket1)
	racket2.getReady()
	racket1.serve()
	time.Sleep(20000 * time.Millisecond)
}

// Ball message passed between two players
type Ball struct {
	timesHit int
}

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

func (thisRacket *Racket) passTheBall() {
	var ball Ball
	for {
		ball = <-*thisRacket.receive
		ball.timesHit++
		fmt.Printf("(Hit %d) %s: %s\n", ball.timesHit, thisRacket.name, thisRacket.motto)
		time.Sleep(1000 * time.Millisecond)
		*thisRacket.opponent.receive <- ball
	}

}

func (thisRacket *Racket) getReady() {
	go thisRacket.passTheBall()
}

func (thisRacket *Racket) serve() {
	go thisRacket.passTheBall()
	*thisRacket.receive <- Ball{0}
}
