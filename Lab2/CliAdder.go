package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println(adder())
}

func adder() int {
	num1 := flag.Int("liczba1", 0, "Number1")
	num2 := flag.Int("liczba2", 0, "Number2")
	flag.Parse()
	intNum1 := *num1
	intNum2 := *num2
	if intNum1 == 0 && intNum2 == 0 {
		fmt.Println("Nie wykryto liczb.\nPodaj liczbę 1:")
		// fmt.Scanf("%d\n", &intNum1)    <- Problem with new line in windows
		fmt.Scanf("%d", &intNum1)
		fmt.Println("Podaj liczbę 2:")
		fmt.Scanf("%d", &intNum2)
	}
	result := intNum1 + intNum2
	return result
}
