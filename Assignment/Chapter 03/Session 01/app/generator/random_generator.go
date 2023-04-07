package generator

import (
	"math/rand"
	"time"
)

func RandomNumber(max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	time.Sleep(2 * time.Second)
	return (r.Intn(max) + 1)
}