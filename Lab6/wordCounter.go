package main

import "fmt"

func main() {
	fmt.Printf("Test\n")
}

// Placeholder for map with counted words
type counterMap map[string]int

// Counter holds total record of words in file, and receives new information from Reader
type Counter struct {
	wordsMap counterMap
	wordsBox chan counterMap
}

func (c *Counter) NewCounter() Counter {
	wMap := map[string]int{}
	wChan := make(chan counterMap)
	return Counter{wMap, wChan}
}

func (c *Counter) MakeReader(line string) Reader {
	return Reader{line, c}
}

// Reader receives lines, counts word in it and sends it back to Counter
type Reader struct {
	line       string
	supervisor *Counter
}
