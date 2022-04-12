package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Settings ------------------------------------------------
var xTerrarium int = 6
var yTerrarium int = 6

func main() {
	Simulation(10, true)
}

// Simulation ----------------------------------------------
func Simulation(iterations int, printEveryStep bool) {
	terrarium := constructTerrarium()
	ants := makeAntList(1)
	leaves := makeLeafList(5)
	terrarium.placeAntsAndLeaves(ants, leaves)
	terrarium.print()
	for i := 0; i < iterations; i++ {

		terrarium.moveAnts(ants, leaves)
		if printEveryStep {

			//cmd := exec.Command("cmd", "/c", "cls")
			//cmd.Stdout = os.Stdout
			//cmd.Run()

			fmt.Printf("\n---------------------\n\n")
			terrarium.print()
			//time.Sleep(1000 * time.Millisecond)
		}
	}
	if !printEveryStep {
		fmt.Printf("\n---------------------\n\n")
		terrarium.print()
	}
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

func (position *Position) move(x, y int) {
	position.x += x
	position.y += y
}

// Ant -----------------------------------------------------
type Ant struct {
	Position
	leaf *Leaf
}

func constructRandomAnt() Ant {
	x := randomNumber(0, xTerrarium-1)
	y := randomNumber(0, yTerrarium-1)
	antPosition := Position{x, y}
	return Ant{antPosition, nil}
}

func (ant *Ant) haveLeaf() bool {
	return ant.leaf != nil && !(ant.leaf.x == -1 && ant.leaf.y == -1)
}

func (ant *Ant) pickUpLeaf(leaf *Leaf) {
	if ant.haveLeaf() {
		ant.dropLeaf()
	}
	ant.leaf = leaf
	leaf.pickedUp = true
}

func (ant *Ant) dropLeaf() {
	ant.leaf.pickedUp = false
	ant.leaf = nil
}

func makeAntList(population int) []Ant {
	antList := make([]Ant, 0, population)
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
	x := randomNumber(0, xTerrarium-1)
	y := randomNumber(0, yTerrarium-1)
	leafPosition := Position{x, y}
	return Leaf{leafPosition, false}
}

func makeLeafList(population int) []Leaf {
	antList := make([]Leaf, 0, population)
	for i := 0; i < population; i++ {
		antList = append(antList, constructRandomLeaf())
	}
	return antList
}

func findLeafOnPositionInList(x, y int, list []Leaf) (int, *Leaf) {
	for index, leaf := range list {
		if leaf.x == x && leaf.y == y {
			return index, &list[index]
		}
	}
	return -1, nil
}

// Terrarium -----------------------------------------------
type Terrarium struct {
	/*
		0 - nothing
		1 - ant
		2 - leaf
		3 - ant caring leaf
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

func (trr *Terrarium) print() {
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

func (trr *Terrarium) placeAntsAndLeaves(ants []Ant, leaves []Leaf) {
	for antIndex, ant := range ants {
		if trr.matrix[ant.x][ant.y] != 0 {
			antTemp := constructRandomAnt()
			for trr.matrix[antTemp.x][antTemp.y] != 0 {
				antTemp = constructRandomAnt()
			}
			ants[antIndex] = antTemp
			ant = antTemp
		}
		trr.matrix[ant.x][ant.y] = 1
	}
	for leafIndex, leaf := range leaves {
		if trr.matrix[leaf.x][leaf.y] != 0 {
			leafTemp := constructRandomLeaf()
			for trr.matrix[leafTemp.x][leafTemp.y] != 0 {
				leafTemp = constructRandomLeaf()
			}
			leaves[leafIndex] = leafTemp
			leaf = leafTemp
		}
		trr.matrix[leaf.x][leaf.y] = 2
	}
}

func (trr *Terrarium) moveAnts(ants []Ant, leaves []Leaf) {
	for index, _ := range ants {
		ant := &ants[index]
		tries := 1
		for {
			// Get random vector
			xMove := randomNumber(-1, 1)
			yMove := randomNumber(-1, 1)
			// Reverse part of vector if move outside terrarium
			if xMove == 0 && yMove == 0 {
				continue
			}
			if ant.x+xMove < 0 || ant.x+xMove >= xTerrarium {
				xMove = -xMove
			}
			if ant.y+yMove < 0 || ant.y+yMove >= yTerrarium {
				yMove = -yMove
			}
			// Perform movement
			switch moved := trr.matrix[ant.x+xMove][ant.y+yMove]; moved {
			case 1, 3:
				// Case where filed was taken by another ant- let's assume it's trying three times, not moving otherwise
				tries++
				if tries != 4 {
					continue
				}
			case 0:
				// Case empty field
				trr.matrix[ant.x][ant.y] = 0
				ant.move(xMove, yMove)
				if ant.haveLeaf() {
					trr.matrix[ant.x][ant.y] = 3
					ant.leaf.move(xMove, yMove)
				} else {
					trr.matrix[ant.x][ant.y] = 1
				}
			case 2:
				// Case met leaf
				if ant.haveLeaf() {
					trr.matrix[ant.x][ant.y] = 2
					ant.leaf.x = ant.x
					ant.leaf.y = ant.y
				} else {
					trr.matrix[ant.x][ant.y] = 0
				}
				ant.move(xMove, yMove)
				_, leaf := findLeafOnPositionInList(ant.x, ant.y, leaves)
				ant.pickUpLeaf(leaf)
				//ant.leaf.move(xMove, yMove)
				trr.matrix[ant.x][ant.y] = 3
			}
			break
		}
	}
}
