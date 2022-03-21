package main

import (
	"flag"
	"fmt"
	"os"
)

var legalUnits []string = []string{"K", "C", "F"}

func main() {
	optionsNumber := len(os.Args)
	if optionsNumber < 3 {
		panic(any("Missing option(s)"))
	} else if optionsNumber > 7 {
		panic(any("To many options"))
	}
	fromFlag := flag.String("from", "C", "The unit you want to convert from (K/C/F)")
	toFlag := flag.String("to", "F", "The unit you want to convert to (K/C/F)")
	valueFlag := flag.Float64("value", 0, "Value to be converted")
	flag.Parse()
	value := *valueFlag
	from := *fromFlag
	to := *toFlag
	fmt.Println(from, to, value)

	if !contains(legalUnits, from) || !contains(legalUnits, to) {
		panic(any("Invalid unit(s)"))
	}

	if from == to {
		fmt.Printf("###\n%4.4f°%s is %4.4f°%s\n###\n", value, from, value, to)
	}

	switch from {

	case "C":
		switch to {

		case "F":
			fmt.Printf("###\n%.4f°C is %.4f°F\n###\n", value, cf(value))

		case "K":
			fmt.Printf("###\n%.4f°C is %.4fK\n###\n", value, ck(value))
		}

	case "F":
		switch to {

		case "C":
			fmt.Printf("###\n%.4f°F is %.4f°C\n###\n", value, fc(value))

		case "K":
			fmt.Printf("###\n%.4f°F is %.4fK\n###\n", value, fk(value))
		}

	case "K":
		switch to {

		case "C":
			fmt.Printf("###\n%.4fK is %.4f°C\n###\n", value, kc(value))

		case "F":
			fmt.Printf("###\n%.4fK is %.4f°F\n###\n", value, kf(value))
		}

	default:
		fmt.Println("Invalid conversion")
	}

}

func cf(value float64) float64 {
	return (1.8 * value) + 32
}

func fc(value float64) float64 {
	return (value - 32) / 1.8
}

func ck(value float64) float64 {
	return value + 273.15
}

func kc(value float64) float64 {
	return value - 273.15
}

func fk(value float64) float64 {
	return (value + 459.67) * 5 / 9
}

func kf(value float64) float64 {
	return cf(kc(value))
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
