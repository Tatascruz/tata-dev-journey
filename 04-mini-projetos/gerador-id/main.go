package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Gerador de ID único")

	id := uuid.New()

	fmt.Println("Seu novo ID é:")
	fmt.Println(id.String())
}
