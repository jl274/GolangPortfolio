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
	trr := constructTerrarium()
	trr.matrix[0][1] = 2
	trr.matrix[2][10] = 1
	trr.matrix[7][11] = 2
	trr.print()
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
	leaf Leaf
}

func constructRandomAnt() Ant {
	x, y := randomNumber(0, xTerrarium), randomNumber(0, yTerrarium)
	antPosition := Position{x, y}
	return Ant{antPosition, Leaf{Position{-1, -1}, false}}
}

func (ant Ant) haveLeaf() bool {
	return !(ant.leaf.x == -1 && ant.leaf.y == -1)
}

func (ant Ant) pickUpLeaf(leaf Leaf) {
	if ant.haveLeaf() {
		ant.dropLeaf()
	}
	ant.leaf = leaf
	leaf.pickedUp = true
}

func (ant Ant) dropLeaf() {
	ant.leaf.pickedUp = false
	ant.leaf = Leaf{Position{-1, -1}, false}
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
	pickedUp bool
}

func constructRandomLeaf() Leaf {
	x, y := randomNumber(0, xTerrarium), randomNumber(0, yTerrarium)
	leafPosition := Position{x, y}
	return Leaf{leafPosition, false}
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

func (trr Terrarium) print() {
	for _, row := range trr.matrix {
		for _, cell := range row {
			switch cell {
			case 0:
				fmt.Printf("🔲")
			case 1:
				fmt.Printf("🐜")
			case 2:
				fmt.Printf("🍃")
			case 3:
				fmt.Printf("🐞")
			}
		}
		fmt.Println("")
	}
}
