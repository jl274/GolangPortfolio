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
	Position
}

func constructRandomAnt() Ant {
	x, y := randomNumber(0, xTerrarium), randomNumber(0, yTerrarium)
	antPosition := Position{x, y}
	return Ant{antPosition}
}

func makeAntList(population int) []Ant {
	antList := make([]Ant, population, population)
	for i := 0; i < population; i++ {
		antList = append(antList, constructRandomAnt())
	}
	return antList
}

// Leaf ----------------------------------------------------
type Leaf struct {
	Position
}

func constructRandomLeaf() Leaf {
	x, y := randomNumber(0, xTerrarium), randomNumber(0, yTerrarium)
	leafPosition := Position{x, y}
	return Leaf{leafPosition}
}

func makeLeafList(population int) []Leaf {
	antList := make([]Leaf, population, population)
	for i := 0; i < population; i++ {
		antList = append(antList, constructRandomLeaf())
	}
	return antList
}

// Terrarium -----------------------------------------------
type Terrarium struct {
	/*
		0 - nothing
		1 - ant
		2 - leaf
	*/
	matrix [][]int
}

func constructTerrarium() Terrarium {
	terrariumMatrix := make([][]int, xTerrarium, xTerrarium)
	for i := 0; i < xTerrarium; i++ {
		terrariumMatrix[i] = make([]int, yTerrarium, yTerrarium)
	}
	return Terrarium{terrariumMatrix}
}
