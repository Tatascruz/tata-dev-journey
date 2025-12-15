package main

import "fmt"

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	}

	return "odd"
}

func main() {
	fmt.Println(EvenOrOdd(87))
	fmt.Println(EvenOrOdd(36))
}
