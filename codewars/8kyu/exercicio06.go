package main

import "fmt"

func Opposite(number int) int {
	return number * -1
}

func main() {
	fmt.Println(Opposite(10))
	fmt.Println(Opposite(-524))
	fmt.Println(Opposite(0))
}
