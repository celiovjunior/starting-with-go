package main

import "fmt"

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func writeMessage(text string, years ...int) {
	for _, year := range years {
		fmt.Println(text, year)
	}
}

func main() {
	// totalSum := sum(1, 2, 3, 4)
	// fmt.Println(totalSum)

	writeMessage("oie", 2012, 2013, 2014)
}
