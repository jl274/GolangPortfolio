package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	textFlag := flag.String("file", "text.txt", "Choose txt file")
	workersFlag := flag.Int("w", 5, "Choose number of threads to handle the task")
	flag.Parse()
	witcherCounter := NewCounter()
	result := witcherCounter.count(*textFlag, *workersFlag)
	printSortResult(result)

}

func printSortResult(scores counterMap) {

	names := make([]string, 0, len(scores))
	for name := range scores {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool {
		return scores[names[i]] > scores[names[j]]
	})
	for _, name := range names {
		fmt.Printf("%s: %d times\n", name, scores[name])
	}
}

// Placeholder for map with counted words
type counterMap map[string]int

// Counter holds total record of words in file, and receives new information from Reader
type Counter struct {
	wordsMap counterMap
	received int
	wordsBox chan counterMap
}

func NewCounter() Counter {
	wMap := make(counterMap)
	wChan := make(chan counterMap)
	return Counter{wMap, 0, wChan}
}

func (c *Counter) MakeReader() Reader {
	return Reader{c, false}
}

// Reader receives lines, counts word in it and sends it back to Counter
type Reader struct {
	supervisor *Counter
	working    bool
}

func (r *Reader) read(line string) {
	line = strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(line, ".", ""),
					",", ""),
				"-", ""),
			"?", ""),
		"'", "")
	line = strings.ToLower(line)
	lineArray := strings.Fields(line)
	resultMap := make(counterMap)
	for _, word := range lineArray {
		resultMap[word]++
	}
	r.supervisor.wordsBox <- resultMap
	r.working = false
}

// File to string array
func (c *Counter) readFile(file string) []string {
	absPath, _ := filepath.Abs(file)
	f, err := os.ReadFile(absPath)
	if err != nil {
		err.Error()
	}
	return strings.Split(string(f), "\n")
}

// Concurrency loop
func (c *Counter) count(file string, workersNumber int) counterMap {
	wordsArray := c.readFile(file)
	maxLen := len(wordsArray)
	workers := make([]Reader, 0, workersNumber)
	for i := 0; i < workersNumber; i++ {
		workers = append(workers, c.MakeReader())
	}
	for {
		if len(wordsArray) > 0 {
			for _, worker := range workers {
				// double check on len bc of concurrency
				if !worker.working && len(wordsArray) > 0 {
					worker.working = true
					go worker.read(wordsArray[0])
					wordsArray = wordsArray[1:]
				}
			}
		}
		if c.received == maxLen {
			break
		}
		select {
		case temp := <-c.wordsBox:
			for k, v := range temp {
				c.wordsMap[k] += v
			}
			c.received++
		}
	}
	return c.wordsMap
}
