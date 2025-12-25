package main

import "fmt"

func RemoveChar(s string) string {
	return s[1 : len(s)-1]
}

func main() {
	fmt.Println(RemoveChar("eloquent"))
	fmt.Println(RemoveChar("complete"))
	fmt.Println(RemoveChar("ab"))
}
