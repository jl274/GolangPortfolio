package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var blockString string
	fmt.Println("Welcome to the quiz app!")
	fmt.Println("You will hear 10 random questions. Each one cant grant you 1 point.")
	fmt.Println("Enter anything to continue!")
	fmt.Scanf("%s\n", &blockString)

	// -----

	questions := makeQuestionList(loadQuestions())

	points := 0
	for round := 1; round <= 10; round++ {
		var answer string
		questionNumber := randomNumber(0, len(questions)-1)
		roundQuestion := questions[questionNumber]
		fmt.Printf("%d: %s?\n", round, roundQuestion.questionText)
		fmt.Scanf("%s\n", &answer)
		if answer == roundQuestion.answer {
			points++
		}
		questions = removeFromSlice(questions, questionNumber)
	}
	fmt.Println("\nThe end!")
	fmt.Printf("***************\n*Score : %2.d/10*\n***************", points)
}

type question struct {
	questionText string
	answer       string
}

func loadQuestions() [][]string {
	absPath, _ := filepath.Abs("data/QuizData.csv")
	f, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func removeFromSlice(slice []question, s int) []question {
	return append(slice[:s], slice[s+1:]...)
}

func makeQuestionList(dataSet [][]string) []question {
	questionlist := make([]question, 0)
	for _, qAndA := range dataSet {
		newQuestion := question{questionText: qAndA[0], answer: qAndA[1]}
		questionlist = append(questionlist, newQuestion)
	}
	return questionlist
}

func randomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
