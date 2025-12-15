package main

import "fmt"

func ReturnNegative(number int) int {
	if number > 0 {
		return -number
	}
	return number
}

func main() {
	fmt.Println(ReturnNegative(14))
	fmt.Println(ReturnNegative(-33))
	fmt.Println(ReturnNegative(0))
}
