package main

import (
	"encoding/csv"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	data := CSVToList(LoadDataFromCSV())
	resultRangeList := printAnalysis(data)
	makePlot(resultRangeList)
}

func printAnalysis(data []Score) []ResultRange {
	fmt.Println("***********************")

	fmt.Println("-----OVERALL STATS-----")

	avgGuess, avgNumber, played := avgTries(data)
	trend := trendTries(data)
	fmt.Printf("There were %d games played\n", played)
	fmt.Printf("Number was guessed average in %.2f round\n", avgGuess)
	fmt.Printf("Median of tries to guess the number was %.2f \n", trend)
	fmt.Printf("Average number to be guessed was %d \n", avgNumber)

	fmt.Println("***********************")

	fmt.Println("-----RANGE  STATS-----")

	rangeAnalysis := mapScoresToResultRange(data, makeRangeList(1, 501, 20))
	for _, subRange := range rangeAnalysis {
		fmt.Printf("In range <%d, %d) number was guessed in average %.2f tries.\n",
			subRange.min, subRange.max, intListAvg(subRange.triesList))
	}
	min, max := getMinMaxFromResultRangeList(rangeAnalysis)
	fmt.Println("************************")
	fmt.Println("---HARDEST VS EASIEST---")
	fmt.Printf("Easiest to gues is range <%d, %d) with average %.2f tries.\n",
		min.min, min.max, intListAvg(min.triesList))
	fmt.Printf("Hardest to gues is range <%d, %d) with average %.2f tries.\n",
		max.min, max.max, intListAvg(max.triesList))

	fmt.Println("************************")

	return rangeAnalysis
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

func trendTries(data []Score) float64 {
	if len(data)%2 == 0 {
		return float64(data[len(data)/2].result)
	} else {
		return (float64(data[len(data)/2].result) + float64(data[(len(data)/2)+1].result)) / 2
	}
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

func intListAvg(list []int) float64 {
	sum := 0
	for _, elem := range list {
		sum += elem
	}
	return float64(sum) / float64(len(list))
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

func getMinMaxFromResultRangeList(rangeList []ResultRange) (ResultRange, ResultRange) {
	var max ResultRange
	var min ResultRange
	isSet := false

	for _, subRange := range rangeList {
		avg := intListAvg(subRange.triesList)

		if math.IsNaN(avg) {
			continue
		} else if !isSet {
			max = subRange
			min = subRange
			isSet = true
		} else {
			if avg > intListAvg(max.triesList) {
				max = subRange
			} else if avg < intListAvg(min.triesList) {
				min = subRange
			}
		}
	}

	return min, max
}

func makePlot(data []ResultRange) {

	p := plot.New()

	p.Title.Text = "Average number of guesses needed to unveil the number"
	p.X.Label.Text = "Number to be guessed"
	p.Y.Label.Text = "Number of tries to guess correctly"

	err := plotutil.AddLines(p,
		makePlotData(data),
	)
	if err != nil {
		panic(err)
	}

	if err := p.Save(5*vg.Inch, 5*vg.Inch, "rangeAnylysis.png"); err != nil {
		panic(err)
	}
}

func makePlotData(data []ResultRange) plotter.XYs {
	pts := make(plotter.XYs, len(data))
	for i, value := range data {
		avg := intListAvg(value.triesList)
		if math.IsNaN(avg) {
			continue
		}
		pts[i].X = float64((value.min + value.max) / 2)
		pts[i].Y = avg
	}
	ptsRefactored := make(plotter.XYs, 0)
	for _, points := range pts {
		if points.X != 0 {
			ptsRefactored = append(ptsRefactored, points)
		}
	}
	return ptsRefactored
}
