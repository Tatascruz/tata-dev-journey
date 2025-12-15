package main

import "fmt"

func Multiply(a int, b int) int {
	return a * b
}

func main() {
	resultado := Multiply(25, 52)
	fmt.Println("A multiplicação é:", resultado)
}
