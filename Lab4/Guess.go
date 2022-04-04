package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	highscore := make([]Score, 0)
	for {
		nickname, gameScore := game()
		if gameScore != 0 {
			highscore = append(highscore, Score{
				nickname: nickname,
				result:   gameScore,
			})
			highscore = sortScores(highscore)
		}
		if !askIfPlayAgain() {
			printScores(highscore)
			break
		}
	}

}

func game() (string, int) {
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
			nickname := ascForNickname()
			return nickname, tries
		}
		tries++
	}
	return "", 0

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

func askIfPlayAgain() bool {
	var answer string
	fmt.Printf("\nPlay again? [YES/NO]\n")
	fmt.Scanf("%s\n", &answer)
	answer = strings.ToUpper(answer)
	for {
		switch answer {

		case "YES", "Y":
			fmt.Println()
			return true

		case "NO", "N":
			fmt.Println()
			return false

		default:
			fmt.Println("Invalid option!")
		}
	}
}

type Score struct {
	nickname string
	result   int
}

func sortScores(scores []Score) []Score {
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].result < scores[j].result
	})
	return scores
}

func printScores(scores []Score) {
	place := 1
	for _, score := range scores {
		fmt.Printf("%d. %s in %d tries\n", place, score.nickname, score.result)
		place++
	}
}

func ascForNickname() string {
	var nickname string
	fmt.Printf("Enter your nickanme:\t")
	fmt.Scanf("%s\n", &nickname)
	fmt.Println()
	return nickname
}
