package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	game()
}

func game() {
	fmt.Printf("Welcome to THE GAME!\nYou have to guess number between 1 and 500.\n")
	numberToBeGuessed := randomNumber(1, 500)
	var guess string
	tries := 1
	for {
		fmt.Printf("Enter number:\t")
		fmt.Scanf("%s\n", &guess)
		intGuess, _ := strconv.Atoi(guess)
		if checkGuess(numberToBeGuessed, intGuess) {
			break
		}
		tries++
	}
	fmt.Printf("You've made in in %d tries", tries)
}

func randomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func checkGuess(number, guess int) bool {
	if number == guess {
		fmt.Printf("Congratulations!\nYou've guessed the number \"%d\"\n", number)
		return true
	} else if guess > number {
		fmt.Printf("No, %d is too big!\n", guess)
		return false
	} else {
		fmt.Printf("No, %d is too small!\n", guess)
		return false
	}
}
