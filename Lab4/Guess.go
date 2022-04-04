package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	game()
}

func game() {
	fmt.Printf("Welcome to THE GAME!\nYou have to guess number between 1 and 500.\n")
	fmt.Println("You can type \"END\" if you want to resign.")
	numberToBeGuessed := randomNumber(1, 500)
	tries := 1
	for {
		guess := makeGuess()
		if end(guess, strconv.Itoa(numberToBeGuessed)) {
			break
		}
		intGuess, _ := strconv.Atoi(guess)
		if checkGuess(numberToBeGuessed, intGuess) {
			fmt.Printf("You've made in in %d tries", tries)
			break
		}
		tries++
	}

}

func randomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func makeGuess() string {
	var guess string
	fmt.Printf("Enter number:\t")
	fmt.Scanf("%s\n", &guess)
	return guess
}

func checkGuess(number, guess int) bool {
	if number == guess {
		fmt.Printf("\nCongratulations!\nYou've guessed the number \"%d\"\n", number)
		return true
	} else if guess > number {
		fmt.Printf("No, %d is too big!\n", guess)
		return false
	} else {
		fmt.Printf("No, %d is too small!\n", guess)
		return false
	}
}

func end(guess, correct string) bool {
	if strings.ToUpper(guess) == "END" {
		fmt.Printf("\nYou've resigned.\nThe number was \"%s\"\n", correct)
		return true
	}
	return false
}
