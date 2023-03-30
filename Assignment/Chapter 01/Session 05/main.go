package main

import (
	"fmt"
	r "ojam-test/fifth-session/random"
	s "ojam-test/fifth-session/straight"
)

var coba = []string{"coba1", "coba2", "coba3"}
var bisa = []string{"bisa1", "bisa2", "bisa3"}

func main() {
	fmt.Println("RANDOM THREAD")
	r.RandomThread(coba, bisa)
	
	fmt.Println()
	
	fmt.Println("STRAIGHT THREAD")
	s.StraightThread(coba, bisa)
}
