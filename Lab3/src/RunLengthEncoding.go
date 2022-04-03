package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	encodeFlag := flag.Bool("encode", false, "Present if wants encoding")
	decodeFlag := flag.Bool("decode", false, "Present if wants decoding")
	stringFlag := flag.String("value", "", "String to be be encoded/decoded")
	flag.Parse()
	encode := *encodeFlag
	decode := *decodeFlag
	stringValue := *stringFlag
	if encode && decode {
		panic(any("Both flags can't be present"))
	} else if stringValue == "" {
		panic(any("Empty string"))
	}
	if encode {
		fmt.Println(RLEncode(stringValue))
	} else if decode {
		fmt.Println(RLDecode(stringValue))
	}
	//fmt.Println(RLEncode("AABCCCDEEEE"))
	//fmt.Println(RLDecode("2AB3CD4E"))
}

func RLEncode(s string) string {
	s = strings.ToUpper(s)
	returnString := ""
	var prev string
	count := 1
	for i, letter := range strings.Split(s, "") {
		if i == 1 {
			prev = letter
			count++
		} else if letter == prev {
			count++
		} else {
			if count == 1 {
				returnString = returnString + prev
			} else {
				returnString = returnString + strconv.Itoa(count) + prev
			}
			count = 1
		}
		prev = letter
	}
	if count == 1 {
		returnString = returnString + prev
	} else {
		returnString = returnString + strconv.Itoa(count) + prev
	}
	return returnString
}

func RLDecode(s string) string {
	decodedString := ""
	count := 1
	for _, letter := range strings.Split(s, "") {
		if count != 1 {
			for i := 1; i <= count; i++ {
				decodedString = decodedString + letter
			}
			count = 1
		} else if !isNumber(letter) {
			decodedString = decodedString + letter
		} else {
			count, _ = strconv.Atoi(letter)
		}
	}
	return decodedString
}

func isNumber(s string) bool {
	match, _ := regexp.MatchString("^[0-9]$", s)
	return match
}
