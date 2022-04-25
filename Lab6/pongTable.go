package main

import "fmt"

func main() {
	fmt.Println("Test")
	//tennisTable := make(chan Ball, 1)
	//racket1 := Racket{"PLayer1", "Ping!", tennisTable}
	//racket2 := Racket{"PLayer2", "Pong!", tennisTable}
	//for i := range [5]int{}{
	//
	//}
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

//func (*Racket) letsPlay(c chan <- Ball){
//
//}
