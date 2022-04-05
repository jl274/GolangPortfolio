package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	data := CSVToList(LoadDataFromCSV())
	printAnalysis(data)

}

func printAnalysis(data []Score) {
	fmt.Println("***********************")
	avgGuess, avgNumber, played := avgTries(data)
	fmt.Printf("There were %d games played\n", played)
	fmt.Printf("Number was guessed average in %f round\n", avgGuess)
	fmt.Printf("Average number to be guessed was %d \n", avgNumber)
	fmt.Println("***********************")
}

func avgTries(data []Score) (float64, int, int) {
	tries := 0
	numberToBeGuessed := 0
	games := 0
	for _, game := range data {
		games++
		numberToBeGuessed += game.guessedNumber
		tries += game.result
	}
	return float64(tries) / float64(games), numberToBeGuessed / games, games
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

type Score struct {
	nickname      string
	result        int
	guessedNumber int
	date          time.Time
}

type ResultRange struct {
	min       int
	max       int
	triesList []int
}

func makeRangeList(min, max, number int) []ResultRange {
	step := max / number
	rangeList := make([]ResultRange, 0)

	for lap := 1; lap <= number; lap++ {
		rangeList = append(rangeList, ResultRange{
			min:       min + (step * (lap - 1)),
			max:       min + (step * lap),
			triesList: make([]int, 0),
		})
	}

	return rangeList
}

func mapScoresToResultRange(data []Score, rangeList []ResultRange) []ResultRange {

	for _, score := range data {
		number, tries := score.guessedNumber, score.result

		for subRangeIndex, subRange := range rangeList {
			if number >= subRange.min && number < subRange.max {
				rangeList[subRangeIndex].triesList = append(rangeList[subRangeIndex].triesList, tries)
			}
		}
	}

	return rangeList
}
