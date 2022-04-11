package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Settings ------------------------------------------------
var xTerrarium int = 15
var yTerrarium int = 15

func main() {
	ant := constructRandomAnt()
	leaf := constructRandomLeaf()
	fmt.Println(ant, leaf)
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

// Ant -----------------------------------------------------
type Ant struct {
	position Position
}

func constructRandomAnt() Ant {
	x, y := randomNumber(0, xTerrarium), randomNumber(0, yTerrarium)
	antPosition := Position{x, y}
	return Ant{antPosition}
}

// Leaf ----------------------------------------------------
type Leaf struct {
	position Position
}

func constructRandomLeaf() Leaf {
	x, y := randomNumber(0, xTerrarium), randomNumber(0, yTerrarium)
	leafPosition := Position{x, y}
	return Leaf{leafPosition}
}
