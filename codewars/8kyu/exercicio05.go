package main

import (
	"fmt"
)

func StringRepeat(number int, word string) string {
	result := ""

	for i := 0; i < number; i++ {
		result = result + word
	}

	return result

}

func main() {
	fmt.Println(StringRepeat(3, "BTS"))
	fmt.Println(StringRepeat(5, "TATA"))
	fmt.Println(StringRepeat(0, "A"))
}
