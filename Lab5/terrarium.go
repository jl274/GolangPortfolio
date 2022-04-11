package main

import "fmt"

func main() {
	fmt.Println("ok")
}

type Position struct {
	x int
	y int
}

func (position Position) move(x, y int) {
	position.x += x
	position.y += y
}
