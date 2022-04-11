package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("ok")
}

// Random number -------------------------------------------
func randomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// Position ------------------------------------------------
type Position struct {
	x int
	y int
}

func (position Position) move(x, y int) {
	position.x += x
	position.y += y
}
