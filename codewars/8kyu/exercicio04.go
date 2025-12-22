package main

import (
	"fmt"
	"strconv"
)

func NumberToString(number int) string {
	return strconv.Itoa(number)
}

func main() {
	fmt.Println(NumberToString(456))
	fmt.Println(NumberToString(-2535))
	fmt.Println(NumberToString(0))
}
