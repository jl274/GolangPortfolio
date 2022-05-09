package main

import "fmt"

func main() {
	fmt.Println("Add: ")
	fmt.Println(add(3, 2))
	fmt.Println()

	fmt.Println("Sum slice: ")
	floatSlice := []float64{33.3, 42.2, 0.999921}
	uintSlice := []uint{1, 2, 3, 4, 5}
	fmt.Println(sumSlice(floatSlice))
	fmt.Println(sumSlice(uintSlice))
	fmt.Println()

	fmt.Println("Print map: ")
	runeIntMap := map[rune]int{'a': 3, 'b': 20, 'c': -1}
	floatStringMap := map[float64]string{0.11: "Very small", 5.4: "Pretty average", 80.0: "Holy guacamoly!"}
	printMap(runeIntMap)
	printMap(floatStringMap)
	fmt.Println()

	fmt.Println("Does slice contains: ")
	fmt.Println(doesSliceContains(uintSlice, 2))
	fmt.Println(doesSliceContains(floatSlice, 11.11))
	fmt.Println()

	fmt.Println("Remove from slice: ")
	fmt.Println(removeElemFromSlice(uintSlice, 2))
	fmt.Println()

	fmt.Println("Print in box: ")
	printInBox(322)
	printInBox("Beautiful weather")
	fmt.Println()

	fmt.Println("Print point: ")
	firstPoint := point2D{2, 10}
	secondPoint := point3D{point2D{3, 3}, 3}
	printPoint(firstPoint)
	printPoint(secondPoint)
	fmt.Println()
}

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
}

func add[T Number](a, b T) T {
	return a + b
}

func sumSlice[N Number](list []N) N {
	var result N
	for _, val := range list {
		result += val
	}
	return result
}

func printMap[T1 Number | string | rune, T2 Number | string](inputMap map[T1]T2) {
	i := 1
	for key, value := range inputMap {
		fmt.Printf("%v. Key - %v: value = %v\n", i, key, value)
		i++
	}
}

func doesSliceContains[TYPE int | uint | string | float64](slice []TYPE, elem TYPE) bool {
	for _, sliceElem := range slice {
		if sliceElem == elem {
			return true
		}
	}
	return false
}

func removeElemFromSlice[T Number | string](inputSlice []T, delete T) []T {
	newSlice := make([]T, 0, len(inputSlice))
	for _, elem := range inputSlice {
		if elem != delete {
			newSlice = append(newSlice, elem)
		}
	}
	return newSlice
}

func printInBox[T any](element T) {
	stringElement := fmt.Sprint(element)
	for i := 0; i < len(stringElement)+4; i++ {
		fmt.Printf("*")
	}
	fmt.Println()
	fmt.Printf("* %v *\n", element)
	for i := 0; i < len(stringElement)+4; i++ {
		fmt.Printf("*")
	}
	fmt.Println()
}

type geometryPoint interface {
	getDimensions() string
	dimension() string
}

type point2D struct {
	x int
	y int
}

func (p point2D) getDimensions() string {
	return fmt.Sprintf("%d x %d", p.x, p.y)
}

func (p point2D) dimension() string {
	return "2"
}

type point3D struct {
	point2D
	z int
}

func (p point3D) getDimensions() string {
	return fmt.Sprintf("%d x %d x %d", p.x, p.y, p.z)
}

func (p point3D) dimension() string {
	return "3"
}

func printPoint[T geometryPoint](point T) {
	fmt.Printf("Point %sD: %s\n", point.dimension(), point.getDimensions())
}
