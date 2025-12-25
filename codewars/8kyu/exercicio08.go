package main

import "fmt"

func BooleanToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func main() {
	fmt.Println(BooleanToString(true))
	fmt.Println(BooleanToString(false))
}
