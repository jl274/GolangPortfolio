package main

import "fmt"

func main() {
	fmt.Printf("Test\n")
}

// Placeholder for map with counted words
type counterMap map[string]int

// Counter holds total record of words in file, and receives new information from Readers
type Counter struct {
	wordsMap counterMap
	wordsBox chan counterMap
}

// Reader receives lines, counts word in it and sends it back to Counter
type Reader struct {
	line       string
	supervisor *Counter
}
