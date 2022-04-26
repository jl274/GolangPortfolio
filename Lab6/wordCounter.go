package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Printf("Test\n")
}

// Placeholder for map with counted words
type counterMap map[string]int

// Counter holds total record of words in file, and receives new information from Reader
type Counter struct {
	wordsMap counterMap
	received int
	wordsBox chan counterMap
}

func (c *Counter) NewCounter() Counter {
	wMap := map[string]int{}
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
	// working
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
	workers := make([]Reader, 0, len(wordsArray))
	for i := 0; i < len(wordsArray); i++ {
		workers = append(workers, c.MakeReader())
	}
	for {
		for _, worker := range workers {
			if !worker.working {
				worker.working = true
				worker.read(wordsArray[0])
				wordsArray = wordsArray[1:]
			}
		}
		if c.received == len(wordsArray) {
			break
		}
		select {
		case temp := <-c.wordsBox:
			for k, v := range temp {
				count := c.wordsMap[k] + v
				c.wordsMap[k] = count
			}
			c.received++
		}
	}
	return c.wordsMap
}
