package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	witcherCounter := NewCounter()
	result := witcherCounter.count("text.txt")
	//for k, v := range result {
	//	fmt.Printf("%s: %d times\n", k, v)
	//}
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
func (c *Counter) count(file string) counterMap {
	wordsArray := c.readFile(file)
	maxLen := len(wordsArray)
	workers := make([]Reader, 0, maxLen)
	for i := 0; i < maxLen; i++ {
		workers = append(workers, c.MakeReader())
	}
	for {
		if len(wordsArray) > 0 {
			for _, worker := range workers {
				if !worker.working {
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
