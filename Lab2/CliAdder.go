package main

/*
Należy za pomocą programu dodać dwie liczby podane jako argumenty wiersza poleceń (CLI)opcje będą takie:

-liczba1 N-liczba2 Mw przypadku gdy nie poda się opcji,
program powinien umożliwić wczytanie liczb z klawiatury
Przykładowe uruchomienie:go run program.go -liczba1 5 -liczba2 712

*/

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
	intNum1 := *num1
	intNum2 := *num2
	flag.Parse()
	if intNum1 == 0 && intNum2 == 0 {
		fmt.Println("Nie wykryto liczb.\n Podaj liczbę 1:")
		fmt.Scanf("%d", &intNum1)
		fmt.Println("Podaj liczbę 2:")
		fmt.Scanf("%d", &intNum2)
	}
	result := intNum1 + intNum2
	return result
}
