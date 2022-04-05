package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	var blockString string
	highscore := CSVToList(LoadDataFromCSV())
	fmt.Println("**********************")
	fmt.Printf("Welcome to THE GAME!\nHighscore (TOP 3):\n")
	printScores(highscore, 3)
	fmt.Printf("**********************\n\nClick Enter to continue...")
	fmt.Scanf("%s\n", &blockString)
	fmt.Println()

	canSave := false

	for {
		nickname, gameScore, numGuessed := game()
		if gameScore != 0 {
			highscore = append(highscore, Score{
				nickname:      nickname,
				result:        gameScore,
				guessedNumber: numGuessed,
				date:          time.Now(),
			})
			highscore = sortScores(highscore)
			canSave = true
		}
		if !askIfPlayAgain() {
			printScores(highscore, 10)
			break
		}
	}
	if canSave {
		writeDataToCSV(scoresToCSVPreparedData(highscore))
	}
	fmt.Println("\nGoodbye! See you soon :)")
}

func game() (string, int, int) {
	fmt.Println("**********************")
	fmt.Printf("New Game!\nYou have to guess number between 1 and 500.\n")
	fmt.Println("You can type \"END\" if you want to resign.")
	numberToBeGuessed := randomNumber(1, 500)
	tries := 1
	for {
		guess := makeGuess()
		if end(guess, strconv.Itoa(numberToBeGuessed)) {
			return "", 0, numberToBeGuessed
		}
		intGuess, _ := strconv.Atoi(guess)
		if checkGuess(numberToBeGuessed, intGuess) {
			fmt.Printf("You've made in in %d tries\n", tries)
			nickname := ascForNickname()
			return nickname, tries, numberToBeGuessed
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
	nickname      string
	result        int
	guessedNumber int
	date          time.Time
}

func sortScores(scores []Score) []Score {
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].result < scores[j].result
	})
	return scores
}

func printScores(scores []Score, howMany int) {
	place := 1
	for index, score := range scores {
		y, m, d := score.date.Date()
		h, min, _ := score.date.Clock()
		fmt.Printf("%d. %s in %d tries (%d-%s-%d %d:%d)\n", place, score.nickname, score.result,
			y, m, d, h, min)
		place++
		if index == howMany-1 {
			break
		}
	}
}

func ascForNickname() string {
	var nickname string
	fmt.Printf("Enter your nickanme:\t")
	fmt.Scanf("%s\n", &nickname)
	fmt.Println()
	return nickname
}

func CSVToList(csvData [][]string) []Score {
	scoreList := make([]Score, 0)
	for _, dataScore := range csvData {
		dateTime, _ := time.Parse("2006-01-02 15:04", dataScore[3][:16])
		guessedNumber, _ := strconv.Atoi(dataScore[2])
		tries, _ := strconv.Atoi(dataScore[1])
		subScore := Score{
			nickname:      dataScore[0],
			result:        tries,
			guessedNumber: guessedNumber,
			date:          dateTime,
		}
		scoreList = append(scoreList, subScore)
	}
	return scoreList
}

func LoadDataFromCSV() [][]string {
	absPath, _ := filepath.Abs("results.csv")
	f, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func writeDataToCSV(data [][]string) {
	absPath, _ := filepath.Abs("results.csv")
	csvFile, err := os.OpenFile(absPath, os.O_WRONLY, 0777)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	err2 := csvwriter.WriteAll(data)
	if err2 != nil {
		log.Fatal(err2)
	}

	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(csvFile)

}

func scoresToCSVPreparedData(scores []Score) [][]string {
	preparedData := make([][]string, 0)

	for _, score := range scores {
		preparedData = append(preparedData, []string{
			score.nickname,
			strconv.Itoa(score.result),
			strconv.Itoa(score.guessedNumber),
			score.date.String(),
		})
	}

	return preparedData
}
