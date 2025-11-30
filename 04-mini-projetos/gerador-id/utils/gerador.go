package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GeradorID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d-%d", time.Now().Unix(), rand.Intn(99999))
}
