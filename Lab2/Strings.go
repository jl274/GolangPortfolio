package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("---Contains char example---")
	fmt.Println(stringContainsChar("pancake", "c"))
	fmt.Println(stringContainsChar("pancake", "x"))

	fmt.Println("---Palindrome Example---")
	fmt.Println(isPalindrome("kamilślimak"))
	fmt.Println(isPalindrome("abfgba"))

	fmt.Println("---Anagram example---")
	fmt.Println(isAnagram("i am lord voldemort", "tom marvolo riddle"))
	fmt.Println(strings.Compare("dff", "cdd"))

	fmt.Println("---Sorting string example---")
	fmt.Println(sortStringAlphabetically("city,brake,is,awesome,indeed"))
}

func stringContainsChar(s string, key string) bool {
	charList := strings.Split(s, "")
	for _, char := range charList {
		if char == key {
			return true
		}
	}
	return false
}

func isPalindrome(s string) bool {

	charList := strings.Split(s, "")
	stringsLength := len(charList)

	for i := 0; i < stringsLength/2; i++ {
		if charList[i] != charList[stringsLength-i-1] {
			return false
		}
	}

	return true
}

func isAnagram(s1, s2 string) bool {

	s1 = strings.ReplaceAll(s1, " ", "")
	s2 = strings.ReplaceAll(s2, " ", "")

	if len(s1) != len(s2) {
		return false
	}

	charListS1 := strings.Split(s1, "")
	for _, char := range charListS1 {
		//fmt.Printf("%s : %d == %d", char, strings.Count(s1, char), strings.Count(s2, char))
		if strings.Count(s1, char) != strings.Count(s2, char) {
			return false
		}
	}

	return true
}

func sortStringAlphabetically(s string) string {
	// sorts string with words separated by commas
	stringsList := strings.Split(s, ",")

	sortedList := make([]string, 0)
	sortedList = append(sortedList, stringsList[0])

	for i := 1; i < len(stringsList); i++ {
		curr := stringsList[i]
		for i2, prevValue := range sortedList {
			if curr < prevValue {
				curr, sortedList[i2] = prevValue, curr
			}
			if i2 == len(sortedList)-1 {
				sortedList = append(sortedList, curr)
			}
		}
	}

	returnString := ""
	for index, val := range sortedList {
		if index == 0 {
			returnString = val
		} else {
			returnString += "," + val
		}
	}
	return returnString
}
