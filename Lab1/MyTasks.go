package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordCount("Lubię bardzo placki"))
	fmt.Println(wordsCountMap("Lubię lubię placki placki a!"))
}

func wordCount(s string) int {
	/*
		Napisz funkcję która liczy ilość wyrazów w zdaniu
		(nie musisz usuwać żadnych znaków interpunkcyjnych)

	*/
	words := 0
	splitString := strings.Fields(s)
	for i := 0; i < len(splitString); i++ {
		words++
	}
	return words
}

/* trudniejszy wariant*/
func wordsCountMap(s string) map[string]int {
	/*
		Teraz stwórz funckję, która liczy wystąpienia każdego wyrazu w zdaniu (stringu)
		i zwraca je w formie mapy (słownika)
		np.:
		fmt.Println(wordsCountMap("Lubię lubię placki placki a!"))
		map[a:1 lubię:2 placki:2]
	*/
	fields := strings.Fields(s)
	wordsMap := make(map[string]int)
	for _, value := range fields {

		word := strings.ToLower(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(
						strings.ReplaceAll(
							value, ",", ""), "?", ""), "!", ""), ".", ""))

		wordsMap[word] += 1
	}
	return wordsMap
}
